// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MspFabricMSPConfig FabricMSPConfig collects all the configuration information for
// a Fabric MSP.
// Here we assume a default certificate validation policy, where
// any certificate signed by any of the listed rootCA certs would
// be considered as valid under this MSP.
// This MSP may or may not come with a signing identity. If it does,
// it can also issue signing identities. If it does not, it can only
// be used to validate and verify certificates.
//
// swagger:model mspFabricMSPConfig
type MspFabricMSPConfig struct {

	// Identity denoting the administrator of this MSP
	Admins []strfmt.Base64 `json:"admins"`

	// FabricCryptoConfig contains the configuration parameters
	// for the cryptographic algorithms used by this MSP
	CryptoConfig *MspFabricCryptoConfig `json:"cryptoConfig,omitempty"`

	// fabric_node_ous contains the configuration to distinguish clients from peers from orderers
	// based on the OUs.
	FabricNodeOus *MspFabricNodeOUs `json:"fabricNodeOus,omitempty"`

	// List of intermediate certificates trusted by this MSP;
	// they are used upon certificate validation as follows:
	// validation attempts to build a path from the certificate
	// to be validated (which is at one end of the path) and
	// one of the certs in the RootCerts field (which is at
	// the other end of the path). If the path is longer than
	// 2, certificates in the middle are searched within the
	// IntermediateCerts pool
	IntermediateCerts []strfmt.Base64 `json:"intermediateCerts"`

	// Name holds the identifier of the MSP; MSP identifier
	// is chosen by the application that governs this MSP.
	// For example, and assuming the default implementation of MSP,
	// that is X.509-based and considers a single Issuer,
	// this can refer to the Subject OU field or the Issuer OU field.
	Name string `json:"name,omitempty"`

	// OrganizationalUnitIdentifiers holds one or more
	// fabric organizational unit identifiers that belong to
	// this MSP configuration
	OrganizationalUnitIdentifiers []*MspFabricOUIdentifier `json:"organizationalUnitIdentifiers"`

	// Identity revocation list
	RevocationList []strfmt.Base64 `json:"revocationList"`

	// List of root certificates trusted by this MSP
	// they are used upon certificate validation (see
	// comment for IntermediateCerts below)
	RootCerts []strfmt.Base64 `json:"rootCerts"`

	// SigningIdentity holds information on the signing identity
	// this peer is to use, and which is to be imported by the
	// MSP defined before
	SigningIdentity *MspSigningIdentityInfo `json:"signingIdentity,omitempty"`

	// List of TLS intermediate certificates trusted by this MSP;
	// They are returned by GetTLSIntermediateCerts.
	TLSIntermediateCerts []strfmt.Base64 `json:"tlsIntermediateCerts"`

	// List of TLS root certificates trusted by this MSP.
	// They are returned by GetTLSRootCerts.
	TLSRootCerts []strfmt.Base64 `json:"tlsRootCerts"`
}

// Validate validates this msp fabric m s p config
func (m *MspFabricMSPConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCryptoConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFabricNodeOus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganizationalUnitIdentifiers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSigningIdentity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MspFabricMSPConfig) validateCryptoConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.CryptoConfig) { // not required
		return nil
	}

	if m.CryptoConfig != nil {
		if err := m.CryptoConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cryptoConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cryptoConfig")
			}
			return err
		}
	}

	return nil
}

func (m *MspFabricMSPConfig) validateFabricNodeOus(formats strfmt.Registry) error {
	if swag.IsZero(m.FabricNodeOus) { // not required
		return nil
	}

	if m.FabricNodeOus != nil {
		if err := m.FabricNodeOus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fabricNodeOus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fabricNodeOus")
			}
			return err
		}
	}

	return nil
}

func (m *MspFabricMSPConfig) validateOrganizationalUnitIdentifiers(formats strfmt.Registry) error {
	if swag.IsZero(m.OrganizationalUnitIdentifiers) { // not required
		return nil
	}

	for i := 0; i < len(m.OrganizationalUnitIdentifiers); i++ {
		if swag.IsZero(m.OrganizationalUnitIdentifiers[i]) { // not required
			continue
		}

		if m.OrganizationalUnitIdentifiers[i] != nil {
			if err := m.OrganizationalUnitIdentifiers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("organizationalUnitIdentifiers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("organizationalUnitIdentifiers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MspFabricMSPConfig) validateSigningIdentity(formats strfmt.Registry) error {
	if swag.IsZero(m.SigningIdentity) { // not required
		return nil
	}

	if m.SigningIdentity != nil {
		if err := m.SigningIdentity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("signingIdentity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("signingIdentity")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this msp fabric m s p config based on the context it is used
func (m *MspFabricMSPConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCryptoConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFabricNodeOus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOrganizationalUnitIdentifiers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSigningIdentity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MspFabricMSPConfig) contextValidateCryptoConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.CryptoConfig != nil {

		if swag.IsZero(m.CryptoConfig) { // not required
			return nil
		}

		if err := m.CryptoConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cryptoConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cryptoConfig")
			}
			return err
		}
	}

	return nil
}

func (m *MspFabricMSPConfig) contextValidateFabricNodeOus(ctx context.Context, formats strfmt.Registry) error {

	if m.FabricNodeOus != nil {

		if swag.IsZero(m.FabricNodeOus) { // not required
			return nil
		}

		if err := m.FabricNodeOus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fabricNodeOus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fabricNodeOus")
			}
			return err
		}
	}

	return nil
}

func (m *MspFabricMSPConfig) contextValidateOrganizationalUnitIdentifiers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.OrganizationalUnitIdentifiers); i++ {

		if m.OrganizationalUnitIdentifiers[i] != nil {

			if swag.IsZero(m.OrganizationalUnitIdentifiers[i]) { // not required
				return nil
			}

			if err := m.OrganizationalUnitIdentifiers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("organizationalUnitIdentifiers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("organizationalUnitIdentifiers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MspFabricMSPConfig) contextValidateSigningIdentity(ctx context.Context, formats strfmt.Registry) error {

	if m.SigningIdentity != nil {

		if swag.IsZero(m.SigningIdentity) { // not required
			return nil
		}

		if err := m.SigningIdentity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("signingIdentity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("signingIdentity")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MspFabricMSPConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MspFabricMSPConfig) UnmarshalBinary(b []byte) error {
	var res MspFabricMSPConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}