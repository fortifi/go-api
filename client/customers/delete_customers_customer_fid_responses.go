// Code generated by go-swagger; DO NOT EDIT.

package customers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteCustomersCustomerFidReader is a Reader for the DeleteCustomersCustomerFid structure.
type DeleteCustomersCustomerFidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCustomersCustomerFidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteCustomersCustomerFidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewDeleteCustomersCustomerFidForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteCustomersCustomerFidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteCustomersCustomerFidOK creates a DeleteCustomersCustomerFidOK with default headers values
func NewDeleteCustomersCustomerFidOK() *DeleteCustomersCustomerFidOK {
	return &DeleteCustomersCustomerFidOK{}
}

/*DeleteCustomersCustomerFidOK handles this case with default header values.

Customer Archived
*/
type DeleteCustomersCustomerFidOK struct {
}

func (o *DeleteCustomersCustomerFidOK) Error() string {
	return fmt.Sprintf("[DELETE /customers/{customerFid}][%d] deleteCustomersCustomerFidOK ", 200)
}

func (o *DeleteCustomersCustomerFidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCustomersCustomerFidForbidden creates a DeleteCustomersCustomerFidForbidden with default headers values
func NewDeleteCustomersCustomerFidForbidden() *DeleteCustomersCustomerFidForbidden {
	return &DeleteCustomersCustomerFidForbidden{}
}

/*DeleteCustomersCustomerFidForbidden handles this case with default header values.

It is not possible to archive this customer
*/
type DeleteCustomersCustomerFidForbidden struct {
}

func (o *DeleteCustomersCustomerFidForbidden) Error() string {
	return fmt.Sprintf("[DELETE /customers/{customerFid}][%d] deleteCustomersCustomerFidForbidden ", 403)
}

func (o *DeleteCustomersCustomerFidForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCustomersCustomerFidNotFound creates a DeleteCustomersCustomerFidNotFound with default headers values
func NewDeleteCustomersCustomerFidNotFound() *DeleteCustomersCustomerFidNotFound {
	return &DeleteCustomersCustomerFidNotFound{}
}

/*DeleteCustomersCustomerFidNotFound handles this case with default header values.

Customer not found
*/
type DeleteCustomersCustomerFidNotFound struct {
}

func (o *DeleteCustomersCustomerFidNotFound) Error() string {
	return fmt.Sprintf("[DELETE /customers/{customerFid}][%d] deleteCustomersCustomerFidNotFound ", 404)
}

func (o *DeleteCustomersCustomerFidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}