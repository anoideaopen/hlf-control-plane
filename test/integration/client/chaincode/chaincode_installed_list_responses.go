// Code generated by go-swagger; DO NOT EDIT.

package chaincode

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/atomyze-foundation/hlf-control-plane/test/integration/models"
)

// ChaincodeInstalledListReader is a Reader for the ChaincodeInstalledList structure.
type ChaincodeInstalledListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChaincodeInstalledListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChaincodeInstalledListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewChaincodeInstalledListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewChaincodeInstalledListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewChaincodeInstalledListOK creates a ChaincodeInstalledListOK with default headers values
func NewChaincodeInstalledListOK() *ChaincodeInstalledListOK {
	return &ChaincodeInstalledListOK{}
}

/*
ChaincodeInstalledListOK describes a response with status code 200, with default header values.

A successful response.
*/
type ChaincodeInstalledListOK struct {
	Payload *models.ProtoChaincodeInstalledResponse
}

// IsSuccess returns true when this chaincode installed list o k response has a 2xx status code
func (o *ChaincodeInstalledListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this chaincode installed list o k response has a 3xx status code
func (o *ChaincodeInstalledListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this chaincode installed list o k response has a 4xx status code
func (o *ChaincodeInstalledListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this chaincode installed list o k response has a 5xx status code
func (o *ChaincodeInstalledListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this chaincode installed list o k response a status code equal to that given
func (o *ChaincodeInstalledListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the chaincode installed list o k response
func (o *ChaincodeInstalledListOK) Code() int {
	return 200
}

func (o *ChaincodeInstalledListOK) Error() string {
	return fmt.Sprintf("[GET /v1/chaincode/installed][%d] chaincodeInstalledListOK  %+v", 200, o.Payload)
}

func (o *ChaincodeInstalledListOK) String() string {
	return fmt.Sprintf("[GET /v1/chaincode/installed][%d] chaincodeInstalledListOK  %+v", 200, o.Payload)
}

func (o *ChaincodeInstalledListOK) GetPayload() *models.ProtoChaincodeInstalledResponse {
	return o.Payload
}

func (o *ChaincodeInstalledListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProtoChaincodeInstalledResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChaincodeInstalledListInternalServerError creates a ChaincodeInstalledListInternalServerError with default headers values
func NewChaincodeInstalledListInternalServerError() *ChaincodeInstalledListInternalServerError {
	return &ChaincodeInstalledListInternalServerError{}
}

/*
ChaincodeInstalledListInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ChaincodeInstalledListInternalServerError struct {
	Payload *models.ProtoErrorResponse
}

// IsSuccess returns true when this chaincode installed list internal server error response has a 2xx status code
func (o *ChaincodeInstalledListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this chaincode installed list internal server error response has a 3xx status code
func (o *ChaincodeInstalledListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this chaincode installed list internal server error response has a 4xx status code
func (o *ChaincodeInstalledListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this chaincode installed list internal server error response has a 5xx status code
func (o *ChaincodeInstalledListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this chaincode installed list internal server error response a status code equal to that given
func (o *ChaincodeInstalledListInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the chaincode installed list internal server error response
func (o *ChaincodeInstalledListInternalServerError) Code() int {
	return 500
}

func (o *ChaincodeInstalledListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/chaincode/installed][%d] chaincodeInstalledListInternalServerError  %+v", 500, o.Payload)
}

func (o *ChaincodeInstalledListInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/chaincode/installed][%d] chaincodeInstalledListInternalServerError  %+v", 500, o.Payload)
}

func (o *ChaincodeInstalledListInternalServerError) GetPayload() *models.ProtoErrorResponse {
	return o.Payload
}

func (o *ChaincodeInstalledListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProtoErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChaincodeInstalledListDefault creates a ChaincodeInstalledListDefault with default headers values
func NewChaincodeInstalledListDefault(code int) *ChaincodeInstalledListDefault {
	return &ChaincodeInstalledListDefault{
		_statusCode: code,
	}
}

/*
ChaincodeInstalledListDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ChaincodeInstalledListDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this chaincode installed list default response has a 2xx status code
func (o *ChaincodeInstalledListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this chaincode installed list default response has a 3xx status code
func (o *ChaincodeInstalledListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this chaincode installed list default response has a 4xx status code
func (o *ChaincodeInstalledListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this chaincode installed list default response has a 5xx status code
func (o *ChaincodeInstalledListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this chaincode installed list default response a status code equal to that given
func (o *ChaincodeInstalledListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the chaincode installed list default response
func (o *ChaincodeInstalledListDefault) Code() int {
	return o._statusCode
}

func (o *ChaincodeInstalledListDefault) Error() string {
	return fmt.Sprintf("[GET /v1/chaincode/installed][%d] chaincodeInstalledList default  %+v", o._statusCode, o.Payload)
}

func (o *ChaincodeInstalledListDefault) String() string {
	return fmt.Sprintf("[GET /v1/chaincode/installed][%d] chaincodeInstalledList default  %+v", o._statusCode, o.Payload)
}

func (o *ChaincodeInstalledListDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *ChaincodeInstalledListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
