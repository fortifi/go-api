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

// DeleteCustomersCustomerFidPaymentMethodsCardsCardFidReader is a Reader for the DeleteCustomersCustomerFidPaymentMethodsCardsCardFid structure.
type DeleteCustomersCustomerFidPaymentMethodsCardsCardFidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCustomersCustomerFidPaymentMethodsCardsCardFidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteCustomersCustomerFidPaymentMethodsCardsCardFidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteCustomersCustomerFidPaymentMethodsCardsCardFidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteCustomersCustomerFidPaymentMethodsCardsCardFidOK creates a DeleteCustomersCustomerFidPaymentMethodsCardsCardFidOK with default headers values
func NewDeleteCustomersCustomerFidPaymentMethodsCardsCardFidOK() *DeleteCustomersCustomerFidPaymentMethodsCardsCardFidOK {
	return &DeleteCustomersCustomerFidPaymentMethodsCardsCardFidOK{}
}

/*DeleteCustomersCustomerFidPaymentMethodsCardsCardFidOK handles this case with default header values.

Card Removed
*/
type DeleteCustomersCustomerFidPaymentMethodsCardsCardFidOK struct {
}

func (o *DeleteCustomersCustomerFidPaymentMethodsCardsCardFidOK) Error() string {
	return fmt.Sprintf("[DELETE /customers/{customerFid}/paymentMethods/cards/{cardFid}][%d] deleteCustomersCustomerFidPaymentMethodsCardsCardFidOK ", 200)
}

func (o *DeleteCustomersCustomerFidPaymentMethodsCardsCardFidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCustomersCustomerFidPaymentMethodsCardsCardFidDefault creates a DeleteCustomersCustomerFidPaymentMethodsCardsCardFidDefault with default headers values
func NewDeleteCustomersCustomerFidPaymentMethodsCardsCardFidDefault(code int) *DeleteCustomersCustomerFidPaymentMethodsCardsCardFidDefault {
	return &DeleteCustomersCustomerFidPaymentMethodsCardsCardFidDefault{
		_statusCode: code,
	}
}

/*DeleteCustomersCustomerFidPaymentMethodsCardsCardFidDefault handles this case with default header values.

Error
*/
type DeleteCustomersCustomerFidPaymentMethodsCardsCardFidDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the delete customers customer fid payment methods cards card fid default response
func (o *DeleteCustomersCustomerFidPaymentMethodsCardsCardFidDefault) Code() int {
	return o._statusCode
}

func (o *DeleteCustomersCustomerFidPaymentMethodsCardsCardFidDefault) Error() string {
	return fmt.Sprintf("[DELETE /customers/{customerFid}/paymentMethods/cards/{cardFid}][%d] DeleteCustomersCustomerFidPaymentMethodsCardsCardFid default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteCustomersCustomerFidPaymentMethodsCardsCardFidDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *DeleteCustomersCustomerFidPaymentMethodsCardsCardFidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
