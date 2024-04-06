package plane

import (
	"context"

	"github.com/golang/protobuf/proto" //nolint:staticcheck
	"github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric/common/channelconfig"
	"gitlab.n-t.io/core/library/hlf-tool/hlf-control-plane/pkg/util"
	pb "gitlab.n-t.io/core/library/hlf-tool/hlf-control-plane/proto"
	"gitlab.n-t.io/core/library/hlf-tool/hlf-control-plane/system/cscc"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *srv) ConfigOrganizationDelete(ctx context.Context, req *pb.ConfigOrganizationDeleteRequest) (*pb.ConfigOrganizationDeleteResponse, error) {
	logger := s.logger.With(zap.String("channel", req.ChannelName))

	// check section is one of: Application, Orderer and Channel values
	// otherwise, current request section value is incorrect
	if req.Section != channelconfig.ApplicationGroupKey && req.Section != channelconfig.OrdererGroupKey && req.Section != channelconfig.ChannelGroupKey {
		return nil, status.Errorf(codes.InvalidArgument, "channel config section is invalid")
	}

	logger.Debug("get channel config", zap.String("channel", req.ChannelName))
	endCli, err := s.peerPool.GetRandomEndorser(ctx, s.mspID)
	if err != nil {
		logger.Error("get endorser failed", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "get endorser failed: %v", err)
	}

	config, err := cscc.NewClient(endCli, s.id).GetChannelConfig(ctx, req.ChannelName)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "get config block: %v", err)
	}

	if _, ok := config.ChannelGroup.Groups[req.Section].Groups[req.Organization]; !ok {
		return &pb.ConfigOrganizationDeleteResponse{Existed: false}, nil
	}

	oldConf, ok := proto.Clone(config).(*common.Config)
	if !ok {
		return nil, status.Errorf(codes.Internal, "wrong config type")
	}

	delete(config.ChannelGroup.Groups[req.Section].Groups, req.Organization)

	upd, err := util.Compute(oldConf, config)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "compute update: %s", err)
	}

	env, err := s.createUpdateEnvelope(req.ChannelName, upd)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "create envelope: %s", err)
	}

	if err = s.sendUpdateAndCompareResult(ctx, func() error {
		return s.channelUpdateToOneOrderer(ctx, config, env)
	}, config, env, req.ChannelName); err != nil {
		return nil, status.Errorf(codes.Internal, "send update: %s", err)
	}

	return &pb.ConfigOrganizationDeleteResponse{Existed: true}, nil
}
