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

// PutCustomersCustomerFidBillingDataReader is a Reader for the PutCustomersCustomerFidBillingData structure.
type PutCustomersCustomerFidBillingDataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutCustomersCustomerFidBillingDataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutCustomersCustomerFidBillingDataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutCustomersCustomerFidBillingDataDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutCustomersCustomerFidBillingDataOK creates a PutCustomersCustomerFidBillingDataOK with default headers values
func NewPutCustomersCustomerFidBillingDataOK() *PutCustomersCustomerFidBillingDataOK {
	return &PutCustomersCustomerFidBillingDataOK{}
}

/*
PutCustomersCustomerFidBillingDataOK describes a response with status code 200, with default header values.

Customer Status Updated
*/
type PutCustomersCustomerFidBillingDataOK struct {
}

// IsSuccess returns true when this put customers customer fid billing data o k response has a 2xx status code
func (o *PutCustomersCustomerFidBillingDataOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put customers customer fid billing data o k response has a 3xx status code
func (o *PutCustomersCustomerFidBillingDataOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put customers customer fid billing data o k response has a 4xx status code
func (o *PutCustomersCustomerFidBillingDataOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put customers customer fid billing data o k response has a 5xx status code
func (o *PutCustomersCustomerFidBillingDataOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put customers customer fid billing data o k response a status code equal to that given
func (o *PutCustomersCustomerFidBillingDataOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the put customers customer fid billing data o k response
func (o *PutCustomersCustomerFidBillingDataOK) Code() int {
	return 200
}

func (o *PutCustomersCustomerFidBillingDataOK) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/billingData][%d] putCustomersCustomerFidBillingDataOK ", 200)
}

func (o *PutCustomersCustomerFidBillingDataOK) String() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/billingData][%d] putCustomersCustomerFidBillingDataOK ", 200)
}

func (o *PutCustomersCustomerFidBillingDataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCustomersCustomerFidBillingDataDefault creates a PutCustomersCustomerFidBillingDataDefault with default headers values
func NewPutCustomersCustomerFidBillingDataDefault(code int) *PutCustomersCustomerFidBillingDataDefault {
	return &PutCustomersCustomerFidBillingDataDefault{
		_statusCode: code,
	}
}

/*
PutCustomersCustomerFidBillingDataDefault describes a response with status code -1, with default header values.

Error
*/
type PutCustomersCustomerFidBillingDataDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// IsSuccess returns true when this put customers customer fid billing data default response has a 2xx status code
func (o *PutCustomersCustomerFidBillingDataDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this put customers customer fid billing data default response has a 3xx status code
func (o *PutCustomersCustomerFidBillingDataDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this put customers customer fid billing data default response has a 4xx status code
func (o *PutCustomersCustomerFidBillingDataDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this put customers customer fid billing data default response has a 5xx status code
func (o *PutCustomersCustomerFidBillingDataDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this put customers customer fid billing data default response a status code equal to that given
func (o *PutCustomersCustomerFidBillingDataDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the put customers customer fid billing data default response
func (o *PutCustomersCustomerFidBillingDataDefault) Code() int {
	return o._statusCode
}

func (o *PutCustomersCustomerFidBillingDataDefault) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/billingData][%d] PutCustomersCustomerFidBillingData default  %+v", o._statusCode, o.Payload)
}

func (o *PutCustomersCustomerFidBillingDataDefault) String() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/billingData][%d] PutCustomersCustomerFidBillingData default  %+v", o._statusCode, o.Payload)
}

func (o *PutCustomersCustomerFidBillingDataDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PutCustomersCustomerFidBillingDataDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
