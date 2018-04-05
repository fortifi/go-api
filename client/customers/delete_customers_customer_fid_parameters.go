// Code generated by go-swagger; DO NOT EDIT.

package customers

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

// NewDeleteCustomersCustomerFidParams creates a new DeleteCustomersCustomerFidParams object
// with the default values initialized.
func NewDeleteCustomersCustomerFidParams() *DeleteCustomersCustomerFidParams {
	var ()
	return &DeleteCustomersCustomerFidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCustomersCustomerFidParamsWithTimeout creates a new DeleteCustomersCustomerFidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteCustomersCustomerFidParamsWithTimeout(timeout time.Duration) *DeleteCustomersCustomerFidParams {
	var ()
	return &DeleteCustomersCustomerFidParams{

		timeout: timeout,
	}
}

// NewDeleteCustomersCustomerFidParamsWithContext creates a new DeleteCustomersCustomerFidParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteCustomersCustomerFidParamsWithContext(ctx context.Context) *DeleteCustomersCustomerFidParams {
	var ()
	return &DeleteCustomersCustomerFidParams{

		Context: ctx,
	}
}

// NewDeleteCustomersCustomerFidParamsWithHTTPClient creates a new DeleteCustomersCustomerFidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteCustomersCustomerFidParamsWithHTTPClient(client *http.Client) *DeleteCustomersCustomerFidParams {
	var ()
	return &DeleteCustomersCustomerFidParams{
		HTTPClient: client,
	}
}

/*DeleteCustomersCustomerFidParams contains all the parameters to send to the API endpoint
for the delete customers customer fid operation typically these are written to a http.Request
*/
type DeleteCustomersCustomerFidParams struct {

	/*CustomerFid
	  Customer FID to use

	*/
	CustomerFid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete customers customer fid params
func (o *DeleteCustomersCustomerFidParams) WithTimeout(timeout time.Duration) *DeleteCustomersCustomerFidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete customers customer fid params
func (o *DeleteCustomersCustomerFidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete customers customer fid params
func (o *DeleteCustomersCustomerFidParams) WithContext(ctx context.Context) *DeleteCustomersCustomerFidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete customers customer fid params
func (o *DeleteCustomersCustomerFidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete customers customer fid params
func (o *DeleteCustomersCustomerFidParams) WithHTTPClient(client *http.Client) *DeleteCustomersCustomerFidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete customers customer fid params
func (o *DeleteCustomersCustomerFidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerFid adds the customerFid to the delete customers customer fid params
func (o *DeleteCustomersCustomerFidParams) WithCustomerFid(customerFid string) *DeleteCustomersCustomerFidParams {
	o.SetCustomerFid(customerFid)
	return o
}

// SetCustomerFid adds the customerFid to the delete customers customer fid params
func (o *DeleteCustomersCustomerFidParams) SetCustomerFid(customerFid string) {
	o.CustomerFid = customerFid
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCustomersCustomerFidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param customerFid
	if err := r.SetPathParam("customerFid", o.CustomerFid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}