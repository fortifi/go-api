// Code generated by go-swagger; DO NOT EDIT.

package entities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/fortifi/go-api/models"
)

// GetEntitiesEntityFidPropertiesValuesPropertyNameReader is a Reader for the GetEntitiesEntityFidPropertiesValuesPropertyName structure.
type GetEntitiesEntityFidPropertiesValuesPropertyNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEntitiesEntityFidPropertiesValuesPropertyNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEntitiesEntityFidPropertiesValuesPropertyNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetEntitiesEntityFidPropertiesValuesPropertyNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetEntitiesEntityFidPropertiesValuesPropertyNameOK creates a GetEntitiesEntityFidPropertiesValuesPropertyNameOK with default headers values
func NewGetEntitiesEntityFidPropertiesValuesPropertyNameOK() *GetEntitiesEntityFidPropertiesValuesPropertyNameOK {
	return &GetEntitiesEntityFidPropertiesValuesPropertyNameOK{}
}

/*GetEntitiesEntityFidPropertiesValuesPropertyNameOK handles this case with default header values.

Property Value
*/
type GetEntitiesEntityFidPropertiesValuesPropertyNameOK struct {
	Payload *GetEntitiesEntityFidPropertiesValuesPropertyNameOKBody
}

func (o *GetEntitiesEntityFidPropertiesValuesPropertyNameOK) Error() string {
	return fmt.Sprintf("[GET /entities/{entityFid}/properties/values/{propertyName}][%d] getEntitiesEntityFidPropertiesValuesPropertyNameOK  %+v", 200, o.Payload)
}

func (o *GetEntitiesEntityFidPropertiesValuesPropertyNameOK) GetPayload() *GetEntitiesEntityFidPropertiesValuesPropertyNameOKBody {
	return o.Payload
}

func (o *GetEntitiesEntityFidPropertiesValuesPropertyNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetEntitiesEntityFidPropertiesValuesPropertyNameOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEntitiesEntityFidPropertiesValuesPropertyNameDefault creates a GetEntitiesEntityFidPropertiesValuesPropertyNameDefault with default headers values
func NewGetEntitiesEntityFidPropertiesValuesPropertyNameDefault(code int) *GetEntitiesEntityFidPropertiesValuesPropertyNameDefault {
	return &GetEntitiesEntityFidPropertiesValuesPropertyNameDefault{
		_statusCode: code,
	}
}

/*GetEntitiesEntityFidPropertiesValuesPropertyNameDefault handles this case with default header values.

Error
*/
type GetEntitiesEntityFidPropertiesValuesPropertyNameDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the get entities entity fid properties values property name default response
func (o *GetEntitiesEntityFidPropertiesValuesPropertyNameDefault) Code() int {
	return o._statusCode
}

func (o *GetEntitiesEntityFidPropertiesValuesPropertyNameDefault) Error() string {
	return fmt.Sprintf("[GET /entities/{entityFid}/properties/values/{propertyName}][%d] GetEntitiesEntityFidPropertiesValuesPropertyName default  %+v", o._statusCode, o.Payload)
}

func (o *GetEntitiesEntityFidPropertiesValuesPropertyNameDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *GetEntitiesEntityFidPropertiesValuesPropertyNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetEntitiesEntityFidPropertiesValuesPropertyNameOKBody get entities entity fid properties values property name o k body
swagger:model GetEntitiesEntityFidPropertiesValuesPropertyNameOKBody
*/
type GetEntitiesEntityFidPropertiesValuesPropertyNameOKBody struct {
	models.Envelope

	// data
	Data *models.PropertyValue `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetEntitiesEntityFidPropertiesValuesPropertyNameOKBody) UnmarshalJSON(raw []byte) error {
	// GetEntitiesEntityFidPropertiesValuesPropertyNameOKBodyAO0
	var getEntitiesEntityFidPropertiesValuesPropertyNameOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &getEntitiesEntityFidPropertiesValuesPropertyNameOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = getEntitiesEntityFidPropertiesValuesPropertyNameOKBodyAO0

	// GetEntitiesEntityFidPropertiesValuesPropertyNameOKBodyAO1
	var dataGetEntitiesEntityFidPropertiesValuesPropertyNameOKBodyAO1 struct {
		Data *models.PropertyValue `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataGetEntitiesEntityFidPropertiesValuesPropertyNameOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataGetEntitiesEntityFidPropertiesValuesPropertyNameOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetEntitiesEntityFidPropertiesValuesPropertyNameOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getEntitiesEntityFidPropertiesValuesPropertyNameOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getEntitiesEntityFidPropertiesValuesPropertyNameOKBodyAO0)

	var dataGetEntitiesEntityFidPropertiesValuesPropertyNameOKBodyAO1 struct {
		Data *models.PropertyValue `json:"data,omitempty"`
	}

	dataGetEntitiesEntityFidPropertiesValuesPropertyNameOKBodyAO1.Data = o.Data

	jsonDataGetEntitiesEntityFidPropertiesValuesPropertyNameOKBodyAO1, errGetEntitiesEntityFidPropertiesValuesPropertyNameOKBodyAO1 := swag.WriteJSON(dataGetEntitiesEntityFidPropertiesValuesPropertyNameOKBodyAO1)
	if errGetEntitiesEntityFidPropertiesValuesPropertyNameOKBodyAO1 != nil {
		return nil, errGetEntitiesEntityFidPropertiesValuesPropertyNameOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetEntitiesEntityFidPropertiesValuesPropertyNameOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get entities entity fid properties values property name o k body
func (o *GetEntitiesEntityFidPropertiesValuesPropertyNameOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.Envelope
	if err := o.Envelope.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetEntitiesEntityFidPropertiesValuesPropertyNameOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getEntitiesEntityFidPropertiesValuesPropertyNameOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetEntitiesEntityFidPropertiesValuesPropertyNameOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetEntitiesEntityFidPropertiesValuesPropertyNameOKBody) UnmarshalBinary(b []byte) error {
	var res GetEntitiesEntityFidPropertiesValuesPropertyNameOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}