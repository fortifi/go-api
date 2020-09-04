// Code generated by go-swagger; DO NOT EDIT.

package entities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new entities API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for entities API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteEntitiesEntityFidPropertiesValuesPropertyName(params *DeleteEntitiesEntityFidPropertiesValuesPropertyNameParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteEntitiesEntityFidPropertiesValuesPropertyNameOK, error)

	GetEntitiesEntityFidLabels(params *GetEntitiesEntityFidLabelsParams, authInfo runtime.ClientAuthInfoWriter) (*GetEntitiesEntityFidLabelsOK, error)

	GetEntitiesEntityFidProperties(params *GetEntitiesEntityFidPropertiesParams, authInfo runtime.ClientAuthInfoWriter) (*GetEntitiesEntityFidPropertiesOK, error)

	GetEntitiesEntityFidPropertiesValuesPropertyName(params *GetEntitiesEntityFidPropertiesValuesPropertyNameParams, authInfo runtime.ClientAuthInfoWriter) (*GetEntitiesEntityFidPropertiesValuesPropertyNameOK, error)

	PostEntitiesEntityFidAttachmentsUploadURL(params *PostEntitiesEntityFidAttachmentsUploadURLParams, authInfo runtime.ClientAuthInfoWriter) (*PostEntitiesEntityFidAttachmentsUploadURLOK, error)

	PostEntitiesEntityFidEvents(params *PostEntitiesEntityFidEventsParams, authInfo runtime.ClientAuthInfoWriter) (*PostEntitiesEntityFidEventsOK, error)

	PostUploadUploadURL(params *PostUploadUploadURLParams, authInfo runtime.ClientAuthInfoWriter) (*PostUploadUploadURLOK, error)

	PutEntitiesEntityFidLabelsAssign(params *PutEntitiesEntityFidLabelsAssignParams, authInfo runtime.ClientAuthInfoWriter) (*PutEntitiesEntityFidLabelsAssignOK, error)

	PutEntitiesEntityFidProperties(params *PutEntitiesEntityFidPropertiesParams, authInfo runtime.ClientAuthInfoWriter) (*PutEntitiesEntityFidPropertiesOK, error)

	PutEntitiesEntityFidPropertiesValuesPropertyName(params *PutEntitiesEntityFidPropertiesValuesPropertyNameParams, authInfo runtime.ClientAuthInfoWriter) (*PutEntitiesEntityFidPropertiesValuesPropertyNameOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteEntitiesEntityFidPropertiesValuesPropertyName removes a value property from an entity
*/
func (a *Client) DeleteEntitiesEntityFidPropertiesValuesPropertyName(params *DeleteEntitiesEntityFidPropertiesValuesPropertyNameParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteEntitiesEntityFidPropertiesValuesPropertyNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteEntitiesEntityFidPropertiesValuesPropertyNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteEntitiesEntityFidPropertiesValuesPropertyName",
		Method:             "DELETE",
		PathPattern:        "/entities/{entityFid}/properties/values/{propertyName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteEntitiesEntityFidPropertiesValuesPropertyNameReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteEntitiesEntityFidPropertiesValuesPropertyNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteEntitiesEntityFidPropertiesValuesPropertyNameDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetEntitiesEntityFidLabels gets all labels for an entity
*/
func (a *Client) GetEntitiesEntityFidLabels(params *GetEntitiesEntityFidLabelsParams, authInfo runtime.ClientAuthInfoWriter) (*GetEntitiesEntityFidLabelsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEntitiesEntityFidLabelsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetEntitiesEntityFidLabels",
		Method:             "GET",
		PathPattern:        "/entities/{entityFid}/labels",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetEntitiesEntityFidLabelsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetEntitiesEntityFidLabelsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetEntitiesEntityFidLabelsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetEntitiesEntityFidProperties gets all properties for an entity
*/
func (a *Client) GetEntitiesEntityFidProperties(params *GetEntitiesEntityFidPropertiesParams, authInfo runtime.ClientAuthInfoWriter) (*GetEntitiesEntityFidPropertiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEntitiesEntityFidPropertiesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetEntitiesEntityFidProperties",
		Method:             "GET",
		PathPattern:        "/entities/{entityFid}/properties",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetEntitiesEntityFidPropertiesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetEntitiesEntityFidPropertiesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetEntitiesEntityFidPropertiesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetEntitiesEntityFidPropertiesValuesPropertyName gets a property value from an entity
*/
func (a *Client) GetEntitiesEntityFidPropertiesValuesPropertyName(params *GetEntitiesEntityFidPropertiesValuesPropertyNameParams, authInfo runtime.ClientAuthInfoWriter) (*GetEntitiesEntityFidPropertiesValuesPropertyNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEntitiesEntityFidPropertiesValuesPropertyNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetEntitiesEntityFidPropertiesValuesPropertyName",
		Method:             "GET",
		PathPattern:        "/entities/{entityFid}/properties/values/{propertyName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetEntitiesEntityFidPropertiesValuesPropertyNameReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetEntitiesEntityFidPropertiesValuesPropertyNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetEntitiesEntityFidPropertiesValuesPropertyNameDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PostEntitiesEntityFidAttachmentsUploadURL creates an upload url
*/
func (a *Client) PostEntitiesEntityFidAttachmentsUploadURL(params *PostEntitiesEntityFidAttachmentsUploadURLParams, authInfo runtime.ClientAuthInfoWriter) (*PostEntitiesEntityFidAttachmentsUploadURLOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostEntitiesEntityFidAttachmentsUploadURLParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostEntitiesEntityFidAttachmentsUploadURL",
		Method:             "POST",
		PathPattern:        "/entities/{entityFid}/attachments/uploadUrl",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostEntitiesEntityFidAttachmentsUploadURLReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostEntitiesEntityFidAttachmentsUploadURLOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostEntitiesEntityFidAttachmentsUploadURLDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PostEntitiesEntityFidEvents triggers a new event
*/
func (a *Client) PostEntitiesEntityFidEvents(params *PostEntitiesEntityFidEventsParams, authInfo runtime.ClientAuthInfoWriter) (*PostEntitiesEntityFidEventsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostEntitiesEntityFidEventsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostEntitiesEntityFidEvents",
		Method:             "POST",
		PathPattern:        "/entities/{entityFid}/events",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostEntitiesEntityFidEventsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostEntitiesEntityFidEventsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostEntitiesEntityFidEventsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PostUploadUploadURL creates an upload attachment url

  This call will give you a URL to PUT files to and a unique filename. You can upload files to the URL like this: ```curl -X PUT -d @filename -H 'content-type: text/plain' 'url'```

*/
func (a *Client) PostUploadUploadURL(params *PostUploadUploadURLParams, authInfo runtime.ClientAuthInfoWriter) (*PostUploadUploadURLOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostUploadUploadURLParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostUploadUploadURL",
		Method:             "POST",
		PathPattern:        "/upload/uploadUrl",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostUploadUploadURLReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostUploadUploadURLOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostUploadUploadURLDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PutEntitiesEntityFidLabelsAssign assigns label to entity
*/
func (a *Client) PutEntitiesEntityFidLabelsAssign(params *PutEntitiesEntityFidLabelsAssignParams, authInfo runtime.ClientAuthInfoWriter) (*PutEntitiesEntityFidLabelsAssignOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutEntitiesEntityFidLabelsAssignParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutEntitiesEntityFidLabelsAssign",
		Method:             "PUT",
		PathPattern:        "/entities/{entityFid}/labels/assign",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutEntitiesEntityFidLabelsAssignReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutEntitiesEntityFidLabelsAssignOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PutEntitiesEntityFidLabelsAssignDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PutEntitiesEntityFidProperties writes multiple entity properties
*/
func (a *Client) PutEntitiesEntityFidProperties(params *PutEntitiesEntityFidPropertiesParams, authInfo runtime.ClientAuthInfoWriter) (*PutEntitiesEntityFidPropertiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutEntitiesEntityFidPropertiesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutEntitiesEntityFidProperties",
		Method:             "PUT",
		PathPattern:        "/entities/{entityFid}/properties",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutEntitiesEntityFidPropertiesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutEntitiesEntityFidPropertiesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PutEntitiesEntityFidPropertiesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PutEntitiesEntityFidPropertiesValuesPropertyName writes an entity value property
*/
func (a *Client) PutEntitiesEntityFidPropertiesValuesPropertyName(params *PutEntitiesEntityFidPropertiesValuesPropertyNameParams, authInfo runtime.ClientAuthInfoWriter) (*PutEntitiesEntityFidPropertiesValuesPropertyNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutEntitiesEntityFidPropertiesValuesPropertyNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutEntitiesEntityFidPropertiesValuesPropertyName",
		Method:             "PUT",
		PathPattern:        "/entities/{entityFid}/properties/values/{propertyName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutEntitiesEntityFidPropertiesValuesPropertyNameReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutEntitiesEntityFidPropertiesValuesPropertyNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PutEntitiesEntityFidPropertiesValuesPropertyNameDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
