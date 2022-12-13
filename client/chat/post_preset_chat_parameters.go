// Code generated by go-swagger; DO NOT EDIT.

package chat

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

// NewPostPresetChatParams creates a new PostPresetChatParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostPresetChatParams() *PostPresetChatParams {
	return &PostPresetChatParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostPresetChatParamsWithTimeout creates a new PostPresetChatParams object
// with the ability to set a timeout on a request.
func NewPostPresetChatParamsWithTimeout(timeout time.Duration) *PostPresetChatParams {
	return &PostPresetChatParams{
		timeout: timeout,
	}
}

// NewPostPresetChatParamsWithContext creates a new PostPresetChatParams object
// with the ability to set a context for a request.
func NewPostPresetChatParamsWithContext(ctx context.Context) *PostPresetChatParams {
	return &PostPresetChatParams{
		Context: ctx,
	}
}

// NewPostPresetChatParamsWithHTTPClient creates a new PostPresetChatParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostPresetChatParamsWithHTTPClient(client *http.Client) *PostPresetChatParams {
	return &PostPresetChatParams{
		HTTPClient: client,
	}
}

/*
PostPresetChatParams contains all the parameters to send to the API endpoint

	for the post preset chat operation.

	Typically these are written to a http.Request.
*/
type PostPresetChatParams struct {

	// Payload.
	Payload *models.CreatePresetChatSessionPayload

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post preset chat params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostPresetChatParams) WithDefaults() *PostPresetChatParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post preset chat params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostPresetChatParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post preset chat params
func (o *PostPresetChatParams) WithTimeout(timeout time.Duration) *PostPresetChatParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post preset chat params
func (o *PostPresetChatParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post preset chat params
func (o *PostPresetChatParams) WithContext(ctx context.Context) *PostPresetChatParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post preset chat params
func (o *PostPresetChatParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post preset chat params
func (o *PostPresetChatParams) WithHTTPClient(client *http.Client) *PostPresetChatParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post preset chat params
func (o *PostPresetChatParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPayload adds the payload to the post preset chat params
func (o *PostPresetChatParams) WithPayload(payload *models.CreatePresetChatSessionPayload) *PostPresetChatParams {
	o.SetPayload(payload)
	return o
}

// SetPayload adds the payload to the post preset chat params
func (o *PostPresetChatParams) SetPayload(payload *models.CreatePresetChatSessionPayload) {
	o.Payload = payload
}

// WriteToRequest writes these params to a swagger request
func (o *PostPresetChatParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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