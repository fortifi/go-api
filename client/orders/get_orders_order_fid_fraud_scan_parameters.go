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

// NewGetOrdersOrderFidFraudScanParams creates a new GetOrdersOrderFidFraudScanParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrdersOrderFidFraudScanParams() *GetOrdersOrderFidFraudScanParams {
	return &GetOrdersOrderFidFraudScanParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrdersOrderFidFraudScanParamsWithTimeout creates a new GetOrdersOrderFidFraudScanParams object
// with the ability to set a timeout on a request.
func NewGetOrdersOrderFidFraudScanParamsWithTimeout(timeout time.Duration) *GetOrdersOrderFidFraudScanParams {
	return &GetOrdersOrderFidFraudScanParams{
		timeout: timeout,
	}
}

// NewGetOrdersOrderFidFraudScanParamsWithContext creates a new GetOrdersOrderFidFraudScanParams object
// with the ability to set a context for a request.
func NewGetOrdersOrderFidFraudScanParamsWithContext(ctx context.Context) *GetOrdersOrderFidFraudScanParams {
	return &GetOrdersOrderFidFraudScanParams{
		Context: ctx,
	}
}

// NewGetOrdersOrderFidFraudScanParamsWithHTTPClient creates a new GetOrdersOrderFidFraudScanParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrdersOrderFidFraudScanParamsWithHTTPClient(client *http.Client) *GetOrdersOrderFidFraudScanParams {
	return &GetOrdersOrderFidFraudScanParams{
		HTTPClient: client,
	}
}

/*
GetOrdersOrderFidFraudScanParams contains all the parameters to send to the API endpoint

	for the get orders order fid fraud scan operation.

	Typically these are written to a http.Request.
*/
type GetOrdersOrderFidFraudScanParams struct {

	// OrderFid.
	OrderFid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get orders order fid fraud scan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrdersOrderFidFraudScanParams) WithDefaults() *GetOrdersOrderFidFraudScanParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get orders order fid fraud scan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrdersOrderFidFraudScanParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get orders order fid fraud scan params
func (o *GetOrdersOrderFidFraudScanParams) WithTimeout(timeout time.Duration) *GetOrdersOrderFidFraudScanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get orders order fid fraud scan params
func (o *GetOrdersOrderFidFraudScanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get orders order fid fraud scan params
func (o *GetOrdersOrderFidFraudScanParams) WithContext(ctx context.Context) *GetOrdersOrderFidFraudScanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get orders order fid fraud scan params
func (o *GetOrdersOrderFidFraudScanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get orders order fid fraud scan params
func (o *GetOrdersOrderFidFraudScanParams) WithHTTPClient(client *http.Client) *GetOrdersOrderFidFraudScanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get orders order fid fraud scan params
func (o *GetOrdersOrderFidFraudScanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrderFid adds the orderFid to the get orders order fid fraud scan params
func (o *GetOrdersOrderFidFraudScanParams) WithOrderFid(orderFid string) *GetOrdersOrderFidFraudScanParams {
	o.SetOrderFid(orderFid)
	return o
}

// SetOrderFid adds the orderFid to the get orders order fid fraud scan params
func (o *GetOrdersOrderFidFraudScanParams) SetOrderFid(orderFid string) {
	o.OrderFid = orderFid
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrdersOrderFidFraudScanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param orderFid
	if err := r.SetPathParam("orderFid", o.OrderFid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
