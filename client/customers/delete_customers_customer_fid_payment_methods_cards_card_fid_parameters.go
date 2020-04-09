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

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteCustomersCustomerFidPaymentMethodsCardsCardFidParams creates a new DeleteCustomersCustomerFidPaymentMethodsCardsCardFidParams object
// with the default values initialized.
func NewDeleteCustomersCustomerFidPaymentMethodsCardsCardFidParams() *DeleteCustomersCustomerFidPaymentMethodsCardsCardFidParams {
	var ()
	return &DeleteCustomersCustomerFidPaymentMethodsCardsCardFidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCustomersCustomerFidPaymentMethodsCardsCardFidParamsWithTimeout creates a new DeleteCustomersCustomerFidPaymentMethodsCardsCardFidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteCustomersCustomerFidPaymentMethodsCardsCardFidParamsWithTimeout(timeout time.Duration) *DeleteCustomersCustomerFidPaymentMethodsCardsCardFidParams {
	var ()
	return &DeleteCustomersCustomerFidPaymentMethodsCardsCardFidParams{

		timeout: timeout,
	}
}

// NewDeleteCustomersCustomerFidPaymentMethodsCardsCardFidParamsWithContext creates a new DeleteCustomersCustomerFidPaymentMethodsCardsCardFidParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteCustomersCustomerFidPaymentMethodsCardsCardFidParamsWithContext(ctx context.Context) *DeleteCustomersCustomerFidPaymentMethodsCardsCardFidParams {
	var ()
	return &DeleteCustomersCustomerFidPaymentMethodsCardsCardFidParams{

		Context: ctx,
	}
}

// NewDeleteCustomersCustomerFidPaymentMethodsCardsCardFidParamsWithHTTPClient creates a new DeleteCustomersCustomerFidPaymentMethodsCardsCardFidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteCustomersCustomerFidPaymentMethodsCardsCardFidParamsWithHTTPClient(client *http.Client) *DeleteCustomersCustomerFidPaymentMethodsCardsCardFidParams {
	var ()
	return &DeleteCustomersCustomerFidPaymentMethodsCardsCardFidParams{
		HTTPClient: client,
	}
}

/*DeleteCustomersCustomerFidPaymentMethodsCardsCardFidParams contains all the parameters to send to the API endpoint
for the delete customers customer fid payment methods cards card fid operation typically these are written to a http.Request
*/
type DeleteCustomersCustomerFidPaymentMethodsCardsCardFidParams struct {

	/*CardFid*/
	CardFid string
	/*CustomerFid
	  Customer FID to use

	*/
	CustomerFid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete customers customer fid payment methods cards card fid params
func (o *DeleteCustomersCustomerFidPaymentMethodsCardsCardFidParams) WithTimeout(timeout time.Duration) *DeleteCustomersCustomerFidPaymentMethodsCardsCardFidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete customers customer fid payment methods cards card fid params
func (o *DeleteCustomersCustomerFidPaymentMethodsCardsCardFidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete customers customer fid payment methods cards card fid params
func (o *DeleteCustomersCustomerFidPaymentMethodsCardsCardFidParams) WithContext(ctx context.Context) *DeleteCustomersCustomerFidPaymentMethodsCardsCardFidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete customers customer fid payment methods cards card fid params
func (o *DeleteCustomersCustomerFidPaymentMethodsCardsCardFidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete customers customer fid payment methods cards card fid params
func (o *DeleteCustomersCustomerFidPaymentMethodsCardsCardFidParams) WithHTTPClient(client *http.Client) *DeleteCustomersCustomerFidPaymentMethodsCardsCardFidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete customers customer fid payment methods cards card fid params
func (o *DeleteCustomersCustomerFidPaymentMethodsCardsCardFidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCardFid adds the cardFid to the delete customers customer fid payment methods cards card fid params
func (o *DeleteCustomersCustomerFidPaymentMethodsCardsCardFidParams) WithCardFid(cardFid string) *DeleteCustomersCustomerFidPaymentMethodsCardsCardFidParams {
	o.SetCardFid(cardFid)
	return o
}

// SetCardFid adds the cardFid to the delete customers customer fid payment methods cards card fid params
func (o *DeleteCustomersCustomerFidPaymentMethodsCardsCardFidParams) SetCardFid(cardFid string) {
	o.CardFid = cardFid
}

// WithCustomerFid adds the customerFid to the delete customers customer fid payment methods cards card fid params
func (o *DeleteCustomersCustomerFidPaymentMethodsCardsCardFidParams) WithCustomerFid(customerFid string) *DeleteCustomersCustomerFidPaymentMethodsCardsCardFidParams {
	o.SetCustomerFid(customerFid)
	return o
}

// SetCustomerFid adds the customerFid to the delete customers customer fid payment methods cards card fid params
func (o *DeleteCustomersCustomerFidPaymentMethodsCardsCardFidParams) SetCustomerFid(customerFid string) {
	o.CustomerFid = customerFid
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCustomersCustomerFidPaymentMethodsCardsCardFidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}