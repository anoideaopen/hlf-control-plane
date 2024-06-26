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

	"gitlab.n-t.io/core/library/hlf-tool/hlf-control-plane/test/integration/models"
)

// NewLifecycleApproveForMyOrgParams creates a new LifecycleApproveForMyOrgParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLifecycleApproveForMyOrgParams() *LifecycleApproveForMyOrgParams {
	return &LifecycleApproveForMyOrgParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLifecycleApproveForMyOrgParamsWithTimeout creates a new LifecycleApproveForMyOrgParams object
// with the ability to set a timeout on a request.
func NewLifecycleApproveForMyOrgParamsWithTimeout(timeout time.Duration) *LifecycleApproveForMyOrgParams {
	return &LifecycleApproveForMyOrgParams{
		timeout: timeout,
	}
}

// NewLifecycleApproveForMyOrgParamsWithContext creates a new LifecycleApproveForMyOrgParams object
// with the ability to set a context for a request.
func NewLifecycleApproveForMyOrgParamsWithContext(ctx context.Context) *LifecycleApproveForMyOrgParams {
	return &LifecycleApproveForMyOrgParams{
		Context: ctx,
	}
}

// NewLifecycleApproveForMyOrgParamsWithHTTPClient creates a new LifecycleApproveForMyOrgParams object
// with the ability to set a custom HTTPClient for a request.
func NewLifecycleApproveForMyOrgParamsWithHTTPClient(client *http.Client) *LifecycleApproveForMyOrgParams {
	return &LifecycleApproveForMyOrgParams{
		HTTPClient: client,
	}
}

/*
LifecycleApproveForMyOrgParams contains all the parameters to send to the API endpoint

	for the lifecycle approve for my org operation.

	Typically these are written to a http.Request.
*/
type LifecycleApproveForMyOrgParams struct {

	// Body.
	Body *models.ProtoLifecycleApproveForMyOrgRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the lifecycle approve for my org params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LifecycleApproveForMyOrgParams) WithDefaults() *LifecycleApproveForMyOrgParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the lifecycle approve for my org params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LifecycleApproveForMyOrgParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the lifecycle approve for my org params
func (o *LifecycleApproveForMyOrgParams) WithTimeout(timeout time.Duration) *LifecycleApproveForMyOrgParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the lifecycle approve for my org params
func (o *LifecycleApproveForMyOrgParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the lifecycle approve for my org params
func (o *LifecycleApproveForMyOrgParams) WithContext(ctx context.Context) *LifecycleApproveForMyOrgParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the lifecycle approve for my org params
func (o *LifecycleApproveForMyOrgParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the lifecycle approve for my org params
func (o *LifecycleApproveForMyOrgParams) WithHTTPClient(client *http.Client) *LifecycleApproveForMyOrgParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the lifecycle approve for my org params
func (o *LifecycleApproveForMyOrgParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the lifecycle approve for my org params
func (o *LifecycleApproveForMyOrgParams) WithBody(body *models.ProtoLifecycleApproveForMyOrgRequest) *LifecycleApproveForMyOrgParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the lifecycle approve for my org params
func (o *LifecycleApproveForMyOrgParams) SetBody(body *models.ProtoLifecycleApproveForMyOrgRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *LifecycleApproveForMyOrgParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
