// Code generated by go-swagger; DO NOT EDIT.

package customers

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

// PutCustomersCustomerFidChargebacksChargebackFidReader is a Reader for the PutCustomersCustomerFidChargebacksChargebackFid structure.
type PutCustomersCustomerFidChargebacksChargebackFidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutCustomersCustomerFidChargebacksChargebackFidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutCustomersCustomerFidChargebacksChargebackFidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutCustomersCustomerFidChargebacksChargebackFidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutCustomersCustomerFidChargebacksChargebackFidOK creates a PutCustomersCustomerFidChargebacksChargebackFidOK with default headers values
func NewPutCustomersCustomerFidChargebacksChargebackFidOK() *PutCustomersCustomerFidChargebacksChargebackFidOK {
	return &PutCustomersCustomerFidChargebacksChargebackFidOK{}
}

/*
PutCustomersCustomerFidChargebacksChargebackFidOK describes a response with status code 200, with default header values.

Chargeback Actioned
*/
type PutCustomersCustomerFidChargebacksChargebackFidOK struct {
	Payload *PutCustomersCustomerFidChargebacksChargebackFidOKBody
}

// IsSuccess returns true when this put customers customer fid chargebacks chargeback fid o k response has a 2xx status code
func (o *PutCustomersCustomerFidChargebacksChargebackFidOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put customers customer fid chargebacks chargeback fid o k response has a 3xx status code
func (o *PutCustomersCustomerFidChargebacksChargebackFidOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put customers customer fid chargebacks chargeback fid o k response has a 4xx status code
func (o *PutCustomersCustomerFidChargebacksChargebackFidOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put customers customer fid chargebacks chargeback fid o k response has a 5xx status code
func (o *PutCustomersCustomerFidChargebacksChargebackFidOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put customers customer fid chargebacks chargeback fid o k response a status code equal to that given
func (o *PutCustomersCustomerFidChargebacksChargebackFidOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the put customers customer fid chargebacks chargeback fid o k response
func (o *PutCustomersCustomerFidChargebacksChargebackFidOK) Code() int {
	return 200
}

func (o *PutCustomersCustomerFidChargebacksChargebackFidOK) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/chargebacks/{chargebackFid}][%d] putCustomersCustomerFidChargebacksChargebackFidOK  %+v", 200, o.Payload)
}

func (o *PutCustomersCustomerFidChargebacksChargebackFidOK) String() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/chargebacks/{chargebackFid}][%d] putCustomersCustomerFidChargebacksChargebackFidOK  %+v", 200, o.Payload)
}

func (o *PutCustomersCustomerFidChargebacksChargebackFidOK) GetPayload() *PutCustomersCustomerFidChargebacksChargebackFidOKBody {
	return o.Payload
}

func (o *PutCustomersCustomerFidChargebacksChargebackFidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PutCustomersCustomerFidChargebacksChargebackFidOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutCustomersCustomerFidChargebacksChargebackFidDefault creates a PutCustomersCustomerFidChargebacksChargebackFidDefault with default headers values
func NewPutCustomersCustomerFidChargebacksChargebackFidDefault(code int) *PutCustomersCustomerFidChargebacksChargebackFidDefault {
	return &PutCustomersCustomerFidChargebacksChargebackFidDefault{
		_statusCode: code,
	}
}

/*
PutCustomersCustomerFidChargebacksChargebackFidDefault describes a response with status code -1, with default header values.

Error
*/
type PutCustomersCustomerFidChargebacksChargebackFidDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// IsSuccess returns true when this put customers customer fid chargebacks chargeback fid default response has a 2xx status code
func (o *PutCustomersCustomerFidChargebacksChargebackFidDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this put customers customer fid chargebacks chargeback fid default response has a 3xx status code
func (o *PutCustomersCustomerFidChargebacksChargebackFidDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this put customers customer fid chargebacks chargeback fid default response has a 4xx status code
func (o *PutCustomersCustomerFidChargebacksChargebackFidDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this put customers customer fid chargebacks chargeback fid default response has a 5xx status code
func (o *PutCustomersCustomerFidChargebacksChargebackFidDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this put customers customer fid chargebacks chargeback fid default response a status code equal to that given
func (o *PutCustomersCustomerFidChargebacksChargebackFidDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the put customers customer fid chargebacks chargeback fid default response
func (o *PutCustomersCustomerFidChargebacksChargebackFidDefault) Code() int {
	return o._statusCode
}

func (o *PutCustomersCustomerFidChargebacksChargebackFidDefault) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/chargebacks/{chargebackFid}][%d] PutCustomersCustomerFidChargebacksChargebackFid default  %+v", o._statusCode, o.Payload)
}

func (o *PutCustomersCustomerFidChargebacksChargebackFidDefault) String() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/chargebacks/{chargebackFid}][%d] PutCustomersCustomerFidChargebacksChargebackFid default  %+v", o._statusCode, o.Payload)
}

func (o *PutCustomersCustomerFidChargebacksChargebackFidDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PutCustomersCustomerFidChargebacksChargebackFidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PutCustomersCustomerFidChargebacksChargebackFidOKBody put customers customer fid chargebacks chargeback fid o k body
swagger:model PutCustomersCustomerFidChargebacksChargebackFidOKBody
*/
type PutCustomersCustomerFidChargebacksChargebackFidOKBody struct {
	models.Envelope

	// data
	Data *models.BoolMessage `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PutCustomersCustomerFidChargebacksChargebackFidOKBody) UnmarshalJSON(raw []byte) error {
	// PutCustomersCustomerFidChargebacksChargebackFidOKBodyAO0
	var putCustomersCustomerFidChargebacksChargebackFidOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &putCustomersCustomerFidChargebacksChargebackFidOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = putCustomersCustomerFidChargebacksChargebackFidOKBodyAO0

	// PutCustomersCustomerFidChargebacksChargebackFidOKBodyAO1
	var dataPutCustomersCustomerFidChargebacksChargebackFidOKBodyAO1 struct {
		Data *models.BoolMessage `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataPutCustomersCustomerFidChargebacksChargebackFidOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataPutCustomersCustomerFidChargebacksChargebackFidOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PutCustomersCustomerFidChargebacksChargebackFidOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	putCustomersCustomerFidChargebacksChargebackFidOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, putCustomersCustomerFidChargebacksChargebackFidOKBodyAO0)
	var dataPutCustomersCustomerFidChargebacksChargebackFidOKBodyAO1 struct {
		Data *models.BoolMessage `json:"data,omitempty"`
	}

	dataPutCustomersCustomerFidChargebacksChargebackFidOKBodyAO1.Data = o.Data

	jsonDataPutCustomersCustomerFidChargebacksChargebackFidOKBodyAO1, errPutCustomersCustomerFidChargebacksChargebackFidOKBodyAO1 := swag.WriteJSON(dataPutCustomersCustomerFidChargebacksChargebackFidOKBodyAO1)
	if errPutCustomersCustomerFidChargebacksChargebackFidOKBodyAO1 != nil {
		return nil, errPutCustomersCustomerFidChargebacksChargebackFidOKBodyAO1
	}
	_parts = append(_parts, jsonDataPutCustomersCustomerFidChargebacksChargebackFidOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this put customers customer fid chargebacks chargeback fid o k body
func (o *PutCustomersCustomerFidChargebacksChargebackFidOKBody) Validate(formats strfmt.Registry) error {
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

func (o *PutCustomersCustomerFidChargebacksChargebackFidOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("putCustomersCustomerFidChargebacksChargebackFidOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("putCustomersCustomerFidChargebacksChargebackFidOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this put customers customer fid chargebacks chargeback fid o k body based on the context it is used
func (o *PutCustomersCustomerFidChargebacksChargebackFidOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *PutCustomersCustomerFidChargebacksChargebackFidOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if swag.IsZero(o.Data) { // not required
			return nil
		}

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("putCustomersCustomerFidChargebacksChargebackFidOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("putCustomersCustomerFidChargebacksChargebackFidOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PutCustomersCustomerFidChargebacksChargebackFidOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutCustomersCustomerFidChargebacksChargebackFidOKBody) UnmarshalBinary(b []byte) error {
	var res PutCustomersCustomerFidChargebacksChargebackFidOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
