package discovery

import (
	"context"

	"github.com/hyperledger/fabric-protos-go/discovery"
	"gitlab.n-t.io/core/library/hlf-tool/hlf-control-plane/proto"
)

type Client interface {
	GetPeers(ctx context.Context, channelName string) ([]*proto.DiscoveryPeer, error)
	GetEndorsers(ctx context.Context, channelName, ccName string) ([]*proto.DiscoveryPeer, error)
	GetConfig(ctx context.Context, channelName string) (*discovery.ConfigResult, error)
}
