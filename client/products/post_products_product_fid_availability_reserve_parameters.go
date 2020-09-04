// Code generated by go-swagger; DO NOT EDIT.

package products

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

// NewPostProductsProductFidAvailabilityReserveParams creates a new PostProductsProductFidAvailabilityReserveParams object
// with the default values initialized.
func NewPostProductsProductFidAvailabilityReserveParams() *PostProductsProductFidAvailabilityReserveParams {
	var ()
	return &PostProductsProductFidAvailabilityReserveParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostProductsProductFidAvailabilityReserveParamsWithTimeout creates a new PostProductsProductFidAvailabilityReserveParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostProductsProductFidAvailabilityReserveParamsWithTimeout(timeout time.Duration) *PostProductsProductFidAvailabilityReserveParams {
	var ()
	return &PostProductsProductFidAvailabilityReserveParams{

		timeout: timeout,
	}
}

// NewPostProductsProductFidAvailabilityReserveParamsWithContext creates a new PostProductsProductFidAvailabilityReserveParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostProductsProductFidAvailabilityReserveParamsWithContext(ctx context.Context) *PostProductsProductFidAvailabilityReserveParams {
	var ()
	return &PostProductsProductFidAvailabilityReserveParams{

		Context: ctx,
	}
}

// NewPostProductsProductFidAvailabilityReserveParamsWithHTTPClient creates a new PostProductsProductFidAvailabilityReserveParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostProductsProductFidAvailabilityReserveParamsWithHTTPClient(client *http.Client) *PostProductsProductFidAvailabilityReserveParams {
	var ()
	return &PostProductsProductFidAvailabilityReserveParams{
		HTTPClient: client,
	}
}

/*PostProductsProductFidAvailabilityReserveParams contains all the parameters to send to the API endpoint
for the post products product fid availability reserve operation typically these are written to a http.Request
*/
type PostProductsProductFidAvailabilityReserveParams struct {

	/*Payload*/
	Payload *models.AvailabilityReservePayload
	/*ProductFid*/
	ProductFid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post products product fid availability reserve params
func (o *PostProductsProductFidAvailabilityReserveParams) WithTimeout(timeout time.Duration) *PostProductsProductFidAvailabilityReserveParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post products product fid availability reserve params
func (o *PostProductsProductFidAvailabilityReserveParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post products product fid availability reserve params
func (o *PostProductsProductFidAvailabilityReserveParams) WithContext(ctx context.Context) *PostProductsProductFidAvailabilityReserveParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post products product fid availability reserve params
func (o *PostProductsProductFidAvailabilityReserveParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post products product fid availability reserve params
func (o *PostProductsProductFidAvailabilityReserveParams) WithHTTPClient(client *http.Client) *PostProductsProductFidAvailabilityReserveParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post products product fid availability reserve params
func (o *PostProductsProductFidAvailabilityReserveParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPayload adds the payload to the post products product fid availability reserve params
func (o *PostProductsProductFidAvailabilityReserveParams) WithPayload(payload *models.AvailabilityReservePayload) *PostProductsProductFidAvailabilityReserveParams {
	o.SetPayload(payload)
	return o
}

// SetPayload adds the payload to the post products product fid availability reserve params
func (o *PostProductsProductFidAvailabilityReserveParams) SetPayload(payload *models.AvailabilityReservePayload) {
	o.Payload = payload
}

// WithProductFid adds the productFid to the post products product fid availability reserve params
func (o *PostProductsProductFidAvailabilityReserveParams) WithProductFid(productFid string) *PostProductsProductFidAvailabilityReserveParams {
	o.SetProductFid(productFid)
	return o
}

// SetProductFid adds the productFid to the post products product fid availability reserve params
func (o *PostProductsProductFidAvailabilityReserveParams) SetProductFid(productFid string) {
	o.ProductFid = productFid
}

// WriteToRequest writes these params to a swagger request
func (o *PostProductsProductFidAvailabilityReserveParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Payload != nil {
		if err := r.SetBodyParam(o.Payload); err != nil {
			return err
		}
	}

	// path param productFid
	if err := r.SetPathParam("productFid", o.ProductFid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
