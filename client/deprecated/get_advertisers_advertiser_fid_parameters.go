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

// NewGetAdvertisersAdvertiserFidParams creates a new GetAdvertisersAdvertiserFidParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAdvertisersAdvertiserFidParams() *GetAdvertisersAdvertiserFidParams {
	return &GetAdvertisersAdvertiserFidParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAdvertisersAdvertiserFidParamsWithTimeout creates a new GetAdvertisersAdvertiserFidParams object
// with the ability to set a timeout on a request.
func NewGetAdvertisersAdvertiserFidParamsWithTimeout(timeout time.Duration) *GetAdvertisersAdvertiserFidParams {
	return &GetAdvertisersAdvertiserFidParams{
		timeout: timeout,
	}
}

// NewGetAdvertisersAdvertiserFidParamsWithContext creates a new GetAdvertisersAdvertiserFidParams object
// with the ability to set a context for a request.
func NewGetAdvertisersAdvertiserFidParamsWithContext(ctx context.Context) *GetAdvertisersAdvertiserFidParams {
	return &GetAdvertisersAdvertiserFidParams{
		Context: ctx,
	}
}

// NewGetAdvertisersAdvertiserFidParamsWithHTTPClient creates a new GetAdvertisersAdvertiserFidParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAdvertisersAdvertiserFidParamsWithHTTPClient(client *http.Client) *GetAdvertisersAdvertiserFidParams {
	return &GetAdvertisersAdvertiserFidParams{
		HTTPClient: client,
	}
}

/*
GetAdvertisersAdvertiserFidParams contains all the parameters to send to the API endpoint

	for the get advertisers advertiser fid operation.

	Typically these are written to a http.Request.
*/
type GetAdvertisersAdvertiserFidParams struct {

	/* AdvertiserFid.

	   Advertiser FID to use
	*/
	AdvertiserFid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get advertisers advertiser fid params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAdvertisersAdvertiserFidParams) WithDefaults() *GetAdvertisersAdvertiserFidParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get advertisers advertiser fid params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAdvertisersAdvertiserFidParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get advertisers advertiser fid params
func (o *GetAdvertisersAdvertiserFidParams) WithTimeout(timeout time.Duration) *GetAdvertisersAdvertiserFidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get advertisers advertiser fid params
func (o *GetAdvertisersAdvertiserFidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get advertisers advertiser fid params
func (o *GetAdvertisersAdvertiserFidParams) WithContext(ctx context.Context) *GetAdvertisersAdvertiserFidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get advertisers advertiser fid params
func (o *GetAdvertisersAdvertiserFidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get advertisers advertiser fid params
func (o *GetAdvertisersAdvertiserFidParams) WithHTTPClient(client *http.Client) *GetAdvertisersAdvertiserFidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get advertisers advertiser fid params
func (o *GetAdvertisersAdvertiserFidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAdvertiserFid adds the advertiserFid to the get advertisers advertiser fid params
func (o *GetAdvertisersAdvertiserFidParams) WithAdvertiserFid(advertiserFid string) *GetAdvertisersAdvertiserFidParams {
	o.SetAdvertiserFid(advertiserFid)
	return o
}

// SetAdvertiserFid adds the advertiserFid to the get advertisers advertiser fid params
func (o *GetAdvertisersAdvertiserFidParams) SetAdvertiserFid(advertiserFid string) {
	o.AdvertiserFid = advertiserFid
}

// WriteToRequest writes these params to a swagger request
func (o *GetAdvertisersAdvertiserFidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
