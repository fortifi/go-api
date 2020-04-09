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

// PostCustomersCustomerFidTicketsTicketFidPostsReader is a Reader for the PostCustomersCustomerFidTicketsTicketFidPosts structure.
type PostCustomersCustomerFidTicketsTicketFidPostsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCustomersCustomerFidTicketsTicketFidPostsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostCustomersCustomerFidTicketsTicketFidPostsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostCustomersCustomerFidTicketsTicketFidPostsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostCustomersCustomerFidTicketsTicketFidPostsOK creates a PostCustomersCustomerFidTicketsTicketFidPostsOK with default headers values
func NewPostCustomersCustomerFidTicketsTicketFidPostsOK() *PostCustomersCustomerFidTicketsTicketFidPostsOK {
	return &PostCustomersCustomerFidTicketsTicketFidPostsOK{}
}

/*PostCustomersCustomerFidTicketsTicketFidPostsOK handles this case with default header values.

Ticket Data
*/
type PostCustomersCustomerFidTicketsTicketFidPostsOK struct {
	Payload *PostCustomersCustomerFidTicketsTicketFidPostsOKBody
}

func (o *PostCustomersCustomerFidTicketsTicketFidPostsOK) Error() string {
	return fmt.Sprintf("[POST /customers/{customerFid}/tickets/{ticketFid}/posts][%d] postCustomersCustomerFidTicketsTicketFidPostsOK  %+v", 200, o.Payload)
}

func (o *PostCustomersCustomerFidTicketsTicketFidPostsOK) GetPayload() *PostCustomersCustomerFidTicketsTicketFidPostsOKBody {
	return o.Payload
}

func (o *PostCustomersCustomerFidTicketsTicketFidPostsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostCustomersCustomerFidTicketsTicketFidPostsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCustomersCustomerFidTicketsTicketFidPostsDefault creates a PostCustomersCustomerFidTicketsTicketFidPostsDefault with default headers values
func NewPostCustomersCustomerFidTicketsTicketFidPostsDefault(code int) *PostCustomersCustomerFidTicketsTicketFidPostsDefault {
	return &PostCustomersCustomerFidTicketsTicketFidPostsDefault{
		_statusCode: code,
	}
}

/*PostCustomersCustomerFidTicketsTicketFidPostsDefault handles this case with default header values.

Error
*/
type PostCustomersCustomerFidTicketsTicketFidPostsDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the post customers customer fid tickets ticket fid posts default response
func (o *PostCustomersCustomerFidTicketsTicketFidPostsDefault) Code() int {
	return o._statusCode
}

func (o *PostCustomersCustomerFidTicketsTicketFidPostsDefault) Error() string {
	return fmt.Sprintf("[POST /customers/{customerFid}/tickets/{ticketFid}/posts][%d] PostCustomersCustomerFidTicketsTicketFidPosts default  %+v", o._statusCode, o.Payload)
}

func (o *PostCustomersCustomerFidTicketsTicketFidPostsDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PostCustomersCustomerFidTicketsTicketFidPostsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostCustomersCustomerFidTicketsTicketFidPostsOKBody post customers customer fid tickets ticket fid posts o k body
swagger:model PostCustomersCustomerFidTicketsTicketFidPostsOKBody
*/
type PostCustomersCustomerFidTicketsTicketFidPostsOKBody struct {
	models.Envelope

	// data
	Data *models.TicketPost `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostCustomersCustomerFidTicketsTicketFidPostsOKBody) UnmarshalJSON(raw []byte) error {
	// PostCustomersCustomerFidTicketsTicketFidPostsOKBodyAO0
	var postCustomersCustomerFidTicketsTicketFidPostsOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &postCustomersCustomerFidTicketsTicketFidPostsOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = postCustomersCustomerFidTicketsTicketFidPostsOKBodyAO0

	// PostCustomersCustomerFidTicketsTicketFidPostsOKBodyAO1
	var dataPostCustomersCustomerFidTicketsTicketFidPostsOKBodyAO1 struct {
		Data *models.TicketPost `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataPostCustomersCustomerFidTicketsTicketFidPostsOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataPostCustomersCustomerFidTicketsTicketFidPostsOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostCustomersCustomerFidTicketsTicketFidPostsOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	postCustomersCustomerFidTicketsTicketFidPostsOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postCustomersCustomerFidTicketsTicketFidPostsOKBodyAO0)

	var dataPostCustomersCustomerFidTicketsTicketFidPostsOKBodyAO1 struct {
		Data *models.TicketPost `json:"data,omitempty"`
	}

	dataPostCustomersCustomerFidTicketsTicketFidPostsOKBodyAO1.Data = o.Data

	jsonDataPostCustomersCustomerFidTicketsTicketFidPostsOKBodyAO1, errPostCustomersCustomerFidTicketsTicketFidPostsOKBodyAO1 := swag.WriteJSON(dataPostCustomersCustomerFidTicketsTicketFidPostsOKBodyAO1)
	if errPostCustomersCustomerFidTicketsTicketFidPostsOKBodyAO1 != nil {
		return nil, errPostCustomersCustomerFidTicketsTicketFidPostsOKBodyAO1
	}
	_parts = append(_parts, jsonDataPostCustomersCustomerFidTicketsTicketFidPostsOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post customers customer fid tickets ticket fid posts o k body
func (o *PostCustomersCustomerFidTicketsTicketFidPostsOKBody) Validate(formats strfmt.Registry) error {
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

func (o *PostCustomersCustomerFidTicketsTicketFidPostsOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postCustomersCustomerFidTicketsTicketFidPostsOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostCustomersCustomerFidTicketsTicketFidPostsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostCustomersCustomerFidTicketsTicketFidPostsOKBody) UnmarshalBinary(b []byte) error {
	var res PostCustomersCustomerFidTicketsTicketFidPostsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}