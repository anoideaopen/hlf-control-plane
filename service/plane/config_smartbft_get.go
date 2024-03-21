package plane

import (
	"context"

	pb "github.com/atomyze-foundation/hlf-control-plane/proto"
	"github.com/atomyze-foundation/hlf-control-plane/system/cscc"
	"github.com/golang/protobuf/proto" //nolint:staticcheck
	protosorderer "github.com/hyperledger/fabric-protos-go/orderer"
	"github.com/hyperledger/fabric-protos-go/orderer/smartbft"
	"github.com/hyperledger/fabric/common/channelconfig"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *srv) ConfigSmartBFTGet(ctx context.Context, req *pb.ConfigSmartBFTGetRequest) (*pb.ConfigSmartBFTGetResponse, error) {
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

	consensusTypeConfigValue := config.ChannelGroup.
		Groups[channelconfig.OrdererGroupKey].
		Values[channelconfig.ConsensusTypeKey]
	consensusTypeValue := &protosorderer.ConsensusType{}
	if err = proto.Unmarshal(consensusTypeConfigValue.Value, consensusTypeValue); err != nil {
		return nil, status.Errorf(codes.Internal, "unmarshal consensus type value: %v", err)
	}

	metadata := &smartbft.ConfigMetadata{}
	if err = proto.Unmarshal(consensusTypeValue.Metadata, metadata); err != nil {
		return nil, status.Errorf(codes.Internal, "unmarshal metadata: %v", err)
	}

	return &pb.ConfigSmartBFTGetResponse{
		Options: metadata.GetOptions(),
	}, nil
}
