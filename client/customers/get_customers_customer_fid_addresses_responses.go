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

// GetCustomersCustomerFidAddressesReader is a Reader for the GetCustomersCustomerFidAddresses structure.
type GetCustomersCustomerFidAddressesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCustomersCustomerFidAddressesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCustomersCustomerFidAddressesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetCustomersCustomerFidAddressesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCustomersCustomerFidAddressesOK creates a GetCustomersCustomerFidAddressesOK with default headers values
func NewGetCustomersCustomerFidAddressesOK() *GetCustomersCustomerFidAddressesOK {
	return &GetCustomersCustomerFidAddressesOK{}
}

/*GetCustomersCustomerFidAddressesOK handles this case with default header values.

List of addresses
*/
type GetCustomersCustomerFidAddressesOK struct {
	Payload *models.Addresses
}

func (o *GetCustomersCustomerFidAddressesOK) Error() string {
	return fmt.Sprintf("[GET /customers/{customerFid}/addresses][%d] getCustomersCustomerFidAddressesOK  %+v", 200, o.Payload)
}

func (o *GetCustomersCustomerFidAddressesOK) GetPayload() *models.Addresses {
	return o.Payload
}

func (o *GetCustomersCustomerFidAddressesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Addresses)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCustomersCustomerFidAddressesDefault creates a GetCustomersCustomerFidAddressesDefault with default headers values
func NewGetCustomersCustomerFidAddressesDefault(code int) *GetCustomersCustomerFidAddressesDefault {
	return &GetCustomersCustomerFidAddressesDefault{
		_statusCode: code,
	}
}

/*GetCustomersCustomerFidAddressesDefault handles this case with default header values.

Error
*/
type GetCustomersCustomerFidAddressesDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the get customers customer fid addresses default response
func (o *GetCustomersCustomerFidAddressesDefault) Code() int {
	return o._statusCode
}

func (o *GetCustomersCustomerFidAddressesDefault) Error() string {
	return fmt.Sprintf("[GET /customers/{customerFid}/addresses][%d] GetCustomersCustomerFidAddresses default  %+v", o._statusCode, o.Payload)
}

func (o *GetCustomersCustomerFidAddressesDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *GetCustomersCustomerFidAddressesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
