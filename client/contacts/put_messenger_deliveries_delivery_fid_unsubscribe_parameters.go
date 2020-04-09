// Code generated by go-swagger; DO NOT EDIT.

package contacts

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

// NewPutMessengerDeliveriesDeliveryFidUnsubscribeParams creates a new PutMessengerDeliveriesDeliveryFidUnsubscribeParams object
// with the default values initialized.
func NewPutMessengerDeliveriesDeliveryFidUnsubscribeParams() *PutMessengerDeliveriesDeliveryFidUnsubscribeParams {
	var ()
	return &PutMessengerDeliveriesDeliveryFidUnsubscribeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutMessengerDeliveriesDeliveryFidUnsubscribeParamsWithTimeout creates a new PutMessengerDeliveriesDeliveryFidUnsubscribeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutMessengerDeliveriesDeliveryFidUnsubscribeParamsWithTimeout(timeout time.Duration) *PutMessengerDeliveriesDeliveryFidUnsubscribeParams {
	var ()
	return &PutMessengerDeliveriesDeliveryFidUnsubscribeParams{

		timeout: timeout,
	}
}

// NewPutMessengerDeliveriesDeliveryFidUnsubscribeParamsWithContext creates a new PutMessengerDeliveriesDeliveryFidUnsubscribeParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutMessengerDeliveriesDeliveryFidUnsubscribeParamsWithContext(ctx context.Context) *PutMessengerDeliveriesDeliveryFidUnsubscribeParams {
	var ()
	return &PutMessengerDeliveriesDeliveryFidUnsubscribeParams{

		Context: ctx,
	}
}

// NewPutMessengerDeliveriesDeliveryFidUnsubscribeParamsWithHTTPClient creates a new PutMessengerDeliveriesDeliveryFidUnsubscribeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutMessengerDeliveriesDeliveryFidUnsubscribeParamsWithHTTPClient(client *http.Client) *PutMessengerDeliveriesDeliveryFidUnsubscribeParams {
	var ()
	return &PutMessengerDeliveriesDeliveryFidUnsubscribeParams{
		HTTPClient: client,
	}
}

