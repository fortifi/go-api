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

// NewPostEntitiesEntityFidEventsParams creates a new PostEntitiesEntityFidEventsParams object
// with the default values initialized.
func NewPostEntitiesEntityFidEventsParams() *PostEntitiesEntityFidEventsParams {
	var ()
	return &PostEntitiesEntityFidEventsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostEntitiesEntityFidEventsParamsWithTimeout creates a new PostEntitiesEntityFidEventsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostEntitiesEntityFidEventsParamsWithTimeout(timeout time.Duration) *PostEntitiesEntityFidEventsParams {
	var ()
	return &PostEntitiesEntityFidEventsParams{

		timeout: timeout,
	}
}

// NewPostEntitiesEntityFidEventsParamsWithContext creates a new PostEntitiesEntityFidEventsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostEntitiesEntityFidEventsParamsWithContext(ctx context.Context) *PostEntitiesEntityFidEventsParams {
	var ()
	return &PostEntitiesEntityFidEventsParams{

		Context: ctx,
	}
}

// NewPostEntitiesEntityFidEventsParamsWithHTTPClient creates a new PostEntitiesEntityFidEventsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostEntitiesEntityFidEventsParamsWithHTTPClient(client *http.Client) *PostEntitiesEntityFidEventsParams {
	var ()
	return &PostEntitiesEntityFidEventsParams{
		HTTPClient: client,
	}
}

/*PostEntitiesEntityFidEventsParams contains all the parameters to send to the API endpoint
for the post entities entity fid events operation typically these are written to a http.Request
*/
type PostEntitiesEntityFidEventsParams struct {

	/*EntityFid
	  Entity FID to use

	*/
	EntityFid string
	/*Payload*/
	Payload *models.TriggerActionPayload

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
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
