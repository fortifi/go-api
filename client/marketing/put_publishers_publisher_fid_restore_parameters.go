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

	strfmt "github.com/go-openapi/strfmt"
)

// NewPutPublishersPublisherFidRestoreParams creates a new PutPublishersPublisherFidRestoreParams object
// with the default values initialized.
func NewPutPublishersPublisherFidRestoreParams() *PutPublishersPublisherFidRestoreParams {
	var ()
	return &PutPublishersPublisherFidRestoreParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutPublishersPublisherFidRestoreParamsWithTimeout creates a new PutPublishersPublisherFidRestoreParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutPublishersPublisherFidRestoreParamsWithTimeout(timeout time.Duration) *PutPublishersPublisherFidRestoreParams {
	var ()
	return &PutPublishersPublisherFidRestoreParams{

		timeout: timeout,
	}
}

// NewPutPublishersPublisherFidRestoreParamsWithContext creates a new PutPublishersPublisherFidRestoreParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutPublishersPublisherFidRestoreParamsWithContext(ctx context.Context) *PutPublishersPublisherFidRestoreParams {
	var ()
	return &PutPublishersPublisherFidRestoreParams{

		Context: ctx,
	}
}

// NewPutPublishersPublisherFidRestoreParamsWithHTTPClient creates a new PutPublishersPublisherFidRestoreParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutPublishersPublisherFidRestoreParamsWithHTTPClient(client *http.Client) *PutPublishersPublisherFidRestoreParams {
	var ()
	return &PutPublishersPublisherFidRestoreParams{
		HTTPClient: client,
	}
}

/*PutPublishersPublisherFidRestoreParams contains all the parameters to send to the API endpoint
for the put publishers publisher fid restore operation typically these are written to a http.Request
*/
type PutPublishersPublisherFidRestoreParams struct {

	/*PublisherFid
	  Publisher FID to use

	*/
	PublisherFid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put publishers publisher fid restore params
func (o *PutPublishersPublisherFidRestoreParams) WithTimeout(timeout time.Duration) *PutPublishersPublisherFidRestoreParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put publishers publisher fid restore params
func (o *PutPublishersPublisherFidRestoreParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put publishers publisher fid restore params
func (o *PutPublishersPublisherFidRestoreParams) WithContext(ctx context.Context) *PutPublishersPublisherFidRestoreParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put publishers publisher fid restore params
func (o *PutPublishersPublisherFidRestoreParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put publishers publisher fid restore params
func (o *PutPublishersPublisherFidRestoreParams) WithHTTPClient(client *http.Client) *PutPublishersPublisherFidRestoreParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put publishers publisher fid restore params
func (o *PutPublishersPublisherFidRestoreParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPublisherFid adds the publisherFid to the put publishers publisher fid restore params
func (o *PutPublishersPublisherFidRestoreParams) WithPublisherFid(publisherFid string) *PutPublishersPublisherFidRestoreParams {
	o.SetPublisherFid(publisherFid)
	return o
}

// SetPublisherFid adds the publisherFid to the put publishers publisher fid restore params
func (o *PutPublishersPublisherFidRestoreParams) SetPublisherFid(publisherFid string) {
	o.PublisherFid = publisherFid
}

// WriteToRequest writes these params to a swagger request
func (o *PutPublishersPublisherFidRestoreParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
