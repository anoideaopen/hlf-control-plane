// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ChannelCreateResponseOrdererResult channel create response orderer result
//
// swagger:model ChannelCreateResponseOrdererResult
type ChannelCreateResponseOrdererResult struct {

	// host
	Host string `json:"host,omitempty"`

	// joined
	Joined bool `json:"joined,omitempty"`

	// port
	Port int64 `json:"port,omitempty"`
}

// Validate validates this channel create response orderer result
func (m *ChannelCreateResponseOrdererResult) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this channel create response orderer result based on context it is used
func (m *ChannelCreateResponseOrdererResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ChannelCreateResponseOrdererResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChannelCreateResponseOrdererResult) UnmarshalBinary(b []byte) error {
	var res ChannelCreateResponseOrdererResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
