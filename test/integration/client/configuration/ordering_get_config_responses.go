// Code generated by go-swagger; DO NOT EDIT.

package configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"gitlab.n-t.io/core/library/hlf-tool/hlf-control-plane/test/integration/models"
)

// OrderingGetConfigReader is a Reader for the OrderingGetConfig structure.
type OrderingGetConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrderingGetConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrderingGetConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewOrderingGetConfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewOrderingGetConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewOrderingGetConfigOK creates a OrderingGetConfigOK with default headers values
func NewOrderingGetConfigOK() *OrderingGetConfigOK {
	return &OrderingGetConfigOK{}
}

/*
OrderingGetConfigOK describes a response with status code 200, with default header values.

A successful response.
*/
type OrderingGetConfigOK struct {
	Payload *models.ProtoConfigOrderingListResponse
}

// IsSuccess returns true when this ordering get config o k response has a 2xx status code
func (o *OrderingGetConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ordering get config o k response has a 3xx status code
func (o *OrderingGetConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ordering get config o k response has a 4xx status code
func (o *OrderingGetConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ordering get config o k response has a 5xx status code
func (o *OrderingGetConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ordering get config o k response a status code equal to that given
func (o *OrderingGetConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ordering get config o k response
func (o *OrderingGetConfigOK) Code() int {
	return 200
}

func (o *OrderingGetConfigOK) Error() string {
	return fmt.Sprintf("[GET /v1/config/{channelName}/ordering][%d] orderingGetConfigOK  %+v", 200, o.Payload)
}

func (o *OrderingGetConfigOK) String() string {
	return fmt.Sprintf("[GET /v1/config/{channelName}/ordering][%d] orderingGetConfigOK  %+v", 200, o.Payload)
}

func (o *OrderingGetConfigOK) GetPayload() *models.ProtoConfigOrderingListResponse {
	return o.Payload
}

func (o *OrderingGetConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProtoConfigOrderingListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrderingGetConfigInternalServerError creates a OrderingGetConfigInternalServerError with default headers values
func NewOrderingGetConfigInternalServerError() *OrderingGetConfigInternalServerError {
	return &OrderingGetConfigInternalServerError{}
}

/*
OrderingGetConfigInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type OrderingGetConfigInternalServerError struct {
	Payload *models.ProtoErrorResponse
}

// IsSuccess returns true when this ordering get config internal server error response has a 2xx status code
func (o *OrderingGetConfigInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ordering get config internal server error response has a 3xx status code
func (o *OrderingGetConfigInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ordering get config internal server error response has a 4xx status code
func (o *OrderingGetConfigInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this ordering get config internal server error response has a 5xx status code
func (o *OrderingGetConfigInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this ordering get config internal server error response a status code equal to that given
func (o *OrderingGetConfigInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the ordering get config internal server error response
func (o *OrderingGetConfigInternalServerError) Code() int {
	return 500
}

func (o *OrderingGetConfigInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/config/{channelName}/ordering][%d] orderingGetConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *OrderingGetConfigInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/config/{channelName}/ordering][%d] orderingGetConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *OrderingGetConfigInternalServerError) GetPayload() *models.ProtoErrorResponse {
	return o.Payload
}

func (o *OrderingGetConfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProtoErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrderingGetConfigDefault creates a OrderingGetConfigDefault with default headers values
func NewOrderingGetConfigDefault(code int) *OrderingGetConfigDefault {
	return &OrderingGetConfigDefault{
		_statusCode: code,
	}
}

/*
OrderingGetConfigDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type OrderingGetConfigDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this ordering get config default response has a 2xx status code
func (o *OrderingGetConfigDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ordering get config default response has a 3xx status code
func (o *OrderingGetConfigDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ordering get config default response has a 4xx status code
func (o *OrderingGetConfigDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ordering get config default response has a 5xx status code
func (o *OrderingGetConfigDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ordering get config default response a status code equal to that given
func (o *OrderingGetConfigDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ordering get config default response
func (o *OrderingGetConfigDefault) Code() int {
	return o._statusCode
}

func (o *OrderingGetConfigDefault) Error() string {
	return fmt.Sprintf("[GET /v1/config/{channelName}/ordering][%d] orderingGetConfig default  %+v", o._statusCode, o.Payload)
}

func (o *OrderingGetConfigDefault) String() string {
	return fmt.Sprintf("[GET /v1/config/{channelName}/ordering][%d] orderingGetConfig default  %+v", o._statusCode, o.Payload)
}

func (o *OrderingGetConfigDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *OrderingGetConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
