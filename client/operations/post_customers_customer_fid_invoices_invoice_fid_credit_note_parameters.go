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

// NewPostCustomersCustomerFidInvoicesInvoiceFidCreditNoteParams creates a new PostCustomersCustomerFidInvoicesInvoiceFidCreditNoteParams object
// with the default values initialized.
func NewPostCustomersCustomerFidInvoicesInvoiceFidCreditNoteParams() *PostCustomersCustomerFidInvoicesInvoiceFidCreditNoteParams {
	var ()
	return &PostCustomersCustomerFidInvoicesInvoiceFidCreditNoteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostCustomersCustomerFidInvoicesInvoiceFidCreditNoteParamsWithTimeout creates a new PostCustomersCustomerFidInvoicesInvoiceFidCreditNoteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostCustomersCustomerFidInvoicesInvoiceFidCreditNoteParamsWithTimeout(timeout time.Duration) *PostCustomersCustomerFidInvoicesInvoiceFidCreditNoteParams {
	var ()
	return &PostCustomersCustomerFidInvoicesInvoiceFidCreditNoteParams{

		timeout: timeout,
	}
}

// NewPostCustomersCustomerFidInvoicesInvoiceFidCreditNoteParamsWithContext creates a new PostCustomersCustomerFidInvoicesInvoiceFidCreditNoteParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostCustomersCustomerFidInvoicesInvoiceFidCreditNoteParamsWithContext(ctx context.Context) *PostCustomersCustomerFidInvoicesInvoiceFidCreditNoteParams {
	var ()
	return &PostCustomersCustomerFidInvoicesInvoiceFidCreditNoteParams{

		Context: ctx,
	}
}

// NewPostCustomersCustomerFidInvoicesInvoiceFidCreditNoteParamsWithHTTPClient creates a new PostCustomersCustomerFidInvoicesInvoiceFidCreditNoteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostCustomersCustomerFidInvoicesInvoiceFidCreditNoteParamsWithHTTPClient(client *http.Client) *PostCustomersCustomerFidInvoicesInvoiceFidCreditNoteParams {
	var ()
	return &PostCustomersCustomerFidInvoicesInvoiceFidCreditNoteParams{
		HTTPClient: client,
	}
}

/*PostCustomersCustomerFidInvoicesInvoiceFidCreditNoteParams contains all the parameters to send to the API endpoint
for the post customers customer fid invoices invoice fid credit note operation typically these are written to a http.Request
*/
type PostCustomersCustomerFidInvoicesInvoiceFidCreditNoteParams struct {

	/*CustomerFid
	  Customer FID to use

	*/
	CustomerFid string
	/*InvoiceFid
	  Invoice Fid to use

	*/
	InvoiceFid string
	/*Payload*/
	Payload *models.InvoiceCreditNotePayload

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post customers customer fid invoices invoice fid credit note params
func (o *PostCustomersCustomerFidInvoicesInvoiceFidCreditNoteParams) WithTimeout(timeout time.Duration) *PostCustomersCustomerFidInvoicesInvoiceFidCreditNoteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post customers customer fid invoices invoice fid credit note params
func (o *PostCustomersCustomerFidInvoicesInvoiceFidCreditNoteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post customers customer fid invoices invoice fid credit note params
func (o *PostCustomersCustomerFidInvoicesInvoiceFidCreditNoteParams) WithContext(ctx context.Context) *PostCustomersCustomerFidInvoicesInvoiceFidCreditNoteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post customers customer fid invoices invoice fid credit note params
func (o *PostCustomersCustomerFidInvoicesInvoiceFidCreditNoteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post customers customer fid invoices invoice fid credit note params
func (o *PostCustomersCustomerFidInvoicesInvoiceFidCreditNoteParams) WithHTTPClient(client *http.Client) *PostCustomersCustomerFidInvoicesInvoiceFidCreditNoteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post customers customer fid invoices invoice fid credit note params
func (o *PostCustomersCustomerFidInvoicesInvoiceFidCreditNoteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerFid adds the customerFid to the post customers customer fid invoices invoice fid credit note params
func (o *PostCustomersCustomerFidInvoicesInvoiceFidCreditNoteParams) WithCustomerFid(customerFid string) *PostCustomersCustomerFidInvoicesInvoiceFidCreditNoteParams {
	o.SetCustomerFid(customerFid)
	return o
}

// SetCustomerFid adds the customerFid to the post customers customer fid invoices invoice fid credit note params
func (o *PostCustomersCustomerFidInvoicesInvoiceFidCreditNoteParams) SetCustomerFid(customerFid string) {
	o.CustomerFid = customerFid
}

// WithInvoiceFid adds the invoiceFid to the post customers customer fid invoices invoice fid credit note params
func (o *PostCustomersCustomerFidInvoicesInvoiceFidCreditNoteParams) WithInvoiceFid(invoiceFid string) *PostCustomersCustomerFidInvoicesInvoiceFidCreditNoteParams {
	o.SetInvoiceFid(invoiceFid)
	return o
}

// SetInvoiceFid adds the invoiceFid to the post customers customer fid invoices invoice fid credit note params
func (o *PostCustomersCustomerFidInvoicesInvoiceFidCreditNoteParams) SetInvoiceFid(invoiceFid string) {
	o.InvoiceFid = invoiceFid
}

// WithPayload adds the payload to the post customers customer fid invoices invoice fid credit note params
func (o *PostCustomersCustomerFidInvoicesInvoiceFidCreditNoteParams) WithPayload(payload *models.InvoiceCreditNotePayload) *PostCustomersCustomerFidInvoicesInvoiceFidCreditNoteParams {
	o.SetPayload(payload)
	return o
}

// SetPayload adds the payload to the post customers customer fid invoices invoice fid credit note params
func (o *PostCustomersCustomerFidInvoicesInvoiceFidCreditNoteParams) SetPayload(payload *models.InvoiceCreditNotePayload) {
	o.Payload = payload
}

// WriteToRequest writes these params to a swagger request
func (o *PostCustomersCustomerFidInvoicesInvoiceFidCreditNoteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param customerFid
	if err := r.SetPathParam("customerFid", o.CustomerFid); err != nil {
		return err
	}

	// path param invoiceFid
	if err := r.SetPathParam("invoiceFid", o.InvoiceFid); err != nil {
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
