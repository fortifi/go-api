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

// GetCustomersCustomerFidOrdersReader is a Reader for the GetCustomersCustomerFidOrders structure.
type GetCustomersCustomerFidOrdersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCustomersCustomerFidOrdersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCustomersCustomerFidOrdersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetCustomersCustomerFidOrdersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCustomersCustomerFidOrdersOK creates a GetCustomersCustomerFidOrdersOK with default headers values
func NewGetCustomersCustomerFidOrdersOK() *GetCustomersCustomerFidOrdersOK {
	return &GetCustomersCustomerFidOrdersOK{}
}

/*
GetCustomersCustomerFidOrdersOK describes a response with status code 200, with default header values.

Customer Orders
*/
type GetCustomersCustomerFidOrdersOK struct {
	Payload *models.Orders
}

// IsSuccess returns true when this get customers customer fid orders o k response has a 2xx status code
func (o *GetCustomersCustomerFidOrdersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get customers customer fid orders o k response has a 3xx status code
func (o *GetCustomersCustomerFidOrdersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get customers customer fid orders o k response has a 4xx status code
func (o *GetCustomersCustomerFidOrdersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get customers customer fid orders o k response has a 5xx status code
func (o *GetCustomersCustomerFidOrdersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get customers customer fid orders o k response a status code equal to that given
func (o *GetCustomersCustomerFidOrdersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get customers customer fid orders o k response
func (o *GetCustomersCustomerFidOrdersOK) Code() int {
	return 200
}

func (o *GetCustomersCustomerFidOrdersOK) Error() string {
	return fmt.Sprintf("[GET /customers/{customerFid}/orders][%d] getCustomersCustomerFidOrdersOK  %+v", 200, o.Payload)
}

func (o *GetCustomersCustomerFidOrdersOK) String() string {
	return fmt.Sprintf("[GET /customers/{customerFid}/orders][%d] getCustomersCustomerFidOrdersOK  %+v", 200, o.Payload)
}

func (o *GetCustomersCustomerFidOrdersOK) GetPayload() *models.Orders {
	return o.Payload
}

func (o *GetCustomersCustomerFidOrdersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Orders)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCustomersCustomerFidOrdersDefault creates a GetCustomersCustomerFidOrdersDefault with default headers values
func NewGetCustomersCustomerFidOrdersDefault(code int) *GetCustomersCustomerFidOrdersDefault {
	return &GetCustomersCustomerFidOrdersDefault{
		_statusCode: code,
	}
}

/*
GetCustomersCustomerFidOrdersDefault describes a response with status code -1, with default header values.

Error
*/
type GetCustomersCustomerFidOrdersDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// IsSuccess returns true when this get customers customer fid orders default response has a 2xx status code
func (o *GetCustomersCustomerFidOrdersDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get customers customer fid orders default response has a 3xx status code
func (o *GetCustomersCustomerFidOrdersDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get customers customer fid orders default response has a 4xx status code
func (o *GetCustomersCustomerFidOrdersDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get customers customer fid orders default response has a 5xx status code
func (o *GetCustomersCustomerFidOrdersDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get customers customer fid orders default response a status code equal to that given
func (o *GetCustomersCustomerFidOrdersDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get customers customer fid orders default response
func (o *GetCustomersCustomerFidOrdersDefault) Code() int {
	return o._statusCode
}

func (o *GetCustomersCustomerFidOrdersDefault) Error() string {
	return fmt.Sprintf("[GET /customers/{customerFid}/orders][%d] GetCustomersCustomerFidOrders default  %+v", o._statusCode, o.Payload)
}

func (o *GetCustomersCustomerFidOrdersDefault) String() string {
	return fmt.Sprintf("[GET /customers/{customerFid}/orders][%d] GetCustomersCustomerFidOrders default  %+v", o._statusCode, o.Payload)
}

func (o *GetCustomersCustomerFidOrdersDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *GetCustomersCustomerFidOrdersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
