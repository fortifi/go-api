// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/fortifi/go-api/models"
)

// GetVersionReader is a Reader for the GetVersion structure.
type GetVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetVersionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetVersionOK creates a GetVersionOK with default headers values
func NewGetVersionOK() *GetVersionOK {
	return &GetVersionOK{}
}

/*
GetVersionOK describes a response with status code 200, with default header values.

Version
*/
type GetVersionOK struct {
	Payload *GetVersionOKBody
}

// IsSuccess returns true when this get version o k response has a 2xx status code
func (o *GetVersionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get version o k response has a 3xx status code
func (o *GetVersionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get version o k response has a 4xx status code
func (o *GetVersionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get version o k response has a 5xx status code
func (o *GetVersionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get version o k response a status code equal to that given
func (o *GetVersionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get version o k response
func (o *GetVersionOK) Code() int {
	return 200
}

func (o *GetVersionOK) Error() string {
	return fmt.Sprintf("[GET /version][%d] getVersionOK  %+v", 200, o.Payload)
}

func (o *GetVersionOK) String() string {
	return fmt.Sprintf("[GET /version][%d] getVersionOK  %+v", 200, o.Payload)
}

func (o *GetVersionOK) GetPayload() *GetVersionOKBody {
	return o.Payload
}

func (o *GetVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetVersionOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVersionDefault creates a GetVersionDefault with default headers values
func NewGetVersionDefault(code int) *GetVersionDefault {
	return &GetVersionDefault{
		_statusCode: code,
	}
}

/*
GetVersionDefault describes a response with status code -1, with default header values.

Error
*/
type GetVersionDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// IsSuccess returns true when this get version default response has a 2xx status code
func (o *GetVersionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get version default response has a 3xx status code
func (o *GetVersionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get version default response has a 4xx status code
func (o *GetVersionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get version default response has a 5xx status code
func (o *GetVersionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get version default response a status code equal to that given
func (o *GetVersionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get version default response
func (o *GetVersionDefault) Code() int {
	return o._statusCode
}

func (o *GetVersionDefault) Error() string {
	return fmt.Sprintf("[GET /version][%d] getVersion default  %+v", o._statusCode, o.Payload)
}

func (o *GetVersionDefault) String() string {
	return fmt.Sprintf("[GET /version][%d] getVersion default  %+v", o._statusCode, o.Payload)
}

func (o *GetVersionDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *GetVersionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetVersionOKBody get version o k body
swagger:model GetVersionOKBody
*/
type GetVersionOKBody struct {
	models.Envelope

	// data
	Data string `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetVersionOKBody) UnmarshalJSON(raw []byte) error {
	// GetVersionOKBodyAO0
	var getVersionOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &getVersionOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = getVersionOKBodyAO0

	// GetVersionOKBodyAO1
	var dataGetVersionOKBodyAO1 struct {
		Data string `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataGetVersionOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataGetVersionOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetVersionOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getVersionOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getVersionOKBodyAO0)
	var dataGetVersionOKBodyAO1 struct {
		Data string `json:"data,omitempty"`
	}

	dataGetVersionOKBodyAO1.Data = o.Data

	jsonDataGetVersionOKBodyAO1, errGetVersionOKBodyAO1 := swag.WriteJSON(dataGetVersionOKBodyAO1)
	if errGetVersionOKBodyAO1 != nil {
		return nil, errGetVersionOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetVersionOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get version o k body
func (o *GetVersionOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.Envelope
	if err := o.Envelope.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this get version o k body based on the context it is used
func (o *GetVersionOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.Envelope
	if err := o.Envelope.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *GetVersionOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetVersionOKBody) UnmarshalBinary(b []byte) error {
	var res GetVersionOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
