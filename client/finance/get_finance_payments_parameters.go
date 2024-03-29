// Code generated by go-swagger; DO NOT EDIT.

package finance

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

// NewGetFinancePaymentsParams creates a new GetFinancePaymentsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetFinancePaymentsParams() *GetFinancePaymentsParams {
	return &GetFinancePaymentsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetFinancePaymentsParamsWithTimeout creates a new GetFinancePaymentsParams object
// with the ability to set a timeout on a request.
func NewGetFinancePaymentsParamsWithTimeout(timeout time.Duration) *GetFinancePaymentsParams {
	return &GetFinancePaymentsParams{
		timeout: timeout,
	}
}

// NewGetFinancePaymentsParamsWithContext creates a new GetFinancePaymentsParams object
// with the ability to set a context for a request.
func NewGetFinancePaymentsParamsWithContext(ctx context.Context) *GetFinancePaymentsParams {
	return &GetFinancePaymentsParams{
		Context: ctx,
	}
}

// NewGetFinancePaymentsParamsWithHTTPClient creates a new GetFinancePaymentsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetFinancePaymentsParamsWithHTTPClient(client *http.Client) *GetFinancePaymentsParams {
	return &GetFinancePaymentsParams{
		HTTPClient: client,
	}
}

/*
GetFinancePaymentsParams contains all the parameters to send to the API endpoint

	for the get finance payments operation.

	Typically these are written to a http.Request.
*/
type GetFinancePaymentsParams struct {

	/* Amount.

	   Transaction amount

	   Format: float
	*/
	Amount *float32

	/* Arn.

	   Acquirer reference number
	*/
	Arn *string

	/* AuthCode.

	   Transaction Auth Code
	*/
	AuthCode *string

	/* Currency.

	   3 Character Currency Code
	*/
	Currency *string

	/* DateTolerance.

	   Number of days tolerance on the transaction date
	*/
	DateTolerance *float64

	/* First6last4.

	   Locate based on the first 6 and last 4 of a card number
	*/
	First6last4 *string

	/* TransactionDate.

	   Transaction Date

	   Format: date-time
	*/
	TransactionDate *strfmt.DateTime

	/* TransactionID.

	   Transaction ID
	*/
	TransactionID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get finance payments params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFinancePaymentsParams) WithDefaults() *GetFinancePaymentsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get finance payments params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFinancePaymentsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get finance payments params
