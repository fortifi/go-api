// Code generated by go-swagger; DO NOT EDIT.

package authentication

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

// GetMeReader is a Reader for the GetMe structure.
type GetMeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetMeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMeOK creates a GetMeOK with default headers values
func NewGetMeOK() *GetMeOK {
	return &GetMeOK{}
}

/*GetMeOK handles this case with default header values.

User Information
*/
type GetMeOK struct {
	Payload *GetMeOKBody
}

func (o *GetMeOK) Error() string {
	return fmt.Sprintf("[GET /me][%d] getMeOK  %+v", 200, o.Payload)
}

func (o *GetMeOK) GetPayload() *GetMeOKBody {
	return o.Payload
}

func (o *GetMeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetMeOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMeDefault creates a GetMeDefault with default headers values
func NewGetMeDefault(code int) *GetMeDefault {
	return &GetMeDefault{
		_statusCode: code,
	}
}

/*GetMeDefault handles this case with default header values.

Error
*/
type GetMeDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the get me default response
func (o *GetMeDefault) Code() int {
	return o._statusCode
}

func (o *GetMeDefault) Error() string {
	return fmt.Sprintf("[GET /me][%d] getMe default  %+v", o._statusCode, o.Payload)
}

func (o *GetMeDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *GetMeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetMeOKBody get me o k body
swagger:model GetMeOKBody
*/
type GetMeOKBody struct {
	models.Envelope

	// data
	Data *models.User `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetMeOKBody) UnmarshalJSON(raw []byte) error {
	// GetMeOKBodyAO0
	var getMeOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &getMeOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = getMeOKBodyAO0

	// GetMeOKBodyAO1
	var dataGetMeOKBodyAO1 struct {
		Data *models.User `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataGetMeOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataGetMeOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetMeOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getMeOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getMeOKBodyAO0)
	var dataGetMeOKBodyAO1 struct {
		Data *models.User `json:"data,omitempty"`
	}

	dataGetMeOKBodyAO1.Data = o.Data

	jsonDataGetMeOKBodyAO1, errGetMeOKBodyAO1 := swag.WriteJSON(dataGetMeOKBodyAO1)
	if errGetMeOKBodyAO1 != nil {
		return nil, errGetMeOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetMeOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get me o k body
func (o *GetMeOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetMeOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getMeOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetMeOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetMeOKBody) UnmarshalBinary(b []byte) error {
	var res GetMeOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
