package integration

import (
	"encoding/base64"
	"testing"

	"github.com/ozontech/allure-go/pkg/allure"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
	"gitlab.n-t.io/core/library/hlf-tool/hlf-control-plane/test/integration/client/chaincode"
	"gitlab.n-t.io/core/library/hlf-tool/hlf-control-plane/test/integration/models"
)

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
}

const (
	aclChaincodePath = "/repository/raw/core/library/chaincode/acl_2.4-8296e079.tar.gz"
)

func TestChaincode(t *testing.T) {
	// skip if environment is not set
	if netRcUser == "" {
		t.Skip()
	}

	var installedCCLen int

	runner.Run(t, "Test chaincode list", func(t provider.T) {
		t.Severity(allure.CRITICAL)
		t.Description("Check work of chaincode installed list method")
		resp, err := org0Cli.Chaincode.ChaincodeInstalledList(&chaincode.ChaincodeInstalledListParams{
			Context: ctx,
		}, auth)
		t.Require().NoError(err)
		t.Require().NotNil(resp.Payload.Result)
		installedCCLen = len(resp.Payload.Result)
	})

	runner.Run(t, "Test chaincode install", func(t provider.T) {
		t.Severity(allure.CRITICAL)
		t.Description("Chaincode installation test suite, Install chaincode using external source - Nexus")
		ccPath := netRcHost + aclChaincodePath
		resp, err := org0Cli.Chaincode.ChaincodeInstall(&chaincode.ChaincodeInstallParams{
			Body: &models.ProtoChaincodeInstallRequest{
				AuthHeaders: map[string]string{
					`Authorization`: basicAuth(netRcUser, netRcPassword),
				},
				Source: &ccPath,
			},
			Context: ctx,
		}, auth)
		t.Require().NoError(err)
		t.Require().NotNil(resp.Payload.Result)
		for _, res := range resp.Payload.Result {
			t.WithNewAsyncStep("Check peer "+res.Peer+" results", func(sCtx provider.StepCtx) {
				t.Require().Empty(res.Err)
				t.Require().Equal(false, res.Existed)
				t.Require().NotEmpty(res.Label)
			})
		}
	})

	runner.Run(t, "Test chaincode list again", func(t provider.T) {
		t.Severity(allure.CRITICAL)
		t.Description("Check work of chaincode installed list method. We expect there is also a new chaincode")
		resp, err := org0Cli.Chaincode.ChaincodeInstalledList(&chaincode.ChaincodeInstalledListParams{
			Context: ctx,
		}, auth)
		t.Require().NoError(err)
		t.Require().NotNil(resp.Payload.Result)
		t.Require().Len(resp.Payload.Result, installedCCLen+1)
	})

}
