// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ProtoDiscoveryPeer proto discovery peer
//
// swagger:model protoDiscoveryPeer
type ProtoDiscoveryPeer struct {

	// block number
	BlockNumber string `json:"blockNumber,omitempty"`

	// cert
	Cert *DiscoveryPeerCertificate `json:"cert,omitempty"`

	// chaincodes
	Chaincodes []string `json:"chaincodes"`

	// host
	Host string `json:"host,omitempty"`

	// msp Id
	MspID string `json:"mspId,omitempty"`

	// port
	Port int32 `json:"port,omitempty"`
}

// Validate validates this proto discovery peer
func (m *ProtoDiscoveryPeer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCert(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProtoDiscoveryPeer) validateCert(formats strfmt.Registry) error {
	if swag.IsZero(m.Cert) { // not required
		return nil
	}

	if m.Cert != nil {
		if err := m.Cert.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cert")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cert")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this proto discovery peer based on the context it is used
func (m *ProtoDiscoveryPeer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCert(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProtoDiscoveryPeer) contextValidateCert(ctx context.Context, formats strfmt.Registry) error {

	if m.Cert != nil {

		if swag.IsZero(m.Cert) { // not required
			return nil
		}

		if err := m.Cert.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cert")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cert")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProtoDiscoveryPeer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProtoDiscoveryPeer) UnmarshalBinary(b []byte) error {
	var res ProtoDiscoveryPeer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
