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

// PostCustomersCustomerFidSarReader is a Reader for the PostCustomersCustomerFidSar structure.
type PostCustomersCustomerFidSarReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCustomersCustomerFidSarReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostCustomersCustomerFidSarOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostCustomersCustomerFidSarDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostCustomersCustomerFidSarOK creates a PostCustomersCustomerFidSarOK with default headers values
func NewPostCustomersCustomerFidSarOK() *PostCustomersCustomerFidSarOK {
	return &PostCustomersCustomerFidSarOK{}
}

/*PostCustomersCustomerFidSarOK handles this case with default header values.

Subject Access Request
*/
type PostCustomersCustomerFidSarOK struct {
	Payload *PostCustomersCustomerFidSarOKBody
}

func (o *PostCustomersCustomerFidSarOK) Error() string {
	return fmt.Sprintf("[POST /customers/{customerFid}/sar][%d] postCustomersCustomerFidSarOK  %+v", 200, o.Payload)
}

func (o *PostCustomersCustomerFidSarOK) GetPayload() *PostCustomersCustomerFidSarOKBody {
	return o.Payload
}

func (o *PostCustomersCustomerFidSarOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostCustomersCustomerFidSarOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCustomersCustomerFidSarDefault creates a PostCustomersCustomerFidSarDefault with default headers values
func NewPostCustomersCustomerFidSarDefault(code int) *PostCustomersCustomerFidSarDefault {
	return &PostCustomersCustomerFidSarDefault{
		_statusCode: code,
	}
}

/*PostCustomersCustomerFidSarDefault handles this case with default header values.

Error
*/
type PostCustomersCustomerFidSarDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the post customers customer fid sar default response
func (o *PostCustomersCustomerFidSarDefault) Code() int {
	return o._statusCode
}

func (o *PostCustomersCustomerFidSarDefault) Error() string {
	return fmt.Sprintf("[POST /customers/{customerFid}/sar][%d] PostCustomersCustomerFidSar default  %+v", o._statusCode, o.Payload)
}

func (o *PostCustomersCustomerFidSarDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PostCustomersCustomerFidSarDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostCustomersCustomerFidSarOKBody post customers customer fid sar o k body
swagger:model PostCustomersCustomerFidSarOKBody
*/
type PostCustomersCustomerFidSarOKBody struct {
	models.Envelope

	// data
	Data *models.BoolMessage `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostCustomersCustomerFidSarOKBody) UnmarshalJSON(raw []byte) error {
	// PostCustomersCustomerFidSarOKBodyAO0
	var postCustomersCustomerFidSarOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &postCustomersCustomerFidSarOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = postCustomersCustomerFidSarOKBodyAO0

	// PostCustomersCustomerFidSarOKBodyAO1
	var dataPostCustomersCustomerFidSarOKBodyAO1 struct {
		Data *models.BoolMessage `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataPostCustomersCustomerFidSarOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataPostCustomersCustomerFidSarOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostCustomersCustomerFidSarOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	postCustomersCustomerFidSarOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postCustomersCustomerFidSarOKBodyAO0)
	var dataPostCustomersCustomerFidSarOKBodyAO1 struct {
		Data *models.BoolMessage `json:"data,omitempty"`
	}

	dataPostCustomersCustomerFidSarOKBodyAO1.Data = o.Data

	jsonDataPostCustomersCustomerFidSarOKBodyAO1, errPostCustomersCustomerFidSarOKBodyAO1 := swag.WriteJSON(dataPostCustomersCustomerFidSarOKBodyAO1)
	if errPostCustomersCustomerFidSarOKBodyAO1 != nil {
		return nil, errPostCustomersCustomerFidSarOKBodyAO1
	}
	_parts = append(_parts, jsonDataPostCustomersCustomerFidSarOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post customers customer fid sar o k body
func (o *PostCustomersCustomerFidSarOKBody) Validate(formats strfmt.Registry) error {
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

func (o *PostCustomersCustomerFidSarOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postCustomersCustomerFidSarOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostCustomersCustomerFidSarOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostCustomersCustomerFidSarOKBody) UnmarshalBinary(b []byte) error {
	var res PostCustomersCustomerFidSarOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
