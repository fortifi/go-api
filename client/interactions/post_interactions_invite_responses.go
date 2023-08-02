// Code generated by go-swagger; DO NOT EDIT.

package interactions

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

// PostInteractionsInviteReader is a Reader for the PostInteractionsInvite structure.
type PostInteractionsInviteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostInteractionsInviteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostInteractionsInviteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostInteractionsInviteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostInteractionsInviteOK creates a PostInteractionsInviteOK with default headers values
func NewPostInteractionsInviteOK() *PostInteractionsInviteOK {
	return &PostInteractionsInviteOK{}
}

/*
PostInteractionsInviteOK describes a response with status code 200, with default header values.

Successfully create an interaction invite
*/
type PostInteractionsInviteOK struct {
	Payload *PostInteractionsInviteOKBody
}

// IsSuccess returns true when this post interactions invite o k response has a 2xx status code
func (o *PostInteractionsInviteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post interactions invite o k response has a 3xx status code
func (o *PostInteractionsInviteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post interactions invite o k response has a 4xx status code
func (o *PostInteractionsInviteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post interactions invite o k response has a 5xx status code
func (o *PostInteractionsInviteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post interactions invite o k response a status code equal to that given
func (o *PostInteractionsInviteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post interactions invite o k response
func (o *PostInteractionsInviteOK) Code() int {
	return 200
}

func (o *PostInteractionsInviteOK) Error() string {
	return fmt.Sprintf("[POST /interactions/invite][%d] postInteractionsInviteOK  %+v", 200, o.Payload)
}

func (o *PostInteractionsInviteOK) String() string {
	return fmt.Sprintf("[POST /interactions/invite][%d] postInteractionsInviteOK  %+v", 200, o.Payload)
}

func (o *PostInteractionsInviteOK) GetPayload() *PostInteractionsInviteOKBody {
	return o.Payload
}

func (o *PostInteractionsInviteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostInteractionsInviteOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostInteractionsInviteDefault creates a PostInteractionsInviteDefault with default headers values
func NewPostInteractionsInviteDefault(code int) *PostInteractionsInviteDefault {
	return &PostInteractionsInviteDefault{
		_statusCode: code,
	}
}

/*
PostInteractionsInviteDefault describes a response with status code -1, with default header values.

Error
*/
type PostInteractionsInviteDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// IsSuccess returns true when this post interactions invite default response has a 2xx status code
func (o *PostInteractionsInviteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this post interactions invite default response has a 3xx status code
func (o *PostInteractionsInviteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this post interactions invite default response has a 4xx status code
func (o *PostInteractionsInviteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this post interactions invite default response has a 5xx status code
func (o *PostInteractionsInviteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this post interactions invite default response a status code equal to that given
func (o *PostInteractionsInviteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the post interactions invite default response
func (o *PostInteractionsInviteDefault) Code() int {
	return o._statusCode
}

func (o *PostInteractionsInviteDefault) Error() string {
	return fmt.Sprintf("[POST /interactions/invite][%d] PostInteractionsInvite default  %+v", o._statusCode, o.Payload)
}

func (o *PostInteractionsInviteDefault) String() string {
	return fmt.Sprintf("[POST /interactions/invite][%d] PostInteractionsInvite default  %+v", o._statusCode, o.Payload)
}

func (o *PostInteractionsInviteDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PostInteractionsInviteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PostInteractionsInviteOKBody post interactions invite o k body
swagger:model PostInteractionsInviteOKBody
*/
type PostInteractionsInviteOKBody struct {
	models.Envelope

	// data
	Data *models.InteractionInviteStatusResponse `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostInteractionsInviteOKBody) UnmarshalJSON(raw []byte) error {
	// PostInteractionsInviteOKBodyAO0
	var postInteractionsInviteOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &postInteractionsInviteOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = postInteractionsInviteOKBodyAO0

	// PostInteractionsInviteOKBodyAO1
	var dataPostInteractionsInviteOKBodyAO1 struct {
		Data *models.InteractionInviteStatusResponse `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataPostInteractionsInviteOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataPostInteractionsInviteOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostInteractionsInviteOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	postInteractionsInviteOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postInteractionsInviteOKBodyAO0)
	var dataPostInteractionsInviteOKBodyAO1 struct {
		Data *models.InteractionInviteStatusResponse `json:"data,omitempty"`
	}

	dataPostInteractionsInviteOKBodyAO1.Data = o.Data

	jsonDataPostInteractionsInviteOKBodyAO1, errPostInteractionsInviteOKBodyAO1 := swag.WriteJSON(dataPostInteractionsInviteOKBodyAO1)
	if errPostInteractionsInviteOKBodyAO1 != nil {
		return nil, errPostInteractionsInviteOKBodyAO1
	}
	_parts = append(_parts, jsonDataPostInteractionsInviteOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post interactions invite o k body
func (o *PostInteractionsInviteOKBody) Validate(formats strfmt.Registry) error {
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

func (o *PostInteractionsInviteOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postInteractionsInviteOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postInteractionsInviteOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post interactions invite o k body based on the context it is used
func (o *PostInteractionsInviteOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *PostInteractionsInviteOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if swag.IsZero(o.Data) { // not required
			return nil
		}

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postInteractionsInviteOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postInteractionsInviteOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostInteractionsInviteOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostInteractionsInviteOKBody) UnmarshalBinary(b []byte) error {
	var res PostInteractionsInviteOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
