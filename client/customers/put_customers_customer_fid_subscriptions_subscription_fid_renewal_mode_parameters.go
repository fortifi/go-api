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

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeParams creates a new PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeParams() *PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeParams {
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeParamsWithTimeout creates a new PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeParams object
// with the ability to set a timeout on a request.
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeParamsWithTimeout(timeout time.Duration) *PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeParams {
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeParams{
		timeout: timeout,
	}
}

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeParamsWithContext creates a new PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeParams object
// with the ability to set a context for a request.
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeParamsWithContext(ctx context.Context) *PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeParams {
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeParams{
		Context: ctx,
	}
}

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeParamsWithHTTPClient creates a new PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeParamsWithHTTPClient(client *http.Client) *PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeParams {
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeParams{
		HTTPClient: client,
	}
}

/*
PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeParams contains all the parameters to send to the API endpoint

	for the put customers customer fid subscriptions subscription fid renewal mode operation.

	Typically these are written to a http.Request.
*/
type PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeParams struct {

	/* CustomerFid.

	   Customer FID to use
	*/
	CustomerFid string

	// Note.
	Note *string

	// Reason.
	Reason *string

	// RenewalMode.
	RenewalMode string

	/* SubscriptionFid.

	   Subscription FID to use
	*/
	SubscriptionFid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put customers customer fid subscriptions subscription fid renewal mode params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeParams) WithDefaults() *PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put customers customer fid subscriptions subscription fid renewal mode params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put customers customer fid subscriptions subscription fid renewal mode params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeParams) WithTimeout(timeout time.Duration) *PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put customers customer fid subscriptions subscription fid renewal mode params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put customers customer fid subscriptions subscription fid renewal mode params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeParams) WithContext(ctx context.Context) *PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put customers customer fid subscriptions subscription fid renewal mode params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put customers customer fid subscriptions subscription fid renewal mode params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeParams) WithHTTPClient(client *http.Client) *PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put customers customer fid subscriptions subscription fid renewal mode params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerFid adds the customerFid to the put customers customer fid subscriptions subscription fid renewal mode params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeParams) WithCustomerFid(customerFid string) *PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeParams {
	o.SetCustomerFid(customerFid)
	return o
}

// SetCustomerFid adds the customerFid to the put customers customer fid subscriptions subscription fid renewal mode params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeParams) SetCustomerFid(customerFid string) {
	o.CustomerFid = customerFid
}

// WithNote adds the note to the put customers customer fid subscriptions subscription fid renewal mode params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeParams) WithNote(note *string) *PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeParams {
	o.SetNote(note)
	return o
}

// SetNote adds the note to the put customers customer fid subscriptions subscription fid renewal mode params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeParams) SetNote(note *string) {
	o.Note = note
}

// WithReason adds the reason to the put customers customer fid subscriptions subscription fid renewal mode params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeParams) WithReason(reason *string) *PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeParams {
	o.SetReason(reason)
	return o
}

// SetReason adds the reason to the put customers customer fid subscriptions subscription fid renewal mode params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeParams) SetReason(reason *string) {
	o.Reason = reason
}

// WithRenewalMode adds the renewalMode to the put customers customer fid subscriptions subscription fid renewal mode params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeParams) WithRenewalMode(renewalMode string) *PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeParams {
	o.SetRenewalMode(renewalMode)
	return o
}

// SetRenewalMode adds the renewalMode to the put customers customer fid subscriptions subscription fid renewal mode params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeParams) SetRenewalMode(renewalMode string) {
	o.RenewalMode = renewalMode
}

// WithSubscriptionFid adds the subscriptionFid to the put customers customer fid subscriptions subscription fid renewal mode params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeParams) WithSubscriptionFid(subscriptionFid string) *PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeParams {
	o.SetSubscriptionFid(subscriptionFid)
	return o
}

// SetSubscriptionFid adds the subscriptionFid to the put customers customer fid subscriptions subscription fid renewal mode params
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeParams) SetSubscriptionFid(subscriptionFid string) {
	o.SubscriptionFid = subscriptionFid
}

// WriteToRequest writes these params to a swagger request
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidRenewalModeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param customerFid
	if err := r.SetPathParam("customerFid", o.CustomerFid); err != nil {
		return err
	}

	if o.Note != nil {

		// form param note
		var frNote string
		if o.Note != nil {
			frNote = *o.Note
		}
		fNote := frNote
		if fNote != "" {
			if err := r.SetFormParam("note", fNote); err != nil {
				return err
			}
		}
	}

	if o.Reason != nil {

		// form param reason
		var frReason string
		if o.Reason != nil {
			frReason = *o.Reason
		}
		fReason := frReason
		if fReason != "" {
			if err := r.SetFormParam("reason", fReason); err != nil {
				return err
			}
		}
	}

	// form param renewalMode
	frRenewalMode := o.RenewalMode
	fRenewalMode := frRenewalMode
	if fRenewalMode != "" {
		if err := r.SetFormParam("renewalMode", fRenewalMode); err != nil {
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
