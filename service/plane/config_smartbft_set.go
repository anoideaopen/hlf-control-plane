package plane

import (
	"context"

	"github.com/golang/protobuf/proto" //nolint:staticcheck
	"github.com/hyperledger/fabric-config/configtx"
	"github.com/hyperledger/fabric-protos-go/common"
	protosorderer "github.com/hyperledger/fabric-protos-go/orderer"
	"github.com/hyperledger/fabric-protos-go/orderer/smartbft"
	"github.com/hyperledger/fabric/common/channelconfig"
	"github.com/hyperledger/fabric/protoutil"
	"gitlab.n-t.io/core/library/hlf-tool/hlf-control-plane/pkg/util"
	pb "gitlab.n-t.io/core/library/hlf-tool/hlf-control-plane/proto"
	"gitlab.n-t.io/core/library/hlf-tool/hlf-control-plane/system/cscc"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *srv) ConfigSmartBFTSet(ctx context.Context, req *pb.ConfigSmartBFTSetRequest) (*pb.ConfigSmartBFTSetResponse, error) {
	logger := s.logger.With(zap.String("channel", req.ChannelName))
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
	updatedConfig, _ := proto.Clone(config).(*common.Config)

	consensusTypeConfigValue := updatedConfig.
		ChannelGroup.Groups[channelconfig.OrdererGroupKey].
		Values[channelconfig.ConsensusTypeKey]
	consensusTypeValue := &protosorderer.ConsensusType{}
	if err = proto.Unmarshal(consensusTypeConfigValue.Value, consensusTypeValue); err != nil {
		return nil, status.Errorf(codes.Internal, "unmarshal consensus type value: %v", err)
	}

	metadata := &smartbft.ConfigMetadata{}
	if err = proto.Unmarshal(consensusTypeValue.Metadata, metadata); err != nil {
		return nil, status.Errorf(codes.Internal, "unmarshal metadata: %v", err)
	}

	metadata.Options = req.Options

	newMetadata, err := proto.Marshal(metadata)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "marshal metadata: %v", err)
	}
	consensusTypeValue.Metadata = newMetadata

	updatedConfig.
		ChannelGroup.Groups[channelconfig.OrdererGroupKey].
		Values[channelconfig.ConsensusTypeKey] = &common.ConfigValue{
		ModPolicy: configtx.AdminsPolicyKey,
		Value:     protoutil.MarshalOrPanic(consensusTypeValue),
	}

	upd, err := util.Compute(config, updatedConfig)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "compute update: %v", err)
	}

	env, err := s.createUpdateEnvelope(req.ChannelName, upd)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "create envelope: %v", err)
	}

	err = s.channelUpdateToOneOrderer(ctx, config, env)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "channel update to one orderer: %v", err)
	}

	return &pb.ConfigSmartBFTSetResponse{}, nil
}
