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

// GetCustomersCustomerFidTicketsTicketFidReader is a Reader for the GetCustomersCustomerFidTicketsTicketFid structure.
type GetCustomersCustomerFidTicketsTicketFidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCustomersCustomerFidTicketsTicketFidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCustomersCustomerFidTicketsTicketFidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetCustomersCustomerFidTicketsTicketFidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCustomersCustomerFidTicketsTicketFidOK creates a GetCustomersCustomerFidTicketsTicketFidOK with default headers values
func NewGetCustomersCustomerFidTicketsTicketFidOK() *GetCustomersCustomerFidTicketsTicketFidOK {
	return &GetCustomersCustomerFidTicketsTicketFidOK{}
}

/*GetCustomersCustomerFidTicketsTicketFidOK handles this case with default header values.

Ticket Data
*/
type GetCustomersCustomerFidTicketsTicketFidOK struct {
	Payload *GetCustomersCustomerFidTicketsTicketFidOKBody
}

func (o *GetCustomersCustomerFidTicketsTicketFidOK) Error() string {
	return fmt.Sprintf("[GET /customers/{customerFid}/tickets/{ticketFid}][%d] getCustomersCustomerFidTicketsTicketFidOK  %+v", 200, o.Payload)
}

func (o *GetCustomersCustomerFidTicketsTicketFidOK) GetPayload() *GetCustomersCustomerFidTicketsTicketFidOKBody {
	return o.Payload
}

func (o *GetCustomersCustomerFidTicketsTicketFidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetCustomersCustomerFidTicketsTicketFidOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCustomersCustomerFidTicketsTicketFidDefault creates a GetCustomersCustomerFidTicketsTicketFidDefault with default headers values
func NewGetCustomersCustomerFidTicketsTicketFidDefault(code int) *GetCustomersCustomerFidTicketsTicketFidDefault {
	return &GetCustomersCustomerFidTicketsTicketFidDefault{
		_statusCode: code,
	}
}

/*GetCustomersCustomerFidTicketsTicketFidDefault handles this case with default header values.

Error
*/
type GetCustomersCustomerFidTicketsTicketFidDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the get customers customer fid tickets ticket fid default response
func (o *GetCustomersCustomerFidTicketsTicketFidDefault) Code() int {
	return o._statusCode
}

func (o *GetCustomersCustomerFidTicketsTicketFidDefault) Error() string {
	return fmt.Sprintf("[GET /customers/{customerFid}/tickets/{ticketFid}][%d] GetCustomersCustomerFidTicketsTicketFid default  %+v", o._statusCode, o.Payload)
}

func (o *GetCustomersCustomerFidTicketsTicketFidDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *GetCustomersCustomerFidTicketsTicketFidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetCustomersCustomerFidTicketsTicketFidOKBody get customers customer fid tickets ticket fid o k body
swagger:model GetCustomersCustomerFidTicketsTicketFidOKBody
*/
type GetCustomersCustomerFidTicketsTicketFidOKBody struct {
	models.Envelope

	// data
	Data *models.Ticket `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetCustomersCustomerFidTicketsTicketFidOKBody) UnmarshalJSON(raw []byte) error {
	// GetCustomersCustomerFidTicketsTicketFidOKBodyAO0
	var getCustomersCustomerFidTicketsTicketFidOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &getCustomersCustomerFidTicketsTicketFidOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = getCustomersCustomerFidTicketsTicketFidOKBodyAO0

	// GetCustomersCustomerFidTicketsTicketFidOKBodyAO1
	var dataGetCustomersCustomerFidTicketsTicketFidOKBodyAO1 struct {
		Data *models.Ticket `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataGetCustomersCustomerFidTicketsTicketFidOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataGetCustomersCustomerFidTicketsTicketFidOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetCustomersCustomerFidTicketsTicketFidOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getCustomersCustomerFidTicketsTicketFidOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getCustomersCustomerFidTicketsTicketFidOKBodyAO0)
	var dataGetCustomersCustomerFidTicketsTicketFidOKBodyAO1 struct {
		Data *models.Ticket `json:"data,omitempty"`
	}

	dataGetCustomersCustomerFidTicketsTicketFidOKBodyAO1.Data = o.Data

	jsonDataGetCustomersCustomerFidTicketsTicketFidOKBodyAO1, errGetCustomersCustomerFidTicketsTicketFidOKBodyAO1 := swag.WriteJSON(dataGetCustomersCustomerFidTicketsTicketFidOKBodyAO1)
	if errGetCustomersCustomerFidTicketsTicketFidOKBodyAO1 != nil {
		return nil, errGetCustomersCustomerFidTicketsTicketFidOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetCustomersCustomerFidTicketsTicketFidOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get customers customer fid tickets ticket fid o k body
func (o *GetCustomersCustomerFidTicketsTicketFidOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetCustomersCustomerFidTicketsTicketFidOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getCustomersCustomerFidTicketsTicketFidOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetCustomersCustomerFidTicketsTicketFidOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCustomersCustomerFidTicketsTicketFidOKBody) UnmarshalBinary(b []byte) error {
	var res GetCustomersCustomerFidTicketsTicketFidOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
