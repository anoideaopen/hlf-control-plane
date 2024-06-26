// Code generated by go-swagger; DO NOT EDIT.

package lifecycle

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"gitlab.n-t.io/core/library/hlf-tool/hlf-control-plane/test/integration/models"
)

// LifecycleFullReader is a Reader for the LifecycleFull structure.
type LifecycleFullReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LifecycleFullReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLifecycleFullOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewLifecycleFullInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewLifecycleFullDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLifecycleFullOK creates a LifecycleFullOK with default headers values
func NewLifecycleFullOK() *LifecycleFullOK {
	return &LifecycleFullOK{}
}

/*
LifecycleFullOK describes a response with status code 200, with default header values.

A successful response.
*/
type LifecycleFullOK struct {
	Payload *models.ProtoLifecycleFullResponse
}

// IsSuccess returns true when this lifecycle full o k response has a 2xx status code
func (o *LifecycleFullOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this lifecycle full o k response has a 3xx status code
func (o *LifecycleFullOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this lifecycle full o k response has a 4xx status code
func (o *LifecycleFullOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this lifecycle full o k response has a 5xx status code
func (o *LifecycleFullOK) IsServerError() bool {
	return false
}

// IsCode returns true when this lifecycle full o k response a status code equal to that given
func (o *LifecycleFullOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the lifecycle full o k response
func (o *LifecycleFullOK) Code() int {
	return 200
}

func (o *LifecycleFullOK) Error() string {
	return fmt.Sprintf("[POST /v1/lifecycle/full][%d] lifecycleFullOK  %+v", 200, o.Payload)
}

func (o *LifecycleFullOK) String() string {
	return fmt.Sprintf("[POST /v1/lifecycle/full][%d] lifecycleFullOK  %+v", 200, o.Payload)
}

func (o *LifecycleFullOK) GetPayload() *models.ProtoLifecycleFullResponse {
	return o.Payload
}

func (o *LifecycleFullOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProtoLifecycleFullResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLifecycleFullInternalServerError creates a LifecycleFullInternalServerError with default headers values
func NewLifecycleFullInternalServerError() *LifecycleFullInternalServerError {
	return &LifecycleFullInternalServerError{}
}

/*
LifecycleFullInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type LifecycleFullInternalServerError struct {
	Payload *models.ProtoErrorResponse
}

// IsSuccess returns true when this lifecycle full internal server error response has a 2xx status code
func (o *LifecycleFullInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this lifecycle full internal server error response has a 3xx status code
func (o *LifecycleFullInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this lifecycle full internal server error response has a 4xx status code
func (o *LifecycleFullInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this lifecycle full internal server error response has a 5xx status code
func (o *LifecycleFullInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this lifecycle full internal server error response a status code equal to that given
func (o *LifecycleFullInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the lifecycle full internal server error response
func (o *LifecycleFullInternalServerError) Code() int {
	return 500
}

func (o *LifecycleFullInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/lifecycle/full][%d] lifecycleFullInternalServerError  %+v", 500, o.Payload)
}

func (o *LifecycleFullInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/lifecycle/full][%d] lifecycleFullInternalServerError  %+v", 500, o.Payload)
}

func (o *LifecycleFullInternalServerError) GetPayload() *models.ProtoErrorResponse {
	return o.Payload
}

func (o *LifecycleFullInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProtoErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLifecycleFullDefault creates a LifecycleFullDefault with default headers values
func NewLifecycleFullDefault(code int) *LifecycleFullDefault {
	return &LifecycleFullDefault{
		_statusCode: code,
	}
}

/*
LifecycleFullDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type LifecycleFullDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this lifecycle full default response has a 2xx status code
func (o *LifecycleFullDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this lifecycle full default response has a 3xx status code
func (o *LifecycleFullDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this lifecycle full default response has a 4xx status code
func (o *LifecycleFullDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this lifecycle full default response has a 5xx status code
func (o *LifecycleFullDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this lifecycle full default response a status code equal to that given
func (o *LifecycleFullDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the lifecycle full default response
func (o *LifecycleFullDefault) Code() int {
	return o._statusCode
}

func (o *LifecycleFullDefault) Error() string {
	return fmt.Sprintf("[POST /v1/lifecycle/full][%d] lifecycleFull default  %+v", o._statusCode, o.Payload)
}

func (o *LifecycleFullDefault) String() string {
	return fmt.Sprintf("[POST /v1/lifecycle/full][%d] lifecycleFull default  %+v", o._statusCode, o.Payload)
}

func (o *LifecycleFullDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *LifecycleFullDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
