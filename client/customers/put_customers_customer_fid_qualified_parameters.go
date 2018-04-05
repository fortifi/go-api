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

// NewPutCustomersCustomerFidQualifiedParams creates a new PutCustomersCustomerFidQualifiedParams object
// with the default values initialized.
func NewPutCustomersCustomerFidQualifiedParams() *PutCustomersCustomerFidQualifiedParams {
	var ()
	return &PutCustomersCustomerFidQualifiedParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutCustomersCustomerFidQualifiedParamsWithTimeout creates a new PutCustomersCustomerFidQualifiedParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutCustomersCustomerFidQualifiedParamsWithTimeout(timeout time.Duration) *PutCustomersCustomerFidQualifiedParams {
	var ()
	return &PutCustomersCustomerFidQualifiedParams{

		timeout: timeout,
	}
}

// NewPutCustomersCustomerFidQualifiedParamsWithContext creates a new PutCustomersCustomerFidQualifiedParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutCustomersCustomerFidQualifiedParamsWithContext(ctx context.Context) *PutCustomersCustomerFidQualifiedParams {
	var ()
	return &PutCustomersCustomerFidQualifiedParams{

		Context: ctx,
	}
}

// NewPutCustomersCustomerFidQualifiedParamsWithHTTPClient creates a new PutCustomersCustomerFidQualifiedParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutCustomersCustomerFidQualifiedParamsWithHTTPClient(client *http.Client) *PutCustomersCustomerFidQualifiedParams {
	var ()
	return &PutCustomersCustomerFidQualifiedParams{
		HTTPClient: client,
	}
}

/*PutCustomersCustomerFidQualifiedParams contains all the parameters to send to the API endpoint
for the put customers customer fid qualified operation typically these are written to a http.Request
*/
type PutCustomersCustomerFidQualifiedParams struct {

	/*CustomerFid
	  Customer FID to use

	*/
	CustomerFid string
	/*IsoTime
	  Time in ISO 8601 standard with optional fractions of a second e.g 2015-12-05T13:11:59.888Z

	*/
	IsoTime *strfmt.DateTime

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put customers customer fid qualified params
func (o *PutCustomersCustomerFidQualifiedParams) WithTimeout(timeout time.Duration) *PutCustomersCustomerFidQualifiedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put customers customer fid qualified params
func (o *PutCustomersCustomerFidQualifiedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put customers customer fid qualified params
func (o *PutCustomersCustomerFidQualifiedParams) WithContext(ctx context.Context) *PutCustomersCustomerFidQualifiedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put customers customer fid qualified params
func (o *PutCustomersCustomerFidQualifiedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put customers customer fid qualified params
func (o *PutCustomersCustomerFidQualifiedParams) WithHTTPClient(client *http.Client) *PutCustomersCustomerFidQualifiedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put customers customer fid qualified params
func (o *PutCustomersCustomerFidQualifiedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerFid adds the customerFid to the put customers customer fid qualified params
func (o *PutCustomersCustomerFidQualifiedParams) WithCustomerFid(customerFid string) *PutCustomersCustomerFidQualifiedParams {
	o.SetCustomerFid(customerFid)
	return o
}

// SetCustomerFid adds the customerFid to the put customers customer fid qualified params
func (o *PutCustomersCustomerFidQualifiedParams) SetCustomerFid(customerFid string) {
	o.CustomerFid = customerFid
}

// WithIsoTime adds the isoTime to the put customers customer fid qualified params
func (o *PutCustomersCustomerFidQualifiedParams) WithIsoTime(isoTime *strfmt.DateTime) *PutCustomersCustomerFidQualifiedParams {
	o.SetIsoTime(isoTime)
	return o
}

// SetIsoTime adds the isoTime to the put customers customer fid qualified params
func (o *PutCustomersCustomerFidQualifiedParams) SetIsoTime(isoTime *strfmt.DateTime) {
	o.IsoTime = isoTime
}

// WriteToRequest writes these params to a swagger request
func (o *PutCustomersCustomerFidQualifiedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param customerFid
	if err := r.SetPathParam("customerFid", o.CustomerFid); err != nil {
		return err
	}

	if o.IsoTime != nil {

		// form param isoTime
		var frIsoTime strfmt.DateTime
		if o.IsoTime != nil {
			frIsoTime = *o.IsoTime
		}
		fIsoTime := frIsoTime.String()
		if fIsoTime != "" {
			if err := r.SetFormParam("isoTime", fIsoTime); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
