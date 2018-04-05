// Code generated by go-swagger; DO NOT EDIT.

package customers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/fortifi/go-api/models"
)

// PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelReader is a Reader for the PutCustomersCustomerFidSubscriptionsSubscriptionFidCancel structure.
type PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutCustomersCustomerFidSubscriptionsSubscriptionFidCancelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPutCustomersCustomerFidSubscriptionsSubscriptionFidCancelBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPutCustomersCustomerFidSubscriptionsSubscriptionFidCancelNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidCancelOK creates a PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelOK with default headers values
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidCancelOK() *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelOK {
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelOK{}
}

/*PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelOK handles this case with default header values.

Subscription cancelled
*/
type PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelOK struct {
	Payload *models.Subscription
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelOK) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/subscriptions/{subscriptionFid}/cancel][%d] putCustomersCustomerFidSubscriptionsSubscriptionFidCancelOK  %+v", 200, o.Payload)
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Subscription)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidCancelBadRequest creates a PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelBadRequest with default headers values
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidCancelBadRequest() *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelBadRequest {
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelBadRequest{}
}

/*PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelBadRequest handles this case with default header values.

Invalid payload data
*/
type PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelBadRequest struct {
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelBadRequest) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/subscriptions/{subscriptionFid}/cancel][%d] putCustomersCustomerFidSubscriptionsSubscriptionFidCancelBadRequest ", 400)
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidCancelNotFound creates a PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelNotFound with default headers values
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidCancelNotFound() *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelNotFound {
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelNotFound{}
}

/*PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelNotFound handles this case with default header values.

Subscription not found
*/
type PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelNotFound struct {
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelNotFound) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/subscriptions/{subscriptionFid}/cancel][%d] putCustomersCustomerFidSubscriptionsSubscriptionFidCancelNotFound ", 404)
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidCancelNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
