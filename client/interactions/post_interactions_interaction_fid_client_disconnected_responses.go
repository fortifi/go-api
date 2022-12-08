// Code generated by go-swagger; DO NOT EDIT.

package interactions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fortifi/go-api/models"
)

// PostInteractionsInteractionFidClientDisconnectedReader is a Reader for the PostInteractionsInteractionFidClientDisconnected structure.
type PostInteractionsInteractionFidClientDisconnectedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostInteractionsInteractionFidClientDisconnectedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostInteractionsInteractionFidClientDisconnectedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostInteractionsInteractionFidClientDisconnectedDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostInteractionsInteractionFidClientDisconnectedOK creates a PostInteractionsInteractionFidClientDisconnectedOK with default headers values
func NewPostInteractionsInteractionFidClientDisconnectedOK() *PostInteractionsInteractionFidClientDisconnectedOK {
	return &PostInteractionsInteractionFidClientDisconnectedOK{}
}

/*
PostInteractionsInteractionFidClientDisconnectedOK describes a response with status code 200, with default header values.

Client Disconnected
*/
type PostInteractionsInteractionFidClientDisconnectedOK struct {
	Payload *models.Envelope
}

// IsSuccess returns true when this post interactions interaction fid client disconnected o k response has a 2xx status code
func (o *PostInteractionsInteractionFidClientDisconnectedOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post interactions interaction fid client disconnected o k response has a 3xx status code
func (o *PostInteractionsInteractionFidClientDisconnectedOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post interactions interaction fid client disconnected o k response has a 4xx status code
func (o *PostInteractionsInteractionFidClientDisconnectedOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post interactions interaction fid client disconnected o k response has a 5xx status code
func (o *PostInteractionsInteractionFidClientDisconnectedOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post interactions interaction fid client disconnected o k response a status code equal to that given
func (o *PostInteractionsInteractionFidClientDisconnectedOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostInteractionsInteractionFidClientDisconnectedOK) Error() string {
	return fmt.Sprintf("[POST /interactions/{interactionFid}/client/disconnected][%d] postInteractionsInteractionFidClientDisconnectedOK  %+v", 200, o.Payload)
}

func (o *PostInteractionsInteractionFidClientDisconnectedOK) String() string {
	return fmt.Sprintf("[POST /interactions/{interactionFid}/client/disconnected][%d] postInteractionsInteractionFidClientDisconnectedOK  %+v", 200, o.Payload)
}

func (o *PostInteractionsInteractionFidClientDisconnectedOK) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PostInteractionsInteractionFidClientDisconnectedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostInteractionsInteractionFidClientDisconnectedDefault creates a PostInteractionsInteractionFidClientDisconnectedDefault with default headers values
func NewPostInteractionsInteractionFidClientDisconnectedDefault(code int) *PostInteractionsInteractionFidClientDisconnectedDefault {
	return &PostInteractionsInteractionFidClientDisconnectedDefault{
		_statusCode: code,
	}
}

/*
PostInteractionsInteractionFidClientDisconnectedDefault describes a response with status code -1, with default header values.

Error
*/
type PostInteractionsInteractionFidClientDisconnectedDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the post interactions interaction fid client disconnected default response
func (o *PostInteractionsInteractionFidClientDisconnectedDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this post interactions interaction fid client disconnected default response has a 2xx status code
func (o *PostInteractionsInteractionFidClientDisconnectedDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this post interactions interaction fid client disconnected default response has a 3xx status code
func (o *PostInteractionsInteractionFidClientDisconnectedDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this post interactions interaction fid client disconnected default response has a 4xx status code
func (o *PostInteractionsInteractionFidClientDisconnectedDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this post interactions interaction fid client disconnected default response has a 5xx status code
func (o *PostInteractionsInteractionFidClientDisconnectedDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this post interactions interaction fid client disconnected default response a status code equal to that given
func (o *PostInteractionsInteractionFidClientDisconnectedDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PostInteractionsInteractionFidClientDisconnectedDefault) Error() string {
	return fmt.Sprintf("[POST /interactions/{interactionFid}/client/disconnected][%d] PostInteractionsInteractionFidClientDisconnected default  %+v", o._statusCode, o.Payload)
}

func (o *PostInteractionsInteractionFidClientDisconnectedDefault) String() string {
	return fmt.Sprintf("[POST /interactions/{interactionFid}/client/disconnected][%d] PostInteractionsInteractionFidClientDisconnected default  %+v", o._statusCode, o.Payload)
}

func (o *PostInteractionsInteractionFidClientDisconnectedDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PostInteractionsInteractionFidClientDisconnectedDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
