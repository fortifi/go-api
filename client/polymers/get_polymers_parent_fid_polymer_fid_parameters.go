// Code generated by go-swagger; DO NOT EDIT.

package polymers

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

// NewGetPolymersParentFidPolymerFidParams creates a new GetPolymersParentFidPolymerFidParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPolymersParentFidPolymerFidParams() *GetPolymersParentFidPolymerFidParams {
	return &GetPolymersParentFidPolymerFidParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPolymersParentFidPolymerFidParamsWithTimeout creates a new GetPolymersParentFidPolymerFidParams object
// with the ability to set a timeout on a request.
func NewGetPolymersParentFidPolymerFidParamsWithTimeout(timeout time.Duration) *GetPolymersParentFidPolymerFidParams {
	return &GetPolymersParentFidPolymerFidParams{
		timeout: timeout,
	}
}

// NewGetPolymersParentFidPolymerFidParamsWithContext creates a new GetPolymersParentFidPolymerFidParams object
// with the ability to set a context for a request.
func NewGetPolymersParentFidPolymerFidParamsWithContext(ctx context.Context) *GetPolymersParentFidPolymerFidParams {
	return &GetPolymersParentFidPolymerFidParams{
		Context: ctx,
	}
}

// NewGetPolymersParentFidPolymerFidParamsWithHTTPClient creates a new GetPolymersParentFidPolymerFidParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPolymersParentFidPolymerFidParamsWithHTTPClient(client *http.Client) *GetPolymersParentFidPolymerFidParams {
	return &GetPolymersParentFidPolymerFidParams{
		HTTPClient: client,
	}
}

/*
GetPolymersParentFidPolymerFidParams contains all the parameters to send to the API endpoint

	for the get polymers parent fid polymer fid operation.

	Typically these are written to a http.Request.
*/
type GetPolymersParentFidPolymerFidParams struct {

	// AllData.
	AllData bool

	// DataKeys.
	DataKeys []string

	// ParentFid.
	ParentFid string

	// PolymerFid.
	PolymerFid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get polymers parent fid polymer fid params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPolymersParentFidPolymerFidParams) WithDefaults() *GetPolymersParentFidPolymerFidParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get polymers parent fid polymer fid params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPolymersParentFidPolymerFidParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get polymers parent fid polymer fid params
func (o *GetPolymersParentFidPolymerFidParams) WithTimeout(timeout time.Duration) *GetPolymersParentFidPolymerFidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get polymers parent fid polymer fid params
func (o *GetPolymersParentFidPolymerFidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get polymers parent fid polymer fid params
func (o *GetPolymersParentFidPolymerFidParams) WithContext(ctx context.Context) *GetPolymersParentFidPolymerFidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get polymers parent fid polymer fid params
func (o *GetPolymersParentFidPolymerFidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get polymers parent fid polymer fid params
func (o *GetPolymersParentFidPolymerFidParams) WithHTTPClient(client *http.Client) *GetPolymersParentFidPolymerFidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get polymers parent fid polymer fid params
func (o *GetPolymersParentFidPolymerFidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAllData adds the allData to the get polymers parent fid polymer fid params
func (o *GetPolymersParentFidPolymerFidParams) WithAllData(allData bool) *GetPolymersParentFidPolymerFidParams {
	o.SetAllData(allData)
	return o
}

// SetAllData adds the allData to the get polymers parent fid polymer fid params
func (o *GetPolymersParentFidPolymerFidParams) SetAllData(allData bool) {
	o.AllData = allData
}

// WithDataKeys adds the dataKeys to the get polymers parent fid polymer fid params
func (o *GetPolymersParentFidPolymerFidParams) WithDataKeys(dataKeys []string) *GetPolymersParentFidPolymerFidParams {
	o.SetDataKeys(dataKeys)
	return o
}

// SetDataKeys adds the dataKeys to the get polymers parent fid polymer fid params
func (o *GetPolymersParentFidPolymerFidParams) SetDataKeys(dataKeys []string) {
	o.DataKeys = dataKeys
}

// WithParentFid adds the parentFid to the get polymers parent fid polymer fid params
func (o *GetPolymersParentFidPolymerFidParams) WithParentFid(parentFid string) *GetPolymersParentFidPolymerFidParams {
	o.SetParentFid(parentFid)
	return o
}

// SetParentFid adds the parentFid to the get polymers parent fid polymer fid params
func (o *GetPolymersParentFidPolymerFidParams) SetParentFid(parentFid string) {
	o.ParentFid = parentFid
}

// WithPolymerFid adds the polymerFid to the get polymers parent fid polymer fid params
func (o *GetPolymersParentFidPolymerFidParams) WithPolymerFid(polymerFid string) *GetPolymersParentFidPolymerFidParams {
	o.SetPolymerFid(polymerFid)
	return o
}

// SetPolymerFid adds the polymerFid to the get polymers parent fid polymer fid params
func (o *GetPolymersParentFidPolymerFidParams) SetPolymerFid(polymerFid string) {
	o.PolymerFid = polymerFid
}

// WriteToRequest writes these params to a swagger request
func (o *GetPolymersParentFidPolymerFidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param allData
	qrAllData := o.AllData
	qAllData := swag.FormatBool(qrAllData)
	if qAllData != "" {

		if err := r.SetQueryParam("allData", qAllData); err != nil {
			return err
		}
	}

	if o.DataKeys != nil {

		// binding items for dataKeys
		joinedDataKeys := o.bindParamDataKeys(reg)

		// query array param dataKeys
		if err := r.SetQueryParam("dataKeys", joinedDataKeys...); err != nil {
			return err
		}
	}

	// path param parentFid
	if err := r.SetPathParam("parentFid", o.ParentFid); err != nil {
		return err
	}

	// path param polymerFid
	if err := r.SetPathParam("polymerFid", o.PolymerFid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetPolymersParentFidPolymerFid binds the parameter dataKeys
func (o *GetPolymersParentFidPolymerFidParams) bindParamDataKeys(formats strfmt.Registry) []string {
	dataKeysIR := o.DataKeys

	var dataKeysIC []string
	for _, dataKeysIIR := range dataKeysIR { // explode []string

		dataKeysIIV := dataKeysIIR // string as string
		dataKeysIC = append(dataKeysIC, dataKeysIIV)
	}

	// items.CollectionFormat: ""
	dataKeysIS := swag.JoinByFormat(dataKeysIC, "")

	return dataKeysIS
}
