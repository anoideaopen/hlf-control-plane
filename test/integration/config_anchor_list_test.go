package integration

import (
	"testing"

	"github.com/atomyze-foundation/hlf-control-plane/test/integration/client/configuration"
	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
)

func TestConfigAnchorList(t *testing.T) {
	runner.Run(t, "Trying to get anchor peer list", func(t provider.T) {
		t.Severity(allure.CRITICAL)
		t.Description("Check method of getting list of current anchor peers")
		t.Tags("anchor")

		t.WithNewAsyncStep("Trying to get anchor peers of org0", func(sCtx provider.StepCtx) {
			resp, err := org0Cli.Configuration.AnchorPeerList(&configuration.AnchorPeerListParams{ChannelName: channelName, Context: ctx}, auth)
			sCtx.Require().NoError(err)
			sCtx.Assert().Len(resp.Payload.Result, 2)
		})
	})
}
