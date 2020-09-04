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
	"github.com/go-openapi/swag"
)

// NewPutCustomersCustomerFidPaymentsPaymentFidRefundParams creates a new PutCustomersCustomerFidPaymentsPaymentFidRefundParams object
// with the default values initialized.
func NewPutCustomersCustomerFidPaymentsPaymentFidRefundParams() *PutCustomersCustomerFidPaymentsPaymentFidRefundParams {
	var ()
	return &PutCustomersCustomerFidPaymentsPaymentFidRefundParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutCustomersCustomerFidPaymentsPaymentFidRefundParamsWithTimeout creates a new PutCustomersCustomerFidPaymentsPaymentFidRefundParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutCustomersCustomerFidPaymentsPaymentFidRefundParamsWithTimeout(timeout time.Duration) *PutCustomersCustomerFidPaymentsPaymentFidRefundParams {
	var ()
	return &PutCustomersCustomerFidPaymentsPaymentFidRefundParams{

		timeout: timeout,
	}
}

// NewPutCustomersCustomerFidPaymentsPaymentFidRefundParamsWithContext creates a new PutCustomersCustomerFidPaymentsPaymentFidRefundParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutCustomersCustomerFidPaymentsPaymentFidRefundParamsWithContext(ctx context.Context) *PutCustomersCustomerFidPaymentsPaymentFidRefundParams {
	var ()
	return &PutCustomersCustomerFidPaymentsPaymentFidRefundParams{

		Context: ctx,
	}
}

// NewPutCustomersCustomerFidPaymentsPaymentFidRefundParamsWithHTTPClient creates a new PutCustomersCustomerFidPaymentsPaymentFidRefundParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutCustomersCustomerFidPaymentsPaymentFidRefundParamsWithHTTPClient(client *http.Client) *PutCustomersCustomerFidPaymentsPaymentFidRefundParams {
	var ()
	return &PutCustomersCustomerFidPaymentsPaymentFidRefundParams{
		HTTPClient: client,
	}
}

