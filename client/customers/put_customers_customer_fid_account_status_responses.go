// Code generated by go-swagger; DO NOT EDIT.

package customers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PutCustomersCustomerFidAccountStatusReader is a Reader for the PutCustomersCustomerFidAccountStatus structure.
type PutCustomersCustomerFidAccountStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutCustomersCustomerFidAccountStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutCustomersCustomerFidAccountStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewPutCustomersCustomerFidAccountStatusNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutCustomersCustomerFidAccountStatusOK creates a PutCustomersCustomerFidAccountStatusOK with default headers values
func NewPutCustomersCustomerFidAccountStatusOK() *PutCustomersCustomerFidAccountStatusOK {
	return &PutCustomersCustomerFidAccountStatusOK{}
}

/*PutCustomersCustomerFidAccountStatusOK handles this case with default header values.

Customer Status Updated
*/
type PutCustomersCustomerFidAccountStatusOK struct {
}

func (o *PutCustomersCustomerFidAccountStatusOK) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/accountStatus][%d] putCustomersCustomerFidAccountStatusOK ", 200)
}

func (o *PutCustomersCustomerFidAccountStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCustomersCustomerFidAccountStatusNotFound creates a PutCustomersCustomerFidAccountStatusNotFound with default headers values
func NewPutCustomersCustomerFidAccountStatusNotFound() *PutCustomersCustomerFidAccountStatusNotFound {
	return &PutCustomersCustomerFidAccountStatusNotFound{}
}

/*PutCustomersCustomerFidAccountStatusNotFound handles this case with default header values.

Customer not found
*/
type PutCustomersCustomerFidAccountStatusNotFound struct {
}

func (o *PutCustomersCustomerFidAccountStatusNotFound) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/accountStatus][%d] putCustomersCustomerFidAccountStatusNotFound ", 404)
}

func (o *PutCustomersCustomerFidAccountStatusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
