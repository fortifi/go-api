// Code generated by go-swagger; DO NOT EDIT.

package orders

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fortifi/go-api/models"
)

// PutOrdersOrderFidProductsOrderProductFidQuantityReader is a Reader for the PutOrdersOrderFidProductsOrderProductFidQuantity structure.
type PutOrdersOrderFidProductsOrderProductFidQuantityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutOrdersOrderFidProductsOrderProductFidQuantityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutOrdersOrderFidProductsOrderProductFidQuantityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutOrdersOrderFidProductsOrderProductFidQuantityDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutOrdersOrderFidProductsOrderProductFidQuantityOK creates a PutOrdersOrderFidProductsOrderProductFidQuantityOK with default headers values
func NewPutOrdersOrderFidProductsOrderProductFidQuantityOK() *PutOrdersOrderFidProductsOrderProductFidQuantityOK {
	return &PutOrdersOrderFidProductsOrderProductFidQuantityOK{}
}

/* PutOrdersOrderFidProductsOrderProductFidQuantityOK describes a response with status code 200, with default header values.

Quantity updated
*/
type PutOrdersOrderFidProductsOrderProductFidQuantityOK struct {
}

func (o *PutOrdersOrderFidProductsOrderProductFidQuantityOK) Error() string {
	return fmt.Sprintf("[PUT /orders/{orderFid}/products/{orderProductFid}/quantity][%d] putOrdersOrderFidProductsOrderProductFidQuantityOK ", 200)
}

func (o *PutOrdersOrderFidProductsOrderProductFidQuantityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutOrdersOrderFidProductsOrderProductFidQuantityDefault creates a PutOrdersOrderFidProductsOrderProductFidQuantityDefault with default headers values
func NewPutOrdersOrderFidProductsOrderProductFidQuantityDefault(code int) *PutOrdersOrderFidProductsOrderProductFidQuantityDefault {
	return &PutOrdersOrderFidProductsOrderProductFidQuantityDefault{
		_statusCode: code,
	}
}

/* PutOrdersOrderFidProductsOrderProductFidQuantityDefault describes a response with status code -1, with default header values.

Error
*/
type PutOrdersOrderFidProductsOrderProductFidQuantityDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the put orders order fid products order product fid quantity default response
func (o *PutOrdersOrderFidProductsOrderProductFidQuantityDefault) Code() int {
	return o._statusCode
}

func (o *PutOrdersOrderFidProductsOrderProductFidQuantityDefault) Error() string {
	return fmt.Sprintf("[PUT /orders/{orderFid}/products/{orderProductFid}/quantity][%d] PutOrdersOrderFidProductsOrderProductFidQuantity default  %+v", o._statusCode, o.Payload)
}
func (o *PutOrdersOrderFidProductsOrderProductFidQuantityDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PutOrdersOrderFidProductsOrderProductFidQuantityDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
