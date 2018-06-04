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

// NewGetCustomersCustomerFidPaymentsPaymentFidParams creates a new GetCustomersCustomerFidPaymentsPaymentFidParams object
// with the default values initialized.
func NewGetCustomersCustomerFidPaymentsPaymentFidParams() *GetCustomersCustomerFidPaymentsPaymentFidParams {
	var ()
	return &GetCustomersCustomerFidPaymentsPaymentFidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCustomersCustomerFidPaymentsPaymentFidParamsWithTimeout creates a new GetCustomersCustomerFidPaymentsPaymentFidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCustomersCustomerFidPaymentsPaymentFidParamsWithTimeout(timeout time.Duration) *GetCustomersCustomerFidPaymentsPaymentFidParams {
	var ()
	return &GetCustomersCustomerFidPaymentsPaymentFidParams{

		timeout: timeout,
	}
}

// NewGetCustomersCustomerFidPaymentsPaymentFidParamsWithContext creates a new GetCustomersCustomerFidPaymentsPaymentFidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCustomersCustomerFidPaymentsPaymentFidParamsWithContext(ctx context.Context) *GetCustomersCustomerFidPaymentsPaymentFidParams {
	var ()
	return &GetCustomersCustomerFidPaymentsPaymentFidParams{

		Context: ctx,
	}
}

// NewGetCustomersCustomerFidPaymentsPaymentFidParamsWithHTTPClient creates a new GetCustomersCustomerFidPaymentsPaymentFidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCustomersCustomerFidPaymentsPaymentFidParamsWithHTTPClient(client *http.Client) *GetCustomersCustomerFidPaymentsPaymentFidParams {
	var ()
	return &GetCustomersCustomerFidPaymentsPaymentFidParams{
		HTTPClient: client,
	}
}

/*GetCustomersCustomerFidPaymentsPaymentFidParams contains all the parameters to send to the API endpoint
for the get customers customer fid payments payment fid operation typically these are written to a http.Request
*/
type GetCustomersCustomerFidPaymentsPaymentFidParams struct {

	/*CustomerFid
	  Customer FID to use

	*/
	CustomerFid string
	/*PaymentFid
	  Payment FID to use

	*/
	PaymentFid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get customers customer fid payments payment fid params
func (o *GetCustomersCustomerFidPaymentsPaymentFidParams) WithTimeout(timeout time.Duration) *GetCustomersCustomerFidPaymentsPaymentFidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get customers customer fid payments payment fid params
func (o *GetCustomersCustomerFidPaymentsPaymentFidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get customers customer fid payments payment fid params
func (o *GetCustomersCustomerFidPaymentsPaymentFidParams) WithContext(ctx context.Context) *GetCustomersCustomerFidPaymentsPaymentFidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get customers customer fid payments payment fid params
func (o *GetCustomersCustomerFidPaymentsPaymentFidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get customers customer fid payments payment fid params
func (o *GetCustomersCustomerFidPaymentsPaymentFidParams) WithHTTPClient(client *http.Client) *GetCustomersCustomerFidPaymentsPaymentFidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get customers customer fid payments payment fid params
func (o *GetCustomersCustomerFidPaymentsPaymentFidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerFid adds the customerFid to the get customers customer fid payments payment fid params
func (o *GetCustomersCustomerFidPaymentsPaymentFidParams) WithCustomerFid(customerFid string) *GetCustomersCustomerFidPaymentsPaymentFidParams {
	o.SetCustomerFid(customerFid)
	return o
}

// SetCustomerFid adds the customerFid to the get customers customer fid payments payment fid params
func (o *GetCustomersCustomerFidPaymentsPaymentFidParams) SetCustomerFid(customerFid string) {
	o.CustomerFid = customerFid
}

// WithPaymentFid adds the paymentFid to the get customers customer fid payments payment fid params
func (o *GetCustomersCustomerFidPaymentsPaymentFidParams) WithPaymentFid(paymentFid string) *GetCustomersCustomerFidPaymentsPaymentFidParams {
	o.SetPaymentFid(paymentFid)
	return o
}

// SetPaymentFid adds the paymentFid to the get customers customer fid payments payment fid params
func (o *GetCustomersCustomerFidPaymentsPaymentFidParams) SetPaymentFid(paymentFid string) {
	o.PaymentFid = paymentFid
}

// WriteToRequest writes these params to a swagger request
func (o *GetCustomersCustomerFidPaymentsPaymentFidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param customerFid
	if err := r.SetPathParam("customerFid", o.CustomerFid); err != nil {
		return err
	}

	// path param paymentFid
	if err := r.SetPathParam("paymentFid", o.PaymentFid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
