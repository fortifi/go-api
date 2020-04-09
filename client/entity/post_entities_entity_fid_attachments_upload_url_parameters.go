// Code generated by go-swagger; DO NOT EDIT.

package entity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/fortifi/go-api/models"
)

// NewPostEntitiesEntityFidAttachmentsUploadURLParams creates a new PostEntitiesEntityFidAttachmentsUploadURLParams object
// with the default values initialized.
func NewPostEntitiesEntityFidAttachmentsUploadURLParams() *PostEntitiesEntityFidAttachmentsUploadURLParams {
	var ()
	return &PostEntitiesEntityFidAttachmentsUploadURLParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostEntitiesEntityFidAttachmentsUploadURLParamsWithTimeout creates a new PostEntitiesEntityFidAttachmentsUploadURLParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostEntitiesEntityFidAttachmentsUploadURLParamsWithTimeout(timeout time.Duration) *PostEntitiesEntityFidAttachmentsUploadURLParams {
	var ()
	return &PostEntitiesEntityFidAttachmentsUploadURLParams{

		timeout: timeout,
	}
}

// NewPostEntitiesEntityFidAttachmentsUploadURLParamsWithContext creates a new PostEntitiesEntityFidAttachmentsUploadURLParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostEntitiesEntityFidAttachmentsUploadURLParamsWithContext(ctx context.Context) *PostEntitiesEntityFidAttachmentsUploadURLParams {
	var ()
	return &PostEntitiesEntityFidAttachmentsUploadURLParams{

		Context: ctx,
	}
}

// NewPostEntitiesEntityFidAttachmentsUploadURLParamsWithHTTPClient creates a new PostEntitiesEntityFidAttachmentsUploadURLParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostEntitiesEntityFidAttachmentsUploadURLParamsWithHTTPClient(client *http.Client) *PostEntitiesEntityFidAttachmentsUploadURLParams {
	var ()
	return &PostEntitiesEntityFidAttachmentsUploadURLParams{
		HTTPClient: client,
	}
}

/*PostEntitiesEntityFidAttachmentsUploadURLParams contains all the parameters to send to the API endpoint
for the post entities entity fid attachments upload URL operation typically these are written to a http.Request
*/
type PostEntitiesEntityFidAttachmentsUploadURLParams struct {

	/*EntityFid
	  Entity FID to use

	*/
	EntityFid string
	/*Payload*/
	Payload *models.RequestUploadURLPayload

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post entities entity fid attachments upload URL params
func (o *PostEntitiesEntityFidAttachmentsUploadURLParams) WithTimeout(timeout time.Duration) *PostEntitiesEntityFidAttachmentsUploadURLParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post entities entity fid attachments upload URL params
func (o *PostEntitiesEntityFidAttachmentsUploadURLParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post entities entity fid attachments upload URL params
func (o *PostEntitiesEntityFidAttachmentsUploadURLParams) WithContext(ctx context.Context) *PostEntitiesEntityFidAttachmentsUploadURLParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post entities entity fid attachments upload URL params
func (o *PostEntitiesEntityFidAttachmentsUploadURLParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post entities entity fid attachments upload URL params
func (o *PostEntitiesEntityFidAttachmentsUploadURLParams) WithHTTPClient(client *http.Client) *PostEntitiesEntityFidAttachmentsUploadURLParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post entities entity fid attachments upload URL params
func (o *PostEntitiesEntityFidAttachmentsUploadURLParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEntityFid adds the entityFid to the post entities entity fid attachments upload URL params
func (o *PostEntitiesEntityFidAttachmentsUploadURLParams) WithEntityFid(entityFid string) *PostEntitiesEntityFidAttachmentsUploadURLParams {
	o.SetEntityFid(entityFid)
	return o
}

// SetEntityFid adds the entityFid to the post entities entity fid attachments upload URL params
func (o *PostEntitiesEntityFidAttachmentsUploadURLParams) SetEntityFid(entityFid string) {
	o.EntityFid = entityFid
}

// WithPayload adds the payload to the post entities entity fid attachments upload URL params
func (o *PostEntitiesEntityFidAttachmentsUploadURLParams) WithPayload(payload *models.RequestUploadURLPayload) *PostEntitiesEntityFidAttachmentsUploadURLParams {
	o.SetPayload(payload)
	return o
}

// SetPayload adds the payload to the post entities entity fid attachments upload URL params
func (o *PostEntitiesEntityFidAttachmentsUploadURLParams) SetPayload(payload *models.RequestUploadURLPayload) {
	o.Payload = payload
}

// WriteToRequest writes these params to a swagger request
func (o *PostEntitiesEntityFidAttachmentsUploadURLParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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