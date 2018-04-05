// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/fortifi/go-api/models"
)

// GetPayCoinbaseReader is a Reader for the GetPayCoinbase structure.
type GetPayCoinbaseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPayCoinbaseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPayCoinbaseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPayCoinbaseOK creates a GetPayCoinbaseOK with default headers values
func NewGetPayCoinbaseOK() *GetPayCoinbaseOK {
	return &GetPayCoinbaseOK{}
}

/*GetPayCoinbaseOK handles this case with default header values.

Coinbase checkout ID
*/
type GetPayCoinbaseOK struct {
	Payload *models.CoinbaseCheckout
}

func (o *GetPayCoinbaseOK) Error() string {
	return fmt.Sprintf("[GET /pay/coinbase][%d] getPayCoinbaseOK  %+v", 200, o.Payload)
}

func (o *GetPayCoinbaseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoinbaseCheckout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}