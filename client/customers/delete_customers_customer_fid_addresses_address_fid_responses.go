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

// DeleteCustomersCustomerFidAddressesAddressFidReader is a Reader for the DeleteCustomersCustomerFidAddressesAddressFid structure.
type DeleteCustomersCustomerFidAddressesAddressFidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCustomersCustomerFidAddressesAddressFidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteCustomersCustomerFidAddressesAddressFidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteCustomersCustomerFidAddressesAddressFidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteCustomersCustomerFidAddressesAddressFidOK creates a DeleteCustomersCustomerFidAddressesAddressFidOK with default headers values
func NewDeleteCustomersCustomerFidAddressesAddressFidOK() *DeleteCustomersCustomerFidAddressesAddressFidOK {
	return &DeleteCustomersCustomerFidAddressesAddressFidOK{}
}

/*
DeleteCustomersCustomerFidAddressesAddressFidOK describes a response with status code 200, with default header values.

Address removed
*/
type DeleteCustomersCustomerFidAddressesAddressFidOK struct {
}

// IsSuccess returns true when this delete customers customer fid addresses address fid o k response has a 2xx status code
func (o *DeleteCustomersCustomerFidAddressesAddressFidOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete customers customer fid addresses address fid o k response has a 3xx status code
func (o *DeleteCustomersCustomerFidAddressesAddressFidOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete customers customer fid addresses address fid o k response has a 4xx status code
func (o *DeleteCustomersCustomerFidAddressesAddressFidOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete customers customer fid addresses address fid o k response has a 5xx status code
func (o *DeleteCustomersCustomerFidAddressesAddressFidOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete customers customer fid addresses address fid o k response a status code equal to that given
func (o *DeleteCustomersCustomerFidAddressesAddressFidOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteCustomersCustomerFidAddressesAddressFidOK) Error() string {
	return fmt.Sprintf("[DELETE /customers/{customerFid}/addresses/{addressFid}][%d] deleteCustomersCustomerFidAddressesAddressFidOK ", 200)
}

func (o *DeleteCustomersCustomerFidAddressesAddressFidOK) String() string {
	return fmt.Sprintf("[DELETE /customers/{customerFid}/addresses/{addressFid}][%d] deleteCustomersCustomerFidAddressesAddressFidOK ", 200)
}

func (o *DeleteCustomersCustomerFidAddressesAddressFidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCustomersCustomerFidAddressesAddressFidDefault creates a DeleteCustomersCustomerFidAddressesAddressFidDefault with default headers values
func NewDeleteCustomersCustomerFidAddressesAddressFidDefault(code int) *DeleteCustomersCustomerFidAddressesAddressFidDefault {
	return &DeleteCustomersCustomerFidAddressesAddressFidDefault{
		_statusCode: code,
	}
}

/*
DeleteCustomersCustomerFidAddressesAddressFidDefault describes a response with status code -1, with default header values.

Error
*/
type DeleteCustomersCustomerFidAddressesAddressFidDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the delete customers customer fid addresses address fid default response
func (o *DeleteCustomersCustomerFidAddressesAddressFidDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this delete customers customer fid addresses address fid default response has a 2xx status code
func (o *DeleteCustomersCustomerFidAddressesAddressFidDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete customers customer fid addresses address fid default response has a 3xx status code
func (o *DeleteCustomersCustomerFidAddressesAddressFidDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete customers customer fid addresses address fid default response has a 4xx status code
func (o *DeleteCustomersCustomerFidAddressesAddressFidDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete customers customer fid addresses address fid default response has a 5xx status code
func (o *DeleteCustomersCustomerFidAddressesAddressFidDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete customers customer fid addresses address fid default response a status code equal to that given
func (o *DeleteCustomersCustomerFidAddressesAddressFidDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DeleteCustomersCustomerFidAddressesAddressFidDefault) Error() string {
	return fmt.Sprintf("[DELETE /customers/{customerFid}/addresses/{addressFid}][%d] DeleteCustomersCustomerFidAddressesAddressFid default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteCustomersCustomerFidAddressesAddressFidDefault) String() string {
	return fmt.Sprintf("[DELETE /customers/{customerFid}/addresses/{addressFid}][%d] DeleteCustomersCustomerFidAddressesAddressFid default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteCustomersCustomerFidAddressesAddressFidDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *DeleteCustomersCustomerFidAddressesAddressFidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}