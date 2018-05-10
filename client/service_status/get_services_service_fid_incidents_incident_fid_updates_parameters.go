// Code generated by go-swagger; DO NOT EDIT.

package service_status

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
)

// NewGetServicesServiceFidIncidentsIncidentFidUpdatesParams creates a new GetServicesServiceFidIncidentsIncidentFidUpdatesParams object
// with the default values initialized.
func NewGetServicesServiceFidIncidentsIncidentFidUpdatesParams() *GetServicesServiceFidIncidentsIncidentFidUpdatesParams {
	var ()
	return &GetServicesServiceFidIncidentsIncidentFidUpdatesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetServicesServiceFidIncidentsIncidentFidUpdatesParamsWithTimeout creates a new GetServicesServiceFidIncidentsIncidentFidUpdatesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetServicesServiceFidIncidentsIncidentFidUpdatesParamsWithTimeout(timeout time.Duration) *GetServicesServiceFidIncidentsIncidentFidUpdatesParams {
	var ()
	return &GetServicesServiceFidIncidentsIncidentFidUpdatesParams{

		timeout: timeout,
	}
}

// NewGetServicesServiceFidIncidentsIncidentFidUpdatesParamsWithContext creates a new GetServicesServiceFidIncidentsIncidentFidUpdatesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetServicesServiceFidIncidentsIncidentFidUpdatesParamsWithContext(ctx context.Context) *GetServicesServiceFidIncidentsIncidentFidUpdatesParams {
	var ()
	return &GetServicesServiceFidIncidentsIncidentFidUpdatesParams{

		Context: ctx,
	}
}

// NewGetServicesServiceFidIncidentsIncidentFidUpdatesParamsWithHTTPClient creates a new GetServicesServiceFidIncidentsIncidentFidUpdatesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetServicesServiceFidIncidentsIncidentFidUpdatesParamsWithHTTPClient(client *http.Client) *GetServicesServiceFidIncidentsIncidentFidUpdatesParams {
	var ()
	return &GetServicesServiceFidIncidentsIncidentFidUpdatesParams{
		HTTPClient: client,
	}
}

/*GetServicesServiceFidIncidentsIncidentFidUpdatesParams contains all the parameters to send to the API endpoint
for the get services service fid incidents incident fid updates operation typically these are written to a http.Request
*/
type GetServicesServiceFidIncidentsIncidentFidUpdatesParams struct {

	/*IncidentFid
	  Incident FID to use

	*/
	IncidentFid string
	/*ServiceFid
	  Service FID to use

	*/
	ServiceFid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get services service fid incidents incident fid updates params
func (o *GetServicesServiceFidIncidentsIncidentFidUpdatesParams) WithTimeout(timeout time.Duration) *GetServicesServiceFidIncidentsIncidentFidUpdatesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get services service fid incidents incident fid updates params
func (o *GetServicesServiceFidIncidentsIncidentFidUpdatesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get services service fid incidents incident fid updates params
func (o *GetServicesServiceFidIncidentsIncidentFidUpdatesParams) WithContext(ctx context.Context) *GetServicesServiceFidIncidentsIncidentFidUpdatesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get services service fid incidents incident fid updates params
func (o *GetServicesServiceFidIncidentsIncidentFidUpdatesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get services service fid incidents incident fid updates params
func (o *GetServicesServiceFidIncidentsIncidentFidUpdatesParams) WithHTTPClient(client *http.Client) *GetServicesServiceFidIncidentsIncidentFidUpdatesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get services service fid incidents incident fid updates params
func (o *GetServicesServiceFidIncidentsIncidentFidUpdatesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIncidentFid adds the incidentFid to the get services service fid incidents incident fid updates params
func (o *GetServicesServiceFidIncidentsIncidentFidUpdatesParams) WithIncidentFid(incidentFid string) *GetServicesServiceFidIncidentsIncidentFidUpdatesParams {
	o.SetIncidentFid(incidentFid)
	return o
}

// SetIncidentFid adds the incidentFid to the get services service fid incidents incident fid updates params
func (o *GetServicesServiceFidIncidentsIncidentFidUpdatesParams) SetIncidentFid(incidentFid string) {
	o.IncidentFid = incidentFid
}

// WithServiceFid adds the serviceFid to the get services service fid incidents incident fid updates params
func (o *GetServicesServiceFidIncidentsIncidentFidUpdatesParams) WithServiceFid(serviceFid string) *GetServicesServiceFidIncidentsIncidentFidUpdatesParams {
	o.SetServiceFid(serviceFid)
	return o
}

// SetServiceFid adds the serviceFid to the get services service fid incidents incident fid updates params
func (o *GetServicesServiceFidIncidentsIncidentFidUpdatesParams) SetServiceFid(serviceFid string) {
	o.ServiceFid = serviceFid
}

// WriteToRequest writes these params to a swagger request
func (o *GetServicesServiceFidIncidentsIncidentFidUpdatesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param incidentFid
	if err := r.SetPathParam("incidentFid", o.IncidentFid); err != nil {
		return err
	}

	// path param serviceFid
	if err := r.SetPathParam("serviceFid", o.ServiceFid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}