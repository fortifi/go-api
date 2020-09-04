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

// PutEntitiesEntityFidPropertiesCountersPropertyNameIncrementReader is a Reader for the PutEntitiesEntityFidPropertiesCountersPropertyNameIncrement structure.
type PutEntitiesEntityFidPropertiesCountersPropertyNameIncrementReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutEntitiesEntityFidPropertiesCountersPropertyNameIncrementReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutEntitiesEntityFidPropertiesCountersPropertyNameIncrementOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutEntitiesEntityFidPropertiesCountersPropertyNameIncrementDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutEntitiesEntityFidPropertiesCountersPropertyNameIncrementOK creates a PutEntitiesEntityFidPropertiesCountersPropertyNameIncrementOK with default headers values
func NewPutEntitiesEntityFidPropertiesCountersPropertyNameIncrementOK() *PutEntitiesEntityFidPropertiesCountersPropertyNameIncrementOK {
	return &PutEntitiesEntityFidPropertiesCountersPropertyNameIncrementOK{}
}

/*PutEntitiesEntityFidPropertiesCountersPropertyNameIncrementOK handles this case with default header values.

Counter Incremented
*/
type PutEntitiesEntityFidPropertiesCountersPropertyNameIncrementOK struct {
}

func (o *PutEntitiesEntityFidPropertiesCountersPropertyNameIncrementOK) Error() string {
	return fmt.Sprintf("[PUT /entities/{entityFid}/properties/counters/{propertyName}/increment][%d] putEntitiesEntityFidPropertiesCountersPropertyNameIncrementOK ", 200)
}

func (o *PutEntitiesEntityFidPropertiesCountersPropertyNameIncrementOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutEntitiesEntityFidPropertiesCountersPropertyNameIncrementDefault creates a PutEntitiesEntityFidPropertiesCountersPropertyNameIncrementDefault with default headers values
func NewPutEntitiesEntityFidPropertiesCountersPropertyNameIncrementDefault(code int) *PutEntitiesEntityFidPropertiesCountersPropertyNameIncrementDefault {
	return &PutEntitiesEntityFidPropertiesCountersPropertyNameIncrementDefault{
		_statusCode: code,
	}
}

/*PutEntitiesEntityFidPropertiesCountersPropertyNameIncrementDefault handles this case with default header values.

Error
*/
type PutEntitiesEntityFidPropertiesCountersPropertyNameIncrementDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the put entities entity fid properties counters property name increment default response
func (o *PutEntitiesEntityFidPropertiesCountersPropertyNameIncrementDefault) Code() int {
	return o._statusCode
}

func (o *PutEntitiesEntityFidPropertiesCountersPropertyNameIncrementDefault) Error() string {
	return fmt.Sprintf("[PUT /entities/{entityFid}/properties/counters/{propertyName}/increment][%d] PutEntitiesEntityFidPropertiesCountersPropertyNameIncrement default  %+v", o._statusCode, o.Payload)
}

func (o *PutEntitiesEntityFidPropertiesCountersPropertyNameIncrementDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PutEntitiesEntityFidPropertiesCountersPropertyNameIncrementDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
