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

// PostInteractionsInteractionFidCustomerClosedReader is a Reader for the PostInteractionsInteractionFidCustomerClosed structure.
type PostInteractionsInteractionFidCustomerClosedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostInteractionsInteractionFidCustomerClosedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostInteractionsInteractionFidCustomerClosedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostInteractionsInteractionFidCustomerClosedDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostInteractionsInteractionFidCustomerClosedOK creates a PostInteractionsInteractionFidCustomerClosedOK with default headers values
func NewPostInteractionsInteractionFidCustomerClosedOK() *PostInteractionsInteractionFidCustomerClosedOK {
	return &PostInteractionsInteractionFidCustomerClosedOK{}
}

/*
PostInteractionsInteractionFidCustomerClosedOK describes a response with status code 200, with default header values.

Interaction Closed Response
*/
type PostInteractionsInteractionFidCustomerClosedOK struct {
	Payload *PostInteractionsInteractionFidCustomerClosedOKBody
}

// IsSuccess returns true when this post interactions interaction fid customer closed o k response has a 2xx status code
func (o *PostInteractionsInteractionFidCustomerClosedOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post interactions interaction fid customer closed o k response has a 3xx status code
func (o *PostInteractionsInteractionFidCustomerClosedOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post interactions interaction fid customer closed o k response has a 4xx status code
func (o *PostInteractionsInteractionFidCustomerClosedOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post interactions interaction fid customer closed o k response has a 5xx status code
func (o *PostInteractionsInteractionFidCustomerClosedOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post interactions interaction fid customer closed o k response a status code equal to that given
func (o *PostInteractionsInteractionFidCustomerClosedOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post interactions interaction fid customer closed o k response
func (o *PostInteractionsInteractionFidCustomerClosedOK) Code() int {
	return 200
}

func (o *PostInteractionsInteractionFidCustomerClosedOK) Error() string {
	return fmt.Sprintf("[POST /interactions/{interactionFid}/customerClosed][%d] postInteractionsInteractionFidCustomerClosedOK  %+v", 200, o.Payload)
}

func (o *PostInteractionsInteractionFidCustomerClosedOK) String() string {
	return fmt.Sprintf("[POST /interactions/{interactionFid}/customerClosed][%d] postInteractionsInteractionFidCustomerClosedOK  %+v", 200, o.Payload)
}

func (o *PostInteractionsInteractionFidCustomerClosedOK) GetPayload() *PostInteractionsInteractionFidCustomerClosedOKBody {
	return o.Payload
}

func (o *PostInteractionsInteractionFidCustomerClosedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostInteractionsInteractionFidCustomerClosedOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostInteractionsInteractionFidCustomerClosedDefault creates a PostInteractionsInteractionFidCustomerClosedDefault with default headers values
func NewPostInteractionsInteractionFidCustomerClosedDefault(code int) *PostInteractionsInteractionFidCustomerClosedDefault {
	return &PostInteractionsInteractionFidCustomerClosedDefault{
		_statusCode: code,
	}
}

/*
PostInteractionsInteractionFidCustomerClosedDefault describes a response with status code -1, with default header values.

Error
*/
type PostInteractionsInteractionFidCustomerClosedDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// IsSuccess returns true when this post interactions interaction fid customer closed default response has a 2xx status code
func (o *PostInteractionsInteractionFidCustomerClosedDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this post interactions interaction fid customer closed default response has a 3xx status code
func (o *PostInteractionsInteractionFidCustomerClosedDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this post interactions interaction fid customer closed default response has a 4xx status code
func (o *PostInteractionsInteractionFidCustomerClosedDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this post interactions interaction fid customer closed default response has a 5xx status code
func (o *PostInteractionsInteractionFidCustomerClosedDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this post interactions interaction fid customer closed default response a status code equal to that given
func (o *PostInteractionsInteractionFidCustomerClosedDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the post interactions interaction fid customer closed default response
func (o *PostInteractionsInteractionFidCustomerClosedDefault) Code() int {
	return o._statusCode
}

func (o *PostInteractionsInteractionFidCustomerClosedDefault) Error() string {
	return fmt.Sprintf("[POST /interactions/{interactionFid}/customerClosed][%d] PostInteractionsInteractionFidCustomerClosed default  %+v", o._statusCode, o.Payload)
}

func (o *PostInteractionsInteractionFidCustomerClosedDefault) String() string {
	return fmt.Sprintf("[POST /interactions/{interactionFid}/customerClosed][%d] PostInteractionsInteractionFidCustomerClosed default  %+v", o._statusCode, o.Payload)
}

func (o *PostInteractionsInteractionFidCustomerClosedDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PostInteractionsInteractionFidCustomerClosedDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PostInteractionsInteractionFidCustomerClosedOKBody post interactions interaction fid customer closed o k body
swagger:model PostInteractionsInteractionFidCustomerClosedOKBody
*/
type PostInteractionsInteractionFidCustomerClosedOKBody struct {
	models.Envelope

	// data
	Data *models.BoolMessage `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostInteractionsInteractionFidCustomerClosedOKBody) UnmarshalJSON(raw []byte) error {
	// PostInteractionsInteractionFidCustomerClosedOKBodyAO0
	var postInteractionsInteractionFidCustomerClosedOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &postInteractionsInteractionFidCustomerClosedOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = postInteractionsInteractionFidCustomerClosedOKBodyAO0

	// PostInteractionsInteractionFidCustomerClosedOKBodyAO1
	var dataPostInteractionsInteractionFidCustomerClosedOKBodyAO1 struct {
		Data *models.BoolMessage `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataPostInteractionsInteractionFidCustomerClosedOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataPostInteractionsInteractionFidCustomerClosedOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostInteractionsInteractionFidCustomerClosedOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	postInteractionsInteractionFidCustomerClosedOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postInteractionsInteractionFidCustomerClosedOKBodyAO0)
	var dataPostInteractionsInteractionFidCustomerClosedOKBodyAO1 struct {
		Data *models.BoolMessage `json:"data,omitempty"`
	}

	dataPostInteractionsInteractionFidCustomerClosedOKBodyAO1.Data = o.Data

	jsonDataPostInteractionsInteractionFidCustomerClosedOKBodyAO1, errPostInteractionsInteractionFidCustomerClosedOKBodyAO1 := swag.WriteJSON(dataPostInteractionsInteractionFidCustomerClosedOKBodyAO1)
	if errPostInteractionsInteractionFidCustomerClosedOKBodyAO1 != nil {
		return nil, errPostInteractionsInteractionFidCustomerClosedOKBodyAO1
	}
	_parts = append(_parts, jsonDataPostInteractionsInteractionFidCustomerClosedOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post interactions interaction fid customer closed o k body
func (o *PostInteractionsInteractionFidCustomerClosedOKBody) Validate(formats strfmt.Registry) error {
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

func (o *PostInteractionsInteractionFidCustomerClosedOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postInteractionsInteractionFidCustomerClosedOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postInteractionsInteractionFidCustomerClosedOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post interactions interaction fid customer closed o k body based on the context it is used
func (o *PostInteractionsInteractionFidCustomerClosedOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *PostInteractionsInteractionFidCustomerClosedOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if swag.IsZero(o.Data) { // not required
			return nil
		}

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postInteractionsInteractionFidCustomerClosedOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postInteractionsInteractionFidCustomerClosedOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostInteractionsInteractionFidCustomerClosedOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostInteractionsInteractionFidCustomerClosedOKBody) UnmarshalBinary(b []byte) error {
	var res PostInteractionsInteractionFidCustomerClosedOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
