package plane

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/atomyze-foundation/hlf-control-plane/pkg/configtx"
	pb "github.com/atomyze-foundation/hlf-control-plane/proto"
	"github.com/hyperledger/fabric/orderer/common/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *srv) ChannelCreate(
	ctx context.Context,
	req *pb.ChannelCreateRequest,
) (*pb.ChannelCreateResponse, error) {
	if req.ChannelName == "" {
		return nil, status.Errorf(codes.Internal, "missing channel name, please specify it")
	}

	if len(req.Organizations) == 0 {
		return nil, status.Errorf(codes.Internal, "missing organization struct, please specify it")
	}

	blockBytes, err := configtx.CreateGenesisBlock(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "create genesis block, %v", err)
	}

	resp := &pb.ChannelCreateResponse{}

	for _, org := range req.GetOrganizations() {

		if org.Id != s.mspID {
			continue
		}

		for _, o := range org.GetOrderers() {
			err = s.joinOrdererAfterCreateChannel(ctx, o.Host, o.AdminPort, blockBytes)
			if err != nil {
				resp.Result = append(resp.Result, &pb.ChannelCreateResponse_OrdererResult{
					Host: o.Host,
					Port: o.Port,
				})

				continue
			}

			resp.Result = append(resp.Result, &pb.ChannelCreateResponse_OrdererResult{
				Host:   o.Host,
				Port:   o.Port,
				Joined: true,
			})
		}

		break
	}

	return resp, nil
}

func (s *srv) joinOrdererAfterCreateChannel(ctx context.Context, host string, port uint32, blockBytes []byte) error {
	httpReq, err := s.createJoinRequest(ctx, fmt.Sprintf(httpRequestMask, s.getURL(host, port)), blockBytes)
	if err != nil {
		return fmt.Errorf("get join http request: %w", err)
	}

	resp, err := s.processRequest(httpReq)
	if err != nil {
		return fmt.Errorf("process request: %w", err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusCreated {
		decoder := json.NewDecoder(resp.Body)
		respErr := &types.ErrorResponse{}
		err = decoder.Decode(respErr)
		if err != nil {
			return fmt.Errorf("process request: %w", err)
		}

		if !strings.Contains(respErr.Error, "already exist") {
			return fmt.Errorf("unexpected status: %d", resp.StatusCode)
		}
	}

	return nil
}
