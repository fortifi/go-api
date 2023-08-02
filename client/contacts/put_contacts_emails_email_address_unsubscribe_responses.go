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

// PutContactsEmailsEmailAddressUnsubscribeReader is a Reader for the PutContactsEmailsEmailAddressUnsubscribe structure.
type PutContactsEmailsEmailAddressUnsubscribeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutContactsEmailsEmailAddressUnsubscribeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutContactsEmailsEmailAddressUnsubscribeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutContactsEmailsEmailAddressUnsubscribeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutContactsEmailsEmailAddressUnsubscribeOK creates a PutContactsEmailsEmailAddressUnsubscribeOK with default headers values
func NewPutContactsEmailsEmailAddressUnsubscribeOK() *PutContactsEmailsEmailAddressUnsubscribeOK {
	return &PutContactsEmailsEmailAddressUnsubscribeOK{}
}

/*
PutContactsEmailsEmailAddressUnsubscribeOK describes a response with status code 200, with default header values.

Email Address Unsubscribed
*/
type PutContactsEmailsEmailAddressUnsubscribeOK struct {
}

// IsSuccess returns true when this put contacts emails email address unsubscribe o k response has a 2xx status code
func (o *PutContactsEmailsEmailAddressUnsubscribeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put contacts emails email address unsubscribe o k response has a 3xx status code
func (o *PutContactsEmailsEmailAddressUnsubscribeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put contacts emails email address unsubscribe o k response has a 4xx status code
func (o *PutContactsEmailsEmailAddressUnsubscribeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put contacts emails email address unsubscribe o k response has a 5xx status code
func (o *PutContactsEmailsEmailAddressUnsubscribeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put contacts emails email address unsubscribe o k response a status code equal to that given
func (o *PutContactsEmailsEmailAddressUnsubscribeOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the put contacts emails email address unsubscribe o k response
func (o *PutContactsEmailsEmailAddressUnsubscribeOK) Code() int {
	return 200
}

func (o *PutContactsEmailsEmailAddressUnsubscribeOK) Error() string {
	return fmt.Sprintf("[PUT /contacts/emails/{emailAddress}/unsubscribe][%d] putContactsEmailsEmailAddressUnsubscribeOK ", 200)
}

func (o *PutContactsEmailsEmailAddressUnsubscribeOK) String() string {
	return fmt.Sprintf("[PUT /contacts/emails/{emailAddress}/unsubscribe][%d] putContactsEmailsEmailAddressUnsubscribeOK ", 200)
}

func (o *PutContactsEmailsEmailAddressUnsubscribeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutContactsEmailsEmailAddressUnsubscribeDefault creates a PutContactsEmailsEmailAddressUnsubscribeDefault with default headers values
func NewPutContactsEmailsEmailAddressUnsubscribeDefault(code int) *PutContactsEmailsEmailAddressUnsubscribeDefault {
	return &PutContactsEmailsEmailAddressUnsubscribeDefault{
		_statusCode: code,
	}
}

/*
PutContactsEmailsEmailAddressUnsubscribeDefault describes a response with status code -1, with default header values.

Error
*/
type PutContactsEmailsEmailAddressUnsubscribeDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// IsSuccess returns true when this put contacts emails email address unsubscribe default response has a 2xx status code
func (o *PutContactsEmailsEmailAddressUnsubscribeDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this put contacts emails email address unsubscribe default response has a 3xx status code
func (o *PutContactsEmailsEmailAddressUnsubscribeDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this put contacts emails email address unsubscribe default response has a 4xx status code
func (o *PutContactsEmailsEmailAddressUnsubscribeDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this put contacts emails email address unsubscribe default response has a 5xx status code
func (o *PutContactsEmailsEmailAddressUnsubscribeDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this put contacts emails email address unsubscribe default response a status code equal to that given
func (o *PutContactsEmailsEmailAddressUnsubscribeDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the put contacts emails email address unsubscribe default response
func (o *PutContactsEmailsEmailAddressUnsubscribeDefault) Code() int {
	return o._statusCode
}

func (o *PutContactsEmailsEmailAddressUnsubscribeDefault) Error() string {
	return fmt.Sprintf("[PUT /contacts/emails/{emailAddress}/unsubscribe][%d] PutContactsEmailsEmailAddressUnsubscribe default  %+v", o._statusCode, o.Payload)
}

func (o *PutContactsEmailsEmailAddressUnsubscribeDefault) String() string {
	return fmt.Sprintf("[PUT /contacts/emails/{emailAddress}/unsubscribe][%d] PutContactsEmailsEmailAddressUnsubscribe default  %+v", o._statusCode, o.Payload)
}

func (o *PutContactsEmailsEmailAddressUnsubscribeDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PutContactsEmailsEmailAddressUnsubscribeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
