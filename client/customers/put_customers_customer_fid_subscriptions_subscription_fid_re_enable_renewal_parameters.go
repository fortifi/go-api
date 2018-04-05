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

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewalParams creates a new PutCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewalParams object
// with the default values initialized.
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewalParams() *PutCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewalParams {
	var ()
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewalParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewalParamsWithTimeout creates a new PutCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewalParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewalParamsWithTimeout(timeout time.Duration) *PutCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewalParams {
	var ()
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewalParams{

		timeout: timeout,
	}
}

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewalParamsWithContext creates a new PutCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewalParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewalParamsWithContext(ctx context.Context) *PutCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewalParams {
	var ()
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewalParams{

		Context: ctx,
	}
}

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewalParamsWithHTTPClient creates a new PutCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewalParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewalParamsWithHTTPClient(client *http.Client) *PutCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewalParams {
	var ()
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewalParams{
		HTTPClient: client,
	}
}

/*PutCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewalParams contains all the parameters to send to the API endpoint
for the put customers customer fid subscriptions subscription fid re enable renewal operation typically these are written to a http.Request
*/
type PutCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewalParams struct {

	/*CustomerFid
	  Customer FID to use

	*/
	CustomerFid string
	/*Reason*/
	Reason string
	/*SubscriptionFid
	  Subscription FID to use

	*/
	SubscriptionFid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put customers customer fid subscriptions subscription fid re enable renewal params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewalParams) WithTimeout(timeout time.Duration) *PutCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewalParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put customers customer fid subscriptions subscription fid re enable renewal params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewalParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put customers customer fid subscriptions subscription fid re enable renewal params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewalParams) WithContext(ctx context.Context) *PutCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewalParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put customers customer fid subscriptions subscription fid re enable renewal params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewalParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put customers customer fid subscriptions subscription fid re enable renewal params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewalParams) WithHTTPClient(client *http.Client) *PutCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewalParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put customers customer fid subscriptions subscription fid re enable renewal params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewalParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerFid adds the customerFid to the put customers customer fid subscriptions subscription fid re enable renewal params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewalParams) WithCustomerFid(customerFid string) *PutCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewalParams {
	o.SetCustomerFid(customerFid)
	return o
}

// SetCustomerFid adds the customerFid to the put customers customer fid subscriptions subscription fid re enable renewal params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewalParams) SetCustomerFid(customerFid string) {
	o.CustomerFid = customerFid
}

// WithReason adds the reason to the put customers customer fid subscriptions subscription fid re enable renewal params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewalParams) WithReason(reason string) *PutCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewalParams {
	o.SetReason(reason)
	return o
}

// SetReason adds the reason to the put customers customer fid subscriptions subscription fid re enable renewal params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewalParams) SetReason(reason string) {
	o.Reason = reason
}

// WithSubscriptionFid adds the subscriptionFid to the put customers customer fid subscriptions subscription fid re enable renewal params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewalParams) WithSubscriptionFid(subscriptionFid string) *PutCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewalParams {
	o.SetSubscriptionFid(subscriptionFid)
	return o
}

// SetSubscriptionFid adds the subscriptionFid to the put customers customer fid subscriptions subscription fid re enable renewal params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewalParams) SetSubscriptionFid(subscriptionFid string) {
	o.SubscriptionFid = subscriptionFid
}

// WriteToRequest writes these params to a swagger request
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewalParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param customerFid
	if err := r.SetPathParam("customerFid", o.CustomerFid); err != nil {
		return err
	}

	// form param reason
	frReason := o.Reason
	fReason := frReason
	if fReason != "" {
		if err := r.SetFormParam("reason", fReason); err != nil {
			return err
		}
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