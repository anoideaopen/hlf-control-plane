package integration

import (
	"strconv"
	"testing"
	"time"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
	"gitlab.n-t.io/core/library/hlf-tool/hlf-control-plane/test/integration/client/configuration"
	"gitlab.n-t.io/core/library/hlf-tool/hlf-control-plane/test/integration/client/lifecycle"
	"gitlab.n-t.io/core/library/hlf-tool/hlf-control-plane/test/integration/models"
)

// TestChannelGetLifecycleNoChaincode - Trying to get lifecycle no chaincode
func TestChannelGetLifecycleNoChaincode(t *testing.T) {
	runner.Run(t, "Trying to get lifecycle no chaincode", func(t provider.T) {
		t.Severity(allure.CRITICAL)
		t.Description("Trying to get lifecycle no chaincode")
		t.Tags("lifecycle")
		t.WithNewAsyncStep("", func(sCtx provider.StepCtx) {
			resp, err := org0Cli.Lifecycle.LifecycleStatus(&lifecycle.LifecycleStatusParams{ChannelName: channelName, Context: ctx}, auth)
			sCtx.Require().NoError(err)
			sCtx.Require().Len(resp.Payload.Chaincodes, 0, "checking that chaincodes is empty")
		})
	})
}

// TestChannelGetLifecycleWrongChannel - Trying to get lifecycle with wrong channel
func TestChannelGetLifecycleWrongChannel(t *testing.T) {
	runner.Run(t, "Trying to get lifecycle with wrong channel", func(t provider.T) {
		t.Severity(allure.CRITICAL)
		t.Description("Trying to get lifecycle with wrong channel")
		t.Tags("lifecycle")
		t.WithNewAsyncStep("", func(sCtx provider.StepCtx) {
			_, err := org0Cli.Lifecycle.LifecycleStatus(&lifecycle.LifecycleStatusParams{ChannelName: "test", Context: ctx}, auth)
			sCtx.Require().Error(err)
		})
	})
}

