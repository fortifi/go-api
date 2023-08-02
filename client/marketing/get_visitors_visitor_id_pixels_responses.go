// Code generated by go-swagger; DO NOT EDIT.

package marketing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/fortifi/go-api/models"
)

// GetVisitorsVisitorIDPixelsReader is a Reader for the GetVisitorsVisitorIDPixels structure.
type GetVisitorsVisitorIDPixelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVisitorsVisitorIDPixelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVisitorsVisitorIDPixelsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetVisitorsVisitorIDPixelsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetVisitorsVisitorIDPixelsOK creates a GetVisitorsVisitorIDPixelsOK with default headers values
func NewGetVisitorsVisitorIDPixelsOK() *GetVisitorsVisitorIDPixelsOK {
	return &GetVisitorsVisitorIDPixelsOK{}
}

/*
GetVisitorsVisitorIDPixelsOK describes a response with status code 200, with default header values.

Pixels
*/
type GetVisitorsVisitorIDPixelsOK struct {
	Payload *GetVisitorsVisitorIDPixelsOKBody
}

// IsSuccess returns true when this get visitors visitor Id pixels o k response has a 2xx status code
func (o *GetVisitorsVisitorIDPixelsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get visitors visitor Id pixels o k response has a 3xx status code
func (o *GetVisitorsVisitorIDPixelsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get visitors visitor Id pixels o k response has a 4xx status code
func (o *GetVisitorsVisitorIDPixelsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get visitors visitor Id pixels o k response has a 5xx status code
func (o *GetVisitorsVisitorIDPixelsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get visitors visitor Id pixels o k response a status code equal to that given
func (o *GetVisitorsVisitorIDPixelsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get visitors visitor Id pixels o k response
func (o *GetVisitorsVisitorIDPixelsOK) Code() int {
	return 200
}

func (o *GetVisitorsVisitorIDPixelsOK) Error() string {
	return fmt.Sprintf("[GET /visitors/{visitorId}/pixels][%d] getVisitorsVisitorIdPixelsOK  %+v", 200, o.Payload)
}

func (o *GetVisitorsVisitorIDPixelsOK) String() string {
	return fmt.Sprintf("[GET /visitors/{visitorId}/pixels][%d] getVisitorsVisitorIdPixelsOK  %+v", 200, o.Payload)
}

func (o *GetVisitorsVisitorIDPixelsOK) GetPayload() *GetVisitorsVisitorIDPixelsOKBody {
	return o.Payload
}

func (o *GetVisitorsVisitorIDPixelsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetVisitorsVisitorIDPixelsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVisitorsVisitorIDPixelsDefault creates a GetVisitorsVisitorIDPixelsDefault with default headers values
func NewGetVisitorsVisitorIDPixelsDefault(code int) *GetVisitorsVisitorIDPixelsDefault {
	return &GetVisitorsVisitorIDPixelsDefault{
		_statusCode: code,
	}
}

/*
GetVisitorsVisitorIDPixelsDefault describes a response with status code -1, with default header values.

Error
*/
type GetVisitorsVisitorIDPixelsDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// IsSuccess returns true when this get visitors visitor ID pixels default response has a 2xx status code
func (o *GetVisitorsVisitorIDPixelsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get visitors visitor ID pixels default response has a 3xx status code
func (o *GetVisitorsVisitorIDPixelsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get visitors visitor ID pixels default response has a 4xx status code
func (o *GetVisitorsVisitorIDPixelsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get visitors visitor ID pixels default response has a 5xx status code
func (o *GetVisitorsVisitorIDPixelsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get visitors visitor ID pixels default response a status code equal to that given
func (o *GetVisitorsVisitorIDPixelsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get visitors visitor ID pixels default response
func (o *GetVisitorsVisitorIDPixelsDefault) Code() int {
	return o._statusCode
}

func (o *GetVisitorsVisitorIDPixelsDefault) Error() string {
	return fmt.Sprintf("[GET /visitors/{visitorId}/pixels][%d] GetVisitorsVisitorIDPixels default  %+v", o._statusCode, o.Payload)
}

func (o *GetVisitorsVisitorIDPixelsDefault) String() string {
	return fmt.Sprintf("[GET /visitors/{visitorId}/pixels][%d] GetVisitorsVisitorIDPixels default  %+v", o._statusCode, o.Payload)
}

func (o *GetVisitorsVisitorIDPixelsDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *GetVisitorsVisitorIDPixelsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetVisitorsVisitorIDPixelsOKBody get visitors visitor ID pixels o k body
swagger:model GetVisitorsVisitorIDPixelsOKBody
*/
type GetVisitorsVisitorIDPixelsOKBody struct {
	models.Envelope

	// data
	Data []*models.AdvertiserPixel `json:"data"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetVisitorsVisitorIDPixelsOKBody) UnmarshalJSON(raw []byte) error {
	// GetVisitorsVisitorIDPixelsOKBodyAO0
	var getVisitorsVisitorIDPixelsOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &getVisitorsVisitorIDPixelsOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = getVisitorsVisitorIDPixelsOKBodyAO0

	// GetVisitorsVisitorIDPixelsOKBodyAO1
	var dataGetVisitorsVisitorIDPixelsOKBodyAO1 struct {
		Data []*models.AdvertiserPixel `json:"data"`
	}
	if err := swag.ReadJSON(raw, &dataGetVisitorsVisitorIDPixelsOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataGetVisitorsVisitorIDPixelsOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetVisitorsVisitorIDPixelsOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getVisitorsVisitorIDPixelsOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getVisitorsVisitorIDPixelsOKBodyAO0)
	var dataGetVisitorsVisitorIDPixelsOKBodyAO1 struct {
		Data []*models.AdvertiserPixel `json:"data"`
	}

	dataGetVisitorsVisitorIDPixelsOKBodyAO1.Data = o.Data

	jsonDataGetVisitorsVisitorIDPixelsOKBodyAO1, errGetVisitorsVisitorIDPixelsOKBodyAO1 := swag.WriteJSON(dataGetVisitorsVisitorIDPixelsOKBodyAO1)
	if errGetVisitorsVisitorIDPixelsOKBodyAO1 != nil {
		return nil, errGetVisitorsVisitorIDPixelsOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetVisitorsVisitorIDPixelsOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get visitors visitor ID pixels o k body
func (o *GetVisitorsVisitorIDPixelsOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetVisitorsVisitorIDPixelsOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	for i := 0; i < len(o.Data); i++ {
		if swag.IsZero(o.Data[i]) { // not required
			continue
		}

		if o.Data[i] != nil {
			if err := o.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getVisitorsVisitorIdPixelsOK" + "." + "data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getVisitorsVisitorIdPixelsOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get visitors visitor ID pixels o k body based on the context it is used
func (o *GetVisitorsVisitorIDPixelsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *GetVisitorsVisitorIDPixelsOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Data); i++ {

		if o.Data[i] != nil {

			if swag.IsZero(o.Data[i]) { // not required
				return nil
			}

			if err := o.Data[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getVisitorsVisitorIdPixelsOK" + "." + "data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getVisitorsVisitorIdPixelsOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetVisitorsVisitorIDPixelsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetVisitorsVisitorIDPixelsOKBody) UnmarshalBinary(b []byte) error {
	var res GetVisitorsVisitorIDPixelsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
