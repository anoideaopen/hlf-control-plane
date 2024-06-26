// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// ProtoConsensusType proto consensus type
//
// swagger:model protoConsensusType
type ProtoConsensusType string

func NewProtoConsensusType(value ProtoConsensusType) *ProtoConsensusType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ProtoConsensusType.
func (m ProtoConsensusType) Pointer() *ProtoConsensusType {
	return &m
}

const (

	// ProtoConsensusTypeCONSENSUSTYPEUNSPECIFIED captures enum value "CONSENSUS_TYPE_UNSPECIFIED"
	ProtoConsensusTypeCONSENSUSTYPEUNSPECIFIED ProtoConsensusType = "CONSENSUS_TYPE_UNSPECIFIED"

	// ProtoConsensusTypeCONSENSUSTYPERAFT captures enum value "CONSENSUS_TYPE_RAFT"
	ProtoConsensusTypeCONSENSUSTYPERAFT ProtoConsensusType = "CONSENSUS_TYPE_RAFT"

	// ProtoConsensusTypeCONSENSUSTYPEBFT captures enum value "CONSENSUS_TYPE_BFT"
	ProtoConsensusTypeCONSENSUSTYPEBFT ProtoConsensusType = "CONSENSUS_TYPE_BFT"
)

// for schema
var protoConsensusTypeEnum []interface{}

func init() {
	var res []ProtoConsensusType
	if err := json.Unmarshal([]byte(`["CONSENSUS_TYPE_UNSPECIFIED","CONSENSUS_TYPE_RAFT","CONSENSUS_TYPE_BFT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		protoConsensusTypeEnum = append(protoConsensusTypeEnum, v)
	}
}

func (m ProtoConsensusType) validateProtoConsensusTypeEnum(path, location string, value ProtoConsensusType) error {
	if err := validate.EnumCase(path, location, value, protoConsensusTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this proto consensus type
func (m ProtoConsensusType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateProtoConsensusTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this proto consensus type based on context it is used
func (m ProtoConsensusType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
