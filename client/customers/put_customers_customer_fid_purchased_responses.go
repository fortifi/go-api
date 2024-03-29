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

// PutCustomersCustomerFidPurchasedReader is a Reader for the PutCustomersCustomerFidPurchased structure.
type PutCustomersCustomerFidPurchasedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutCustomersCustomerFidPurchasedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutCustomersCustomerFidPurchasedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutCustomersCustomerFidPurchasedDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutCustomersCustomerFidPurchasedOK creates a PutCustomersCustomerFidPurchasedOK with default headers values
func NewPutCustomersCustomerFidPurchasedOK() *PutCustomersCustomerFidPurchasedOK {
	return &PutCustomersCustomerFidPurchasedOK{}
}

/*
PutCustomersCustomerFidPurchasedOK describes a response with status code 200, with default header values.

Customer Marked
*/
type PutCustomersCustomerFidPurchasedOK struct {
}

// IsSuccess returns true when this put customers customer fid purchased o k response has a 2xx status code
func (o *PutCustomersCustomerFidPurchasedOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put customers customer fid purchased o k response has a 3xx status code
func (o *PutCustomersCustomerFidPurchasedOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put customers customer fid purchased o k response has a 4xx status code
func (o *PutCustomersCustomerFidPurchasedOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put customers customer fid purchased o k response has a 5xx status code
func (o *PutCustomersCustomerFidPurchasedOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put customers customer fid purchased o k response a status code equal to that given
func (o *PutCustomersCustomerFidPurchasedOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the put customers customer fid purchased o k response
func (o *PutCustomersCustomerFidPurchasedOK) Code() int {
	return 200
}

func (o *PutCustomersCustomerFidPurchasedOK) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/purchased][%d] putCustomersCustomerFidPurchasedOK ", 200)
}

func (o *PutCustomersCustomerFidPurchasedOK) String() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/purchased][%d] putCustomersCustomerFidPurchasedOK ", 200)
}

func (o *PutCustomersCustomerFidPurchasedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCustomersCustomerFidPurchasedDefault creates a PutCustomersCustomerFidPurchasedDefault with default headers values
func NewPutCustomersCustomerFidPurchasedDefault(code int) *PutCustomersCustomerFidPurchasedDefault {
	return &PutCustomersCustomerFidPurchasedDefault{
		_statusCode: code,
	}
}

/*
PutCustomersCustomerFidPurchasedDefault describes a response with status code -1, with default header values.

Error
*/
type PutCustomersCustomerFidPurchasedDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// IsSuccess returns true when this put customers customer fid purchased default response has a 2xx status code
func (o *PutCustomersCustomerFidPurchasedDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this put customers customer fid purchased default response has a 3xx status code
func (o *PutCustomersCustomerFidPurchasedDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this put customers customer fid purchased default response has a 4xx status code
func (o *PutCustomersCustomerFidPurchasedDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this put customers customer fid purchased default response has a 5xx status code
func (o *PutCustomersCustomerFidPurchasedDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this put customers customer fid purchased default response a status code equal to that given
func (o *PutCustomersCustomerFidPurchasedDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the put customers customer fid purchased default response
func (o *PutCustomersCustomerFidPurchasedDefault) Code() int {
	return o._statusCode
}

func (o *PutCustomersCustomerFidPurchasedDefault) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/purchased][%d] PutCustomersCustomerFidPurchased default  %+v", o._statusCode, o.Payload)
}

func (o *PutCustomersCustomerFidPurchasedDefault) String() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/purchased][%d] PutCustomersCustomerFidPurchased default  %+v", o._statusCode, o.Payload)
}

func (o *PutCustomersCustomerFidPurchasedDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PutCustomersCustomerFidPurchasedDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
