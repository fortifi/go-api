// Code generated by go-swagger; DO NOT EDIT.

package authentication

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

// GetServiceAuthTokenReader is a Reader for the GetServiceAuthToken structure.
type GetServiceAuthTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServiceAuthTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetServiceAuthTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetServiceAuthTokenDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetServiceAuthTokenOK creates a GetServiceAuthTokenOK with default headers values
func NewGetServiceAuthTokenOK() *GetServiceAuthTokenOK {
	return &GetServiceAuthTokenOK{}
}

/*GetServiceAuthTokenOK handles this case with default header values.

Token
*/
type GetServiceAuthTokenOK struct {
	Payload *GetServiceAuthTokenOKBody
}

func (o *GetServiceAuthTokenOK) Error() string {
	return fmt.Sprintf("[POST /svcauth/verify][%d] getServiceAuthTokenOK  %+v", 200, o.Payload)
}

func (o *GetServiceAuthTokenOK) GetPayload() *GetServiceAuthTokenOKBody {
	return o.Payload
}

func (o *GetServiceAuthTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetServiceAuthTokenOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServiceAuthTokenDefault creates a GetServiceAuthTokenDefault with default headers values
func NewGetServiceAuthTokenDefault(code int) *GetServiceAuthTokenDefault {
	return &GetServiceAuthTokenDefault{
		_statusCode: code,
	}
}

/*GetServiceAuthTokenDefault handles this case with default header values.

Error
*/
type GetServiceAuthTokenDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the get service auth token default response
func (o *GetServiceAuthTokenDefault) Code() int {
	return o._statusCode
}

func (o *GetServiceAuthTokenDefault) Error() string {
	return fmt.Sprintf("[POST /svcauth/verify][%d] getServiceAuthToken default  %+v", o._statusCode, o.Payload)
}

func (o *GetServiceAuthTokenDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *GetServiceAuthTokenDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetServiceAuthTokenOKBody get service auth token o k body
swagger:model GetServiceAuthTokenOKBody
*/
type GetServiceAuthTokenOKBody struct {
	models.Envelope

	// data
	Data *models.AuthToken `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetServiceAuthTokenOKBody) UnmarshalJSON(raw []byte) error {
	// GetServiceAuthTokenOKBodyAO0
	var getServiceAuthTokenOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &getServiceAuthTokenOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = getServiceAuthTokenOKBodyAO0

	// GetServiceAuthTokenOKBodyAO1
	var dataGetServiceAuthTokenOKBodyAO1 struct {
		Data *models.AuthToken `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataGetServiceAuthTokenOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataGetServiceAuthTokenOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetServiceAuthTokenOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getServiceAuthTokenOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getServiceAuthTokenOKBodyAO0)

	var dataGetServiceAuthTokenOKBodyAO1 struct {
		Data *models.AuthToken `json:"data,omitempty"`
	}

	dataGetServiceAuthTokenOKBodyAO1.Data = o.Data

	jsonDataGetServiceAuthTokenOKBodyAO1, errGetServiceAuthTokenOKBodyAO1 := swag.WriteJSON(dataGetServiceAuthTokenOKBodyAO1)
	if errGetServiceAuthTokenOKBodyAO1 != nil {
		return nil, errGetServiceAuthTokenOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetServiceAuthTokenOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get service auth token o k body
func (o *GetServiceAuthTokenOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetServiceAuthTokenOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getServiceAuthTokenOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetServiceAuthTokenOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetServiceAuthTokenOKBody) UnmarshalBinary(b []byte) error {
	var res GetServiceAuthTokenOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}