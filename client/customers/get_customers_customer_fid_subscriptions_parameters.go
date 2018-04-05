// Code generated by go-swagger; DO NOT EDIT.

package customers

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

// NewGetCustomersCustomerFidSubscriptionsParams creates a new GetCustomersCustomerFidSubscriptionsParams object
// with the default values initialized.
func NewGetCustomersCustomerFidSubscriptionsParams() *GetCustomersCustomerFidSubscriptionsParams {
	var ()
	return &GetCustomersCustomerFidSubscriptionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCustomersCustomerFidSubscriptionsParamsWithTimeout creates a new GetCustomersCustomerFidSubscriptionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCustomersCustomerFidSubscriptionsParamsWithTimeout(timeout time.Duration) *GetCustomersCustomerFidSubscriptionsParams {
	var ()
	return &GetCustomersCustomerFidSubscriptionsParams{

		timeout: timeout,
	}
}

// NewGetCustomersCustomerFidSubscriptionsParamsWithContext creates a new GetCustomersCustomerFidSubscriptionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCustomersCustomerFidSubscriptionsParamsWithContext(ctx context.Context) *GetCustomersCustomerFidSubscriptionsParams {
	var ()
	return &GetCustomersCustomerFidSubscriptionsParams{

		Context: ctx,
	}
}

// NewGetCustomersCustomerFidSubscriptionsParamsWithHTTPClient creates a new GetCustomersCustomerFidSubscriptionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCustomersCustomerFidSubscriptionsParamsWithHTTPClient(client *http.Client) *GetCustomersCustomerFidSubscriptionsParams {
	var ()
	return &GetCustomersCustomerFidSubscriptionsParams{
		HTTPClient: client,
	}
}

/*GetCustomersCustomerFidSubscriptionsParams contains all the parameters to send to the API endpoint
for the get customers customer fid subscriptions operation typically these are written to a http.Request
*/
type GetCustomersCustomerFidSubscriptionsParams struct {

	/*CustomerFid
	  Customer FID to use

	*/
	CustomerFid string
	/*Limit
	  Maximum number of records per page (default: 20)

	*/
	Limit *string
	/*Page
	  Page number (default: 0)

	*/
	Page *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get customers customer fid subscriptions params
func (o *GetCustomersCustomerFidSubscriptionsParams) WithTimeout(timeout time.Duration) *GetCustomersCustomerFidSubscriptionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get customers customer fid subscriptions params
func (o *GetCustomersCustomerFidSubscriptionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get customers customer fid subscriptions params
func (o *GetCustomersCustomerFidSubscriptionsParams) WithContext(ctx context.Context) *GetCustomersCustomerFidSubscriptionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get customers customer fid subscriptions params
func (o *GetCustomersCustomerFidSubscriptionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get customers customer fid subscriptions params
func (o *GetCustomersCustomerFidSubscriptionsParams) WithHTTPClient(client *http.Client) *GetCustomersCustomerFidSubscriptionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get customers customer fid subscriptions params
func (o *GetCustomersCustomerFidSubscriptionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerFid adds the customerFid to the get customers customer fid subscriptions params
func (o *GetCustomersCustomerFidSubscriptionsParams) WithCustomerFid(customerFid string) *GetCustomersCustomerFidSubscriptionsParams {
	o.SetCustomerFid(customerFid)
	return o
}

// SetCustomerFid adds the customerFid to the get customers customer fid subscriptions params
func (o *GetCustomersCustomerFidSubscriptionsParams) SetCustomerFid(customerFid string) {
	o.CustomerFid = customerFid
}

// WithLimit adds the limit to the get customers customer fid subscriptions params
func (o *GetCustomersCustomerFidSubscriptionsParams) WithLimit(limit *string) *GetCustomersCustomerFidSubscriptionsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get customers customer fid subscriptions params
func (o *GetCustomersCustomerFidSubscriptionsParams) SetLimit(limit *string) {
	o.Limit = limit
}

// WithPage adds the page to the get customers customer fid subscriptions params
func (o *GetCustomersCustomerFidSubscriptionsParams) WithPage(page *string) *GetCustomersCustomerFidSubscriptionsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get customers customer fid subscriptions params
func (o *GetCustomersCustomerFidSubscriptionsParams) SetPage(page *string) {
	o.Page = page
}

// WriteToRequest writes these params to a swagger request
func (o *GetCustomersCustomerFidSubscriptionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param customerFid
	if err := r.SetPathParam("customerFid", o.CustomerFid); err != nil {
		return err
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit string
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := qrLimit
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Page != nil {

		// query param page
		var qrPage string
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := qrPage
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
