// Code generated by go-swagger; DO NOT EDIT.

package customers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteCustomersCustomerFidLoyalReader is a Reader for the DeleteCustomersCustomerFidLoyal structure.
type DeleteCustomersCustomerFidLoyalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCustomersCustomerFidLoyalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteCustomersCustomerFidLoyalOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewDeleteCustomersCustomerFidLoyalNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteCustomersCustomerFidLoyalOK creates a DeleteCustomersCustomerFidLoyalOK with default headers values
func NewDeleteCustomersCustomerFidLoyalOK() *DeleteCustomersCustomerFidLoyalOK {
	return &DeleteCustomersCustomerFidLoyalOK{}
}

/*DeleteCustomersCustomerFidLoyalOK handles this case with default header values.

Customer No Longer Loyal
*/
type DeleteCustomersCustomerFidLoyalOK struct {
}

func (o *DeleteCustomersCustomerFidLoyalOK) Error() string {
	return fmt.Sprintf("[DELETE /customers/{customerFid}/loyal][%d] deleteCustomersCustomerFidLoyalOK ", 200)
}

func (o *DeleteCustomersCustomerFidLoyalOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCustomersCustomerFidLoyalNotFound creates a DeleteCustomersCustomerFidLoyalNotFound with default headers values
func NewDeleteCustomersCustomerFidLoyalNotFound() *DeleteCustomersCustomerFidLoyalNotFound {
	return &DeleteCustomersCustomerFidLoyalNotFound{}
}

/*DeleteCustomersCustomerFidLoyalNotFound handles this case with default header values.

Customer not found
*/
type DeleteCustomersCustomerFidLoyalNotFound struct {
}

func (o *DeleteCustomersCustomerFidLoyalNotFound) Error() string {
	return fmt.Sprintf("[DELETE /customers/{customerFid}/loyal][%d] deleteCustomersCustomerFidLoyalNotFound ", 404)
}

func (o *DeleteCustomersCustomerFidLoyalNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
