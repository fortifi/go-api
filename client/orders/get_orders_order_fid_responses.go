// Code generated by go-swagger; DO NOT EDIT.

package orders

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/fortifi/go-api/models"
)

// GetOrdersOrderFidReader is a Reader for the GetOrdersOrderFid structure.
type GetOrdersOrderFidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrdersOrderFidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetOrdersOrderFidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetOrdersOrderFidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetOrdersOrderFidOK creates a GetOrdersOrderFidOK with default headers values
func NewGetOrdersOrderFidOK() *GetOrdersOrderFidOK {
	return &GetOrdersOrderFidOK{}
}

/*GetOrdersOrderFidOK handles this case with default header values.

Order retrieved
*/
type GetOrdersOrderFidOK struct {
	Payload *models.GetOrdersOrderFidOKBody
}

func (o *GetOrdersOrderFidOK) Error() string {
	return fmt.Sprintf("[GET /orders/{orderFid}][%d] getOrdersOrderFidOK  %+v", 200, o.Payload)
}

func (o *GetOrdersOrderFidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetOrdersOrderFidOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrdersOrderFidDefault creates a GetOrdersOrderFidDefault with default headers values
func NewGetOrdersOrderFidDefault(code int) *GetOrdersOrderFidDefault {
	return &GetOrdersOrderFidDefault{
		_statusCode: code,
	}
}

/*GetOrdersOrderFidDefault handles this case with default header values.

Error
*/
type GetOrdersOrderFidDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the get orders order fid default response
func (o *GetOrdersOrderFidDefault) Code() int {
	return o._statusCode
}

func (o *GetOrdersOrderFidDefault) Error() string {
	return fmt.Sprintf("[GET /orders/{orderFid}][%d] GetOrdersOrderFid default  %+v", o._statusCode, o.Payload)
}

func (o *GetOrdersOrderFidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
