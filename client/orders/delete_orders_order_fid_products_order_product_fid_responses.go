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

// DeleteOrdersOrderFidProductsOrderProductFidReader is a Reader for the DeleteOrdersOrderFidProductsOrderProductFid structure.
type DeleteOrdersOrderFidProductsOrderProductFidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOrdersOrderFidProductsOrderProductFidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteOrdersOrderFidProductsOrderProductFidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteOrdersOrderFidProductsOrderProductFidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteOrdersOrderFidProductsOrderProductFidOK creates a DeleteOrdersOrderFidProductsOrderProductFidOK with default headers values
func NewDeleteOrdersOrderFidProductsOrderProductFidOK() *DeleteOrdersOrderFidProductsOrderProductFidOK {
	return &DeleteOrdersOrderFidProductsOrderProductFidOK{}
}

/*
DeleteOrdersOrderFidProductsOrderProductFidOK describes a response with status code 200, with default header values.

Product removed from the order successfully
*/
type DeleteOrdersOrderFidProductsOrderProductFidOK struct {
}

// IsSuccess returns true when this delete orders order fid products order product fid o k response has a 2xx status code
func (o *DeleteOrdersOrderFidProductsOrderProductFidOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete orders order fid products order product fid o k response has a 3xx status code
func (o *DeleteOrdersOrderFidProductsOrderProductFidOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete orders order fid products order product fid o k response has a 4xx status code
func (o *DeleteOrdersOrderFidProductsOrderProductFidOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete orders order fid products order product fid o k response has a 5xx status code
func (o *DeleteOrdersOrderFidProductsOrderProductFidOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete orders order fid products order product fid o k response a status code equal to that given
func (o *DeleteOrdersOrderFidProductsOrderProductFidOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete orders order fid products order product fid o k response
func (o *DeleteOrdersOrderFidProductsOrderProductFidOK) Code() int {
	return 200
}

func (o *DeleteOrdersOrderFidProductsOrderProductFidOK) Error() string {
	return fmt.Sprintf("[DELETE /orders/{orderFid}/products/{orderProductFid}][%d] deleteOrdersOrderFidProductsOrderProductFidOK ", 200)
}

func (o *DeleteOrdersOrderFidProductsOrderProductFidOK) String() string {
	return fmt.Sprintf("[DELETE /orders/{orderFid}/products/{orderProductFid}][%d] deleteOrdersOrderFidProductsOrderProductFidOK ", 200)
}

func (o *DeleteOrdersOrderFidProductsOrderProductFidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteOrdersOrderFidProductsOrderProductFidDefault creates a DeleteOrdersOrderFidProductsOrderProductFidDefault with default headers values
func NewDeleteOrdersOrderFidProductsOrderProductFidDefault(code int) *DeleteOrdersOrderFidProductsOrderProductFidDefault {
	return &DeleteOrdersOrderFidProductsOrderProductFidDefault{
		_statusCode: code,
	}
}

/*
DeleteOrdersOrderFidProductsOrderProductFidDefault describes a response with status code -1, with default header values.

Error
*/
type DeleteOrdersOrderFidProductsOrderProductFidDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// IsSuccess returns true when this delete orders order fid products order product fid default response has a 2xx status code
func (o *DeleteOrdersOrderFidProductsOrderProductFidDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete orders order fid products order product fid default response has a 3xx status code
func (o *DeleteOrdersOrderFidProductsOrderProductFidDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete orders order fid products order product fid default response has a 4xx status code
func (o *DeleteOrdersOrderFidProductsOrderProductFidDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete orders order fid products order product fid default response has a 5xx status code
func (o *DeleteOrdersOrderFidProductsOrderProductFidDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete orders order fid products order product fid default response a status code equal to that given
func (o *DeleteOrdersOrderFidProductsOrderProductFidDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete orders order fid products order product fid default response
func (o *DeleteOrdersOrderFidProductsOrderProductFidDefault) Code() int {
	return o._statusCode
}

func (o *DeleteOrdersOrderFidProductsOrderProductFidDefault) Error() string {
	return fmt.Sprintf("[DELETE /orders/{orderFid}/products/{orderProductFid}][%d] DeleteOrdersOrderFidProductsOrderProductFid default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteOrdersOrderFidProductsOrderProductFidDefault) String() string {
	return fmt.Sprintf("[DELETE /orders/{orderFid}/products/{orderProductFid}][%d] DeleteOrdersOrderFidProductsOrderProductFid default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteOrdersOrderFidProductsOrderProductFidDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *DeleteOrdersOrderFidProductsOrderProductFidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
