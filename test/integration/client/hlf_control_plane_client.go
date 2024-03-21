// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"gitlab.n-t.io/core/library/hlf-tool/hlf-control-plane/test/integration/client/chaincode"
	"gitlab.n-t.io/core/library/hlf-tool/hlf-control-plane/test/integration/client/channels"
	"gitlab.n-t.io/core/library/hlf-tool/hlf-control-plane/test/integration/client/configuration"
	"gitlab.n-t.io/core/library/hlf-tool/hlf-control-plane/test/integration/client/discovery"
	"gitlab.n-t.io/core/library/hlf-tool/hlf-control-plane/test/integration/client/lifecycle"
	"gitlab.n-t.io/core/library/hlf-tool/hlf-control-plane/test/integration/client/ordering"
)

// Default hlf control plane HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http"}

// NewHTTPClient creates a new hlf control plane HTTP client.
func NewHTTPClient(formats strfmt.Registry) *HlfControlPlane {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new hlf control plane HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *HlfControlPlane {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new hlf control plane client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *HlfControlPlane {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(HlfControlPlane)
	cli.Transport = transport
	cli.Chaincode = chaincode.New(transport, formats)
	cli.Channels = channels.New(transport, formats)
	cli.Configuration = configuration.New(transport, formats)
	cli.Discovery = discovery.New(transport, formats)
	cli.Lifecycle = lifecycle.New(transport, formats)
	cli.Ordering = ordering.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// HlfControlPlane is a client for hlf control plane
type HlfControlPlane struct {
	Chaincode chaincode.ClientService

	Channels channels.ClientService

	Configuration configuration.ClientService

	Discovery discovery.ClientService

	Lifecycle lifecycle.ClientService

	Ordering ordering.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *HlfControlPlane) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Chaincode.SetTransport(transport)
	c.Channels.SetTransport(transport)
	c.Configuration.SetTransport(transport)
	c.Discovery.SetTransport(transport)
	c.Lifecycle.SetTransport(transport)
	c.Ordering.SetTransport(transport)
}
