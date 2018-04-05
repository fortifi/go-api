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

// GetProductsOffersReader is a Reader for the GetProductsOffers structure.
type GetProductsOffersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProductsOffersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetProductsOffersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetProductsOffersOK creates a GetProductsOffersOK with default headers values
func NewGetProductsOffersOK() *GetProductsOffersOK {
	return &GetProductsOffersOK{}
}

/*GetProductsOffersOK handles this case with default header values.

Available Offers
*/
type GetProductsOffersOK struct {
	Payload *models.ProductOffers
}

func (o *GetProductsOffersOK) Error() string {
	return fmt.Sprintf("[GET /products/offers][%d] getProductsOffersOK  %+v", 200, o.Payload)
}

func (o *GetProductsOffersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProductOffers)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
