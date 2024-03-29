// Code generated by go-swagger; DO NOT EDIT.

package interactions

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

// NewPostInteractionsInteractionFidCustomerClosedParams creates a new PostInteractionsInteractionFidCustomerClosedParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostInteractionsInteractionFidCustomerClosedParams() *PostInteractionsInteractionFidCustomerClosedParams {
	return &PostInteractionsInteractionFidCustomerClosedParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostInteractionsInteractionFidCustomerClosedParamsWithTimeout creates a new PostInteractionsInteractionFidCustomerClosedParams object
// with the ability to set a timeout on a request.
func NewPostInteractionsInteractionFidCustomerClosedParamsWithTimeout(timeout time.Duration) *PostInteractionsInteractionFidCustomerClosedParams {
	return &PostInteractionsInteractionFidCustomerClosedParams{
		timeout: timeout,
	}
}

// NewPostInteractionsInteractionFidCustomerClosedParamsWithContext creates a new PostInteractionsInteractionFidCustomerClosedParams object
// with the ability to set a context for a request.
func NewPostInteractionsInteractionFidCustomerClosedParamsWithContext(ctx context.Context) *PostInteractionsInteractionFidCustomerClosedParams {
	return &PostInteractionsInteractionFidCustomerClosedParams{
		Context: ctx,
	}
}

// NewPostInteractionsInteractionFidCustomerClosedParamsWithHTTPClient creates a new PostInteractionsInteractionFidCustomerClosedParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostInteractionsInteractionFidCustomerClosedParamsWithHTTPClient(client *http.Client) *PostInteractionsInteractionFidCustomerClosedParams {
	return &PostInteractionsInteractionFidCustomerClosedParams{
		HTTPClient: client,
	}
}

/*
PostInteractionsInteractionFidCustomerClosedParams contains all the parameters to send to the API endpoint

	for the post interactions interaction fid customer closed operation.

	Typically these are written to a http.Request.
*/
type PostInteractionsInteractionFidCustomerClosedParams struct {

	// InteractionFid.
	InteractionFid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post interactions interaction fid customer closed params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostInteractionsInteractionFidCustomerClosedParams) WithDefaults() *PostInteractionsInteractionFidCustomerClosedParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post interactions interaction fid customer closed params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostInteractionsInteractionFidCustomerClosedParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post interactions interaction fid customer closed params
func (o *PostInteractionsInteractionFidCustomerClosedParams) WithTimeout(timeout time.Duration) *PostInteractionsInteractionFidCustomerClosedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post interactions interaction fid customer closed params
func (o *PostInteractionsInteractionFidCustomerClosedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post interactions interaction fid customer closed params
func (o *PostInteractionsInteractionFidCustomerClosedParams) WithContext(ctx context.Context) *PostInteractionsInteractionFidCustomerClosedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post interactions interaction fid customer closed params
func (o *PostInteractionsInteractionFidCustomerClosedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post interactions interaction fid customer closed params
func (o *PostInteractionsInteractionFidCustomerClosedParams) WithHTTPClient(client *http.Client) *PostInteractionsInteractionFidCustomerClosedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post interactions interaction fid customer closed params
func (o *PostInteractionsInteractionFidCustomerClosedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInteractionFid adds the interactionFid to the post interactions interaction fid customer closed params
func (o *PostInteractionsInteractionFidCustomerClosedParams) WithInteractionFid(interactionFid string) *PostInteractionsInteractionFidCustomerClosedParams {
	o.SetInteractionFid(interactionFid)
	return o
}

// SetInteractionFid adds the interactionFid to the post interactions interaction fid customer closed params
func (o *PostInteractionsInteractionFidCustomerClosedParams) SetInteractionFid(interactionFid string) {
	o.InteractionFid = interactionFid
}

// WriteToRequest writes these params to a swagger request
func (o *PostInteractionsInteractionFidCustomerClosedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param interactionFid
	if err := r.SetPathParam("interactionFid", o.InteractionFid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
