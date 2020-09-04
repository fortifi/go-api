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

// PostCustomersCustomerFidAddressesReader is a Reader for the PostCustomersCustomerFidAddresses structure.
type PostCustomersCustomerFidAddressesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCustomersCustomerFidAddressesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostCustomersCustomerFidAddressesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostCustomersCustomerFidAddressesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostCustomersCustomerFidAddressesOK creates a PostCustomersCustomerFidAddressesOK with default headers values
func NewPostCustomersCustomerFidAddressesOK() *PostCustomersCustomerFidAddressesOK {
	return &PostCustomersCustomerFidAddressesOK{}
}

/*PostCustomersCustomerFidAddressesOK handles this case with default header values.

Address Added
*/
type PostCustomersCustomerFidAddressesOK struct {
}

func (o *PostCustomersCustomerFidAddressesOK) Error() string {
	return fmt.Sprintf("[POST /customers/{customerFid}/addresses][%d] postCustomersCustomerFidAddressesOK ", 200)
}

func (o *PostCustomersCustomerFidAddressesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostCustomersCustomerFidAddressesDefault creates a PostCustomersCustomerFidAddressesDefault with default headers values
func NewPostCustomersCustomerFidAddressesDefault(code int) *PostCustomersCustomerFidAddressesDefault {
	return &PostCustomersCustomerFidAddressesDefault{
		_statusCode: code,
	}
}

/*PostCustomersCustomerFidAddressesDefault handles this case with default header values.

Error
*/
type PostCustomersCustomerFidAddressesDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the post customers customer fid addresses default response
func (o *PostCustomersCustomerFidAddressesDefault) Code() int {
	return o._statusCode
}

func (o *PostCustomersCustomerFidAddressesDefault) Error() string {
	return fmt.Sprintf("[POST /customers/{customerFid}/addresses][%d] PostCustomersCustomerFidAddresses default  %+v", o._statusCode, o.Payload)
}

func (o *PostCustomersCustomerFidAddressesDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PostCustomersCustomerFidAddressesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
