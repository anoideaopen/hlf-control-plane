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

// LifecycleCheckCommitReadinessReader is a Reader for the LifecycleCheckCommitReadiness structure.
type LifecycleCheckCommitReadinessReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LifecycleCheckCommitReadinessReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLifecycleCheckCommitReadinessOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewLifecycleCheckCommitReadinessInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewLifecycleCheckCommitReadinessDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLifecycleCheckCommitReadinessOK creates a LifecycleCheckCommitReadinessOK with default headers values
func NewLifecycleCheckCommitReadinessOK() *LifecycleCheckCommitReadinessOK {
	return &LifecycleCheckCommitReadinessOK{}
}

/*
LifecycleCheckCommitReadinessOK describes a response with status code 200, with default header values.

A successful response.
*/
type LifecycleCheckCommitReadinessOK struct {
	Payload *models.ProtoLifecycleCheckCommitReadinessResponse
}

// IsSuccess returns true when this lifecycle check commit readiness o k response has a 2xx status code
func (o *LifecycleCheckCommitReadinessOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this lifecycle check commit readiness o k response has a 3xx status code
func (o *LifecycleCheckCommitReadinessOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this lifecycle check commit readiness o k response has a 4xx status code
func (o *LifecycleCheckCommitReadinessOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this lifecycle check commit readiness o k response has a 5xx status code
func (o *LifecycleCheckCommitReadinessOK) IsServerError() bool {
	return false
}

// IsCode returns true when this lifecycle check commit readiness o k response a status code equal to that given
func (o *LifecycleCheckCommitReadinessOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the lifecycle check commit readiness o k response
func (o *LifecycleCheckCommitReadinessOK) Code() int {
	return 200
}

func (o *LifecycleCheckCommitReadinessOK) Error() string {
	return fmt.Sprintf("[POST /v1/lifecycle/check-commit-readiness][%d] lifecycleCheckCommitReadinessOK  %+v", 200, o.Payload)
}

func (o *LifecycleCheckCommitReadinessOK) String() string {
	return fmt.Sprintf("[POST /v1/lifecycle/check-commit-readiness][%d] lifecycleCheckCommitReadinessOK  %+v", 200, o.Payload)
}

func (o *LifecycleCheckCommitReadinessOK) GetPayload() *models.ProtoLifecycleCheckCommitReadinessResponse {
	return o.Payload
}

func (o *LifecycleCheckCommitReadinessOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProtoLifecycleCheckCommitReadinessResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLifecycleCheckCommitReadinessInternalServerError creates a LifecycleCheckCommitReadinessInternalServerError with default headers values
func NewLifecycleCheckCommitReadinessInternalServerError() *LifecycleCheckCommitReadinessInternalServerError {
	return &LifecycleCheckCommitReadinessInternalServerError{}
}

/*
LifecycleCheckCommitReadinessInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type LifecycleCheckCommitReadinessInternalServerError struct {
	Payload *models.ProtoErrorResponse
}

// IsSuccess returns true when this lifecycle check commit readiness internal server error response has a 2xx status code
func (o *LifecycleCheckCommitReadinessInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this lifecycle check commit readiness internal server error response has a 3xx status code
func (o *LifecycleCheckCommitReadinessInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this lifecycle check commit readiness internal server error response has a 4xx status code
func (o *LifecycleCheckCommitReadinessInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this lifecycle check commit readiness internal server error response has a 5xx status code
func (o *LifecycleCheckCommitReadinessInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this lifecycle check commit readiness internal server error response a status code equal to that given
func (o *LifecycleCheckCommitReadinessInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the lifecycle check commit readiness internal server error response
func (o *LifecycleCheckCommitReadinessInternalServerError) Code() int {
	return 500
}

func (o *LifecycleCheckCommitReadinessInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/lifecycle/check-commit-readiness][%d] lifecycleCheckCommitReadinessInternalServerError  %+v", 500, o.Payload)
}

func (o *LifecycleCheckCommitReadinessInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/lifecycle/check-commit-readiness][%d] lifecycleCheckCommitReadinessInternalServerError  %+v", 500, o.Payload)
}

func (o *LifecycleCheckCommitReadinessInternalServerError) GetPayload() *models.ProtoErrorResponse {
	return o.Payload
}

func (o *LifecycleCheckCommitReadinessInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProtoErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLifecycleCheckCommitReadinessDefault creates a LifecycleCheckCommitReadinessDefault with default headers values
func NewLifecycleCheckCommitReadinessDefault(code int) *LifecycleCheckCommitReadinessDefault {
	return &LifecycleCheckCommitReadinessDefault{
		_statusCode: code,
	}
}

/*
LifecycleCheckCommitReadinessDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type LifecycleCheckCommitReadinessDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this lifecycle check commit readiness default response has a 2xx status code
func (o *LifecycleCheckCommitReadinessDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this lifecycle check commit readiness default response has a 3xx status code
func (o *LifecycleCheckCommitReadinessDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this lifecycle check commit readiness default response has a 4xx status code
func (o *LifecycleCheckCommitReadinessDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this lifecycle check commit readiness default response has a 5xx status code
func (o *LifecycleCheckCommitReadinessDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this lifecycle check commit readiness default response a status code equal to that given
func (o *LifecycleCheckCommitReadinessDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the lifecycle check commit readiness default response
func (o *LifecycleCheckCommitReadinessDefault) Code() int {
	return o._statusCode
}

func (o *LifecycleCheckCommitReadinessDefault) Error() string {
	return fmt.Sprintf("[POST /v1/lifecycle/check-commit-readiness][%d] lifecycleCheckCommitReadiness default  %+v", o._statusCode, o.Payload)
}

func (o *LifecycleCheckCommitReadinessDefault) String() string {
	return fmt.Sprintf("[POST /v1/lifecycle/check-commit-readiness][%d] lifecycleCheckCommitReadiness default  %+v", o._statusCode, o.Payload)
}

func (o *LifecycleCheckCommitReadinessDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *LifecycleCheckCommitReadinessDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}