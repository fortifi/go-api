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

// GetCustomersCustomerFidContactsReader is a Reader for the GetCustomersCustomerFidContacts structure.
type GetCustomersCustomerFidContactsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCustomersCustomerFidContactsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCustomersCustomerFidContactsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetCustomersCustomerFidContactsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCustomersCustomerFidContactsOK creates a GetCustomersCustomerFidContactsOK with default headers values
func NewGetCustomersCustomerFidContactsOK() *GetCustomersCustomerFidContactsOK {
	return &GetCustomersCustomerFidContactsOK{}
}

/*GetCustomersCustomerFidContactsOK handles this case with default header values.

List of people
*/
type GetCustomersCustomerFidContactsOK struct {
	Payload *GetCustomersCustomerFidContactsOKBody
}

func (o *GetCustomersCustomerFidContactsOK) Error() string {
	return fmt.Sprintf("[GET /customers/{customerFid}/contacts][%d] getCustomersCustomerFidContactsOK  %+v", 200, o.Payload)
}

func (o *GetCustomersCustomerFidContactsOK) GetPayload() *GetCustomersCustomerFidContactsOKBody {
	return o.Payload
}

func (o *GetCustomersCustomerFidContactsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetCustomersCustomerFidContactsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCustomersCustomerFidContactsDefault creates a GetCustomersCustomerFidContactsDefault with default headers values
func NewGetCustomersCustomerFidContactsDefault(code int) *GetCustomersCustomerFidContactsDefault {
	return &GetCustomersCustomerFidContactsDefault{
		_statusCode: code,
	}
}

/*GetCustomersCustomerFidContactsDefault handles this case with default header values.

Error
*/
type GetCustomersCustomerFidContactsDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the get customers customer fid contacts default response
func (o *GetCustomersCustomerFidContactsDefault) Code() int {
	return o._statusCode
}

func (o *GetCustomersCustomerFidContactsDefault) Error() string {
	return fmt.Sprintf("[GET /customers/{customerFid}/contacts][%d] GetCustomersCustomerFidContacts default  %+v", o._statusCode, o.Payload)
}

func (o *GetCustomersCustomerFidContactsDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *GetCustomersCustomerFidContactsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetCustomersCustomerFidContactsOKBody get customers customer fid contacts o k body
swagger:model GetCustomersCustomerFidContactsOKBody
*/
type GetCustomersCustomerFidContactsOKBody struct {
	models.Envelope

	// data
	Data *models.People `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetCustomersCustomerFidContactsOKBody) UnmarshalJSON(raw []byte) error {
	// GetCustomersCustomerFidContactsOKBodyAO0
	var getCustomersCustomerFidContactsOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &getCustomersCustomerFidContactsOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = getCustomersCustomerFidContactsOKBodyAO0

	// GetCustomersCustomerFidContactsOKBodyAO1
	var dataGetCustomersCustomerFidContactsOKBodyAO1 struct {
		Data *models.People `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataGetCustomersCustomerFidContactsOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataGetCustomersCustomerFidContactsOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetCustomersCustomerFidContactsOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getCustomersCustomerFidContactsOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getCustomersCustomerFidContactsOKBodyAO0)

	var dataGetCustomersCustomerFidContactsOKBodyAO1 struct {
		Data *models.People `json:"data,omitempty"`
	}

	dataGetCustomersCustomerFidContactsOKBodyAO1.Data = o.Data

	jsonDataGetCustomersCustomerFidContactsOKBodyAO1, errGetCustomersCustomerFidContactsOKBodyAO1 := swag.WriteJSON(dataGetCustomersCustomerFidContactsOKBodyAO1)
	if errGetCustomersCustomerFidContactsOKBodyAO1 != nil {
		return nil, errGetCustomersCustomerFidContactsOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetCustomersCustomerFidContactsOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get customers customer fid contacts o k body
func (o *GetCustomersCustomerFidContactsOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetCustomersCustomerFidContactsOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getCustomersCustomerFidContactsOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetCustomersCustomerFidContactsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCustomersCustomerFidContactsOKBody) UnmarshalBinary(b []byte) error {
	var res GetCustomersCustomerFidContactsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}