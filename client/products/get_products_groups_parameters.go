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
)

// NewGetProductsGroupsParams creates a new GetProductsGroupsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetProductsGroupsParams() *GetProductsGroupsParams {
	return &GetProductsGroupsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetProductsGroupsParamsWithTimeout creates a new GetProductsGroupsParams object
// with the ability to set a timeout on a request.
func NewGetProductsGroupsParamsWithTimeout(timeout time.Duration) *GetProductsGroupsParams {
	return &GetProductsGroupsParams{
		timeout: timeout,
	}
}

// NewGetProductsGroupsParamsWithContext creates a new GetProductsGroupsParams object
// with the ability to set a context for a request.
func NewGetProductsGroupsParamsWithContext(ctx context.Context) *GetProductsGroupsParams {
	return &GetProductsGroupsParams{
		Context: ctx,
	}
}

// NewGetProductsGroupsParamsWithHTTPClient creates a new GetProductsGroupsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetProductsGroupsParamsWithHTTPClient(client *http.Client) *GetProductsGroupsParams {
	return &GetProductsGroupsParams{
		HTTPClient: client,
	}
}

/*
GetProductsGroupsParams contains all the parameters to send to the API endpoint

	for the get products groups operation.

	Typically these are written to a http.Request.
*/
type GetProductsGroupsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get products groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProductsGroupsParams) WithDefaults() *GetProductsGroupsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get products groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProductsGroupsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get products groups params
func (o *GetProductsGroupsParams) WithTimeout(timeout time.Duration) *GetProductsGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get products groups params
func (o *GetProductsGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get products groups params
func (o *GetProductsGroupsParams) WithContext(ctx context.Context) *GetProductsGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get products groups params
func (o *GetProductsGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get products groups params
func (o *GetProductsGroupsParams) WithHTTPClient(client *http.Client) *GetProductsGroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get products groups params
func (o *GetProductsGroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetProductsGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
