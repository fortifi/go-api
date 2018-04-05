// Code generated by go-swagger; DO NOT EDIT.

package customers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PutCustomersCustomerFidCurrencyReader is a Reader for the PutCustomersCustomerFidCurrency structure.
type PutCustomersCustomerFidCurrencyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutCustomersCustomerFidCurrencyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutCustomersCustomerFidCurrencyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewPutCustomersCustomerFidCurrencyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutCustomersCustomerFidCurrencyOK creates a PutCustomersCustomerFidCurrencyOK with default headers values
func NewPutCustomersCustomerFidCurrencyOK() *PutCustomersCustomerFidCurrencyOK {
	return &PutCustomersCustomerFidCurrencyOK{}
}

/*PutCustomersCustomerFidCurrencyOK handles this case with default header values.

Customer Currency Updated
*/
type PutCustomersCustomerFidCurrencyOK struct {
}

func (o *PutCustomersCustomerFidCurrencyOK) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/currency][%d] putCustomersCustomerFidCurrencyOK ", 200)
}

func (o *PutCustomersCustomerFidCurrencyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCustomersCustomerFidCurrencyNotFound creates a PutCustomersCustomerFidCurrencyNotFound with default headers values
func NewPutCustomersCustomerFidCurrencyNotFound() *PutCustomersCustomerFidCurrencyNotFound {
	return &PutCustomersCustomerFidCurrencyNotFound{}
}

/*PutCustomersCustomerFidCurrencyNotFound handles this case with default header values.

Customer not found
*/
type PutCustomersCustomerFidCurrencyNotFound struct {
}

func (o *PutCustomersCustomerFidCurrencyNotFound) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/currency][%d] putCustomersCustomerFidCurrencyNotFound ", 404)
}

func (o *PutCustomersCustomerFidCurrencyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}