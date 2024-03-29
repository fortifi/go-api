// Code generated by go-swagger; DO NOT EDIT.

package polymers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new polymers API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for polymers API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetPolymersParentFidPolymerFid(params *GetPolymersParentFidPolymerFidParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPolymersParentFidPolymerFidOK, error)

	PostPolymersParentFid(params *PostPolymersParentFidParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostPolymersParentFidOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetPolymersParentFidPolymerFid reads a polymer
*/
func (a *Client) GetPolymersParentFidPolymerFid(params *GetPolymersParentFidPolymerFidParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPolymersParentFidPolymerFidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPolymersParentFidPolymerFidParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetPolymersParentFidPolymerFid",
		Method:             "GET",
		PathPattern:        "/polymers/{parentFid}/{polymerFid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPolymersParentFidPolymerFidReader{formats: a.formats},
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
	success, ok := result.(*GetPolymersParentFidPolymerFidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetPolymersParentFidPolymerFidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostPolymersParentFid creates a new polymer
*/
func (a *Client) PostPolymersParentFid(params *PostPolymersParentFidParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostPolymersParentFidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostPolymersParentFidParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostPolymersParentFid",
		Method:             "POST",
		PathPattern:        "/polymers/{parentFid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostPolymersParentFidReader{formats: a.formats},
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
	success, ok := result.(*PostPolymersParentFidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostPolymersParentFidDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