/*PutCustomersCustomerFidPaymentsPaymentFidRefundParams contains all the parameters to send to the API endpoint
for the put customers customer fid payments payment fid refund operation typically these are written to a http.Request
*/
type PutCustomersCustomerFidPaymentsPaymentFidRefundParams struct {

	/*AddCreditToInvoice
	  Add Credit to Invoice

	*/
	AddCreditToInvoice bool
	/*CustomerFid
	  Customer FID to use

	*/
	CustomerFid string
	/*PaymentFid
	  Payment FID to use

	*/
	PaymentFid string
	/*ReasonFid
	  Reason FID

	*/
	ReasonFid string
	/*RefundType
	  Refund Type

	*/
	RefundType string
	/*TotalRefund
	  Total Refund

	*/
	TotalRefund float64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put customers customer fid payments payment fid refund params
func (o *PutCustomersCustomerFidPaymentsPaymentFidRefundParams) WithTimeout(timeout time.Duration) *PutCustomersCustomerFidPaymentsPaymentFidRefundParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put customers customer fid payments payment fid refund params
func (o *PutCustomersCustomerFidPaymentsPaymentFidRefundParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put customers customer fid payments payment fid refund params
func (o *PutCustomersCustomerFidPaymentsPaymentFidRefundParams) WithContext(ctx context.Context) *PutCustomersCustomerFidPaymentsPaymentFidRefundParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put customers customer fid payments payment fid refund params
func (o *PutCustomersCustomerFidPaymentsPaymentFidRefundParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put customers customer fid payments payment fid refund params
func (o *PutCustomersCustomerFidPaymentsPaymentFidRefundParams) WithHTTPClient(client *http.Client) *PutCustomersCustomerFidPaymentsPaymentFidRefundParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put customers customer fid payments payment fid refund params
func (o *PutCustomersCustomerFidPaymentsPaymentFidRefundParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddCreditToInvoice adds the addCreditToInvoice to the put customers customer fid payments payment fid refund params
func (o *PutCustomersCustomerFidPaymentsPaymentFidRefundParams) WithAddCreditToInvoice(addCreditToInvoice bool) *PutCustomersCustomerFidPaymentsPaymentFidRefundParams {
	o.SetAddCreditToInvoice(addCreditToInvoice)
	return o
}

// SetAddCreditToInvoice adds the addCreditToInvoice to the put customers customer fid payments payment fid refund params
func (o *PutCustomersCustomerFidPaymentsPaymentFidRefundParams) SetAddCreditToInvoice(addCreditToInvoice bool) {
	o.AddCreditToInvoice = addCreditToInvoice
}

// WithCustomerFid adds the customerFid to the put customers customer fid payments payment fid refund params
func (o *PutCustomersCustomerFidPaymentsPaymentFidRefundParams) WithCustomerFid(customerFid string) *PutCustomersCustomerFidPaymentsPaymentFidRefundParams {
	o.SetCustomerFid(customerFid)
	return o
}

// SetCustomerFid adds the customerFid to the put customers customer fid payments payment fid refund params
func (o *PutCustomersCustomerFidPaymentsPaymentFidRefundParams) SetCustomerFid(customerFid string) {
	o.CustomerFid = customerFid
}

// WithPaymentFid adds the paymentFid to the put customers customer fid payments payment fid refund params
func (o *PutCustomersCustomerFidPaymentsPaymentFidRefundParams) WithPaymentFid(paymentFid string) *PutCustomersCustomerFidPaymentsPaymentFidRefundParams {
	o.SetPaymentFid(paymentFid)
	return o
}

// SetPaymentFid adds the paymentFid to the put customers customer fid payments payment fid refund params
func (o *PutCustomersCustomerFidPaymentsPaymentFidRefundParams) SetPaymentFid(paymentFid string) {
	o.PaymentFid = paymentFid
}

// WithReasonFid adds the reasonFid to the put customers customer fid payments payment fid refund params
func (o *PutCustomersCustomerFidPaymentsPaymentFidRefundParams) WithReasonFid(reasonFid string) *PutCustomersCustomerFidPaymentsPaymentFidRefundParams {
	o.SetReasonFid(reasonFid)
	return o
}

// SetReasonFid adds the reasonFid to the put customers customer fid payments payment fid refund params
func (o *PutCustomersCustomerFidPaymentsPaymentFidRefundParams) SetReasonFid(reasonFid string) {
	o.ReasonFid = reasonFid
}

// WithRefundType adds the refundType to the put customers customer fid payments payment fid refund params
func (o *PutCustomersCustomerFidPaymentsPaymentFidRefundParams) WithRefundType(refundType string) *PutCustomersCustomerFidPaymentsPaymentFidRefundParams {
	o.SetRefundType(refundType)
	return o
}

// SetRefundType adds the refundType to the put customers customer fid payments payment fid refund params
func (o *PutCustomersCustomerFidPaymentsPaymentFidRefundParams) SetRefundType(refundType string) {
	o.RefundType = refundType
}

// WithTotalRefund adds the totalRefund to the put customers customer fid payments payment fid refund params
func (o *PutCustomersCustomerFidPaymentsPaymentFidRefundParams) WithTotalRefund(totalRefund float64) *PutCustomersCustomerFidPaymentsPaymentFidRefundParams {
	o.SetTotalRefund(totalRefund)
	return o
}

// SetTotalRefund adds the totalRefund to the put customers customer fid payments payment fid refund params
func (o *PutCustomersCustomerFidPaymentsPaymentFidRefundParams) SetTotalRefund(totalRefund float64) {
	o.TotalRefund = totalRefund
}

// WriteToRequest writes these params to a swagger request
func (o *PutCustomersCustomerFidPaymentsPaymentFidRefundParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// form param addCreditToInvoice
	frAddCreditToInvoice := o.AddCreditToInvoice
	fAddCreditToInvoice := swag.FormatBool(frAddCreditToInvoice)
	if fAddCreditToInvoice != "" {
		if err := r.SetFormParam("addCreditToInvoice", fAddCreditToInvoice); err != nil {
			return err
		}
	}

	// path param customerFid
	if err := r.SetPathParam("customerFid", o.CustomerFid); err != nil {
		return err
	}

	// path param paymentFid
	if err := r.SetPathParam("paymentFid", o.PaymentFid); err != nil {
		return err
	}

	// form param reasonFid
	frReasonFid := o.ReasonFid
	fReasonFid := frReasonFid
	if fReasonFid != "" {
		if err := r.SetFormParam("reasonFid", fReasonFid); err != nil {
			return err
		}
	}

	// form param refundType
	frRefundType := o.RefundType
	fRefundType := frRefundType
	if fRefundType != "" {
		if err := r.SetFormParam("refundType", fRefundType); err != nil {
			return err
		}
	}

	// form param totalRefund
	frTotalRefund := o.TotalRefund
	fTotalRefund := swag.FormatFloat64(frTotalRefund)
	if fTotalRefund != "" {
		if err := r.SetFormParam("totalRefund", fTotalRefund); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
