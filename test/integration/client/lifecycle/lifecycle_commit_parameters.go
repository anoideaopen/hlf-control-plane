// Code generated by go-swagger; DO NOT EDIT.

package lifecycle

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

	"github.com/atomyze-foundation/hlf-control-plane/test/integration/models"
)

// NewLifecycleCommitParams creates a new LifecycleCommitParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLifecycleCommitParams() *LifecycleCommitParams {
	return &LifecycleCommitParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLifecycleCommitParamsWithTimeout creates a new LifecycleCommitParams object
// with the ability to set a timeout on a request.
func NewLifecycleCommitParamsWithTimeout(timeout time.Duration) *LifecycleCommitParams {
	return &LifecycleCommitParams{
		timeout: timeout,
	}
}

// NewLifecycleCommitParamsWithContext creates a new LifecycleCommitParams object
// with the ability to set a context for a request.
func NewLifecycleCommitParamsWithContext(ctx context.Context) *LifecycleCommitParams {
	return &LifecycleCommitParams{
		Context: ctx,
	}
}

// NewLifecycleCommitParamsWithHTTPClient creates a new LifecycleCommitParams object
// with the ability to set a custom HTTPClient for a request.
func NewLifecycleCommitParamsWithHTTPClient(client *http.Client) *LifecycleCommitParams {
	return &LifecycleCommitParams{
		HTTPClient: client,
	}
}

/*
LifecycleCommitParams contains all the parameters to send to the API endpoint

	for the lifecycle commit operation.

	Typically these are written to a http.Request.
*/
type LifecycleCommitParams struct {

	// Body.
	Body *models.ProtoLifecycleCommitRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the lifecycle commit params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LifecycleCommitParams) WithDefaults() *LifecycleCommitParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the lifecycle commit params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LifecycleCommitParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the lifecycle commit params
func (o *LifecycleCommitParams) WithTimeout(timeout time.Duration) *LifecycleCommitParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the lifecycle commit params
func (o *LifecycleCommitParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the lifecycle commit params
func (o *LifecycleCommitParams) WithContext(ctx context.Context) *LifecycleCommitParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the lifecycle commit params
func (o *LifecycleCommitParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the lifecycle commit params
func (o *LifecycleCommitParams) WithHTTPClient(client *http.Client) *LifecycleCommitParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the lifecycle commit params
func (o *LifecycleCommitParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the lifecycle commit params
func (o *LifecycleCommitParams) WithBody(body *models.ProtoLifecycleCommitRequest) *LifecycleCommitParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the lifecycle commit params
func (o *LifecycleCommitParams) SetBody(body *models.ProtoLifecycleCommitRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *LifecycleCommitParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}