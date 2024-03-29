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

// NewPutCustomersCustomerFidLoyalParams creates a new PutCustomersCustomerFidLoyalParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutCustomersCustomerFidLoyalParams() *PutCustomersCustomerFidLoyalParams {
	return &PutCustomersCustomerFidLoyalParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutCustomersCustomerFidLoyalParamsWithTimeout creates a new PutCustomersCustomerFidLoyalParams object
// with the ability to set a timeout on a request.
func NewPutCustomersCustomerFidLoyalParamsWithTimeout(timeout time.Duration) *PutCustomersCustomerFidLoyalParams {
	return &PutCustomersCustomerFidLoyalParams{
		timeout: timeout,
	}
}

// NewPutCustomersCustomerFidLoyalParamsWithContext creates a new PutCustomersCustomerFidLoyalParams object
// with the ability to set a context for a request.
func NewPutCustomersCustomerFidLoyalParamsWithContext(ctx context.Context) *PutCustomersCustomerFidLoyalParams {
	return &PutCustomersCustomerFidLoyalParams{
		Context: ctx,
	}
}

// NewPutCustomersCustomerFidLoyalParamsWithHTTPClient creates a new PutCustomersCustomerFidLoyalParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutCustomersCustomerFidLoyalParamsWithHTTPClient(client *http.Client) *PutCustomersCustomerFidLoyalParams {
	return &PutCustomersCustomerFidLoyalParams{
		HTTPClient: client,
	}
}

/*
PutCustomersCustomerFidLoyalParams contains all the parameters to send to the API endpoint

	for the put customers customer fid loyal operation.

	Typically these are written to a http.Request.
*/
type PutCustomersCustomerFidLoyalParams struct {

	/* CustomerFid.

	   Customer FID to use
	*/
	CustomerFid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put customers customer fid loyal params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutCustomersCustomerFidLoyalParams) WithDefaults() *PutCustomersCustomerFidLoyalParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put customers customer fid loyal params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutCustomersCustomerFidLoyalParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put customers customer fid loyal params
func (o *PutCustomersCustomerFidLoyalParams) WithTimeout(timeout time.Duration) *PutCustomersCustomerFidLoyalParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put customers customer fid loyal params
func (o *PutCustomersCustomerFidLoyalParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put customers customer fid loyal params
func (o *PutCustomersCustomerFidLoyalParams) WithContext(ctx context.Context) *PutCustomersCustomerFidLoyalParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put customers customer fid loyal params
func (o *PutCustomersCustomerFidLoyalParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put customers customer fid loyal params
func (o *PutCustomersCustomerFidLoyalParams) WithHTTPClient(client *http.Client) *PutCustomersCustomerFidLoyalParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put customers customer fid loyal params
func (o *PutCustomersCustomerFidLoyalParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerFid adds the customerFid to the put customers customer fid loyal params
func (o *PutCustomersCustomerFidLoyalParams) WithCustomerFid(customerFid string) *PutCustomersCustomerFidLoyalParams {
	o.SetCustomerFid(customerFid)
	return o
}

// SetCustomerFid adds the customerFid to the put customers customer fid loyal params
func (o *PutCustomersCustomerFidLoyalParams) SetCustomerFid(customerFid string) {
	o.CustomerFid = customerFid
}

// WriteToRequest writes these params to a swagger request
func (o *PutCustomersCustomerFidLoyalParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
