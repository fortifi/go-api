// Code generated by go-swagger; DO NOT EDIT.

package polymers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetPolymersPolymerFidParams creates a new GetPolymersPolymerFidParams object
// with the default values initialized.
func NewGetPolymersPolymerFidParams() *GetPolymersPolymerFidParams {
	var ()
	return &GetPolymersPolymerFidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPolymersPolymerFidParamsWithTimeout creates a new GetPolymersPolymerFidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPolymersPolymerFidParamsWithTimeout(timeout time.Duration) *GetPolymersPolymerFidParams {
	var ()
	return &GetPolymersPolymerFidParams{

		timeout: timeout,
	}
}

// NewGetPolymersPolymerFidParamsWithContext creates a new GetPolymersPolymerFidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPolymersPolymerFidParamsWithContext(ctx context.Context) *GetPolymersPolymerFidParams {
	var ()
	return &GetPolymersPolymerFidParams{

		Context: ctx,
	}
}

// NewGetPolymersPolymerFidParamsWithHTTPClient creates a new GetPolymersPolymerFidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPolymersPolymerFidParamsWithHTTPClient(client *http.Client) *GetPolymersPolymerFidParams {
	var ()
	return &GetPolymersPolymerFidParams{
		HTTPClient: client,
	}
}

/*GetPolymersPolymerFidParams contains all the parameters to send to the API endpoint
for the get polymers polymer fid operation typically these are written to a http.Request
*/
type GetPolymersPolymerFidParams struct {

	/*PolymerFid*/
	PolymerFid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get polymers polymer fid params
func (o *GetPolymersPolymerFidParams) WithTimeout(timeout time.Duration) *GetPolymersPolymerFidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get polymers polymer fid params
func (o *GetPolymersPolymerFidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get polymers polymer fid params
func (o *GetPolymersPolymerFidParams) WithContext(ctx context.Context) *GetPolymersPolymerFidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get polymers polymer fid params
func (o *GetPolymersPolymerFidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get polymers polymer fid params
func (o *GetPolymersPolymerFidParams) WithHTTPClient(client *http.Client) *GetPolymersPolymerFidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get polymers polymer fid params
func (o *GetPolymersPolymerFidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPolymerFid adds the polymerFid to the get polymers polymer fid params
func (o *GetPolymersPolymerFidParams) WithPolymerFid(polymerFid string) *GetPolymersPolymerFidParams {
	o.SetPolymerFid(polymerFid)
	return o
}

// SetPolymerFid adds the polymerFid to the get polymers polymer fid params
func (o *GetPolymersPolymerFidParams) SetPolymerFid(polymerFid string) {
	o.PolymerFid = polymerFid
}

// WriteToRequest writes these params to a swagger request
func (o *GetPolymersPolymerFidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param polymerFid
	if err := r.SetPathParam("polymerFid", o.PolymerFid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
