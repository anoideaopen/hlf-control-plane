package config

import (
	"fmt"

	"github.com/atomyze-foundation/hlf-control-plane/system/pkcs11"
	"github.com/hyperledger/fabric/bccsp/factory"
	"github.com/hyperledger/fabric/cmd/common/signer"
	"github.com/hyperledger/fabric/protoutil"
	"go.uber.org/zap"
)

const dummyMsg = "dummyMsg"

type Identity struct {
	Cert  string              `yaml:"cert"`
	Key   string              `yaml:"key"`
	BCCSP factory.FactoryOpts `yaml:"bccsp"`
	MspID string              `yaml:"mspId"`
}

func (i *Identity) Load(log *zap.Logger, mspID string) (protoutil.Signer, error) {
	var (
		id  protoutil.Signer
		err error
	)
	if i.MspID != "" {
		mspID = i.MspID
	}
	if i.Key != "" {
		// create signing identity
		id, err = signer.NewSigner(signer.Config{
			MSPID:        mspID,
			IdentityPath: i.Cert,
			KeyPath:      i.Key,
		})
		if err != nil {
			return nil, fmt.Errorf("get cert-based identity: %w", err)
		}
		log.Debug("id added", zap.String("mspId", mspID))
	} else {
		id, err = pkcs11.NewPKCS11Signer(log, mspID, i.Cert, &i.BCCSP)
		if err != nil {
			return nil, fmt.Errorf("get pkcs11 based identity: %w", err)
		}
	}

	// check identity to sign
	if _, err = id.Sign([]byte(dummyMsg)); err != nil {
		return nil, fmt.Errorf("check identity ability to sign: %w", err)
	}
	return id, nil
}
