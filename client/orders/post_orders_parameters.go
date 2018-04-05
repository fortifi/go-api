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

// NewPostOrdersParams creates a new PostOrdersParams object
// with the default values initialized.
func NewPostOrdersParams() *PostOrdersParams {
	var ()
	return &PostOrdersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostOrdersParamsWithTimeout creates a new PostOrdersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostOrdersParamsWithTimeout(timeout time.Duration) *PostOrdersParams {
	var ()
	return &PostOrdersParams{

		timeout: timeout,
	}
}

// NewPostOrdersParamsWithContext creates a new PostOrdersParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostOrdersParamsWithContext(ctx context.Context) *PostOrdersParams {
	var ()
	return &PostOrdersParams{

		Context: ctx,
	}
}

// NewPostOrdersParamsWithHTTPClient creates a new PostOrdersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostOrdersParamsWithHTTPClient(client *http.Client) *PostOrdersParams {
	var ()
	return &PostOrdersParams{
		HTTPClient: client,
	}
}

/*PostOrdersParams contains all the parameters to send to the API endpoint
for the post orders operation typically these are written to a http.Request
*/
type PostOrdersParams struct {

	/*Payload*/
	Payload *models.CreateOrderPayload

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post orders params
func (o *PostOrdersParams) WithTimeout(timeout time.Duration) *PostOrdersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post orders params
func (o *PostOrdersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post orders params
func (o *PostOrdersParams) WithContext(ctx context.Context) *PostOrdersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post orders params
func (o *PostOrdersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post orders params
func (o *PostOrdersParams) WithHTTPClient(client *http.Client) *PostOrdersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post orders params
func (o *PostOrdersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPayload adds the payload to the post orders params
func (o *PostOrdersParams) WithPayload(payload *models.CreateOrderPayload) *PostOrdersParams {
	o.SetPayload(payload)
	return o
}

// SetPayload adds the payload to the post orders params
func (o *PostOrdersParams) SetPayload(payload *models.CreateOrderPayload) {
	o.Payload = payload
}

// WriteToRequest writes these params to a swagger request
func (o *PostOrdersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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