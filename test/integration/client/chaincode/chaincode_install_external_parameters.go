// Code generated by go-swagger; DO NOT EDIT.

package chaincode

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

// NewChaincodeInstallExternalParams creates a new ChaincodeInstallExternalParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewChaincodeInstallExternalParams() *ChaincodeInstallExternalParams {
	return &ChaincodeInstallExternalParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewChaincodeInstallExternalParamsWithTimeout creates a new ChaincodeInstallExternalParams object
// with the ability to set a timeout on a request.
func NewChaincodeInstallExternalParamsWithTimeout(timeout time.Duration) *ChaincodeInstallExternalParams {
	return &ChaincodeInstallExternalParams{
		timeout: timeout,
	}
}

// NewChaincodeInstallExternalParamsWithContext creates a new ChaincodeInstallExternalParams object
// with the ability to set a context for a request.
func NewChaincodeInstallExternalParamsWithContext(ctx context.Context) *ChaincodeInstallExternalParams {
	return &ChaincodeInstallExternalParams{
		Context: ctx,
	}
}

// NewChaincodeInstallExternalParamsWithHTTPClient creates a new ChaincodeInstallExternalParams object
// with the ability to set a custom HTTPClient for a request.
func NewChaincodeInstallExternalParamsWithHTTPClient(client *http.Client) *ChaincodeInstallExternalParams {
	return &ChaincodeInstallExternalParams{
		HTTPClient: client,
	}
}

/*
ChaincodeInstallExternalParams contains all the parameters to send to the API endpoint

	for the chaincode install external operation.

	Typically these are written to a http.Request.
*/
type ChaincodeInstallExternalParams struct {

	// Body.
	Body *models.ProtoChaincodeInstallExternalRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the chaincode install external params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChaincodeInstallExternalParams) WithDefaults() *ChaincodeInstallExternalParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the chaincode install external params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChaincodeInstallExternalParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the chaincode install external params
func (o *ChaincodeInstallExternalParams) WithTimeout(timeout time.Duration) *ChaincodeInstallExternalParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the chaincode install external params
func (o *ChaincodeInstallExternalParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the chaincode install external params
func (o *ChaincodeInstallExternalParams) WithContext(ctx context.Context) *ChaincodeInstallExternalParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the chaincode install external params
func (o *ChaincodeInstallExternalParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the chaincode install external params
func (o *ChaincodeInstallExternalParams) WithHTTPClient(client *http.Client) *ChaincodeInstallExternalParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the chaincode install external params
func (o *ChaincodeInstallExternalParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the chaincode install external params
func (o *ChaincodeInstallExternalParams) WithBody(body *models.ProtoChaincodeInstallExternalRequest) *ChaincodeInstallExternalParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the chaincode install external params
func (o *ChaincodeInstallExternalParams) SetBody(body *models.ProtoChaincodeInstallExternalRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ChaincodeInstallExternalParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
