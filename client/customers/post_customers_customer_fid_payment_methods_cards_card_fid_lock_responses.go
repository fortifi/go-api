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

// PostCustomersCustomerFidPaymentMethodsCardsCardFidLockReader is a Reader for the PostCustomersCustomerFidPaymentMethodsCardsCardFidLock structure.
type PostCustomersCustomerFidPaymentMethodsCardsCardFidLockReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCustomersCustomerFidPaymentMethodsCardsCardFidLockReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostCustomersCustomerFidPaymentMethodsCardsCardFidLockOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostCustomersCustomerFidPaymentMethodsCardsCardFidLockDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostCustomersCustomerFidPaymentMethodsCardsCardFidLockOK creates a PostCustomersCustomerFidPaymentMethodsCardsCardFidLockOK with default headers values
func NewPostCustomersCustomerFidPaymentMethodsCardsCardFidLockOK() *PostCustomersCustomerFidPaymentMethodsCardsCardFidLockOK {
	return &PostCustomersCustomerFidPaymentMethodsCardsCardFidLockOK{}
}

/*
PostCustomersCustomerFidPaymentMethodsCardsCardFidLockOK describes a response with status code 200, with default header values.

Card Locked
*/
type PostCustomersCustomerFidPaymentMethodsCardsCardFidLockOK struct {
}

// IsSuccess returns true when this post customers customer fid payment methods cards card fid lock o k response has a 2xx status code
func (o *PostCustomersCustomerFidPaymentMethodsCardsCardFidLockOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post customers customer fid payment methods cards card fid lock o k response has a 3xx status code
func (o *PostCustomersCustomerFidPaymentMethodsCardsCardFidLockOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post customers customer fid payment methods cards card fid lock o k response has a 4xx status code
func (o *PostCustomersCustomerFidPaymentMethodsCardsCardFidLockOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post customers customer fid payment methods cards card fid lock o k response has a 5xx status code
func (o *PostCustomersCustomerFidPaymentMethodsCardsCardFidLockOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post customers customer fid payment methods cards card fid lock o k response a status code equal to that given
func (o *PostCustomersCustomerFidPaymentMethodsCardsCardFidLockOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post customers customer fid payment methods cards card fid lock o k response
func (o *PostCustomersCustomerFidPaymentMethodsCardsCardFidLockOK) Code() int {
	return 200
}

func (o *PostCustomersCustomerFidPaymentMethodsCardsCardFidLockOK) Error() string {
	return fmt.Sprintf("[POST /customers/{customerFid}/paymentMethods/cards/{cardFid}/lock][%d] postCustomersCustomerFidPaymentMethodsCardsCardFidLockOK ", 200)
}

func (o *PostCustomersCustomerFidPaymentMethodsCardsCardFidLockOK) String() string {
	return fmt.Sprintf("[POST /customers/{customerFid}/paymentMethods/cards/{cardFid}/lock][%d] postCustomersCustomerFidPaymentMethodsCardsCardFidLockOK ", 200)
}

func (o *PostCustomersCustomerFidPaymentMethodsCardsCardFidLockOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostCustomersCustomerFidPaymentMethodsCardsCardFidLockDefault creates a PostCustomersCustomerFidPaymentMethodsCardsCardFidLockDefault with default headers values
func NewPostCustomersCustomerFidPaymentMethodsCardsCardFidLockDefault(code int) *PostCustomersCustomerFidPaymentMethodsCardsCardFidLockDefault {
	return &PostCustomersCustomerFidPaymentMethodsCardsCardFidLockDefault{
		_statusCode: code,
	}
}

/*
PostCustomersCustomerFidPaymentMethodsCardsCardFidLockDefault describes a response with status code -1, with default header values.

Error
*/
type PostCustomersCustomerFidPaymentMethodsCardsCardFidLockDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// IsSuccess returns true when this post customers customer fid payment methods cards card fid lock default response has a 2xx status code
func (o *PostCustomersCustomerFidPaymentMethodsCardsCardFidLockDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this post customers customer fid payment methods cards card fid lock default response has a 3xx status code
func (o *PostCustomersCustomerFidPaymentMethodsCardsCardFidLockDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this post customers customer fid payment methods cards card fid lock default response has a 4xx status code
func (o *PostCustomersCustomerFidPaymentMethodsCardsCardFidLockDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this post customers customer fid payment methods cards card fid lock default response has a 5xx status code
func (o *PostCustomersCustomerFidPaymentMethodsCardsCardFidLockDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this post customers customer fid payment methods cards card fid lock default response a status code equal to that given
func (o *PostCustomersCustomerFidPaymentMethodsCardsCardFidLockDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the post customers customer fid payment methods cards card fid lock default response
func (o *PostCustomersCustomerFidPaymentMethodsCardsCardFidLockDefault) Code() int {
	return o._statusCode
}

func (o *PostCustomersCustomerFidPaymentMethodsCardsCardFidLockDefault) Error() string {
	return fmt.Sprintf("[POST /customers/{customerFid}/paymentMethods/cards/{cardFid}/lock][%d] PostCustomersCustomerFidPaymentMethodsCardsCardFidLock default  %+v", o._statusCode, o.Payload)
}

func (o *PostCustomersCustomerFidPaymentMethodsCardsCardFidLockDefault) String() string {
	return fmt.Sprintf("[POST /customers/{customerFid}/paymentMethods/cards/{cardFid}/lock][%d] PostCustomersCustomerFidPaymentMethodsCardsCardFidLock default  %+v", o._statusCode, o.Payload)
}

func (o *PostCustomersCustomerFidPaymentMethodsCardsCardFidLockDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PostCustomersCustomerFidPaymentMethodsCardsCardFidLockDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}