func (o *GetFinancePaymentsParams) WithTimeout(timeout time.Duration) *GetFinancePaymentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get finance payments params
func (o *GetFinancePaymentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get finance payments params
func (o *GetFinancePaymentsParams) WithContext(ctx context.Context) *GetFinancePaymentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get finance payments params
func (o *GetFinancePaymentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get finance payments params
func (o *GetFinancePaymentsParams) WithHTTPClient(client *http.Client) *GetFinancePaymentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get finance payments params
func (o *GetFinancePaymentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAmount adds the amount to the get finance payments params
func (o *GetFinancePaymentsParams) WithAmount(amount *float32) *GetFinancePaymentsParams {
	o.SetAmount(amount)
	return o
}

// SetAmount adds the amount to the get finance payments params
func (o *GetFinancePaymentsParams) SetAmount(amount *float32) {
	o.Amount = amount
}

// WithArn adds the arn to the get finance payments params
func (o *GetFinancePaymentsParams) WithArn(arn *string) *GetFinancePaymentsParams {
	o.SetArn(arn)
	return o
}

// SetArn adds the arn to the get finance payments params
func (o *GetFinancePaymentsParams) SetArn(arn *string) {
	o.Arn = arn
}

// WithAuthCode adds the authCode to the get finance payments params
func (o *GetFinancePaymentsParams) WithAuthCode(authCode *string) *GetFinancePaymentsParams {
	o.SetAuthCode(authCode)
	return o
}

// SetAuthCode adds the authCode to the get finance payments params
func (o *GetFinancePaymentsParams) SetAuthCode(authCode *string) {
	o.AuthCode = authCode
}

// WithCurrency adds the currency to the get finance payments params
func (o *GetFinancePaymentsParams) WithCurrency(currency *string) *GetFinancePaymentsParams {
	o.SetCurrency(currency)
	return o
}

// SetCurrency adds the currency to the get finance payments params
func (o *GetFinancePaymentsParams) SetCurrency(currency *string) {
	o.Currency = currency
}

// WithDateTolerance adds the dateTolerance to the get finance payments params
func (o *GetFinancePaymentsParams) WithDateTolerance(dateTolerance *float64) *GetFinancePaymentsParams {
	o.SetDateTolerance(dateTolerance)
	return o
}

// SetDateTolerance adds the dateTolerance to the get finance payments params
func (o *GetFinancePaymentsParams) SetDateTolerance(dateTolerance *float64) {
	o.DateTolerance = dateTolerance
}

// WithFirst6last4 adds the first6last4 to the get finance payments params
func (o *GetFinancePaymentsParams) WithFirst6last4(first6last4 *string) *GetFinancePaymentsParams {
	o.SetFirst6last4(first6last4)
	return o
}

// SetFirst6last4 adds the first6last4 to the get finance payments params
func (o *GetFinancePaymentsParams) SetFirst6last4(first6last4 *string) {
	o.First6last4 = first6last4
}

// WithTransactionDate adds the transactionDate to the get finance payments params
func (o *GetFinancePaymentsParams) WithTransactionDate(transactionDate *strfmt.DateTime) *GetFinancePaymentsParams {
	o.SetTransactionDate(transactionDate)
	return o
}

// SetTransactionDate adds the transactionDate to the get finance payments params
func (o *GetFinancePaymentsParams) SetTransactionDate(transactionDate *strfmt.DateTime) {
	o.TransactionDate = transactionDate
}

// WithTransactionID adds the transactionID to the get finance payments params
func (o *GetFinancePaymentsParams) WithTransactionID(transactionID *string) *GetFinancePaymentsParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the get finance payments params
func (o *GetFinancePaymentsParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetFinancePaymentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Amount != nil {

		// query param amount
		var qrAmount float32

		if o.Amount != nil {
			qrAmount = *o.Amount
		}
		qAmount := swag.FormatFloat32(qrAmount)
		if qAmount != "" {

			if err := r.SetQueryParam("amount", qAmount); err != nil {
				return err
			}
		}
	}

	if o.Arn != nil {

		// query param arn
		var qrArn string

		if o.Arn != nil {
			qrArn = *o.Arn
		}
		qArn := qrArn
		if qArn != "" {

			if err := r.SetQueryParam("arn", qArn); err != nil {
				return err
			}
		}
	}

	if o.AuthCode != nil {

		// query param authCode
		var qrAuthCode string

		if o.AuthCode != nil {
			qrAuthCode = *o.AuthCode
		}
		qAuthCode := qrAuthCode
		if qAuthCode != "" {

			if err := r.SetQueryParam("authCode", qAuthCode); err != nil {
				return err
			}
		}
	}

	if o.Currency != nil {

		// query param currency
		var qrCurrency string

		if o.Currency != nil {
			qrCurrency = *o.Currency
		}
		qCurrency := qrCurrency
		if qCurrency != "" {

			if err := r.SetQueryParam("currency", qCurrency); err != nil {
				return err
			}
		}
	}

	if o.DateTolerance != nil {

		// query param dateTolerance
		var qrDateTolerance float64

		if o.DateTolerance != nil {
			qrDateTolerance = *o.DateTolerance
		}
		qDateTolerance := swag.FormatFloat64(qrDateTolerance)
		if qDateTolerance != "" {

			if err := r.SetQueryParam("dateTolerance", qDateTolerance); err != nil {
				return err
			}
		}
	}

	if o.First6last4 != nil {

		// query param first6last4
		var qrFirst6last4 string

		if o.First6last4 != nil {
			qrFirst6last4 = *o.First6last4
		}
		qFirst6last4 := qrFirst6last4
		if qFirst6last4 != "" {

			if err := r.SetQueryParam("first6last4", qFirst6last4); err != nil {
				return err
			}
		}
	}

	if o.TransactionDate != nil {

		// query param transactionDate
		var qrTransactionDate strfmt.DateTime

		if o.TransactionDate != nil {
			qrTransactionDate = *o.TransactionDate
		}
		qTransactionDate := qrTransactionDate.String()
		if qTransactionDate != "" {

			if err := r.SetQueryParam("transactionDate", qTransactionDate); err != nil {
				return err
			}
		}
	}

	if o.TransactionID != nil {

		// query param transactionId
		var qrTransactionID string

		if o.TransactionID != nil {
			qrTransactionID = *o.TransactionID
		}
		qTransactionID := qrTransactionID
		if qTransactionID != "" {

			if err := r.SetQueryParam("transactionId", qTransactionID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
