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

// NewGetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchParams creates a new GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchParams() *GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchParams {
	return &GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchParamsWithTimeout creates a new GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchParams object
// with the ability to set a timeout on a request.
func NewGetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchParamsWithTimeout(timeout time.Duration) *GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchParams {
	return &GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchParams{
		timeout: timeout,
	}
}

// NewGetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchParamsWithContext creates a new GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchParams object
// with the ability to set a context for a request.
func NewGetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchParamsWithContext(ctx context.Context) *GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchParams {
	return &GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchParams{
		Context: ctx,
	}
}

// NewGetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchParamsWithHTTPClient creates a new GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchParamsWithHTTPClient(client *http.Client) *GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchParams {
	return &GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchParams{
		HTTPClient: client,
	}
}

/*
GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchParams contains all the parameters to send to the API endpoint

	for the get customers customer fid subscriptions subscription fid retention flow flow search operation.

	Typically these are written to a http.Request.
*/
type GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchParams struct {

	/* CustomerFid.

	   Customer FID to use
	*/
	CustomerFid string

	/* FlowSearch.

	   The unique code or fid for the flow you wish to retrieve
	*/
	FlowSearch string

	/* SubscriptionFid.

	   Subscription FID to use
	*/
	SubscriptionFid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get customers customer fid subscriptions subscription fid retention flow flow search params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchParams) WithDefaults() *GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get customers customer fid subscriptions subscription fid retention flow flow search params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get customers customer fid subscriptions subscription fid retention flow flow search params
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchParams) WithTimeout(timeout time.Duration) *GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get customers customer fid subscriptions subscription fid retention flow flow search params
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get customers customer fid subscriptions subscription fid retention flow flow search params
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchParams) WithContext(ctx context.Context) *GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get customers customer fid subscriptions subscription fid retention flow flow search params
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get customers customer fid subscriptions subscription fid retention flow flow search params
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchParams) WithHTTPClient(client *http.Client) *GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get customers customer fid subscriptions subscription fid retention flow flow search params
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerFid adds the customerFid to the get customers customer fid subscriptions subscription fid retention flow flow search params
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchParams) WithCustomerFid(customerFid string) *GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchParams {
	o.SetCustomerFid(customerFid)
	return o
}

// SetCustomerFid adds the customerFid to the get customers customer fid subscriptions subscription fid retention flow flow search params
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchParams) SetCustomerFid(customerFid string) {
	o.CustomerFid = customerFid
}

// WithFlowSearch adds the flowSearch to the get customers customer fid subscriptions subscription fid retention flow flow search params
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchParams) WithFlowSearch(flowSearch string) *GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchParams {
	o.SetFlowSearch(flowSearch)
	return o
}

// SetFlowSearch adds the flowSearch to the get customers customer fid subscriptions subscription fid retention flow flow search params
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchParams) SetFlowSearch(flowSearch string) {
	o.FlowSearch = flowSearch
}

// WithSubscriptionFid adds the subscriptionFid to the get customers customer fid subscriptions subscription fid retention flow flow search params
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchParams) WithSubscriptionFid(subscriptionFid string) *GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchParams {
	o.SetSubscriptionFid(subscriptionFid)
	return o
}

// SetSubscriptionFid adds the subscriptionFid to the get customers customer fid subscriptions subscription fid retention flow flow search params
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchParams) SetSubscriptionFid(subscriptionFid string) {
	o.SubscriptionFid = subscriptionFid
}

// WriteToRequest writes these params to a swagger request
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param subscriptionFid
	if err := r.SetPathParam("subscriptionFid", o.SubscriptionFid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
