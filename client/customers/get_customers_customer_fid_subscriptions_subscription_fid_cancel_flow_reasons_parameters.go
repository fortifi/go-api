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

// NewGetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowReasonsParams creates a new GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowReasonsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowReasonsParams() *GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowReasonsParams {
	return &GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowReasonsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowReasonsParamsWithTimeout creates a new GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowReasonsParams object
// with the ability to set a timeout on a request.
func NewGetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowReasonsParamsWithTimeout(timeout time.Duration) *GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowReasonsParams {
	return &GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowReasonsParams{
		timeout: timeout,
	}
}

// NewGetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowReasonsParamsWithContext creates a new GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowReasonsParams object
// with the ability to set a context for a request.
func NewGetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowReasonsParamsWithContext(ctx context.Context) *GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowReasonsParams {
	return &GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowReasonsParams{
		Context: ctx,
	}
}

// NewGetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowReasonsParamsWithHTTPClient creates a new GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowReasonsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowReasonsParamsWithHTTPClient(client *http.Client) *GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowReasonsParams {
	return &GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowReasonsParams{
		HTTPClient: client,
	}
}

/*
GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowReasonsParams contains all the parameters to send to the API endpoint

	for the get customers customer fid subscriptions subscription fid cancel flow reasons operation.

	Typically these are written to a http.Request.
*/
type GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowReasonsParams struct {

	/* CustomerFid.

	   Customer FID to use
	*/
	CustomerFid string

	/* SubscriptionFid.

	   Subscription FID to use
	*/
	SubscriptionFid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get customers customer fid subscriptions subscription fid cancel flow reasons params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowReasonsParams) WithDefaults() *GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowReasonsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get customers customer fid subscriptions subscription fid cancel flow reasons params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowReasonsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get customers customer fid subscriptions subscription fid cancel flow reasons params
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowReasonsParams) WithTimeout(timeout time.Duration) *GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowReasonsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get customers customer fid subscriptions subscription fid cancel flow reasons params
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowReasonsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get customers customer fid subscriptions subscription fid cancel flow reasons params
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowReasonsParams) WithContext(ctx context.Context) *GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowReasonsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get customers customer fid subscriptions subscription fid cancel flow reasons params
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowReasonsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get customers customer fid subscriptions subscription fid cancel flow reasons params
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowReasonsParams) WithHTTPClient(client *http.Client) *GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowReasonsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get customers customer fid subscriptions subscription fid cancel flow reasons params
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowReasonsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerFid adds the customerFid to the get customers customer fid subscriptions subscription fid cancel flow reasons params
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowReasonsParams) WithCustomerFid(customerFid string) *GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowReasonsParams {
	o.SetCustomerFid(customerFid)
	return o
}

// SetCustomerFid adds the customerFid to the get customers customer fid subscriptions subscription fid cancel flow reasons params
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowReasonsParams) SetCustomerFid(customerFid string) {
	o.CustomerFid = customerFid
}

// WithSubscriptionFid adds the subscriptionFid to the get customers customer fid subscriptions subscription fid cancel flow reasons params
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowReasonsParams) WithSubscriptionFid(subscriptionFid string) *GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowReasonsParams {
	o.SetSubscriptionFid(subscriptionFid)
	return o
}

// SetSubscriptionFid adds the subscriptionFid to the get customers customer fid subscriptions subscription fid cancel flow reasons params
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowReasonsParams) SetSubscriptionFid(subscriptionFid string) {
	o.SubscriptionFid = subscriptionFid
}

// WriteToRequest writes these params to a swagger request
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowReasonsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
