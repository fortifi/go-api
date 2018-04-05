// Code generated by go-swagger; DO NOT EDIT.

package customers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PutCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewalReader is a Reader for the PutCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewal structure.
type PutCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewalOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewPutCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewalNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewalOK creates a PutCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewalOK with default headers values
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewalOK() *PutCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewalOK {
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewalOK{}
}

/*PutCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewalOK handles this case with default header values.

Renewals Enabled
*/
type PutCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewalOK struct {
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewalOK) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/subscriptions/{subscriptionFid}/reEnableRenewal][%d] putCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewalOK ", 200)
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewalOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewalNotFound creates a PutCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewalNotFound with default headers values
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewalNotFound() *PutCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewalNotFound {
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewalNotFound{}
}

/*PutCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewalNotFound handles this case with default header values.

Subscription not found
*/
type PutCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewalNotFound struct {
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewalNotFound) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/subscriptions/{subscriptionFid}/reEnableRenewal][%d] putCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewalNotFound ", 404)
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidReEnableRenewalNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}