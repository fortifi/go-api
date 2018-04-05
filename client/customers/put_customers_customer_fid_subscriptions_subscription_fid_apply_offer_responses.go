// Code generated by go-swagger; DO NOT EDIT.

package customers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PutCustomersCustomerFidSubscriptionsSubscriptionFidApplyOfferReader is a Reader for the PutCustomersCustomerFidSubscriptionsSubscriptionFidApplyOffer structure.
type PutCustomersCustomerFidSubscriptionsSubscriptionFidApplyOfferReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidApplyOfferReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutCustomersCustomerFidSubscriptionsSubscriptionFidApplyOfferOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutCustomersCustomerFidSubscriptionsSubscriptionFidApplyOfferOK creates a PutCustomersCustomerFidSubscriptionsSubscriptionFidApplyOfferOK with default headers values
func NewPutCustomersCustomerFidSubscriptionsSubscriptionFidApplyOfferOK() *PutCustomersCustomerFidSubscriptionsSubscriptionFidApplyOfferOK {
	return &PutCustomersCustomerFidSubscriptionsSubscriptionFidApplyOfferOK{}
}

/*PutCustomersCustomerFidSubscriptionsSubscriptionFidApplyOfferOK handles this case with default header values.

Offer applied
*/
type PutCustomersCustomerFidSubscriptionsSubscriptionFidApplyOfferOK struct {
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidApplyOfferOK) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/subscriptions/{subscriptionFid}/applyOffer][%d] putCustomersCustomerFidSubscriptionsSubscriptionFidApplyOfferOK ", 200)
}

func (o *PutCustomersCustomerFidSubscriptionsSubscriptionFidApplyOfferOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
