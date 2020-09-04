// Code generated by go-swagger; DO NOT EDIT.

package customers

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

// NewPutCustomersCustomerFidParams creates a new PutCustomersCustomerFidParams object
// with the default values initialized.
func NewPutCustomersCustomerFidParams() *PutCustomersCustomerFidParams {
	var ()
	return &PutCustomersCustomerFidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutCustomersCustomerFidParamsWithTimeout creates a new PutCustomersCustomerFidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutCustomersCustomerFidParamsWithTimeout(timeout time.Duration) *PutCustomersCustomerFidParams {
	var ()
	return &PutCustomersCustomerFidParams{

		timeout: timeout,
	}
}

// NewPutCustomersCustomerFidParamsWithContext creates a new PutCustomersCustomerFidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutCustomersCustomerFidParamsWithContext(ctx context.Context) *PutCustomersCustomerFidParams {
	var ()
	return &PutCustomersCustomerFidParams{

		Context: ctx,
	}
}

// NewPutCustomersCustomerFidParamsWithHTTPClient creates a new PutCustomersCustomerFidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutCustomersCustomerFidParamsWithHTTPClient(client *http.Client) *PutCustomersCustomerFidParams {
	var ()
	return &PutCustomersCustomerFidParams{
		HTTPClient: client,
	}
}

/*PutCustomersCustomerFidParams contains all the parameters to send to the API endpoint
for the put customers customer fid operation typically these are written to a http.Request
*/
type PutCustomersCustomerFidParams struct {

	/*CustomerFid
	  Customer FID to use

	*/
	CustomerFid string
	/*DisplayName*/
	DisplayName *string
	/*FirstName*/
	FirstName *string
	/*LastName*/
	LastName *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put customers customer fid params
func (o *PutCustomersCustomerFidParams) WithTimeout(timeout time.Duration) *PutCustomersCustomerFidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put customers customer fid params
func (o *PutCustomersCustomerFidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put customers customer fid params
func (o *PutCustomersCustomerFidParams) WithContext(ctx context.Context) *PutCustomersCustomerFidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put customers customer fid params
func (o *PutCustomersCustomerFidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put customers customer fid params
func (o *PutCustomersCustomerFidParams) WithHTTPClient(client *http.Client) *PutCustomersCustomerFidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put customers customer fid params
func (o *PutCustomersCustomerFidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerFid adds the customerFid to the put customers customer fid params
func (o *PutCustomersCustomerFidParams) WithCustomerFid(customerFid string) *PutCustomersCustomerFidParams {
	o.SetCustomerFid(customerFid)
	return o
}

// SetCustomerFid adds the customerFid to the put customers customer fid params
func (o *PutCustomersCustomerFidParams) SetCustomerFid(customerFid string) {
	o.CustomerFid = customerFid
}

// WithDisplayName adds the displayName to the put customers customer fid params
func (o *PutCustomersCustomerFidParams) WithDisplayName(displayName *string) *PutCustomersCustomerFidParams {
	o.SetDisplayName(displayName)
	return o
}

// SetDisplayName adds the displayName to the put customers customer fid params
func (o *PutCustomersCustomerFidParams) SetDisplayName(displayName *string) {
	o.DisplayName = displayName
}

// WithFirstName adds the firstName to the put customers customer fid params
func (o *PutCustomersCustomerFidParams) WithFirstName(firstName *string) *PutCustomersCustomerFidParams {
	o.SetFirstName(firstName)
	return o
}

// SetFirstName adds the firstName to the put customers customer fid params
func (o *PutCustomersCustomerFidParams) SetFirstName(firstName *string) {
	o.FirstName = firstName
}

// WithLastName adds the lastName to the put customers customer fid params
func (o *PutCustomersCustomerFidParams) WithLastName(lastName *string) *PutCustomersCustomerFidParams {
	o.SetLastName(lastName)
	return o
}

// SetLastName adds the lastName to the put customers customer fid params
func (o *PutCustomersCustomerFidParams) SetLastName(lastName *string) {
	o.LastName = lastName
}

// WriteToRequest writes these params to a swagger request
func (o *PutCustomersCustomerFidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param customerFid
	if err := r.SetPathParam("customerFid", o.CustomerFid); err != nil {
		return err
	}

	if o.DisplayName != nil {

		// form param displayName
		var frDisplayName string
		if o.DisplayName != nil {
			frDisplayName = *o.DisplayName
		}
		fDisplayName := frDisplayName
		if fDisplayName != "" {
			if err := r.SetFormParam("displayName", fDisplayName); err != nil {
				return err
			}
		}

	}

	if o.FirstName != nil {

		// form param firstName
		var frFirstName string
		if o.FirstName != nil {
			frFirstName = *o.FirstName
		}
		fFirstName := frFirstName
		if fFirstName != "" {
			if err := r.SetFormParam("firstName", fFirstName); err != nil {
				return err
			}
		}

	}

	if o.LastName != nil {

		// form param lastName
		var frLastName string
		if o.LastName != nil {
			frLastName = *o.LastName
		}
		fLastName := frLastName
		if fLastName != "" {
			if err := r.SetFormParam("lastName", fLastName); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
