// Code generated by go-swagger; DO NOT EDIT.

package marketing

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
	"github.com/go-openapi/swag"
)

// NewGetPublishersParams creates a new GetPublishersParams object
// with the default values initialized.
func NewGetPublishersParams() *GetPublishersParams {
	var ()
	return &GetPublishersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPublishersParamsWithTimeout creates a new GetPublishersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPublishersParamsWithTimeout(timeout time.Duration) *GetPublishersParams {
	var ()
	return &GetPublishersParams{

		timeout: timeout,
	}
}

// NewGetPublishersParamsWithContext creates a new GetPublishersParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPublishersParamsWithContext(ctx context.Context) *GetPublishersParams {
	var ()
	return &GetPublishersParams{

		Context: ctx,
	}
}

// NewGetPublishersParamsWithHTTPClient creates a new GetPublishersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPublishersParamsWithHTTPClient(client *http.Client) *GetPublishersParams {
	var ()
	return &GetPublishersParams{
		HTTPClient: client,
	}
}

/*GetPublishersParams contains all the parameters to send to the API endpoint
for the get publishers operation typically these are written to a http.Request
*/
type GetPublishersParams struct {

	/*Limit
	  Maximum number of records per page (default: 20)

	*/
	Limit *int64
	/*Page
	  Page number (default: 1)

	*/
	Page *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get publishers params
func (o *GetPublishersParams) WithTimeout(timeout time.Duration) *GetPublishersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get publishers params
func (o *GetPublishersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get publishers params
func (o *GetPublishersParams) WithContext(ctx context.Context) *GetPublishersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get publishers params
func (o *GetPublishersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get publishers params
func (o *GetPublishersParams) WithHTTPClient(client *http.Client) *GetPublishersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get publishers params
func (o *GetPublishersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the get publishers params
func (o *GetPublishersParams) WithLimit(limit *int64) *GetPublishersParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get publishers params
func (o *GetPublishersParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithPage adds the page to the get publishers params
func (o *GetPublishersParams) WithPage(page *int64) *GetPublishersParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get publishers params
func (o *GetPublishersParams) SetPage(page *int64) {
	o.Page = page
}

// WriteToRequest writes these params to a swagger request
func (o *GetPublishersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Page != nil {

		// query param page
		var qrPage int64
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
