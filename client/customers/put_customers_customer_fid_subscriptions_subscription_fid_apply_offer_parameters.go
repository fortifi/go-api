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

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidApplyOfferParams creates a new PutCustomersCustomerFidSubscriptionsSubscriptionFidApplyOfferParams object
// with the default values initialized.
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidApplyOfferParams() *PutCustomersCustomerFidSubscriptionsSubscriptionFidApplyOfferParams {
	var ()
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidApplyOfferParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidApplyOfferParamsWithTimeout creates a new PutCustomersCustomerFidSubscriptionsSubscriptionFidApplyOfferParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidApplyOfferParamsWithTimeout(timeout time.Duration) *PutCustomersCustomerFidSubscriptionsSubscriptionFidApplyOfferParams {
	var ()
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidApplyOfferParams{

		timeout: timeout,
	}
}

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidApplyOfferParamsWithContext creates a new PutCustomersCustomerFidSubscriptionsSubscriptionFidApplyOfferParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidApplyOfferParamsWithContext(ctx context.Context) *PutCustomersCustomerFidSubscriptionsSubscriptionFidApplyOfferParams {
	var ()
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidApplyOfferParams{

		Context: ctx,
	}
}

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidApplyOfferParamsWithHTTPClient creates a new PutCustomersCustomerFidSubscriptionsSubscriptionFidApplyOfferParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidApplyOfferParamsWithHTTPClient(client *http.Client) *PutCustomersCustomerFidSubscriptionsSubscriptionFidApplyOfferParams {
	var ()
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidApplyOfferParams{
		HTTPClient: client,
	}
}

/*PutCustomersCustomerFidSubscriptionsSubscriptionFidApplyOfferParams contains all the parameters to send to the API endpoint
for the put customers customer fid subscriptions subscription fid apply offer operation typically these are written to a http.Request
*/
type PutCustomersCustomerFidSubscriptionsSubscriptionFidApplyOfferParams struct {

	/*CustomerFid
	  Customer FID to use

	*/
	CustomerFid string
	/*OfferFid*/
	OfferFid string
	/*SubscriptionFid
	  Subscription FID to use

	*/
	SubscriptionFid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put customers customer fid subscriptions subscription fid apply offer params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidApplyOfferParams) WithTimeout(timeout time.Duration) *PutCustomersCustomerFidSubscriptionsSubscriptionFidApplyOfferParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put customers customer fid subscriptions subscription fid apply offer params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidApplyOfferParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put customers customer fid subscriptions subscription fid apply offer params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidApplyOfferParams) WithContext(ctx context.Context) *PutCustomersCustomerFidSubscriptionsSubscriptionFidApplyOfferParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put customers customer fid subscriptions subscription fid apply offer params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidApplyOfferParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put customers customer fid subscriptions subscription fid apply offer params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidApplyOfferParams) WithHTTPClient(client *http.Client) *PutCustomersCustomerFidSubscriptionsSubscriptionFidApplyOfferParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put customers customer fid subscriptions subscription fid apply offer params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidApplyOfferParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerFid adds the customerFid to the put customers customer fid subscriptions subscription fid apply offer params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidApplyOfferParams) WithCustomerFid(customerFid string) *PutCustomersCustomerFidSubscriptionsSubscriptionFidApplyOfferParams {
	o.SetCustomerFid(customerFid)
	return o
}

// SetCustomerFid adds the customerFid to the put customers customer fid subscriptions subscription fid apply offer params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidApplyOfferParams) SetCustomerFid(customerFid string) {
	o.CustomerFid = customerFid
}

// WithOfferFid adds the offerFid to the put customers customer fid subscriptions subscription fid apply offer params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidApplyOfferParams) WithOfferFid(offerFid string) *PutCustomersCustomerFidSubscriptionsSubscriptionFidApplyOfferParams {
	o.SetOfferFid(offerFid)
	return o
}

// SetOfferFid adds the offerFid to the put customers customer fid subscriptions subscription fid apply offer params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidApplyOfferParams) SetOfferFid(offerFid string) {
	o.OfferFid = offerFid
}

// WithSubscriptionFid adds the subscriptionFid to the put customers customer fid subscriptions subscription fid apply offer params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidApplyOfferParams) WithSubscriptionFid(subscriptionFid string) *PutCustomersCustomerFidSubscriptionsSubscriptionFidApplyOfferParams {
	o.SetSubscriptionFid(subscriptionFid)
	return o
}

// SetSubscriptionFid adds the subscriptionFid to the put customers customer fid subscriptions subscription fid apply offer params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidApplyOfferParams) SetSubscriptionFid(subscriptionFid string) {
	o.SubscriptionFid = subscriptionFid
}

// WriteToRequest writes these params to a swagger request
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidApplyOfferParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param customerFid
	if err := r.SetPathParam("customerFid", o.CustomerFid); err != nil {
		return err
	}

	// form param offerFid
	frOfferFid := o.OfferFid
	fOfferFid := frOfferFid
	if fOfferFid != "" {
		if err := r.SetFormParam("offerFid", fOfferFid); err != nil {
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