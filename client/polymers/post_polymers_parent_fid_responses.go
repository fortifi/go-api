// Code generated by go-swagger; DO NOT EDIT.

package polymers

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

// PostPolymersParentFidReader is a Reader for the PostPolymersParentFid structure.
type PostPolymersParentFidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostPolymersParentFidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostPolymersParentFidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostPolymersParentFidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostPolymersParentFidOK creates a PostPolymersParentFidOK with default headers values
func NewPostPolymersParentFidOK() *PostPolymersParentFidOK {
	return &PostPolymersParentFidOK{}
}

/* PostPolymersParentFidOK describes a response with status code 200, with default header values.

Polymer created
*/
type PostPolymersParentFidOK struct {
	Payload *PostPolymersParentFidOKBody
}

func (o *PostPolymersParentFidOK) Error() string {
	return fmt.Sprintf("[POST /polymers/{parentFid}][%d] postPolymersParentFidOK  %+v", 200, o.Payload)
}
func (o *PostPolymersParentFidOK) GetPayload() *PostPolymersParentFidOKBody {
	return o.Payload
}

func (o *PostPolymersParentFidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostPolymersParentFidOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPolymersParentFidDefault creates a PostPolymersParentFidDefault with default headers values
func NewPostPolymersParentFidDefault(code int) *PostPolymersParentFidDefault {
	return &PostPolymersParentFidDefault{
		_statusCode: code,
	}
}

/* PostPolymersParentFidDefault describes a response with status code -1, with default header values.

Error
*/
type PostPolymersParentFidDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the post polymers parent fid default response
func (o *PostPolymersParentFidDefault) Code() int {
	return o._statusCode
}

func (o *PostPolymersParentFidDefault) Error() string {
	return fmt.Sprintf("[POST /polymers/{parentFid}][%d] PostPolymersParentFid default  %+v", o._statusCode, o.Payload)
}
func (o *PostPolymersParentFidDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PostPolymersParentFidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostPolymersParentFidOKBody post polymers parent fid o k body
swagger:model PostPolymersParentFidOKBody
*/
type PostPolymersParentFidOKBody struct {
	models.Envelope

	// data
	Data *models.Fid `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostPolymersParentFidOKBody) UnmarshalJSON(raw []byte) error {
	// PostPolymersParentFidOKBodyAO0
	var postPolymersParentFidOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &postPolymersParentFidOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = postPolymersParentFidOKBodyAO0

	// PostPolymersParentFidOKBodyAO1
	var dataPostPolymersParentFidOKBodyAO1 struct {
		Data *models.Fid `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataPostPolymersParentFidOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataPostPolymersParentFidOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostPolymersParentFidOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	postPolymersParentFidOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postPolymersParentFidOKBodyAO0)
	var dataPostPolymersParentFidOKBodyAO1 struct {
		Data *models.Fid `json:"data,omitempty"`
	}

	dataPostPolymersParentFidOKBodyAO1.Data = o.Data

	jsonDataPostPolymersParentFidOKBodyAO1, errPostPolymersParentFidOKBodyAO1 := swag.WriteJSON(dataPostPolymersParentFidOKBodyAO1)
	if errPostPolymersParentFidOKBodyAO1 != nil {
		return nil, errPostPolymersParentFidOKBodyAO1
	}
	_parts = append(_parts, jsonDataPostPolymersParentFidOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post polymers parent fid o k body
func (o *PostPolymersParentFidOKBody) Validate(formats strfmt.Registry) error {
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

func (o *PostPolymersParentFidOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postPolymersParentFidOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post polymers parent fid o k body based on the context it is used
func (o *PostPolymersParentFidOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.Envelope
	if err := o.Envelope.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostPolymersParentFidOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {
		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postPolymersParentFidOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostPolymersParentFidOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostPolymersParentFidOKBody) UnmarshalBinary(b []byte) error {
	var res PostPolymersParentFidOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}