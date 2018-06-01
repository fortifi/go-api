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

	models "github.com/fortifi/go-api/models"
)

// NewPutCustomersCustomerFidPaymentMethodsCardsCardFidParams creates a new PutCustomersCustomerFidPaymentMethodsCardsCardFidParams object
// with the default values initialized.
func NewPutCustomersCustomerFidPaymentMethodsCardsCardFidParams() *PutCustomersCustomerFidPaymentMethodsCardsCardFidParams {
	var ()
	return &PutCustomersCustomerFidPaymentMethodsCardsCardFidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutCustomersCustomerFidPaymentMethodsCardsCardFidParamsWithTimeout creates a new PutCustomersCustomerFidPaymentMethodsCardsCardFidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutCustomersCustomerFidPaymentMethodsCardsCardFidParamsWithTimeout(timeout time.Duration) *PutCustomersCustomerFidPaymentMethodsCardsCardFidParams {
	var ()
	return &PutCustomersCustomerFidPaymentMethodsCardsCardFidParams{

		timeout: timeout,
	}
}

// NewPutCustomersCustomerFidPaymentMethodsCardsCardFidParamsWithContext creates a new PutCustomersCustomerFidPaymentMethodsCardsCardFidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutCustomersCustomerFidPaymentMethodsCardsCardFidParamsWithContext(ctx context.Context) *PutCustomersCustomerFidPaymentMethodsCardsCardFidParams {
	var ()
	return &PutCustomersCustomerFidPaymentMethodsCardsCardFidParams{

		Context: ctx,
	}
}

// NewPutCustomersCustomerFidPaymentMethodsCardsCardFidParamsWithHTTPClient creates a new PutCustomersCustomerFidPaymentMethodsCardsCardFidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutCustomersCustomerFidPaymentMethodsCardsCardFidParamsWithHTTPClient(client *http.Client) *PutCustomersCustomerFidPaymentMethodsCardsCardFidParams {
	var ()
	return &PutCustomersCustomerFidPaymentMethodsCardsCardFidParams{
		HTTPClient: client,
	}
}

/*PutCustomersCustomerFidPaymentMethodsCardsCardFidParams contains all the parameters to send to the API endpoint
for the put customers customer fid payment methods cards card fid operation typically these are written to a http.Request
*/
type PutCustomersCustomerFidPaymentMethodsCardsCardFidParams struct {

	/*CardFid*/
	CardFid string
	/*CustomerFid
	  Customer FID to use

	*/
	CustomerFid string
	/*Payload*/
	Payload *models.CardUpdatePayload

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put customers customer fid payment methods cards card fid params
func (o *PutCustomersCustomerFidPaymentMethodsCardsCardFidParams) WithTimeout(timeout time.Duration) *PutCustomersCustomerFidPaymentMethodsCardsCardFidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put customers customer fid payment methods cards card fid params
func (o *PutCustomersCustomerFidPaymentMethodsCardsCardFidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put customers customer fid payment methods cards card fid params
func (o *PutCustomersCustomerFidPaymentMethodsCardsCardFidParams) WithContext(ctx context.Context) *PutCustomersCustomerFidPaymentMethodsCardsCardFidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put customers customer fid payment methods cards card fid params
func (o *PutCustomersCustomerFidPaymentMethodsCardsCardFidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put customers customer fid payment methods cards card fid params
func (o *PutCustomersCustomerFidPaymentMethodsCardsCardFidParams) WithHTTPClient(client *http.Client) *PutCustomersCustomerFidPaymentMethodsCardsCardFidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put customers customer fid payment methods cards card fid params
func (o *PutCustomersCustomerFidPaymentMethodsCardsCardFidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCardFid adds the cardFid to the put customers customer fid payment methods cards card fid params
func (o *PutCustomersCustomerFidPaymentMethodsCardsCardFidParams) WithCardFid(cardFid string) *PutCustomersCustomerFidPaymentMethodsCardsCardFidParams {
	o.SetCardFid(cardFid)
	return o
}

// SetCardFid adds the cardFid to the put customers customer fid payment methods cards card fid params
func (o *PutCustomersCustomerFidPaymentMethodsCardsCardFidParams) SetCardFid(cardFid string) {
	o.CardFid = cardFid
}

// WithCustomerFid adds the customerFid to the put customers customer fid payment methods cards card fid params
func (o *PutCustomersCustomerFidPaymentMethodsCardsCardFidParams) WithCustomerFid(customerFid string) *PutCustomersCustomerFidPaymentMethodsCardsCardFidParams {
	o.SetCustomerFid(customerFid)
	return o
}

// SetCustomerFid adds the customerFid to the put customers customer fid payment methods cards card fid params
func (o *PutCustomersCustomerFidPaymentMethodsCardsCardFidParams) SetCustomerFid(customerFid string) {
	o.CustomerFid = customerFid
}

// WithPayload adds the payload to the put customers customer fid payment methods cards card fid params
func (o *PutCustomersCustomerFidPaymentMethodsCardsCardFidParams) WithPayload(payload *models.CardUpdatePayload) *PutCustomersCustomerFidPaymentMethodsCardsCardFidParams {
	o.SetPayload(payload)
	return o
}

// SetPayload adds the payload to the put customers customer fid payment methods cards card fid params
func (o *PutCustomersCustomerFidPaymentMethodsCardsCardFidParams) SetPayload(payload *models.CardUpdatePayload) {
	o.Payload = payload
}

// WriteToRequest writes these params to a swagger request
func (o *PutCustomersCustomerFidPaymentMethodsCardsCardFidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cardFid
	if err := r.SetPathParam("cardFid", o.CardFid); err != nil {
		return err
	}

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
