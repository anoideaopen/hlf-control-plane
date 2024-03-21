// Code generated by go-swagger; DO NOT EDIT.

package chaincode

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"gitlab.n-t.io/core/library/hlf-tool/hlf-control-plane/test/integration/models"
)

// ChaincodeInstallExternalReader is a Reader for the ChaincodeInstallExternal structure.
type ChaincodeInstallExternalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChaincodeInstallExternalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChaincodeInstallExternalOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewChaincodeInstallExternalInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewChaincodeInstallExternalDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewChaincodeInstallExternalOK creates a ChaincodeInstallExternalOK with default headers values
func NewChaincodeInstallExternalOK() *ChaincodeInstallExternalOK {
	return &ChaincodeInstallExternalOK{}
}

/*
ChaincodeInstallExternalOK describes a response with status code 200, with default header values.

A successful response.
*/
type ChaincodeInstallExternalOK struct {
	Payload *models.ProtoChaincodeInstallResponse
}

// IsSuccess returns true when this chaincode install external o k response has a 2xx status code
func (o *ChaincodeInstallExternalOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this chaincode install external o k response has a 3xx status code
func (o *ChaincodeInstallExternalOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this chaincode install external o k response has a 4xx status code
func (o *ChaincodeInstallExternalOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this chaincode install external o k response has a 5xx status code
func (o *ChaincodeInstallExternalOK) IsServerError() bool {
	return false
}

// IsCode returns true when this chaincode install external o k response a status code equal to that given
func (o *ChaincodeInstallExternalOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the chaincode install external o k response
func (o *ChaincodeInstallExternalOK) Code() int {
	return 200
}

func (o *ChaincodeInstallExternalOK) Error() string {
	return fmt.Sprintf("[POST /v1/chaincode/install-external][%d] chaincodeInstallExternalOK  %+v", 200, o.Payload)
}

func (o *ChaincodeInstallExternalOK) String() string {
	return fmt.Sprintf("[POST /v1/chaincode/install-external][%d] chaincodeInstallExternalOK  %+v", 200, o.Payload)
}

func (o *ChaincodeInstallExternalOK) GetPayload() *models.ProtoChaincodeInstallResponse {
	return o.Payload
}

func (o *ChaincodeInstallExternalOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProtoChaincodeInstallResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChaincodeInstallExternalInternalServerError creates a ChaincodeInstallExternalInternalServerError with default headers values
func NewChaincodeInstallExternalInternalServerError() *ChaincodeInstallExternalInternalServerError {
	return &ChaincodeInstallExternalInternalServerError{}
}

/*
ChaincodeInstallExternalInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ChaincodeInstallExternalInternalServerError struct {
	Payload *models.ProtoErrorResponse
}

// IsSuccess returns true when this chaincode install external internal server error response has a 2xx status code
func (o *ChaincodeInstallExternalInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this chaincode install external internal server error response has a 3xx status code
func (o *ChaincodeInstallExternalInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this chaincode install external internal server error response has a 4xx status code
func (o *ChaincodeInstallExternalInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this chaincode install external internal server error response has a 5xx status code
func (o *ChaincodeInstallExternalInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this chaincode install external internal server error response a status code equal to that given
func (o *ChaincodeInstallExternalInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the chaincode install external internal server error response
func (o *ChaincodeInstallExternalInternalServerError) Code() int {
	return 500
}

func (o *ChaincodeInstallExternalInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/chaincode/install-external][%d] chaincodeInstallExternalInternalServerError  %+v", 500, o.Payload)
}

func (o *ChaincodeInstallExternalInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/chaincode/install-external][%d] chaincodeInstallExternalInternalServerError  %+v", 500, o.Payload)
}

func (o *ChaincodeInstallExternalInternalServerError) GetPayload() *models.ProtoErrorResponse {
	return o.Payload
}

func (o *ChaincodeInstallExternalInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProtoErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChaincodeInstallExternalDefault creates a ChaincodeInstallExternalDefault with default headers values
func NewChaincodeInstallExternalDefault(code int) *ChaincodeInstallExternalDefault {
	return &ChaincodeInstallExternalDefault{
		_statusCode: code,
	}
}

/*
ChaincodeInstallExternalDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ChaincodeInstallExternalDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this chaincode install external default response has a 2xx status code
func (o *ChaincodeInstallExternalDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this chaincode install external default response has a 3xx status code
func (o *ChaincodeInstallExternalDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this chaincode install external default response has a 4xx status code
func (o *ChaincodeInstallExternalDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this chaincode install external default response has a 5xx status code
func (o *ChaincodeInstallExternalDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this chaincode install external default response a status code equal to that given
func (o *ChaincodeInstallExternalDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the chaincode install external default response
func (o *ChaincodeInstallExternalDefault) Code() int {
	return o._statusCode
}

func (o *ChaincodeInstallExternalDefault) Error() string {
	return fmt.Sprintf("[POST /v1/chaincode/install-external][%d] chaincodeInstallExternal default  %+v", o._statusCode, o.Payload)
}

func (o *ChaincodeInstallExternalDefault) String() string {
	return fmt.Sprintf("[POST /v1/chaincode/install-external][%d] chaincodeInstallExternal default  %+v", o._statusCode, o.Payload)
}

func (o *ChaincodeInstallExternalDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *ChaincodeInstallExternalDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
