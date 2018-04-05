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

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetOrdersOrderFidParams creates a new GetOrdersOrderFidParams object
// with the default values initialized.
func NewGetOrdersOrderFidParams() *GetOrdersOrderFidParams {
	var ()
	return &GetOrdersOrderFidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrdersOrderFidParamsWithTimeout creates a new GetOrdersOrderFidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOrdersOrderFidParamsWithTimeout(timeout time.Duration) *GetOrdersOrderFidParams {
	var ()
	return &GetOrdersOrderFidParams{

		timeout: timeout,
	}
}

// NewGetOrdersOrderFidParamsWithContext creates a new GetOrdersOrderFidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOrdersOrderFidParamsWithContext(ctx context.Context) *GetOrdersOrderFidParams {
	var ()
	return &GetOrdersOrderFidParams{

		Context: ctx,
	}
}

// NewGetOrdersOrderFidParamsWithHTTPClient creates a new GetOrdersOrderFidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOrdersOrderFidParamsWithHTTPClient(client *http.Client) *GetOrdersOrderFidParams {
	var ()
	return &GetOrdersOrderFidParams{
		HTTPClient: client,
	}
}

/*GetOrdersOrderFidParams contains all the parameters to send to the API endpoint
for the get orders order fid operation typically these are written to a http.Request
*/
type GetOrdersOrderFidParams struct {

	/*OrderFid*/
	OrderFid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get orders order fid params
func (o *GetOrdersOrderFidParams) WithTimeout(timeout time.Duration) *GetOrdersOrderFidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get orders order fid params
func (o *GetOrdersOrderFidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get orders order fid params
func (o *GetOrdersOrderFidParams) WithContext(ctx context.Context) *GetOrdersOrderFidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get orders order fid params
func (o *GetOrdersOrderFidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get orders order fid params
func (o *GetOrdersOrderFidParams) WithHTTPClient(client *http.Client) *GetOrdersOrderFidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get orders order fid params
func (o *GetOrdersOrderFidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrderFid adds the orderFid to the get orders order fid params
func (o *GetOrdersOrderFidParams) WithOrderFid(orderFid string) *GetOrdersOrderFidParams {
	o.SetOrderFid(orderFid)
	return o
}

// SetOrderFid adds the orderFid to the get orders order fid params
func (o *GetOrdersOrderFidParams) SetOrderFid(orderFid string) {
	o.OrderFid = orderFid
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrdersOrderFidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
