// Code generated by go-swagger; DO NOT EDIT.

package custom_properties

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteEntitiesEntityFidPropertiesValuesPropertyNameReader is a Reader for the DeleteEntitiesEntityFidPropertiesValuesPropertyName structure.
type DeleteEntitiesEntityFidPropertiesValuesPropertyNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteEntitiesEntityFidPropertiesValuesPropertyNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteEntitiesEntityFidPropertiesValuesPropertyNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteEntitiesEntityFidPropertiesValuesPropertyNameOK creates a DeleteEntitiesEntityFidPropertiesValuesPropertyNameOK with default headers values
func NewDeleteEntitiesEntityFidPropertiesValuesPropertyNameOK() *DeleteEntitiesEntityFidPropertiesValuesPropertyNameOK {
	return &DeleteEntitiesEntityFidPropertiesValuesPropertyNameOK{}
}

/*DeleteEntitiesEntityFidPropertiesValuesPropertyNameOK handles this case with default header values.

Property Deleted
*/
type DeleteEntitiesEntityFidPropertiesValuesPropertyNameOK struct {
}

func (o *DeleteEntitiesEntityFidPropertiesValuesPropertyNameOK) Error() string {
	return fmt.Sprintf("[DELETE /entities/{entityFid}/properties/values/{propertyName}][%d] deleteEntitiesEntityFidPropertiesValuesPropertyNameOK ", 200)
}

func (o *DeleteEntitiesEntityFidPropertiesValuesPropertyNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
