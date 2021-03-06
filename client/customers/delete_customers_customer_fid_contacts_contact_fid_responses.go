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

// DeleteCustomersCustomerFidContactsContactFidReader is a Reader for the DeleteCustomersCustomerFidContactsContactFid structure.
type DeleteCustomersCustomerFidContactsContactFidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCustomersCustomerFidContactsContactFidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteCustomersCustomerFidContactsContactFidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteCustomersCustomerFidContactsContactFidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteCustomersCustomerFidContactsContactFidOK creates a DeleteCustomersCustomerFidContactsContactFidOK with default headers values
func NewDeleteCustomersCustomerFidContactsContactFidOK() *DeleteCustomersCustomerFidContactsContactFidOK {
	return &DeleteCustomersCustomerFidContactsContactFidOK{}
}

/* DeleteCustomersCustomerFidContactsContactFidOK describes a response with status code 200, with default header values.

Contact removed
*/
type DeleteCustomersCustomerFidContactsContactFidOK struct {
}

func (o *DeleteCustomersCustomerFidContactsContactFidOK) Error() string {
	return fmt.Sprintf("[DELETE /customers/{customerFid}/contacts/{contactFid}][%d] deleteCustomersCustomerFidContactsContactFidOK ", 200)
}

func (o *DeleteCustomersCustomerFidContactsContactFidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCustomersCustomerFidContactsContactFidDefault creates a DeleteCustomersCustomerFidContactsContactFidDefault with default headers values
func NewDeleteCustomersCustomerFidContactsContactFidDefault(code int) *DeleteCustomersCustomerFidContactsContactFidDefault {
	return &DeleteCustomersCustomerFidContactsContactFidDefault{
		_statusCode: code,
	}
}

/* DeleteCustomersCustomerFidContactsContactFidDefault describes a response with status code -1, with default header values.

Error
*/
type DeleteCustomersCustomerFidContactsContactFidDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the delete customers customer fid contacts contact fid default response
func (o *DeleteCustomersCustomerFidContactsContactFidDefault) Code() int {
	return o._statusCode
}

func (o *DeleteCustomersCustomerFidContactsContactFidDefault) Error() string {
	return fmt.Sprintf("[DELETE /customers/{customerFid}/contacts/{contactFid}][%d] DeleteCustomersCustomerFidContactsContactFid default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteCustomersCustomerFidContactsContactFidDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *DeleteCustomersCustomerFidContactsContactFidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
