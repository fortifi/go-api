// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPostCustomersCustomerFidPaymentMethodsPaypalInitialiseParams creates a new PostCustomersCustomerFidPaymentMethodsPaypalInitialiseParams object
// with the default values initialized.
func NewPostCustomersCustomerFidPaymentMethodsPaypalInitialiseParams() *PostCustomersCustomerFidPaymentMethodsPaypalInitialiseParams {
	var ()
	return &PostCustomersCustomerFidPaymentMethodsPaypalInitialiseParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostCustomersCustomerFidPaymentMethodsPaypalInitialiseParamsWithTimeout creates a new PostCustomersCustomerFidPaymentMethodsPaypalInitialiseParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostCustomersCustomerFidPaymentMethodsPaypalInitialiseParamsWithTimeout(timeout time.Duration) *PostCustomersCustomerFidPaymentMethodsPaypalInitialiseParams {
	var ()
	return &PostCustomersCustomerFidPaymentMethodsPaypalInitialiseParams{

		timeout: timeout,
	}
}

// NewPostCustomersCustomerFidPaymentMethodsPaypalInitialiseParamsWithContext creates a new PostCustomersCustomerFidPaymentMethodsPaypalInitialiseParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostCustomersCustomerFidPaymentMethodsPaypalInitialiseParamsWithContext(ctx context.Context) *PostCustomersCustomerFidPaymentMethodsPaypalInitialiseParams {
	var ()
	return &PostCustomersCustomerFidPaymentMethodsPaypalInitialiseParams{

		Context: ctx,
	}
}

// NewPostCustomersCustomerFidPaymentMethodsPaypalInitialiseParamsWithHTTPClient creates a new PostCustomersCustomerFidPaymentMethodsPaypalInitialiseParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostCustomersCustomerFidPaymentMethodsPaypalInitialiseParamsWithHTTPClient(client *http.Client) *PostCustomersCustomerFidPaymentMethodsPaypalInitialiseParams {
	var ()
	return &PostCustomersCustomerFidPaymentMethodsPaypalInitialiseParams{
		HTTPClient: client,
	}
}

