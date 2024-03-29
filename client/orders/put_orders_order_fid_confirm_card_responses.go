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

// PutOrdersOrderFidConfirmCardReader is a Reader for the PutOrdersOrderFidConfirmCard structure.
type PutOrdersOrderFidConfirmCardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutOrdersOrderFidConfirmCardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutOrdersOrderFidConfirmCardOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutOrdersOrderFidConfirmCardDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutOrdersOrderFidConfirmCardOK creates a PutOrdersOrderFidConfirmCardOK with default headers values
func NewPutOrdersOrderFidConfirmCardOK() *PutOrdersOrderFidConfirmCardOK {
	return &PutOrdersOrderFidConfirmCardOK{}
}

/*
PutOrdersOrderFidConfirmCardOK describes a response with status code 200, with default header values.

Order confirmed and payment authroized
*/
type PutOrdersOrderFidConfirmCardOK struct {
	Payload *PutOrdersOrderFidConfirmCardOKBody
}

// IsSuccess returns true when this put orders order fid confirm card o k response has a 2xx status code
func (o *PutOrdersOrderFidConfirmCardOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put orders order fid confirm card o k response has a 3xx status code
func (o *PutOrdersOrderFidConfirmCardOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put orders order fid confirm card o k response has a 4xx status code
func (o *PutOrdersOrderFidConfirmCardOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put orders order fid confirm card o k response has a 5xx status code
func (o *PutOrdersOrderFidConfirmCardOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put orders order fid confirm card o k response a status code equal to that given
func (o *PutOrdersOrderFidConfirmCardOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the put orders order fid confirm card o k response
func (o *PutOrdersOrderFidConfirmCardOK) Code() int {
	return 200
}

func (o *PutOrdersOrderFidConfirmCardOK) Error() string {
	return fmt.Sprintf("[PUT /orders/{orderFid}/confirmCard][%d] putOrdersOrderFidConfirmCardOK  %+v", 200, o.Payload)
}

func (o *PutOrdersOrderFidConfirmCardOK) String() string {
	return fmt.Sprintf("[PUT /orders/{orderFid}/confirmCard][%d] putOrdersOrderFidConfirmCardOK  %+v", 200, o.Payload)
}

func (o *PutOrdersOrderFidConfirmCardOK) GetPayload() *PutOrdersOrderFidConfirmCardOKBody {
	return o.Payload
}

func (o *PutOrdersOrderFidConfirmCardOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PutOrdersOrderFidConfirmCardOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOrdersOrderFidConfirmCardDefault creates a PutOrdersOrderFidConfirmCardDefault with default headers values
func NewPutOrdersOrderFidConfirmCardDefault(code int) *PutOrdersOrderFidConfirmCardDefault {
	return &PutOrdersOrderFidConfirmCardDefault{
		_statusCode: code,
	}
}

/*
PutOrdersOrderFidConfirmCardDefault describes a response with status code -1, with default header values.

Error
*/
type PutOrdersOrderFidConfirmCardDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// IsSuccess returns true when this put orders order fid confirm card default response has a 2xx status code
func (o *PutOrdersOrderFidConfirmCardDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this put orders order fid confirm card default response has a 3xx status code
func (o *PutOrdersOrderFidConfirmCardDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this put orders order fid confirm card default response has a 4xx status code
func (o *PutOrdersOrderFidConfirmCardDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this put orders order fid confirm card default response has a 5xx status code
func (o *PutOrdersOrderFidConfirmCardDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this put orders order fid confirm card default response a status code equal to that given
func (o *PutOrdersOrderFidConfirmCardDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the put orders order fid confirm card default response
func (o *PutOrdersOrderFidConfirmCardDefault) Code() int {
	return o._statusCode
}

func (o *PutOrdersOrderFidConfirmCardDefault) Error() string {
	return fmt.Sprintf("[PUT /orders/{orderFid}/confirmCard][%d] PutOrdersOrderFidConfirmCard default  %+v", o._statusCode, o.Payload)
}

func (o *PutOrdersOrderFidConfirmCardDefault) String() string {
	return fmt.Sprintf("[PUT /orders/{orderFid}/confirmCard][%d] PutOrdersOrderFidConfirmCard default  %+v", o._statusCode, o.Payload)
}

func (o *PutOrdersOrderFidConfirmCardDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PutOrdersOrderFidConfirmCardDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PutOrdersOrderFidConfirmCardOKBody put orders order fid confirm card o k body
swagger:model PutOrdersOrderFidConfirmCardOKBody
*/
type PutOrdersOrderFidConfirmCardOKBody struct {
	models.Envelope

	// data
	Data *models.OrderConfirmation `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PutOrdersOrderFidConfirmCardOKBody) UnmarshalJSON(raw []byte) error {
	// PutOrdersOrderFidConfirmCardOKBodyAO0
	var putOrdersOrderFidConfirmCardOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &putOrdersOrderFidConfirmCardOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = putOrdersOrderFidConfirmCardOKBodyAO0

	// PutOrdersOrderFidConfirmCardOKBodyAO1
	var dataPutOrdersOrderFidConfirmCardOKBodyAO1 struct {
		Data *models.OrderConfirmation `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataPutOrdersOrderFidConfirmCardOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataPutOrdersOrderFidConfirmCardOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PutOrdersOrderFidConfirmCardOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	putOrdersOrderFidConfirmCardOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, putOrdersOrderFidConfirmCardOKBodyAO0)
	var dataPutOrdersOrderFidConfirmCardOKBodyAO1 struct {
		Data *models.OrderConfirmation `json:"data,omitempty"`
	}

	dataPutOrdersOrderFidConfirmCardOKBodyAO1.Data = o.Data

	jsonDataPutOrdersOrderFidConfirmCardOKBodyAO1, errPutOrdersOrderFidConfirmCardOKBodyAO1 := swag.WriteJSON(dataPutOrdersOrderFidConfirmCardOKBodyAO1)
	if errPutOrdersOrderFidConfirmCardOKBodyAO1 != nil {
		return nil, errPutOrdersOrderFidConfirmCardOKBodyAO1
	}
	_parts = append(_parts, jsonDataPutOrdersOrderFidConfirmCardOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this put orders order fid confirm card o k body
func (o *PutOrdersOrderFidConfirmCardOKBody) Validate(formats strfmt.Registry) error {
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

func (o *PutOrdersOrderFidConfirmCardOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("putOrdersOrderFidConfirmCardOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("putOrdersOrderFidConfirmCardOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this put orders order fid confirm card o k body based on the context it is used
func (o *PutOrdersOrderFidConfirmCardOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *PutOrdersOrderFidConfirmCardOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if swag.IsZero(o.Data) { // not required
			return nil
		}

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("putOrdersOrderFidConfirmCardOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("putOrdersOrderFidConfirmCardOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PutOrdersOrderFidConfirmCardOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutOrdersOrderFidConfirmCardOKBody) UnmarshalBinary(b []byte) error {
	var res PutOrdersOrderFidConfirmCardOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
