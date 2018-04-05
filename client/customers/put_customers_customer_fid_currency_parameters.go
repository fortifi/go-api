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

// NewPutCustomersCustomerFidCurrencyParams creates a new PutCustomersCustomerFidCurrencyParams object
// with the default values initialized.
func NewPutCustomersCustomerFidCurrencyParams() *PutCustomersCustomerFidCurrencyParams {
	var ()
	return &PutCustomersCustomerFidCurrencyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutCustomersCustomerFidCurrencyParamsWithTimeout creates a new PutCustomersCustomerFidCurrencyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutCustomersCustomerFidCurrencyParamsWithTimeout(timeout time.Duration) *PutCustomersCustomerFidCurrencyParams {
	var ()
	return &PutCustomersCustomerFidCurrencyParams{

		timeout: timeout,
	}
}

// NewPutCustomersCustomerFidCurrencyParamsWithContext creates a new PutCustomersCustomerFidCurrencyParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutCustomersCustomerFidCurrencyParamsWithContext(ctx context.Context) *PutCustomersCustomerFidCurrencyParams {
	var ()
	return &PutCustomersCustomerFidCurrencyParams{

		Context: ctx,
	}
}

// NewPutCustomersCustomerFidCurrencyParamsWithHTTPClient creates a new PutCustomersCustomerFidCurrencyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutCustomersCustomerFidCurrencyParamsWithHTTPClient(client *http.Client) *PutCustomersCustomerFidCurrencyParams {
	var ()
	return &PutCustomersCustomerFidCurrencyParams{
		HTTPClient: client,
	}
}

/*PutCustomersCustomerFidCurrencyParams contains all the parameters to send to the API endpoint
for the put customers customer fid currency operation typically these are written to a http.Request
*/
type PutCustomersCustomerFidCurrencyParams struct {

	/*Currency
	  Currency (ISO 4217, 3 Character Code)

	*/
	Currency string
	/*CustomerFid
	  Customer FID to use

	*/
	CustomerFid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put customers customer fid currency params
func (o *PutCustomersCustomerFidCurrencyParams) WithTimeout(timeout time.Duration) *PutCustomersCustomerFidCurrencyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put customers customer fid currency params
func (o *PutCustomersCustomerFidCurrencyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put customers customer fid currency params
func (o *PutCustomersCustomerFidCurrencyParams) WithContext(ctx context.Context) *PutCustomersCustomerFidCurrencyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put customers customer fid currency params
func (o *PutCustomersCustomerFidCurrencyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put customers customer fid currency params
func (o *PutCustomersCustomerFidCurrencyParams) WithHTTPClient(client *http.Client) *PutCustomersCustomerFidCurrencyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put customers customer fid currency params
func (o *PutCustomersCustomerFidCurrencyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCurrency adds the currency to the put customers customer fid currency params
func (o *PutCustomersCustomerFidCurrencyParams) WithCurrency(currency string) *PutCustomersCustomerFidCurrencyParams {
	o.SetCurrency(currency)
	return o
}

// SetCurrency adds the currency to the put customers customer fid currency params
func (o *PutCustomersCustomerFidCurrencyParams) SetCurrency(currency string) {
	o.Currency = currency
}

// WithCustomerFid adds the customerFid to the put customers customer fid currency params
func (o *PutCustomersCustomerFidCurrencyParams) WithCustomerFid(customerFid string) *PutCustomersCustomerFidCurrencyParams {
	o.SetCustomerFid(customerFid)
	return o
}

// SetCustomerFid adds the customerFid to the put customers customer fid currency params
func (o *PutCustomersCustomerFidCurrencyParams) SetCustomerFid(customerFid string) {
	o.CustomerFid = customerFid
}

// WriteToRequest writes these params to a swagger request
func (o *PutCustomersCustomerFidCurrencyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// form param currency
	frCurrency := o.Currency
	fCurrency := frCurrency
	if fCurrency != "" {
		if err := r.SetFormParam("currency", fCurrency); err != nil {
			return err
		}
	}

	// path param customerFid
	if err := r.SetPathParam("customerFid", o.CustomerFid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
