// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/fortifi/go-api/models"
)

// GetProductsReader is a Reader for the GetProducts structure.
type GetProductsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProductsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetProductsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetProductsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetProductsOK creates a GetProductsOK with default headers values
func NewGetProductsOK() *GetProductsOK {
	return &GetProductsOK{}
}

/*GetProductsOK handles this case with default header values.

Products retrieved
*/
type GetProductsOK struct {
	Payload *models.GetProductsOKBody
}

func (o *GetProductsOK) Error() string {
	return fmt.Sprintf("[GET /products][%d] getProductsOK  %+v", 200, o.Payload)
}

func (o *GetProductsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetProductsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProductsDefault creates a GetProductsDefault with default headers values
func NewGetProductsDefault(code int) *GetProductsDefault {
	return &GetProductsDefault{
		_statusCode: code,
	}
}

/*GetProductsDefault handles this case with default header values.

Error
*/
type GetProductsDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the get products default response
func (o *GetProductsDefault) Code() int {
	return o._statusCode
}

func (o *GetProductsDefault) Error() string {
	return fmt.Sprintf("[GET /products][%d] GetProducts default  %+v", o._statusCode, o.Payload)
}

func (o *GetProductsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
