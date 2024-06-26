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

// ProtoChannelCreateRequest Request and response for channel create method
//
// swagger:model protoChannelCreateRequest
type ProtoChannelCreateRequest struct {

	// channel name
	ChannelName string `json:"channelName,omitempty"`

	// organizations
	Organizations []*ProtoOrganization `json:"organizations"`
}

// Validate validates this proto channel create request
func (m *ProtoChannelCreateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrganizations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProtoChannelCreateRequest) validateOrganizations(formats strfmt.Registry) error {
	if swag.IsZero(m.Organizations) { // not required
		return nil
	}

	for i := 0; i < len(m.Organizations); i++ {
		if swag.IsZero(m.Organizations[i]) { // not required
			continue
		}

		if m.Organizations[i] != nil {
			if err := m.Organizations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("organizations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("organizations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this proto channel create request based on the context it is used
func (m *ProtoChannelCreateRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOrganizations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProtoChannelCreateRequest) contextValidateOrganizations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Organizations); i++ {

		if m.Organizations[i] != nil {

			if swag.IsZero(m.Organizations[i]) { // not required
				return nil
			}

			if err := m.Organizations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("organizations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("organizations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProtoChannelCreateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProtoChannelCreateRequest) UnmarshalBinary(b []byte) error {
	var res ProtoChannelCreateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
