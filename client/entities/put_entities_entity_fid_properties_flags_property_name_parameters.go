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

// NewPutEntitiesEntityFidPropertiesFlagsPropertyNameParams creates a new PutEntitiesEntityFidPropertiesFlagsPropertyNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutEntitiesEntityFidPropertiesFlagsPropertyNameParams() *PutEntitiesEntityFidPropertiesFlagsPropertyNameParams {
	return &PutEntitiesEntityFidPropertiesFlagsPropertyNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutEntitiesEntityFidPropertiesFlagsPropertyNameParamsWithTimeout creates a new PutEntitiesEntityFidPropertiesFlagsPropertyNameParams object
// with the ability to set a timeout on a request.
func NewPutEntitiesEntityFidPropertiesFlagsPropertyNameParamsWithTimeout(timeout time.Duration) *PutEntitiesEntityFidPropertiesFlagsPropertyNameParams {
	return &PutEntitiesEntityFidPropertiesFlagsPropertyNameParams{
		timeout: timeout,
	}
}

// NewPutEntitiesEntityFidPropertiesFlagsPropertyNameParamsWithContext creates a new PutEntitiesEntityFidPropertiesFlagsPropertyNameParams object
// with the ability to set a context for a request.
func NewPutEntitiesEntityFidPropertiesFlagsPropertyNameParamsWithContext(ctx context.Context) *PutEntitiesEntityFidPropertiesFlagsPropertyNameParams {
	return &PutEntitiesEntityFidPropertiesFlagsPropertyNameParams{
		Context: ctx,
	}
}

// NewPutEntitiesEntityFidPropertiesFlagsPropertyNameParamsWithHTTPClient creates a new PutEntitiesEntityFidPropertiesFlagsPropertyNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutEntitiesEntityFidPropertiesFlagsPropertyNameParamsWithHTTPClient(client *http.Client) *PutEntitiesEntityFidPropertiesFlagsPropertyNameParams {
	return &PutEntitiesEntityFidPropertiesFlagsPropertyNameParams{
		HTTPClient: client,
	}
}

/*
PutEntitiesEntityFidPropertiesFlagsPropertyNameParams contains all the parameters to send to the API endpoint

	for the put entities entity fid properties flags property name operation.

	Typically these are written to a http.Request.
*/
type PutEntitiesEntityFidPropertiesFlagsPropertyNameParams struct {

	/* EntityFid.

	   Entity FID to use
	*/
	EntityFid string

	// Payload.
	Payload *models.PropertyFlagPayload

	/* PropertyName.

	   Property Name
	*/
	PropertyName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put entities entity fid properties flags property name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutEntitiesEntityFidPropertiesFlagsPropertyNameParams) WithDefaults() *PutEntitiesEntityFidPropertiesFlagsPropertyNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put entities entity fid properties flags property name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutEntitiesEntityFidPropertiesFlagsPropertyNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put entities entity fid properties flags property name params
func (o *PutEntitiesEntityFidPropertiesFlagsPropertyNameParams) WithTimeout(timeout time.Duration) *PutEntitiesEntityFidPropertiesFlagsPropertyNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put entities entity fid properties flags property name params
func (o *PutEntitiesEntityFidPropertiesFlagsPropertyNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put entities entity fid properties flags property name params
func (o *PutEntitiesEntityFidPropertiesFlagsPropertyNameParams) WithContext(ctx context.Context) *PutEntitiesEntityFidPropertiesFlagsPropertyNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put entities entity fid properties flags property name params
func (o *PutEntitiesEntityFidPropertiesFlagsPropertyNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put entities entity fid properties flags property name params
func (o *PutEntitiesEntityFidPropertiesFlagsPropertyNameParams) WithHTTPClient(client *http.Client) *PutEntitiesEntityFidPropertiesFlagsPropertyNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put entities entity fid properties flags property name params
func (o *PutEntitiesEntityFidPropertiesFlagsPropertyNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEntityFid adds the entityFid to the put entities entity fid properties flags property name params
func (o *PutEntitiesEntityFidPropertiesFlagsPropertyNameParams) WithEntityFid(entityFid string) *PutEntitiesEntityFidPropertiesFlagsPropertyNameParams {
	o.SetEntityFid(entityFid)
	return o
}

// SetEntityFid adds the entityFid to the put entities entity fid properties flags property name params
func (o *PutEntitiesEntityFidPropertiesFlagsPropertyNameParams) SetEntityFid(entityFid string) {
	o.EntityFid = entityFid
}

// WithPayload adds the payload to the put entities entity fid properties flags property name params
func (o *PutEntitiesEntityFidPropertiesFlagsPropertyNameParams) WithPayload(payload *models.PropertyFlagPayload) *PutEntitiesEntityFidPropertiesFlagsPropertyNameParams {
	o.SetPayload(payload)
	return o
}

// SetPayload adds the payload to the put entities entity fid properties flags property name params
func (o *PutEntitiesEntityFidPropertiesFlagsPropertyNameParams) SetPayload(payload *models.PropertyFlagPayload) {
	o.Payload = payload
}

// WithPropertyName adds the propertyName to the put entities entity fid properties flags property name params
func (o *PutEntitiesEntityFidPropertiesFlagsPropertyNameParams) WithPropertyName(propertyName string) *PutEntitiesEntityFidPropertiesFlagsPropertyNameParams {
	o.SetPropertyName(propertyName)
	return o
}

// SetPropertyName adds the propertyName to the put entities entity fid properties flags property name params
func (o *PutEntitiesEntityFidPropertiesFlagsPropertyNameParams) SetPropertyName(propertyName string) {
	o.PropertyName = propertyName
}

// WriteToRequest writes these params to a swagger request
func (o *PutEntitiesEntityFidPropertiesFlagsPropertyNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param propertyName
	if err := r.SetPathParam("propertyName", o.PropertyName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
