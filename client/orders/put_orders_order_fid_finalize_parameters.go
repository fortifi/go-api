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

	"github.com/fortifi/go-api/models"
)

// NewPutOrdersOrderFidFinalizeParams creates a new PutOrdersOrderFidFinalizeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutOrdersOrderFidFinalizeParams() *PutOrdersOrderFidFinalizeParams {
	return &PutOrdersOrderFidFinalizeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutOrdersOrderFidFinalizeParamsWithTimeout creates a new PutOrdersOrderFidFinalizeParams object
// with the ability to set a timeout on a request.
func NewPutOrdersOrderFidFinalizeParamsWithTimeout(timeout time.Duration) *PutOrdersOrderFidFinalizeParams {
	return &PutOrdersOrderFidFinalizeParams{
		timeout: timeout,
	}
}

// NewPutOrdersOrderFidFinalizeParamsWithContext creates a new PutOrdersOrderFidFinalizeParams object
// with the ability to set a context for a request.
func NewPutOrdersOrderFidFinalizeParamsWithContext(ctx context.Context) *PutOrdersOrderFidFinalizeParams {
	return &PutOrdersOrderFidFinalizeParams{
		Context: ctx,
	}
}

// NewPutOrdersOrderFidFinalizeParamsWithHTTPClient creates a new PutOrdersOrderFidFinalizeParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutOrdersOrderFidFinalizeParamsWithHTTPClient(client *http.Client) *PutOrdersOrderFidFinalizeParams {
	return &PutOrdersOrderFidFinalizeParams{
		HTTPClient: client,
	}
}

/*
PutOrdersOrderFidFinalizeParams contains all the parameters to send to the API endpoint

	for the put orders order fid finalize operation.

	Typically these are written to a http.Request.
*/
type PutOrdersOrderFidFinalizeParams struct {

	// OrderFid.
	OrderFid string

	// Payload.
	Payload *models.FinalizeOrderPayload

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put orders order fid finalize params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutOrdersOrderFidFinalizeParams) WithDefaults() *PutOrdersOrderFidFinalizeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put orders order fid finalize params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutOrdersOrderFidFinalizeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put orders order fid finalize params
func (o *PutOrdersOrderFidFinalizeParams) WithTimeout(timeout time.Duration) *PutOrdersOrderFidFinalizeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put orders order fid finalize params
func (o *PutOrdersOrderFidFinalizeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put orders order fid finalize params
func (o *PutOrdersOrderFidFinalizeParams) WithContext(ctx context.Context) *PutOrdersOrderFidFinalizeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put orders order fid finalize params
func (o *PutOrdersOrderFidFinalizeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put orders order fid finalize params
func (o *PutOrdersOrderFidFinalizeParams) WithHTTPClient(client *http.Client) *PutOrdersOrderFidFinalizeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put orders order fid finalize params
func (o *PutOrdersOrderFidFinalizeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrderFid adds the orderFid to the put orders order fid finalize params
func (o *PutOrdersOrderFidFinalizeParams) WithOrderFid(orderFid string) *PutOrdersOrderFidFinalizeParams {
	o.SetOrderFid(orderFid)
	return o
}

// SetOrderFid adds the orderFid to the put orders order fid finalize params
func (o *PutOrdersOrderFidFinalizeParams) SetOrderFid(orderFid string) {
	o.OrderFid = orderFid
}

// WithPayload adds the payload to the put orders order fid finalize params
func (o *PutOrdersOrderFidFinalizeParams) WithPayload(payload *models.FinalizeOrderPayload) *PutOrdersOrderFidFinalizeParams {
	o.SetPayload(payload)
	return o
}

// SetPayload adds the payload to the put orders order fid finalize params
func (o *PutOrdersOrderFidFinalizeParams) SetPayload(payload *models.FinalizeOrderPayload) {
	o.Payload = payload
}

// WriteToRequest writes these params to a swagger request
func (o *PutOrdersOrderFidFinalizeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param orderFid
	if err := r.SetPathParam("orderFid", o.OrderFid); err != nil {
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
