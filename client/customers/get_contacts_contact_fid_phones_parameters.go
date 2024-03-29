// Code generated by go-swagger; DO NOT EDIT.

package customers

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

// NewGetContactsContactFidPhonesParams creates a new GetContactsContactFidPhonesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetContactsContactFidPhonesParams() *GetContactsContactFidPhonesParams {
	return &GetContactsContactFidPhonesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetContactsContactFidPhonesParamsWithTimeout creates a new GetContactsContactFidPhonesParams object
// with the ability to set a timeout on a request.
func NewGetContactsContactFidPhonesParamsWithTimeout(timeout time.Duration) *GetContactsContactFidPhonesParams {
	return &GetContactsContactFidPhonesParams{
		timeout: timeout,
	}
}

// NewGetContactsContactFidPhonesParamsWithContext creates a new GetContactsContactFidPhonesParams object
// with the ability to set a context for a request.
func NewGetContactsContactFidPhonesParamsWithContext(ctx context.Context) *GetContactsContactFidPhonesParams {
	return &GetContactsContactFidPhonesParams{
		Context: ctx,
	}
}

// NewGetContactsContactFidPhonesParamsWithHTTPClient creates a new GetContactsContactFidPhonesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetContactsContactFidPhonesParamsWithHTTPClient(client *http.Client) *GetContactsContactFidPhonesParams {
	return &GetContactsContactFidPhonesParams{
		HTTPClient: client,
	}
}

/*
GetContactsContactFidPhonesParams contains all the parameters to send to the API endpoint

	for the get contacts contact fid phones operation.

	Typically these are written to a http.Request.
*/
type GetContactsContactFidPhonesParams struct {

	/* ContactFid.

	   Contact FID to use
	*/
	ContactFid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get contacts contact fid phones params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetContactsContactFidPhonesParams) WithDefaults() *GetContactsContactFidPhonesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get contacts contact fid phones params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetContactsContactFidPhonesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get contacts contact fid phones params
func (o *GetContactsContactFidPhonesParams) WithTimeout(timeout time.Duration) *GetContactsContactFidPhonesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get contacts contact fid phones params
func (o *GetContactsContactFidPhonesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get contacts contact fid phones params
func (o *GetContactsContactFidPhonesParams) WithContext(ctx context.Context) *GetContactsContactFidPhonesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get contacts contact fid phones params
func (o *GetContactsContactFidPhonesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get contacts contact fid phones params
func (o *GetContactsContactFidPhonesParams) WithHTTPClient(client *http.Client) *GetContactsContactFidPhonesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get contacts contact fid phones params
func (o *GetContactsContactFidPhonesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContactFid adds the contactFid to the get contacts contact fid phones params
func (o *GetContactsContactFidPhonesParams) WithContactFid(contactFid string) *GetContactsContactFidPhonesParams {
	o.SetContactFid(contactFid)
	return o
}

// SetContactFid adds the contactFid to the get contacts contact fid phones params
func (o *GetContactsContactFidPhonesParams) SetContactFid(contactFid string) {
	o.ContactFid = contactFid
}

// WriteToRequest writes these params to a swagger request
func (o *GetContactsContactFidPhonesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param contactFid
	if err := r.SetPathParam("contactFid", o.ContactFid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
