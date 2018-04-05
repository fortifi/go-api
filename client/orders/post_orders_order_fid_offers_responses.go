// Code generated by go-swagger; DO NOT EDIT.

package orders

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PostOrdersOrderFidOffersReader is a Reader for the PostOrdersOrderFidOffers structure.
type PostOrdersOrderFidOffersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostOrdersOrderFidOffersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostOrdersOrderFidOffersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewPostOrdersOrderFidOffersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostOrdersOrderFidOffersOK creates a PostOrdersOrderFidOffersOK with default headers values
func NewPostOrdersOrderFidOffersOK() *PostOrdersOrderFidOffersOK {
	return &PostOrdersOrderFidOffersOK{}
}

/*PostOrdersOrderFidOffersOK handles this case with default header values.

Offer added to the order successfully
*/
type PostOrdersOrderFidOffersOK struct {
}

func (o *PostOrdersOrderFidOffersOK) Error() string {
	return fmt.Sprintf("[POST /orders/{orderFid}/offers][%d] postOrdersOrderFidOffersOK ", 200)
}

func (o *PostOrdersOrderFidOffersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostOrdersOrderFidOffersNotFound creates a PostOrdersOrderFidOffersNotFound with default headers values
func NewPostOrdersOrderFidOffersNotFound() *PostOrdersOrderFidOffersNotFound {
	return &PostOrdersOrderFidOffersNotFound{}
}

/*PostOrdersOrderFidOffersNotFound handles this case with default header values.

Order not found
*/
type PostOrdersOrderFidOffersNotFound struct {
}

func (o *PostOrdersOrderFidOffersNotFound) Error() string {
	return fmt.Sprintf("[POST /orders/{orderFid}/offers][%d] postOrdersOrderFidOffersNotFound ", 404)
}

func (o *PostOrdersOrderFidOffersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
