// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewPostCustomersCustomerFidPaymentMethodsCreatePendingParams creates a new PostCustomersCustomerFidPaymentMethodsCreatePendingParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostCustomersCustomerFidPaymentMethodsCreatePendingParams() *PostCustomersCustomerFidPaymentMethodsCreatePendingParams {
	return &PostCustomersCustomerFidPaymentMethodsCreatePendingParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostCustomersCustomerFidPaymentMethodsCreatePendingParamsWithTimeout creates a new PostCustomersCustomerFidPaymentMethodsCreatePendingParams object
// with the ability to set a timeout on a request.
func NewPostCustomersCustomerFidPaymentMethodsCreatePendingParamsWithTimeout(timeout time.Duration) *PostCustomersCustomerFidPaymentMethodsCreatePendingParams {
	return &PostCustomersCustomerFidPaymentMethodsCreatePendingParams{
		timeout: timeout,
	}
}

// NewPostCustomersCustomerFidPaymentMethodsCreatePendingParamsWithContext creates a new PostCustomersCustomerFidPaymentMethodsCreatePendingParams object
// with the ability to set a context for a request.
func NewPostCustomersCustomerFidPaymentMethodsCreatePendingParamsWithContext(ctx context.Context) *PostCustomersCustomerFidPaymentMethodsCreatePendingParams {
	return &PostCustomersCustomerFidPaymentMethodsCreatePendingParams{
		Context: ctx,
	}
}

// NewPostCustomersCustomerFidPaymentMethodsCreatePendingParamsWithHTTPClient creates a new PostCustomersCustomerFidPaymentMethodsCreatePendingParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostCustomersCustomerFidPaymentMethodsCreatePendingParamsWithHTTPClient(client *http.Client) *PostCustomersCustomerFidPaymentMethodsCreatePendingParams {
	return &PostCustomersCustomerFidPaymentMethodsCreatePendingParams{
		HTTPClient: client,
	}
}

/*
PostCustomersCustomerFidPaymentMethodsCreatePendingParams contains all the parameters to send to the API endpoint

	for the post customers customer fid payment methods create pending operation.

	Typically these are written to a http.Request.
*/
type PostCustomersCustomerFidPaymentMethodsCreatePendingParams struct {

	/* CustomerFid.

	   Customer FID to use
	*/
	CustomerFid string

	// Payload.
	Payload *models.CreatePendingPaymentMethodPayload

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post customers customer fid payment methods create pending params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostCustomersCustomerFidPaymentMethodsCreatePendingParams) WithDefaults() *PostCustomersCustomerFidPaymentMethodsCreatePendingParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post customers customer fid payment methods create pending params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostCustomersCustomerFidPaymentMethodsCreatePendingParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post customers customer fid payment methods create pending params
func (o *PostCustomersCustomerFidPaymentMethodsCreatePendingParams) WithTimeout(timeout time.Duration) *PostCustomersCustomerFidPaymentMethodsCreatePendingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post customers customer fid payment methods create pending params
func (o *PostCustomersCustomerFidPaymentMethodsCreatePendingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post customers customer fid payment methods create pending params
func (o *PostCustomersCustomerFidPaymentMethodsCreatePendingParams) WithContext(ctx context.Context) *PostCustomersCustomerFidPaymentMethodsCreatePendingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post customers customer fid payment methods create pending params
func (o *PostCustomersCustomerFidPaymentMethodsCreatePendingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post customers customer fid payment methods create pending params
func (o *PostCustomersCustomerFidPaymentMethodsCreatePendingParams) WithHTTPClient(client *http.Client) *PostCustomersCustomerFidPaymentMethodsCreatePendingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post customers customer fid payment methods create pending params
func (o *PostCustomersCustomerFidPaymentMethodsCreatePendingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerFid adds the customerFid to the post customers customer fid payment methods create pending params
func (o *PostCustomersCustomerFidPaymentMethodsCreatePendingParams) WithCustomerFid(customerFid string) *PostCustomersCustomerFidPaymentMethodsCreatePendingParams {
	o.SetCustomerFid(customerFid)
	return o
}

// SetCustomerFid adds the customerFid to the post customers customer fid payment methods create pending params
func (o *PostCustomersCustomerFidPaymentMethodsCreatePendingParams) SetCustomerFid(customerFid string) {
	o.CustomerFid = customerFid
}

// WithPayload adds the payload to the post customers customer fid payment methods create pending params
func (o *PostCustomersCustomerFidPaymentMethodsCreatePendingParams) WithPayload(payload *models.CreatePendingPaymentMethodPayload) *PostCustomersCustomerFidPaymentMethodsCreatePendingParams {
	o.SetPayload(payload)
	return o
}

// SetPayload adds the payload to the post customers customer fid payment methods create pending params
func (o *PostCustomersCustomerFidPaymentMethodsCreatePendingParams) SetPayload(payload *models.CreatePendingPaymentMethodPayload) {
	o.Payload = payload
}

// WriteToRequest writes these params to a swagger request
func (o *PostCustomersCustomerFidPaymentMethodsCreatePendingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
