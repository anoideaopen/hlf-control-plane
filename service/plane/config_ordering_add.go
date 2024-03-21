package plane

import (
	"context"
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"time"

	"github.com/atomyze-foundation/hlf-control-plane/pkg/orderer"
	"github.com/atomyze-foundation/hlf-control-plane/pkg/util"
	pb "github.com/atomyze-foundation/hlf-control-plane/proto"
	"github.com/atomyze-foundation/hlf-control-plane/system/cscc"
	"github.com/golang/protobuf/proto" //nolint:staticcheck
	"github.com/hyperledger/fabric-protos-go/common"
	fbMsp "github.com/hyperledger/fabric-protos-go/msp"
	pbord "github.com/hyperledger/fabric-protos-go/orderer"
	"github.com/hyperledger/fabric-protos-go/orderer/etcdraft"
	"github.com/hyperledger/fabric-protos-go/orderer/smartbft"
	"github.com/hyperledger/fabric/common/channelconfig"
	"github.com/hyperledger/fabric/common/configtx"
	"github.com/hyperledger/fabric/msp"
	"github.com/hyperledger/fabric/protoutil"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *srv) ConfigOrderingAdd(ctx context.Context, req *pb.ConfigOrderingAddRequest) (*pb.ConfigOrderingAddResponse, error) {
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

	for _, o := range orderers {
		if o.Host == req.Orderer.Host && o.Port == req.Orderer.Port {
			// orderer found, nothing to do
			return &pb.ConfigOrderingAddResponse{}, nil
		}
	}

	orderers = append(orderers, req.Orderer)
	if err = s.proceedOrderingConsenterUpdate(ctx, req.ChannelName, config, orderers); err != nil {
		return nil, status.Errorf(codes.Internal, "proceed update: %v", err)
	}

	return &pb.ConfigOrderingAddResponse{}, nil
}

func (s *srv) proceedOrderingConsenterUpdate(ctx context.Context, channelName string, config *common.Config, orderers []*pb.Orderer, excludeSendOrderer ...*pb.Orderer) error {
	updated, _ := proto.Clone(config).(*common.Config)

	ordererGroup, ok := updated.ChannelGroup.Groups[channelconfig.OrdererGroupKey]
	if !ok {
		return fmt.Errorf("orderer group not found")
	}

	cType := new(pbord.ConsensusType)
	if err := proto.Unmarshal(ordererGroup.Values[channelconfig.ConsensusTypeKey].Value, cType); err != nil {
		return fmt.Errorf("unmarshal consensus value: %w", err)
	}

	var (
		mdBytes []byte
		err     error
	)

	switch util.ConvertConsensusType(cType) {
	case pb.ConsensusType_CONSENSUS_TYPE_RAFT:
		if mdBytes, err = s.createRaftOrderersUpdEnvelope(cType.Metadata, orderers); err != nil {
			return fmt.Errorf("create raft update envelope: %w", err)
		}
	case pb.ConsensusType_CONSENSUS_TYPE_BFT:
		var addr common.OrdererAddresses
		if err = proto.Unmarshal(updated.ChannelGroup.Values[channelconfig.OrdererAddressesKey].Value, &addr); err != nil {
			return fmt.Errorf("get endpoints value: %w", err)
		}

		address := make([]string, 0, len(orderers))
		for _, ord := range orderers {
			address = append(address, net.JoinHostPort(ord.Host, strconv.Itoa(int(ord.Port))))
		}
		addr.Addresses = address
		if updated.ChannelGroup.Values[channelconfig.OrdererAddressesKey].Value, err = proto.Marshal(&addr); err != nil {
			return fmt.Errorf("marshal endpoints: %w", err)
		}

		if mdBytes, err = s.createBftOrderersUpdEnvelope(cType.Metadata, orderers); err != nil {
			return fmt.Errorf("create bft update envelope: %w", err)
		}
	}

	cType.Metadata = mdBytes
	if updated.ChannelGroup.Groups[channelconfig.OrdererGroupKey].Values[channelconfig.ConsensusTypeKey].Value, err = proto.Marshal(cType); err != nil {
		return fmt.Errorf("marshal metadata: %w", err)
	}

	upd, err := util.Compute(config, updated)
	if err != nil {
		return fmt.Errorf("compute update: %w", err)
	}

	env, err := s.createUpdateEnvelope(channelName, upd)
	if err != nil {
		return fmt.Errorf("create envelope: %w", err)
	}

	return s.sendUpdateAndCompareResult(ctx, func() error {
		return s.channelUpdateToOneOrderer(ctx, config, env, excludeSendOrderer...)
	}, config, env, channelName)
}

