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

// PutCustomersCustomerFidReader is a Reader for the PutCustomersCustomerFid structure.
type PutCustomersCustomerFidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutCustomersCustomerFidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutCustomersCustomerFidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPutCustomersCustomerFidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutCustomersCustomerFidOK creates a PutCustomersCustomerFidOK with default headers values
func NewPutCustomersCustomerFidOK() *PutCustomersCustomerFidOK {
	return &PutCustomersCustomerFidOK{}
}

/*PutCustomersCustomerFidOK handles this case with default header values.

Customer Updated
*/
type PutCustomersCustomerFidOK struct {
}

func (o *PutCustomersCustomerFidOK) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}][%d] putCustomersCustomerFidOK ", 200)
}

func (o *PutCustomersCustomerFidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCustomersCustomerFidDefault creates a PutCustomersCustomerFidDefault with default headers values
func NewPutCustomersCustomerFidDefault(code int) *PutCustomersCustomerFidDefault {
	return &PutCustomersCustomerFidDefault{
		_statusCode: code,
	}
}

/*PutCustomersCustomerFidDefault handles this case with default header values.

Error
*/
type PutCustomersCustomerFidDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the put customers customer fid default response
func (o *PutCustomersCustomerFidDefault) Code() int {
	return o._statusCode
}

func (o *PutCustomersCustomerFidDefault) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}][%d] PutCustomersCustomerFid default  %+v", o._statusCode, o.Payload)
}

func (o *PutCustomersCustomerFidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
