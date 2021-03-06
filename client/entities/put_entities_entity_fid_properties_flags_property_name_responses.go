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

// PutEntitiesEntityFidPropertiesFlagsPropertyNameReader is a Reader for the PutEntitiesEntityFidPropertiesFlagsPropertyName structure.
type PutEntitiesEntityFidPropertiesFlagsPropertyNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutEntitiesEntityFidPropertiesFlagsPropertyNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutEntitiesEntityFidPropertiesFlagsPropertyNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutEntitiesEntityFidPropertiesFlagsPropertyNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutEntitiesEntityFidPropertiesFlagsPropertyNameOK creates a PutEntitiesEntityFidPropertiesFlagsPropertyNameOK with default headers values
func NewPutEntitiesEntityFidPropertiesFlagsPropertyNameOK() *PutEntitiesEntityFidPropertiesFlagsPropertyNameOK {
	return &PutEntitiesEntityFidPropertiesFlagsPropertyNameOK{}
}

/* PutEntitiesEntityFidPropertiesFlagsPropertyNameOK describes a response with status code 200, with default header values.

Flag Saved
*/
type PutEntitiesEntityFidPropertiesFlagsPropertyNameOK struct {
}

func (o *PutEntitiesEntityFidPropertiesFlagsPropertyNameOK) Error() string {
	return fmt.Sprintf("[PUT /entities/{entityFid}/properties/flags/{propertyName}][%d] putEntitiesEntityFidPropertiesFlagsPropertyNameOK ", 200)
}

func (o *PutEntitiesEntityFidPropertiesFlagsPropertyNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutEntitiesEntityFidPropertiesFlagsPropertyNameDefault creates a PutEntitiesEntityFidPropertiesFlagsPropertyNameDefault with default headers values
func NewPutEntitiesEntityFidPropertiesFlagsPropertyNameDefault(code int) *PutEntitiesEntityFidPropertiesFlagsPropertyNameDefault {
	return &PutEntitiesEntityFidPropertiesFlagsPropertyNameDefault{
		_statusCode: code,
	}
}

/* PutEntitiesEntityFidPropertiesFlagsPropertyNameDefault describes a response with status code -1, with default header values.

Error
*/
type PutEntitiesEntityFidPropertiesFlagsPropertyNameDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the put entities entity fid properties flags property name default response
func (o *PutEntitiesEntityFidPropertiesFlagsPropertyNameDefault) Code() int {
	return o._statusCode
}

func (o *PutEntitiesEntityFidPropertiesFlagsPropertyNameDefault) Error() string {
	return fmt.Sprintf("[PUT /entities/{entityFid}/properties/flags/{propertyName}][%d] PutEntitiesEntityFidPropertiesFlagsPropertyName default  %+v", o._statusCode, o.Payload)
}
func (o *PutEntitiesEntityFidPropertiesFlagsPropertyNameDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PutEntitiesEntityFidPropertiesFlagsPropertyNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
