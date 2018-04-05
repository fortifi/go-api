// Code generated by go-swagger; DO NOT EDIT.

package customers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/fortifi/go-api/models"
)

// GetCustomersCustomerFidInvoicesReader is a Reader for the GetCustomersCustomerFidInvoices structure.
type GetCustomersCustomerFidInvoicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCustomersCustomerFidInvoicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCustomersCustomerFidInvoicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCustomersCustomerFidInvoicesOK creates a GetCustomersCustomerFidInvoicesOK with default headers values
func NewGetCustomersCustomerFidInvoicesOK() *GetCustomersCustomerFidInvoicesOK {
	return &GetCustomersCustomerFidInvoicesOK{}
}

/*GetCustomersCustomerFidInvoicesOK handles this case with default header values.

List of invoices summaries
*/
type GetCustomersCustomerFidInvoicesOK struct {
	Payload *models.Invoices
}

func (o *GetCustomersCustomerFidInvoicesOK) Error() string {
	return fmt.Sprintf("[GET /customers/{customerFid}/invoices][%d] getCustomersCustomerFidInvoicesOK  %+v", 200, o.Payload)
}

func (o *GetCustomersCustomerFidInvoicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Invoices)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
