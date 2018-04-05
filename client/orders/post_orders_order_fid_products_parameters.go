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

	"github.com/fortifi/go-api/models"
)

// NewPostOrdersOrderFidProductsParams creates a new PostOrdersOrderFidProductsParams object
// with the default values initialized.
func NewPostOrdersOrderFidProductsParams() *PostOrdersOrderFidProductsParams {
	var ()
	return &PostOrdersOrderFidProductsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostOrdersOrderFidProductsParamsWithTimeout creates a new PostOrdersOrderFidProductsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostOrdersOrderFidProductsParamsWithTimeout(timeout time.Duration) *PostOrdersOrderFidProductsParams {
	var ()
	return &PostOrdersOrderFidProductsParams{

		timeout: timeout,
	}
}

// NewPostOrdersOrderFidProductsParamsWithContext creates a new PostOrdersOrderFidProductsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostOrdersOrderFidProductsParamsWithContext(ctx context.Context) *PostOrdersOrderFidProductsParams {
	var ()
	return &PostOrdersOrderFidProductsParams{

		Context: ctx,
	}
}

// NewPostOrdersOrderFidProductsParamsWithHTTPClient creates a new PostOrdersOrderFidProductsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostOrdersOrderFidProductsParamsWithHTTPClient(client *http.Client) *PostOrdersOrderFidProductsParams {
	var ()
	return &PostOrdersOrderFidProductsParams{
		HTTPClient: client,
	}
}

/*PostOrdersOrderFidProductsParams contains all the parameters to send to the API endpoint
for the post orders order fid products operation typically these are written to a http.Request
*/
type PostOrdersOrderFidProductsParams struct {

	/*OrderFid*/
	OrderFid string
	/*Payload*/
	Payload *models.AddOrderProductsPayload

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post orders order fid products params
func (o *PostOrdersOrderFidProductsParams) WithTimeout(timeout time.Duration) *PostOrdersOrderFidProductsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post orders order fid products params
func (o *PostOrdersOrderFidProductsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post orders order fid products params
func (o *PostOrdersOrderFidProductsParams) WithContext(ctx context.Context) *PostOrdersOrderFidProductsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post orders order fid products params
func (o *PostOrdersOrderFidProductsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post orders order fid products params
func (o *PostOrdersOrderFidProductsParams) WithHTTPClient(client *http.Client) *PostOrdersOrderFidProductsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post orders order fid products params
func (o *PostOrdersOrderFidProductsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrderFid adds the orderFid to the post orders order fid products params
func (o *PostOrdersOrderFidProductsParams) WithOrderFid(orderFid string) *PostOrdersOrderFidProductsParams {
	o.SetOrderFid(orderFid)
	return o
}

// SetOrderFid adds the orderFid to the post orders order fid products params
func (o *PostOrdersOrderFidProductsParams) SetOrderFid(orderFid string) {
	o.OrderFid = orderFid
}

// WithPayload adds the payload to the post orders order fid products params
func (o *PostOrdersOrderFidProductsParams) WithPayload(payload *models.AddOrderProductsPayload) *PostOrdersOrderFidProductsParams {
	o.SetPayload(payload)
	return o
}

// SetPayload adds the payload to the post orders order fid products params
func (o *PostOrdersOrderFidProductsParams) SetPayload(payload *models.AddOrderProductsPayload) {
	o.Payload = payload
}

// WriteToRequest writes these params to a swagger request
func (o *PostOrdersOrderFidProductsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
