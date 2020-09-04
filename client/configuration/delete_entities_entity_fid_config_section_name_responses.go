// Code generated by go-swagger; DO NOT EDIT.

package configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fortifi/go-api/models"
)

// DeleteEntitiesEntityFidConfigSectionNameReader is a Reader for the DeleteEntitiesEntityFidConfigSectionName structure.
type DeleteEntitiesEntityFidConfigSectionNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteEntitiesEntityFidConfigSectionNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteEntitiesEntityFidConfigSectionNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteEntitiesEntityFidConfigSectionNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteEntitiesEntityFidConfigSectionNameOK creates a DeleteEntitiesEntityFidConfigSectionNameOK with default headers values
func NewDeleteEntitiesEntityFidConfigSectionNameOK() *DeleteEntitiesEntityFidConfigSectionNameOK {
	return &DeleteEntitiesEntityFidConfigSectionNameOK{}
}

/*DeleteEntitiesEntityFidConfigSectionNameOK handles this case with default header values.

Config Item Deleted
*/
type DeleteEntitiesEntityFidConfigSectionNameOK struct {
}

func (o *DeleteEntitiesEntityFidConfigSectionNameOK) Error() string {
	return fmt.Sprintf("[DELETE /entities/{entityFid}/config/{sectionName}][%d] deleteEntitiesEntityFidConfigSectionNameOK ", 200)
}

func (o *DeleteEntitiesEntityFidConfigSectionNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteEntitiesEntityFidConfigSectionNameDefault creates a DeleteEntitiesEntityFidConfigSectionNameDefault with default headers values
func NewDeleteEntitiesEntityFidConfigSectionNameDefault(code int) *DeleteEntitiesEntityFidConfigSectionNameDefault {
	return &DeleteEntitiesEntityFidConfigSectionNameDefault{
		_statusCode: code,
	}
}

/*DeleteEntitiesEntityFidConfigSectionNameDefault handles this case with default header values.

Error
*/
type DeleteEntitiesEntityFidConfigSectionNameDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the delete entities entity fid config section name default response
func (o *DeleteEntitiesEntityFidConfigSectionNameDefault) Code() int {
	return o._statusCode
}

func (o *DeleteEntitiesEntityFidConfigSectionNameDefault) Error() string {
	return fmt.Sprintf("[DELETE /entities/{entityFid}/config/{sectionName}][%d] DeleteEntitiesEntityFidConfigSectionName default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteEntitiesEntityFidConfigSectionNameDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *DeleteEntitiesEntityFidConfigSectionNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
