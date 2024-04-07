package plane

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/hyperledger/fabric/integration/channelparticipation"
	pb "gitlab.n-t.io/core/library/hlf-tool/hlf-control-plane/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *srv) OrderingListInfo(ctx context.Context, req *pb.OrderingRequest) (*pb.OrderingListInfoResponse, error) { //nolint:funlen
	var (
		logger *zap.Logger
		url    string
	)

	if len(req.ChannelName) == 0 {
		logger = s.logger.With(zap.String("channel", "no channel"))
		url = fmt.Sprintf(httpRequestMask, s.getURL(req.Host, req.Port))
	} else {
		logger = s.logger.With(zap.String("channel", req.ChannelName))
		url = fmt.Sprintf(httpRequestMask+"/%s", s.getURL(req.Host, req.Port), req.ChannelName)
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		logger.Error("get list http request", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "get list http request: %v", err)
	}

	resp, err := s.processRequest(httpReq)
	if err != nil {
		logger.Error("process request", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "process request: %v", err)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error("read all", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "read all: %v", err)
	}

	err = resp.Body.Close()
	if err != nil {
		logger.Error("body close", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "body close: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		logger.Error("unexpected status", zap.Int("status", resp.StatusCode), zap.Error(err))
		return nil, status.Errorf(codes.Internal, "unexpected status: %d", resp.StatusCode)
	}

	out := &pb.OrderingListInfoResponse{}

	if len(req.ChannelName) == 0 {
		cl := &channelparticipation.ChannelList{}
		err = json.Unmarshal(bodyBytes, cl)
		if err != nil {
			logger.Error("unmarshal json", zap.Error(err))
			return nil, status.Errorf(codes.Internal, "unmarshal json: %v", err)
		}
		for _, c := range cl.Channels {
			out.Info = append(out.Info, &pb.OrderingListInfoResponse_ChannelInfo{
				Name: c.Name,
			})
		}
	} else {
		c := &channelparticipation.ChannelInfo{}
		err = json.Unmarshal(bodyBytes, c)
		if err != nil {
			logger.Error("unmarshal json", zap.Error(err))
			return nil, status.Errorf(codes.Internal, "unmarshal json: %v", err)
		}
		out.Info = append(out.Info, &pb.OrderingListInfoResponse_ChannelInfo{
			Name:              c.Name,
			Status:            c.Status,
			ConsensusRelation: c.ConsensusRelation,
			Height:            c.Height,
		})
	}

	return out, nil
}
