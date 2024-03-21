package integration

import (
	"context"
	"os"

	"github.com/jdxcode/netrc"

	"github.com/atomyze-foundation/hlf-control-plane/test/integration/client"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

const (
	channelName       = "fiat"
	controlPlaneToken = "test"
	ordererHost       = "test-orderer-001.org0"
	ordererPort       = int64(7050)
)

var (
	org0Cli *client.HlfControlPlane
	org1Cli *client.HlfControlPlane

	auth runtime.ClientAuthInfoWriter
	ctx  = context.Background()

	// settings for access nexus
	netRcContent  = os.Getenv("NETRC_CONTENT")
	netRcMachine  = "prod-nexus-001"
	netRcHost     = "https://prod-raw.cicd.prod.core.n-t.io"
	netRcUser     string
	netRcPassword string

	// settings for chaincode init
	ski1   = "c3711405a81e1a7a345b282ee9e924f834eff9a4cbbc7bf8be6cd84c98d512c2"
	ski2   = "3253b0b9788e8e900bf49c00606373aee2afbbdcbd46601c2f00d5eb7668c30f"
	issuer = "2d16MSkrG5BhHXSfykzpwi5nadHG1PFKR26kYtmn6RTZtLa5ta"
)

func init() {
	org0Transport := httptransport.New(os.Getenv("HLF_CONTROL_PLANE_API_ORG0"), "", nil)
	org0Cli = client.New(org0Transport, strfmt.Default)

	org1Transport := httptransport.New(os.Getenv("HLF_CONTROL_PLANE_API_ORG1"), "", nil)
	org1Cli = client.New(org1Transport, strfmt.Default)

	auth = httptransport.APIKeyAuth("Authorization", "header", controlPlaneToken)

	nc, err := netrc.ParseString(netRcContent)
	if err != nil {
		panic(err)
	}

	// setting for local development
	if m := os.Getenv("NETRC_MACHINE"); m != "" {
		netRcMachine = m
	}
	if h := os.Getenv("NETRC_HOST"); h != "" {
		netRcHost = h
	}

	m := nc.Machine(netRcMachine)
	if m != nil {
		netRcUser = m.Get("login")
		netRcPassword = m.Get("password")
	}
}
