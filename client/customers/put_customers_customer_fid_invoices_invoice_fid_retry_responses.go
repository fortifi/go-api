// Code generated by go-swagger; DO NOT EDIT.

package customers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fortifi/go-api/models"
)

// PutCustomersCustomerFidInvoicesInvoiceFidRetryReader is a Reader for the PutCustomersCustomerFidInvoicesInvoiceFidRetry structure.
type PutCustomersCustomerFidInvoicesInvoiceFidRetryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutCustomersCustomerFidInvoicesInvoiceFidRetryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutCustomersCustomerFidInvoicesInvoiceFidRetryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutCustomersCustomerFidInvoicesInvoiceFidRetryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutCustomersCustomerFidInvoicesInvoiceFidRetryOK creates a PutCustomersCustomerFidInvoicesInvoiceFidRetryOK with default headers values
func NewPutCustomersCustomerFidInvoicesInvoiceFidRetryOK() *PutCustomersCustomerFidInvoicesInvoiceFidRetryOK {
	return &PutCustomersCustomerFidInvoicesInvoiceFidRetryOK{}
}

/*
PutCustomersCustomerFidInvoicesInvoiceFidRetryOK describes a response with status code 200, with default header values.

Invoice set to retry payment
*/
type PutCustomersCustomerFidInvoicesInvoiceFidRetryOK struct {
}

// IsSuccess returns true when this put customers customer fid invoices invoice fid retry o k response has a 2xx status code
func (o *PutCustomersCustomerFidInvoicesInvoiceFidRetryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put customers customer fid invoices invoice fid retry o k response has a 3xx status code
func (o *PutCustomersCustomerFidInvoicesInvoiceFidRetryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put customers customer fid invoices invoice fid retry o k response has a 4xx status code
func (o *PutCustomersCustomerFidInvoicesInvoiceFidRetryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put customers customer fid invoices invoice fid retry o k response has a 5xx status code
func (o *PutCustomersCustomerFidInvoicesInvoiceFidRetryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put customers customer fid invoices invoice fid retry o k response a status code equal to that given
func (o *PutCustomersCustomerFidInvoicesInvoiceFidRetryOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the put customers customer fid invoices invoice fid retry o k response
func (o *PutCustomersCustomerFidInvoicesInvoiceFidRetryOK) Code() int {
	return 200
}

func (o *PutCustomersCustomerFidInvoicesInvoiceFidRetryOK) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/invoices/{invoiceFid}/retry][%d] putCustomersCustomerFidInvoicesInvoiceFidRetryOK ", 200)
}

func (o *PutCustomersCustomerFidInvoicesInvoiceFidRetryOK) String() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/invoices/{invoiceFid}/retry][%d] putCustomersCustomerFidInvoicesInvoiceFidRetryOK ", 200)
}

func (o *PutCustomersCustomerFidInvoicesInvoiceFidRetryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCustomersCustomerFidInvoicesInvoiceFidRetryDefault creates a PutCustomersCustomerFidInvoicesInvoiceFidRetryDefault with default headers values
func NewPutCustomersCustomerFidInvoicesInvoiceFidRetryDefault(code int) *PutCustomersCustomerFidInvoicesInvoiceFidRetryDefault {
	return &PutCustomersCustomerFidInvoicesInvoiceFidRetryDefault{
		_statusCode: code,
	}
}

/*
PutCustomersCustomerFidInvoicesInvoiceFidRetryDefault describes a response with status code -1, with default header values.

Error
*/
type PutCustomersCustomerFidInvoicesInvoiceFidRetryDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// IsSuccess returns true when this put customers customer fid invoices invoice fid retry default response has a 2xx status code
func (o *PutCustomersCustomerFidInvoicesInvoiceFidRetryDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this put customers customer fid invoices invoice fid retry default response has a 3xx status code
func (o *PutCustomersCustomerFidInvoicesInvoiceFidRetryDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this put customers customer fid invoices invoice fid retry default response has a 4xx status code
func (o *PutCustomersCustomerFidInvoicesInvoiceFidRetryDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this put customers customer fid invoices invoice fid retry default response has a 5xx status code
func (o *PutCustomersCustomerFidInvoicesInvoiceFidRetryDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this put customers customer fid invoices invoice fid retry default response a status code equal to that given
func (o *PutCustomersCustomerFidInvoicesInvoiceFidRetryDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the put customers customer fid invoices invoice fid retry default response
func (o *PutCustomersCustomerFidInvoicesInvoiceFidRetryDefault) Code() int {
	return o._statusCode
}

func (o *PutCustomersCustomerFidInvoicesInvoiceFidRetryDefault) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/invoices/{invoiceFid}/retry][%d] PutCustomersCustomerFidInvoicesInvoiceFidRetry default  %+v", o._statusCode, o.Payload)
}

func (o *PutCustomersCustomerFidInvoicesInvoiceFidRetryDefault) String() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/invoices/{invoiceFid}/retry][%d] PutCustomersCustomerFidInvoicesInvoiceFidRetry default  %+v", o._statusCode, o.Payload)
}

func (o *PutCustomersCustomerFidInvoicesInvoiceFidRetryDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PutCustomersCustomerFidInvoicesInvoiceFidRetryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
