//nolint:gomnd
package configtx

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"time"

	"github.com/golang/protobuf/proto" //nolint:staticcheck
	"github.com/hyperledger/fabric-config/configtx"
	"github.com/hyperledger/fabric-config/configtx/membership"
	"github.com/hyperledger/fabric-config/configtx/orderer"
	"github.com/hyperledger/fabric-protos-go/common"
	fbMsp "github.com/hyperledger/fabric-protos-go/msp"
	"github.com/hyperledger/fabric-protos-go/orderer/smartbft"
	"github.com/hyperledger/fabric/msp"
	"github.com/hyperledger/fabric/protoutil"
	pb "gitlab.n-t.io/core/library/hlf-tool/hlf-control-plane/proto"
)

func CreateGenesisBlock(req *pb.ChannelCreateRequest) ([]byte, error) {
	orgs, err := organisationWithMSP(req.GetOrganizations())
	if err != nil {
		return nil, err
	}

	cons, err := consenters(req.Organizations)
	if err != nil {
		return nil, err
	}

	c := configtx.Channel{
		Application: configtx.Application{
			Organizations: orgs,
			Capabilities:  []string{ChannelV2_0},
			Policies:      standardPoliciesApplication(),
			ModPolicy:     configtx.AdminsPolicyKey,
		},
		Orderer: configtx.Orderer{
			OrdererType:  ConsensusTypeSmartBFT,
			Addresses:    ordererAddresses(req.Organizations),
			BatchTimeout: 500 * time.Millisecond,
			BatchSize: orderer.BatchSize{
				MaxMessageCount:   500,
				AbsoluteMaxBytes:  99 * 1024 * 1024,
				PreferredMaxBytes: 2 * 1024 * 1024,
			},
			SmartBFT: &smartbft.ConfigMetadata{
				Consenters: cons,
				Options:    standartSmartbftOptions(),
			},
			Organizations: orgs,
			MaxChannels:   0,
			Capabilities:  []string{ChannelV2_0},
			Policies:      standardPoliciesOrderer(),
		},
		Capabilities: []string{ChannelV2_0},
		Consortium:   DefaultConsortiumValue,
		Policies:     standardPolicies(),
		ModPolicy:    configtx.AdminsPolicyKey,
	}

	genesisBlock, err := configtx.NewApplicationChannelGenesisBlock(c, req.ChannelName)
	if err != nil {
		return nil, fmt.Errorf("error new application channel genesis block, %w", err)
	}

	return protoutil.Marshal(genesisBlock)
}

func organisationWithMSP(orgsReq []*pb.Organization) (orgs []configtx.Organization, err error) {
	for _, oReq := range orgsReq {
		rootCert, err := parseCertificateListFromBytes([][]byte{oReq.RootCerts})
		if err != nil {
			return nil, err
		}
		tlsRootCert, err := parseCertificateListFromBytes([][]byte{oReq.TlsRootCerts})
		if err != nil {
			return nil, err
		}
		polisies := parsePolicies(oReq.Policies)
		o := configtx.Organization{
			Name:     oReq.Name,
			Policies: polisies,
			MSP: configtx.MSP{
				Name:      oReq.Id,
				RootCerts: rootCert,
				CryptoConfig: membership.CryptoConfig{
					SignatureHashFamily:            SHA2,
					IdentityIdentifierHashFunction: SHA256,
				},
				TLSRootCerts: tlsRootCert,
				NodeOUs: membership.NodeOUs{
					Enable: true,
					ClientOUIdentifier: membership.OUIdentifier{
						Certificate:                  rootCert[0],
						OrganizationalUnitIdentifier: RoleClient,
					},
					PeerOUIdentifier: membership.OUIdentifier{
						Certificate:                  rootCert[0],
						OrganizationalUnitIdentifier: RolePeer,
					},
					AdminOUIdentifier: membership.OUIdentifier{
						Certificate:                  rootCert[0],
						OrganizationalUnitIdentifier: RoleAdmin,
					},
					OrdererOUIdentifier: membership.OUIdentifier{
						Certificate:                  rootCert[0],
						OrganizationalUnitIdentifier: RoleOrderer,
					},
				},
			},
			ModPolicy: configtx.AdminsPolicyKey,
		}
		orgs = append(orgs, o)
	}

	return
}

func ordererAddresses(orgs []*pb.Organization) (address []string) {
	for _, org := range orgs {
		for _, o := range org.Orderers {
			address = append(address, fmt.Sprintf("%s:%d", o.Host, o.Port))
		}
	}

	return
}

func standartSmartbftOptions() *smartbft.Options {
	return &smartbft.Options{
		RequestBatchMaxCount:      100,
		RequestBatchMaxBytes:      10 * 1024 * 1024,
		RequestBatchMaxInterval:   "500ms",
		IncomingMessageBufferSize: 1000,
		RequestPoolSize:           100,
		RequestForwardTimeout:     "2s",
		RequestComplainTimeout:    "20s",
		RequestAutoRemoveTimeout:  "3m",
		ViewChangeResendInterval:  "5s",
		ViewChangeTimeout:         "20s",
		LeaderHeartbeatTimeout:    "1m",
		LeaderHeartbeatCount:      10,
		CollectTimeout:            "1s",
		SyncOnStart:               false,
		SpeedUpViewChange:         false,
		LeaderRotation:            1,
		DecisionsPerLeader:        0,
		RequestMaxBytes:           500 * 1024,
		RequestPoolSubmitTimeout:  "5s",
	}
}

