// Code generated by go-swagger; DO NOT EDIT.

package deprecated

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

// NewDeleteAdvertisersAdvertiserFidApprovedParams creates a new DeleteAdvertisersAdvertiserFidApprovedParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteAdvertisersAdvertiserFidApprovedParams() *DeleteAdvertisersAdvertiserFidApprovedParams {
	return &DeleteAdvertisersAdvertiserFidApprovedParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAdvertisersAdvertiserFidApprovedParamsWithTimeout creates a new DeleteAdvertisersAdvertiserFidApprovedParams object
// with the ability to set a timeout on a request.
func NewDeleteAdvertisersAdvertiserFidApprovedParamsWithTimeout(timeout time.Duration) *DeleteAdvertisersAdvertiserFidApprovedParams {
	return &DeleteAdvertisersAdvertiserFidApprovedParams{
		timeout: timeout,
	}
}

// NewDeleteAdvertisersAdvertiserFidApprovedParamsWithContext creates a new DeleteAdvertisersAdvertiserFidApprovedParams object
// with the ability to set a context for a request.
func NewDeleteAdvertisersAdvertiserFidApprovedParamsWithContext(ctx context.Context) *DeleteAdvertisersAdvertiserFidApprovedParams {
	return &DeleteAdvertisersAdvertiserFidApprovedParams{
		Context: ctx,
	}
}

// NewDeleteAdvertisersAdvertiserFidApprovedParamsWithHTTPClient creates a new DeleteAdvertisersAdvertiserFidApprovedParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteAdvertisersAdvertiserFidApprovedParamsWithHTTPClient(client *http.Client) *DeleteAdvertisersAdvertiserFidApprovedParams {
	return &DeleteAdvertisersAdvertiserFidApprovedParams{
		HTTPClient: client,
	}
}

/*
DeleteAdvertisersAdvertiserFidApprovedParams contains all the parameters to send to the API endpoint

	for the delete advertisers advertiser fid approved operation.

	Typically these are written to a http.Request.
*/
type DeleteAdvertisersAdvertiserFidApprovedParams struct {

	/* AdvertiserFid.

	   Advertiser FID to use
	*/
	AdvertiserFid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete advertisers advertiser fid approved params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAdvertisersAdvertiserFidApprovedParams) WithDefaults() *DeleteAdvertisersAdvertiserFidApprovedParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete advertisers advertiser fid approved params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAdvertisersAdvertiserFidApprovedParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete advertisers advertiser fid approved params
func (o *DeleteAdvertisersAdvertiserFidApprovedParams) WithTimeout(timeout time.Duration) *DeleteAdvertisersAdvertiserFidApprovedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete advertisers advertiser fid approved params
func (o *DeleteAdvertisersAdvertiserFidApprovedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete advertisers advertiser fid approved params
func (o *DeleteAdvertisersAdvertiserFidApprovedParams) WithContext(ctx context.Context) *DeleteAdvertisersAdvertiserFidApprovedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete advertisers advertiser fid approved params
func (o *DeleteAdvertisersAdvertiserFidApprovedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete advertisers advertiser fid approved params
func (o *DeleteAdvertisersAdvertiserFidApprovedParams) WithHTTPClient(client *http.Client) *DeleteAdvertisersAdvertiserFidApprovedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete advertisers advertiser fid approved params
func (o *DeleteAdvertisersAdvertiserFidApprovedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAdvertiserFid adds the advertiserFid to the delete advertisers advertiser fid approved params
func (o *DeleteAdvertisersAdvertiserFidApprovedParams) WithAdvertiserFid(advertiserFid string) *DeleteAdvertisersAdvertiserFidApprovedParams {
	o.SetAdvertiserFid(advertiserFid)
	return o
}

// SetAdvertiserFid adds the advertiserFid to the delete advertisers advertiser fid approved params
func (o *DeleteAdvertisersAdvertiserFidApprovedParams) SetAdvertiserFid(advertiserFid string) {
	o.AdvertiserFid = advertiserFid
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAdvertisersAdvertiserFidApprovedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param advertiserFid
	if err := r.SetPathParam("advertiserFid", o.AdvertiserFid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
