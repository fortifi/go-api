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

// NewGetCustomersCustomerFidSubscriptionsSubscriptionFidParams creates a new GetCustomersCustomerFidSubscriptionsSubscriptionFidParams object
// with the default values initialized.
func NewGetCustomersCustomerFidSubscriptionsSubscriptionFidParams() *GetCustomersCustomerFidSubscriptionsSubscriptionFidParams {
	var ()
	return &GetCustomersCustomerFidSubscriptionsSubscriptionFidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCustomersCustomerFidSubscriptionsSubscriptionFidParamsWithTimeout creates a new GetCustomersCustomerFidSubscriptionsSubscriptionFidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCustomersCustomerFidSubscriptionsSubscriptionFidParamsWithTimeout(timeout time.Duration) *GetCustomersCustomerFidSubscriptionsSubscriptionFidParams {
	var ()
	return &GetCustomersCustomerFidSubscriptionsSubscriptionFidParams{

		timeout: timeout,
	}
}

// NewGetCustomersCustomerFidSubscriptionsSubscriptionFidParamsWithContext creates a new GetCustomersCustomerFidSubscriptionsSubscriptionFidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCustomersCustomerFidSubscriptionsSubscriptionFidParamsWithContext(ctx context.Context) *GetCustomersCustomerFidSubscriptionsSubscriptionFidParams {
	var ()
	return &GetCustomersCustomerFidSubscriptionsSubscriptionFidParams{

		Context: ctx,
	}
}

// NewGetCustomersCustomerFidSubscriptionsSubscriptionFidParamsWithHTTPClient creates a new GetCustomersCustomerFidSubscriptionsSubscriptionFidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCustomersCustomerFidSubscriptionsSubscriptionFidParamsWithHTTPClient(client *http.Client) *GetCustomersCustomerFidSubscriptionsSubscriptionFidParams {
	var ()
	return &GetCustomersCustomerFidSubscriptionsSubscriptionFidParams{
		HTTPClient: client,
	}
}

/*GetCustomersCustomerFidSubscriptionsSubscriptionFidParams contains all the parameters to send to the API endpoint
for the get customers customer fid subscriptions subscription fid operation typically these are written to a http.Request
*/
type GetCustomersCustomerFidSubscriptionsSubscriptionFidParams struct {

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

// WithTimeout adds the timeout to the get customers customer fid subscriptions subscription fid params
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidParams) WithTimeout(timeout time.Duration) *GetCustomersCustomerFidSubscriptionsSubscriptionFidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get customers customer fid subscriptions subscription fid params
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get customers customer fid subscriptions subscription fid params
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidParams) WithContext(ctx context.Context) *GetCustomersCustomerFidSubscriptionsSubscriptionFidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get customers customer fid subscriptions subscription fid params
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get customers customer fid subscriptions subscription fid params
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidParams) WithHTTPClient(client *http.Client) *GetCustomersCustomerFidSubscriptionsSubscriptionFidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get customers customer fid subscriptions subscription fid params
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerFid adds the customerFid to the get customers customer fid subscriptions subscription fid params
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidParams) WithCustomerFid(customerFid string) *GetCustomersCustomerFidSubscriptionsSubscriptionFidParams {
	o.SetCustomerFid(customerFid)
	return o
}

// SetCustomerFid adds the customerFid to the get customers customer fid subscriptions subscription fid params
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidParams) SetCustomerFid(customerFid string) {
	o.CustomerFid = customerFid
}

// WithSubscriptionFid adds the subscriptionFid to the get customers customer fid subscriptions subscription fid params
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidParams) WithSubscriptionFid(subscriptionFid string) *GetCustomersCustomerFidSubscriptionsSubscriptionFidParams {
	o.SetSubscriptionFid(subscriptionFid)
	return o
}

// SetSubscriptionFid adds the subscriptionFid to the get customers customer fid subscriptions subscription fid params
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidParams) SetSubscriptionFid(subscriptionFid string) {
	o.SubscriptionFid = subscriptionFid
}

// WriteToRequest writes these params to a swagger request
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
