package cscc

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/proto" //nolint:staticcheck
	"github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric-protos-go/peer"
	"gitlab.n-t.io/core/library/hlf-tool/hlf-control-plane/pkg/util"
)

func (c *cli) GetChannelConfigBlock(ctx context.Context, channelName string) (*common.Block, error) {
	// create chaincode endorsement proposal
	prop, err := c.getChannelConfigBlockProposal(channelName)
	if err != nil {
		return nil, fmt.Errorf("get proposal: %w", err)
	}

	// create signed proposal using identity
	signedProp, err := util.SignProposal(prop, c.id)
	if err != nil {
		return nil, fmt.Errorf("create singed proposal: %w", err)
	}

	// process signed proposal on peer
	config, err := c.processGetChannelConfigBlockProposal(ctx, signedProp)
	if err != nil {
		return nil, fmt.Errorf("process proposal: %w", err)
	}
	return config, nil
}

func (c *cli) processGetChannelConfigBlockProposal(ctx context.Context, prop *peer.SignedProposal) (*common.Block, error) {
	resp, err := c.cli.ProcessProposal(ctx, prop)
	if err != nil {
		return nil, fmt.Errorf("process proposal: %w", err)
	}

	if resp.Response == nil || resp.Response.Status != 200 {
		return nil, fmt.Errorf("received bad response, status %d: %s", resp.Response.Status, resp.Response.Message)
	}

	block := &common.Block{}
	if err = proto.Unmarshal(resp.Response.Payload, block); err != nil {
		return nil, fmt.Errorf("proto unmarshal: %w", err)
	}

	return block, nil
}
