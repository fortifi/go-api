// Code generated by go-swagger; DO NOT EDIT.

package support

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/fortifi/go-api/models"
)

// PostTicketsReader is a Reader for the PostTickets structure.
type PostTicketsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostTicketsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostTicketsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostTicketsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostTicketsOK creates a PostTicketsOK with default headers values
func NewPostTicketsOK() *PostTicketsOK {
	return &PostTicketsOK{}
}

/*PostTicketsOK handles this case with default header values.

Ticket Information
*/
type PostTicketsOK struct {
	Payload *models.PostTicketsOKBody
}

func (o *PostTicketsOK) Error() string {
	return fmt.Sprintf("[POST /tickets][%d] postTicketsOK  %+v", 200, o.Payload)
}

func (o *PostTicketsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostTicketsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTicketsDefault creates a PostTicketsDefault with default headers values
func NewPostTicketsDefault(code int) *PostTicketsDefault {
	return &PostTicketsDefault{
		_statusCode: code,
	}
}

/*PostTicketsDefault handles this case with default header values.

Error
*/
type PostTicketsDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the post tickets default response
func (o *PostTicketsDefault) Code() int {
	return o._statusCode
}

func (o *PostTicketsDefault) Error() string {
	return fmt.Sprintf("[POST /tickets][%d] PostTickets default  %+v", o._statusCode, o.Payload)
}

func (o *PostTicketsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
