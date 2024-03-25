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

// NewLifecycleCheckCommitReadinessParams creates a new LifecycleCheckCommitReadinessParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLifecycleCheckCommitReadinessParams() *LifecycleCheckCommitReadinessParams {
	return &LifecycleCheckCommitReadinessParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLifecycleCheckCommitReadinessParamsWithTimeout creates a new LifecycleCheckCommitReadinessParams object
// with the ability to set a timeout on a request.
func NewLifecycleCheckCommitReadinessParamsWithTimeout(timeout time.Duration) *LifecycleCheckCommitReadinessParams {
	return &LifecycleCheckCommitReadinessParams{
		timeout: timeout,
	}
}

// NewLifecycleCheckCommitReadinessParamsWithContext creates a new LifecycleCheckCommitReadinessParams object
// with the ability to set a context for a request.
func NewLifecycleCheckCommitReadinessParamsWithContext(ctx context.Context) *LifecycleCheckCommitReadinessParams {
	return &LifecycleCheckCommitReadinessParams{
		Context: ctx,
	}
}

// NewLifecycleCheckCommitReadinessParamsWithHTTPClient creates a new LifecycleCheckCommitReadinessParams object
// with the ability to set a custom HTTPClient for a request.
func NewLifecycleCheckCommitReadinessParamsWithHTTPClient(client *http.Client) *LifecycleCheckCommitReadinessParams {
	return &LifecycleCheckCommitReadinessParams{
		HTTPClient: client,
	}
}

/*
LifecycleCheckCommitReadinessParams contains all the parameters to send to the API endpoint

	for the lifecycle check commit readiness operation.

	Typically these are written to a http.Request.
*/
type LifecycleCheckCommitReadinessParams struct {

	// Body.
	Body *models.ProtoLifecycleCheckCommitReadinessRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the lifecycle check commit readiness params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LifecycleCheckCommitReadinessParams) WithDefaults() *LifecycleCheckCommitReadinessParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the lifecycle check commit readiness params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LifecycleCheckCommitReadinessParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the lifecycle check commit readiness params
func (o *LifecycleCheckCommitReadinessParams) WithTimeout(timeout time.Duration) *LifecycleCheckCommitReadinessParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the lifecycle check commit readiness params
func (o *LifecycleCheckCommitReadinessParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the lifecycle check commit readiness params
func (o *LifecycleCheckCommitReadinessParams) WithContext(ctx context.Context) *LifecycleCheckCommitReadinessParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the lifecycle check commit readiness params
func (o *LifecycleCheckCommitReadinessParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the lifecycle check commit readiness params
func (o *LifecycleCheckCommitReadinessParams) WithHTTPClient(client *http.Client) *LifecycleCheckCommitReadinessParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the lifecycle check commit readiness params
func (o *LifecycleCheckCommitReadinessParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the lifecycle check commit readiness params
func (o *LifecycleCheckCommitReadinessParams) WithBody(body *models.ProtoLifecycleCheckCommitReadinessRequest) *LifecycleCheckCommitReadinessParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the lifecycle check commit readiness params
func (o *LifecycleCheckCommitReadinessParams) SetBody(body *models.ProtoLifecycleCheckCommitReadinessRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *LifecycleCheckCommitReadinessParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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