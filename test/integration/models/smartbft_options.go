// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SmartbftOptions Options to be specified for all the smartbft nodes. These can be modified on a
// per-channel basis.
//
// swagger:model smartbftOptions
type SmartbftOptions struct {

	// collect timeout
	CollectTimeout string `json:"collectTimeout,omitempty"`

	// decisions per leader
	DecisionsPerLeader string `json:"decisionsPerLeader,omitempty"`

	// incoming message buffer size
	IncomingMessageBufferSize string `json:"incomingMessageBufferSize,omitempty"`

	// leader heartbeat count
	LeaderHeartbeatCount string `json:"leaderHeartbeatCount,omitempty"`

	// leader heartbeat timeout
	LeaderHeartbeatTimeout string `json:"leaderHeartbeatTimeout,omitempty"`

	// leader rotation
	LeaderRotation *OptionsRotation `json:"leaderRotation,omitempty"`

	// request auto remove timeout
	RequestAutoRemoveTimeout string `json:"requestAutoRemoveTimeout,omitempty"`

	// request batch max bytes
	RequestBatchMaxBytes string `json:"requestBatchMaxBytes,omitempty"`

	// request batch max count
	RequestBatchMaxCount string `json:"requestBatchMaxCount,omitempty"`

	// request batch max interval
	RequestBatchMaxInterval string `json:"requestBatchMaxInterval,omitempty"`

	// request complain timeout
	RequestComplainTimeout string `json:"requestComplainTimeout,omitempty"`

	// request forward timeout
	RequestForwardTimeout string `json:"requestForwardTimeout,omitempty"`

	// request max bytes
	RequestMaxBytes string `json:"requestMaxBytes,omitempty"`

	// request pool size
	RequestPoolSize string `json:"requestPoolSize,omitempty"`

	// request pool submit timeout
	RequestPoolSubmitTimeout string `json:"requestPoolSubmitTimeout,omitempty"`

	// speed up view change
	SpeedUpViewChange bool `json:"speedUpViewChange,omitempty"`

	// sync on start
	SyncOnStart bool `json:"syncOnStart,omitempty"`

	// view change resend interval
	ViewChangeResendInterval string `json:"viewChangeResendInterval,omitempty"`

	// view change timeout
	ViewChangeTimeout string `json:"viewChangeTimeout,omitempty"`
}

// Validate validates this smartbft options
func (m *SmartbftOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLeaderRotation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SmartbftOptions) validateLeaderRotation(formats strfmt.Registry) error {
	if swag.IsZero(m.LeaderRotation) { // not required
		return nil
	}

	if m.LeaderRotation != nil {
		if err := m.LeaderRotation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("leaderRotation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("leaderRotation")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this smartbft options based on the context it is used
func (m *SmartbftOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLeaderRotation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SmartbftOptions) contextValidateLeaderRotation(ctx context.Context, formats strfmt.Registry) error {

	if m.LeaderRotation != nil {

		if swag.IsZero(m.LeaderRotation) { // not required
			return nil
		}

		if err := m.LeaderRotation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("leaderRotation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("leaderRotation")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SmartbftOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SmartbftOptions) UnmarshalBinary(b []byte) error {
	var res SmartbftOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