func (s *srv) createRaftOrderersUpdEnvelope(md []byte, orderers []*pb.Orderer) ([]byte, error) {
	metadata := new(etcdraft.ConfigMetadata)
	if err := proto.Unmarshal(md, metadata); err != nil {
		return nil, fmt.Errorf("unmarshal bft metadata: %w", err)
	}

	cons := make([]*etcdraft.Consenter, 0, len(orderers))
	for _, ord := range orderers {
		cons = append(cons, &etcdraft.Consenter{
			Host:          ord.Host,
			Port:          ord.Port,
			ClientTlsCert: ord.Cert,
			ServerTlsCert: ord.Cert,
		})
	}

	metadata.Consenters = cons
	mdBytes, err := proto.Marshal(metadata)
	if err != nil {
		return nil, fmt.Errorf("marshal metadata: %w", err)
	}

	return mdBytes, nil
}

func (s *srv) createBftOrderersUpdEnvelope(md []byte, orderers []*pb.Orderer) ([]byte, error) {
	metadata := new(smartbft.ConfigMetadata)
	if err := proto.Unmarshal(md, metadata); err != nil {
		return nil, fmt.Errorf("unmarshal bft metadata: %w", err)
	}

	var (
		consID    uint64
		consIDMap = make(map[string]uint64)
	)
	// populate max value and consenterId map for future usage
	for _, ord := range metadata.Consenters {
		consIDMap[fmt.Sprintf("%s:%d", ord.Host, ord.Port)] = ord.ConsenterId
		if ord.ConsenterId > consID {
			consID = ord.ConsenterId
		}
	}
	// populate updated consenters value
	cons := make([]*smartbft.Consenter, 0, len(orderers))
	for _, ord := range orderers {
		cID, ok := consIDMap[fmt.Sprintf("%s:%d", ord.Host, ord.Port)]
		if !ok {
			consID++
			cID = consID
		}

		ordID := ord.Identity
		var id fbMsp.SerializedIdentity
		if err := proto.Unmarshal(ord.Identity, &id); err != nil {
			ordID, err = msp.NewSerializedIdentity(ord.MspId, ord.Identity)
			if err != nil {
				return nil, fmt.Errorf("marshal id: %w", err)
			}
		}

		cons = append(cons, &smartbft.Consenter{
			Host:          ord.Host,
			Port:          ord.Port,
			ClientTlsCert: ord.Cert,
			ServerTlsCert: ord.Cert,
			Identity:      ordID,
			MspId:         ord.MspId,
			ConsenterId:   cID,
		})
	}

	metadata.Consenters = cons
	mdBytes, err := proto.Marshal(metadata)
	if err != nil {
		return nil, fmt.Errorf("marshal metadata: %w", err)
	}

	return mdBytes, nil
}

func (s *srv) channelUpdateToOneOrderer(ctx context.Context, conf *common.Config, env *common.Envelope, excludeSendOrderer ...*pb.Orderer) error {
	orderers, _, err := util.GetOrdererConfig(conf)
	if err != nil {
		return fmt.Errorf("get orderer config: %w", err)
	}

external:
	for i := range excludeSendOrderer {
		for j := range orderers {
			if excludeSendOrderer[i].Host == orderers[j].Host && excludeSendOrderer[i].Port == orderers[j].Port ||
				excludeSendOrderer[i].ConsenterId == orderers[j].ConsenterId {
				orderers[j] = orderers[len(orderers)-1]
				orderers[len(orderers)-1] = nil
				orderers = orderers[:len(orderers)-1]

				continue external
			}
		}
	}

	if len(orderers) == 0 {
		return fmt.Errorf("get orderer from pool: pool length is zero")
	}

	ordRnd := orderers[rand.Intn(len(orderers)-1)]
	ordCli, err := s.ordPool.Get(&orderer.Orderer{
		Host:         ordRnd.Host,
		Port:         ordRnd.Port,
		Certificates: ordRnd.CaCerts,
	})
	if err != nil {
		return fmt.Errorf("get orderer from pool: %w", err)
	}
	return util.OrdererBroadcast(ctx, s.logger, env, 1, ordCli)
}

const (
	timeoutInterval = time.Minute * 2
)

