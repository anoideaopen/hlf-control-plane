package plane

import (
	"crypto/tls"

	"github.com/hyperledger/fabric/protoutil"
	sd "gitlab.n-t.io/core/library/hlf-tool/hlf-control-plane/pkg/delivery"
	"gitlab.n-t.io/core/library/hlf-tool/hlf-control-plane/pkg/discovery"
	"gitlab.n-t.io/core/library/hlf-tool/hlf-control-plane/pkg/orderer"
	"gitlab.n-t.io/core/library/hlf-tool/hlf-control-plane/pkg/peer"
	"gitlab.n-t.io/core/library/hlf-tool/hlf-control-plane/proto"
	"go.uber.org/zap"
)

// check server interface during compilation
var _ proto.ChaincodeServiceServer = &srv{}

type srv struct {
	mspID      string
	localPeers []*peer.Peer
	logger     *zap.Logger
	ordPool    orderer.Pool
	peerPool   peer.Pool
	id         protoutil.Signer
	dCli       sd.Client
	discCli    discovery.Client
	addIds     []protoutil.Signer
	tlsCli     *tls.Config

	proto.UnimplementedChaincodeServiceServer
}

func NewService(mspID string, signer protoutil.Signer, tlsConf *tls.Config, logger *zap.Logger, ordPool orderer.Pool, peerPool peer.Pool, discoveryCli discovery.Client, localPeers []*peer.Peer, addIds []protoutil.Signer) proto.ChaincodeServiceServer {
	return &srv{
		mspID:      mspID,
		localPeers: localPeers,
		logger:     logger.Named("plane").With(zap.String("mspId", mspID)),
		ordPool:    ordPool,
		id:         signer,
		tlsCli:     tlsConf,
		peerPool:   peerPool,
		discCli:    discoveryCli,
		dCli:       sd.NewPeer(logger, peerPool, localPeers, signer),
		addIds:     addIds,
	}
}
