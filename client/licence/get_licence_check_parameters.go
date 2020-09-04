// Code generated by go-swagger; DO NOT EDIT.

package licence

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

// NewGetLicenceCheckParams creates a new GetLicenceCheckParams object
// with the default values initialized.
func NewGetLicenceCheckParams() *GetLicenceCheckParams {
	var ()
	return &GetLicenceCheckParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLicenceCheckParamsWithTimeout creates a new GetLicenceCheckParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLicenceCheckParamsWithTimeout(timeout time.Duration) *GetLicenceCheckParams {
	var ()
	return &GetLicenceCheckParams{

		timeout: timeout,
	}
}

// NewGetLicenceCheckParamsWithContext creates a new GetLicenceCheckParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLicenceCheckParamsWithContext(ctx context.Context) *GetLicenceCheckParams {
	var ()
	return &GetLicenceCheckParams{

		Context: ctx,
	}
}

// NewGetLicenceCheckParamsWithHTTPClient creates a new GetLicenceCheckParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLicenceCheckParamsWithHTTPClient(client *http.Client) *GetLicenceCheckParams {
	var ()
	return &GetLicenceCheckParams{
		HTTPClient: client,
	}
}

/*GetLicenceCheckParams contains all the parameters to send to the API endpoint
for the get licence check operation typically these are written to a http.Request
*/
type GetLicenceCheckParams struct {

	/*Key*/
	Key *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get licence check params
func (o *GetLicenceCheckParams) WithTimeout(timeout time.Duration) *GetLicenceCheckParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get licence check params
func (o *GetLicenceCheckParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get licence check params
func (o *GetLicenceCheckParams) WithContext(ctx context.Context) *GetLicenceCheckParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get licence check params
func (o *GetLicenceCheckParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get licence check params
func (o *GetLicenceCheckParams) WithHTTPClient(client *http.Client) *GetLicenceCheckParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get licence check params
func (o *GetLicenceCheckParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKey adds the key to the get licence check params
func (o *GetLicenceCheckParams) WithKey(key *string) *GetLicenceCheckParams {
	o.SetKey(key)
	return o
}

// SetKey adds the key to the get licence check params
func (o *GetLicenceCheckParams) SetKey(key *string) {
	o.Key = key
}

// WriteToRequest writes these params to a swagger request
func (o *GetLicenceCheckParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Key != nil {

		// query param key
		var qrKey string
		if o.Key != nil {
			qrKey = *o.Key
		}
		qKey := qrKey
		if qKey != "" {
			if err := r.SetQueryParam("key", qKey); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
