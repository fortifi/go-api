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

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchParams creates a new PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchParams() *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchParams {
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchParamsWithTimeout creates a new PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchParams object
// with the ability to set a timeout on a request.
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchParamsWithTimeout(timeout time.Duration) *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchParams {
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchParams{
		timeout: timeout,
	}
}

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchParamsWithContext creates a new PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchParams object
// with the ability to set a context for a request.
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchParamsWithContext(ctx context.Context) *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchParams {
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchParams{
		Context: ctx,
	}
}

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchParamsWithHTTPClient creates a new PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchParamsWithHTTPClient(client *http.Client) *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchParams {
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchParams{
		HTTPClient: client,
	}
}

/*
PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchParams contains all the parameters to send to the API endpoint

	for the put customers customer fid subscriptions subscription fid cancel flow flow search operation.

	Typically these are written to a http.Request.
*/
type PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchParams struct {

	/* CustomerFid.

	   Customer FID to use
	*/
	CustomerFid string

	/* FlowSearch.

	   The unique code or fid for the flow you wish to retrieve
	*/
	FlowSearch string

	// Payload.
	Payload *models.ActionCancelFlowPayload

	/* SubscriptionFid.

	   Subscription FID to use
	*/
	SubscriptionFid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put customers customer fid subscriptions subscription fid cancel flow flow search params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchParams) WithDefaults() *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put customers customer fid subscriptions subscription fid cancel flow flow search params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put customers customer fid subscriptions subscription fid cancel flow flow search params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchParams) WithTimeout(timeout time.Duration) *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put customers customer fid subscriptions subscription fid cancel flow flow search params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put customers customer fid subscriptions subscription fid cancel flow flow search params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchParams) WithContext(ctx context.Context) *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put customers customer fid subscriptions subscription fid cancel flow flow search params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put customers customer fid subscriptions subscription fid cancel flow flow search params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchParams) WithHTTPClient(client *http.Client) *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put customers customer fid subscriptions subscription fid cancel flow flow search params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerFid adds the customerFid to the put customers customer fid subscriptions subscription fid cancel flow flow search params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchParams) WithCustomerFid(customerFid string) *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchParams {
	o.SetCustomerFid(customerFid)
	return o
}

// SetCustomerFid adds the customerFid to the put customers customer fid subscriptions subscription fid cancel flow flow search params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchParams) SetCustomerFid(customerFid string) {
	o.CustomerFid = customerFid
}

// WithFlowSearch adds the flowSearch to the put customers customer fid subscriptions subscription fid cancel flow flow search params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchParams) WithFlowSearch(flowSearch string) *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchParams {
	o.SetFlowSearch(flowSearch)
	return o
}

// SetFlowSearch adds the flowSearch to the put customers customer fid subscriptions subscription fid cancel flow flow search params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchParams) SetFlowSearch(flowSearch string) {
	o.FlowSearch = flowSearch
}

// WithPayload adds the payload to the put customers customer fid subscriptions subscription fid cancel flow flow search params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchParams) WithPayload(payload *models.ActionCancelFlowPayload) *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchParams {
	o.SetPayload(payload)
	return o
}

// SetPayload adds the payload to the put customers customer fid subscriptions subscription fid cancel flow flow search params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchParams) SetPayload(payload *models.ActionCancelFlowPayload) {
	o.Payload = payload
}

// WithSubscriptionFid adds the subscriptionFid to the put customers customer fid subscriptions subscription fid cancel flow flow search params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchParams) WithSubscriptionFid(subscriptionFid string) *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchParams {
	o.SetSubscriptionFid(subscriptionFid)
	return o
}

// SetSubscriptionFid adds the subscriptionFid to the put customers customer fid subscriptions subscription fid cancel flow flow search params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchParams) SetSubscriptionFid(subscriptionFid string) {
	o.SubscriptionFid = subscriptionFid
}

// WriteToRequest writes these params to a swagger request
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param customerFid
	if err := r.SetPathParam("customerFid", o.CustomerFid); err != nil {
		return err
	}

	// path param flowSearch
	if err := r.SetPathParam("flowSearch", o.FlowSearch); err != nil {
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
