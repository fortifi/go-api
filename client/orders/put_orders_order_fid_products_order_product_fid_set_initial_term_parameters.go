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

// NewPutOrdersOrderFidProductsOrderProductFidSetInitialTermParams creates a new PutOrdersOrderFidProductsOrderProductFidSetInitialTermParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutOrdersOrderFidProductsOrderProductFidSetInitialTermParams() *PutOrdersOrderFidProductsOrderProductFidSetInitialTermParams {
	return &PutOrdersOrderFidProductsOrderProductFidSetInitialTermParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutOrdersOrderFidProductsOrderProductFidSetInitialTermParamsWithTimeout creates a new PutOrdersOrderFidProductsOrderProductFidSetInitialTermParams object
// with the ability to set a timeout on a request.
func NewPutOrdersOrderFidProductsOrderProductFidSetInitialTermParamsWithTimeout(timeout time.Duration) *PutOrdersOrderFidProductsOrderProductFidSetInitialTermParams {
	return &PutOrdersOrderFidProductsOrderProductFidSetInitialTermParams{
		timeout: timeout,
	}
}

// NewPutOrdersOrderFidProductsOrderProductFidSetInitialTermParamsWithContext creates a new PutOrdersOrderFidProductsOrderProductFidSetInitialTermParams object
// with the ability to set a context for a request.
func NewPutOrdersOrderFidProductsOrderProductFidSetInitialTermParamsWithContext(ctx context.Context) *PutOrdersOrderFidProductsOrderProductFidSetInitialTermParams {
	return &PutOrdersOrderFidProductsOrderProductFidSetInitialTermParams{
		Context: ctx,
	}
}

// NewPutOrdersOrderFidProductsOrderProductFidSetInitialTermParamsWithHTTPClient creates a new PutOrdersOrderFidProductsOrderProductFidSetInitialTermParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutOrdersOrderFidProductsOrderProductFidSetInitialTermParamsWithHTTPClient(client *http.Client) *PutOrdersOrderFidProductsOrderProductFidSetInitialTermParams {
	return &PutOrdersOrderFidProductsOrderProductFidSetInitialTermParams{
		HTTPClient: client,
	}
}

/*
PutOrdersOrderFidProductsOrderProductFidSetInitialTermParams contains all the parameters to send to the API endpoint

	for the put orders order fid products order product fid set initial term operation.

	Typically these are written to a http.Request.
*/
type PutOrdersOrderFidProductsOrderProductFidSetInitialTermParams struct {

	// EndDate.
	EndDate *string

	// OrderFid.
	OrderFid string

	// OrderProductFid.
	OrderProductFid string

	// StartDate.
	StartDate *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put orders order fid products order product fid set initial term params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutOrdersOrderFidProductsOrderProductFidSetInitialTermParams) WithDefaults() *PutOrdersOrderFidProductsOrderProductFidSetInitialTermParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put orders order fid products order product fid set initial term params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutOrdersOrderFidProductsOrderProductFidSetInitialTermParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put orders order fid products order product fid set initial term params
func (o *PutOrdersOrderFidProductsOrderProductFidSetInitialTermParams) WithTimeout(timeout time.Duration) *PutOrdersOrderFidProductsOrderProductFidSetInitialTermParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put orders order fid products order product fid set initial term params
func (o *PutOrdersOrderFidProductsOrderProductFidSetInitialTermParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put orders order fid products order product fid set initial term params
func (o *PutOrdersOrderFidProductsOrderProductFidSetInitialTermParams) WithContext(ctx context.Context) *PutOrdersOrderFidProductsOrderProductFidSetInitialTermParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put orders order fid products order product fid set initial term params
func (o *PutOrdersOrderFidProductsOrderProductFidSetInitialTermParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put orders order fid products order product fid set initial term params
func (o *PutOrdersOrderFidProductsOrderProductFidSetInitialTermParams) WithHTTPClient(client *http.Client) *PutOrdersOrderFidProductsOrderProductFidSetInitialTermParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put orders order fid products order product fid set initial term params
func (o *PutOrdersOrderFidProductsOrderProductFidSetInitialTermParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndDate adds the endDate to the put orders order fid products order product fid set initial term params
func (o *PutOrdersOrderFidProductsOrderProductFidSetInitialTermParams) WithEndDate(endDate *string) *PutOrdersOrderFidProductsOrderProductFidSetInitialTermParams {
	o.SetEndDate(endDate)
	return o
}

// SetEndDate adds the endDate to the put orders order fid products order product fid set initial term params
func (o *PutOrdersOrderFidProductsOrderProductFidSetInitialTermParams) SetEndDate(endDate *string) {
	o.EndDate = endDate
}

// WithOrderFid adds the orderFid to the put orders order fid products order product fid set initial term params
func (o *PutOrdersOrderFidProductsOrderProductFidSetInitialTermParams) WithOrderFid(orderFid string) *PutOrdersOrderFidProductsOrderProductFidSetInitialTermParams {
	o.SetOrderFid(orderFid)
	return o
}

// SetOrderFid adds the orderFid to the put orders order fid products order product fid set initial term params
func (o *PutOrdersOrderFidProductsOrderProductFidSetInitialTermParams) SetOrderFid(orderFid string) {
	o.OrderFid = orderFid
}

// WithOrderProductFid adds the orderProductFid to the put orders order fid products order product fid set initial term params
func (o *PutOrdersOrderFidProductsOrderProductFidSetInitialTermParams) WithOrderProductFid(orderProductFid string) *PutOrdersOrderFidProductsOrderProductFidSetInitialTermParams {
	o.SetOrderProductFid(orderProductFid)
	return o
}

// SetOrderProductFid adds the orderProductFid to the put orders order fid products order product fid set initial term params
func (o *PutOrdersOrderFidProductsOrderProductFidSetInitialTermParams) SetOrderProductFid(orderProductFid string) {
	o.OrderProductFid = orderProductFid
}

// WithStartDate adds the startDate to the put orders order fid products order product fid set initial term params
func (o *PutOrdersOrderFidProductsOrderProductFidSetInitialTermParams) WithStartDate(startDate *string) *PutOrdersOrderFidProductsOrderProductFidSetInitialTermParams {
	o.SetStartDate(startDate)
	return o
}

// SetStartDate adds the startDate to the put orders order fid products order product fid set initial term params
func (o *PutOrdersOrderFidProductsOrderProductFidSetInitialTermParams) SetStartDate(startDate *string) {
	o.StartDate = startDate
}

// WriteToRequest writes these params to a swagger request
func (o *PutOrdersOrderFidProductsOrderProductFidSetInitialTermParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.EndDate != nil {

		// form param endDate
		var frEndDate string
		if o.EndDate != nil {
			frEndDate = *o.EndDate
		}
		fEndDate := frEndDate
		if fEndDate != "" {
			if err := r.SetFormParam("endDate", fEndDate); err != nil {
				return err
			}
		}
	}

	// path param orderFid
	if err := r.SetPathParam("orderFid", o.OrderFid); err != nil {
		return err
	}

	// path param orderProductFid
	if err := r.SetPathParam("orderProductFid", o.OrderProductFid); err != nil {
		return err
	}

	if o.StartDate != nil {

		// form param startDate
		var frStartDate string
		if o.StartDate != nil {
			frStartDate = *o.StartDate
		}
		fStartDate := frStartDate
		if fStartDate != "" {
			if err := r.SetFormParam("startDate", fStartDate); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
