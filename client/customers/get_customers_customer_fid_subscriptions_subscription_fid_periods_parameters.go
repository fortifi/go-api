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

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsParams creates a new GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsParams object
// with the default values initialized.
func NewGetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsParams() *GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsParams {
	var ()
	return &GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsParamsWithTimeout creates a new GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsParamsWithTimeout(timeout time.Duration) *GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsParams {
	var ()
	return &GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsParams{

		timeout: timeout,
	}
}

// NewGetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsParamsWithContext creates a new GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsParamsWithContext(ctx context.Context) *GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsParams {
	var ()
	return &GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsParams{

		Context: ctx,
	}
}

// NewGetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsParamsWithHTTPClient creates a new GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsParamsWithHTTPClient(client *http.Client) *GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsParams {
	var ()
	return &GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsParams{
		HTTPClient: client,
	}
}

/*GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsParams contains all the parameters to send to the API endpoint
for the get customers customer fid subscriptions subscription fid periods operation typically these are written to a http.Request
*/
type GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsParams struct {

	/*CustomerFid
	  Customer FID to use

	*/
	CustomerFid string
	/*SubscriptionFid
	  Subscription FID to use

	*/
	SubscriptionFid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get customers customer fid subscriptions subscription fid periods params
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsParams) WithTimeout(timeout time.Duration) *GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get customers customer fid subscriptions subscription fid periods params
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get customers customer fid subscriptions subscription fid periods params
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsParams) WithContext(ctx context.Context) *GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get customers customer fid subscriptions subscription fid periods params
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get customers customer fid subscriptions subscription fid periods params
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsParams) WithHTTPClient(client *http.Client) *GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get customers customer fid subscriptions subscription fid periods params
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerFid adds the customerFid to the get customers customer fid subscriptions subscription fid periods params
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsParams) WithCustomerFid(customerFid string) *GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsParams {
	o.SetCustomerFid(customerFid)
	return o
}

// SetCustomerFid adds the customerFid to the get customers customer fid subscriptions subscription fid periods params
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsParams) SetCustomerFid(customerFid string) {
	o.CustomerFid = customerFid
}

// WithSubscriptionFid adds the subscriptionFid to the get customers customer fid subscriptions subscription fid periods params
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsParams) WithSubscriptionFid(subscriptionFid string) *GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsParams {
	o.SetSubscriptionFid(subscriptionFid)
	return o
}

// SetSubscriptionFid adds the subscriptionFid to the get customers customer fid subscriptions subscription fid periods params
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsParams) SetSubscriptionFid(subscriptionFid string) {
	o.SubscriptionFid = subscriptionFid
}

// WriteToRequest writes these params to a swagger request
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param customerFid
	if err := r.SetPathParam("customerFid", o.CustomerFid); err != nil {
		return err
	}

	// path param subscriptionFid
	if err := r.SetPathParam("subscriptionFid", o.SubscriptionFid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}