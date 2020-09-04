// Code generated by go-swagger; DO NOT EDIT.

package marketing

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

// NewDeletePublishersPublisherFidParams creates a new DeletePublishersPublisherFidParams object
// with the default values initialized.
func NewDeletePublishersPublisherFidParams() *DeletePublishersPublisherFidParams {
	var ()
	return &DeletePublishersPublisherFidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePublishersPublisherFidParamsWithTimeout creates a new DeletePublishersPublisherFidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeletePublishersPublisherFidParamsWithTimeout(timeout time.Duration) *DeletePublishersPublisherFidParams {
	var ()
	return &DeletePublishersPublisherFidParams{

		timeout: timeout,
	}
}

// NewDeletePublishersPublisherFidParamsWithContext creates a new DeletePublishersPublisherFidParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeletePublishersPublisherFidParamsWithContext(ctx context.Context) *DeletePublishersPublisherFidParams {
	var ()
	return &DeletePublishersPublisherFidParams{

		Context: ctx,
	}
}

// NewDeletePublishersPublisherFidParamsWithHTTPClient creates a new DeletePublishersPublisherFidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeletePublishersPublisherFidParamsWithHTTPClient(client *http.Client) *DeletePublishersPublisherFidParams {
	var ()
	return &DeletePublishersPublisherFidParams{
		HTTPClient: client,
	}
}

/*DeletePublishersPublisherFidParams contains all the parameters to send to the API endpoint
for the delete publishers publisher fid operation typically these are written to a http.Request
*/
type DeletePublishersPublisherFidParams struct {

	/*PublisherFid
	  Publisher FID to use

	*/
	PublisherFid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete publishers publisher fid params
func (o *DeletePublishersPublisherFidParams) WithTimeout(timeout time.Duration) *DeletePublishersPublisherFidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete publishers publisher fid params
func (o *DeletePublishersPublisherFidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete publishers publisher fid params
func (o *DeletePublishersPublisherFidParams) WithContext(ctx context.Context) *DeletePublishersPublisherFidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete publishers publisher fid params
func (o *DeletePublishersPublisherFidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete publishers publisher fid params
func (o *DeletePublishersPublisherFidParams) WithHTTPClient(client *http.Client) *DeletePublishersPublisherFidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete publishers publisher fid params
func (o *DeletePublishersPublisherFidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPublisherFid adds the publisherFid to the delete publishers publisher fid params
func (o *DeletePublishersPublisherFidParams) WithPublisherFid(publisherFid string) *DeletePublishersPublisherFidParams {
	o.SetPublisherFid(publisherFid)
	return o
}

// SetPublisherFid adds the publisherFid to the delete publishers publisher fid params
func (o *DeletePublishersPublisherFidParams) SetPublisherFid(publisherFid string) {
	o.PublisherFid = publisherFid
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePublishersPublisherFidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param publisherFid
	if err := r.SetPathParam("publisherFid", o.PublisherFid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
