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

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteEntitiesEntityFidConfigSectionName(params *DeleteEntitiesEntityFidConfigSectionNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteEntitiesEntityFidConfigSectionNameOK, error)

	DeleteEntitiesEntityFidPropertiesCountersPropertyName(params *DeleteEntitiesEntityFidPropertiesCountersPropertyNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteEntitiesEntityFidPropertiesCountersPropertyNameOK, error)

	DeleteEntitiesEntityFidPropertiesFlagsPropertyName(params *DeleteEntitiesEntityFidPropertiesFlagsPropertyNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteEntitiesEntityFidPropertiesFlagsPropertyNameOK, error)

	DeleteEntitiesEntityFidPropertiesValuesPropertyName(params *DeleteEntitiesEntityFidPropertiesValuesPropertyNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteEntitiesEntityFidPropertiesValuesPropertyNameOK, error)

	GetEntitiesEntityFidConfigSectionName(params *GetEntitiesEntityFidConfigSectionNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEntitiesEntityFidConfigSectionNameOK, error)

	GetEntitiesEntityFidConfigSectionNameItemsItemName(params *GetEntitiesEntityFidConfigSectionNameItemsItemNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEntitiesEntityFidConfigSectionNameItemsItemNameOK, error)

	GetEntitiesEntityFidLabels(params *GetEntitiesEntityFidLabelsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEntitiesEntityFidLabelsOK, error)

	GetEntitiesEntityFidProperties(params *GetEntitiesEntityFidPropertiesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEntitiesEntityFidPropertiesOK, error)

	GetEntitiesEntityFidPropertiesFlagsPropertyName(params *GetEntitiesEntityFidPropertiesFlagsPropertyNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEntitiesEntityFidPropertiesFlagsPropertyNameOK, error)

	GetEntitiesEntityFidPropertiesValuesPropertyName(params *GetEntitiesEntityFidPropertiesValuesPropertyNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEntitiesEntityFidPropertiesValuesPropertyNameOK, error)

	PostEntitiesEntityFidAttachmentsUploadURL(params *PostEntitiesEntityFidAttachmentsUploadURLParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostEntitiesEntityFidAttachmentsUploadURLOK, error)

	PostEntitiesEntityFidConfigSectionName(params *PostEntitiesEntityFidConfigSectionNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostEntitiesEntityFidConfigSectionNameOK, error)

	PostEntitiesEntityFidEvents(params *PostEntitiesEntityFidEventsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostEntitiesEntityFidEventsOK, error)

	PostUploadUploadURL(params *PostUploadUploadURLParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostUploadUploadURLOK, error)

	PutEntitiesEntityFidLabelsAssign(params *PutEntitiesEntityFidLabelsAssignParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutEntitiesEntityFidLabelsAssignOK, error)

	PutEntitiesEntityFidProperties(params *PutEntitiesEntityFidPropertiesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutEntitiesEntityFidPropertiesOK, error)

	PutEntitiesEntityFidPropertiesCountersPropertyNameDecrement(params *PutEntitiesEntityFidPropertiesCountersPropertyNameDecrementParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutEntitiesEntityFidPropertiesCountersPropertyNameDecrementOK, error)

	PutEntitiesEntityFidPropertiesCountersPropertyNameIncrement(params *PutEntitiesEntityFidPropertiesCountersPropertyNameIncrementParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutEntitiesEntityFidPropertiesCountersPropertyNameIncrementOK, error)

	PutEntitiesEntityFidPropertiesFlagsPropertyName(params *PutEntitiesEntityFidPropertiesFlagsPropertyNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutEntitiesEntityFidPropertiesFlagsPropertyNameOK, error)

	PutEntitiesEntityFidPropertiesValuesPropertyName(params *PutEntitiesEntityFidPropertiesValuesPropertyNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutEntitiesEntityFidPropertiesValuesPropertyNameOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DeleteEntitiesEntityFidConfigSectionName removes a config section or property from an entity
*/
func (a *Client) DeleteEntitiesEntityFidConfigSectionName(params *DeleteEntitiesEntityFidConfigSectionNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteEntitiesEntityFidConfigSectionNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteEntitiesEntityFidConfigSectionNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteEntitiesEntityFidConfigSectionName",
		Method:             "DELETE",
		PathPattern:        "/entities/{entityFid}/config/{sectionName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteEntitiesEntityFidConfigSectionNameReader{formats: a.formats},
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
	success, ok := result.(*DeleteEntitiesEntityFidConfigSectionNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteEntitiesEntityFidConfigSectionNameDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteEntitiesEntityFidPropertiesCountersPropertyName removes a counter from an entity
*/
func (a *Client) DeleteEntitiesEntityFidPropertiesCountersPropertyName(params *DeleteEntitiesEntityFidPropertiesCountersPropertyNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteEntitiesEntityFidPropertiesCountersPropertyNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteEntitiesEntityFidPropertiesCountersPropertyNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteEntitiesEntityFidPropertiesCountersPropertyName",
		Method:             "DELETE",
		PathPattern:        "/entities/{entityFid}/properties/counters/{propertyName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteEntitiesEntityFidPropertiesCountersPropertyNameReader{formats: a.formats},
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
	success, ok := result.(*DeleteEntitiesEntityFidPropertiesCountersPropertyNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteEntitiesEntityFidPropertiesCountersPropertyNameDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteEntitiesEntityFidPropertiesFlagsPropertyName removes a flag from an entity
*/
func (a *Client) DeleteEntitiesEntityFidPropertiesFlagsPropertyName(params *DeleteEntitiesEntityFidPropertiesFlagsPropertyNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteEntitiesEntityFidPropertiesFlagsPropertyNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteEntitiesEntityFidPropertiesFlagsPropertyNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteEntitiesEntityFidPropertiesFlagsPropertyName",
		Method:             "DELETE",
		PathPattern:        "/entities/{entityFid}/properties/flags/{propertyName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteEntitiesEntityFidPropertiesFlagsPropertyNameReader{formats: a.formats},
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
	success, ok := result.(*DeleteEntitiesEntityFidPropertiesFlagsPropertyNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteEntitiesEntityFidPropertiesFlagsPropertyNameDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteEntitiesEntityFidPropertiesValuesPropertyName removes a value property from an entity
*/
func (a *Client) DeleteEntitiesEntityFidPropertiesValuesPropertyName(params *DeleteEntitiesEntityFidPropertiesValuesPropertyNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteEntitiesEntityFidPropertiesValuesPropertyNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteEntitiesEntityFidPropertiesValuesPropertyNameParams()
	}
	op := &runtime.ClientOperation{
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
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
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
GetEntitiesEntityFidConfigSectionName retrieves a config section
*/
func (a *Client) GetEntitiesEntityFidConfigSectionName(params *GetEntitiesEntityFidConfigSectionNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEntitiesEntityFidConfigSectionNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEntitiesEntityFidConfigSectionNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetEntitiesEntityFidConfigSectionName",
		Method:             "GET",
		PathPattern:        "/entities/{entityFid}/config/{sectionName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetEntitiesEntityFidConfigSectionNameReader{formats: a.formats},
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
	success, ok := result.(*GetEntitiesEntityFidConfigSectionNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetEntitiesEntityFidConfigSectionNameDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetEntitiesEntityFidConfigSectionNameItemsItemName retrieves a config item
*/
func (a *Client) GetEntitiesEntityFidConfigSectionNameItemsItemName(params *GetEntitiesEntityFidConfigSectionNameItemsItemNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEntitiesEntityFidConfigSectionNameItemsItemNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEntitiesEntityFidConfigSectionNameItemsItemNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetEntitiesEntityFidConfigSectionNameItemsItemName",
		Method:             "GET",
		PathPattern:        "/entities/{entityFid}/config/{sectionName}/items/{itemName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetEntitiesEntityFidConfigSectionNameItemsItemNameReader{formats: a.formats},
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
	success, ok := result.(*GetEntitiesEntityFidConfigSectionNameItemsItemNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetEntitiesEntityFidConfigSectionNameItemsItemNameDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetEntitiesEntityFidLabels gets all labels for an entity
*/
func (a *Client) GetEntitiesEntityFidLabels(params *GetEntitiesEntityFidLabelsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEntitiesEntityFidLabelsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEntitiesEntityFidLabelsParams()
	}
	op := &runtime.ClientOperation{
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
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
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
func (a *Client) GetEntitiesEntityFidProperties(params *GetEntitiesEntityFidPropertiesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEntitiesEntityFidPropertiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEntitiesEntityFidPropertiesParams()
	}
	op := &runtime.ClientOperation{
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
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
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
GetEntitiesEntityFidPropertiesFlagsPropertyName gets a flag property from an entity
*/
func (a *Client) GetEntitiesEntityFidPropertiesFlagsPropertyName(params *GetEntitiesEntityFidPropertiesFlagsPropertyNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEntitiesEntityFidPropertiesFlagsPropertyNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEntitiesEntityFidPropertiesFlagsPropertyNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetEntitiesEntityFidPropertiesFlagsPropertyName",
		Method:             "GET",
		PathPattern:        "/entities/{entityFid}/properties/flags/{propertyName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetEntitiesEntityFidPropertiesFlagsPropertyNameReader{formats: a.formats},
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
	success, ok := result.(*GetEntitiesEntityFidPropertiesFlagsPropertyNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetEntitiesEntityFidPropertiesFlagsPropertyNameDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetEntitiesEntityFidPropertiesValuesPropertyName gets a property value from an entity
*/
func (a *Client) GetEntitiesEntityFidPropertiesValuesPropertyName(params *GetEntitiesEntityFidPropertiesValuesPropertyNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEntitiesEntityFidPropertiesValuesPropertyNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEntitiesEntityFidPropertiesValuesPropertyNameParams()
	}
	op := &runtime.ClientOperation{
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
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
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
func (a *Client) PostEntitiesEntityFidAttachmentsUploadURL(params *PostEntitiesEntityFidAttachmentsUploadURLParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostEntitiesEntityFidAttachmentsUploadURLOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostEntitiesEntityFidAttachmentsUploadURLParams()
	}
	op := &runtime.ClientOperation{
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
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
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
PostEntitiesEntityFidConfigSectionName writes a config item
*/
func (a *Client) PostEntitiesEntityFidConfigSectionName(params *PostEntitiesEntityFidConfigSectionNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostEntitiesEntityFidConfigSectionNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostEntitiesEntityFidConfigSectionNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostEntitiesEntityFidConfigSectionName",
		Method:             "POST",
		PathPattern:        "/entities/{entityFid}/config/{sectionName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostEntitiesEntityFidConfigSectionNameReader{formats: a.formats},
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
	success, ok := result.(*PostEntitiesEntityFidConfigSectionNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostEntitiesEntityFidConfigSectionNameDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostEntitiesEntityFidEvents triggers a new event
*/
func (a *Client) PostEntitiesEntityFidEvents(params *PostEntitiesEntityFidEventsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostEntitiesEntityFidEventsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostEntitiesEntityFidEventsParams()
	}
	op := &runtime.ClientOperation{
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
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
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
func (a *Client) PostUploadUploadURL(params *PostUploadUploadURLParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostUploadUploadURLOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostUploadUploadURLParams()
	}
	op := &runtime.ClientOperation{
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
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
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
func (a *Client) PutEntitiesEntityFidLabelsAssign(params *PutEntitiesEntityFidLabelsAssignParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutEntitiesEntityFidLabelsAssignOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutEntitiesEntityFidLabelsAssignParams()
	}
	op := &runtime.ClientOperation{
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
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
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
func (a *Client) PutEntitiesEntityFidProperties(params *PutEntitiesEntityFidPropertiesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutEntitiesEntityFidPropertiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutEntitiesEntityFidPropertiesParams()
	}
	op := &runtime.ClientOperation{
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
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
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
PutEntitiesEntityFidPropertiesCountersPropertyNameDecrement decrements an entity counter
*/
func (a *Client) PutEntitiesEntityFidPropertiesCountersPropertyNameDecrement(params *PutEntitiesEntityFidPropertiesCountersPropertyNameDecrementParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutEntitiesEntityFidPropertiesCountersPropertyNameDecrementOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutEntitiesEntityFidPropertiesCountersPropertyNameDecrementParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PutEntitiesEntityFidPropertiesCountersPropertyNameDecrement",
		Method:             "PUT",
		PathPattern:        "/entities/{entityFid}/properties/counters/{propertyName}/decrement",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutEntitiesEntityFidPropertiesCountersPropertyNameDecrementReader{formats: a.formats},
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
	success, ok := result.(*PutEntitiesEntityFidPropertiesCountersPropertyNameDecrementOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PutEntitiesEntityFidPropertiesCountersPropertyNameDecrementDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PutEntitiesEntityFidPropertiesCountersPropertyNameIncrement increments an entity counter
*/
func (a *Client) PutEntitiesEntityFidPropertiesCountersPropertyNameIncrement(params *PutEntitiesEntityFidPropertiesCountersPropertyNameIncrementParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutEntitiesEntityFidPropertiesCountersPropertyNameIncrementOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutEntitiesEntityFidPropertiesCountersPropertyNameIncrementParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PutEntitiesEntityFidPropertiesCountersPropertyNameIncrement",
		Method:             "PUT",
		PathPattern:        "/entities/{entityFid}/properties/counters/{propertyName}/increment",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutEntitiesEntityFidPropertiesCountersPropertyNameIncrementReader{formats: a.formats},
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
	success, ok := result.(*PutEntitiesEntityFidPropertiesCountersPropertyNameIncrementOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PutEntitiesEntityFidPropertiesCountersPropertyNameIncrementDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PutEntitiesEntityFidPropertiesFlagsPropertyName writes an entity flag
*/
func (a *Client) PutEntitiesEntityFidPropertiesFlagsPropertyName(params *PutEntitiesEntityFidPropertiesFlagsPropertyNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutEntitiesEntityFidPropertiesFlagsPropertyNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutEntitiesEntityFidPropertiesFlagsPropertyNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PutEntitiesEntityFidPropertiesFlagsPropertyName",
		Method:             "PUT",
		PathPattern:        "/entities/{entityFid}/properties/flags/{propertyName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutEntitiesEntityFidPropertiesFlagsPropertyNameReader{formats: a.formats},
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
	success, ok := result.(*PutEntitiesEntityFidPropertiesFlagsPropertyNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PutEntitiesEntityFidPropertiesFlagsPropertyNameDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PutEntitiesEntityFidPropertiesValuesPropertyName writes an entity value property
*/
func (a *Client) PutEntitiesEntityFidPropertiesValuesPropertyName(params *PutEntitiesEntityFidPropertiesValuesPropertyNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutEntitiesEntityFidPropertiesValuesPropertyNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutEntitiesEntityFidPropertiesValuesPropertyNameParams()
	}
	op := &runtime.ClientOperation{
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
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
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
