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

// LifecycleApprovedReader is a Reader for the LifecycleApproved structure.
type LifecycleApprovedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LifecycleApprovedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLifecycleApprovedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewLifecycleApprovedInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewLifecycleApprovedDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLifecycleApprovedOK creates a LifecycleApprovedOK with default headers values
func NewLifecycleApprovedOK() *LifecycleApprovedOK {
	return &LifecycleApprovedOK{}
}

/*
LifecycleApprovedOK describes a response with status code 200, with default header values.

A successful response.
*/
type LifecycleApprovedOK struct {
	Payload *models.ProtoLifecycleApprovedResponse
}

// IsSuccess returns true when this lifecycle approved o k response has a 2xx status code
func (o *LifecycleApprovedOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this lifecycle approved o k response has a 3xx status code
func (o *LifecycleApprovedOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this lifecycle approved o k response has a 4xx status code
func (o *LifecycleApprovedOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this lifecycle approved o k response has a 5xx status code
func (o *LifecycleApprovedOK) IsServerError() bool {
	return false
}

// IsCode returns true when this lifecycle approved o k response a status code equal to that given
func (o *LifecycleApprovedOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the lifecycle approved o k response
func (o *LifecycleApprovedOK) Code() int {
	return 200
}

func (o *LifecycleApprovedOK) Error() string {
	return fmt.Sprintf("[GET /v1/lifecycle/{channelName}/approved/{chaincodeName}][%d] lifecycleApprovedOK  %+v", 200, o.Payload)
}

func (o *LifecycleApprovedOK) String() string {
	return fmt.Sprintf("[GET /v1/lifecycle/{channelName}/approved/{chaincodeName}][%d] lifecycleApprovedOK  %+v", 200, o.Payload)
}

func (o *LifecycleApprovedOK) GetPayload() *models.ProtoLifecycleApprovedResponse {
	return o.Payload
}

func (o *LifecycleApprovedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProtoLifecycleApprovedResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLifecycleApprovedInternalServerError creates a LifecycleApprovedInternalServerError with default headers values
func NewLifecycleApprovedInternalServerError() *LifecycleApprovedInternalServerError {
	return &LifecycleApprovedInternalServerError{}
}

/*
LifecycleApprovedInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type LifecycleApprovedInternalServerError struct {
	Payload *models.ProtoErrorResponse
}

// IsSuccess returns true when this lifecycle approved internal server error response has a 2xx status code
func (o *LifecycleApprovedInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this lifecycle approved internal server error response has a 3xx status code
func (o *LifecycleApprovedInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this lifecycle approved internal server error response has a 4xx status code
func (o *LifecycleApprovedInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this lifecycle approved internal server error response has a 5xx status code
func (o *LifecycleApprovedInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this lifecycle approved internal server error response a status code equal to that given
func (o *LifecycleApprovedInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the lifecycle approved internal server error response
func (o *LifecycleApprovedInternalServerError) Code() int {
	return 500
}

func (o *LifecycleApprovedInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/lifecycle/{channelName}/approved/{chaincodeName}][%d] lifecycleApprovedInternalServerError  %+v", 500, o.Payload)
}

func (o *LifecycleApprovedInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/lifecycle/{channelName}/approved/{chaincodeName}][%d] lifecycleApprovedInternalServerError  %+v", 500, o.Payload)
}

func (o *LifecycleApprovedInternalServerError) GetPayload() *models.ProtoErrorResponse {
	return o.Payload
}

func (o *LifecycleApprovedInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProtoErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLifecycleApprovedDefault creates a LifecycleApprovedDefault with default headers values
func NewLifecycleApprovedDefault(code int) *LifecycleApprovedDefault {
	return &LifecycleApprovedDefault{
		_statusCode: code,
	}
}

/*
LifecycleApprovedDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type LifecycleApprovedDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this lifecycle approved default response has a 2xx status code
func (o *LifecycleApprovedDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this lifecycle approved default response has a 3xx status code
func (o *LifecycleApprovedDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this lifecycle approved default response has a 4xx status code
func (o *LifecycleApprovedDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this lifecycle approved default response has a 5xx status code
func (o *LifecycleApprovedDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this lifecycle approved default response a status code equal to that given
func (o *LifecycleApprovedDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the lifecycle approved default response
func (o *LifecycleApprovedDefault) Code() int {
	return o._statusCode
}

func (o *LifecycleApprovedDefault) Error() string {
	return fmt.Sprintf("[GET /v1/lifecycle/{channelName}/approved/{chaincodeName}][%d] lifecycleApproved default  %+v", o._statusCode, o.Payload)
}

func (o *LifecycleApprovedDefault) String() string {
	return fmt.Sprintf("[GET /v1/lifecycle/{channelName}/approved/{chaincodeName}][%d] lifecycleApproved default  %+v", o._statusCode, o.Payload)
}

func (o *LifecycleApprovedDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *LifecycleApprovedDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
