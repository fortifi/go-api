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

// PutOrdersOrderFidItemsProductFidOrderItemFidReader is a Reader for the PutOrdersOrderFidItemsProductFidOrderItemFid structure.
type PutOrdersOrderFidItemsProductFidOrderItemFidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutOrdersOrderFidItemsProductFidOrderItemFidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutOrdersOrderFidItemsProductFidOrderItemFidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutOrdersOrderFidItemsProductFidOrderItemFidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutOrdersOrderFidItemsProductFidOrderItemFidOK creates a PutOrdersOrderFidItemsProductFidOrderItemFidOK with default headers values
func NewPutOrdersOrderFidItemsProductFidOrderItemFidOK() *PutOrdersOrderFidItemsProductFidOrderItemFidOK {
	return &PutOrdersOrderFidItemsProductFidOrderItemFidOK{}
}

/*
PutOrdersOrderFidItemsProductFidOrderItemFidOK describes a response with status code 200, with default header values.

Order retrieved
*/
type PutOrdersOrderFidItemsProductFidOrderItemFidOK struct {
	Payload *PutOrdersOrderFidItemsProductFidOrderItemFidOKBody
}

// IsSuccess returns true when this put orders order fid items product fid order item fid o k response has a 2xx status code
func (o *PutOrdersOrderFidItemsProductFidOrderItemFidOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put orders order fid items product fid order item fid o k response has a 3xx status code
func (o *PutOrdersOrderFidItemsProductFidOrderItemFidOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put orders order fid items product fid order item fid o k response has a 4xx status code
func (o *PutOrdersOrderFidItemsProductFidOrderItemFidOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put orders order fid items product fid order item fid o k response has a 5xx status code
func (o *PutOrdersOrderFidItemsProductFidOrderItemFidOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put orders order fid items product fid order item fid o k response a status code equal to that given
func (o *PutOrdersOrderFidItemsProductFidOrderItemFidOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the put orders order fid items product fid order item fid o k response
func (o *PutOrdersOrderFidItemsProductFidOrderItemFidOK) Code() int {
	return 200
}

func (o *PutOrdersOrderFidItemsProductFidOrderItemFidOK) Error() string {
	return fmt.Sprintf("[PUT /orders/{orderFid}/items/{productFid}/{orderItemFid}][%d] putOrdersOrderFidItemsProductFidOrderItemFidOK  %+v", 200, o.Payload)
}

func (o *PutOrdersOrderFidItemsProductFidOrderItemFidOK) String() string {
	return fmt.Sprintf("[PUT /orders/{orderFid}/items/{productFid}/{orderItemFid}][%d] putOrdersOrderFidItemsProductFidOrderItemFidOK  %+v", 200, o.Payload)
}

func (o *PutOrdersOrderFidItemsProductFidOrderItemFidOK) GetPayload() *PutOrdersOrderFidItemsProductFidOrderItemFidOKBody {
	return o.Payload
}

func (o *PutOrdersOrderFidItemsProductFidOrderItemFidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PutOrdersOrderFidItemsProductFidOrderItemFidOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOrdersOrderFidItemsProductFidOrderItemFidDefault creates a PutOrdersOrderFidItemsProductFidOrderItemFidDefault with default headers values
func NewPutOrdersOrderFidItemsProductFidOrderItemFidDefault(code int) *PutOrdersOrderFidItemsProductFidOrderItemFidDefault {
	return &PutOrdersOrderFidItemsProductFidOrderItemFidDefault{
		_statusCode: code,
	}
}

/*
PutOrdersOrderFidItemsProductFidOrderItemFidDefault describes a response with status code -1, with default header values.

Error
*/
type PutOrdersOrderFidItemsProductFidOrderItemFidDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// IsSuccess returns true when this put orders order fid items product fid order item fid default response has a 2xx status code
func (o *PutOrdersOrderFidItemsProductFidOrderItemFidDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this put orders order fid items product fid order item fid default response has a 3xx status code
func (o *PutOrdersOrderFidItemsProductFidOrderItemFidDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this put orders order fid items product fid order item fid default response has a 4xx status code
func (o *PutOrdersOrderFidItemsProductFidOrderItemFidDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this put orders order fid items product fid order item fid default response has a 5xx status code
func (o *PutOrdersOrderFidItemsProductFidOrderItemFidDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this put orders order fid items product fid order item fid default response a status code equal to that given
func (o *PutOrdersOrderFidItemsProductFidOrderItemFidDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the put orders order fid items product fid order item fid default response
func (o *PutOrdersOrderFidItemsProductFidOrderItemFidDefault) Code() int {
	return o._statusCode
}

func (o *PutOrdersOrderFidItemsProductFidOrderItemFidDefault) Error() string {
	return fmt.Sprintf("[PUT /orders/{orderFid}/items/{productFid}/{orderItemFid}][%d] PutOrdersOrderFidItemsProductFidOrderItemFid default  %+v", o._statusCode, o.Payload)
}

func (o *PutOrdersOrderFidItemsProductFidOrderItemFidDefault) String() string {
	return fmt.Sprintf("[PUT /orders/{orderFid}/items/{productFid}/{orderItemFid}][%d] PutOrdersOrderFidItemsProductFidOrderItemFid default  %+v", o._statusCode, o.Payload)
}

func (o *PutOrdersOrderFidItemsProductFidOrderItemFidDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PutOrdersOrderFidItemsProductFidOrderItemFidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PutOrdersOrderFidItemsProductFidOrderItemFidOKBody put orders order fid items product fid order item fid o k body
swagger:model PutOrdersOrderFidItemsProductFidOrderItemFidOKBody
*/
type PutOrdersOrderFidItemsProductFidOrderItemFidOKBody struct {
	models.Envelope

	// data
	Data *models.BoolMessage `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PutOrdersOrderFidItemsProductFidOrderItemFidOKBody) UnmarshalJSON(raw []byte) error {
	// PutOrdersOrderFidItemsProductFidOrderItemFidOKBodyAO0
	var putOrdersOrderFidItemsProductFidOrderItemFidOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &putOrdersOrderFidItemsProductFidOrderItemFidOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = putOrdersOrderFidItemsProductFidOrderItemFidOKBodyAO0

	// PutOrdersOrderFidItemsProductFidOrderItemFidOKBodyAO1
	var dataPutOrdersOrderFidItemsProductFidOrderItemFidOKBodyAO1 struct {
		Data *models.BoolMessage `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataPutOrdersOrderFidItemsProductFidOrderItemFidOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataPutOrdersOrderFidItemsProductFidOrderItemFidOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PutOrdersOrderFidItemsProductFidOrderItemFidOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	putOrdersOrderFidItemsProductFidOrderItemFidOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, putOrdersOrderFidItemsProductFidOrderItemFidOKBodyAO0)
	var dataPutOrdersOrderFidItemsProductFidOrderItemFidOKBodyAO1 struct {
		Data *models.BoolMessage `json:"data,omitempty"`
	}

	dataPutOrdersOrderFidItemsProductFidOrderItemFidOKBodyAO1.Data = o.Data

	jsonDataPutOrdersOrderFidItemsProductFidOrderItemFidOKBodyAO1, errPutOrdersOrderFidItemsProductFidOrderItemFidOKBodyAO1 := swag.WriteJSON(dataPutOrdersOrderFidItemsProductFidOrderItemFidOKBodyAO1)
	if errPutOrdersOrderFidItemsProductFidOrderItemFidOKBodyAO1 != nil {
		return nil, errPutOrdersOrderFidItemsProductFidOrderItemFidOKBodyAO1
	}
	_parts = append(_parts, jsonDataPutOrdersOrderFidItemsProductFidOrderItemFidOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this put orders order fid items product fid order item fid o k body
func (o *PutOrdersOrderFidItemsProductFidOrderItemFidOKBody) Validate(formats strfmt.Registry) error {
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

func (o *PutOrdersOrderFidItemsProductFidOrderItemFidOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("putOrdersOrderFidItemsProductFidOrderItemFidOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("putOrdersOrderFidItemsProductFidOrderItemFidOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this put orders order fid items product fid order item fid o k body based on the context it is used
func (o *PutOrdersOrderFidItemsProductFidOrderItemFidOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *PutOrdersOrderFidItemsProductFidOrderItemFidOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if swag.IsZero(o.Data) { // not required
			return nil
		}

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("putOrdersOrderFidItemsProductFidOrderItemFidOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("putOrdersOrderFidItemsProductFidOrderItemFidOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PutOrdersOrderFidItemsProductFidOrderItemFidOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutOrdersOrderFidItemsProductFidOrderItemFidOKBody) UnmarshalBinary(b []byte) error {
	var res PutOrdersOrderFidItemsProductFidOrderItemFidOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
