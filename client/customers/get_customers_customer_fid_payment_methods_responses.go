// Code generated by go-swagger; DO NOT EDIT.

package customers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/fortifi/go-api/models"
)

// GetCustomersCustomerFidPaymentMethodsReader is a Reader for the GetCustomersCustomerFidPaymentMethods structure.
type GetCustomersCustomerFidPaymentMethodsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCustomersCustomerFidPaymentMethodsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCustomersCustomerFidPaymentMethodsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetCustomersCustomerFidPaymentMethodsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCustomersCustomerFidPaymentMethodsOK creates a GetCustomersCustomerFidPaymentMethodsOK with default headers values
func NewGetCustomersCustomerFidPaymentMethodsOK() *GetCustomersCustomerFidPaymentMethodsOK {
	return &GetCustomersCustomerFidPaymentMethodsOK{}
}

/*GetCustomersCustomerFidPaymentMethodsOK handles this case with default header values.

List of payment methods
*/
type GetCustomersCustomerFidPaymentMethodsOK struct {
	Payload *GetCustomersCustomerFidPaymentMethodsOKBody
}

func (o *GetCustomersCustomerFidPaymentMethodsOK) Error() string {
	return fmt.Sprintf("[GET /customers/{customerFid}/paymentMethods][%d] getCustomersCustomerFidPaymentMethodsOK  %+v", 200, o.Payload)
}

func (o *GetCustomersCustomerFidPaymentMethodsOK) GetPayload() *GetCustomersCustomerFidPaymentMethodsOKBody {
	return o.Payload
}

func (o *GetCustomersCustomerFidPaymentMethodsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetCustomersCustomerFidPaymentMethodsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCustomersCustomerFidPaymentMethodsDefault creates a GetCustomersCustomerFidPaymentMethodsDefault with default headers values
func NewGetCustomersCustomerFidPaymentMethodsDefault(code int) *GetCustomersCustomerFidPaymentMethodsDefault {
	return &GetCustomersCustomerFidPaymentMethodsDefault{
		_statusCode: code,
	}
}

/*GetCustomersCustomerFidPaymentMethodsDefault handles this case with default header values.

Error
*/
type GetCustomersCustomerFidPaymentMethodsDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the get customers customer fid payment methods default response
func (o *GetCustomersCustomerFidPaymentMethodsDefault) Code() int {
	return o._statusCode
}

func (o *GetCustomersCustomerFidPaymentMethodsDefault) Error() string {
	return fmt.Sprintf("[GET /customers/{customerFid}/paymentMethods][%d] GetCustomersCustomerFidPaymentMethods default  %+v", o._statusCode, o.Payload)
}

func (o *GetCustomersCustomerFidPaymentMethodsDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *GetCustomersCustomerFidPaymentMethodsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetCustomersCustomerFidPaymentMethodsOKBody get customers customer fid payment methods o k body
swagger:model GetCustomersCustomerFidPaymentMethodsOKBody
*/
type GetCustomersCustomerFidPaymentMethodsOKBody struct {
	models.Envelope

	// data
	Data *models.PaymentMethods `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetCustomersCustomerFidPaymentMethodsOKBody) UnmarshalJSON(raw []byte) error {
	// GetCustomersCustomerFidPaymentMethodsOKBodyAO0
	var getCustomersCustomerFidPaymentMethodsOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &getCustomersCustomerFidPaymentMethodsOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = getCustomersCustomerFidPaymentMethodsOKBodyAO0

	// GetCustomersCustomerFidPaymentMethodsOKBodyAO1
	var dataGetCustomersCustomerFidPaymentMethodsOKBodyAO1 struct {
		Data *models.PaymentMethods `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataGetCustomersCustomerFidPaymentMethodsOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataGetCustomersCustomerFidPaymentMethodsOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetCustomersCustomerFidPaymentMethodsOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getCustomersCustomerFidPaymentMethodsOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getCustomersCustomerFidPaymentMethodsOKBodyAO0)
	var dataGetCustomersCustomerFidPaymentMethodsOKBodyAO1 struct {
		Data *models.PaymentMethods `json:"data,omitempty"`
	}

	dataGetCustomersCustomerFidPaymentMethodsOKBodyAO1.Data = o.Data

	jsonDataGetCustomersCustomerFidPaymentMethodsOKBodyAO1, errGetCustomersCustomerFidPaymentMethodsOKBodyAO1 := swag.WriteJSON(dataGetCustomersCustomerFidPaymentMethodsOKBodyAO1)
	if errGetCustomersCustomerFidPaymentMethodsOKBodyAO1 != nil {
		return nil, errGetCustomersCustomerFidPaymentMethodsOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetCustomersCustomerFidPaymentMethodsOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get customers customer fid payment methods o k body
func (o *GetCustomersCustomerFidPaymentMethodsOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetCustomersCustomerFidPaymentMethodsOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getCustomersCustomerFidPaymentMethodsOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetCustomersCustomerFidPaymentMethodsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCustomersCustomerFidPaymentMethodsOKBody) UnmarshalBinary(b []byte) error {
	var res GetCustomersCustomerFidPaymentMethodsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
