// Code generated by go-swagger; DO NOT EDIT.

package ordering

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewOrderingRemoveParams creates a new OrderingRemoveParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOrderingRemoveParams() *OrderingRemoveParams {
	return &OrderingRemoveParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOrderingRemoveParamsWithTimeout creates a new OrderingRemoveParams object
// with the ability to set a timeout on a request.
func NewOrderingRemoveParamsWithTimeout(timeout time.Duration) *OrderingRemoveParams {
	return &OrderingRemoveParams{
		timeout: timeout,
	}
}

// NewOrderingRemoveParamsWithContext creates a new OrderingRemoveParams object
// with the ability to set a context for a request.
func NewOrderingRemoveParamsWithContext(ctx context.Context) *OrderingRemoveParams {
	return &OrderingRemoveParams{
		Context: ctx,
	}
}

// NewOrderingRemoveParamsWithHTTPClient creates a new OrderingRemoveParams object
// with the ability to set a custom HTTPClient for a request.
func NewOrderingRemoveParamsWithHTTPClient(client *http.Client) *OrderingRemoveParams {
	return &OrderingRemoveParams{
		HTTPClient: client,
	}
}

/*
OrderingRemoveParams contains all the parameters to send to the API endpoint

	for the ordering remove operation.

	Typically these are written to a http.Request.
*/
type OrderingRemoveParams struct {

	// Body.
	Body OrderingRemoveBody

	// ChannelName.
	ChannelName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ordering remove params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OrderingRemoveParams) WithDefaults() *OrderingRemoveParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ordering remove params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OrderingRemoveParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ordering remove params
func (o *OrderingRemoveParams) WithTimeout(timeout time.Duration) *OrderingRemoveParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ordering remove params
func (o *OrderingRemoveParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ordering remove params
func (o *OrderingRemoveParams) WithContext(ctx context.Context) *OrderingRemoveParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ordering remove params
func (o *OrderingRemoveParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ordering remove params
func (o *OrderingRemoveParams) WithHTTPClient(client *http.Client) *OrderingRemoveParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ordering remove params
func (o *OrderingRemoveParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the ordering remove params
func (o *OrderingRemoveParams) WithBody(body OrderingRemoveBody) *OrderingRemoveParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the ordering remove params
func (o *OrderingRemoveParams) SetBody(body OrderingRemoveBody) {
	o.Body = body
}

// WithChannelName adds the channelName to the ordering remove params
func (o *OrderingRemoveParams) WithChannelName(channelName string) *OrderingRemoveParams {
	o.SetChannelName(channelName)
	return o
}

// SetChannelName adds the channelName to the ordering remove params
func (o *OrderingRemoveParams) SetChannelName(channelName string) {
	o.ChannelName = channelName
}

// WriteToRequest writes these params to a swagger request
func (o *OrderingRemoveParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param channelName
	if err := r.SetPathParam("channelName", o.ChannelName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
