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

// DeleteEntitiesEntityFidPropertiesFlagsPropertyNameReader is a Reader for the DeleteEntitiesEntityFidPropertiesFlagsPropertyName structure.
type DeleteEntitiesEntityFidPropertiesFlagsPropertyNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteEntitiesEntityFidPropertiesFlagsPropertyNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteEntitiesEntityFidPropertiesFlagsPropertyNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteEntitiesEntityFidPropertiesFlagsPropertyNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteEntitiesEntityFidPropertiesFlagsPropertyNameOK creates a DeleteEntitiesEntityFidPropertiesFlagsPropertyNameOK with default headers values
func NewDeleteEntitiesEntityFidPropertiesFlagsPropertyNameOK() *DeleteEntitiesEntityFidPropertiesFlagsPropertyNameOK {
	return &DeleteEntitiesEntityFidPropertiesFlagsPropertyNameOK{}
}

/*
DeleteEntitiesEntityFidPropertiesFlagsPropertyNameOK describes a response with status code 200, with default header values.

Property Deleted
*/
type DeleteEntitiesEntityFidPropertiesFlagsPropertyNameOK struct {
}

// IsSuccess returns true when this delete entities entity fid properties flags property name o k response has a 2xx status code
func (o *DeleteEntitiesEntityFidPropertiesFlagsPropertyNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete entities entity fid properties flags property name o k response has a 3xx status code
func (o *DeleteEntitiesEntityFidPropertiesFlagsPropertyNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete entities entity fid properties flags property name o k response has a 4xx status code
func (o *DeleteEntitiesEntityFidPropertiesFlagsPropertyNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete entities entity fid properties flags property name o k response has a 5xx status code
func (o *DeleteEntitiesEntityFidPropertiesFlagsPropertyNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete entities entity fid properties flags property name o k response a status code equal to that given
func (o *DeleteEntitiesEntityFidPropertiesFlagsPropertyNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete entities entity fid properties flags property name o k response
func (o *DeleteEntitiesEntityFidPropertiesFlagsPropertyNameOK) Code() int {
	return 200
}

func (o *DeleteEntitiesEntityFidPropertiesFlagsPropertyNameOK) Error() string {
	return fmt.Sprintf("[DELETE /entities/{entityFid}/properties/flags/{propertyName}][%d] deleteEntitiesEntityFidPropertiesFlagsPropertyNameOK ", 200)
}

func (o *DeleteEntitiesEntityFidPropertiesFlagsPropertyNameOK) String() string {
	return fmt.Sprintf("[DELETE /entities/{entityFid}/properties/flags/{propertyName}][%d] deleteEntitiesEntityFidPropertiesFlagsPropertyNameOK ", 200)
}

func (o *DeleteEntitiesEntityFidPropertiesFlagsPropertyNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteEntitiesEntityFidPropertiesFlagsPropertyNameDefault creates a DeleteEntitiesEntityFidPropertiesFlagsPropertyNameDefault with default headers values
func NewDeleteEntitiesEntityFidPropertiesFlagsPropertyNameDefault(code int) *DeleteEntitiesEntityFidPropertiesFlagsPropertyNameDefault {
	return &DeleteEntitiesEntityFidPropertiesFlagsPropertyNameDefault{
		_statusCode: code,
	}
}

/*
DeleteEntitiesEntityFidPropertiesFlagsPropertyNameDefault describes a response with status code -1, with default header values.

Error
*/
type DeleteEntitiesEntityFidPropertiesFlagsPropertyNameDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// IsSuccess returns true when this delete entities entity fid properties flags property name default response has a 2xx status code
func (o *DeleteEntitiesEntityFidPropertiesFlagsPropertyNameDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete entities entity fid properties flags property name default response has a 3xx status code
func (o *DeleteEntitiesEntityFidPropertiesFlagsPropertyNameDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete entities entity fid properties flags property name default response has a 4xx status code
func (o *DeleteEntitiesEntityFidPropertiesFlagsPropertyNameDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete entities entity fid properties flags property name default response has a 5xx status code
func (o *DeleteEntitiesEntityFidPropertiesFlagsPropertyNameDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete entities entity fid properties flags property name default response a status code equal to that given
func (o *DeleteEntitiesEntityFidPropertiesFlagsPropertyNameDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete entities entity fid properties flags property name default response
func (o *DeleteEntitiesEntityFidPropertiesFlagsPropertyNameDefault) Code() int {
	return o._statusCode
}

func (o *DeleteEntitiesEntityFidPropertiesFlagsPropertyNameDefault) Error() string {
	return fmt.Sprintf("[DELETE /entities/{entityFid}/properties/flags/{propertyName}][%d] DeleteEntitiesEntityFidPropertiesFlagsPropertyName default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteEntitiesEntityFidPropertiesFlagsPropertyNameDefault) String() string {
	return fmt.Sprintf("[DELETE /entities/{entityFid}/properties/flags/{propertyName}][%d] DeleteEntitiesEntityFidPropertiesFlagsPropertyName default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteEntitiesEntityFidPropertiesFlagsPropertyNameDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *DeleteEntitiesEntityFidPropertiesFlagsPropertyNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
