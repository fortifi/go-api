// Code generated by go-swagger; DO NOT EDIT.

package orders

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/fortifi/go-api/models"
)

// PostOrdersOrderFidProductsReader is a Reader for the PostOrdersOrderFidProducts structure.
type PostOrdersOrderFidProductsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostOrdersOrderFidProductsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostOrdersOrderFidProductsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostOrdersOrderFidProductsOK creates a PostOrdersOrderFidProductsOK with default headers values
func NewPostOrdersOrderFidProductsOK() *PostOrdersOrderFidProductsOK {
	return &PostOrdersOrderFidProductsOK{}
}

/*PostOrdersOrderFidProductsOK handles this case with default header values.

Product added to the order successfully
*/
type PostOrdersOrderFidProductsOK struct {
	Payload *models.OrderAddProducts
}

func (o *PostOrdersOrderFidProductsOK) Error() string {
	return fmt.Sprintf("[POST /orders/{orderFid}/products][%d] postOrdersOrderFidProductsOK  %+v", 200, o.Payload)
}

func (o *PostOrdersOrderFidProductsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OrderAddProducts)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
