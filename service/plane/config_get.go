package plane

import (
	"bytes"
	"context"

	"github.com/hyperledger/fabric-config/protolator"
	pb "gitlab.n-t.io/core/library/hlf-tool/hlf-control-plane/proto"
	"gitlab.n-t.io/core/library/hlf-tool/hlf-control-plane/system/cscc"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *srv) ConfigGet(ctx context.Context, req *pb.ConfigGetRequest) (*pb.ConfigGetResponse, error) {
	logger := s.logger.With(zap.String("channel", req.ChannelName))
	logger.Debug("get channel config", zap.String("channel", req.ChannelName))
	endCli, err := s.peerPool.GetRandomEndorser(ctx, s.mspID)
	if err != nil {
		logger.Error("get endorser failed", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "get endorser failed: %v", err)
	}
	conf, err := cscc.NewClient(endCli, s.id).GetChannelConfig(ctx, req.ChannelName)
	if err != nil {
		logger.Error("get config failed", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "get channel config: %v", err)
	}

	b := new(bytes.Buffer)
	if err = protolator.DeepMarshalJSON(b, conf); err != nil {
		return nil, err
	}

	return &pb.ConfigGetResponse{Config: b.Bytes()}, nil
}
