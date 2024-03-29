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
)

// NewGetPresetChatSessionIDParams creates a new GetPresetChatSessionIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPresetChatSessionIDParams() *GetPresetChatSessionIDParams {
	return &GetPresetChatSessionIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPresetChatSessionIDParamsWithTimeout creates a new GetPresetChatSessionIDParams object
// with the ability to set a timeout on a request.
func NewGetPresetChatSessionIDParamsWithTimeout(timeout time.Duration) *GetPresetChatSessionIDParams {
	return &GetPresetChatSessionIDParams{
		timeout: timeout,
	}
}

// NewGetPresetChatSessionIDParamsWithContext creates a new GetPresetChatSessionIDParams object
// with the ability to set a context for a request.
func NewGetPresetChatSessionIDParamsWithContext(ctx context.Context) *GetPresetChatSessionIDParams {
	return &GetPresetChatSessionIDParams{
		Context: ctx,
	}
}

// NewGetPresetChatSessionIDParamsWithHTTPClient creates a new GetPresetChatSessionIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPresetChatSessionIDParamsWithHTTPClient(client *http.Client) *GetPresetChatSessionIDParams {
	return &GetPresetChatSessionIDParams{
		HTTPClient: client,
	}
}

/*
GetPresetChatSessionIDParams contains all the parameters to send to the API endpoint

	for the get preset chat session ID operation.

	Typically these are written to a http.Request.
*/
type GetPresetChatSessionIDParams struct {

	/* SessionID.

	   The session id of a fortifi chat
	*/
	SessionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get preset chat session ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPresetChatSessionIDParams) WithDefaults() *GetPresetChatSessionIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get preset chat session ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPresetChatSessionIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get preset chat session ID params
func (o *GetPresetChatSessionIDParams) WithTimeout(timeout time.Duration) *GetPresetChatSessionIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get preset chat session ID params
func (o *GetPresetChatSessionIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get preset chat session ID params
func (o *GetPresetChatSessionIDParams) WithContext(ctx context.Context) *GetPresetChatSessionIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get preset chat session ID params
func (o *GetPresetChatSessionIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get preset chat session ID params
func (o *GetPresetChatSessionIDParams) WithHTTPClient(client *http.Client) *GetPresetChatSessionIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get preset chat session ID params
func (o *GetPresetChatSessionIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSessionID adds the sessionID to the get preset chat session ID params
func (o *GetPresetChatSessionIDParams) WithSessionID(sessionID string) *GetPresetChatSessionIDParams {
	o.SetSessionID(sessionID)
	return o
}

// SetSessionID adds the sessionId to the get preset chat session ID params
func (o *GetPresetChatSessionIDParams) SetSessionID(sessionID string) {
	o.SessionID = sessionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetPresetChatSessionIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param sessionId
	if err := r.SetPathParam("sessionId", o.SessionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
