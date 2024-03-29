// Code generated by go-swagger; DO NOT EDIT.

package customers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fortifi/go-api/models"
)

// PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeReader is a Reader for the PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoCharge structure.
type PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeOK creates a PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeOK with default headers values
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeOK() *PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeOK {
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeOK{}
}

/*
PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeOK describes a response with status code 200, with default header values.

Auto charge enabled
*/
type PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeOK struct {
}

// IsSuccess returns true when this put customers customer fid subscriptions subscription fid enable auto charge o k response has a 2xx status code
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put customers customer fid subscriptions subscription fid enable auto charge o k response has a 3xx status code
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put customers customer fid subscriptions subscription fid enable auto charge o k response has a 4xx status code
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put customers customer fid subscriptions subscription fid enable auto charge o k response has a 5xx status code
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put customers customer fid subscriptions subscription fid enable auto charge o k response a status code equal to that given
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the put customers customer fid subscriptions subscription fid enable auto charge o k response
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeOK) Code() int {
	return 200
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeOK) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/subscriptions/{subscriptionFid}/enableAutoCharge][%d] putCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeOK ", 200)
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeOK) String() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/subscriptions/{subscriptionFid}/enableAutoCharge][%d] putCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeOK ", 200)
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeDefault creates a PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeDefault with default headers values
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeDefault(code int) *PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeDefault {
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeDefault{
		_statusCode: code,
	}
}

/*
PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeDefault describes a response with status code -1, with default header values.

Error
*/
type PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// IsSuccess returns true when this put customers customer fid subscriptions subscription fid enable auto charge default response has a 2xx status code
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this put customers customer fid subscriptions subscription fid enable auto charge default response has a 3xx status code
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this put customers customer fid subscriptions subscription fid enable auto charge default response has a 4xx status code
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this put customers customer fid subscriptions subscription fid enable auto charge default response has a 5xx status code
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this put customers customer fid subscriptions subscription fid enable auto charge default response a status code equal to that given
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the put customers customer fid subscriptions subscription fid enable auto charge default response
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeDefault) Code() int {
	return o._statusCode
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeDefault) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/subscriptions/{subscriptionFid}/enableAutoCharge][%d] PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoCharge default  %+v", o._statusCode, o.Payload)
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeDefault) String() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/subscriptions/{subscriptionFid}/enableAutoCharge][%d] PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoCharge default  %+v", o._statusCode, o.Payload)
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidEnableAutoChargeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
