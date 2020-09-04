// Code generated by go-swagger; DO NOT EDIT.

package deprecated

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

// GetAdvertisersReader is a Reader for the GetAdvertisers structure.
type GetAdvertisersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAdvertisersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAdvertisersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetAdvertisersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAdvertisersOK creates a GetAdvertisersOK with default headers values
func NewGetAdvertisersOK() *GetAdvertisersOK {
	return &GetAdvertisersOK{}
}

/*GetAdvertisersOK handles this case with default header values.

List of advertisers
*/
type GetAdvertisersOK struct {
	Payload *GetAdvertisersOKBody
}

func (o *GetAdvertisersOK) Error() string {
	return fmt.Sprintf("[GET /advertisers][%d] getAdvertisersOK  %+v", 200, o.Payload)
}

func (o *GetAdvertisersOK) GetPayload() *GetAdvertisersOKBody {
	return o.Payload
}

func (o *GetAdvertisersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetAdvertisersOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAdvertisersDefault creates a GetAdvertisersDefault with default headers values
func NewGetAdvertisersDefault(code int) *GetAdvertisersDefault {
	return &GetAdvertisersDefault{
		_statusCode: code,
	}
}

/*GetAdvertisersDefault handles this case with default header values.

Error
*/
type GetAdvertisersDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the get advertisers default response
func (o *GetAdvertisersDefault) Code() int {
	return o._statusCode
}

func (o *GetAdvertisersDefault) Error() string {
	return fmt.Sprintf("[GET /advertisers][%d] GetAdvertisers default  %+v", o._statusCode, o.Payload)
}

func (o *GetAdvertisersDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *GetAdvertisersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetAdvertisersOKBody get advertisers o k body
swagger:model GetAdvertisersOKBody
*/
type GetAdvertisersOKBody struct {
	models.Envelope

	// data
	Data *models.Advertisers `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetAdvertisersOKBody) UnmarshalJSON(raw []byte) error {
	// GetAdvertisersOKBodyAO0
	var getAdvertisersOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &getAdvertisersOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = getAdvertisersOKBodyAO0

	// GetAdvertisersOKBodyAO1
	var dataGetAdvertisersOKBodyAO1 struct {
		Data *models.Advertisers `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataGetAdvertisersOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataGetAdvertisersOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetAdvertisersOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getAdvertisersOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getAdvertisersOKBodyAO0)
	var dataGetAdvertisersOKBodyAO1 struct {
		Data *models.Advertisers `json:"data,omitempty"`
	}

	dataGetAdvertisersOKBodyAO1.Data = o.Data

	jsonDataGetAdvertisersOKBodyAO1, errGetAdvertisersOKBodyAO1 := swag.WriteJSON(dataGetAdvertisersOKBodyAO1)
	if errGetAdvertisersOKBodyAO1 != nil {
		return nil, errGetAdvertisersOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetAdvertisersOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get advertisers o k body
func (o *GetAdvertisersOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetAdvertisersOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getAdvertisersOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetAdvertisersOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAdvertisersOKBody) UnmarshalBinary(b []byte) error {
	var res GetAdvertisersOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
