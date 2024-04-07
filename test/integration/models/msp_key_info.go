// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MspKeyInfo KeyInfo represents a (secret) key that is either already stored
// in the bccsp/keystore or key material to be imported to the
// bccsp key-store. In later versions it may contain also a
// keystore identifier
//
// swagger:model mspKeyInfo
type MspKeyInfo struct {

	// Identifier of the key inside the default keystore; this for
	// the case of Software BCCSP as well as the HSM BCCSP would be
	// the SKI of the key
	KeyIdentifier string `json:"keyIdentifier,omitempty"`

	// KeyMaterial (optional) for the key to be imported; this is
	// properly encoded key bytes, prefixed by the type of the key
	// Format: byte
	KeyMaterial strfmt.Base64 `json:"keyMaterial,omitempty"`
}

// Validate validates this msp key info
func (m *MspKeyInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this msp key info based on context it is used
func (m *MspKeyInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MspKeyInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MspKeyInfo) UnmarshalBinary(b []byte) error {
	var res MspKeyInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}