/*PostCustomersCustomerFidPaymentMethodsPaypalInitialiseParams contains all the parameters to send to the API endpoint
for the post customers customer fid payment methods paypal initialise operation typically these are written to a http.Request
*/
type PostCustomersCustomerFidPaymentMethodsPaypalInitialiseParams struct {

	/*CancelURL*/
	CancelURL string
	/*CustomerFid
	  Customer FID to use

	*/
	CustomerFid string
	/*FailURL*/
	FailURL string
	/*PaymentServiceFid
	  Payment Service FID to use

	*/
	PaymentServiceFid string
	/*SubscriptionFids
	  Subscription FIDs

	*/
	SubscriptionFids []string
	/*SuccessURL*/
	SuccessURL string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post customers customer fid payment methods paypal initialise params
func (o *PostCustomersCustomerFidPaymentMethodsPaypalInitialiseParams) WithTimeout(timeout time.Duration) *PostCustomersCustomerFidPaymentMethodsPaypalInitialiseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post customers customer fid payment methods paypal initialise params
func (o *PostCustomersCustomerFidPaymentMethodsPaypalInitialiseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post customers customer fid payment methods paypal initialise params
func (o *PostCustomersCustomerFidPaymentMethodsPaypalInitialiseParams) WithContext(ctx context.Context) *PostCustomersCustomerFidPaymentMethodsPaypalInitialiseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post customers customer fid payment methods paypal initialise params
func (o *PostCustomersCustomerFidPaymentMethodsPaypalInitialiseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post customers customer fid payment methods paypal initialise params
func (o *PostCustomersCustomerFidPaymentMethodsPaypalInitialiseParams) WithHTTPClient(client *http.Client) *PostCustomersCustomerFidPaymentMethodsPaypalInitialiseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post customers customer fid payment methods paypal initialise params
func (o *PostCustomersCustomerFidPaymentMethodsPaypalInitialiseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCancelURL adds the cancelURL to the post customers customer fid payment methods paypal initialise params
func (o *PostCustomersCustomerFidPaymentMethodsPaypalInitialiseParams) WithCancelURL(cancelURL string) *PostCustomersCustomerFidPaymentMethodsPaypalInitialiseParams {
	o.SetCancelURL(cancelURL)
	return o
}

// SetCancelURL adds the cancelUrl to the post customers customer fid payment methods paypal initialise params
func (o *PostCustomersCustomerFidPaymentMethodsPaypalInitialiseParams) SetCancelURL(cancelURL string) {
	o.CancelURL = cancelURL
}

// WithCustomerFid adds the customerFid to the post customers customer fid payment methods paypal initialise params
func (o *PostCustomersCustomerFidPaymentMethodsPaypalInitialiseParams) WithCustomerFid(customerFid string) *PostCustomersCustomerFidPaymentMethodsPaypalInitialiseParams {
	o.SetCustomerFid(customerFid)
	return o
}

// SetCustomerFid adds the customerFid to the post customers customer fid payment methods paypal initialise params
func (o *PostCustomersCustomerFidPaymentMethodsPaypalInitialiseParams) SetCustomerFid(customerFid string) {
	o.CustomerFid = customerFid
}

// WithFailURL adds the failURL to the post customers customer fid payment methods paypal initialise params
func (o *PostCustomersCustomerFidPaymentMethodsPaypalInitialiseParams) WithFailURL(failURL string) *PostCustomersCustomerFidPaymentMethodsPaypalInitialiseParams {
	o.SetFailURL(failURL)
	return o
}

// SetFailURL adds the failUrl to the post customers customer fid payment methods paypal initialise params
func (o *PostCustomersCustomerFidPaymentMethodsPaypalInitialiseParams) SetFailURL(failURL string) {
	o.FailURL = failURL
}

// WithPaymentServiceFid adds the paymentServiceFid to the post customers customer fid payment methods paypal initialise params
func (o *PostCustomersCustomerFidPaymentMethodsPaypalInitialiseParams) WithPaymentServiceFid(paymentServiceFid string) *PostCustomersCustomerFidPaymentMethodsPaypalInitialiseParams {
	o.SetPaymentServiceFid(paymentServiceFid)
	return o
}

// SetPaymentServiceFid adds the paymentServiceFid to the post customers customer fid payment methods paypal initialise params
func (o *PostCustomersCustomerFidPaymentMethodsPaypalInitialiseParams) SetPaymentServiceFid(paymentServiceFid string) {
	o.PaymentServiceFid = paymentServiceFid
}

// WithSubscriptionFids adds the subscriptionFids to the post customers customer fid payment methods paypal initialise params
func (o *PostCustomersCustomerFidPaymentMethodsPaypalInitialiseParams) WithSubscriptionFids(subscriptionFids []string) *PostCustomersCustomerFidPaymentMethodsPaypalInitialiseParams {
	o.SetSubscriptionFids(subscriptionFids)
	return o
}

// SetSubscriptionFids adds the subscriptionFids to the post customers customer fid payment methods paypal initialise params
func (o *PostCustomersCustomerFidPaymentMethodsPaypalInitialiseParams) SetSubscriptionFids(subscriptionFids []string) {
	o.SubscriptionFids = subscriptionFids
}

// WithSuccessURL adds the successURL to the post customers customer fid payment methods paypal initialise params
func (o *PostCustomersCustomerFidPaymentMethodsPaypalInitialiseParams) WithSuccessURL(successURL string) *PostCustomersCustomerFidPaymentMethodsPaypalInitialiseParams {
	o.SetSuccessURL(successURL)
	return o
}

// SetSuccessURL adds the successUrl to the post customers customer fid payment methods paypal initialise params
func (o *PostCustomersCustomerFidPaymentMethodsPaypalInitialiseParams) SetSuccessURL(successURL string) {
	o.SuccessURL = successURL
}

// WriteToRequest writes these params to a swagger request
func (o *PostCustomersCustomerFidPaymentMethodsPaypalInitialiseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// form param cancelUrl
	frCancelURL := o.CancelURL
	fCancelURL := frCancelURL
	if fCancelURL != "" {
		if err := r.SetFormParam("cancelUrl", fCancelURL); err != nil {
			return err
		}
	}

	// path param customerFid
	if err := r.SetPathParam("customerFid", o.CustomerFid); err != nil {
		return err
	}

	// form param failUrl
	frFailURL := o.FailURL
	fFailURL := frFailURL
	if fFailURL != "" {
		if err := r.SetFormParam("failUrl", fFailURL); err != nil {
			return err
		}
	}

	// form param paymentServiceFid
	frPaymentServiceFid := o.PaymentServiceFid
	fPaymentServiceFid := frPaymentServiceFid
	if fPaymentServiceFid != "" {
		if err := r.SetFormParam("paymentServiceFid", fPaymentServiceFid); err != nil {
			return err
		}
	}

	valuesSubscriptionFids := o.SubscriptionFids

	joinedSubscriptionFids := swag.JoinByFormat(valuesSubscriptionFids, "")
	// form array param subscriptionFids
	if err := r.SetFormParam("subscriptionFids", joinedSubscriptionFids...); err != nil {
		return err
	}

	// form param successUrl
	frSuccessURL := o.SuccessURL
	fSuccessURL := frSuccessURL
	if fSuccessURL != "" {
		if err := r.SetFormParam("successUrl", fSuccessURL); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
