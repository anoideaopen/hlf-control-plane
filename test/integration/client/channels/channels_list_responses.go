// Code generated by go-swagger; DO NOT EDIT.

package channels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"gitlab.n-t.io/core/library/hlf-tool/hlf-control-plane/test/integration/models"
)

// ChannelsListReader is a Reader for the ChannelsList structure.
type ChannelsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChannelsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChannelsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewChannelsListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewChannelsListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewChannelsListOK creates a ChannelsListOK with default headers values
func NewChannelsListOK() *ChannelsListOK {
	return &ChannelsListOK{}
}

/*
ChannelsListOK describes a response with status code 200, with default header values.

A successful response.
*/
type ChannelsListOK struct {
	Payload *models.ProtoChannelJoinedResponse
}

// IsSuccess returns true when this channels list o k response has a 2xx status code
func (o *ChannelsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this channels list o k response has a 3xx status code
func (o *ChannelsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this channels list o k response has a 4xx status code
func (o *ChannelsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this channels list o k response has a 5xx status code
func (o *ChannelsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this channels list o k response a status code equal to that given
func (o *ChannelsListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the channels list o k response
func (o *ChannelsListOK) Code() int {
	return 200
}

func (o *ChannelsListOK) Error() string {
	return fmt.Sprintf("[GET /v1/channel/joined][%d] channelsListOK  %+v", 200, o.Payload)
}

func (o *ChannelsListOK) String() string {
	return fmt.Sprintf("[GET /v1/channel/joined][%d] channelsListOK  %+v", 200, o.Payload)
}

func (o *ChannelsListOK) GetPayload() *models.ProtoChannelJoinedResponse {
	return o.Payload
}

func (o *ChannelsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProtoChannelJoinedResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChannelsListInternalServerError creates a ChannelsListInternalServerError with default headers values
func NewChannelsListInternalServerError() *ChannelsListInternalServerError {
	return &ChannelsListInternalServerError{}
}

/*
ChannelsListInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ChannelsListInternalServerError struct {
	Payload *models.ProtoErrorResponse
}

// IsSuccess returns true when this channels list internal server error response has a 2xx status code
func (o *ChannelsListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this channels list internal server error response has a 3xx status code
func (o *ChannelsListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this channels list internal server error response has a 4xx status code
func (o *ChannelsListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this channels list internal server error response has a 5xx status code
func (o *ChannelsListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this channels list internal server error response a status code equal to that given
func (o *ChannelsListInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the channels list internal server error response
func (o *ChannelsListInternalServerError) Code() int {
	return 500
}

func (o *ChannelsListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/channel/joined][%d] channelsListInternalServerError  %+v", 500, o.Payload)
}

func (o *ChannelsListInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/channel/joined][%d] channelsListInternalServerError  %+v", 500, o.Payload)
}

func (o *ChannelsListInternalServerError) GetPayload() *models.ProtoErrorResponse {
	return o.Payload
}

func (o *ChannelsListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProtoErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChannelsListDefault creates a ChannelsListDefault with default headers values
func NewChannelsListDefault(code int) *ChannelsListDefault {
	return &ChannelsListDefault{
		_statusCode: code,
	}
}

/*
ChannelsListDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ChannelsListDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this channels list default response has a 2xx status code
func (o *ChannelsListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this channels list default response has a 3xx status code
func (o *ChannelsListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this channels list default response has a 4xx status code
func (o *ChannelsListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this channels list default response has a 5xx status code
func (o *ChannelsListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this channels list default response a status code equal to that given
func (o *ChannelsListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the channels list default response
func (o *ChannelsListDefault) Code() int {
	return o._statusCode
}

func (o *ChannelsListDefault) Error() string {
	return fmt.Sprintf("[GET /v1/channel/joined][%d] channelsList default  %+v", o._statusCode, o.Payload)
}

func (o *ChannelsListDefault) String() string {
	return fmt.Sprintf("[GET /v1/channel/joined][%d] channelsList default  %+v", o._statusCode, o.Payload)
}

func (o *ChannelsListDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *ChannelsListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
