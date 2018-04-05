// Code generated by go-swagger; DO NOT EDIT.

package custom_properties

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/fortifi/go-api/models"
)

// NewPutEntitiesEntityFidPropertiesCountersPropertyNameDecrementParams creates a new PutEntitiesEntityFidPropertiesCountersPropertyNameDecrementParams object
// with the default values initialized.
func NewPutEntitiesEntityFidPropertiesCountersPropertyNameDecrementParams() *PutEntitiesEntityFidPropertiesCountersPropertyNameDecrementParams {
	var ()
	return &PutEntitiesEntityFidPropertiesCountersPropertyNameDecrementParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutEntitiesEntityFidPropertiesCountersPropertyNameDecrementParamsWithTimeout creates a new PutEntitiesEntityFidPropertiesCountersPropertyNameDecrementParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutEntitiesEntityFidPropertiesCountersPropertyNameDecrementParamsWithTimeout(timeout time.Duration) *PutEntitiesEntityFidPropertiesCountersPropertyNameDecrementParams {
	var ()
	return &PutEntitiesEntityFidPropertiesCountersPropertyNameDecrementParams{

		timeout: timeout,
	}
}

// NewPutEntitiesEntityFidPropertiesCountersPropertyNameDecrementParamsWithContext creates a new PutEntitiesEntityFidPropertiesCountersPropertyNameDecrementParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutEntitiesEntityFidPropertiesCountersPropertyNameDecrementParamsWithContext(ctx context.Context) *PutEntitiesEntityFidPropertiesCountersPropertyNameDecrementParams {
	var ()
	return &PutEntitiesEntityFidPropertiesCountersPropertyNameDecrementParams{

		Context: ctx,
	}
}

// NewPutEntitiesEntityFidPropertiesCountersPropertyNameDecrementParamsWithHTTPClient creates a new PutEntitiesEntityFidPropertiesCountersPropertyNameDecrementParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutEntitiesEntityFidPropertiesCountersPropertyNameDecrementParamsWithHTTPClient(client *http.Client) *PutEntitiesEntityFidPropertiesCountersPropertyNameDecrementParams {
	var ()
	return &PutEntitiesEntityFidPropertiesCountersPropertyNameDecrementParams{
		HTTPClient: client,
	}
}

/*PutEntitiesEntityFidPropertiesCountersPropertyNameDecrementParams contains all the parameters to send to the API endpoint
for the put entities entity fid properties counters property name decrement operation typically these are written to a http.Request
*/
type PutEntitiesEntityFidPropertiesCountersPropertyNameDecrementParams struct {

	/*EntityFid
	  Entity FID to use

	*/
	EntityFid string
	/*Payload*/
	Payload *models.PropertyCounterPayload
	/*PropertyName
	  Property Name

	*/
	PropertyName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put entities entity fid properties counters property name decrement params
func (o *PutEntitiesEntityFidPropertiesCountersPropertyNameDecrementParams) WithTimeout(timeout time.Duration) *PutEntitiesEntityFidPropertiesCountersPropertyNameDecrementParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put entities entity fid properties counters property name decrement params
func (o *PutEntitiesEntityFidPropertiesCountersPropertyNameDecrementParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put entities entity fid properties counters property name decrement params
func (o *PutEntitiesEntityFidPropertiesCountersPropertyNameDecrementParams) WithContext(ctx context.Context) *PutEntitiesEntityFidPropertiesCountersPropertyNameDecrementParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put entities entity fid properties counters property name decrement params
func (o *PutEntitiesEntityFidPropertiesCountersPropertyNameDecrementParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put entities entity fid properties counters property name decrement params
func (o *PutEntitiesEntityFidPropertiesCountersPropertyNameDecrementParams) WithHTTPClient(client *http.Client) *PutEntitiesEntityFidPropertiesCountersPropertyNameDecrementParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put entities entity fid properties counters property name decrement params
func (o *PutEntitiesEntityFidPropertiesCountersPropertyNameDecrementParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEntityFid adds the entityFid to the put entities entity fid properties counters property name decrement params
func (o *PutEntitiesEntityFidPropertiesCountersPropertyNameDecrementParams) WithEntityFid(entityFid string) *PutEntitiesEntityFidPropertiesCountersPropertyNameDecrementParams {
	o.SetEntityFid(entityFid)
	return o
}

// SetEntityFid adds the entityFid to the put entities entity fid properties counters property name decrement params
func (o *PutEntitiesEntityFidPropertiesCountersPropertyNameDecrementParams) SetEntityFid(entityFid string) {
	o.EntityFid = entityFid
}

// WithPayload adds the payload to the put entities entity fid properties counters property name decrement params
func (o *PutEntitiesEntityFidPropertiesCountersPropertyNameDecrementParams) WithPayload(payload *models.PropertyCounterPayload) *PutEntitiesEntityFidPropertiesCountersPropertyNameDecrementParams {
	o.SetPayload(payload)
	return o
}

// SetPayload adds the payload to the put entities entity fid properties counters property name decrement params
func (o *PutEntitiesEntityFidPropertiesCountersPropertyNameDecrementParams) SetPayload(payload *models.PropertyCounterPayload) {
	o.Payload = payload
}

// WithPropertyName adds the propertyName to the put entities entity fid properties counters property name decrement params
func (o *PutEntitiesEntityFidPropertiesCountersPropertyNameDecrementParams) WithPropertyName(propertyName string) *PutEntitiesEntityFidPropertiesCountersPropertyNameDecrementParams {
	o.SetPropertyName(propertyName)
	return o
}

// SetPropertyName adds the propertyName to the put entities entity fid properties counters property name decrement params
func (o *PutEntitiesEntityFidPropertiesCountersPropertyNameDecrementParams) SetPropertyName(propertyName string) {
	o.PropertyName = propertyName
}

// WriteToRequest writes these params to a swagger request
func (o *PutEntitiesEntityFidPropertiesCountersPropertyNameDecrementParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
