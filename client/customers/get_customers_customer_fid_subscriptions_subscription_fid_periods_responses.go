// Code generated by go-swagger; DO NOT EDIT.

package customers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/fortifi/go-api/models"
)

// GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsReader is a Reader for the GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriods structure.
type GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsOK creates a GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsOK with default headers values
func NewGetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsOK() *GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsOK {
	return &GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsOK{}
}

/*GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsOK handles this case with default header values.

A purchase period
*/
type GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsOK struct {
	Payload *GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsOKBody
}

func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsOK) Error() string {
	return fmt.Sprintf("[GET /customers/{customerFid}/subscriptions/{subscriptionFid}/periods][%d] getCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsOK  %+v", 200, o.Payload)
}

func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsOK) GetPayload() *GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsOKBody {
	return o.Payload
}

func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsDefault creates a GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsDefault with default headers values
func NewGetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsDefault(code int) *GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsDefault {
	return &GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsDefault{
		_statusCode: code,
	}
}

/*GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsDefault handles this case with default header values.

Error
*/
type GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the get customers customer fid subscriptions subscription fid periods default response
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsDefault) Code() int {
	return o._statusCode
}

func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsDefault) Error() string {
	return fmt.Sprintf("[GET /customers/{customerFid}/subscriptions/{subscriptionFid}/periods][%d] GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriods default  %+v", o._statusCode, o.Payload)
}

func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsOKBody get customers customer fid subscriptions subscription fid periods o k body
swagger:model GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsOKBody
*/
type GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsOKBody struct {
	models.Envelope

	// data
	Data *models.Periods `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsOKBody) UnmarshalJSON(raw []byte) error {
	// GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsOKBodyAO0
	var getCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &getCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = getCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsOKBodyAO0

	// GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsOKBodyAO1
	var dataGetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsOKBodyAO1 struct {
		Data *models.Periods `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataGetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataGetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsOKBodyAO0)

	var dataGetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsOKBodyAO1 struct {
		Data *models.Periods `json:"data,omitempty"`
	}

	dataGetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsOKBodyAO1.Data = o.Data

	jsonDataGetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsOKBodyAO1, errGetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsOKBodyAO1 := swag.WriteJSON(dataGetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsOKBodyAO1)
	if errGetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsOKBodyAO1 != nil {
		return nil, errGetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get customers customer fid subscriptions subscription fid periods o k body
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsOKBody) UnmarshalBinary(b []byte) error {
	var res GetCustomersCustomerFidSubscriptionsSubscriptionFidPeriodsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
