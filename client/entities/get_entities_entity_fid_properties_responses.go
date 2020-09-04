// Code generated by go-swagger; DO NOT EDIT.

package entities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/fortifi/go-api/models"
)

// GetEntitiesEntityFidPropertiesReader is a Reader for the GetEntitiesEntityFidProperties structure.
type GetEntitiesEntityFidPropertiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEntitiesEntityFidPropertiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEntitiesEntityFidPropertiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetEntitiesEntityFidPropertiesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetEntitiesEntityFidPropertiesOK creates a GetEntitiesEntityFidPropertiesOK with default headers values
func NewGetEntitiesEntityFidPropertiesOK() *GetEntitiesEntityFidPropertiesOK {
	return &GetEntitiesEntityFidPropertiesOK{}
}

/*GetEntitiesEntityFidPropertiesOK handles this case with default header values.

Entity Properties
*/
type GetEntitiesEntityFidPropertiesOK struct {
	Payload *GetEntitiesEntityFidPropertiesOKBody
}

func (o *GetEntitiesEntityFidPropertiesOK) Error() string {
	return fmt.Sprintf("[GET /entities/{entityFid}/properties][%d] getEntitiesEntityFidPropertiesOK  %+v", 200, o.Payload)
}

func (o *GetEntitiesEntityFidPropertiesOK) GetPayload() *GetEntitiesEntityFidPropertiesOKBody {
	return o.Payload
}

func (o *GetEntitiesEntityFidPropertiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetEntitiesEntityFidPropertiesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEntitiesEntityFidPropertiesDefault creates a GetEntitiesEntityFidPropertiesDefault with default headers values
func NewGetEntitiesEntityFidPropertiesDefault(code int) *GetEntitiesEntityFidPropertiesDefault {
	return &GetEntitiesEntityFidPropertiesDefault{
		_statusCode: code,
	}
}

/*GetEntitiesEntityFidPropertiesDefault handles this case with default header values.

Error
*/
type GetEntitiesEntityFidPropertiesDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the get entities entity fid properties default response
func (o *GetEntitiesEntityFidPropertiesDefault) Code() int {
	return o._statusCode
}

func (o *GetEntitiesEntityFidPropertiesDefault) Error() string {
	return fmt.Sprintf("[GET /entities/{entityFid}/properties][%d] GetEntitiesEntityFidProperties default  %+v", o._statusCode, o.Payload)
}

func (o *GetEntitiesEntityFidPropertiesDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *GetEntitiesEntityFidPropertiesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetEntitiesEntityFidPropertiesOKBody get entities entity fid properties o k body
swagger:model GetEntitiesEntityFidPropertiesOKBody
*/
type GetEntitiesEntityFidPropertiesOKBody struct {
	models.Envelope

	// data
	Data *models.Properties `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetEntitiesEntityFidPropertiesOKBody) UnmarshalJSON(raw []byte) error {
	// GetEntitiesEntityFidPropertiesOKBodyAO0
	var getEntitiesEntityFidPropertiesOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &getEntitiesEntityFidPropertiesOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = getEntitiesEntityFidPropertiesOKBodyAO0

	// GetEntitiesEntityFidPropertiesOKBodyAO1
	var dataGetEntitiesEntityFidPropertiesOKBodyAO1 struct {
		Data *models.Properties `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataGetEntitiesEntityFidPropertiesOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataGetEntitiesEntityFidPropertiesOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetEntitiesEntityFidPropertiesOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getEntitiesEntityFidPropertiesOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getEntitiesEntityFidPropertiesOKBodyAO0)
	var dataGetEntitiesEntityFidPropertiesOKBodyAO1 struct {
		Data *models.Properties `json:"data,omitempty"`
	}

	dataGetEntitiesEntityFidPropertiesOKBodyAO1.Data = o.Data

	jsonDataGetEntitiesEntityFidPropertiesOKBodyAO1, errGetEntitiesEntityFidPropertiesOKBodyAO1 := swag.WriteJSON(dataGetEntitiesEntityFidPropertiesOKBodyAO1)
	if errGetEntitiesEntityFidPropertiesOKBodyAO1 != nil {
		return nil, errGetEntitiesEntityFidPropertiesOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetEntitiesEntityFidPropertiesOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get entities entity fid properties o k body
func (o *GetEntitiesEntityFidPropertiesOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetEntitiesEntityFidPropertiesOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getEntitiesEntityFidPropertiesOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetEntitiesEntityFidPropertiesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetEntitiesEntityFidPropertiesOKBody) UnmarshalBinary(b []byte) error {
	var res GetEntitiesEntityFidPropertiesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
