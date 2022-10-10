// Code generated by go-swagger; DO NOT EDIT.

package orders

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

// NewPutOrdersOrderFidExternalReferenceParams creates a new PutOrdersOrderFidExternalReferenceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutOrdersOrderFidExternalReferenceParams() *PutOrdersOrderFidExternalReferenceParams {
	return &PutOrdersOrderFidExternalReferenceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutOrdersOrderFidExternalReferenceParamsWithTimeout creates a new PutOrdersOrderFidExternalReferenceParams object
// with the ability to set a timeout on a request.
func NewPutOrdersOrderFidExternalReferenceParamsWithTimeout(timeout time.Duration) *PutOrdersOrderFidExternalReferenceParams {
	return &PutOrdersOrderFidExternalReferenceParams{
		timeout: timeout,
	}
}

// NewPutOrdersOrderFidExternalReferenceParamsWithContext creates a new PutOrdersOrderFidExternalReferenceParams object
// with the ability to set a context for a request.
func NewPutOrdersOrderFidExternalReferenceParamsWithContext(ctx context.Context) *PutOrdersOrderFidExternalReferenceParams {
	return &PutOrdersOrderFidExternalReferenceParams{
		Context: ctx,
	}
}

// NewPutOrdersOrderFidExternalReferenceParamsWithHTTPClient creates a new PutOrdersOrderFidExternalReferenceParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutOrdersOrderFidExternalReferenceParamsWithHTTPClient(client *http.Client) *PutOrdersOrderFidExternalReferenceParams {
	return &PutOrdersOrderFidExternalReferenceParams{
		HTTPClient: client,
	}
}

/* PutOrdersOrderFidExternalReferenceParams contains all the parameters to send to the API endpoint
   for the put orders order fid external reference operation.

   Typically these are written to a http.Request.
*/
type PutOrdersOrderFidExternalReferenceParams struct {

	/* DisplayName.

	   New Display Name
	*/
	DisplayName *string

	/* ExternalReference.

	   New External Reference
	*/
	ExternalReference string

	// OrderFid.
	OrderFid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put orders order fid external reference params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutOrdersOrderFidExternalReferenceParams) WithDefaults() *PutOrdersOrderFidExternalReferenceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put orders order fid external reference params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutOrdersOrderFidExternalReferenceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put orders order fid external reference params
func (o *PutOrdersOrderFidExternalReferenceParams) WithTimeout(timeout time.Duration) *PutOrdersOrderFidExternalReferenceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put orders order fid external reference params
func (o *PutOrdersOrderFidExternalReferenceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put orders order fid external reference params
func (o *PutOrdersOrderFidExternalReferenceParams) WithContext(ctx context.Context) *PutOrdersOrderFidExternalReferenceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put orders order fid external reference params
func (o *PutOrdersOrderFidExternalReferenceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put orders order fid external reference params
func (o *PutOrdersOrderFidExternalReferenceParams) WithHTTPClient(client *http.Client) *PutOrdersOrderFidExternalReferenceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put orders order fid external reference params
func (o *PutOrdersOrderFidExternalReferenceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDisplayName adds the displayName to the put orders order fid external reference params
func (o *PutOrdersOrderFidExternalReferenceParams) WithDisplayName(displayName *string) *PutOrdersOrderFidExternalReferenceParams {
	o.SetDisplayName(displayName)
	return o
}

// SetDisplayName adds the displayName to the put orders order fid external reference params
func (o *PutOrdersOrderFidExternalReferenceParams) SetDisplayName(displayName *string) {
	o.DisplayName = displayName
}

// WithExternalReference adds the externalReference to the put orders order fid external reference params
func (o *PutOrdersOrderFidExternalReferenceParams) WithExternalReference(externalReference string) *PutOrdersOrderFidExternalReferenceParams {
	o.SetExternalReference(externalReference)
	return o
}

// SetExternalReference adds the externalReference to the put orders order fid external reference params
func (o *PutOrdersOrderFidExternalReferenceParams) SetExternalReference(externalReference string) {
	o.ExternalReference = externalReference
}

// WithOrderFid adds the orderFid to the put orders order fid external reference params
func (o *PutOrdersOrderFidExternalReferenceParams) WithOrderFid(orderFid string) *PutOrdersOrderFidExternalReferenceParams {
	o.SetOrderFid(orderFid)
	return o
}

// SetOrderFid adds the orderFid to the put orders order fid external reference params
func (o *PutOrdersOrderFidExternalReferenceParams) SetOrderFid(orderFid string) {
	o.OrderFid = orderFid
}

// WriteToRequest writes these params to a swagger request
func (o *PutOrdersOrderFidExternalReferenceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DisplayName != nil {

		// form param displayName
		var frDisplayName string
		if o.DisplayName != nil {
			frDisplayName = *o.DisplayName
		}
		fDisplayName := frDisplayName
		if fDisplayName != "" {
			if err := r.SetFormParam("displayName", fDisplayName); err != nil {
				return err
			}
		}
	}

	// form param externalReference
	frExternalReference := o.ExternalReference
	fExternalReference := frExternalReference
	if fExternalReference != "" {
		if err := r.SetFormParam("externalReference", fExternalReference); err != nil {
			return err
		}
	}

	// path param orderFid
	if err := r.SetPathParam("orderFid", o.OrderFid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}