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

// GetIntegrationsVerifyUserReader is a Reader for the GetIntegrationsVerifyUser structure.
type GetIntegrationsVerifyUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIntegrationsVerifyUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIntegrationsVerifyUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetIntegrationsVerifyUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetIntegrationsVerifyUserOK creates a GetIntegrationsVerifyUserOK with default headers values
func NewGetIntegrationsVerifyUserOK() *GetIntegrationsVerifyUserOK {
	return &GetIntegrationsVerifyUserOK{}
}

/*GetIntegrationsVerifyUserOK handles this case with default header values.

Integration User
*/
type GetIntegrationsVerifyUserOK struct {
	Payload *GetIntegrationsVerifyUserOKBody
}

func (o *GetIntegrationsVerifyUserOK) Error() string {
	return fmt.Sprintf("[GET /integrations/verifyUser][%d] getIntegrationsVerifyUserOK  %+v", 200, o.Payload)
}

func (o *GetIntegrationsVerifyUserOK) GetPayload() *GetIntegrationsVerifyUserOKBody {
	return o.Payload
}

func (o *GetIntegrationsVerifyUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetIntegrationsVerifyUserOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsVerifyUserDefault creates a GetIntegrationsVerifyUserDefault with default headers values
func NewGetIntegrationsVerifyUserDefault(code int) *GetIntegrationsVerifyUserDefault {
	return &GetIntegrationsVerifyUserDefault{
		_statusCode: code,
	}
}

/*GetIntegrationsVerifyUserDefault handles this case with default header values.

Error
*/
type GetIntegrationsVerifyUserDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the get integrations verify user default response
func (o *GetIntegrationsVerifyUserDefault) Code() int {
	return o._statusCode
}

func (o *GetIntegrationsVerifyUserDefault) Error() string {
	return fmt.Sprintf("[GET /integrations/verifyUser][%d] GetIntegrationsVerifyUser default  %+v", o._statusCode, o.Payload)
}

func (o *GetIntegrationsVerifyUserDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *GetIntegrationsVerifyUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetIntegrationsVerifyUserOKBody get integrations verify user o k body
swagger:model GetIntegrationsVerifyUserOKBody
*/
type GetIntegrationsVerifyUserOKBody struct {
	models.Envelope

	// data
	Data *models.IntegrationUser `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetIntegrationsVerifyUserOKBody) UnmarshalJSON(raw []byte) error {
	// GetIntegrationsVerifyUserOKBodyAO0
	var getIntegrationsVerifyUserOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &getIntegrationsVerifyUserOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = getIntegrationsVerifyUserOKBodyAO0

	// GetIntegrationsVerifyUserOKBodyAO1
	var dataGetIntegrationsVerifyUserOKBodyAO1 struct {
		Data *models.IntegrationUser `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataGetIntegrationsVerifyUserOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataGetIntegrationsVerifyUserOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetIntegrationsVerifyUserOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getIntegrationsVerifyUserOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getIntegrationsVerifyUserOKBodyAO0)
	var dataGetIntegrationsVerifyUserOKBodyAO1 struct {
		Data *models.IntegrationUser `json:"data,omitempty"`
	}

	dataGetIntegrationsVerifyUserOKBodyAO1.Data = o.Data

	jsonDataGetIntegrationsVerifyUserOKBodyAO1, errGetIntegrationsVerifyUserOKBodyAO1 := swag.WriteJSON(dataGetIntegrationsVerifyUserOKBodyAO1)
	if errGetIntegrationsVerifyUserOKBodyAO1 != nil {
		return nil, errGetIntegrationsVerifyUserOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetIntegrationsVerifyUserOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get integrations verify user o k body
func (o *GetIntegrationsVerifyUserOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetIntegrationsVerifyUserOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getIntegrationsVerifyUserOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetIntegrationsVerifyUserOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetIntegrationsVerifyUserOKBody) UnmarshalBinary(b []byte) error {
	var res GetIntegrationsVerifyUserOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
