// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fortifi/go-api/models"
)

// PostAccountVerificationReader is a Reader for the PostAccountVerification structure.
type PostAccountVerificationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAccountVerificationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostAccountVerificationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostAccountVerificationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostAccountVerificationOK creates a PostAccountVerificationOK with default headers values
func NewPostAccountVerificationOK() *PostAccountVerificationOK {
	return &PostAccountVerificationOK{}
}

/*
PostAccountVerificationOK describes a response with status code 200, with default header values.

Account Verification
*/
type PostAccountVerificationOK struct {
	Payload *models.AccountVerificationResponse
}

// IsSuccess returns true when this post account verification o k response has a 2xx status code
func (o *PostAccountVerificationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post account verification o k response has a 3xx status code
func (o *PostAccountVerificationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post account verification o k response has a 4xx status code
func (o *PostAccountVerificationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post account verification o k response has a 5xx status code
func (o *PostAccountVerificationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post account verification o k response a status code equal to that given
func (o *PostAccountVerificationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post account verification o k response
func (o *PostAccountVerificationOK) Code() int {
	return 200
}

func (o *PostAccountVerificationOK) Error() string {
	return fmt.Sprintf("[POST /account/verification][%d] postAccountVerificationOK  %+v", 200, o.Payload)
}

func (o *PostAccountVerificationOK) String() string {
	return fmt.Sprintf("[POST /account/verification][%d] postAccountVerificationOK  %+v", 200, o.Payload)
}

func (o *PostAccountVerificationOK) GetPayload() *models.AccountVerificationResponse {
	return o.Payload
}

func (o *PostAccountVerificationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccountVerificationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAccountVerificationDefault creates a PostAccountVerificationDefault with default headers values
func NewPostAccountVerificationDefault(code int) *PostAccountVerificationDefault {
	return &PostAccountVerificationDefault{
		_statusCode: code,
	}
}

/*
PostAccountVerificationDefault describes a response with status code -1, with default header values.

Error
*/
type PostAccountVerificationDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// IsSuccess returns true when this post account verification default response has a 2xx status code
func (o *PostAccountVerificationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this post account verification default response has a 3xx status code
func (o *PostAccountVerificationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this post account verification default response has a 4xx status code
func (o *PostAccountVerificationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this post account verification default response has a 5xx status code
func (o *PostAccountVerificationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this post account verification default response a status code equal to that given
func (o *PostAccountVerificationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the post account verification default response
func (o *PostAccountVerificationDefault) Code() int {
	return o._statusCode
}

func (o *PostAccountVerificationDefault) Error() string {
	return fmt.Sprintf("[POST /account/verification][%d] PostAccountVerification default  %+v", o._statusCode, o.Payload)
}

func (o *PostAccountVerificationDefault) String() string {
	return fmt.Sprintf("[POST /account/verification][%d] PostAccountVerification default  %+v", o._statusCode, o.Payload)
}

func (o *PostAccountVerificationDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PostAccountVerificationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
