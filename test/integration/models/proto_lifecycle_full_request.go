// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ProtoLifecycleFullRequest proto lifecycle full request
//
// swagger:model protoLifecycleFullRequest
type ProtoLifecycleFullRequest struct {

	// chaincode label
	ChaincodeLabel string `json:"chaincodeLabel,omitempty"`

	// chaincode name
	ChaincodeName string `json:"chaincodeName,omitempty"`

	// channel name
	ChannelName string `json:"channelName,omitempty"`

	// commit force
	CommitForce bool `json:"commitForce,omitempty"`

	// init args
	InitArgs []string `json:"initArgs"`

	// init required
	InitRequired bool `json:"initRequired,omitempty"`

	// policy
	Policy string `json:"policy,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this proto lifecycle full request
func (m *ProtoLifecycleFullRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this proto lifecycle full request based on context it is used
func (m *ProtoLifecycleFullRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProtoLifecycleFullRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProtoLifecycleFullRequest) UnmarshalBinary(b []byte) error {
	var res ProtoLifecycleFullRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}