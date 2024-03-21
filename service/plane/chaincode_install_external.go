package plane

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/atomyze-foundation/hlf-control-plane/proto"
	"github.com/hashicorp/go-multierror"
	"github.com/hyperledger/fabric/core/chaincode/persistence"
	"github.com/hyperledger/fabric/core/container/externalbuilder"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

const (
	ccFilePerm          = 0o600
	externalDefaultType = "ccaas"
	defaultDuration     = 5 * 1000 * 1000 * 1000

	tmplTLSCert     = `{{ .tlsCert }}`
	tmplTLSKey      = `{{ .tlsKey }}`
	tmplTLSRootCert = `{{ .tlsRootCert }}`
)

func (s *srv) ChaincodeInstallExternal(ctx context.Context, req *proto.ChaincodeInstallExternalRequest) (*proto.ChaincodeInstallResponse, error) {
	if req.Type == "" {
		req.Type = externalDefaultType
	}

	ccPkg, err := s.createExternalCCPackage(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "create external cc: %v", err)
	}

	pkgID := persistence.PackageID(req.Label, ccPkg)

	var wg sync.WaitGroup
	resChan := make(chan *proto.ChaincodeInstallResponse_Result)

	for _, p := range s.localPeers {
		p := p
		wg.Add(1)
		go func() {
			defer wg.Done()
			s.logger.Debug("processing package on peer", zap.String("label", req.Label), zap.String("type", req.Type), zap.Int("size", len(ccPkg)))
			endCli, err := s.peerPool.GetEndorser(ctx, p)
			if err != nil {
				resChan <- &proto.ChaincodeInstallResponse_Result{
					Peer:   p.String(),
					Result: &proto.ChaincodeInstallResponse_Result_Err{Err: fmt.Sprintf("get endorser: %s", err)},
				}
				return
			}
			s.processPeerInstall(ctx, req.Label, ccPkg, p.String(), endCli, resChan)
		}()
	}

	result := make([]*proto.ChaincodeInstallResponse_Result, 0)
	go func() {
		wg.Wait()
		close(resChan)
	}()

	for res := range resChan {
		result = append(result, res)
	}

	return &proto.ChaincodeInstallResponse{PackageId: pkgID, Result: result}, nil
}

func (s *srv) createExternalCCPackage(ctx context.Context, req *proto.ChaincodeInstallExternalRequest) ([]byte, error) {
	codePkg, err := s.createCodePackage(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("create code package: %w", err)
	}

	ccPkg, err := s.createChaincodePackage(codePkg, req.Type, req.Label)
	if err != nil {
		return nil, fmt.Errorf("create cc: %w", err)
	}
	return ccPkg, nil
}

func (s *srv) createCodePackage(ctx context.Context, req *proto.ChaincodeInstallExternalRequest) (codeBytes []byte, err error) { //nolint:funlen
	var (
		buf   = new(bytes.Buffer)
		gzWr  = gzip.NewWriter(buf)
		tarWr = tar.NewWriter(gzWr)
	)

	defer func() {
		codeBytes = buf.Bytes()
	}()

	defer func() {
		if gzErr := gzWr.Close(); err != nil {
			err = multierror.Append(fmt.Errorf("gzip close: %w", gzErr), err)
			return
		}
	}()

	defer func() {
		if tarErr := tarWr.Close(); err != nil {
			err = multierror.Append(fmt.Errorf("tar close: %w", tarErr), err)
			return
		}
	}()

	var connData externalbuilder.ChaincodeServerUserData
	connData.TLSRequired = req.TlsRequired
	connData.Address = req.Address
	var duration externalbuilder.Duration
	if err = duration.UnmarshalJSON([]byte(req.Timeout)); err != nil {
		duration = defaultDuration
	}
	if req.TlsRequired {
		connData.ClientAuthRequired = req.TlsClientAuth

		connData.ClientCert = tmplTLSCert
		connData.ClientKey = tmplTLSKey
		connData.RootCert = tmplTLSRootCert
	}

	// check connection to chaincode with presented crypto materials
	if req.EnableConnCheck {
		tlsCert, err := tls.X509KeyPair([]byte(connData.ClientCert), []byte(connData.ClientKey))
		if err != nil {
			return nil, fmt.Errorf("failed to create tls credentials: %w", err)
		}

		certPool := x509.NewCertPool()
		if ok := certPool.AppendCertsFromPEM([]byte(connData.RootCert)); ok {
			return nil, fmt.Errorf("append cert to pool failed")
		}

		config := &tls.Config{
			Certificates: []tls.Certificate{tlsCert},
			ClientCAs:    certPool,
		}

		ctx, cancel := context.WithTimeout(ctx, defaultDuration)
		defer cancel()

		conn, err := grpc.DialContext(ctx, connData.Address, grpc.WithBlock(), grpc.WithTransportCredentials(credentials.NewTLS(config)))
		if err != nil {
			return nil, fmt.Errorf("failed to init grpc connection: %w", err)
		}
		defer func() {
			_ = conn.Close()
		}()
	}

	connDataBytes, err := json.Marshal(connData)
	if err != nil {
		return nil, fmt.Errorf("marshal connection data: %w", err)
	}

	if err = tarWr.WriteHeader(&tar.Header{
		Name: "connection.json",
		Mode: ccFilePerm,
		Size: int64(len(connDataBytes)),
	}); err != nil {
		err = fmt.Errorf("write header for connection.json: %w", err)
		return
	}

	if _, err = tarWr.Write(connDataBytes); err != nil {
		err = fmt.Errorf("write connection.json bytes: %w", err)
		return
	}

	return
}

func (s *srv) createChaincodePackage(codePkg []byte, ccType, ccLabel string) (pkgBytes []byte, err error) {
	var (
		buf   = new(bytes.Buffer)
		gzWr  = gzip.NewWriter(buf)
		tarWr = tar.NewWriter(gzWr)
	)

	defer func() {
		pkgBytes = buf.Bytes()
	}()

	defer func() {
		if gzErr := gzWr.Close(); err != nil {
			err = multierror.Append(fmt.Errorf("gzip close: %w", gzErr), err)
			return
		}
	}()

	defer func() {
		if tarErr := tarWr.Close(); err != nil {
			err = multierror.Append(fmt.Errorf("tar close: %w", tarErr), err)
			return
		}
	}()

	mdBytes, err := json.Marshal(persistence.ChaincodePackageMetadata{
		Type:  ccType,
		Label: ccLabel,
	})
	if err != nil {
		err = fmt.Errorf("marshal metadata: %w", err)
		return
	}

	if err = tarWr.WriteHeader(&tar.Header{
		Name: persistence.CodePackageFile,
		Mode: ccFilePerm,
		Size: int64(len(codePkg)),
	}); err != nil {
		err = fmt.Errorf("tar write header for code package: %w", err)
		return
	}
	if _, err = tarWr.Write(codePkg); err != nil {
		err = fmt.Errorf("tar write code package: %w", err)
		return
	}

	if err = tarWr.WriteHeader(&tar.Header{
		Name: persistence.MetadataFile,
		Mode: ccFilePerm,
		Size: int64(len(mdBytes)),
	}); err != nil {
		err = fmt.Errorf("tar write header for metadata: %w", err)
		return
	}
	if _, err = tarWr.Write(mdBytes); err != nil {
		err = fmt.Errorf("tar write metadata: %w", err)
		return
	}

	return
}
