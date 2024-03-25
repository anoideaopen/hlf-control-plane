// Code generated by go-swagger; DO NOT EDIT.

package ordering

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/atomyze-foundation/hlf-control-plane/test/integration/models"
)

// OrderingJoinReader is a Reader for the OrderingJoin structure.
type OrderingJoinReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrderingJoinReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrderingJoinOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewOrderingJoinInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewOrderingJoinDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewOrderingJoinOK creates a OrderingJoinOK with default headers values
func NewOrderingJoinOK() *OrderingJoinOK {
	return &OrderingJoinOK{}
}

/*
OrderingJoinOK describes a response with status code 200, with default header values.

A successful response.
*/
type OrderingJoinOK struct {
	Payload *models.ProtoOrderingJoinResponse
}

// IsSuccess returns true when this ordering join o k response has a 2xx status code
func (o *OrderingJoinOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ordering join o k response has a 3xx status code
func (o *OrderingJoinOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ordering join o k response has a 4xx status code
func (o *OrderingJoinOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ordering join o k response has a 5xx status code
func (o *OrderingJoinOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ordering join o k response a status code equal to that given
func (o *OrderingJoinOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ordering join o k response
func (o *OrderingJoinOK) Code() int {
	return 200
}

func (o *OrderingJoinOK) Error() string {
	return fmt.Sprintf("[POST /v1/ordering/{channelName}/join][%d] orderingJoinOK  %+v", 200, o.Payload)
}

func (o *OrderingJoinOK) String() string {
	return fmt.Sprintf("[POST /v1/ordering/{channelName}/join][%d] orderingJoinOK  %+v", 200, o.Payload)
}

func (o *OrderingJoinOK) GetPayload() *models.ProtoOrderingJoinResponse {
	return o.Payload
}

func (o *OrderingJoinOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProtoOrderingJoinResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrderingJoinInternalServerError creates a OrderingJoinInternalServerError with default headers values
func NewOrderingJoinInternalServerError() *OrderingJoinInternalServerError {
	return &OrderingJoinInternalServerError{}
}

/*
OrderingJoinInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type OrderingJoinInternalServerError struct {
	Payload *models.ProtoErrorResponse
}

// IsSuccess returns true when this ordering join internal server error response has a 2xx status code
func (o *OrderingJoinInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ordering join internal server error response has a 3xx status code
func (o *OrderingJoinInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ordering join internal server error response has a 4xx status code
func (o *OrderingJoinInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this ordering join internal server error response has a 5xx status code
func (o *OrderingJoinInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this ordering join internal server error response a status code equal to that given
func (o *OrderingJoinInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the ordering join internal server error response
func (o *OrderingJoinInternalServerError) Code() int {
	return 500
}

func (o *OrderingJoinInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/ordering/{channelName}/join][%d] orderingJoinInternalServerError  %+v", 500, o.Payload)
}

func (o *OrderingJoinInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/ordering/{channelName}/join][%d] orderingJoinInternalServerError  %+v", 500, o.Payload)
}

func (o *OrderingJoinInternalServerError) GetPayload() *models.ProtoErrorResponse {
	return o.Payload
}

func (o *OrderingJoinInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProtoErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrderingJoinDefault creates a OrderingJoinDefault with default headers values
func NewOrderingJoinDefault(code int) *OrderingJoinDefault {
	return &OrderingJoinDefault{
		_statusCode: code,
	}
}

/*
OrderingJoinDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type OrderingJoinDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this ordering join default response has a 2xx status code
func (o *OrderingJoinDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ordering join default response has a 3xx status code
func (o *OrderingJoinDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ordering join default response has a 4xx status code
func (o *OrderingJoinDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ordering join default response has a 5xx status code
func (o *OrderingJoinDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ordering join default response a status code equal to that given
func (o *OrderingJoinDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ordering join default response
func (o *OrderingJoinDefault) Code() int {
	return o._statusCode
}

func (o *OrderingJoinDefault) Error() string {
	return fmt.Sprintf("[POST /v1/ordering/{channelName}/join][%d] orderingJoin default  %+v", o._statusCode, o.Payload)
}

func (o *OrderingJoinDefault) String() string {
	return fmt.Sprintf("[POST /v1/ordering/{channelName}/join][%d] orderingJoin default  %+v", o._statusCode, o.Payload)
}

func (o *OrderingJoinDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *OrderingJoinDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
OrderingJoinBody ordering join body
swagger:model OrderingJoinBody
*/
type OrderingJoinBody struct {

	// host
	Host string `json:"host,omitempty"`

	// joined orderer
	JoinedOrderer []*models.OrderingRequestJoined `json:"joinedOrderer"`

	// port
	Port int64 `json:"port,omitempty"`
}

// Validate validates this ordering join body
func (o *OrderingJoinBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateJoinedOrderer(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *OrderingJoinBody) validateJoinedOrderer(formats strfmt.Registry) error {
	if swag.IsZero(o.JoinedOrderer) { // not required
		return nil
	}

	for i := 0; i < len(o.JoinedOrderer); i++ {
		if swag.IsZero(o.JoinedOrderer[i]) { // not required
			continue
		}

		if o.JoinedOrderer[i] != nil {
			if err := o.JoinedOrderer[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "joinedOrderer" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("body" + "." + "joinedOrderer" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this ordering join body based on the context it is used
func (o *OrderingJoinBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateJoinedOrderer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *OrderingJoinBody) contextValidateJoinedOrderer(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.JoinedOrderer); i++ {

		if o.JoinedOrderer[i] != nil {

			if swag.IsZero(o.JoinedOrderer[i]) { // not required
				return nil
			}

			if err := o.JoinedOrderer[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "joinedOrderer" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("body" + "." + "joinedOrderer" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *OrderingJoinBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *OrderingJoinBody) UnmarshalBinary(b []byte) error {
	var res OrderingJoinBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}