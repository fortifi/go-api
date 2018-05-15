// Code generated by go-swagger; DO NOT EDIT.

package marketing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/fortifi/go-api/models"
)

// PostVisitorsVisitorIDActionsActionKeyReader is a Reader for the PostVisitorsVisitorIDActionsActionKey structure.
type PostVisitorsVisitorIDActionsActionKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostVisitorsVisitorIDActionsActionKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostVisitorsVisitorIDActionsActionKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostVisitorsVisitorIDActionsActionKeyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostVisitorsVisitorIDActionsActionKeyOK creates a PostVisitorsVisitorIDActionsActionKeyOK with default headers values
func NewPostVisitorsVisitorIDActionsActionKeyOK() *PostVisitorsVisitorIDActionsActionKeyOK {
	return &PostVisitorsVisitorIDActionsActionKeyOK{}
}

/*PostVisitorsVisitorIDActionsActionKeyOK handles this case with default header values.

Action Tracked
*/
type PostVisitorsVisitorIDActionsActionKeyOK struct {
	Payload *models.PostVisitorsVisitorIDActionsActionKeyOKBody
}

func (o *PostVisitorsVisitorIDActionsActionKeyOK) Error() string {
	return fmt.Sprintf("[POST /visitors/{visitorId}/actions/{actionKey}][%d] postVisitorsVisitorIdActionsActionKeyOK  %+v", 200, o.Payload)
}

func (o *PostVisitorsVisitorIDActionsActionKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostVisitorsVisitorIDActionsActionKeyOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostVisitorsVisitorIDActionsActionKeyDefault creates a PostVisitorsVisitorIDActionsActionKeyDefault with default headers values
func NewPostVisitorsVisitorIDActionsActionKeyDefault(code int) *PostVisitorsVisitorIDActionsActionKeyDefault {
	return &PostVisitorsVisitorIDActionsActionKeyDefault{
		_statusCode: code,
	}
}

/*PostVisitorsVisitorIDActionsActionKeyDefault handles this case with default header values.

Error
*/
type PostVisitorsVisitorIDActionsActionKeyDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the post visitors visitor ID actions action key default response
func (o *PostVisitorsVisitorIDActionsActionKeyDefault) Code() int {
	return o._statusCode
}

func (o *PostVisitorsVisitorIDActionsActionKeyDefault) Error() string {
	return fmt.Sprintf("[POST /visitors/{visitorId}/actions/{actionKey}][%d] PostVisitorsVisitorIDActionsActionKey default  %+v", o._statusCode, o.Payload)
}

func (o *PostVisitorsVisitorIDActionsActionKeyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
