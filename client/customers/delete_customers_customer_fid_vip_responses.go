// Code generated by go-swagger; DO NOT EDIT.

package customers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fortifi/go-api/models"
)

// DeleteCustomersCustomerFidVipReader is a Reader for the DeleteCustomersCustomerFidVip structure.
type DeleteCustomersCustomerFidVipReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCustomersCustomerFidVipReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteCustomersCustomerFidVipOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteCustomersCustomerFidVipDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteCustomersCustomerFidVipOK creates a DeleteCustomersCustomerFidVipOK with default headers values
func NewDeleteCustomersCustomerFidVipOK() *DeleteCustomersCustomerFidVipOK {
	return &DeleteCustomersCustomerFidVipOK{}
}

/* DeleteCustomersCustomerFidVipOK describes a response with status code 200, with default header values.

Customer No Longer VIP
*/
type DeleteCustomersCustomerFidVipOK struct {
}

func (o *DeleteCustomersCustomerFidVipOK) Error() string {
	return fmt.Sprintf("[DELETE /customers/{customerFid}/vip][%d] deleteCustomersCustomerFidVipOK ", 200)
}

func (o *DeleteCustomersCustomerFidVipOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCustomersCustomerFidVipDefault creates a DeleteCustomersCustomerFidVipDefault with default headers values
func NewDeleteCustomersCustomerFidVipDefault(code int) *DeleteCustomersCustomerFidVipDefault {
	return &DeleteCustomersCustomerFidVipDefault{
		_statusCode: code,
	}
}

/* DeleteCustomersCustomerFidVipDefault describes a response with status code -1, with default header values.

Error
*/
type DeleteCustomersCustomerFidVipDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the delete customers customer fid vip default response
func (o *DeleteCustomersCustomerFidVipDefault) Code() int {
	return o._statusCode
}

func (o *DeleteCustomersCustomerFidVipDefault) Error() string {
	return fmt.Sprintf("[DELETE /customers/{customerFid}/vip][%d] DeleteCustomersCustomerFidVip default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteCustomersCustomerFidVipDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *DeleteCustomersCustomerFidVipDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
