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

	"github.com/fortifi/go-api/models"
)

// NewPutEntitiesEntityFidPropertiesParams creates a new PutEntitiesEntityFidPropertiesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutEntitiesEntityFidPropertiesParams() *PutEntitiesEntityFidPropertiesParams {
	return &PutEntitiesEntityFidPropertiesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutEntitiesEntityFidPropertiesParamsWithTimeout creates a new PutEntitiesEntityFidPropertiesParams object
// with the ability to set a timeout on a request.
func NewPutEntitiesEntityFidPropertiesParamsWithTimeout(timeout time.Duration) *PutEntitiesEntityFidPropertiesParams {
	return &PutEntitiesEntityFidPropertiesParams{
		timeout: timeout,
	}
}

// NewPutEntitiesEntityFidPropertiesParamsWithContext creates a new PutEntitiesEntityFidPropertiesParams object
// with the ability to set a context for a request.
func NewPutEntitiesEntityFidPropertiesParamsWithContext(ctx context.Context) *PutEntitiesEntityFidPropertiesParams {
	return &PutEntitiesEntityFidPropertiesParams{
		Context: ctx,
	}
}

// NewPutEntitiesEntityFidPropertiesParamsWithHTTPClient creates a new PutEntitiesEntityFidPropertiesParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutEntitiesEntityFidPropertiesParamsWithHTTPClient(client *http.Client) *PutEntitiesEntityFidPropertiesParams {
	return &PutEntitiesEntityFidPropertiesParams{
		HTTPClient: client,
	}
}

/*
PutEntitiesEntityFidPropertiesParams contains all the parameters to send to the API endpoint

	for the put entities entity fid properties operation.

	Typically these are written to a http.Request.
*/
type PutEntitiesEntityFidPropertiesParams struct {

	/* EntityFid.

	   Entity FID to use
	*/
	EntityFid string

	// Payload.
	Payload *models.PropertyBulkSetPayload

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put entities entity fid properties params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutEntitiesEntityFidPropertiesParams) WithDefaults() *PutEntitiesEntityFidPropertiesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put entities entity fid properties params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutEntitiesEntityFidPropertiesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put entities entity fid properties params
func (o *PutEntitiesEntityFidPropertiesParams) WithTimeout(timeout time.Duration) *PutEntitiesEntityFidPropertiesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put entities entity fid properties params
func (o *PutEntitiesEntityFidPropertiesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put entities entity fid properties params
func (o *PutEntitiesEntityFidPropertiesParams) WithContext(ctx context.Context) *PutEntitiesEntityFidPropertiesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put entities entity fid properties params
func (o *PutEntitiesEntityFidPropertiesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put entities entity fid properties params
func (o *PutEntitiesEntityFidPropertiesParams) WithHTTPClient(client *http.Client) *PutEntitiesEntityFidPropertiesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put entities entity fid properties params
func (o *PutEntitiesEntityFidPropertiesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEntityFid adds the entityFid to the put entities entity fid properties params
func (o *PutEntitiesEntityFidPropertiesParams) WithEntityFid(entityFid string) *PutEntitiesEntityFidPropertiesParams {
	o.SetEntityFid(entityFid)
	return o
}

// SetEntityFid adds the entityFid to the put entities entity fid properties params
func (o *PutEntitiesEntityFidPropertiesParams) SetEntityFid(entityFid string) {
	o.EntityFid = entityFid
}

// WithPayload adds the payload to the put entities entity fid properties params
func (o *PutEntitiesEntityFidPropertiesParams) WithPayload(payload *models.PropertyBulkSetPayload) *PutEntitiesEntityFidPropertiesParams {
	o.SetPayload(payload)
	return o
}

// SetPayload adds the payload to the put entities entity fid properties params
func (o *PutEntitiesEntityFidPropertiesParams) SetPayload(payload *models.PropertyBulkSetPayload) {
	o.Payload = payload
}

// WriteToRequest writes these params to a swagger request
func (o *PutEntitiesEntityFidPropertiesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param entityFid
	if err := r.SetPathParam("entityFid", o.EntityFid); err != nil {
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
