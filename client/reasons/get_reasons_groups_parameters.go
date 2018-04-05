// Code generated by go-swagger; DO NOT EDIT.

package reasons

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

// NewGetReasonsGroupsParams creates a new GetReasonsGroupsParams object
// with the default values initialized.
func NewGetReasonsGroupsParams() *GetReasonsGroupsParams {
	var ()
	return &GetReasonsGroupsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetReasonsGroupsParamsWithTimeout creates a new GetReasonsGroupsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetReasonsGroupsParamsWithTimeout(timeout time.Duration) *GetReasonsGroupsParams {
	var ()
	return &GetReasonsGroupsParams{

		timeout: timeout,
	}
}

// NewGetReasonsGroupsParamsWithContext creates a new GetReasonsGroupsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetReasonsGroupsParamsWithContext(ctx context.Context) *GetReasonsGroupsParams {
	var ()
	return &GetReasonsGroupsParams{

		Context: ctx,
	}
}

// NewGetReasonsGroupsParamsWithHTTPClient creates a new GetReasonsGroupsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetReasonsGroupsParamsWithHTTPClient(client *http.Client) *GetReasonsGroupsParams {
	var ()
	return &GetReasonsGroupsParams{
		HTTPClient: client,
	}
}

/*GetReasonsGroupsParams contains all the parameters to send to the API endpoint
for the get reasons groups operation typically these are written to a http.Request
*/
type GetReasonsGroupsParams struct {

	/*ReasonGroupType*/
	ReasonGroupType *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get reasons groups params
func (o *GetReasonsGroupsParams) WithTimeout(timeout time.Duration) *GetReasonsGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get reasons groups params
func (o *GetReasonsGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get reasons groups params
func (o *GetReasonsGroupsParams) WithContext(ctx context.Context) *GetReasonsGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get reasons groups params
func (o *GetReasonsGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get reasons groups params
func (o *GetReasonsGroupsParams) WithHTTPClient(client *http.Client) *GetReasonsGroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get reasons groups params
func (o *GetReasonsGroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReasonGroupType adds the reasonGroupType to the get reasons groups params
func (o *GetReasonsGroupsParams) WithReasonGroupType(reasonGroupType *string) *GetReasonsGroupsParams {
	o.SetReasonGroupType(reasonGroupType)
	return o
}

// SetReasonGroupType adds the reasonGroupType to the get reasons groups params
func (o *GetReasonsGroupsParams) SetReasonGroupType(reasonGroupType *string) {
	o.ReasonGroupType = reasonGroupType
}

// WriteToRequest writes these params to a swagger request
func (o *GetReasonsGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ReasonGroupType != nil {

		// query param reasonGroupType
		var qrReasonGroupType string
		if o.ReasonGroupType != nil {
			qrReasonGroupType = *o.ReasonGroupType
		}
		qReasonGroupType := qrReasonGroupType
		if qReasonGroupType != "" {
			if err := r.SetQueryParam("reasonGroupType", qReasonGroupType); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}