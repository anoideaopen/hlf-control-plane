// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ProtoChannelJoinResponsePeerResult proto channel join response peer result
//
// swagger:model protoChannelJoinResponsePeerResult
type ProtoChannelJoinResponsePeerResult struct {

	// err
	Err string `json:"err,omitempty"`

	// existed
	Existed bool `json:"existed,omitempty"`

	// peer
	Peer string `json:"peer,omitempty"`
}

// Validate validates this proto channel join response peer result
func (m *ProtoChannelJoinResponsePeerResult) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this proto channel join response peer result based on context it is used
func (m *ProtoChannelJoinResponsePeerResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProtoChannelJoinResponsePeerResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProtoChannelJoinResponsePeerResult) UnmarshalBinary(b []byte) error {
	var res ProtoChannelJoinResponsePeerResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
