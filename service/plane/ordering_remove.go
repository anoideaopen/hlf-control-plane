package plane

import (
	"context"
	"fmt"
	"net/http"

	"github.com/atomyze-foundation/hlf-control-plane/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *srv) OrderingRemove(ctx context.Context, req *proto.OrderingRequest) (*proto.OrderingRemoveResponse, error) {
	logger := s.logger.With(zap.String("channel", req.ChannelName))

	url := fmt.Sprintf(httpRequestMask+"/%s", s.getURL(req.Host, req.Port), req.ChannelName)

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodDelete, url, nil)
	if err != nil {
		logger.Error("get remove http request", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "get remove http request: %v", err)
	}

	resp, err := s.processRequest(httpReq)
	if err != nil {
		logger.Error("process request", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "process request: %v", err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusNoContent {
		logger.Error("unexpected status", zap.Int("status", resp.StatusCode), zap.Error(err))
		return nil, status.Errorf(codes.Internal, "unexpected status: %d", resp.StatusCode)
	}

	logger.Debug("remove orderer from channel", zap.String("channel", req.ChannelName))
	return &proto.OrderingRemoveResponse{}, nil
}
