// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewGetPayPublicKeyParams creates a new GetPayPublicKeyParams object
// with the default values initialized.
func NewGetPayPublicKeyParams() *GetPayPublicKeyParams {
	var ()
	return &GetPayPublicKeyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPayPublicKeyParamsWithTimeout creates a new GetPayPublicKeyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPayPublicKeyParamsWithTimeout(timeout time.Duration) *GetPayPublicKeyParams {
	var ()
	return &GetPayPublicKeyParams{

		timeout: timeout,
	}
}

// NewGetPayPublicKeyParamsWithContext creates a new GetPayPublicKeyParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPayPublicKeyParamsWithContext(ctx context.Context) *GetPayPublicKeyParams {
	var ()
	return &GetPayPublicKeyParams{

		Context: ctx,
	}
}

// NewGetPayPublicKeyParamsWithHTTPClient creates a new GetPayPublicKeyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPayPublicKeyParamsWithHTTPClient(client *http.Client) *GetPayPublicKeyParams {
	var ()
	return &GetPayPublicKeyParams{
		HTTPClient: client,
	}
}

/*GetPayPublicKeyParams contains all the parameters to send to the API endpoint
for the get pay public key operation typically these are written to a http.Request
*/
type GetPayPublicKeyParams struct {

	/*Format
	  Format for the generated key xml, raw, pkcs1 or pkcs8.

	*/
	Format *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get pay public key params
func (o *GetPayPublicKeyParams) WithTimeout(timeout time.Duration) *GetPayPublicKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get pay public key params
func (o *GetPayPublicKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get pay public key params
func (o *GetPayPublicKeyParams) WithContext(ctx context.Context) *GetPayPublicKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get pay public key params
func (o *GetPayPublicKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get pay public key params
func (o *GetPayPublicKeyParams) WithHTTPClient(client *http.Client) *GetPayPublicKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get pay public key params
func (o *GetPayPublicKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFormat adds the format to the get pay public key params
func (o *GetPayPublicKeyParams) WithFormat(format *string) *GetPayPublicKeyParams {
	o.SetFormat(format)
	return o
}

// SetFormat adds the format to the get pay public key params
func (o *GetPayPublicKeyParams) SetFormat(format *string) {
	o.Format = format
}

// WriteToRequest writes these params to a swagger request
func (o *GetPayPublicKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Format != nil {

		// query param format
		var qrFormat string
		if o.Format != nil {
			qrFormat = *o.Format
		}
		qFormat := qrFormat
		if qFormat != "" {
			if err := r.SetQueryParam("format", qFormat); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