// TestChannelPostLifecycleFullAndGetIt - Add new chaincode version than commit it and get committed chaincode on channel
func TestChannelPostLifecycleFullAndGetIt(t *testing.T) {
	runner.Run(t, "Add new chaincode version than commit it and get committed chaincode on channel", func(t provider.T) {
		t.Severity(allure.CRITICAL)
		t.Description("Add new chaincode version than commit it and get committed chaincode on channel")
		t.Tags("lifecycle")
		t.WithNewStep("Add new chaincode version than commit it and get committed chaincode on channel", func(sCtx provider.StepCtx) {
			industrialChannel := "industrial"
			cVersion := "1676556954"
			cPolicy := "AND('org0.member')"

			peersToSet := []*models.ProtosAnchorPeer{
				{Host: "test-peer-001.org0", Port: 7051},
			}

			peersToSet2 := []*models.ProtosAnchorPeer{
				{Host: "test-peer-001.org1", Port: 7051},
			}

			sCtx.WithNewStep("Modify anchor peers", func(sCtx provider.StepCtx) {
				_, err := org0Cli.Configuration.AnchorPeerModify(&configuration.AnchorPeerModifyParams{
					ChannelName: industrialChannel,
					Body: configuration.AnchorPeerModifyBody{
						Peers: peersToSet,
					},
					Context: ctx,
				}, auth)
				sCtx.Require().NoError(err)
			})
			sCtx.WithNewStep("Modify anchor peers", func(sCtx provider.StepCtx) {
				_, err := org1Cli.Configuration.AnchorPeerModify(&configuration.AnchorPeerModifyParams{
					ChannelName: industrialChannel,
					Body: configuration.AnchorPeerModifyBody{
						Peers: peersToSet2,
					},
					Context: ctx,
				}, auth)
				sCtx.Require().NoError(err)
			})

			time.Sleep(1 * time.Second)

			requestBody := models.ProtoLifecycleFullRequest{
				ChannelName:    industrialChannel,
				ChaincodeName:  industrialChannel,
				ChaincodeLabel: "industrial_1676556920",
				Policy:         cPolicy,
				Version:        cVersion,
				InitArgs:       []string{ski2, ski1, issuer},
				CommitForce:    false,
			}

			sCtx.WithNewStep("Commit from peer org1", func(sCtx provider.StepCtx) {
				resp, err := org1Cli.Lifecycle.LifecycleFull(&lifecycle.LifecycleFullParams{Body: &requestBody, Context: ctx}, auth)
				sCtx.Require().NoError(err)
				sCtx.Require().Equal(false, resp.GetPayload().Committed, "checking that committed is false")
				sCtx.Require().Len(resp.GetPayload().Approvals, 2)
			})

			sCtx.WithNewStep("Commit from peer org0", func(sCtx provider.StepCtx) {
				resp, err := org0Cli.Lifecycle.LifecycleFull(&lifecycle.LifecycleFullParams{Body: &requestBody, Context: ctx}, auth)
				sCtx.Require().NoError(err)
				sCtx.Require().Equal(true, resp.GetPayload().Committed, "checking that committed is true")
				sCtx.Require().Len(resp.GetPayload().Approvals, 0)
			})

			sCtx.WithNewStep("Trying to get lifecycle", func(sCtx provider.StepCtx) {
				var s models.ProtoLifecycleChaincode
				resp, err := org0Cli.Lifecycle.LifecycleStatus(&lifecycle.LifecycleStatusParams{ChannelName: industrialChannel, Context: ctx}, auth)
				sCtx.Require().NoError(err)
				s = *resp.GetPayload().Chaincodes[0]
				sCtx.Require().Equal(true, s.InitRequired)
				sCtx.Require().Equal(industrialChannel, s.Name, "checking that name equal %s", industrialChannel)
				sCtx.Require().Equal(cVersion, s.Version, "checking version")
				seq, err := strconv.Atoi(s.Sequence)
				sCtx.Require().NoError(err)
				sCtx.Require().True(seq > 0)
			})

			requestBody.Version = cVersion + "_1"
			requestBody.InitRequired = false

			sCtx.WithNewStep("Commit from peer org1 with next version", func(sCtx provider.StepCtx) {
				resp, err := org1Cli.Lifecycle.LifecycleFull(&lifecycle.LifecycleFullParams{Body: &requestBody, Context: ctx}, auth)
				sCtx.Require().NoError(err)
				sCtx.Require().Equal(false, resp.GetPayload().Committed, "checking that committed is false")
				sCtx.Require().Len(resp.GetPayload().Approvals, 2)
			})

			requestBody.CommitForce = true

			sCtx.WithNewStep("Commit from peer org0 with next version", func(sCtx provider.StepCtx) {
				resp, err := org0Cli.Lifecycle.LifecycleFull(&lifecycle.LifecycleFullParams{Body: &requestBody, Context: ctx}, auth)
				sCtx.Require().NoError(err)
				sCtx.Require().Equal(true, resp.GetPayload().Committed, "checking that committed is true")
				sCtx.Require().Len(resp.GetPayload().Approvals, 0)
			})

			var seq int

			sCtx.WithNewStep("Trying to get lifecycle", func(sCtx provider.StepCtx) {
				var s models.ProtoLifecycleChaincode
				resp, err := org0Cli.Lifecycle.LifecycleStatus(&lifecycle.LifecycleStatusParams{ChannelName: industrialChannel, Context: ctx}, auth)
				sCtx.Require().NoError(err)
				s = *resp.GetPayload().Chaincodes[0]
				sCtx.Require().Equal(true, s.InitRequired)
				sCtx.Require().Equal(industrialChannel, s.Name, "checking that name equal %s", industrialChannel)
				sCtx.Require().Equal(cVersion+"_1", s.Version, "checking version")
				seq, err = strconv.Atoi(s.Sequence)
				sCtx.Require().NoError(err)
				sCtx.Require().True(seq > 0)
			})

			sCtx.WithNewStep("Send ApproveForMyOrg", func(sCtx provider.StepCtx) {
				_, err := org0Cli.Lifecycle.LifecycleApproveForMyOrg(&lifecycle.LifecycleApproveForMyOrgParams{
					Body: &models.ProtoLifecycleApproveForMyOrgRequest{
						ChaincodeLabel: "industrial_1676556920",
						ChaincodeName:  industrialChannel,
						ChannelName:    industrialChannel,
						InitRequired:   true,
						Policy:         cPolicy,
						Sequence:       strconv.Itoa(seq + 1),
						Version:        cVersion + "_2",
					},
					Context: ctx,
				}, auth)
				sCtx.Require().NoError(err)
			})

			sCtx.WithNewStep("Send ApproveForMyOrg", func(sCtx provider.StepCtx) {
				_, err := org1Cli.Lifecycle.LifecycleApproveForMyOrg(&lifecycle.LifecycleApproveForMyOrgParams{
					Body: &models.ProtoLifecycleApproveForMyOrgRequest{
						ChaincodeLabel: "industrial_1676556920",
						ChaincodeName:  industrialChannel,
						ChannelName:    industrialChannel,
						InitRequired:   true,
						Policy:         cPolicy,
						Sequence:       strconv.Itoa(seq + 1),
						Version:        cVersion + "_2",
					},
					Context: ctx,
				}, auth)
				sCtx.Require().NoError(err)
			})

			sCtx.WithNewStep("Get Approved", func(sCtx provider.StepCtx) {
				var s models.ProtoLifecycleApprovedResponse
				resp, err := org0Cli.Lifecycle.LifecycleApproved(&lifecycle.LifecycleApprovedParams{
					ChaincodeName: industrialChannel,
					ChannelName:   industrialChannel,
					Context:       ctx,
				}, auth)
				sCtx.Require().NoError(err)
				s = *resp.GetPayload()
				sCtx.Require().Equal(true, s.Chaincode.InitRequired)
				sCtx.Require().Equal(industrialChannel, s.Chaincode.Name, "checking that name equal %s", industrialChannel)
				sCtx.Require().Equal(cVersion+"_2", s.Chaincode.Version, "checking version")
				seqResp, err := strconv.Atoi(s.Chaincode.Sequence)
				sCtx.Require().NoError(err)
				sCtx.Require().Equal(seq+1, seqResp)
				sCtx.Log("package id ", s.PackageID)
			})

			sCtx.WithNewStep("Get CheckCommitReadiness", func(sCtx provider.StepCtx) {
				resp, err := org0Cli.Lifecycle.LifecycleCheckCommitReadiness(&lifecycle.LifecycleCheckCommitReadinessParams{
					Body: &models.ProtoLifecycleCheckCommitReadinessRequest{
						ChaincodeName: industrialChannel,
						ChannelName:   industrialChannel,
						InitRequired:  true,
						Policy:        cPolicy,
						Sequence:      strconv.Itoa(seq + 1),
						Version:       cVersion + "_2",
					},
					Context: ctx,
				}, auth)
				sCtx.Require().NoError(err)
				s := *resp.GetPayload()
				for o, b := range s.Approvals {
					sCtx.Require().True(b, "check approvals for %s", o)
				}
			})

			sCtx.WithNewStep("Send Commit", func(sCtx provider.StepCtx) {
				_, err := org0Cli.Lifecycle.LifecycleCommit(&lifecycle.LifecycleCommitParams{
					Body: &models.ProtoLifecycleCommitRequest{
						ChaincodeName: industrialChannel,
						ChannelName:   industrialChannel,
						InitRequired:  true,
						Policy:        cPolicy,
						Sequence:      strconv.Itoa(seq + 1),
						Version:       cVersion + "_2",
					},
					Context: ctx,
				}, auth)
				sCtx.Require().NoError(err)
			})

			sCtx.WithNewStep("Trying to get lifecycle", func(sCtx provider.StepCtx) {
				var s models.ProtoLifecycleChaincode
				resp, err := org0Cli.Lifecycle.LifecycleStatus(&lifecycle.LifecycleStatusParams{ChannelName: industrialChannel, Context: ctx}, auth)
				sCtx.Require().NoError(err)
				s = *resp.GetPayload().Chaincodes[0]
				sCtx.Require().Equal(true, s.InitRequired)
				sCtx.Require().Equal(industrialChannel, s.Name, "checking that name equal %s", industrialChannel)
				sCtx.Require().Equal(cVersion+"_2", s.Version, "checking version")
				seqResp, err := strconv.Atoi(s.Sequence)
				sCtx.Require().NoError(err)
				sCtx.Require().Equal(seq+1, seqResp)
			})

			sCtx.WithNewStep("Send Init", func(sCtx provider.StepCtx) {
				_, err := org0Cli.Lifecycle.LifecycleInit(&lifecycle.LifecycleInitParams{
					Body: &models.ProtoLifecycleInitRequest{
						ChaincodeName: industrialChannel,
						ChannelName:   industrialChannel,
						InitArgs:      []string{ski2, ski1, issuer},
					},
					Context: ctx,
				}, auth)
				sCtx.Require().NoError(err)
			})
		})
	})
}
