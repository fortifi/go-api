// Code generated by go-swagger; DO NOT EDIT.

package account

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

// NewPostAccountVerificationParams creates a new PostAccountVerificationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostAccountVerificationParams() *PostAccountVerificationParams {
	return &PostAccountVerificationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostAccountVerificationParamsWithTimeout creates a new PostAccountVerificationParams object
// with the ability to set a timeout on a request.
func NewPostAccountVerificationParamsWithTimeout(timeout time.Duration) *PostAccountVerificationParams {
	return &PostAccountVerificationParams{
		timeout: timeout,
	}
}

// NewPostAccountVerificationParamsWithContext creates a new PostAccountVerificationParams object
// with the ability to set a context for a request.
func NewPostAccountVerificationParamsWithContext(ctx context.Context) *PostAccountVerificationParams {
	return &PostAccountVerificationParams{
		Context: ctx,
	}
}

// NewPostAccountVerificationParamsWithHTTPClient creates a new PostAccountVerificationParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostAccountVerificationParamsWithHTTPClient(client *http.Client) *PostAccountVerificationParams {
	return &PostAccountVerificationParams{
		HTTPClient: client,
	}
}

/*
PostAccountVerificationParams contains all the parameters to send to the API endpoint

	for the post account verification operation.

	Typically these are written to a http.Request.
*/
type PostAccountVerificationParams struct {

	// Payload.
	Payload *models.AccountVerificationPayload

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post account verification params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAccountVerificationParams) WithDefaults() *PostAccountVerificationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post account verification params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAccountVerificationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post account verification params
func (o *PostAccountVerificationParams) WithTimeout(timeout time.Duration) *PostAccountVerificationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post account verification params
func (o *PostAccountVerificationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post account verification params
func (o *PostAccountVerificationParams) WithContext(ctx context.Context) *PostAccountVerificationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post account verification params
func (o *PostAccountVerificationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post account verification params
func (o *PostAccountVerificationParams) WithHTTPClient(client *http.Client) *PostAccountVerificationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post account verification params
func (o *PostAccountVerificationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPayload adds the payload to the post account verification params
func (o *PostAccountVerificationParams) WithPayload(payload *models.AccountVerificationPayload) *PostAccountVerificationParams {
	o.SetPayload(payload)
	return o
}

// SetPayload adds the payload to the post account verification params
func (o *PostAccountVerificationParams) SetPayload(payload *models.AccountVerificationPayload) {
	o.Payload = payload
}

// WriteToRequest writes these params to a swagger request
func (o *PostAccountVerificationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
