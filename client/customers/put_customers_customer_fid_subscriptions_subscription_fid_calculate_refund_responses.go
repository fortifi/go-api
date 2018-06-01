// Code generated by go-swagger; DO NOT EDIT.

package customers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/fortifi/go-api/models"
)

// PutCustomersCustomerFidSubscriptionsSubscriptionFidCalculateRefundReader is a Reader for the PutCustomersCustomerFidSubscriptionsSubscriptionFidCalculateRefund structure.
type PutCustomersCustomerFidSubscriptionsSubscriptionFidCalculateRefundReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCalculateRefundReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutCustomersCustomerFidSubscriptionsSubscriptionFidCalculateRefundOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPutCustomersCustomerFidSubscriptionsSubscriptionFidCalculateRefundDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidCalculateRefundOK creates a PutCustomersCustomerFidSubscriptionsSubscriptionFidCalculateRefundOK with default headers values
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidCalculateRefundOK() *PutCustomersCustomerFidSubscriptionsSubscriptionFidCalculateRefundOK {
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidCalculateRefundOK{}
}

/*PutCustomersCustomerFidSubscriptionsSubscriptionFidCalculateRefundOK handles this case with default header values.

Returns calculated refund amounts
*/
type PutCustomersCustomerFidSubscriptionsSubscriptionFidCalculateRefundOK struct {
	Payload *models.PutCustomersCustomerFidSubscriptionsSubscriptionFidCalculateRefundOKBody
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCalculateRefundOK) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/subscriptions/{subscriptionFid}/calculateRefund][%d] putCustomersCustomerFidSubscriptionsSubscriptionFidCalculateRefundOK  %+v", 200, o.Payload)
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCalculateRefundOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PutCustomersCustomerFidSubscriptionsSubscriptionFidCalculateRefundOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidCalculateRefundDefault creates a PutCustomersCustomerFidSubscriptionsSubscriptionFidCalculateRefundDefault with default headers values
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidCalculateRefundDefault(code int) *PutCustomersCustomerFidSubscriptionsSubscriptionFidCalculateRefundDefault {
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidCalculateRefundDefault{
		_statusCode: code,
	}
}

/*PutCustomersCustomerFidSubscriptionsSubscriptionFidCalculateRefundDefault handles this case with default header values.

Error
*/
type PutCustomersCustomerFidSubscriptionsSubscriptionFidCalculateRefundDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the put customers customer fid subscriptions subscription fid calculate refund default response
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCalculateRefundDefault) Code() int {
	return o._statusCode
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCalculateRefundDefault) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/subscriptions/{subscriptionFid}/calculateRefund][%d] PutCustomersCustomerFidSubscriptionsSubscriptionFidCalculateRefund default  %+v", o._statusCode, o.Payload)
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCalculateRefundDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
