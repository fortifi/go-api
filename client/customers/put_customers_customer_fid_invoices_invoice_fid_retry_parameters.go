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

// NewPutCustomersCustomerFidInvoicesInvoiceFidRetryParams creates a new PutCustomersCustomerFidInvoicesInvoiceFidRetryParams object
// with the default values initialized.
func NewPutCustomersCustomerFidInvoicesInvoiceFidRetryParams() *PutCustomersCustomerFidInvoicesInvoiceFidRetryParams {
	var ()
	return &PutCustomersCustomerFidInvoicesInvoiceFidRetryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutCustomersCustomerFidInvoicesInvoiceFidRetryParamsWithTimeout creates a new PutCustomersCustomerFidInvoicesInvoiceFidRetryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutCustomersCustomerFidInvoicesInvoiceFidRetryParamsWithTimeout(timeout time.Duration) *PutCustomersCustomerFidInvoicesInvoiceFidRetryParams {
	var ()
	return &PutCustomersCustomerFidInvoicesInvoiceFidRetryParams{

		timeout: timeout,
	}
}

// NewPutCustomersCustomerFidInvoicesInvoiceFidRetryParamsWithContext creates a new PutCustomersCustomerFidInvoicesInvoiceFidRetryParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutCustomersCustomerFidInvoicesInvoiceFidRetryParamsWithContext(ctx context.Context) *PutCustomersCustomerFidInvoicesInvoiceFidRetryParams {
	var ()
	return &PutCustomersCustomerFidInvoicesInvoiceFidRetryParams{

		Context: ctx,
	}
}

// NewPutCustomersCustomerFidInvoicesInvoiceFidRetryParamsWithHTTPClient creates a new PutCustomersCustomerFidInvoicesInvoiceFidRetryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutCustomersCustomerFidInvoicesInvoiceFidRetryParamsWithHTTPClient(client *http.Client) *PutCustomersCustomerFidInvoicesInvoiceFidRetryParams {
	var ()
	return &PutCustomersCustomerFidInvoicesInvoiceFidRetryParams{
		HTTPClient: client,
	}
}

/*PutCustomersCustomerFidInvoicesInvoiceFidRetryParams contains all the parameters to send to the API endpoint
for the put customers customer fid invoices invoice fid retry operation typically these are written to a http.Request
*/
type PutCustomersCustomerFidInvoicesInvoiceFidRetryParams struct {

	/*CustomerFid
	  Customer FID to use

	*/
	CustomerFid string
	/*InvoiceFid
	  Invoice Fid to use

	*/
	InvoiceFid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put customers customer fid invoices invoice fid retry params
func (o *PutCustomersCustomerFidInvoicesInvoiceFidRetryParams) WithTimeout(timeout time.Duration) *PutCustomersCustomerFidInvoicesInvoiceFidRetryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put customers customer fid invoices invoice fid retry params
func (o *PutCustomersCustomerFidInvoicesInvoiceFidRetryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put customers customer fid invoices invoice fid retry params
func (o *PutCustomersCustomerFidInvoicesInvoiceFidRetryParams) WithContext(ctx context.Context) *PutCustomersCustomerFidInvoicesInvoiceFidRetryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put customers customer fid invoices invoice fid retry params
func (o *PutCustomersCustomerFidInvoicesInvoiceFidRetryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put customers customer fid invoices invoice fid retry params
func (o *PutCustomersCustomerFidInvoicesInvoiceFidRetryParams) WithHTTPClient(client *http.Client) *PutCustomersCustomerFidInvoicesInvoiceFidRetryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put customers customer fid invoices invoice fid retry params
func (o *PutCustomersCustomerFidInvoicesInvoiceFidRetryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerFid adds the customerFid to the put customers customer fid invoices invoice fid retry params
func (o *PutCustomersCustomerFidInvoicesInvoiceFidRetryParams) WithCustomerFid(customerFid string) *PutCustomersCustomerFidInvoicesInvoiceFidRetryParams {
	o.SetCustomerFid(customerFid)
	return o
}

// SetCustomerFid adds the customerFid to the put customers customer fid invoices invoice fid retry params
func (o *PutCustomersCustomerFidInvoicesInvoiceFidRetryParams) SetCustomerFid(customerFid string) {
	o.CustomerFid = customerFid
}

// WithInvoiceFid adds the invoiceFid to the put customers customer fid invoices invoice fid retry params
func (o *PutCustomersCustomerFidInvoicesInvoiceFidRetryParams) WithInvoiceFid(invoiceFid string) *PutCustomersCustomerFidInvoicesInvoiceFidRetryParams {
	o.SetInvoiceFid(invoiceFid)
	return o
}

// SetInvoiceFid adds the invoiceFid to the put customers customer fid invoices invoice fid retry params
func (o *PutCustomersCustomerFidInvoicesInvoiceFidRetryParams) SetInvoiceFid(invoiceFid string) {
	o.InvoiceFid = invoiceFid
}

// WriteToRequest writes these params to a swagger request
func (o *PutCustomersCustomerFidInvoicesInvoiceFidRetryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param customerFid
	if err := r.SetPathParam("customerFid", o.CustomerFid); err != nil {
		return err
	}

	// path param invoiceFid
	if err := r.SetPathParam("invoiceFid", o.InvoiceFid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
