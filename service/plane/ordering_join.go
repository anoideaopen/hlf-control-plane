package plane

import (
	"bytes"
	"context"
	"fmt"
	"mime/multipart"
	"net"
	"net/http"
	"strconv"

	"github.com/atomyze-foundation/hlf-control-plane/pkg/orderer"
	"github.com/atomyze-foundation/hlf-control-plane/pkg/util"
	"github.com/atomyze-foundation/hlf-control-plane/proto"
	"github.com/atomyze-foundation/hlf-control-plane/system/cscc"
	pb "github.com/golang/protobuf/proto" //nolint:staticcheck
	"github.com/hyperledger/fabric-protos-go/common"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const httpRequestMask = "%s/participation/v1/channels"

func (s *srv) OrderingJoin(ctx context.Context, req *proto.OrderingRequest) (*proto.OrderingJoinResponse, error) {
	logger := s.logger.With(zap.String("channel", req.ChannelName))
	logger.Debug("get channel config", zap.String("channel", req.ChannelName))

	var (
		block *common.Block
		err   error
	)

	if len(req.JoinedOrderer) != 0 {
		for _, jord := range req.JoinedOrderer {
			block, err = s.getGenesisBlockFromOrderer(ctx, jord.Host, jord.Port, req.ChannelName)
			if err != nil {
				logger.Error("get block from orderer", zap.String("error", err.Error()))
				continue
			}

			break
		}
	}

	if block == nil {
		endCli, err := s.peerPool.GetRandomEndorser(ctx, s.mspID)
		if err != nil {
			logger.Error("get endorser failed", zap.Error(err))
			return nil, status.Errorf(codes.Internal, "get endorser failed: %v", err)
		}
		block, err = cscc.NewClient(endCli, s.id).GetChannelConfigBlock(ctx, req.ChannelName)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "get config block: %v", err)
		}
	}

	blockBytes, err := pb.Marshal(block)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "marshal block: %v", err)
	}

	httpReq, err := s.createJoinRequest(ctx, fmt.Sprintf(httpRequestMask, s.getURL(req.Host, req.Port)), blockBytes)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "get join http request: %v", err)
	}

	resp, err := s.processRequest(httpReq)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "process request: %v", err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusCreated {
		return nil, status.Errorf(codes.Internal, "unexpected status: %d", resp.StatusCode)
	}
	return &proto.OrderingJoinResponse{Exists: false}, nil
}

func (s *srv) processRequest(req *http.Request) (*http.Response, error) {
	cli := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: s.tlsCli,
		},
	}
	return cli.Do(req)
}

func (s *srv) createJoinRequest(ctx context.Context, url string, blockBytes []byte) (*http.Request, error) {
	joinBody := new(bytes.Buffer)
	writer := multipart.NewWriter(joinBody)
	part, err := writer.CreateFormFile("config-block", "config.block")
	if err != nil {
		return nil, err
	}
	if _, err = part.Write(blockBytes); err != nil {
		return nil, err
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, joinBody)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	return req, nil
}

func (s *srv) getURL(host string, port uint32) string {
	return "https://" + net.JoinHostPort(host, strconv.Itoa(int(port)))
}

func (s *srv) getGenesisBlockFromOrderer(ctx context.Context, host string, port uint32, channelName string) (*common.Block, error) {
	ordCli, err := s.ordPool.Get(&orderer.Orderer{
		Host: host,
		Port: port,
	})
	if err != nil {
		s.logger.Error("get orderer", zap.String("error", err.Error()))
		return nil, fmt.Errorf("get orderer: %w", err)
	}

	// get envelope for orderer deliver blocks
	env, err := util.GetSeekOldestEnvelope(channelName, s.id)
	if err != nil {
		s.logger.Error("get seek envelope", zap.String("error", err.Error()))
		return nil, fmt.Errorf("get seek envelope: %w", err)
	}

	// get deliver client and send envelope
	deliverCli, err := ordCli.Deliver(ctx)
	if err != nil {
		s.logger.Error("get orderer deliver client", zap.String("error", err.Error()))
		return nil, fmt.Errorf("get orderer deliver client: %w", err)
	}

	defer func() {
		if err = deliverCli.CloseSend(); err != nil {
			s.logger.Error("close deliver stream error", zap.String("orderer", host), zap.Error(err))
		}
	}()

	if err = deliverCli.Send(env); err != nil {
		s.logger.Error("deliver send", zap.String("error", err.Error()))
		return nil, fmt.Errorf("deliver send: %w", err)
	}

	// get block from deliver stream
	block, err := util.GetBlockFromDeliverClient(deliverCli)
	if err != nil {
		s.logger.Error("get block", zap.String("error", err.Error()))
		return nil, fmt.Errorf("get block: %w", err)
	}

	if block == nil {
		s.logger.Error("deliver send", zap.String("error", "block is nil"))
		return nil, fmt.Errorf("get block: error block is nil")
	}

	return block, nil
}

