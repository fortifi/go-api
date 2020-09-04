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

// PutCustomersCustomerFidContactsContactFidReader is a Reader for the PutCustomersCustomerFidContactsContactFid structure.
type PutCustomersCustomerFidContactsContactFidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutCustomersCustomerFidContactsContactFidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutCustomersCustomerFidContactsContactFidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutCustomersCustomerFidContactsContactFidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutCustomersCustomerFidContactsContactFidOK creates a PutCustomersCustomerFidContactsContactFidOK with default headers values
func NewPutCustomersCustomerFidContactsContactFidOK() *PutCustomersCustomerFidContactsContactFidOK {
	return &PutCustomersCustomerFidContactsContactFidOK{}
}

/*PutCustomersCustomerFidContactsContactFidOK handles this case with default header values.

Contact Updated
*/
type PutCustomersCustomerFidContactsContactFidOK struct {
}

func (o *PutCustomersCustomerFidContactsContactFidOK) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/contacts/{contactFid}][%d] putCustomersCustomerFidContactsContactFidOK ", 200)
}

func (o *PutCustomersCustomerFidContactsContactFidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCustomersCustomerFidContactsContactFidDefault creates a PutCustomersCustomerFidContactsContactFidDefault with default headers values
func NewPutCustomersCustomerFidContactsContactFidDefault(code int) *PutCustomersCustomerFidContactsContactFidDefault {
	return &PutCustomersCustomerFidContactsContactFidDefault{
		_statusCode: code,
	}
}

/*PutCustomersCustomerFidContactsContactFidDefault handles this case with default header values.

Error
*/
type PutCustomersCustomerFidContactsContactFidDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the put customers customer fid contacts contact fid default response
func (o *PutCustomersCustomerFidContactsContactFidDefault) Code() int {
	return o._statusCode
}

func (o *PutCustomersCustomerFidContactsContactFidDefault) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/contacts/{contactFid}][%d] PutCustomersCustomerFidContactsContactFid default  %+v", o._statusCode, o.Payload)
}

func (o *PutCustomersCustomerFidContactsContactFidDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PutCustomersCustomerFidContactsContactFidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
