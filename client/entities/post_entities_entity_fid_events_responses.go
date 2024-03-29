// Code generated by go-swagger; DO NOT EDIT.

package entities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fortifi/go-api/models"
)

// PostEntitiesEntityFidEventsReader is a Reader for the PostEntitiesEntityFidEvents structure.
type PostEntitiesEntityFidEventsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostEntitiesEntityFidEventsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostEntitiesEntityFidEventsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostEntitiesEntityFidEventsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostEntitiesEntityFidEventsOK creates a PostEntitiesEntityFidEventsOK with default headers values
func NewPostEntitiesEntityFidEventsOK() *PostEntitiesEntityFidEventsOK {
	return &PostEntitiesEntityFidEventsOK{}
}

/*
PostEntitiesEntityFidEventsOK describes a response with status code 200, with default header values.

Event Triggered
*/
type PostEntitiesEntityFidEventsOK struct {
}

// IsSuccess returns true when this post entities entity fid events o k response has a 2xx status code
func (o *PostEntitiesEntityFidEventsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post entities entity fid events o k response has a 3xx status code
func (o *PostEntitiesEntityFidEventsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post entities entity fid events o k response has a 4xx status code
func (o *PostEntitiesEntityFidEventsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post entities entity fid events o k response has a 5xx status code
func (o *PostEntitiesEntityFidEventsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post entities entity fid events o k response a status code equal to that given
func (o *PostEntitiesEntityFidEventsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post entities entity fid events o k response
func (o *PostEntitiesEntityFidEventsOK) Code() int {
	return 200
}

func (o *PostEntitiesEntityFidEventsOK) Error() string {
	return fmt.Sprintf("[POST /entities/{entityFid}/events][%d] postEntitiesEntityFidEventsOK ", 200)
}

func (o *PostEntitiesEntityFidEventsOK) String() string {
	return fmt.Sprintf("[POST /entities/{entityFid}/events][%d] postEntitiesEntityFidEventsOK ", 200)
}

func (o *PostEntitiesEntityFidEventsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostEntitiesEntityFidEventsDefault creates a PostEntitiesEntityFidEventsDefault with default headers values
func NewPostEntitiesEntityFidEventsDefault(code int) *PostEntitiesEntityFidEventsDefault {
	return &PostEntitiesEntityFidEventsDefault{
		_statusCode: code,
	}
}

/*
PostEntitiesEntityFidEventsDefault describes a response with status code -1, with default header values.

Error
*/
type PostEntitiesEntityFidEventsDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// IsSuccess returns true when this post entities entity fid events default response has a 2xx status code
func (o *PostEntitiesEntityFidEventsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this post entities entity fid events default response has a 3xx status code
func (o *PostEntitiesEntityFidEventsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this post entities entity fid events default response has a 4xx status code
func (o *PostEntitiesEntityFidEventsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this post entities entity fid events default response has a 5xx status code
func (o *PostEntitiesEntityFidEventsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this post entities entity fid events default response a status code equal to that given
func (o *PostEntitiesEntityFidEventsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the post entities entity fid events default response
func (o *PostEntitiesEntityFidEventsDefault) Code() int {
	return o._statusCode
}

func (o *PostEntitiesEntityFidEventsDefault) Error() string {
	return fmt.Sprintf("[POST /entities/{entityFid}/events][%d] PostEntitiesEntityFidEvents default  %+v", o._statusCode, o.Payload)
}

func (o *PostEntitiesEntityFidEventsDefault) String() string {
	return fmt.Sprintf("[POST /entities/{entityFid}/events][%d] PostEntitiesEntityFidEvents default  %+v", o._statusCode, o.Payload)
}

func (o *PostEntitiesEntityFidEventsDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PostEntitiesEntityFidEventsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