func (s *srv) sendUpdateAndCompareResult( //nolint:funlen
	ctx context.Context,
	send func() error,
	config *common.Config,
	update *common.Envelope,
	channelName string,
) error {
	blockNumber, err := s.lastBlockNumberFromOrderer(ctx, config, channelName)
	if err != nil {
		return fmt.Errorf("get last block number: %w", err)
	}

	if err = send(); err != nil {
		return fmt.Errorf("send update: %w", err)
	}

	ctx, cancel := context.WithTimeout(ctx, timeoutInterval)
	defer func() {
		cancel()
	}()

	blockNumber, err = s.getNewConfigBlockNumberFromOrderer(ctx, config, channelName, blockNumber)
	if err != nil {
		return fmt.Errorf("get new config block numbere: %w", err)
	}

	select {
	case <-ctx.Done():
		return fmt.Errorf("context is done")
	default:
	}

	orderers, _, err := util.GetOrdererConfig(config)
	if err != nil {
		s.logger.Error("get orderer config failed", zap.Error(err))
		return fmt.Errorf("get orderer config: %w", err)
	}

	resChan := make(chan error, len(orderers))

	for _, ord := range orderers {
		go func(ord *pb.Orderer) {
			err = s.getAndCompareEnv(ctx, update, ord, channelName, blockNumber)
			if err != nil {
				s.logger.Error("get and compare env", zap.String("error", err.Error()))
				resChan <- fmt.Errorf("get and compare env: %w", err)
				return
			}

			resChan <- nil
		}(ord)
	}

	var goodCount int

	nOrders := len(orderers)
	quorum := s.quorumBft(nOrders)
	var res error
	for i := 0; i < nOrders; i++ {
		select {
		case res = <-resChan:
		case <-ctx.Done():
			s.logger.Debug("context is done")
			return fmt.Errorf("context is done")
		}

		if res != nil {
			err = res
			continue
		}

		goodCount++
		if goodCount >= quorum {
			s.logger.Debug("quorum is reached", zap.Int("good", goodCount), zap.Int("quorum", quorum))
			return nil
		}
	}

	if goodCount > 0 {
		s.logger.Debug("quorum is not reached but there is success", zap.Int("good", goodCount), zap.Int("quorum", quorum))
		return nil
	}

	return err
}

func (s *srv) getAndCompareEnv(ctx context.Context, update *common.Envelope, ord *pb.Orderer, channelName string, number uint64) error {
	block, err := s.getBlocksFromNumberFromOrderer(ctx, ord.Host, ord.Port, ord.CaCerts, channelName, number)
	if err != nil {
		return fmt.Errorf("get block from orderer: %w", err)
	}

	envelopeConfig, err := protoutil.ExtractEnvelope(block, 0)
	if err != nil {
		return fmt.Errorf("extract envelope: %w", err)
	}

	envPayload, err := protoutil.UnmarshalPayload(envelopeConfig.Payload)
	if err != nil {
		return fmt.Errorf("get payload: %w", err)
	}

	configEnvelope, err := configtx.UnmarshalConfigEnvelope(envPayload.Data)
	if err != nil {
		return fmt.Errorf("get config envelope: %w", err)
	}

	if !proto.Equal(configEnvelope.LastUpdate, update) {
		return fmt.Errorf("two envelops aren't equal")
	}

	return nil
}

func (s *srv) getNewConfigBlockNumberFromOrderer(ctx context.Context, config *common.Config, channel string, number uint64) (uint64, error) {
	orderers, _, err := util.GetOrdererConfig(config)
	if err != nil {
		return 0, fmt.Errorf("get orderer config: %w", err)
	}

	if len(orderers) == 0 {
		return 0, fmt.Errorf("len orderer is zero")
	}

	for {
		select {
		case <-ctx.Done():
			return 0, fmt.Errorf("context is done")
		default:
		}

		number++
		block, err := s.getBlocksFromNumberFromOrderer(ctx, orderers[0].Host, orderers[0].Port, orderers[0].CaCerts, channel, number)
		if err != nil {
			s.logger.Error("getBlocksFromNumberFromOrderer", zap.String("error", err.Error()))
			return 0, fmt.Errorf("getBlocksFromNumberFromOrderer: %w", err)
		}

		if protoutil.IsConfigBlock(block) {
			return block.Header.Number, nil
		}
	}
}

func (s *srv) lastBlockNumberFromOrderer(ctx context.Context, config *common.Config, channel string) (uint64, error) {
	orderers, _, err := util.GetOrdererConfig(config)
	if err != nil {
		return 0, fmt.Errorf("get orderer config: %w", err)
	}

	if len(orderers) == 0 {
		return 0, fmt.Errorf("len orderer is zero")
	}

	block, err := s.getLastBlockFromOrderer(ctx, orderers[0].Host, orderers[0].Port, orderers[0].CaCerts, channel)
	if err != nil {
		s.logger.Error("getLastBlockFromOrderer", zap.String("error", err.Error()))
		return 0, fmt.Errorf("getLastBlockFromOrderer: %w", err)
	}

	return block.Header.Number, nil
}
