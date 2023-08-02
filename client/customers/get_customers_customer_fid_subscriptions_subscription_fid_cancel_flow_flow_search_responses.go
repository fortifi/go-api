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

// GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchReader is a Reader for the GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearch structure.
type GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchOK creates a GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchOK with default headers values
func NewGetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchOK() *GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchOK {
	return &GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchOK{}
}

/*
GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchOK describes a response with status code 200, with default header values.

Cancel Flow State
*/
type GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchOK struct {
	Payload *GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchOKBody
}

// IsSuccess returns true when this get customers customer fid subscriptions subscription fid cancel flow flow search o k response has a 2xx status code
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get customers customer fid subscriptions subscription fid cancel flow flow search o k response has a 3xx status code
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get customers customer fid subscriptions subscription fid cancel flow flow search o k response has a 4xx status code
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get customers customer fid subscriptions subscription fid cancel flow flow search o k response has a 5xx status code
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get customers customer fid subscriptions subscription fid cancel flow flow search o k response a status code equal to that given
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get customers customer fid subscriptions subscription fid cancel flow flow search o k response
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchOK) Code() int {
	return 200
}

func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchOK) Error() string {
	return fmt.Sprintf("[GET /customers/{customerFid}/subscriptions/{subscriptionFid}/cancelFlow/{flowSearch}][%d] getCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchOK  %+v", 200, o.Payload)
}

func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchOK) String() string {
	return fmt.Sprintf("[GET /customers/{customerFid}/subscriptions/{subscriptionFid}/cancelFlow/{flowSearch}][%d] getCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchOK  %+v", 200, o.Payload)
}

func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchOK) GetPayload() *GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchOKBody {
	return o.Payload
}

func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchDefault creates a GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchDefault with default headers values
func NewGetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchDefault(code int) *GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchDefault {
	return &GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchDefault{
		_statusCode: code,
	}
}

/*
GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchDefault describes a response with status code -1, with default header values.

Error
*/
type GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// IsSuccess returns true when this get customers customer fid subscriptions subscription fid cancel flow flow search default response has a 2xx status code
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get customers customer fid subscriptions subscription fid cancel flow flow search default response has a 3xx status code
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get customers customer fid subscriptions subscription fid cancel flow flow search default response has a 4xx status code
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get customers customer fid subscriptions subscription fid cancel flow flow search default response has a 5xx status code
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get customers customer fid subscriptions subscription fid cancel flow flow search default response a status code equal to that given
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get customers customer fid subscriptions subscription fid cancel flow flow search default response
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchDefault) Code() int {
	return o._statusCode
}

func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchDefault) Error() string {
	return fmt.Sprintf("[GET /customers/{customerFid}/subscriptions/{subscriptionFid}/cancelFlow/{flowSearch}][%d] GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearch default  %+v", o._statusCode, o.Payload)
}

func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchDefault) String() string {
	return fmt.Sprintf("[GET /customers/{customerFid}/subscriptions/{subscriptionFid}/cancelFlow/{flowSearch}][%d] GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearch default  %+v", o._statusCode, o.Payload)
}

func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchOKBody get customers customer fid subscriptions subscription fid cancel flow flow search o k body
swagger:model GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchOKBody
*/
type GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchOKBody struct {
	models.Envelope

	// data
	Data *models.CancelFlowState `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchOKBody) UnmarshalJSON(raw []byte) error {
	// GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchOKBodyAO0
	var getCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &getCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = getCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchOKBodyAO0

	// GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchOKBodyAO1
	var dataGetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchOKBodyAO1 struct {
		Data *models.CancelFlowState `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataGetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataGetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchOKBodyAO0)
	var dataGetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchOKBodyAO1 struct {
		Data *models.CancelFlowState `json:"data,omitempty"`
	}

	dataGetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchOKBodyAO1.Data = o.Data

	jsonDataGetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchOKBodyAO1, errGetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchOKBodyAO1 := swag.WriteJSON(dataGetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchOKBodyAO1)
	if errGetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchOKBodyAO1 != nil {
		return nil, errGetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get customers customer fid subscriptions subscription fid cancel flow flow search o k body
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get customers customer fid subscriptions subscription fid cancel flow flow search o k body based on the context it is used
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if swag.IsZero(o.Data) { // not required
			return nil
		}

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchOKBody) UnmarshalBinary(b []byte) error {
	var res GetCustomersCustomerFidSubscriptionsSubscriptionFidCancelFlowFlowSearchOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
