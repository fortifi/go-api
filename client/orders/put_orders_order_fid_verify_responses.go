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

// PutOrdersOrderFidVerifyReader is a Reader for the PutOrdersOrderFidVerify structure.
type PutOrdersOrderFidVerifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutOrdersOrderFidVerifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutOrdersOrderFidVerifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutOrdersOrderFidVerifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutOrdersOrderFidVerifyOK creates a PutOrdersOrderFidVerifyOK with default headers values
func NewPutOrdersOrderFidVerifyOK() *PutOrdersOrderFidVerifyOK {
	return &PutOrdersOrderFidVerifyOK{}
}

/*
PutOrdersOrderFidVerifyOK describes a response with status code 200, with default header values.

Order payment account verified
*/
type PutOrdersOrderFidVerifyOK struct {
	Payload *PutOrdersOrderFidVerifyOKBody
}

// IsSuccess returns true when this put orders order fid verify o k response has a 2xx status code
func (o *PutOrdersOrderFidVerifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put orders order fid verify o k response has a 3xx status code
func (o *PutOrdersOrderFidVerifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put orders order fid verify o k response has a 4xx status code
func (o *PutOrdersOrderFidVerifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put orders order fid verify o k response has a 5xx status code
func (o *PutOrdersOrderFidVerifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put orders order fid verify o k response a status code equal to that given
func (o *PutOrdersOrderFidVerifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the put orders order fid verify o k response
func (o *PutOrdersOrderFidVerifyOK) Code() int {
	return 200
}

func (o *PutOrdersOrderFidVerifyOK) Error() string {
	return fmt.Sprintf("[PUT /orders/{orderFid}/verify][%d] putOrdersOrderFidVerifyOK  %+v", 200, o.Payload)
}

func (o *PutOrdersOrderFidVerifyOK) String() string {
	return fmt.Sprintf("[PUT /orders/{orderFid}/verify][%d] putOrdersOrderFidVerifyOK  %+v", 200, o.Payload)
}

func (o *PutOrdersOrderFidVerifyOK) GetPayload() *PutOrdersOrderFidVerifyOKBody {
	return o.Payload
}

func (o *PutOrdersOrderFidVerifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PutOrdersOrderFidVerifyOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOrdersOrderFidVerifyDefault creates a PutOrdersOrderFidVerifyDefault with default headers values
func NewPutOrdersOrderFidVerifyDefault(code int) *PutOrdersOrderFidVerifyDefault {
	return &PutOrdersOrderFidVerifyDefault{
		_statusCode: code,
	}
}

/*
PutOrdersOrderFidVerifyDefault describes a response with status code -1, with default header values.

Error
*/
type PutOrdersOrderFidVerifyDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// IsSuccess returns true when this put orders order fid verify default response has a 2xx status code
func (o *PutOrdersOrderFidVerifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this put orders order fid verify default response has a 3xx status code
func (o *PutOrdersOrderFidVerifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this put orders order fid verify default response has a 4xx status code
func (o *PutOrdersOrderFidVerifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this put orders order fid verify default response has a 5xx status code
func (o *PutOrdersOrderFidVerifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this put orders order fid verify default response a status code equal to that given
func (o *PutOrdersOrderFidVerifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the put orders order fid verify default response
func (o *PutOrdersOrderFidVerifyDefault) Code() int {
	return o._statusCode
}

func (o *PutOrdersOrderFidVerifyDefault) Error() string {
	return fmt.Sprintf("[PUT /orders/{orderFid}/verify][%d] PutOrdersOrderFidVerify default  %+v", o._statusCode, o.Payload)
}

func (o *PutOrdersOrderFidVerifyDefault) String() string {
	return fmt.Sprintf("[PUT /orders/{orderFid}/verify][%d] PutOrdersOrderFidVerify default  %+v", o._statusCode, o.Payload)
}

func (o *PutOrdersOrderFidVerifyDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PutOrdersOrderFidVerifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PutOrdersOrderFidVerifyOKBody put orders order fid verify o k body
swagger:model PutOrdersOrderFidVerifyOKBody
*/
type PutOrdersOrderFidVerifyOKBody struct {
	models.Envelope

	// data
	Data *models.OrderVerification `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PutOrdersOrderFidVerifyOKBody) UnmarshalJSON(raw []byte) error {
	// PutOrdersOrderFidVerifyOKBodyAO0
	var putOrdersOrderFidVerifyOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &putOrdersOrderFidVerifyOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = putOrdersOrderFidVerifyOKBodyAO0

	// PutOrdersOrderFidVerifyOKBodyAO1
	var dataPutOrdersOrderFidVerifyOKBodyAO1 struct {
		Data *models.OrderVerification `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataPutOrdersOrderFidVerifyOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataPutOrdersOrderFidVerifyOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PutOrdersOrderFidVerifyOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	putOrdersOrderFidVerifyOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, putOrdersOrderFidVerifyOKBodyAO0)
	var dataPutOrdersOrderFidVerifyOKBodyAO1 struct {
		Data *models.OrderVerification `json:"data,omitempty"`
	}

	dataPutOrdersOrderFidVerifyOKBodyAO1.Data = o.Data

	jsonDataPutOrdersOrderFidVerifyOKBodyAO1, errPutOrdersOrderFidVerifyOKBodyAO1 := swag.WriteJSON(dataPutOrdersOrderFidVerifyOKBodyAO1)
	if errPutOrdersOrderFidVerifyOKBodyAO1 != nil {
		return nil, errPutOrdersOrderFidVerifyOKBodyAO1
	}
	_parts = append(_parts, jsonDataPutOrdersOrderFidVerifyOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this put orders order fid verify o k body
func (o *PutOrdersOrderFidVerifyOKBody) Validate(formats strfmt.Registry) error {
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

func (o *PutOrdersOrderFidVerifyOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("putOrdersOrderFidVerifyOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("putOrdersOrderFidVerifyOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this put orders order fid verify o k body based on the context it is used
func (o *PutOrdersOrderFidVerifyOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *PutOrdersOrderFidVerifyOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if swag.IsZero(o.Data) { // not required
			return nil
		}

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("putOrdersOrderFidVerifyOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("putOrdersOrderFidVerifyOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PutOrdersOrderFidVerifyOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutOrdersOrderFidVerifyOKBody) UnmarshalBinary(b []byte) error {
	var res PutOrdersOrderFidVerifyOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
