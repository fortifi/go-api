// Code generated by go-swagger; DO NOT EDIT.

package orders

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fortifi/go-api/models"
)

// PutOrdersOrderFidExternalReferenceReader is a Reader for the PutOrdersOrderFidExternalReference structure.
type PutOrdersOrderFidExternalReferenceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutOrdersOrderFidExternalReferenceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutOrdersOrderFidExternalReferenceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutOrdersOrderFidExternalReferenceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutOrdersOrderFidExternalReferenceOK creates a PutOrdersOrderFidExternalReferenceOK with default headers values
func NewPutOrdersOrderFidExternalReferenceOK() *PutOrdersOrderFidExternalReferenceOK {
	return &PutOrdersOrderFidExternalReferenceOK{}
}

/*
PutOrdersOrderFidExternalReferenceOK describes a response with status code 200, with default header values.

External Reference Updated
*/
type PutOrdersOrderFidExternalReferenceOK struct {
	Payload *models.Envelope
}

// IsSuccess returns true when this put orders order fid external reference o k response has a 2xx status code
func (o *PutOrdersOrderFidExternalReferenceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put orders order fid external reference o k response has a 3xx status code
func (o *PutOrdersOrderFidExternalReferenceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put orders order fid external reference o k response has a 4xx status code
func (o *PutOrdersOrderFidExternalReferenceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put orders order fid external reference o k response has a 5xx status code
func (o *PutOrdersOrderFidExternalReferenceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put orders order fid external reference o k response a status code equal to that given
func (o *PutOrdersOrderFidExternalReferenceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the put orders order fid external reference o k response
func (o *PutOrdersOrderFidExternalReferenceOK) Code() int {
	return 200
}

func (o *PutOrdersOrderFidExternalReferenceOK) Error() string {
	return fmt.Sprintf("[PUT /orders/{orderFid}/externalReference][%d] putOrdersOrderFidExternalReferenceOK  %+v", 200, o.Payload)
}

func (o *PutOrdersOrderFidExternalReferenceOK) String() string {
	return fmt.Sprintf("[PUT /orders/{orderFid}/externalReference][%d] putOrdersOrderFidExternalReferenceOK  %+v", 200, o.Payload)
}

func (o *PutOrdersOrderFidExternalReferenceOK) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PutOrdersOrderFidExternalReferenceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOrdersOrderFidExternalReferenceDefault creates a PutOrdersOrderFidExternalReferenceDefault with default headers values
func NewPutOrdersOrderFidExternalReferenceDefault(code int) *PutOrdersOrderFidExternalReferenceDefault {
	return &PutOrdersOrderFidExternalReferenceDefault{
		_statusCode: code,
	}
}

/*
PutOrdersOrderFidExternalReferenceDefault describes a response with status code -1, with default header values.

Error
*/
type PutOrdersOrderFidExternalReferenceDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// IsSuccess returns true when this put orders order fid external reference default response has a 2xx status code
func (o *PutOrdersOrderFidExternalReferenceDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this put orders order fid external reference default response has a 3xx status code
func (o *PutOrdersOrderFidExternalReferenceDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this put orders order fid external reference default response has a 4xx status code
func (o *PutOrdersOrderFidExternalReferenceDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this put orders order fid external reference default response has a 5xx status code
func (o *PutOrdersOrderFidExternalReferenceDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this put orders order fid external reference default response a status code equal to that given
func (o *PutOrdersOrderFidExternalReferenceDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the put orders order fid external reference default response
func (o *PutOrdersOrderFidExternalReferenceDefault) Code() int {
	return o._statusCode
}

func (o *PutOrdersOrderFidExternalReferenceDefault) Error() string {
	return fmt.Sprintf("[PUT /orders/{orderFid}/externalReference][%d] PutOrdersOrderFidExternalReference default  %+v", o._statusCode, o.Payload)
}

func (o *PutOrdersOrderFidExternalReferenceDefault) String() string {
	return fmt.Sprintf("[PUT /orders/{orderFid}/externalReference][%d] PutOrdersOrderFidExternalReference default  %+v", o._statusCode, o.Payload)
}

func (o *PutOrdersOrderFidExternalReferenceDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PutOrdersOrderFidExternalReferenceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
