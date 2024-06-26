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

// ProtoOrderingListInfoResponse proto ordering list info response
//
// swagger:model protoOrderingListInfoResponse
type ProtoOrderingListInfoResponse struct {

	// info
	Info []*OrderingListInfoResponseChannelInfo `json:"info"`
}

// Validate validates this proto ordering list info response
func (m *ProtoOrderingListInfoResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProtoOrderingListInfoResponse) validateInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.Info) { // not required
		return nil
	}

	for i := 0; i < len(m.Info); i++ {
		if swag.IsZero(m.Info[i]) { // not required
			continue
		}

		if m.Info[i] != nil {
			if err := m.Info[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("info" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this proto ordering list info response based on the context it is used
func (m *ProtoOrderingListInfoResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProtoOrderingListInfoResponse) contextValidateInfo(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Info); i++ {

		if m.Info[i] != nil {

			if swag.IsZero(m.Info[i]) { // not required
				return nil
			}

			if err := m.Info[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("info" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProtoOrderingListInfoResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProtoOrderingListInfoResponse) UnmarshalBinary(b []byte) error {
	var res ProtoOrderingListInfoResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
