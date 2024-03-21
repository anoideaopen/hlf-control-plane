package integration

import (
	"testing"

	"github.com/atomyze-foundation/hlf-control-plane/test/integration/client/channels"
	"github.com/atomyze-foundation/hlf-control-plane/test/integration/models"
	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
)

const expectedPeer = "test-peer-001.org0:7051"

// TestChannelGetJoined - Trying to get list of channels
func TestChannelGetJoined(t *testing.T) {
	runner.Run(t, "Trying to get list of channels", func(t provider.T) {
		t.Severity(allure.CRITICAL)
		t.Description("Trying to get list of channels")
		t.Tags("positive", "channel")
		t.WithNewAsyncStep("", func(sCtx provider.StepCtx) {
			resp, err := org0Cli.Channels.ChannelsList(&channels.ChannelsListParams{Context: ctx}, auth)
			sCtx.Require().NoError(err)
			for _, s := range resp.GetPayload().Result {
				peers := make([]string, 0, len(s.Peers))
				for _, p := range s.Peers {
					peers = append(peers, p.Peer)
				}
				if s.Name == "fiat" {
					sCtx.Require().Contains(peers, expectedPeer, "checking that peer is %s", expectedPeer)
				}
			}
		})
	})
}

// TestChannelPostJoin - Trying to join channel
func TestChannelPostJoin(t *testing.T) {
	runner.Run(t, "Trying to join channel", func(t provider.T) {
		t.Severity(allure.CRITICAL)
		t.Description("Trying to join channel")
		t.Tags("positive", "channel")
		t.WithNewAsyncStep("", func(sCtx provider.StepCtx) {
			orderer := models.ProtoChannelJoinRequestOrderer{
				Host: ordererHost,
				Port: ordererPort,
			}
			joinRequest := models.ProtoChannelJoinRequest{
				ChannelName: "cc",
				Orderer:     &orderer,
			}
			var s models.ProtoChannelJoinResponsePeerResult
			resp, err := org0Cli.Channels.ChannelJoin(&channels.ChannelJoinParams{Body: &joinRequest, Context: ctx}, auth)
			sCtx.Require().NoError(err)
			s = *resp.GetPayload().Result[0]
			sCtx.Require().Empty(s.Err, "checking that error is empty")
			sCtx.Require().Equal(false, s.Existed, "checking that peer is false")
			sCtx.Require().Equal(expectedPeer, s.Peer, "checking that peer is %s", expectedPeer)
		})
	})
}

// TestChannelPostJoinToAlreadyJoinedChannel - Trying to join to already joined channel
func TestChannelPostJoinToAlreadyJoinedChannel(t *testing.T) {
	runner.Run(t, "Trying to join to already joined channel", func(t provider.T) {
		t.Severity(allure.CRITICAL)
		t.Description("Trying to join to already joined channel")
		t.Tags("channel")
		t.WithNewAsyncStep("Trying to join to already joined channel", func(sCtx provider.StepCtx) {
			orderer := models.ProtoChannelJoinRequestOrderer{
				Host: ordererHost,
				Port: ordererPort,
			}
			joinRequest := models.ProtoChannelJoinRequest{
				ChannelName: "fiat",
				Orderer:     &orderer,
			}
			var s models.ProtoChannelJoinResponsePeerResult
			resp, err := org0Cli.Channels.ChannelJoin(&channels.ChannelJoinParams{Body: &joinRequest, Context: ctx}, auth)
			sCtx.Require().NoError(err)
			s = *resp.GetPayload().Result[0]
			sCtx.Require().Empty(s.Err, "checking that error is empty")
			sCtx.Require().Equal(true, s.Existed, "checking that peer is true")
			sCtx.Require().Equal(expectedPeer, s.Peer, "checking that peer is %s", expectedPeer)
		})
	})
}

// TestChannelPostJoinWrongChannel - Join channel with wrong channel and check that we get an error
func TestChannelPostJoinWrongChannel(t *testing.T) {
	runner.Run(t, "Trying to join channel with wrong channel", func(t provider.T) {
		t.Severity(allure.CRITICAL)
		t.Description("Trying to join channel with wrong channel")
		t.Tags("channel")
		t.WithNewAsyncStep("Join channel with wrong channel and check that we get an error", func(sCtx provider.StepCtx) {
			orderer := models.ProtoChannelJoinRequestOrderer{
				Host: ordererHost,
				Port: ordererPort,
			}
			joinRequest := models.ProtoChannelJoinRequest{
				ChannelName: "test",
				Orderer:     &orderer,
			}
			_, err := org0Cli.Channels.ChannelJoin(&channels.ChannelJoinParams{Body: &joinRequest, Context: ctx}, auth)
			sCtx.Require().Error(err)
		})
	})
}

// TestChannelPostJoinWrongHost - Join channel with wrong host and check that we get an error
func TestChannelPostJoinWrongHost(t *testing.T) {
	runner.Run(t, "Trying to join channel with wrong host", func(t provider.T) {
		t.Severity(allure.CRITICAL)
		t.Description("Trying to join channel with wrong host")
		t.Tags("channel")
		t.WithNewAsyncStep("Join channel with wrong host and check that we get an error", func(sCtx provider.StepCtx) {
			orderer := models.ProtoChannelJoinRequestOrderer{
				Host: "test",
				Port: ordererPort,
			}
			joinRequest := models.ProtoChannelJoinRequest{
				ChannelName: "cc",
				Orderer:     &orderer,
			}
			_, err := org0Cli.Channels.ChannelJoin(&channels.ChannelJoinParams{Body: &joinRequest, Context: ctx}, auth)
			sCtx.Require().Error(err)
		})
	})
}

// TestChannelPostJoinWrongPort - Join channel with wrong port and check that we get an error
func TestChannelPostJoinWrongPort(t *testing.T) {
	runner.Run(t, "Trying to join channel with wrong port", func(t provider.T) {
		t.Severity(allure.CRITICAL)
		t.Description("Trying to join channel with wrong or empty port")
		t.Tags("channel")
		t.WithNewAsyncStep("Join channel with wrong port and check that we get an error", func(sCtx provider.StepCtx) {
			orderer := models.ProtoChannelJoinRequestOrderer{
				Host: ordererHost,
				Port: int64(0),
			}
			joinRequest := models.ProtoChannelJoinRequest{
				ChannelName: "cc",
				Orderer:     &orderer,
			}
			_, err := org0Cli.Channels.ChannelJoin(&channels.ChannelJoinParams{Body: &joinRequest, Context: ctx}, auth)
			sCtx.Require().Error(err)
		})
	})
}
