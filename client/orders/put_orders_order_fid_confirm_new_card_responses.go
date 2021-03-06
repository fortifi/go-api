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

// PutOrdersOrderFidConfirmNewCardReader is a Reader for the PutOrdersOrderFidConfirmNewCard structure.
type PutOrdersOrderFidConfirmNewCardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutOrdersOrderFidConfirmNewCardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutOrdersOrderFidConfirmNewCardOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutOrdersOrderFidConfirmNewCardDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutOrdersOrderFidConfirmNewCardOK creates a PutOrdersOrderFidConfirmNewCardOK with default headers values
func NewPutOrdersOrderFidConfirmNewCardOK() *PutOrdersOrderFidConfirmNewCardOK {
	return &PutOrdersOrderFidConfirmNewCardOK{}
}

/* PutOrdersOrderFidConfirmNewCardOK describes a response with status code 200, with default header values.

Order confirmed and payment authroized
*/
type PutOrdersOrderFidConfirmNewCardOK struct {
	Payload *PutOrdersOrderFidConfirmNewCardOKBody
}

func (o *PutOrdersOrderFidConfirmNewCardOK) Error() string {
	return fmt.Sprintf("[PUT /orders/{orderFid}/confirmNewCard][%d] putOrdersOrderFidConfirmNewCardOK  %+v", 200, o.Payload)
}
func (o *PutOrdersOrderFidConfirmNewCardOK) GetPayload() *PutOrdersOrderFidConfirmNewCardOKBody {
	return o.Payload
}

func (o *PutOrdersOrderFidConfirmNewCardOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PutOrdersOrderFidConfirmNewCardOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOrdersOrderFidConfirmNewCardDefault creates a PutOrdersOrderFidConfirmNewCardDefault with default headers values
func NewPutOrdersOrderFidConfirmNewCardDefault(code int) *PutOrdersOrderFidConfirmNewCardDefault {
	return &PutOrdersOrderFidConfirmNewCardDefault{
		_statusCode: code,
	}
}

/* PutOrdersOrderFidConfirmNewCardDefault describes a response with status code -1, with default header values.

Error
*/
type PutOrdersOrderFidConfirmNewCardDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the put orders order fid confirm new card default response
func (o *PutOrdersOrderFidConfirmNewCardDefault) Code() int {
	return o._statusCode
}

func (o *PutOrdersOrderFidConfirmNewCardDefault) Error() string {
	return fmt.Sprintf("[PUT /orders/{orderFid}/confirmNewCard][%d] PutOrdersOrderFidConfirmNewCard default  %+v", o._statusCode, o.Payload)
}
func (o *PutOrdersOrderFidConfirmNewCardDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PutOrdersOrderFidConfirmNewCardDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PutOrdersOrderFidConfirmNewCardOKBody put orders order fid confirm new card o k body
swagger:model PutOrdersOrderFidConfirmNewCardOKBody
*/
type PutOrdersOrderFidConfirmNewCardOKBody struct {
	models.Envelope

	// data
	Data *models.OrderConfirmation `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PutOrdersOrderFidConfirmNewCardOKBody) UnmarshalJSON(raw []byte) error {
	// PutOrdersOrderFidConfirmNewCardOKBodyAO0
	var putOrdersOrderFidConfirmNewCardOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &putOrdersOrderFidConfirmNewCardOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = putOrdersOrderFidConfirmNewCardOKBodyAO0

	// PutOrdersOrderFidConfirmNewCardOKBodyAO1
	var dataPutOrdersOrderFidConfirmNewCardOKBodyAO1 struct {
		Data *models.OrderConfirmation `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataPutOrdersOrderFidConfirmNewCardOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataPutOrdersOrderFidConfirmNewCardOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PutOrdersOrderFidConfirmNewCardOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	putOrdersOrderFidConfirmNewCardOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, putOrdersOrderFidConfirmNewCardOKBodyAO0)
	var dataPutOrdersOrderFidConfirmNewCardOKBodyAO1 struct {
		Data *models.OrderConfirmation `json:"data,omitempty"`
	}

	dataPutOrdersOrderFidConfirmNewCardOKBodyAO1.Data = o.Data

	jsonDataPutOrdersOrderFidConfirmNewCardOKBodyAO1, errPutOrdersOrderFidConfirmNewCardOKBodyAO1 := swag.WriteJSON(dataPutOrdersOrderFidConfirmNewCardOKBodyAO1)
	if errPutOrdersOrderFidConfirmNewCardOKBodyAO1 != nil {
		return nil, errPutOrdersOrderFidConfirmNewCardOKBodyAO1
	}
	_parts = append(_parts, jsonDataPutOrdersOrderFidConfirmNewCardOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this put orders order fid confirm new card o k body
func (o *PutOrdersOrderFidConfirmNewCardOKBody) Validate(formats strfmt.Registry) error {
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

func (o *PutOrdersOrderFidConfirmNewCardOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("putOrdersOrderFidConfirmNewCardOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this put orders order fid confirm new card o k body based on the context it is used
func (o *PutOrdersOrderFidConfirmNewCardOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *PutOrdersOrderFidConfirmNewCardOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {
		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("putOrdersOrderFidConfirmNewCardOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PutOrdersOrderFidConfirmNewCardOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutOrdersOrderFidConfirmNewCardOKBody) UnmarshalBinary(b []byte) error {
	var res PutOrdersOrderFidConfirmNewCardOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
