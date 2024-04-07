package plane

import (
	"context"

	"gitlab.n-t.io/core/library/hlf-tool/hlf-control-plane/pkg/util"
	pb "gitlab.n-t.io/core/library/hlf-tool/hlf-control-plane/proto"
	"gitlab.n-t.io/core/library/hlf-tool/hlf-control-plane/system/cscc"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *srv) ConfigOrderingDelete(ctx context.Context, req *pb.ConfigOrderingDeleteRequest) (*pb.ConfigOrderingDeleteResponse, error) {
	logger := s.logger.With(zap.String("channel", req.ChannelName))

	logger.Debug("get channel config", zap.String("channel", req.ChannelName))
	endCli, err := s.peerPool.GetRandomEndorser(ctx, s.mspID)
	if err != nil {
		logger.Error("get endorser failed", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "get endorser failed: %v", err)
	}

	config, err := cscc.NewClient(endCli, s.id).GetChannelConfig(ctx, req.ChannelName)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "get config: %v", err)
	}

	s.logger.Debug("get orderer config")
	orderers, consType, err := util.GetOrdererConfig(config)
	if err != nil {
		s.logger.Error("get orderer config failed", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "get orderer config: %v", err)
	}
	s.logger.Debug("got orderer config", zap.Int("orderers", len(orderers)), zap.String("consensus", consType.String()))

	newOrd := make([]*pb.Orderer, 0)
	var editOrd *pb.Orderer
	for _, ord := range orderers {
		if ord.Host == req.Orderer.Host && ord.Port == req.Orderer.Port ||
			ord.ConsenterId == req.Orderer.ConsenterId {
			editOrd = ord
			continue
		}
		newOrd = append(newOrd, ord)
	}

	if err = s.proceedOrderingConsenterUpdate(ctx, req.ChannelName, config, newOrd, editOrd); err != nil {
		return nil, status.Errorf(codes.Internal, "proceed update: %v", err)
	}
	return &pb.ConfigOrderingDeleteResponse{}, nil
}
