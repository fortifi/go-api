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

// NewPutCustomersCustomerFidPurchasedParams creates a new PutCustomersCustomerFidPurchasedParams object
// with the default values initialized.
func NewPutCustomersCustomerFidPurchasedParams() *PutCustomersCustomerFidPurchasedParams {
	var ()
	return &PutCustomersCustomerFidPurchasedParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutCustomersCustomerFidPurchasedParamsWithTimeout creates a new PutCustomersCustomerFidPurchasedParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutCustomersCustomerFidPurchasedParamsWithTimeout(timeout time.Duration) *PutCustomersCustomerFidPurchasedParams {
	var ()
	return &PutCustomersCustomerFidPurchasedParams{

		timeout: timeout,
	}
}

// NewPutCustomersCustomerFidPurchasedParamsWithContext creates a new PutCustomersCustomerFidPurchasedParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutCustomersCustomerFidPurchasedParamsWithContext(ctx context.Context) *PutCustomersCustomerFidPurchasedParams {
	var ()
	return &PutCustomersCustomerFidPurchasedParams{

		Context: ctx,
	}
}

// NewPutCustomersCustomerFidPurchasedParamsWithHTTPClient creates a new PutCustomersCustomerFidPurchasedParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutCustomersCustomerFidPurchasedParamsWithHTTPClient(client *http.Client) *PutCustomersCustomerFidPurchasedParams {
	var ()
	return &PutCustomersCustomerFidPurchasedParams{
		HTTPClient: client,
	}
}

/*PutCustomersCustomerFidPurchasedParams contains all the parameters to send to the API endpoint
for the put customers customer fid purchased operation typically these are written to a http.Request
*/
type PutCustomersCustomerFidPurchasedParams struct {

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

// WithTimeout adds the timeout to the put customers customer fid purchased params
func (o *PutCustomersCustomerFidPurchasedParams) WithTimeout(timeout time.Duration) *PutCustomersCustomerFidPurchasedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put customers customer fid purchased params
func (o *PutCustomersCustomerFidPurchasedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put customers customer fid purchased params
func (o *PutCustomersCustomerFidPurchasedParams) WithContext(ctx context.Context) *PutCustomersCustomerFidPurchasedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put customers customer fid purchased params
func (o *PutCustomersCustomerFidPurchasedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put customers customer fid purchased params
func (o *PutCustomersCustomerFidPurchasedParams) WithHTTPClient(client *http.Client) *PutCustomersCustomerFidPurchasedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put customers customer fid purchased params
func (o *PutCustomersCustomerFidPurchasedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerFid adds the customerFid to the put customers customer fid purchased params
func (o *PutCustomersCustomerFidPurchasedParams) WithCustomerFid(customerFid string) *PutCustomersCustomerFidPurchasedParams {
	o.SetCustomerFid(customerFid)
	return o
}

// SetCustomerFid adds the customerFid to the put customers customer fid purchased params
func (o *PutCustomersCustomerFidPurchasedParams) SetCustomerFid(customerFid string) {
	o.CustomerFid = customerFid
}

// WithIsoTime adds the isoTime to the put customers customer fid purchased params
func (o *PutCustomersCustomerFidPurchasedParams) WithIsoTime(isoTime *strfmt.DateTime) *PutCustomersCustomerFidPurchasedParams {
	o.SetIsoTime(isoTime)
	return o
}

// SetIsoTime adds the isoTime to the put customers customer fid purchased params
func (o *PutCustomersCustomerFidPurchasedParams) SetIsoTime(isoTime *strfmt.DateTime) {
	o.IsoTime = isoTime
}

// WriteToRequest writes these params to a swagger request
func (o *PutCustomersCustomerFidPurchasedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
