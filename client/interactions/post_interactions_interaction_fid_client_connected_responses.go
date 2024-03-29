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

// PostInteractionsInteractionFidClientConnectedReader is a Reader for the PostInteractionsInteractionFidClientConnected structure.
type PostInteractionsInteractionFidClientConnectedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostInteractionsInteractionFidClientConnectedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostInteractionsInteractionFidClientConnectedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostInteractionsInteractionFidClientConnectedDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostInteractionsInteractionFidClientConnectedOK creates a PostInteractionsInteractionFidClientConnectedOK with default headers values
func NewPostInteractionsInteractionFidClientConnectedOK() *PostInteractionsInteractionFidClientConnectedOK {
	return &PostInteractionsInteractionFidClientConnectedOK{}
}

/*
PostInteractionsInteractionFidClientConnectedOK describes a response with status code 200, with default header values.

Client Connected
*/
type PostInteractionsInteractionFidClientConnectedOK struct {
	Payload *models.Envelope
}

// IsSuccess returns true when this post interactions interaction fid client connected o k response has a 2xx status code
func (o *PostInteractionsInteractionFidClientConnectedOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post interactions interaction fid client connected o k response has a 3xx status code
func (o *PostInteractionsInteractionFidClientConnectedOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post interactions interaction fid client connected o k response has a 4xx status code
func (o *PostInteractionsInteractionFidClientConnectedOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post interactions interaction fid client connected o k response has a 5xx status code
func (o *PostInteractionsInteractionFidClientConnectedOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post interactions interaction fid client connected o k response a status code equal to that given
func (o *PostInteractionsInteractionFidClientConnectedOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post interactions interaction fid client connected o k response
func (o *PostInteractionsInteractionFidClientConnectedOK) Code() int {
	return 200
}

func (o *PostInteractionsInteractionFidClientConnectedOK) Error() string {
	return fmt.Sprintf("[POST /interactions/{interactionFid}/client/connected][%d] postInteractionsInteractionFidClientConnectedOK  %+v", 200, o.Payload)
}

func (o *PostInteractionsInteractionFidClientConnectedOK) String() string {
	return fmt.Sprintf("[POST /interactions/{interactionFid}/client/connected][%d] postInteractionsInteractionFidClientConnectedOK  %+v", 200, o.Payload)
}

func (o *PostInteractionsInteractionFidClientConnectedOK) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PostInteractionsInteractionFidClientConnectedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostInteractionsInteractionFidClientConnectedDefault creates a PostInteractionsInteractionFidClientConnectedDefault with default headers values
func NewPostInteractionsInteractionFidClientConnectedDefault(code int) *PostInteractionsInteractionFidClientConnectedDefault {
	return &PostInteractionsInteractionFidClientConnectedDefault{
		_statusCode: code,
	}
}

/*
PostInteractionsInteractionFidClientConnectedDefault describes a response with status code -1, with default header values.

Error
*/
type PostInteractionsInteractionFidClientConnectedDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// IsSuccess returns true when this post interactions interaction fid client connected default response has a 2xx status code
func (o *PostInteractionsInteractionFidClientConnectedDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this post interactions interaction fid client connected default response has a 3xx status code
func (o *PostInteractionsInteractionFidClientConnectedDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this post interactions interaction fid client connected default response has a 4xx status code
func (o *PostInteractionsInteractionFidClientConnectedDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this post interactions interaction fid client connected default response has a 5xx status code
func (o *PostInteractionsInteractionFidClientConnectedDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this post interactions interaction fid client connected default response a status code equal to that given
func (o *PostInteractionsInteractionFidClientConnectedDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the post interactions interaction fid client connected default response
func (o *PostInteractionsInteractionFidClientConnectedDefault) Code() int {
	return o._statusCode
}

func (o *PostInteractionsInteractionFidClientConnectedDefault) Error() string {
	return fmt.Sprintf("[POST /interactions/{interactionFid}/client/connected][%d] PostInteractionsInteractionFidClientConnected default  %+v", o._statusCode, o.Payload)
}

func (o *PostInteractionsInteractionFidClientConnectedDefault) String() string {
	return fmt.Sprintf("[POST /interactions/{interactionFid}/client/connected][%d] PostInteractionsInteractionFidClientConnected default  %+v", o._statusCode, o.Payload)
}

func (o *PostInteractionsInteractionFidClientConnectedDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PostInteractionsInteractionFidClientConnectedDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
