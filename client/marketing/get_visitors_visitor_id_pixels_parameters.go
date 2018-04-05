// Code generated by go-swagger; DO NOT EDIT.

package marketing

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

// NewGetVisitorsVisitorIDPixelsParams creates a new GetVisitorsVisitorIDPixelsParams object
// with the default values initialized.
func NewGetVisitorsVisitorIDPixelsParams() *GetVisitorsVisitorIDPixelsParams {
	var ()
	return &GetVisitorsVisitorIDPixelsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetVisitorsVisitorIDPixelsParamsWithTimeout creates a new GetVisitorsVisitorIDPixelsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetVisitorsVisitorIDPixelsParamsWithTimeout(timeout time.Duration) *GetVisitorsVisitorIDPixelsParams {
	var ()
	return &GetVisitorsVisitorIDPixelsParams{

		timeout: timeout,
	}
}

// NewGetVisitorsVisitorIDPixelsParamsWithContext creates a new GetVisitorsVisitorIDPixelsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetVisitorsVisitorIDPixelsParamsWithContext(ctx context.Context) *GetVisitorsVisitorIDPixelsParams {
	var ()
	return &GetVisitorsVisitorIDPixelsParams{

		Context: ctx,
	}
}

// NewGetVisitorsVisitorIDPixelsParamsWithHTTPClient creates a new GetVisitorsVisitorIDPixelsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetVisitorsVisitorIDPixelsParamsWithHTTPClient(client *http.Client) *GetVisitorsVisitorIDPixelsParams {
	var ()
	return &GetVisitorsVisitorIDPixelsParams{
		HTTPClient: client,
	}
}

/*GetVisitorsVisitorIDPixelsParams contains all the parameters to send to the API endpoint
for the get visitors visitor ID pixels operation typically these are written to a http.Request
*/
type GetVisitorsVisitorIDPixelsParams struct {

	/*VisitorID
	  "Visitor ID from the cookie.
	If providing a pre-linked external reference, should be set to 'byref'.
	If no visitor ID is known, client IP should be provided and visitorId should be set to 'unknown'"


	*/
	VisitorID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get visitors visitor ID pixels params
func (o *GetVisitorsVisitorIDPixelsParams) WithTimeout(timeout time.Duration) *GetVisitorsVisitorIDPixelsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get visitors visitor ID pixels params
func (o *GetVisitorsVisitorIDPixelsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get visitors visitor ID pixels params
func (o *GetVisitorsVisitorIDPixelsParams) WithContext(ctx context.Context) *GetVisitorsVisitorIDPixelsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get visitors visitor ID pixels params
func (o *GetVisitorsVisitorIDPixelsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get visitors visitor ID pixels params
func (o *GetVisitorsVisitorIDPixelsParams) WithHTTPClient(client *http.Client) *GetVisitorsVisitorIDPixelsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get visitors visitor ID pixels params
func (o *GetVisitorsVisitorIDPixelsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithVisitorID adds the visitorID to the get visitors visitor ID pixels params
func (o *GetVisitorsVisitorIDPixelsParams) WithVisitorID(visitorID string) *GetVisitorsVisitorIDPixelsParams {
	o.SetVisitorID(visitorID)
	return o
}

// SetVisitorID adds the visitorId to the get visitors visitor ID pixels params
func (o *GetVisitorsVisitorIDPixelsParams) SetVisitorID(visitorID string) {
	o.VisitorID = visitorID
}

// WriteToRequest writes these params to a swagger request
func (o *GetVisitorsVisitorIDPixelsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param visitorId
	if err := r.SetPathParam("visitorId", o.VisitorID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}