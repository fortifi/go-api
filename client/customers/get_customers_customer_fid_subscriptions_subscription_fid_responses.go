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

// GetCustomersCustomerFidSubscriptionsSubscriptionFidReader is a Reader for the GetCustomersCustomerFidSubscriptionsSubscriptionFid structure.
type GetCustomersCustomerFidSubscriptionsSubscriptionFidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCustomersCustomerFidSubscriptionsSubscriptionFidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetCustomersCustomerFidSubscriptionsSubscriptionFidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCustomersCustomerFidSubscriptionsSubscriptionFidOK creates a GetCustomersCustomerFidSubscriptionsSubscriptionFidOK with default headers values
func NewGetCustomersCustomerFidSubscriptionsSubscriptionFidOK() *GetCustomersCustomerFidSubscriptionsSubscriptionFidOK {
	return &GetCustomersCustomerFidSubscriptionsSubscriptionFidOK{}
}

/*GetCustomersCustomerFidSubscriptionsSubscriptionFidOK handles this case with default header values.

Loaded subscription
*/
type GetCustomersCustomerFidSubscriptionsSubscriptionFidOK struct {
	Payload *GetCustomersCustomerFidSubscriptionsSubscriptionFidOKBody
}

func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidOK) Error() string {
	return fmt.Sprintf("[GET /customers/{customerFid}/subscriptions/{subscriptionFid}][%d] getCustomersCustomerFidSubscriptionsSubscriptionFidOK  %+v", 200, o.Payload)
}

func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidOK) GetPayload() *GetCustomersCustomerFidSubscriptionsSubscriptionFidOKBody {
	return o.Payload
}

func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetCustomersCustomerFidSubscriptionsSubscriptionFidOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCustomersCustomerFidSubscriptionsSubscriptionFidDefault creates a GetCustomersCustomerFidSubscriptionsSubscriptionFidDefault with default headers values
func NewGetCustomersCustomerFidSubscriptionsSubscriptionFidDefault(code int) *GetCustomersCustomerFidSubscriptionsSubscriptionFidDefault {
	return &GetCustomersCustomerFidSubscriptionsSubscriptionFidDefault{
		_statusCode: code,
	}
}

/*GetCustomersCustomerFidSubscriptionsSubscriptionFidDefault handles this case with default header values.

Error
*/
type GetCustomersCustomerFidSubscriptionsSubscriptionFidDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the get customers customer fid subscriptions subscription fid default response
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidDefault) Code() int {
	return o._statusCode
}

func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidDefault) Error() string {
	return fmt.Sprintf("[GET /customers/{customerFid}/subscriptions/{subscriptionFid}][%d] GetCustomersCustomerFidSubscriptionsSubscriptionFid default  %+v", o._statusCode, o.Payload)
}

func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetCustomersCustomerFidSubscriptionsSubscriptionFidOKBody get customers customer fid subscriptions subscription fid o k body
swagger:model GetCustomersCustomerFidSubscriptionsSubscriptionFidOKBody
*/
type GetCustomersCustomerFidSubscriptionsSubscriptionFidOKBody struct {
	models.Envelope

	// data
	Data *models.Subscription `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidOKBody) UnmarshalJSON(raw []byte) error {
	// GetCustomersCustomerFidSubscriptionsSubscriptionFidOKBodyAO0
	var getCustomersCustomerFidSubscriptionsSubscriptionFidOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &getCustomersCustomerFidSubscriptionsSubscriptionFidOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = getCustomersCustomerFidSubscriptionsSubscriptionFidOKBodyAO0

	// GetCustomersCustomerFidSubscriptionsSubscriptionFidOKBodyAO1
	var dataGetCustomersCustomerFidSubscriptionsSubscriptionFidOKBodyAO1 struct {
		Data *models.Subscription `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataGetCustomersCustomerFidSubscriptionsSubscriptionFidOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataGetCustomersCustomerFidSubscriptionsSubscriptionFidOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetCustomersCustomerFidSubscriptionsSubscriptionFidOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getCustomersCustomerFidSubscriptionsSubscriptionFidOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getCustomersCustomerFidSubscriptionsSubscriptionFidOKBodyAO0)
	var dataGetCustomersCustomerFidSubscriptionsSubscriptionFidOKBodyAO1 struct {
		Data *models.Subscription `json:"data,omitempty"`
	}

	dataGetCustomersCustomerFidSubscriptionsSubscriptionFidOKBodyAO1.Data = o.Data

	jsonDataGetCustomersCustomerFidSubscriptionsSubscriptionFidOKBodyAO1, errGetCustomersCustomerFidSubscriptionsSubscriptionFidOKBodyAO1 := swag.WriteJSON(dataGetCustomersCustomerFidSubscriptionsSubscriptionFidOKBodyAO1)
	if errGetCustomersCustomerFidSubscriptionsSubscriptionFidOKBodyAO1 != nil {
		return nil, errGetCustomersCustomerFidSubscriptionsSubscriptionFidOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetCustomersCustomerFidSubscriptionsSubscriptionFidOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get customers customer fid subscriptions subscription fid o k body
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getCustomersCustomerFidSubscriptionsSubscriptionFidOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidOKBody) UnmarshalBinary(b []byte) error {
	var res GetCustomersCustomerFidSubscriptionsSubscriptionFidOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