func standardPolicies() map[string]configtx.Policy {
	return map[string]configtx.Policy{
		configtx.ReadersPolicyKey: {
			Type:      configtx.ImplicitMetaPolicyType,
			Rule:      AnyReaders,
			ModPolicy: configtx.AdminsPolicyKey,
		},
		configtx.WritersPolicyKey: {
			Type:      configtx.ImplicitMetaPolicyType,
			Rule:      AnyWriters,
			ModPolicy: configtx.AdminsPolicyKey,
		},
		configtx.AdminsPolicyKey: {
			Type:      configtx.ImplicitMetaPolicyType,
			Rule:      MajorityAdmins,
			ModPolicy: configtx.AdminsPolicyKey,
		},
	}
}

func standardPoliciesOrderer() map[string]configtx.Policy {
	return map[string]configtx.Policy{
		configtx.ReadersPolicyKey: {
			Type:      configtx.ImplicitMetaPolicyType,
			Rule:      AnyReaders,
			ModPolicy: configtx.AdminsPolicyKey,
		},
		configtx.WritersPolicyKey: {
			Type:      configtx.ImplicitMetaPolicyType,
			Rule:      AnyWriters,
			ModPolicy: configtx.AdminsPolicyKey,
		},
		configtx.AdminsPolicyKey: {
			Type:      configtx.ImplicitMetaPolicyType,
			Rule:      MajorityAdmins,
			ModPolicy: configtx.AdminsPolicyKey,
		},
		configtx.BlockValidationPolicyKey: {
			Type:      configtx.ImplicitOrdererPolicyType,
			Rule:      common.ImplicitOrdererPolicy_SMARTBFT.String(),
			ModPolicy: configtx.AdminsPolicyKey,
		},
	}
}

func standardPoliciesApplication() map[string]configtx.Policy {
	return map[string]configtx.Policy{
		configtx.LifecycleEndorsementPolicyKey: {
			Type:      configtx.ImplicitMetaPolicyType,
			Rule:      MajorityEndorsement,
			ModPolicy: configtx.AdminsPolicyKey,
		},
		configtx.EndorsementPolicyKey: {
			Type:      configtx.ImplicitMetaPolicyType,
			Rule:      MajorityEndorsement,
			ModPolicy: configtx.AdminsPolicyKey,
		},
		configtx.ReadersPolicyKey: {
			Type:      configtx.ImplicitMetaPolicyType,
			Rule:      AnyReaders,
			ModPolicy: configtx.AdminsPolicyKey,
		},
		configtx.WritersPolicyKey: {
			Type:      configtx.ImplicitMetaPolicyType,
			Rule:      AnyWriters,
			ModPolicy: configtx.AdminsPolicyKey,
		},
		configtx.AdminsPolicyKey: {
			Type:      configtx.ImplicitMetaPolicyType,
			Rule:      MajorityAdmins,
			ModPolicy: configtx.AdminsPolicyKey,
		},
	}
}

func parseCertificateListFromBytes(certs [][]byte) ([]*x509.Certificate, error) {
	certificateList := make([]*x509.Certificate, 0, len(certs))

	for _, cert := range certs {
		certificate, err := parseCertificateFromBytes(cert)
		if err != nil {
			return certificateList, err
		}

		certificateList = append(certificateList, certificate)
	}

	return certificateList, nil
}

func parseCertificateFromBytes(cert []byte) (*x509.Certificate, error) {
	pemBlock, _ := pem.Decode(cert)
	if pemBlock == nil {
		return &x509.Certificate{}, fmt.Errorf("no PEM data found in cert[% x]", cert)
	}

	certificate, err := x509.ParseCertificate(pemBlock.Bytes)
	if err != nil {
		return &x509.Certificate{}, err
	}

	return certificate, nil
}

func parsePolicies(polisies []*pb.Policy) map[string]configtx.Policy {
	pols := make(map[string]configtx.Policy)
	for _, p := range polisies {
		pols[p.Name] = configtx.Policy{
			Type:      p.Type,
			Rule:      p.Rule,
			ModPolicy: configtx.AdminsPolicyKey,
		}
	}

	return pols
}

func consenters(orgs []*pb.Organization) ([]*smartbft.Consenter, error) {
	var (
		cons []*smartbft.Consenter
		id   uint64
	)

	for _, org := range orgs {
		for _, o := range org.Orderers {
			id++

			ordID := o.Identity
			var sid fbMsp.SerializedIdentity
			if err := proto.Unmarshal(o.Identity, &sid); err != nil {
				ordID, err = msp.NewSerializedIdentity(org.Id, o.Identity)
				if err != nil {
					return nil, fmt.Errorf("marshal id: %w", err)
				}
			}

			cons = append(cons, &smartbft.Consenter{
				ConsenterId:   id,
				Host:          o.Host,
				Port:          o.Port,
				MspId:         org.Id,
				Identity:      ordID,
				ClientTlsCert: o.Cert,
				ServerTlsCert: o.Cert,
			})
		}
	}

	return cons, nil
}
