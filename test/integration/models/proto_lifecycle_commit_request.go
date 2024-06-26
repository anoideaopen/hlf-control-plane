// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ProtoLifecycleCommitRequest proto lifecycle commit request
//
// swagger:model protoLifecycleCommitRequest
type ProtoLifecycleCommitRequest struct {

	// chaincode name
	ChaincodeName string `json:"chaincodeName,omitempty"`

	// channel name
	ChannelName string `json:"channelName,omitempty"`

	// init required
	InitRequired bool `json:"initRequired,omitempty"`

	// policy
	Policy string `json:"policy,omitempty"`

	// sequence
	Sequence string `json:"sequence,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this proto lifecycle commit request
func (m *ProtoLifecycleCommitRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this proto lifecycle commit request based on context it is used
func (m *ProtoLifecycleCommitRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProtoLifecycleCommitRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProtoLifecycleCommitRequest) UnmarshalBinary(b []byte) error {
	var res ProtoLifecycleCommitRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
