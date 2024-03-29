// Code generated by go-swagger; DO NOT EDIT.

package entities

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

	"github.com/fortifi/go-api/models"
)

// NewPostEntitiesEntityFidEventsParams creates a new PostEntitiesEntityFidEventsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostEntitiesEntityFidEventsParams() *PostEntitiesEntityFidEventsParams {
	return &PostEntitiesEntityFidEventsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostEntitiesEntityFidEventsParamsWithTimeout creates a new PostEntitiesEntityFidEventsParams object
// with the ability to set a timeout on a request.
func NewPostEntitiesEntityFidEventsParamsWithTimeout(timeout time.Duration) *PostEntitiesEntityFidEventsParams {
	return &PostEntitiesEntityFidEventsParams{
		timeout: timeout,
	}
}

// NewPostEntitiesEntityFidEventsParamsWithContext creates a new PostEntitiesEntityFidEventsParams object
// with the ability to set a context for a request.
func NewPostEntitiesEntityFidEventsParamsWithContext(ctx context.Context) *PostEntitiesEntityFidEventsParams {
	return &PostEntitiesEntityFidEventsParams{
		Context: ctx,
	}
}

// NewPostEntitiesEntityFidEventsParamsWithHTTPClient creates a new PostEntitiesEntityFidEventsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostEntitiesEntityFidEventsParamsWithHTTPClient(client *http.Client) *PostEntitiesEntityFidEventsParams {
	return &PostEntitiesEntityFidEventsParams{
		HTTPClient: client,
	}
}

/*
PostEntitiesEntityFidEventsParams contains all the parameters to send to the API endpoint

	for the post entities entity fid events operation.

	Typically these are written to a http.Request.
*/
type PostEntitiesEntityFidEventsParams struct {

	/* EntityFid.

	   Entity FID to use
	*/
	EntityFid string

	// Payload.
	Payload *models.TriggerActionPayload

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post entities entity fid events params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostEntitiesEntityFidEventsParams) WithDefaults() *PostEntitiesEntityFidEventsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post entities entity fid events params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostEntitiesEntityFidEventsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post entities entity fid events params
func (o *PostEntitiesEntityFidEventsParams) WithTimeout(timeout time.Duration) *PostEntitiesEntityFidEventsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post entities entity fid events params
func (o *PostEntitiesEntityFidEventsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post entities entity fid events params
func (o *PostEntitiesEntityFidEventsParams) WithContext(ctx context.Context) *PostEntitiesEntityFidEventsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post entities entity fid events params
func (o *PostEntitiesEntityFidEventsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post entities entity fid events params
func (o *PostEntitiesEntityFidEventsParams) WithHTTPClient(client *http.Client) *PostEntitiesEntityFidEventsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post entities entity fid events params
func (o *PostEntitiesEntityFidEventsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEntityFid adds the entityFid to the post entities entity fid events params
func (o *PostEntitiesEntityFidEventsParams) WithEntityFid(entityFid string) *PostEntitiesEntityFidEventsParams {
	o.SetEntityFid(entityFid)
	return o
}

// SetEntityFid adds the entityFid to the post entities entity fid events params
func (o *PostEntitiesEntityFidEventsParams) SetEntityFid(entityFid string) {
	o.EntityFid = entityFid
}

// WithPayload adds the payload to the post entities entity fid events params
func (o *PostEntitiesEntityFidEventsParams) WithPayload(payload *models.TriggerActionPayload) *PostEntitiesEntityFidEventsParams {
	o.SetPayload(payload)
	return o
}

// SetPayload adds the payload to the post entities entity fid events params
func (o *PostEntitiesEntityFidEventsParams) SetPayload(payload *models.TriggerActionPayload) {
	o.Payload = payload
}

// WriteToRequest writes these params to a swagger request
func (o *PostEntitiesEntityFidEventsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param entityFid
	if err := r.SetPathParam("entityFid", o.EntityFid); err != nil {
		return err
	}
	if o.Payload != nil {
		if err := r.SetBodyParam(o.Payload); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
