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

// NewGetContactsContactFidEmailsParams creates a new GetContactsContactFidEmailsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetContactsContactFidEmailsParams() *GetContactsContactFidEmailsParams {
	return &GetContactsContactFidEmailsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetContactsContactFidEmailsParamsWithTimeout creates a new GetContactsContactFidEmailsParams object
// with the ability to set a timeout on a request.
func NewGetContactsContactFidEmailsParamsWithTimeout(timeout time.Duration) *GetContactsContactFidEmailsParams {
	return &GetContactsContactFidEmailsParams{
		timeout: timeout,
	}
}

// NewGetContactsContactFidEmailsParamsWithContext creates a new GetContactsContactFidEmailsParams object
// with the ability to set a context for a request.
func NewGetContactsContactFidEmailsParamsWithContext(ctx context.Context) *GetContactsContactFidEmailsParams {
	return &GetContactsContactFidEmailsParams{
		Context: ctx,
	}
}

// NewGetContactsContactFidEmailsParamsWithHTTPClient creates a new GetContactsContactFidEmailsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetContactsContactFidEmailsParamsWithHTTPClient(client *http.Client) *GetContactsContactFidEmailsParams {
	return &GetContactsContactFidEmailsParams{
		HTTPClient: client,
	}
}

/*
GetContactsContactFidEmailsParams contains all the parameters to send to the API endpoint

	for the get contacts contact fid emails operation.

	Typically these are written to a http.Request.
*/
type GetContactsContactFidEmailsParams struct {

	/* ContactFid.

	   Contact FID to use
	*/
	ContactFid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get contacts contact fid emails params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetContactsContactFidEmailsParams) WithDefaults() *GetContactsContactFidEmailsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get contacts contact fid emails params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetContactsContactFidEmailsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get contacts contact fid emails params
func (o *GetContactsContactFidEmailsParams) WithTimeout(timeout time.Duration) *GetContactsContactFidEmailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get contacts contact fid emails params
func (o *GetContactsContactFidEmailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get contacts contact fid emails params
func (o *GetContactsContactFidEmailsParams) WithContext(ctx context.Context) *GetContactsContactFidEmailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get contacts contact fid emails params
func (o *GetContactsContactFidEmailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get contacts contact fid emails params
func (o *GetContactsContactFidEmailsParams) WithHTTPClient(client *http.Client) *GetContactsContactFidEmailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get contacts contact fid emails params
func (o *GetContactsContactFidEmailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContactFid adds the contactFid to the get contacts contact fid emails params
func (o *GetContactsContactFidEmailsParams) WithContactFid(contactFid string) *GetContactsContactFidEmailsParams {
	o.SetContactFid(contactFid)
	return o
}

// SetContactFid adds the contactFid to the get contacts contact fid emails params
func (o *GetContactsContactFidEmailsParams) SetContactFid(contactFid string) {
	o.ContactFid = contactFid
}

// WriteToRequest writes these params to a swagger request
func (o *GetContactsContactFidEmailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
