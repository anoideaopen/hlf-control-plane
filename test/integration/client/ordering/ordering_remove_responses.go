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

	"gitlab.n-t.io/core/library/hlf-tool/hlf-control-plane/test/integration/models"
)

// OrderingRemoveReader is a Reader for the OrderingRemove structure.
type OrderingRemoveReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrderingRemoveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrderingRemoveOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewOrderingRemoveInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewOrderingRemoveDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewOrderingRemoveOK creates a OrderingRemoveOK with default headers values
func NewOrderingRemoveOK() *OrderingRemoveOK {
	return &OrderingRemoveOK{}
}

/*
OrderingRemoveOK describes a response with status code 200, with default header values.

A successful response.
*/
type OrderingRemoveOK struct {
	Payload models.ProtoOrderingRemoveResponse
}

// IsSuccess returns true when this ordering remove o k response has a 2xx status code
func (o *OrderingRemoveOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ordering remove o k response has a 3xx status code
func (o *OrderingRemoveOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ordering remove o k response has a 4xx status code
func (o *OrderingRemoveOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ordering remove o k response has a 5xx status code
func (o *OrderingRemoveOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ordering remove o k response a status code equal to that given
func (o *OrderingRemoveOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ordering remove o k response
func (o *OrderingRemoveOK) Code() int {
	return 200
}

func (o *OrderingRemoveOK) Error() string {
	return fmt.Sprintf("[POST /v1/ordering/{channelName}/remove][%d] orderingRemoveOK  %+v", 200, o.Payload)
}

func (o *OrderingRemoveOK) String() string {
	return fmt.Sprintf("[POST /v1/ordering/{channelName}/remove][%d] orderingRemoveOK  %+v", 200, o.Payload)
}

func (o *OrderingRemoveOK) GetPayload() models.ProtoOrderingRemoveResponse {
	return o.Payload
}

func (o *OrderingRemoveOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrderingRemoveInternalServerError creates a OrderingRemoveInternalServerError with default headers values
func NewOrderingRemoveInternalServerError() *OrderingRemoveInternalServerError {
	return &OrderingRemoveInternalServerError{}
}

/*
OrderingRemoveInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type OrderingRemoveInternalServerError struct {
	Payload *models.ProtoErrorResponse
}

// IsSuccess returns true when this ordering remove internal server error response has a 2xx status code
func (o *OrderingRemoveInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ordering remove internal server error response has a 3xx status code
func (o *OrderingRemoveInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ordering remove internal server error response has a 4xx status code
func (o *OrderingRemoveInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this ordering remove internal server error response has a 5xx status code
func (o *OrderingRemoveInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this ordering remove internal server error response a status code equal to that given
func (o *OrderingRemoveInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the ordering remove internal server error response
func (o *OrderingRemoveInternalServerError) Code() int {
	return 500
}

func (o *OrderingRemoveInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/ordering/{channelName}/remove][%d] orderingRemoveInternalServerError  %+v", 500, o.Payload)
}

func (o *OrderingRemoveInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/ordering/{channelName}/remove][%d] orderingRemoveInternalServerError  %+v", 500, o.Payload)
}

func (o *OrderingRemoveInternalServerError) GetPayload() *models.ProtoErrorResponse {
	return o.Payload
}

func (o *OrderingRemoveInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProtoErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrderingRemoveDefault creates a OrderingRemoveDefault with default headers values
func NewOrderingRemoveDefault(code int) *OrderingRemoveDefault {
	return &OrderingRemoveDefault{
		_statusCode: code,
	}
}

/*
OrderingRemoveDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type OrderingRemoveDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this ordering remove default response has a 2xx status code
func (o *OrderingRemoveDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ordering remove default response has a 3xx status code
func (o *OrderingRemoveDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ordering remove default response has a 4xx status code
func (o *OrderingRemoveDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ordering remove default response has a 5xx status code
func (o *OrderingRemoveDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ordering remove default response a status code equal to that given
func (o *OrderingRemoveDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ordering remove default response
func (o *OrderingRemoveDefault) Code() int {
	return o._statusCode
}

func (o *OrderingRemoveDefault) Error() string {
	return fmt.Sprintf("[POST /v1/ordering/{channelName}/remove][%d] orderingRemove default  %+v", o._statusCode, o.Payload)
}

func (o *OrderingRemoveDefault) String() string {
	return fmt.Sprintf("[POST /v1/ordering/{channelName}/remove][%d] orderingRemove default  %+v", o._statusCode, o.Payload)
}

func (o *OrderingRemoveDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *OrderingRemoveDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
OrderingRemoveBody ordering remove body
swagger:model OrderingRemoveBody
*/
type OrderingRemoveBody struct {

	// host
	Host string `json:"host,omitempty"`

	// joined orderer
	JoinedOrderer []*models.OrderingRequestJoined `json:"joinedOrderer"`

	// port
	Port int64 `json:"port,omitempty"`
}

// Validate validates this ordering remove body
func (o *OrderingRemoveBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateJoinedOrderer(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *OrderingRemoveBody) validateJoinedOrderer(formats strfmt.Registry) error {
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

// ContextValidate validate this ordering remove body based on the context it is used
func (o *OrderingRemoveBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateJoinedOrderer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *OrderingRemoveBody) contextValidateJoinedOrderer(ctx context.Context, formats strfmt.Registry) error {

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
func (o *OrderingRemoveBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *OrderingRemoveBody) UnmarshalBinary(b []byte) error {
	var res OrderingRemoveBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
