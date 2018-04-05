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

	strfmt "github.com/go-openapi/strfmt"

	"github.com/fortifi/go-api/models"
)

// NewPostCustomersCustomerFidPaymentMethodsCardsParams creates a new PostCustomersCustomerFidPaymentMethodsCardsParams object
// with the default values initialized.
func NewPostCustomersCustomerFidPaymentMethodsCardsParams() *PostCustomersCustomerFidPaymentMethodsCardsParams {
	var ()
	return &PostCustomersCustomerFidPaymentMethodsCardsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostCustomersCustomerFidPaymentMethodsCardsParamsWithTimeout creates a new PostCustomersCustomerFidPaymentMethodsCardsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostCustomersCustomerFidPaymentMethodsCardsParamsWithTimeout(timeout time.Duration) *PostCustomersCustomerFidPaymentMethodsCardsParams {
	var ()
	return &PostCustomersCustomerFidPaymentMethodsCardsParams{

		timeout: timeout,
	}
}

// NewPostCustomersCustomerFidPaymentMethodsCardsParamsWithContext creates a new PostCustomersCustomerFidPaymentMethodsCardsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostCustomersCustomerFidPaymentMethodsCardsParamsWithContext(ctx context.Context) *PostCustomersCustomerFidPaymentMethodsCardsParams {
	var ()
	return &PostCustomersCustomerFidPaymentMethodsCardsParams{

		Context: ctx,
	}
}

// NewPostCustomersCustomerFidPaymentMethodsCardsParamsWithHTTPClient creates a new PostCustomersCustomerFidPaymentMethodsCardsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostCustomersCustomerFidPaymentMethodsCardsParamsWithHTTPClient(client *http.Client) *PostCustomersCustomerFidPaymentMethodsCardsParams {
	var ()
	return &PostCustomersCustomerFidPaymentMethodsCardsParams{
		HTTPClient: client,
	}
}

/*PostCustomersCustomerFidPaymentMethodsCardsParams contains all the parameters to send to the API endpoint
for the post customers customer fid payment methods cards operation typically these are written to a http.Request
*/
type PostCustomersCustomerFidPaymentMethodsCardsParams struct {

	/*CustomerFid
	  Customer FID to use

	*/
	CustomerFid string
	/*Payload*/
	Payload *models.CardDataPayload

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post customers customer fid payment methods cards params
func (o *PostCustomersCustomerFidPaymentMethodsCardsParams) WithTimeout(timeout time.Duration) *PostCustomersCustomerFidPaymentMethodsCardsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post customers customer fid payment methods cards params
func (o *PostCustomersCustomerFidPaymentMethodsCardsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post customers customer fid payment methods cards params
func (o *PostCustomersCustomerFidPaymentMethodsCardsParams) WithContext(ctx context.Context) *PostCustomersCustomerFidPaymentMethodsCardsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post customers customer fid payment methods cards params
func (o *PostCustomersCustomerFidPaymentMethodsCardsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post customers customer fid payment methods cards params
func (o *PostCustomersCustomerFidPaymentMethodsCardsParams) WithHTTPClient(client *http.Client) *PostCustomersCustomerFidPaymentMethodsCardsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post customers customer fid payment methods cards params
func (o *PostCustomersCustomerFidPaymentMethodsCardsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerFid adds the customerFid to the post customers customer fid payment methods cards params
func (o *PostCustomersCustomerFidPaymentMethodsCardsParams) WithCustomerFid(customerFid string) *PostCustomersCustomerFidPaymentMethodsCardsParams {
	o.SetCustomerFid(customerFid)
	return o
}

// SetCustomerFid adds the customerFid to the post customers customer fid payment methods cards params
func (o *PostCustomersCustomerFidPaymentMethodsCardsParams) SetCustomerFid(customerFid string) {
	o.CustomerFid = customerFid
}

// WithPayload adds the payload to the post customers customer fid payment methods cards params
func (o *PostCustomersCustomerFidPaymentMethodsCardsParams) WithPayload(payload *models.CardDataPayload) *PostCustomersCustomerFidPaymentMethodsCardsParams {
	o.SetPayload(payload)
	return o
}

// SetPayload adds the payload to the post customers customer fid payment methods cards params
func (o *PostCustomersCustomerFidPaymentMethodsCardsParams) SetPayload(payload *models.CardDataPayload) {
	o.Payload = payload
}

// WriteToRequest writes these params to a swagger request
func (o *PostCustomersCustomerFidPaymentMethodsCardsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
