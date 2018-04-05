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

// GetCustomersCustomerFidSubscriptionsReader is a Reader for the GetCustomersCustomerFidSubscriptions structure.
type GetCustomersCustomerFidSubscriptionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCustomersCustomerFidSubscriptionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCustomersCustomerFidSubscriptionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCustomersCustomerFidSubscriptionsOK creates a GetCustomersCustomerFidSubscriptionsOK with default headers values
func NewGetCustomersCustomerFidSubscriptionsOK() *GetCustomersCustomerFidSubscriptionsOK {
	return &GetCustomersCustomerFidSubscriptionsOK{}
}

/*GetCustomersCustomerFidSubscriptionsOK handles this case with default header values.

List of subscription summaries
*/
type GetCustomersCustomerFidSubscriptionsOK struct {
	Payload *models.Subscriptions
}

func (o *GetCustomersCustomerFidSubscriptionsOK) Error() string {
	return fmt.Sprintf("[GET /customers/{customerFid}/subscriptions][%d] getCustomersCustomerFidSubscriptionsOK  %+v", 200, o.Payload)
}

func (o *GetCustomersCustomerFidSubscriptionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Subscriptions)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
