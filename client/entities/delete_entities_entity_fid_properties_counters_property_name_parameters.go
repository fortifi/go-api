// Code generated by go-swagger; DO NOT EDIT.

package entities

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

// NewDeleteEntitiesEntityFidPropertiesCountersPropertyNameParams creates a new DeleteEntitiesEntityFidPropertiesCountersPropertyNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteEntitiesEntityFidPropertiesCountersPropertyNameParams() *DeleteEntitiesEntityFidPropertiesCountersPropertyNameParams {
	return &DeleteEntitiesEntityFidPropertiesCountersPropertyNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteEntitiesEntityFidPropertiesCountersPropertyNameParamsWithTimeout creates a new DeleteEntitiesEntityFidPropertiesCountersPropertyNameParams object
// with the ability to set a timeout on a request.
func NewDeleteEntitiesEntityFidPropertiesCountersPropertyNameParamsWithTimeout(timeout time.Duration) *DeleteEntitiesEntityFidPropertiesCountersPropertyNameParams {
	return &DeleteEntitiesEntityFidPropertiesCountersPropertyNameParams{
		timeout: timeout,
	}
}

// NewDeleteEntitiesEntityFidPropertiesCountersPropertyNameParamsWithContext creates a new DeleteEntitiesEntityFidPropertiesCountersPropertyNameParams object
// with the ability to set a context for a request.
func NewDeleteEntitiesEntityFidPropertiesCountersPropertyNameParamsWithContext(ctx context.Context) *DeleteEntitiesEntityFidPropertiesCountersPropertyNameParams {
	return &DeleteEntitiesEntityFidPropertiesCountersPropertyNameParams{
		Context: ctx,
	}
}

// NewDeleteEntitiesEntityFidPropertiesCountersPropertyNameParamsWithHTTPClient creates a new DeleteEntitiesEntityFidPropertiesCountersPropertyNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteEntitiesEntityFidPropertiesCountersPropertyNameParamsWithHTTPClient(client *http.Client) *DeleteEntitiesEntityFidPropertiesCountersPropertyNameParams {
	return &DeleteEntitiesEntityFidPropertiesCountersPropertyNameParams{
		HTTPClient: client,
	}
}

/*
DeleteEntitiesEntityFidPropertiesCountersPropertyNameParams contains all the parameters to send to the API endpoint

	for the delete entities entity fid properties counters property name operation.

	Typically these are written to a http.Request.
*/
type DeleteEntitiesEntityFidPropertiesCountersPropertyNameParams struct {

	/* EntityFid.

	   Entity FID to use
	*/
	EntityFid string

	/* PropertyName.

	   Property Name
	*/
	PropertyName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete entities entity fid properties counters property name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteEntitiesEntityFidPropertiesCountersPropertyNameParams) WithDefaults() *DeleteEntitiesEntityFidPropertiesCountersPropertyNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete entities entity fid properties counters property name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteEntitiesEntityFidPropertiesCountersPropertyNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete entities entity fid properties counters property name params
func (o *DeleteEntitiesEntityFidPropertiesCountersPropertyNameParams) WithTimeout(timeout time.Duration) *DeleteEntitiesEntityFidPropertiesCountersPropertyNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete entities entity fid properties counters property name params
func (o *DeleteEntitiesEntityFidPropertiesCountersPropertyNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete entities entity fid properties counters property name params
func (o *DeleteEntitiesEntityFidPropertiesCountersPropertyNameParams) WithContext(ctx context.Context) *DeleteEntitiesEntityFidPropertiesCountersPropertyNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete entities entity fid properties counters property name params
func (o *DeleteEntitiesEntityFidPropertiesCountersPropertyNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete entities entity fid properties counters property name params
func (o *DeleteEntitiesEntityFidPropertiesCountersPropertyNameParams) WithHTTPClient(client *http.Client) *DeleteEntitiesEntityFidPropertiesCountersPropertyNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete entities entity fid properties counters property name params
func (o *DeleteEntitiesEntityFidPropertiesCountersPropertyNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEntityFid adds the entityFid to the delete entities entity fid properties counters property name params
func (o *DeleteEntitiesEntityFidPropertiesCountersPropertyNameParams) WithEntityFid(entityFid string) *DeleteEntitiesEntityFidPropertiesCountersPropertyNameParams {
	o.SetEntityFid(entityFid)
	return o
}

// SetEntityFid adds the entityFid to the delete entities entity fid properties counters property name params
func (o *DeleteEntitiesEntityFidPropertiesCountersPropertyNameParams) SetEntityFid(entityFid string) {
	o.EntityFid = entityFid
}

// WithPropertyName adds the propertyName to the delete entities entity fid properties counters property name params
func (o *DeleteEntitiesEntityFidPropertiesCountersPropertyNameParams) WithPropertyName(propertyName string) *DeleteEntitiesEntityFidPropertiesCountersPropertyNameParams {
	o.SetPropertyName(propertyName)
	return o
}

// SetPropertyName adds the propertyName to the delete entities entity fid properties counters property name params
func (o *DeleteEntitiesEntityFidPropertiesCountersPropertyNameParams) SetPropertyName(propertyName string) {
	o.PropertyName = propertyName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteEntitiesEntityFidPropertiesCountersPropertyNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param entityFid
	if err := r.SetPathParam("entityFid", o.EntityFid); err != nil {
		return err
	}

	// path param propertyName
	if err := r.SetPathParam("propertyName", o.PropertyName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
