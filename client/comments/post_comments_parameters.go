// Code generated by go-swagger; DO NOT EDIT.

package comments

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

// NewPostCommentsParams creates a new PostCommentsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostCommentsParams() *PostCommentsParams {
	return &PostCommentsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostCommentsParamsWithTimeout creates a new PostCommentsParams object
// with the ability to set a timeout on a request.
func NewPostCommentsParamsWithTimeout(timeout time.Duration) *PostCommentsParams {
	return &PostCommentsParams{
		timeout: timeout,
	}
}

// NewPostCommentsParamsWithContext creates a new PostCommentsParams object
// with the ability to set a context for a request.
func NewPostCommentsParamsWithContext(ctx context.Context) *PostCommentsParams {
	return &PostCommentsParams{
		Context: ctx,
	}
}

// NewPostCommentsParamsWithHTTPClient creates a new PostCommentsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostCommentsParamsWithHTTPClient(client *http.Client) *PostCommentsParams {
	return &PostCommentsParams{
		HTTPClient: client,
	}
}

/*
PostCommentsParams contains all the parameters to send to the API endpoint

	for the post comments operation.

	Typically these are written to a http.Request.
*/
type PostCommentsParams struct {

	// Payload.
	Payload *models.CreateCommentPayload

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post comments params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostCommentsParams) WithDefaults() *PostCommentsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post comments params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostCommentsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post comments params
func (o *PostCommentsParams) WithTimeout(timeout time.Duration) *PostCommentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post comments params
func (o *PostCommentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post comments params
func (o *PostCommentsParams) WithContext(ctx context.Context) *PostCommentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post comments params
func (o *PostCommentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post comments params
func (o *PostCommentsParams) WithHTTPClient(client *http.Client) *PostCommentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post comments params
func (o *PostCommentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPayload adds the payload to the post comments params
func (o *PostCommentsParams) WithPayload(payload *models.CreateCommentPayload) *PostCommentsParams {
	o.SetPayload(payload)
	return o
}

// SetPayload adds the payload to the post comments params
func (o *PostCommentsParams) SetPayload(payload *models.CreateCommentPayload) {
	o.Payload = payload
}

// WriteToRequest writes these params to a swagger request
func (o *PostCommentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
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
