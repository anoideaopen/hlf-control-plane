// Code generated by go-swagger; DO NOT EDIT.

package discovery

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/atomyze-foundation/hlf-control-plane/test/integration/models"
)

// DiscoveryConfigReader is a Reader for the DiscoveryConfig structure.
type DiscoveryConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DiscoveryConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDiscoveryConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewDiscoveryConfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDiscoveryConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDiscoveryConfigOK creates a DiscoveryConfigOK with default headers values
func NewDiscoveryConfigOK() *DiscoveryConfigOK {
	return &DiscoveryConfigOK{}
}

/*
DiscoveryConfigOK describes a response with status code 200, with default header values.

A successful response.
*/
type DiscoveryConfigOK struct {
	Payload *models.ProtoDiscoveryConfigResponse
}

// IsSuccess returns true when this discovery config o k response has a 2xx status code
func (o *DiscoveryConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this discovery config o k response has a 3xx status code
func (o *DiscoveryConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this discovery config o k response has a 4xx status code
func (o *DiscoveryConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this discovery config o k response has a 5xx status code
func (o *DiscoveryConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this discovery config o k response a status code equal to that given
func (o *DiscoveryConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the discovery config o k response
func (o *DiscoveryConfigOK) Code() int {
	return 200
}

func (o *DiscoveryConfigOK) Error() string {
	return fmt.Sprintf("[POST /v1/discovery/{channelName}/config][%d] discoveryConfigOK  %+v", 200, o.Payload)
}

func (o *DiscoveryConfigOK) String() string {
	return fmt.Sprintf("[POST /v1/discovery/{channelName}/config][%d] discoveryConfigOK  %+v", 200, o.Payload)
}

func (o *DiscoveryConfigOK) GetPayload() *models.ProtoDiscoveryConfigResponse {
	return o.Payload
}

func (o *DiscoveryConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProtoDiscoveryConfigResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDiscoveryConfigInternalServerError creates a DiscoveryConfigInternalServerError with default headers values
func NewDiscoveryConfigInternalServerError() *DiscoveryConfigInternalServerError {
	return &DiscoveryConfigInternalServerError{}
}

/*
DiscoveryConfigInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type DiscoveryConfigInternalServerError struct {
	Payload *models.ProtoErrorResponse
}

// IsSuccess returns true when this discovery config internal server error response has a 2xx status code
func (o *DiscoveryConfigInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this discovery config internal server error response has a 3xx status code
func (o *DiscoveryConfigInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this discovery config internal server error response has a 4xx status code
func (o *DiscoveryConfigInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this discovery config internal server error response has a 5xx status code
func (o *DiscoveryConfigInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this discovery config internal server error response a status code equal to that given
func (o *DiscoveryConfigInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the discovery config internal server error response
func (o *DiscoveryConfigInternalServerError) Code() int {
	return 500
}

func (o *DiscoveryConfigInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/discovery/{channelName}/config][%d] discoveryConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *DiscoveryConfigInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/discovery/{channelName}/config][%d] discoveryConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *DiscoveryConfigInternalServerError) GetPayload() *models.ProtoErrorResponse {
	return o.Payload
}

func (o *DiscoveryConfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProtoErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDiscoveryConfigDefault creates a DiscoveryConfigDefault with default headers values
func NewDiscoveryConfigDefault(code int) *DiscoveryConfigDefault {
	return &DiscoveryConfigDefault{
		_statusCode: code,
	}
}

/*
DiscoveryConfigDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type DiscoveryConfigDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this discovery config default response has a 2xx status code
func (o *DiscoveryConfigDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this discovery config default response has a 3xx status code
func (o *DiscoveryConfigDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this discovery config default response has a 4xx status code
func (o *DiscoveryConfigDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this discovery config default response has a 5xx status code
func (o *DiscoveryConfigDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this discovery config default response a status code equal to that given
func (o *DiscoveryConfigDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the discovery config default response
func (o *DiscoveryConfigDefault) Code() int {
	return o._statusCode
}

func (o *DiscoveryConfigDefault) Error() string {
	return fmt.Sprintf("[POST /v1/discovery/{channelName}/config][%d] discoveryConfig default  %+v", o._statusCode, o.Payload)
}

func (o *DiscoveryConfigDefault) String() string {
	return fmt.Sprintf("[POST /v1/discovery/{channelName}/config][%d] discoveryConfig default  %+v", o._statusCode, o.Payload)
}

func (o *DiscoveryConfigDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *DiscoveryConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}