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

// PostOrdersReader is a Reader for the PostOrders structure.
type PostOrdersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostOrdersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostOrdersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostOrdersOK creates a PostOrdersOK with default headers values
func NewPostOrdersOK() *PostOrdersOK {
	return &PostOrdersOK{}
}

/*PostOrdersOK handles this case with default header values.

Order created successfully
*/
type PostOrdersOK struct {
	Payload *models.Fid
}

func (o *PostOrdersOK) Error() string {
	return fmt.Sprintf("[POST /orders][%d] postOrdersOK  %+v", 200, o.Payload)
}

func (o *PostOrdersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Fid)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}