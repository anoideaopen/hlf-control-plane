// Code generated by go-swagger; DO NOT EDIT.

package configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/atomyze-foundation/hlf-control-plane/test/integration/models"
)

// OrderingConfigUpdateReader is a Reader for the OrderingConfigUpdate structure.
type OrderingConfigUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrderingConfigUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrderingConfigUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewOrderingConfigUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewOrderingConfigUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewOrderingConfigUpdateOK creates a OrderingConfigUpdateOK with default headers values
func NewOrderingConfigUpdateOK() *OrderingConfigUpdateOK {
	return &OrderingConfigUpdateOK{}
}

/*
OrderingConfigUpdateOK describes a response with status code 200, with default header values.

A successful response.
*/
type OrderingConfigUpdateOK struct {
	Payload models.ProtoConfigOrderingUpdateResponse
}

// IsSuccess returns true when this ordering config update o k response has a 2xx status code
func (o *OrderingConfigUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ordering config update o k response has a 3xx status code
func (o *OrderingConfigUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ordering config update o k response has a 4xx status code
func (o *OrderingConfigUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ordering config update o k response has a 5xx status code
func (o *OrderingConfigUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ordering config update o k response a status code equal to that given
func (o *OrderingConfigUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ordering config update o k response
func (o *OrderingConfigUpdateOK) Code() int {
	return 200
}

func (o *OrderingConfigUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /v1/config/{channelName}/ordering][%d] orderingConfigUpdateOK  %+v", 200, o.Payload)
}

func (o *OrderingConfigUpdateOK) String() string {
	return fmt.Sprintf("[PUT /v1/config/{channelName}/ordering][%d] orderingConfigUpdateOK  %+v", 200, o.Payload)
}

func (o *OrderingConfigUpdateOK) GetPayload() models.ProtoConfigOrderingUpdateResponse {
	return o.Payload
}

func (o *OrderingConfigUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrderingConfigUpdateInternalServerError creates a OrderingConfigUpdateInternalServerError with default headers values
func NewOrderingConfigUpdateInternalServerError() *OrderingConfigUpdateInternalServerError {
	return &OrderingConfigUpdateInternalServerError{}
}

/*
OrderingConfigUpdateInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type OrderingConfigUpdateInternalServerError struct {
	Payload *models.ProtoErrorResponse
}

// IsSuccess returns true when this ordering config update internal server error response has a 2xx status code
func (o *OrderingConfigUpdateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ordering config update internal server error response has a 3xx status code
func (o *OrderingConfigUpdateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ordering config update internal server error response has a 4xx status code
func (o *OrderingConfigUpdateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this ordering config update internal server error response has a 5xx status code
func (o *OrderingConfigUpdateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this ordering config update internal server error response a status code equal to that given
func (o *OrderingConfigUpdateInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the ordering config update internal server error response
func (o *OrderingConfigUpdateInternalServerError) Code() int {
	return 500
}

func (o *OrderingConfigUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/config/{channelName}/ordering][%d] orderingConfigUpdateInternalServerError  %+v", 500, o.Payload)
}

func (o *OrderingConfigUpdateInternalServerError) String() string {
	return fmt.Sprintf("[PUT /v1/config/{channelName}/ordering][%d] orderingConfigUpdateInternalServerError  %+v", 500, o.Payload)
}

func (o *OrderingConfigUpdateInternalServerError) GetPayload() *models.ProtoErrorResponse {
	return o.Payload
}

func (o *OrderingConfigUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProtoErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrderingConfigUpdateDefault creates a OrderingConfigUpdateDefault with default headers values
func NewOrderingConfigUpdateDefault(code int) *OrderingConfigUpdateDefault {
	return &OrderingConfigUpdateDefault{
		_statusCode: code,
	}
}

/*
OrderingConfigUpdateDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type OrderingConfigUpdateDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this ordering config update default response has a 2xx status code
func (o *OrderingConfigUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ordering config update default response has a 3xx status code
func (o *OrderingConfigUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ordering config update default response has a 4xx status code
func (o *OrderingConfigUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ordering config update default response has a 5xx status code
func (o *OrderingConfigUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ordering config update default response a status code equal to that given
func (o *OrderingConfigUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ordering config update default response
func (o *OrderingConfigUpdateDefault) Code() int {
	return o._statusCode
}

func (o *OrderingConfigUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/config/{channelName}/ordering][%d] orderingConfigUpdate default  %+v", o._statusCode, o.Payload)
}

func (o *OrderingConfigUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /v1/config/{channelName}/ordering][%d] orderingConfigUpdate default  %+v", o._statusCode, o.Payload)
}

func (o *OrderingConfigUpdateDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *OrderingConfigUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
OrderingConfigUpdateBody Request and response for ordering modify method
swagger:model OrderingConfigUpdateBody
*/
type OrderingConfigUpdateBody struct {

	// orderer
	Orderer *models.ProtoOrderer `json:"orderer,omitempty"`
}

// Validate validates this ordering config update body
func (o *OrderingConfigUpdateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateOrderer(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *OrderingConfigUpdateBody) validateOrderer(formats strfmt.Registry) error {
	if swag.IsZero(o.Orderer) { // not required
		return nil
	}

	if o.Orderer != nil {
		if err := o.Orderer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "orderer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "orderer")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ordering config update body based on the context it is used
func (o *OrderingConfigUpdateBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateOrderer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *OrderingConfigUpdateBody) contextValidateOrderer(ctx context.Context, formats strfmt.Registry) error {

	if o.Orderer != nil {

		if swag.IsZero(o.Orderer) { // not required
			return nil
		}

		if err := o.Orderer.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "orderer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "orderer")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *OrderingConfigUpdateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *OrderingConfigUpdateBody) UnmarshalBinary(b []byte) error {
	var res OrderingConfigUpdateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