/*PutMessengerDeliveriesDeliveryFidUnsubscribeParams contains all the parameters to send to the API endpoint
for the put messenger deliveries delivery fid unsubscribe operation typically these are written to a http.Request
*/
type PutMessengerDeliveriesDeliveryFidUnsubscribeParams struct {

	/*ClientIP
	  IP Address of the visitor

	*/
	ClientIP *string
	/*DeliveryFid
	  Delivery FID

	*/
	DeliveryFid string
	/*Encoding
	  Encoding from the visitors browser 'HTTP_ACCEPT_ENCODING'

	*/
	Encoding *string
	/*Language
	  Language from visitors browser 'HTTP_ACCEPT_LANGUAGE'

	*/
	Language *string
	/*UserAgent
	  User Agent of the visitors browser 'HTTP_USER_AGENT'

	*/
	UserAgent *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put messenger deliveries delivery fid unsubscribe params
func (o *PutMessengerDeliveriesDeliveryFidUnsubscribeParams) WithTimeout(timeout time.Duration) *PutMessengerDeliveriesDeliveryFidUnsubscribeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put messenger deliveries delivery fid unsubscribe params
func (o *PutMessengerDeliveriesDeliveryFidUnsubscribeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put messenger deliveries delivery fid unsubscribe params
func (o *PutMessengerDeliveriesDeliveryFidUnsubscribeParams) WithContext(ctx context.Context) *PutMessengerDeliveriesDeliveryFidUnsubscribeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put messenger deliveries delivery fid unsubscribe params
func (o *PutMessengerDeliveriesDeliveryFidUnsubscribeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put messenger deliveries delivery fid unsubscribe params
func (o *PutMessengerDeliveriesDeliveryFidUnsubscribeParams) WithHTTPClient(client *http.Client) *PutMessengerDeliveriesDeliveryFidUnsubscribeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put messenger deliveries delivery fid unsubscribe params
func (o *PutMessengerDeliveriesDeliveryFidUnsubscribeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClientIP adds the clientIP to the put messenger deliveries delivery fid unsubscribe params
func (o *PutMessengerDeliveriesDeliveryFidUnsubscribeParams) WithClientIP(clientIP *string) *PutMessengerDeliveriesDeliveryFidUnsubscribeParams {
	o.SetClientIP(clientIP)
	return o
}

// SetClientIP adds the clientIp to the put messenger deliveries delivery fid unsubscribe params
func (o *PutMessengerDeliveriesDeliveryFidUnsubscribeParams) SetClientIP(clientIP *string) {
	o.ClientIP = clientIP
}

// WithDeliveryFid adds the deliveryFid to the put messenger deliveries delivery fid unsubscribe params
func (o *PutMessengerDeliveriesDeliveryFidUnsubscribeParams) WithDeliveryFid(deliveryFid string) *PutMessengerDeliveriesDeliveryFidUnsubscribeParams {
	o.SetDeliveryFid(deliveryFid)
	return o
}

// SetDeliveryFid adds the deliveryFid to the put messenger deliveries delivery fid unsubscribe params
func (o *PutMessengerDeliveriesDeliveryFidUnsubscribeParams) SetDeliveryFid(deliveryFid string) {
	o.DeliveryFid = deliveryFid
}

// WithEncoding adds the encoding to the put messenger deliveries delivery fid unsubscribe params
func (o *PutMessengerDeliveriesDeliveryFidUnsubscribeParams) WithEncoding(encoding *string) *PutMessengerDeliveriesDeliveryFidUnsubscribeParams {
	o.SetEncoding(encoding)
	return o
}

// SetEncoding adds the encoding to the put messenger deliveries delivery fid unsubscribe params
func (o *PutMessengerDeliveriesDeliveryFidUnsubscribeParams) SetEncoding(encoding *string) {
	o.Encoding = encoding
}

// WithLanguage adds the language to the put messenger deliveries delivery fid unsubscribe params
func (o *PutMessengerDeliveriesDeliveryFidUnsubscribeParams) WithLanguage(language *string) *PutMessengerDeliveriesDeliveryFidUnsubscribeParams {
	o.SetLanguage(language)
	return o
}

// SetLanguage adds the language to the put messenger deliveries delivery fid unsubscribe params
func (o *PutMessengerDeliveriesDeliveryFidUnsubscribeParams) SetLanguage(language *string) {
	o.Language = language
}

// WithUserAgent adds the userAgent to the put messenger deliveries delivery fid unsubscribe params
func (o *PutMessengerDeliveriesDeliveryFidUnsubscribeParams) WithUserAgent(userAgent *string) *PutMessengerDeliveriesDeliveryFidUnsubscribeParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the put messenger deliveries delivery fid unsubscribe params
func (o *PutMessengerDeliveriesDeliveryFidUnsubscribeParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WriteToRequest writes these params to a swagger request
func (o *PutMessengerDeliveriesDeliveryFidUnsubscribeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ClientIP != nil {

		// form param clientIp
		var frClientIP string
		if o.ClientIP != nil {
			frClientIP = *o.ClientIP
		}
		fClientIP := frClientIP
		if fClientIP != "" {
			if err := r.SetFormParam("clientIp", fClientIP); err != nil {
				return err
			}
		}

	}

	// path param deliveryFid
	if err := r.SetPathParam("deliveryFid", o.DeliveryFid); err != nil {
		return err
	}

	if o.Encoding != nil {

		// form param encoding
		var frEncoding string
		if o.Encoding != nil {
			frEncoding = *o.Encoding
		}
		fEncoding := frEncoding
		if fEncoding != "" {
			if err := r.SetFormParam("encoding", fEncoding); err != nil {
				return err
			}
		}

	}

	if o.Language != nil {

		// form param language
		var frLanguage string
		if o.Language != nil {
			frLanguage = *o.Language
		}
		fLanguage := frLanguage
		if fLanguage != "" {
			if err := r.SetFormParam("language", fLanguage); err != nil {
				return err
			}
		}

	}

	if o.UserAgent != nil {

		// form param userAgent
		var frUserAgent string
		if o.UserAgent != nil {
			frUserAgent = *o.UserAgent
		}
		fUserAgent := frUserAgent
		if fUserAgent != "" {
			if err := r.SetFormParam("userAgent", fUserAgent); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}