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

	"github.com/fortifi/go-api/models"
)

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidCancelParams creates a new PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelParams object
// with the default values initialized.
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidCancelParams() *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelParams {
	var ()
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidCancelParamsWithTimeout creates a new PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidCancelParamsWithTimeout(timeout time.Duration) *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelParams {
	var ()
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelParams{

		timeout: timeout,
	}
}

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidCancelParamsWithContext creates a new PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidCancelParamsWithContext(ctx context.Context) *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelParams {
	var ()
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelParams{

		Context: ctx,
	}
}

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidCancelParamsWithHTTPClient creates a new PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidCancelParamsWithHTTPClient(client *http.Client) *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelParams {
	var ()
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelParams{
		HTTPClient: client,
	}
}

/*PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelParams contains all the parameters to send to the API endpoint
for the put customers customer fid subscriptions subscription fid cancel operation typically these are written to a http.Request
*/
type PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelParams struct {

	/*CustomerFid
	  Customer FID to use

	*/
	CustomerFid string
	/*Payload*/
	Payload *models.SubscriptionCancelPayload
	/*SubscriptionFid
	  Subscription FID to use

	*/
	SubscriptionFid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put customers customer fid subscriptions subscription fid cancel params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelParams) WithTimeout(timeout time.Duration) *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put customers customer fid subscriptions subscription fid cancel params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put customers customer fid subscriptions subscription fid cancel params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelParams) WithContext(ctx context.Context) *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put customers customer fid subscriptions subscription fid cancel params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put customers customer fid subscriptions subscription fid cancel params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelParams) WithHTTPClient(client *http.Client) *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put customers customer fid subscriptions subscription fid cancel params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerFid adds the customerFid to the put customers customer fid subscriptions subscription fid cancel params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelParams) WithCustomerFid(customerFid string) *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelParams {
	o.SetCustomerFid(customerFid)
	return o
}

// SetCustomerFid adds the customerFid to the put customers customer fid subscriptions subscription fid cancel params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelParams) SetCustomerFid(customerFid string) {
	o.CustomerFid = customerFid
}

// WithPayload adds the payload to the put customers customer fid subscriptions subscription fid cancel params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelParams) WithPayload(payload *models.SubscriptionCancelPayload) *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelParams {
	o.SetPayload(payload)
	return o
}

// SetPayload adds the payload to the put customers customer fid subscriptions subscription fid cancel params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelParams) SetPayload(payload *models.SubscriptionCancelPayload) {
	o.Payload = payload
}

// WithSubscriptionFid adds the subscriptionFid to the put customers customer fid subscriptions subscription fid cancel params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelParams) WithSubscriptionFid(subscriptionFid string) *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelParams {
	o.SetSubscriptionFid(subscriptionFid)
	return o
}

// SetSubscriptionFid adds the subscriptionFid to the put customers customer fid subscriptions subscription fid cancel params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelParams) SetSubscriptionFid(subscriptionFid string) {
	o.SubscriptionFid = subscriptionFid
}

// WriteToRequest writes these params to a swagger request
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param subscriptionFid
	if err := r.SetPathParam("subscriptionFid", o.SubscriptionFid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}