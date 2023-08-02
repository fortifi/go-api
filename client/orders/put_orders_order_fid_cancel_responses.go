// Code generated by go-swagger; DO NOT EDIT.

package orders

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

// PutOrdersOrderFidCancelReader is a Reader for the PutOrdersOrderFidCancel structure.
type PutOrdersOrderFidCancelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutOrdersOrderFidCancelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutOrdersOrderFidCancelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutOrdersOrderFidCancelDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutOrdersOrderFidCancelOK creates a PutOrdersOrderFidCancelOK with default headers values
func NewPutOrdersOrderFidCancelOK() *PutOrdersOrderFidCancelOK {
	return &PutOrdersOrderFidCancelOK{}
}

/*
PutOrdersOrderFidCancelOK describes a response with status code 200, with default header values.

Order cancelled
*/
type PutOrdersOrderFidCancelOK struct {
	Payload *PutOrdersOrderFidCancelOKBody
}

// IsSuccess returns true when this put orders order fid cancel o k response has a 2xx status code
func (o *PutOrdersOrderFidCancelOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put orders order fid cancel o k response has a 3xx status code
func (o *PutOrdersOrderFidCancelOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put orders order fid cancel o k response has a 4xx status code
func (o *PutOrdersOrderFidCancelOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put orders order fid cancel o k response has a 5xx status code
func (o *PutOrdersOrderFidCancelOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put orders order fid cancel o k response a status code equal to that given
func (o *PutOrdersOrderFidCancelOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the put orders order fid cancel o k response
func (o *PutOrdersOrderFidCancelOK) Code() int {
	return 200
}

func (o *PutOrdersOrderFidCancelOK) Error() string {
	return fmt.Sprintf("[PUT /orders/{orderFid}/cancel][%d] putOrdersOrderFidCancelOK  %+v", 200, o.Payload)
}

func (o *PutOrdersOrderFidCancelOK) String() string {
	return fmt.Sprintf("[PUT /orders/{orderFid}/cancel][%d] putOrdersOrderFidCancelOK  %+v", 200, o.Payload)
}

func (o *PutOrdersOrderFidCancelOK) GetPayload() *PutOrdersOrderFidCancelOKBody {
	return o.Payload
}

func (o *PutOrdersOrderFidCancelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PutOrdersOrderFidCancelOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOrdersOrderFidCancelDefault creates a PutOrdersOrderFidCancelDefault with default headers values
func NewPutOrdersOrderFidCancelDefault(code int) *PutOrdersOrderFidCancelDefault {
	return &PutOrdersOrderFidCancelDefault{
		_statusCode: code,
	}
}

/*
PutOrdersOrderFidCancelDefault describes a response with status code -1, with default header values.

Error
*/
type PutOrdersOrderFidCancelDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// IsSuccess returns true when this put orders order fid cancel default response has a 2xx status code
func (o *PutOrdersOrderFidCancelDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this put orders order fid cancel default response has a 3xx status code
func (o *PutOrdersOrderFidCancelDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this put orders order fid cancel default response has a 4xx status code
func (o *PutOrdersOrderFidCancelDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this put orders order fid cancel default response has a 5xx status code
func (o *PutOrdersOrderFidCancelDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this put orders order fid cancel default response a status code equal to that given
func (o *PutOrdersOrderFidCancelDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the put orders order fid cancel default response
func (o *PutOrdersOrderFidCancelDefault) Code() int {
	return o._statusCode
}

func (o *PutOrdersOrderFidCancelDefault) Error() string {
	return fmt.Sprintf("[PUT /orders/{orderFid}/cancel][%d] PutOrdersOrderFidCancel default  %+v", o._statusCode, o.Payload)
}

func (o *PutOrdersOrderFidCancelDefault) String() string {
	return fmt.Sprintf("[PUT /orders/{orderFid}/cancel][%d] PutOrdersOrderFidCancel default  %+v", o._statusCode, o.Payload)
}

func (o *PutOrdersOrderFidCancelDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PutOrdersOrderFidCancelDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PutOrdersOrderFidCancelOKBody put orders order fid cancel o k body
swagger:model PutOrdersOrderFidCancelOKBody
*/
type PutOrdersOrderFidCancelOKBody struct {
	models.Envelope

	// data
	Data *models.BoolMessage `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PutOrdersOrderFidCancelOKBody) UnmarshalJSON(raw []byte) error {
	// PutOrdersOrderFidCancelOKBodyAO0
	var putOrdersOrderFidCancelOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &putOrdersOrderFidCancelOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = putOrdersOrderFidCancelOKBodyAO0

	// PutOrdersOrderFidCancelOKBodyAO1
	var dataPutOrdersOrderFidCancelOKBodyAO1 struct {
		Data *models.BoolMessage `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataPutOrdersOrderFidCancelOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataPutOrdersOrderFidCancelOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PutOrdersOrderFidCancelOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	putOrdersOrderFidCancelOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, putOrdersOrderFidCancelOKBodyAO0)
	var dataPutOrdersOrderFidCancelOKBodyAO1 struct {
		Data *models.BoolMessage `json:"data,omitempty"`
	}

	dataPutOrdersOrderFidCancelOKBodyAO1.Data = o.Data

	jsonDataPutOrdersOrderFidCancelOKBodyAO1, errPutOrdersOrderFidCancelOKBodyAO1 := swag.WriteJSON(dataPutOrdersOrderFidCancelOKBodyAO1)
	if errPutOrdersOrderFidCancelOKBodyAO1 != nil {
		return nil, errPutOrdersOrderFidCancelOKBodyAO1
	}
	_parts = append(_parts, jsonDataPutOrdersOrderFidCancelOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this put orders order fid cancel o k body
func (o *PutOrdersOrderFidCancelOKBody) Validate(formats strfmt.Registry) error {
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

func (o *PutOrdersOrderFidCancelOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("putOrdersOrderFidCancelOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("putOrdersOrderFidCancelOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this put orders order fid cancel o k body based on the context it is used
func (o *PutOrdersOrderFidCancelOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *PutOrdersOrderFidCancelOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if swag.IsZero(o.Data) { // not required
			return nil
		}

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("putOrdersOrderFidCancelOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("putOrdersOrderFidCancelOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PutOrdersOrderFidCancelOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutOrdersOrderFidCancelOKBody) UnmarshalBinary(b []byte) error {
	var res PutOrdersOrderFidCancelOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
