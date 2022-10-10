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

// GetContactsContactFidAddressesReader is a Reader for the GetContactsContactFidAddresses structure.
type GetContactsContactFidAddressesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetContactsContactFidAddressesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetContactsContactFidAddressesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetContactsContactFidAddressesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetContactsContactFidAddressesOK creates a GetContactsContactFidAddressesOK with default headers values
func NewGetContactsContactFidAddressesOK() *GetContactsContactFidAddressesOK {
	return &GetContactsContactFidAddressesOK{}
}

/*
GetContactsContactFidAddressesOK describes a response with status code 200, with default header values.

List of addresses
*/
type GetContactsContactFidAddressesOK struct {
	Payload *models.Addresses
}

// IsSuccess returns true when this get contacts contact fid addresses o k response has a 2xx status code
func (o *GetContactsContactFidAddressesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get contacts contact fid addresses o k response has a 3xx status code
func (o *GetContactsContactFidAddressesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get contacts contact fid addresses o k response has a 4xx status code
func (o *GetContactsContactFidAddressesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get contacts contact fid addresses o k response has a 5xx status code
func (o *GetContactsContactFidAddressesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get contacts contact fid addresses o k response a status code equal to that given
func (o *GetContactsContactFidAddressesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetContactsContactFidAddressesOK) Error() string {
	return fmt.Sprintf("[GET /contacts/{contactFid}/addresses][%d] getContactsContactFidAddressesOK  %+v", 200, o.Payload)
}

func (o *GetContactsContactFidAddressesOK) String() string {
	return fmt.Sprintf("[GET /contacts/{contactFid}/addresses][%d] getContactsContactFidAddressesOK  %+v", 200, o.Payload)
}

func (o *GetContactsContactFidAddressesOK) GetPayload() *models.Addresses {
	return o.Payload
}

func (o *GetContactsContactFidAddressesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Addresses)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContactsContactFidAddressesDefault creates a GetContactsContactFidAddressesDefault with default headers values
func NewGetContactsContactFidAddressesDefault(code int) *GetContactsContactFidAddressesDefault {
	return &GetContactsContactFidAddressesDefault{
		_statusCode: code,
	}
}

/*
GetContactsContactFidAddressesDefault describes a response with status code -1, with default header values.

Error
*/
type GetContactsContactFidAddressesDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the get contacts contact fid addresses default response
func (o *GetContactsContactFidAddressesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get contacts contact fid addresses default response has a 2xx status code
func (o *GetContactsContactFidAddressesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get contacts contact fid addresses default response has a 3xx status code
func (o *GetContactsContactFidAddressesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get contacts contact fid addresses default response has a 4xx status code
func (o *GetContactsContactFidAddressesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get contacts contact fid addresses default response has a 5xx status code
func (o *GetContactsContactFidAddressesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get contacts contact fid addresses default response a status code equal to that given
func (o *GetContactsContactFidAddressesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetContactsContactFidAddressesDefault) Error() string {
	return fmt.Sprintf("[GET /contacts/{contactFid}/addresses][%d] GetContactsContactFidAddresses default  %+v", o._statusCode, o.Payload)
}

func (o *GetContactsContactFidAddressesDefault) String() string {
	return fmt.Sprintf("[GET /contacts/{contactFid}/addresses][%d] GetContactsContactFidAddresses default  %+v", o._statusCode, o.Payload)
}

func (o *GetContactsContactFidAddressesDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *GetContactsContactFidAddressesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}