// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ProtoOrderingJoinResponse proto ordering join response
//
// swagger:model protoOrderingJoinResponse
type ProtoOrderingJoinResponse struct {

	// exists
	Exists bool `json:"exists,omitempty"`
}

// Validate validates this proto ordering join response
func (m *ProtoOrderingJoinResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this proto ordering join response based on context it is used
func (m *ProtoOrderingJoinResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProtoOrderingJoinResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProtoOrderingJoinResponse) UnmarshalBinary(b []byte) error {
	var res ProtoOrderingJoinResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
