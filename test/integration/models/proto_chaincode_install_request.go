// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ProtoChaincodeInstallRequest InstallChaincodeRequest
//
// # Request for downloading and installation of chaincodes
//
// swagger:model protoChaincodeInstallRequest
type ProtoChaincodeInstallRequest struct {

	// URL link to package with chaincode
	AuthHeaders map[string]string `json:"authHeaders,omitempty"`

	// URL link to package with chaincode
	// Required: true
	Source *string `json:"source"`
}

// Validate validates this proto chaincode install request
func (m *ProtoChaincodeInstallRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProtoChaincodeInstallRequest) validateSource(formats strfmt.Registry) error {

	if err := validate.Required("source", "body", m.Source); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this proto chaincode install request based on context it is used
func (m *ProtoChaincodeInstallRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProtoChaincodeInstallRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProtoChaincodeInstallRequest) UnmarshalBinary(b []byte) error {
	var res ProtoChaincodeInstallRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
