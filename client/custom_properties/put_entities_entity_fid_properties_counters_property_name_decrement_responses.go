// Code generated by go-swagger; DO NOT EDIT.

package custom_properties

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
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

	case 400:
		result := NewPutEntitiesEntityFidPropertiesCountersPropertyNameDecrementBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
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

// NewPutEntitiesEntityFidPropertiesCountersPropertyNameDecrementBadRequest creates a PutEntitiesEntityFidPropertiesCountersPropertyNameDecrementBadRequest with default headers values
func NewPutEntitiesEntityFidPropertiesCountersPropertyNameDecrementBadRequest() *PutEntitiesEntityFidPropertiesCountersPropertyNameDecrementBadRequest {
	return &PutEntitiesEntityFidPropertiesCountersPropertyNameDecrementBadRequest{}
}

/*PutEntitiesEntityFidPropertiesCountersPropertyNameDecrementBadRequest handles this case with default header values.

Invalid Entity
*/
type PutEntitiesEntityFidPropertiesCountersPropertyNameDecrementBadRequest struct {
}

func (o *PutEntitiesEntityFidPropertiesCountersPropertyNameDecrementBadRequest) Error() string {
	return fmt.Sprintf("[PUT /entities/{entityFid}/properties/counters/{propertyName}/decrement][%d] putEntitiesEntityFidPropertiesCountersPropertyNameDecrementBadRequest ", 400)
}

func (o *PutEntitiesEntityFidPropertiesCountersPropertyNameDecrementBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}