// Code generated by go-swagger; DO NOT EDIT.

package contacts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fortifi/go-api/models"
)

// PutContactsDeviceHardwareIDSubscribeReader is a Reader for the PutContactsDeviceHardwareIDSubscribe structure.
type PutContactsDeviceHardwareIDSubscribeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutContactsDeviceHardwareIDSubscribeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutContactsDeviceHardwareIDSubscribeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutContactsDeviceHardwareIDSubscribeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutContactsDeviceHardwareIDSubscribeOK creates a PutContactsDeviceHardwareIDSubscribeOK with default headers values
func NewPutContactsDeviceHardwareIDSubscribeOK() *PutContactsDeviceHardwareIDSubscribeOK {
	return &PutContactsDeviceHardwareIDSubscribeOK{}
}

/*
PutContactsDeviceHardwareIDSubscribeOK describes a response with status code 200, with default header values.

Device Subscribed
*/
type PutContactsDeviceHardwareIDSubscribeOK struct {
}

// IsSuccess returns true when this put contacts device hardware Id subscribe o k response has a 2xx status code
func (o *PutContactsDeviceHardwareIDSubscribeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put contacts device hardware Id subscribe o k response has a 3xx status code
func (o *PutContactsDeviceHardwareIDSubscribeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put contacts device hardware Id subscribe o k response has a 4xx status code
func (o *PutContactsDeviceHardwareIDSubscribeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put contacts device hardware Id subscribe o k response has a 5xx status code
func (o *PutContactsDeviceHardwareIDSubscribeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put contacts device hardware Id subscribe o k response a status code equal to that given
func (o *PutContactsDeviceHardwareIDSubscribeOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the put contacts device hardware Id subscribe o k response
func (o *PutContactsDeviceHardwareIDSubscribeOK) Code() int {
	return 200
}

func (o *PutContactsDeviceHardwareIDSubscribeOK) Error() string {
	return fmt.Sprintf("[PUT /contacts/device/{hardwareId}/subscribe][%d] putContactsDeviceHardwareIdSubscribeOK ", 200)
}

func (o *PutContactsDeviceHardwareIDSubscribeOK) String() string {
	return fmt.Sprintf("[PUT /contacts/device/{hardwareId}/subscribe][%d] putContactsDeviceHardwareIdSubscribeOK ", 200)
}

func (o *PutContactsDeviceHardwareIDSubscribeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutContactsDeviceHardwareIDSubscribeDefault creates a PutContactsDeviceHardwareIDSubscribeDefault with default headers values
func NewPutContactsDeviceHardwareIDSubscribeDefault(code int) *PutContactsDeviceHardwareIDSubscribeDefault {
	return &PutContactsDeviceHardwareIDSubscribeDefault{
		_statusCode: code,
	}
}

/*
PutContactsDeviceHardwareIDSubscribeDefault describes a response with status code -1, with default header values.

Error
*/
type PutContactsDeviceHardwareIDSubscribeDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// IsSuccess returns true when this put contacts device hardware ID subscribe default response has a 2xx status code
func (o *PutContactsDeviceHardwareIDSubscribeDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this put contacts device hardware ID subscribe default response has a 3xx status code
func (o *PutContactsDeviceHardwareIDSubscribeDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this put contacts device hardware ID subscribe default response has a 4xx status code
func (o *PutContactsDeviceHardwareIDSubscribeDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this put contacts device hardware ID subscribe default response has a 5xx status code
func (o *PutContactsDeviceHardwareIDSubscribeDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this put contacts device hardware ID subscribe default response a status code equal to that given
func (o *PutContactsDeviceHardwareIDSubscribeDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the put contacts device hardware ID subscribe default response
func (o *PutContactsDeviceHardwareIDSubscribeDefault) Code() int {
	return o._statusCode
}

func (o *PutContactsDeviceHardwareIDSubscribeDefault) Error() string {
	return fmt.Sprintf("[PUT /contacts/device/{hardwareId}/subscribe][%d] PutContactsDeviceHardwareIDSubscribe default  %+v", o._statusCode, o.Payload)
}

func (o *PutContactsDeviceHardwareIDSubscribeDefault) String() string {
	return fmt.Sprintf("[PUT /contacts/device/{hardwareId}/subscribe][%d] PutContactsDeviceHardwareIDSubscribe default  %+v", o._statusCode, o.Payload)
}

func (o *PutContactsDeviceHardwareIDSubscribeDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PutContactsDeviceHardwareIDSubscribeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
