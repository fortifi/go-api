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

	"github.com/fortifi/go-api/models"
)

// NewPutCustomersCustomerFidAccountTypeParams creates a new PutCustomersCustomerFidAccountTypeParams object
// with the default values initialized.
func NewPutCustomersCustomerFidAccountTypeParams() *PutCustomersCustomerFidAccountTypeParams {
	var ()
	return &PutCustomersCustomerFidAccountTypeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutCustomersCustomerFidAccountTypeParamsWithTimeout creates a new PutCustomersCustomerFidAccountTypeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutCustomersCustomerFidAccountTypeParamsWithTimeout(timeout time.Duration) *PutCustomersCustomerFidAccountTypeParams {
	var ()
	return &PutCustomersCustomerFidAccountTypeParams{

		timeout: timeout,
	}
}

// NewPutCustomersCustomerFidAccountTypeParamsWithContext creates a new PutCustomersCustomerFidAccountTypeParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutCustomersCustomerFidAccountTypeParamsWithContext(ctx context.Context) *PutCustomersCustomerFidAccountTypeParams {
	var ()
	return &PutCustomersCustomerFidAccountTypeParams{

		Context: ctx,
	}
}

// NewPutCustomersCustomerFidAccountTypeParamsWithHTTPClient creates a new PutCustomersCustomerFidAccountTypeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutCustomersCustomerFidAccountTypeParamsWithHTTPClient(client *http.Client) *PutCustomersCustomerFidAccountTypeParams {
	var ()
	return &PutCustomersCustomerFidAccountTypeParams{
		HTTPClient: client,
	}
}

/*PutCustomersCustomerFidAccountTypeParams contains all the parameters to send to the API endpoint
for the put customers customer fid account type operation typically these are written to a http.Request
*/
type PutCustomersCustomerFidAccountTypeParams struct {

	/*CustomerFid
	  Customer FID to use

	*/
	CustomerFid string
	/*Payload*/
	Payload *models.SetAccountTypePayload

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put customers customer fid account type params
func (o *PutCustomersCustomerFidAccountTypeParams) WithTimeout(timeout time.Duration) *PutCustomersCustomerFidAccountTypeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put customers customer fid account type params
func (o *PutCustomersCustomerFidAccountTypeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put customers customer fid account type params
func (o *PutCustomersCustomerFidAccountTypeParams) WithContext(ctx context.Context) *PutCustomersCustomerFidAccountTypeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put customers customer fid account type params
func (o *PutCustomersCustomerFidAccountTypeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put customers customer fid account type params
func (o *PutCustomersCustomerFidAccountTypeParams) WithHTTPClient(client *http.Client) *PutCustomersCustomerFidAccountTypeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put customers customer fid account type params
func (o *PutCustomersCustomerFidAccountTypeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerFid adds the customerFid to the put customers customer fid account type params
func (o *PutCustomersCustomerFidAccountTypeParams) WithCustomerFid(customerFid string) *PutCustomersCustomerFidAccountTypeParams {
	o.SetCustomerFid(customerFid)
	return o
}

// SetCustomerFid adds the customerFid to the put customers customer fid account type params
func (o *PutCustomersCustomerFidAccountTypeParams) SetCustomerFid(customerFid string) {
	o.CustomerFid = customerFid
}

// WithPayload adds the payload to the put customers customer fid account type params
func (o *PutCustomersCustomerFidAccountTypeParams) WithPayload(payload *models.SetAccountTypePayload) *PutCustomersCustomerFidAccountTypeParams {
	o.SetPayload(payload)
	return o
}

// SetPayload adds the payload to the put customers customer fid account type params
func (o *PutCustomersCustomerFidAccountTypeParams) SetPayload(payload *models.SetAccountTypePayload) {
	o.Payload = payload
}

// WriteToRequest writes these params to a swagger request
func (o *PutCustomersCustomerFidAccountTypeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param customerFid
	if err := r.SetPathParam("customerFid", o.CustomerFid); err != nil {
		return err
	}

	if o.Payload != nil {
		if err := r.SetBodyParam(o.Payload); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
