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

// PostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsReader is a Reader for the PostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptions structure.
type PostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsOK creates a PostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsOK with default headers values
func NewPostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsOK() *PostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsOK {
	return &PostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsOK{}
}

/*
PostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsOK describes a response with status code 200, with default header values.

Create Retention Flow Options
*/
type PostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsOK struct {
	Payload *PostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsOKBody
}

// IsSuccess returns true when this post customers customer fid subscriptions subscription fid retention flow flow search options o k response has a 2xx status code
func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post customers customer fid subscriptions subscription fid retention flow flow search options o k response has a 3xx status code
func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post customers customer fid subscriptions subscription fid retention flow flow search options o k response has a 4xx status code
func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post customers customer fid subscriptions subscription fid retention flow flow search options o k response has a 5xx status code
func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post customers customer fid subscriptions subscription fid retention flow flow search options o k response a status code equal to that given
func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post customers customer fid subscriptions subscription fid retention flow flow search options o k response
func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsOK) Code() int {
	return 200
}

func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsOK) Error() string {
	return fmt.Sprintf("[POST /customers/{customerFid}/subscriptions/{subscriptionFid}/retentionFlow/{flowSearch}/options][%d] postCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsOK  %+v", 200, o.Payload)
}

func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsOK) String() string {
	return fmt.Sprintf("[POST /customers/{customerFid}/subscriptions/{subscriptionFid}/retentionFlow/{flowSearch}/options][%d] postCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsOK  %+v", 200, o.Payload)
}

func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsOK) GetPayload() *PostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsOKBody {
	return o.Payload
}

func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsDefault creates a PostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsDefault with default headers values
func NewPostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsDefault(code int) *PostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsDefault {
	return &PostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsDefault{
		_statusCode: code,
	}
}

/*
PostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsDefault describes a response with status code -1, with default header values.

Error
*/
type PostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// IsSuccess returns true when this post customers customer fid subscriptions subscription fid retention flow flow search options default response has a 2xx status code
func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this post customers customer fid subscriptions subscription fid retention flow flow search options default response has a 3xx status code
func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this post customers customer fid subscriptions subscription fid retention flow flow search options default response has a 4xx status code
func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this post customers customer fid subscriptions subscription fid retention flow flow search options default response has a 5xx status code
func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this post customers customer fid subscriptions subscription fid retention flow flow search options default response a status code equal to that given
func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the post customers customer fid subscriptions subscription fid retention flow flow search options default response
func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsDefault) Code() int {
	return o._statusCode
}

func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsDefault) Error() string {
	return fmt.Sprintf("[POST /customers/{customerFid}/subscriptions/{subscriptionFid}/retentionFlow/{flowSearch}/options][%d] PostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptions default  %+v", o._statusCode, o.Payload)
}

func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsDefault) String() string {
	return fmt.Sprintf("[POST /customers/{customerFid}/subscriptions/{subscriptionFid}/retentionFlow/{flowSearch}/options][%d] PostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptions default  %+v", o._statusCode, o.Payload)
}

func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsOKBody post customers customer fid subscriptions subscription fid retention flow flow search options o k body
swagger:model PostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsOKBody
*/
type PostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsOKBody struct {
	models.Envelope

	// data
	Data *models.RetentionFlowOptions `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsOKBody) UnmarshalJSON(raw []byte) error {
	// PostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsOKBodyAO0
	var postCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &postCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = postCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsOKBodyAO0

	// PostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsOKBodyAO1
	var dataPostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsOKBodyAO1 struct {
		Data *models.RetentionFlowOptions `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataPostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataPostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	postCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsOKBodyAO0)
	var dataPostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsOKBodyAO1 struct {
		Data *models.RetentionFlowOptions `json:"data,omitempty"`
	}

	dataPostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsOKBodyAO1.Data = o.Data

	jsonDataPostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsOKBodyAO1, errPostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsOKBodyAO1 := swag.WriteJSON(dataPostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsOKBodyAO1)
	if errPostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsOKBodyAO1 != nil {
		return nil, errPostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsOKBodyAO1
	}
	_parts = append(_parts, jsonDataPostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post customers customer fid subscriptions subscription fid retention flow flow search options o k body
func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsOKBody) Validate(formats strfmt.Registry) error {
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

func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post customers customer fid subscriptions subscription fid retention flow flow search options o k body based on the context it is used
func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if swag.IsZero(o.Data) { // not required
			return nil
		}

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsOKBody) UnmarshalBinary(b []byte) error {
	var res PostCustomersCustomerFidSubscriptionsSubscriptionFidRetentionFlowFlowSearchOptionsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
