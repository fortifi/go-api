// Code generated by go-swagger; DO NOT EDIT.

package configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/fortifi/go-api/models"
)

// GetEntitiesEntityFidConfigSectionNameReader is a Reader for the GetEntitiesEntityFidConfigSectionName structure.
type GetEntitiesEntityFidConfigSectionNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEntitiesEntityFidConfigSectionNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetEntitiesEntityFidConfigSectionNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetEntitiesEntityFidConfigSectionNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetEntitiesEntityFidConfigSectionNameOK creates a GetEntitiesEntityFidConfigSectionNameOK with default headers values
func NewGetEntitiesEntityFidConfigSectionNameOK() *GetEntitiesEntityFidConfigSectionNameOK {
	return &GetEntitiesEntityFidConfigSectionNameOK{}
}

/*GetEntitiesEntityFidConfigSectionNameOK handles this case with default header values.

List Of Config Items
*/
type GetEntitiesEntityFidConfigSectionNameOK struct {
	Payload *models.GetEntitiesEntityFidConfigSectionNameOKBody
}

func (o *GetEntitiesEntityFidConfigSectionNameOK) Error() string {
	return fmt.Sprintf("[GET /entities/{entityFid}/config/{sectionName}][%d] getEntitiesEntityFidConfigSectionNameOK  %+v", 200, o.Payload)
}

func (o *GetEntitiesEntityFidConfigSectionNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetEntitiesEntityFidConfigSectionNameOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEntitiesEntityFidConfigSectionNameDefault creates a GetEntitiesEntityFidConfigSectionNameDefault with default headers values
func NewGetEntitiesEntityFidConfigSectionNameDefault(code int) *GetEntitiesEntityFidConfigSectionNameDefault {
	return &GetEntitiesEntityFidConfigSectionNameDefault{
		_statusCode: code,
	}
}

/*GetEntitiesEntityFidConfigSectionNameDefault handles this case with default header values.

Error
*/
type GetEntitiesEntityFidConfigSectionNameDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the get entities entity fid config section name default response
func (o *GetEntitiesEntityFidConfigSectionNameDefault) Code() int {
	return o._statusCode
}

func (o *GetEntitiesEntityFidConfigSectionNameDefault) Error() string {
	return fmt.Sprintf("[GET /entities/{entityFid}/config/{sectionName}][%d] GetEntitiesEntityFidConfigSectionName default  %+v", o._statusCode, o.Payload)
}

func (o *GetEntitiesEntityFidConfigSectionNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
