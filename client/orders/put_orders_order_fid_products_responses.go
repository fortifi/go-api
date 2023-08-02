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

// PutOrdersOrderFidProductsReader is a Reader for the PutOrdersOrderFidProducts structure.
type PutOrdersOrderFidProductsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutOrdersOrderFidProductsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutOrdersOrderFidProductsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutOrdersOrderFidProductsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutOrdersOrderFidProductsOK creates a PutOrdersOrderFidProductsOK with default headers values
func NewPutOrdersOrderFidProductsOK() *PutOrdersOrderFidProductsOK {
	return &PutOrdersOrderFidProductsOK{}
}

/*
PutOrdersOrderFidProductsOK describes a response with status code 200, with default header values.

Order products successfully updated
*/
type PutOrdersOrderFidProductsOK struct {
	Payload *PutOrdersOrderFidProductsOKBody
}

// IsSuccess returns true when this put orders order fid products o k response has a 2xx status code
func (o *PutOrdersOrderFidProductsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put orders order fid products o k response has a 3xx status code
func (o *PutOrdersOrderFidProductsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put orders order fid products o k response has a 4xx status code
func (o *PutOrdersOrderFidProductsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put orders order fid products o k response has a 5xx status code
func (o *PutOrdersOrderFidProductsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put orders order fid products o k response a status code equal to that given
func (o *PutOrdersOrderFidProductsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the put orders order fid products o k response
func (o *PutOrdersOrderFidProductsOK) Code() int {
	return 200
}

func (o *PutOrdersOrderFidProductsOK) Error() string {
	return fmt.Sprintf("[PUT /orders/{orderFid}/products][%d] putOrdersOrderFidProductsOK  %+v", 200, o.Payload)
}

func (o *PutOrdersOrderFidProductsOK) String() string {
	return fmt.Sprintf("[PUT /orders/{orderFid}/products][%d] putOrdersOrderFidProductsOK  %+v", 200, o.Payload)
}

func (o *PutOrdersOrderFidProductsOK) GetPayload() *PutOrdersOrderFidProductsOKBody {
	return o.Payload
}

func (o *PutOrdersOrderFidProductsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PutOrdersOrderFidProductsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOrdersOrderFidProductsDefault creates a PutOrdersOrderFidProductsDefault with default headers values
func NewPutOrdersOrderFidProductsDefault(code int) *PutOrdersOrderFidProductsDefault {
	return &PutOrdersOrderFidProductsDefault{
		_statusCode: code,
	}
}

/*
PutOrdersOrderFidProductsDefault describes a response with status code -1, with default header values.

Error
*/
type PutOrdersOrderFidProductsDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// IsSuccess returns true when this put orders order fid products default response has a 2xx status code
func (o *PutOrdersOrderFidProductsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this put orders order fid products default response has a 3xx status code
func (o *PutOrdersOrderFidProductsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this put orders order fid products default response has a 4xx status code
func (o *PutOrdersOrderFidProductsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this put orders order fid products default response has a 5xx status code
func (o *PutOrdersOrderFidProductsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this put orders order fid products default response a status code equal to that given
func (o *PutOrdersOrderFidProductsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the put orders order fid products default response
func (o *PutOrdersOrderFidProductsDefault) Code() int {
	return o._statusCode
}

func (o *PutOrdersOrderFidProductsDefault) Error() string {
	return fmt.Sprintf("[PUT /orders/{orderFid}/products][%d] PutOrdersOrderFidProducts default  %+v", o._statusCode, o.Payload)
}

func (o *PutOrdersOrderFidProductsDefault) String() string {
	return fmt.Sprintf("[PUT /orders/{orderFid}/products][%d] PutOrdersOrderFidProducts default  %+v", o._statusCode, o.Payload)
}

func (o *PutOrdersOrderFidProductsDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PutOrdersOrderFidProductsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PutOrdersOrderFidProductsOKBody put orders order fid products o k body
swagger:model PutOrdersOrderFidProductsOKBody
*/
type PutOrdersOrderFidProductsOKBody struct {
	models.Envelope

	// data
	Data *models.OrderAddProducts `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PutOrdersOrderFidProductsOKBody) UnmarshalJSON(raw []byte) error {
	// PutOrdersOrderFidProductsOKBodyAO0
	var putOrdersOrderFidProductsOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &putOrdersOrderFidProductsOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = putOrdersOrderFidProductsOKBodyAO0

	// PutOrdersOrderFidProductsOKBodyAO1
	var dataPutOrdersOrderFidProductsOKBodyAO1 struct {
		Data *models.OrderAddProducts `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataPutOrdersOrderFidProductsOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataPutOrdersOrderFidProductsOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PutOrdersOrderFidProductsOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	putOrdersOrderFidProductsOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, putOrdersOrderFidProductsOKBodyAO0)
	var dataPutOrdersOrderFidProductsOKBodyAO1 struct {
		Data *models.OrderAddProducts `json:"data,omitempty"`
	}

	dataPutOrdersOrderFidProductsOKBodyAO1.Data = o.Data

	jsonDataPutOrdersOrderFidProductsOKBodyAO1, errPutOrdersOrderFidProductsOKBodyAO1 := swag.WriteJSON(dataPutOrdersOrderFidProductsOKBodyAO1)
	if errPutOrdersOrderFidProductsOKBodyAO1 != nil {
		return nil, errPutOrdersOrderFidProductsOKBodyAO1
	}
	_parts = append(_parts, jsonDataPutOrdersOrderFidProductsOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this put orders order fid products o k body
func (o *PutOrdersOrderFidProductsOKBody) Validate(formats strfmt.Registry) error {
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

func (o *PutOrdersOrderFidProductsOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("putOrdersOrderFidProductsOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("putOrdersOrderFidProductsOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this put orders order fid products o k body based on the context it is used
func (o *PutOrdersOrderFidProductsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *PutOrdersOrderFidProductsOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if swag.IsZero(o.Data) { // not required
			return nil
		}

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("putOrdersOrderFidProductsOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("putOrdersOrderFidProductsOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PutOrdersOrderFidProductsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutOrdersOrderFidProductsOKBody) UnmarshalBinary(b []byte) error {
	var res PutOrdersOrderFidProductsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
