// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/fortifi/go-api/client/account"
	"github.com/fortifi/go-api/client/accountant"
	"github.com/fortifi/go-api/client/appstore"
	"github.com/fortifi/go-api/client/attachment"
	"github.com/fortifi/go-api/client/authentication"
	"github.com/fortifi/go-api/client/brands"
	"github.com/fortifi/go-api/client/chat"
	"github.com/fortifi/go-api/client/comments"
	"github.com/fortifi/go-api/client/contacts"
	"github.com/fortifi/go-api/client/customers"
	"github.com/fortifi/go-api/client/deprecated"
	"github.com/fortifi/go-api/client/devices"
	"github.com/fortifi/go-api/client/entities"
	"github.com/fortifi/go-api/client/finance"
	"github.com/fortifi/go-api/client/interactions"
	"github.com/fortifi/go-api/client/licence"
	"github.com/fortifi/go-api/client/marketing"
	"github.com/fortifi/go-api/client/messenger"
	"github.com/fortifi/go-api/client/operations"
	"github.com/fortifi/go-api/client/orders"
	"github.com/fortifi/go-api/client/polymer"
	"github.com/fortifi/go-api/client/polymers"
	"github.com/fortifi/go-api/client/products"
	"github.com/fortifi/go-api/client/properties"
	"github.com/fortifi/go-api/client/reasons"
	"github.com/fortifi/go-api/client/reservations"
	"github.com/fortifi/go-api/client/reviews"
	"github.com/fortifi/go-api/client/service_status"
	"github.com/fortifi/go-api/client/support"
)

// Default fortifi API HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "api.fortifi.io"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/v1"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new fortifi API HTTP client.
func NewHTTPClient(formats strfmt.Registry) *FortifiAPI {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new fortifi API HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *FortifiAPI {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new fortifi API client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *FortifiAPI {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(FortifiAPI)
	cli.Transport = transport
	cli.Account = account.New(transport, formats)
	cli.Accountant = accountant.New(transport, formats)
	cli.Appstore = appstore.New(transport, formats)
	cli.Attachment = attachment.New(transport, formats)
	cli.Authentication = authentication.New(transport, formats)
	cli.Brands = brands.New(transport, formats)
	cli.Chat = chat.New(transport, formats)
	cli.Comments = comments.New(transport, formats)
	cli.Contacts = contacts.New(transport, formats)
	cli.Customers = customers.New(transport, formats)
	cli.Deprecated = deprecated.New(transport, formats)
	cli.Devices = devices.New(transport, formats)
	cli.Entities = entities.New(transport, formats)
	cli.Finance = finance.New(transport, formats)
	cli.Interactions = interactions.New(transport, formats)
	cli.Licence = licence.New(transport, formats)
	cli.Marketing = marketing.New(transport, formats)
	cli.Messenger = messenger.New(transport, formats)
	cli.Operations = operations.New(transport, formats)
	cli.Orders = orders.New(transport, formats)
	cli.Polymer = polymer.New(transport, formats)
	cli.Polymers = polymers.New(transport, formats)
	cli.Products = products.New(transport, formats)
	cli.Properties = properties.New(transport, formats)
	cli.Reasons = reasons.New(transport, formats)
	cli.Reservations = reservations.New(transport, formats)
	cli.Reviews = reviews.New(transport, formats)
	cli.ServiceStatus = service_status.New(transport, formats)
	cli.Support = support.New(transport, formats)
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

// FortifiAPI is a client for fortifi API
type FortifiAPI struct {
	Account account.ClientService

	Accountant accountant.ClientService

	Appstore appstore.ClientService

	Attachment attachment.ClientService

	Authentication authentication.ClientService

	Brands brands.ClientService

	Chat chat.ClientService

	Comments comments.ClientService

	Contacts contacts.ClientService

	Customers customers.ClientService

	Deprecated deprecated.ClientService

	Devices devices.ClientService

	Entities entities.ClientService

	Finance finance.ClientService

	Interactions interactions.ClientService

	Licence licence.ClientService

	Marketing marketing.ClientService

	Messenger messenger.ClientService

	Operations operations.ClientService

	Orders orders.ClientService

	Polymer polymer.ClientService

	Polymers polymers.ClientService

	Products products.ClientService

	Properties properties.ClientService

	Reasons reasons.ClientService

	Reservations reservations.ClientService

	Reviews reviews.ClientService

	ServiceStatus service_status.ClientService

	Support support.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *FortifiAPI) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Account.SetTransport(transport)
	c.Accountant.SetTransport(transport)
	c.Appstore.SetTransport(transport)
	c.Attachment.SetTransport(transport)
	c.Authentication.SetTransport(transport)
	c.Brands.SetTransport(transport)
	c.Chat.SetTransport(transport)
	c.Comments.SetTransport(transport)
	c.Contacts.SetTransport(transport)
	c.Customers.SetTransport(transport)
	c.Deprecated.SetTransport(transport)
	c.Devices.SetTransport(transport)
	c.Entities.SetTransport(transport)
	c.Finance.SetTransport(transport)
	c.Interactions.SetTransport(transport)
	c.Licence.SetTransport(transport)
	c.Marketing.SetTransport(transport)
	c.Messenger.SetTransport(transport)
	c.Operations.SetTransport(transport)
	c.Orders.SetTransport(transport)
	c.Polymer.SetTransport(transport)
	c.Polymers.SetTransport(transport)
	c.Products.SetTransport(transport)
	c.Properties.SetTransport(transport)
	c.Reasons.SetTransport(transport)
	c.Reservations.SetTransport(transport)
	c.Reviews.SetTransport(transport)
	c.ServiceStatus.SetTransport(transport)
	c.Support.SetTransport(transport)
}