func (s *srv) getLastBlockFromOrderer(ctx context.Context, host string, port uint32, certs [][]byte, channelName string) (*common.Block, error) {
	ordCli, err := s.ordPool.Get(&orderer.Orderer{
		Host:         host,
		Port:         port,
		Certificates: certs,
	})
	if err != nil {
		s.logger.Error("get orderer", zap.String("error", err.Error()))
		return nil, fmt.Errorf("get orderer: %w", err)
	}

	// get envelope for orderer deliver blocks
	env, err := util.GetSeekNewestEnvelopeOnlyOne(channelName, s.id, nil)
	if err != nil {
		s.logger.Error("get seek envelope", zap.String("error", err.Error()))
		return nil, fmt.Errorf("get seek envelope: %w", err)
	}

	// get deliver client and send envelope
	deliverCli, err := ordCli.Deliver(ctx)
	if err != nil {
		s.logger.Error("get orderer deliver client", zap.String("error", err.Error()))
		return nil, fmt.Errorf("get orderer deliver client: %w", err)
	}

	defer func() {
		if err = deliverCli.CloseSend(); err != nil {
			s.logger.Error("close deliver stream error", zap.String("orderer", host), zap.Error(err))
		}
	}()

	if err = deliverCli.Send(env); err != nil {
		s.logger.Error("deliver send", zap.String("error", err.Error()))
		return nil, fmt.Errorf("deliver send: %w", err)
	}

	// get block from deliver stream
	block, err := util.GetBlockFromDeliverClient(deliverCli)
	if err != nil {
		s.logger.Error("get block", zap.String("error", err.Error()))
		return nil, fmt.Errorf("get block: %w", err)
	}

	if block == nil {
		s.logger.Error("deliver send", zap.String("error", "block is nil"))
		return nil, fmt.Errorf("get block: error block is nil")
	}

	return block, nil
}

func (s *srv) getBlocksFromNumberFromOrderer(ctx context.Context, host string, port uint32, certs [][]byte, channelName string, number uint64) (*common.Block, error) {
	ordCli, err := s.ordPool.Get(&orderer.Orderer{
		Host:         host,
		Port:         port,
		Certificates: certs,
	})
	if err != nil {
		s.logger.Error("get orderer", zap.String("error", err.Error()))
		return nil, fmt.Errorf("get orderer: %w", err)
	}

	// get envelope for orderer deliver blocks
	env, err := util.GetSeekSpecifiedEnvelope(channelName, s.id, number, nil)
	if err != nil {
		s.logger.Error("get seek envelope", zap.String("error", err.Error()))
		return nil, fmt.Errorf("get seek envelope: %w", err)
	}

	// get deliver client and send envelope
	deliverCli, err := ordCli.Deliver(ctx)
	if err != nil {
		s.logger.Error("get orderer deliver client", zap.String("error", err.Error()))
		return nil, fmt.Errorf("get orderer deliver client: %w", err)
	}

	defer func() {
		if err = deliverCli.CloseSend(); err != nil {
			s.logger.Error("close deliver stream error", zap.String("orderer", host), zap.Error(err))
		}
	}()

	if err = deliverCli.Send(env); err != nil {
		s.logger.Error("deliver send", zap.String("error", err.Error()))
		return nil, fmt.Errorf("deliver send: %w", err)
	}

	// get block from deliver stream
	block, err := util.GetBlockFromDeliverClient(deliverCli)
	if err != nil {
		s.logger.Error("get block", zap.String("error", err.Error()))
		return nil, fmt.Errorf("get block: %w", err)
	}

	if block == nil {
		s.logger.Error("deliver send", zap.String("error", "block is nil"))
		return nil, fmt.Errorf("get block: error block is nil")
	}

	return block, nil
}
