// Code generated by go-swagger; DO NOT EDIT.

package customers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PutCustomersCustomerFidVipReader is a Reader for the PutCustomersCustomerFidVip structure.
type PutCustomersCustomerFidVipReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutCustomersCustomerFidVipReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutCustomersCustomerFidVipOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewPutCustomersCustomerFidVipNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutCustomersCustomerFidVipOK creates a PutCustomersCustomerFidVipOK with default headers values
func NewPutCustomersCustomerFidVipOK() *PutCustomersCustomerFidVipOK {
	return &PutCustomersCustomerFidVipOK{}
}

/*PutCustomersCustomerFidVipOK handles this case with default header values.

Customer Now VIP
*/
type PutCustomersCustomerFidVipOK struct {
}

func (o *PutCustomersCustomerFidVipOK) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/vip][%d] putCustomersCustomerFidVipOK ", 200)
}

func (o *PutCustomersCustomerFidVipOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCustomersCustomerFidVipNotFound creates a PutCustomersCustomerFidVipNotFound with default headers values
func NewPutCustomersCustomerFidVipNotFound() *PutCustomersCustomerFidVipNotFound {
	return &PutCustomersCustomerFidVipNotFound{}
}

/*PutCustomersCustomerFidVipNotFound handles this case with default header values.

Customer not found
*/
type PutCustomersCustomerFidVipNotFound struct {
}

func (o *PutCustomersCustomerFidVipNotFound) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/vip][%d] putCustomersCustomerFidVipNotFound ", 404)
}

func (o *PutCustomersCustomerFidVipNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}