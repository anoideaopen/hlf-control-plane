// Code generated by go-swagger; DO NOT EDIT.

package chaincode

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new chaincode API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for chaincode API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ChaincodeInstall(params *ChaincodeInstallParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ChaincodeInstallOK, error)

	ChaincodeInstallExternal(params *ChaincodeInstallExternalParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ChaincodeInstallExternalOK, error)

	ChaincodeInstalledList(params *ChaincodeInstalledListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ChaincodeInstalledListOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ChaincodeInstall installations of chaincode package
*/
func (a *Client) ChaincodeInstall(params *ChaincodeInstallParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ChaincodeInstallOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChaincodeInstallParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "chaincodeInstall",
		Method:             "POST",
		PathPattern:        "/v1/chaincode/install",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ChaincodeInstallReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ChaincodeInstallOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ChaincodeInstallDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ChaincodeInstallExternal installations of external chaincode package
*/
func (a *Client) ChaincodeInstallExternal(params *ChaincodeInstallExternalParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ChaincodeInstallExternalOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChaincodeInstallExternalParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "chaincodeInstallExternal",
		Method:             "POST",
		PathPattern:        "/v1/chaincode/install-external",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ChaincodeInstallExternalReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ChaincodeInstallExternalOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ChaincodeInstallExternalDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ChaincodeInstalledList gets list of installed chaincodes
*/
func (a *Client) ChaincodeInstalledList(params *ChaincodeInstalledListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ChaincodeInstalledListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChaincodeInstalledListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "chaincodeInstalledList",
		Method:             "GET",
		PathPattern:        "/v1/chaincode/installed",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ChaincodeInstalledListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ChaincodeInstalledListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ChaincodeInstalledListDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
