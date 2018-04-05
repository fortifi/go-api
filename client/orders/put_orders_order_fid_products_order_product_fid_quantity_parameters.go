// Code generated by go-swagger; DO NOT EDIT.

package orders

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPutOrdersOrderFidProductsOrderProductFidQuantityParams creates a new PutOrdersOrderFidProductsOrderProductFidQuantityParams object
// with the default values initialized.
func NewPutOrdersOrderFidProductsOrderProductFidQuantityParams() *PutOrdersOrderFidProductsOrderProductFidQuantityParams {
	var ()
	return &PutOrdersOrderFidProductsOrderProductFidQuantityParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutOrdersOrderFidProductsOrderProductFidQuantityParamsWithTimeout creates a new PutOrdersOrderFidProductsOrderProductFidQuantityParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutOrdersOrderFidProductsOrderProductFidQuantityParamsWithTimeout(timeout time.Duration) *PutOrdersOrderFidProductsOrderProductFidQuantityParams {
	var ()
	return &PutOrdersOrderFidProductsOrderProductFidQuantityParams{

		timeout: timeout,
	}
}

// NewPutOrdersOrderFidProductsOrderProductFidQuantityParamsWithContext creates a new PutOrdersOrderFidProductsOrderProductFidQuantityParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutOrdersOrderFidProductsOrderProductFidQuantityParamsWithContext(ctx context.Context) *PutOrdersOrderFidProductsOrderProductFidQuantityParams {
	var ()
	return &PutOrdersOrderFidProductsOrderProductFidQuantityParams{

		Context: ctx,
	}
}

// NewPutOrdersOrderFidProductsOrderProductFidQuantityParamsWithHTTPClient creates a new PutOrdersOrderFidProductsOrderProductFidQuantityParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutOrdersOrderFidProductsOrderProductFidQuantityParamsWithHTTPClient(client *http.Client) *PutOrdersOrderFidProductsOrderProductFidQuantityParams {
	var ()
	return &PutOrdersOrderFidProductsOrderProductFidQuantityParams{
		HTTPClient: client,
	}
}

/*PutOrdersOrderFidProductsOrderProductFidQuantityParams contains all the parameters to send to the API endpoint
for the put orders order fid products order product fid quantity operation typically these are written to a http.Request
*/
type PutOrdersOrderFidProductsOrderProductFidQuantityParams struct {

	/*OrderFid*/
	OrderFid string
	/*OrderProductFid*/
	OrderProductFid string
	/*Quantity*/
	Quantity float64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put orders order fid products order product fid quantity params
func (o *PutOrdersOrderFidProductsOrderProductFidQuantityParams) WithTimeout(timeout time.Duration) *PutOrdersOrderFidProductsOrderProductFidQuantityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put orders order fid products order product fid quantity params
func (o *PutOrdersOrderFidProductsOrderProductFidQuantityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put orders order fid products order product fid quantity params
func (o *PutOrdersOrderFidProductsOrderProductFidQuantityParams) WithContext(ctx context.Context) *PutOrdersOrderFidProductsOrderProductFidQuantityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put orders order fid products order product fid quantity params
func (o *PutOrdersOrderFidProductsOrderProductFidQuantityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put orders order fid products order product fid quantity params
func (o *PutOrdersOrderFidProductsOrderProductFidQuantityParams) WithHTTPClient(client *http.Client) *PutOrdersOrderFidProductsOrderProductFidQuantityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put orders order fid products order product fid quantity params
func (o *PutOrdersOrderFidProductsOrderProductFidQuantityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrderFid adds the orderFid to the put orders order fid products order product fid quantity params
func (o *PutOrdersOrderFidProductsOrderProductFidQuantityParams) WithOrderFid(orderFid string) *PutOrdersOrderFidProductsOrderProductFidQuantityParams {
	o.SetOrderFid(orderFid)
	return o
}

// SetOrderFid adds the orderFid to the put orders order fid products order product fid quantity params
func (o *PutOrdersOrderFidProductsOrderProductFidQuantityParams) SetOrderFid(orderFid string) {
	o.OrderFid = orderFid
}

// WithOrderProductFid adds the orderProductFid to the put orders order fid products order product fid quantity params
func (o *PutOrdersOrderFidProductsOrderProductFidQuantityParams) WithOrderProductFid(orderProductFid string) *PutOrdersOrderFidProductsOrderProductFidQuantityParams {
	o.SetOrderProductFid(orderProductFid)
	return o
}

// SetOrderProductFid adds the orderProductFid to the put orders order fid products order product fid quantity params
func (o *PutOrdersOrderFidProductsOrderProductFidQuantityParams) SetOrderProductFid(orderProductFid string) {
	o.OrderProductFid = orderProductFid
}

// WithQuantity adds the quantity to the put orders order fid products order product fid quantity params
func (o *PutOrdersOrderFidProductsOrderProductFidQuantityParams) WithQuantity(quantity float64) *PutOrdersOrderFidProductsOrderProductFidQuantityParams {
	o.SetQuantity(quantity)
	return o
}

// SetQuantity adds the quantity to the put orders order fid products order product fid quantity params
func (o *PutOrdersOrderFidProductsOrderProductFidQuantityParams) SetQuantity(quantity float64) {
	o.Quantity = quantity
}

// WriteToRequest writes these params to a swagger request
func (o *PutOrdersOrderFidProductsOrderProductFidQuantityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param orderFid
	if err := r.SetPathParam("orderFid", o.OrderFid); err != nil {
		return err
	}

	// path param orderProductFid
	if err := r.SetPathParam("orderProductFid", o.OrderProductFid); err != nil {
		return err
	}

	// form param quantity
	frQuantity := o.Quantity
	fQuantity := swag.FormatFloat64(frQuantity)
	if fQuantity != "" {
		if err := r.SetFormParam("quantity", fQuantity); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}