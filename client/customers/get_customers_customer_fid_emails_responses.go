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

// GetCustomersCustomerFidEmailsReader is a Reader for the GetCustomersCustomerFidEmails structure.
type GetCustomersCustomerFidEmailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCustomersCustomerFidEmailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCustomersCustomerFidEmailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetCustomersCustomerFidEmailsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCustomersCustomerFidEmailsOK creates a GetCustomersCustomerFidEmailsOK with default headers values
func NewGetCustomersCustomerFidEmailsOK() *GetCustomersCustomerFidEmailsOK {
	return &GetCustomersCustomerFidEmailsOK{}
}

/*
GetCustomersCustomerFidEmailsOK describes a response with status code 200, with default header values.

List of emails
*/
type GetCustomersCustomerFidEmailsOK struct {
	Payload *models.Emails
}

// IsSuccess returns true when this get customers customer fid emails o k response has a 2xx status code
func (o *GetCustomersCustomerFidEmailsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get customers customer fid emails o k response has a 3xx status code
func (o *GetCustomersCustomerFidEmailsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get customers customer fid emails o k response has a 4xx status code
func (o *GetCustomersCustomerFidEmailsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get customers customer fid emails o k response has a 5xx status code
func (o *GetCustomersCustomerFidEmailsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get customers customer fid emails o k response a status code equal to that given
func (o *GetCustomersCustomerFidEmailsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get customers customer fid emails o k response
func (o *GetCustomersCustomerFidEmailsOK) Code() int {
	return 200
}

func (o *GetCustomersCustomerFidEmailsOK) Error() string {
	return fmt.Sprintf("[GET /customers/{customerFid}/emails][%d] getCustomersCustomerFidEmailsOK  %+v", 200, o.Payload)
}

func (o *GetCustomersCustomerFidEmailsOK) String() string {
	return fmt.Sprintf("[GET /customers/{customerFid}/emails][%d] getCustomersCustomerFidEmailsOK  %+v", 200, o.Payload)
}

func (o *GetCustomersCustomerFidEmailsOK) GetPayload() *models.Emails {
	return o.Payload
}

func (o *GetCustomersCustomerFidEmailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Emails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCustomersCustomerFidEmailsDefault creates a GetCustomersCustomerFidEmailsDefault with default headers values
func NewGetCustomersCustomerFidEmailsDefault(code int) *GetCustomersCustomerFidEmailsDefault {
	return &GetCustomersCustomerFidEmailsDefault{
		_statusCode: code,
	}
}

/*
GetCustomersCustomerFidEmailsDefault describes a response with status code -1, with default header values.

Error
*/
type GetCustomersCustomerFidEmailsDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// IsSuccess returns true when this get customers customer fid emails default response has a 2xx status code
func (o *GetCustomersCustomerFidEmailsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get customers customer fid emails default response has a 3xx status code
func (o *GetCustomersCustomerFidEmailsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get customers customer fid emails default response has a 4xx status code
func (o *GetCustomersCustomerFidEmailsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get customers customer fid emails default response has a 5xx status code
func (o *GetCustomersCustomerFidEmailsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get customers customer fid emails default response a status code equal to that given
func (o *GetCustomersCustomerFidEmailsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get customers customer fid emails default response
func (o *GetCustomersCustomerFidEmailsDefault) Code() int {
	return o._statusCode
}

func (o *GetCustomersCustomerFidEmailsDefault) Error() string {
	return fmt.Sprintf("[GET /customers/{customerFid}/emails][%d] GetCustomersCustomerFidEmails default  %+v", o._statusCode, o.Payload)
}

func (o *GetCustomersCustomerFidEmailsDefault) String() string {
	return fmt.Sprintf("[GET /customers/{customerFid}/emails][%d] GetCustomersCustomerFidEmails default  %+v", o._statusCode, o.Payload)
}

func (o *GetCustomersCustomerFidEmailsDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *GetCustomersCustomerFidEmailsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
