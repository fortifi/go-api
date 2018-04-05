// Code generated by go-swagger; DO NOT EDIT.

package customers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PutCustomersCustomerFidAccountTypeReader is a Reader for the PutCustomersCustomerFidAccountType structure.
type PutCustomersCustomerFidAccountTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutCustomersCustomerFidAccountTypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutCustomersCustomerFidAccountTypeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewPutCustomersCustomerFidAccountTypeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutCustomersCustomerFidAccountTypeOK creates a PutCustomersCustomerFidAccountTypeOK with default headers values
func NewPutCustomersCustomerFidAccountTypeOK() *PutCustomersCustomerFidAccountTypeOK {
	return &PutCustomersCustomerFidAccountTypeOK{}
}

/*PutCustomersCustomerFidAccountTypeOK handles this case with default header values.

Customer Status Updated
*/
type PutCustomersCustomerFidAccountTypeOK struct {
}

func (o *PutCustomersCustomerFidAccountTypeOK) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/accountType][%d] putCustomersCustomerFidAccountTypeOK ", 200)
}

func (o *PutCustomersCustomerFidAccountTypeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCustomersCustomerFidAccountTypeNotFound creates a PutCustomersCustomerFidAccountTypeNotFound with default headers values
func NewPutCustomersCustomerFidAccountTypeNotFound() *PutCustomersCustomerFidAccountTypeNotFound {
	return &PutCustomersCustomerFidAccountTypeNotFound{}
}

/*PutCustomersCustomerFidAccountTypeNotFound handles this case with default header values.

Customer not found
*/
type PutCustomersCustomerFidAccountTypeNotFound struct {
}

func (o *PutCustomersCustomerFidAccountTypeNotFound) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/accountType][%d] putCustomersCustomerFidAccountTypeNotFound ", 404)
}

func (o *PutCustomersCustomerFidAccountTypeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
