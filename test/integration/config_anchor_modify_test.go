package integration

import (
	"testing"
	"time"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
	"gitlab.n-t.io/core/library/hlf-tool/hlf-control-plane/test/integration/client/configuration"
	"gitlab.n-t.io/core/library/hlf-tool/hlf-control-plane/test/integration/models"
)

func TestConfigAnchorModify(t *testing.T) {
	runner.Run(t, "Try to modify anchor peer list and return back", func(t provider.T) {
		t.Severity(allure.CRITICAL)
		t.Description("Check method of modifying current list of anchor peers")
		t.Tags("anchor")

		t.WithNewStep("Trying to modify and revert anchor peers of org0", func(sCtx provider.StepCtx) {
			peersToSet := []*models.ProtosAnchorPeer{
				{Host: "test-peer-001.org0", Port: 7051},
			}

			var curSet []*models.ProtosAnchorPeer

			sCtx.WithNewStep("Get current anchor peer values", func(sCtx provider.StepCtx) {
				resp, err := org0Cli.Configuration.AnchorPeerList(&configuration.AnchorPeerListParams{ChannelName: channelName, Context: ctx}, auth)
				sCtx.Require().NoError(err)
				sCtx.Assert().NotNil(resp.Payload.Result)
				curSet = append([]*models.ProtosAnchorPeer{}, resp.Payload.Result...)
			})

			sCtx.WithNewStep("Modify anchor peers", func(sCtx provider.StepCtx) {
				resp, err := org0Cli.Configuration.AnchorPeerModify(&configuration.AnchorPeerModifyParams{
					ChannelName: channelName,
					Body: configuration.AnchorPeerModifyBody{
						Peers: peersToSet,
					},
					Context: ctx,
				}, auth)
				sCtx.Require().NoError(err)
				sCtx.Require().NotNil(resp.Payload)
			})

			time.Sleep(time.Second * 3)

			sCtx.WithNewStep("Get current anchor peer values", func(sCtx provider.StepCtx) {
				resp, err := org0Cli.Configuration.AnchorPeerList(&configuration.AnchorPeerListParams{ChannelName: channelName, Context: ctx}, auth)
				sCtx.Require().NoError(err)
				sCtx.Assert().NotNil(resp.Payload.Result)
				sCtx.Assert().Len(resp.Payload.Result, 1)
			})

			sCtx.WithNewStep("Revert changes back", func(sCtx provider.StepCtx) {
				resp, err := org0Cli.Configuration.AnchorPeerModify(&configuration.AnchorPeerModifyParams{
					ChannelName: channelName,
					Body: configuration.AnchorPeerModifyBody{
						Peers: curSet,
					},
					Context: ctx,
				}, auth)
				sCtx.Require().NoError(err)
				sCtx.Require().NotNil(resp.Payload)
			})
		})

		// t.WithNewStep("Trying to modify and revert anchor peers of org1", func(sCtx provider.StepCtx) {
		// 	peersToSet := []*models.ProtosAnchorPeer{
		// 		{Host: "test-peer-001.org1", Port: 7051},
		// 		{Host: "test-peer-002.org1", Port: 7051},
		// 	}
		//
		// 	var curSet []*models.ProtosAnchorPeer
		//
		// 	sCtx.WithNewStep("Get current anchor peer values", func(sCtx provider.StepCtx) {
		// 		resp, err := org1Cli.Configuration.AnchorPeerList(&configuration.AnchorPeerListParams{ChannelName: channelName, Context: ctx}, auth)
		// 		sCtx.Assert().NoError(err)
		// 		sCtx.Assert().NotNil(resp.Payload.Result)
		// 		curSet = resp.Payload.Result
		// 	})
		//
		// 	sCtx.WithNewStep("Modify anchor peers", func(sCtx provider.StepCtx) {
		// 		resp, err := org1Cli.Configuration.AnchorPeerModify(&configuration.AnchorPeerModifyParams{
		// 			ChannelName: channelName,
		// 			Body: configuration.AnchorPeerModifyBody{
		// 				Peers: peersToSet,
		// 			},
		// 			Context: ctx,
		// 		}, auth)
		// 		sCtx.Assert().NoError(err)
		// 		sCtx.Assert().NotNil(resp.Payload)
		// 		sCtx.Assert().Len(resp.Payload.New, 1)
		// 		sCtx.Assert().Len(resp.Payload.Existed, 1)
		// 	})
		//
		// 	sCtx.WithNewStep("Revert changes back", func(sCtx provider.StepCtx) {
		// 		resp, err := org1Cli.Configuration.AnchorPeerModify(&configuration.AnchorPeerModifyParams{
		// 			ChannelName: channelName,
		// 			Body: configuration.AnchorPeerModifyBody{
		// 				Peers: curSet,
		// 			},
		// 			Context: ctx,
		// 		}, auth)
		// 		sCtx.Assert().NoError(err)
		// 		sCtx.Assert().Len(resp.Payload.Deleted, 1)
		// 		sCtx.Assert().Len(resp.Payload.Existed, 1)
		// 	})
		// })
	})
}
