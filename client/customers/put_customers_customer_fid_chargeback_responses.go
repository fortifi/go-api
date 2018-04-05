// Code generated by go-swagger; DO NOT EDIT.

package customers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PutCustomersCustomerFidChargebackReader is a Reader for the PutCustomersCustomerFidChargeback structure.
type PutCustomersCustomerFidChargebackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutCustomersCustomerFidChargebackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutCustomersCustomerFidChargebackOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewPutCustomersCustomerFidChargebackNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutCustomersCustomerFidChargebackOK creates a PutCustomersCustomerFidChargebackOK with default headers values
func NewPutCustomersCustomerFidChargebackOK() *PutCustomersCustomerFidChargebackOK {
	return &PutCustomersCustomerFidChargebackOK{}
}

/*PutCustomersCustomerFidChargebackOK handles this case with default header values.

Customer Marked
*/
type PutCustomersCustomerFidChargebackOK struct {
}

func (o *PutCustomersCustomerFidChargebackOK) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/chargeback][%d] putCustomersCustomerFidChargebackOK ", 200)
}

func (o *PutCustomersCustomerFidChargebackOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCustomersCustomerFidChargebackNotFound creates a PutCustomersCustomerFidChargebackNotFound with default headers values
func NewPutCustomersCustomerFidChargebackNotFound() *PutCustomersCustomerFidChargebackNotFound {
	return &PutCustomersCustomerFidChargebackNotFound{}
}

/*PutCustomersCustomerFidChargebackNotFound handles this case with default header values.

Customer not found
*/
type PutCustomersCustomerFidChargebackNotFound struct {
}

func (o *PutCustomersCustomerFidChargebackNotFound) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/chargeback][%d] putCustomersCustomerFidChargebackNotFound ", 404)
}

func (o *PutCustomersCustomerFidChargebackNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
