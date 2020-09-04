// Code generated by go-swagger; DO NOT EDIT.

package custom_properties

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fortifi/go-api/models"
)

// PutEntitiesEntityFidPropertiesCountersPropertyNameDecrementReader is a Reader for the PutEntitiesEntityFidPropertiesCountersPropertyNameDecrement structure.
type PutEntitiesEntityFidPropertiesCountersPropertyNameDecrementReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutEntitiesEntityFidPropertiesCountersPropertyNameDecrementReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutEntitiesEntityFidPropertiesCountersPropertyNameDecrementOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutEntitiesEntityFidPropertiesCountersPropertyNameDecrementDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutEntitiesEntityFidPropertiesCountersPropertyNameDecrementOK creates a PutEntitiesEntityFidPropertiesCountersPropertyNameDecrementOK with default headers values
func NewPutEntitiesEntityFidPropertiesCountersPropertyNameDecrementOK() *PutEntitiesEntityFidPropertiesCountersPropertyNameDecrementOK {
	return &PutEntitiesEntityFidPropertiesCountersPropertyNameDecrementOK{}
}

/*PutEntitiesEntityFidPropertiesCountersPropertyNameDecrementOK handles this case with default header values.

Counter Decremented
*/
type PutEntitiesEntityFidPropertiesCountersPropertyNameDecrementOK struct {
}

func (o *PutEntitiesEntityFidPropertiesCountersPropertyNameDecrementOK) Error() string {
	return fmt.Sprintf("[PUT /entities/{entityFid}/properties/counters/{propertyName}/decrement][%d] putEntitiesEntityFidPropertiesCountersPropertyNameDecrementOK ", 200)
}

func (o *PutEntitiesEntityFidPropertiesCountersPropertyNameDecrementOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutEntitiesEntityFidPropertiesCountersPropertyNameDecrementDefault creates a PutEntitiesEntityFidPropertiesCountersPropertyNameDecrementDefault with default headers values
func NewPutEntitiesEntityFidPropertiesCountersPropertyNameDecrementDefault(code int) *PutEntitiesEntityFidPropertiesCountersPropertyNameDecrementDefault {
	return &PutEntitiesEntityFidPropertiesCountersPropertyNameDecrementDefault{
		_statusCode: code,
	}
}

/*PutEntitiesEntityFidPropertiesCountersPropertyNameDecrementDefault handles this case with default header values.

Error
*/
type PutEntitiesEntityFidPropertiesCountersPropertyNameDecrementDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the put entities entity fid properties counters property name decrement default response
func (o *PutEntitiesEntityFidPropertiesCountersPropertyNameDecrementDefault) Code() int {
	return o._statusCode
}

func (o *PutEntitiesEntityFidPropertiesCountersPropertyNameDecrementDefault) Error() string {
	return fmt.Sprintf("[PUT /entities/{entityFid}/properties/counters/{propertyName}/decrement][%d] PutEntitiesEntityFidPropertiesCountersPropertyNameDecrement default  %+v", o._statusCode, o.Payload)
}

func (o *PutEntitiesEntityFidPropertiesCountersPropertyNameDecrementDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PutEntitiesEntityFidPropertiesCountersPropertyNameDecrementDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
