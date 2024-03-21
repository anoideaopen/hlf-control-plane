package lifecycle

import (
	"context"
	"fmt"
	"strings"

	"github.com/golang/protobuf/proto" //nolint:staticcheck
	cb "github.com/hyperledger/fabric-protos-go/common"
	pb "github.com/hyperledger/fabric-protos-go/peer"
	lb "github.com/hyperledger/fabric-protos-go/peer/lifecycle"
	"github.com/hyperledger/fabric/protoutil"
	"github.com/pkg/errors"
	"gitlab.n-t.io/core/library/hlf-tool/hlf-control-plane/pkg/util"
)

const queryApprovedFunc = "QueryApprovedChaincodeDefinition"

func (c *cli) QueryApproved(ctx context.Context, channelName, chaincodeName string) (*lb.QueryApprovedChaincodeDefinitionResult, error) {
	prop, err := c.createQueryApprovedProposal(channelName, chaincodeName)
	if err != nil {
		return nil, fmt.Errorf("create proposal: %w", err)
	}

	signedProp, err := util.SignProposal(prop, c.id)
	if err != nil {
		return nil, fmt.Errorf("sign proposal: %w", err)
	}

	return c.processQueryApprovedProposal(ctx, signedProp)
}

func (c *cli) createQueryApprovedProposal(channelName, chaincodeName string) (*pb.Proposal, error) {
	argsBytes, err := proto.Marshal(&lb.QueryApprovedChaincodeDefinitionArgs{
		Name: chaincodeName,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal args")
	}

	ccInput := &pb.ChaincodeInput{Args: [][]byte{[]byte(queryApprovedFunc), argsBytes}}

	cis := &pb.ChaincodeInvocationSpec{
		ChaincodeSpec: &pb.ChaincodeSpec{
			ChaincodeId: &pb.ChaincodeID{Name: CcName},
			Input:       ccInput,
		},
	}

	signer, err := c.id.Serialize()
	if err != nil {
		return nil, errors.WithMessage(err, "failed to serialize identity")
	}

	proposal, _, err := protoutil.CreateProposalFromCIS(cb.HeaderType_ENDORSER_TRANSACTION, channelName, cis, signer)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to create ChaincodeInvocationSpec proposal")
	}

	return proposal, nil
}

func (c *cli) processQueryApprovedProposal(ctx context.Context, singedProp *pb.SignedProposal) (*lb.QueryApprovedChaincodeDefinitionResult, error) {
	resp, err := c.cli.ProcessProposal(ctx, singedProp)
	if err != nil {
		return nil, fmt.Errorf("endorse proposal: %w", err)
	}

	if resp == nil {
		return nil, fmt.Errorf("received nil proposal response")
	}

	if resp.Response == nil {
		return nil, fmt.Errorf("received proposal response with nil response")
	}

	if resp.Response.Status != int32(cb.Status_SUCCESS) {
		if strings.Contains(resp.Response.Message, "could not fetch approved chaincode definition") {
			return &lb.QueryApprovedChaincodeDefinitionResult{}, nil
		}
		return nil, fmt.Errorf("query failed with status: %d - %s", resp.Response.Status, resp.Response.Message)
	}

	var result lb.QueryApprovedChaincodeDefinitionResult
	if err = proto.Unmarshal(resp.Response.Payload, &result); err != nil {
		return nil, fmt.Errorf("proto unmarshal: %w", err)
	}

	return &result, nil
}
