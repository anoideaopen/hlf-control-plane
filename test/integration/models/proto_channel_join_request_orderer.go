// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ProtoChannelJoinRequestOrderer proto channel join request orderer
//
// swagger:model protoChannelJoinRequestOrderer
type ProtoChannelJoinRequestOrderer struct {

	// host
	Host string `json:"host,omitempty"`

	// port
	Port int64 `json:"port,omitempty"`
}

// Validate validates this proto channel join request orderer
func (m *ProtoChannelJoinRequestOrderer) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this proto channel join request orderer based on context it is used
func (m *ProtoChannelJoinRequestOrderer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProtoChannelJoinRequestOrderer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProtoChannelJoinRequestOrderer) UnmarshalBinary(b []byte) error {
	var res ProtoChannelJoinRequestOrderer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
