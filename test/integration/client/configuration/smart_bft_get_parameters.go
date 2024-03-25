// Code generated by go-swagger; DO NOT EDIT.

package configuration

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

// NewSmartBftGetParams creates a new SmartBftGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSmartBftGetParams() *SmartBftGetParams {
	return &SmartBftGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSmartBftGetParamsWithTimeout creates a new SmartBftGetParams object
// with the ability to set a timeout on a request.
func NewSmartBftGetParamsWithTimeout(timeout time.Duration) *SmartBftGetParams {
	return &SmartBftGetParams{
		timeout: timeout,
	}
}

// NewSmartBftGetParamsWithContext creates a new SmartBftGetParams object
// with the ability to set a context for a request.
func NewSmartBftGetParamsWithContext(ctx context.Context) *SmartBftGetParams {
	return &SmartBftGetParams{
		Context: ctx,
	}
}

// NewSmartBftGetParamsWithHTTPClient creates a new SmartBftGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewSmartBftGetParamsWithHTTPClient(client *http.Client) *SmartBftGetParams {
	return &SmartBftGetParams{
		HTTPClient: client,
	}
}

/*
SmartBftGetParams contains all the parameters to send to the API endpoint

	for the smart bft get operation.

	Typically these are written to a http.Request.
*/
type SmartBftGetParams struct {

	// ChannelName.
	ChannelName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the smart bft get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SmartBftGetParams) WithDefaults() *SmartBftGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the smart bft get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SmartBftGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the smart bft get params
func (o *SmartBftGetParams) WithTimeout(timeout time.Duration) *SmartBftGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the smart bft get params
func (o *SmartBftGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the smart bft get params
func (o *SmartBftGetParams) WithContext(ctx context.Context) *SmartBftGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the smart bft get params
func (o *SmartBftGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the smart bft get params
func (o *SmartBftGetParams) WithHTTPClient(client *http.Client) *SmartBftGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the smart bft get params
func (o *SmartBftGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChannelName adds the channelName to the smart bft get params
func (o *SmartBftGetParams) WithChannelName(channelName string) *SmartBftGetParams {
	o.SetChannelName(channelName)
	return o
}

// SetChannelName adds the channelName to the smart bft get params
func (o *SmartBftGetParams) SetChannelName(channelName string) {
	o.ChannelName = channelName
}

// WriteToRequest writes these params to a swagger request
func (o *SmartBftGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param channelName
	if err := r.SetPathParam("channelName", o.ChannelName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}