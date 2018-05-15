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

// PostCustomersCustomerFidSarReader is a Reader for the PostCustomersCustomerFidSar structure.
type PostCustomersCustomerFidSarReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCustomersCustomerFidSarReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostCustomersCustomerFidSarOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostCustomersCustomerFidSarDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostCustomersCustomerFidSarOK creates a PostCustomersCustomerFidSarOK with default headers values
func NewPostCustomersCustomerFidSarOK() *PostCustomersCustomerFidSarOK {
	return &PostCustomersCustomerFidSarOK{}
}

/*PostCustomersCustomerFidSarOK handles this case with default header values.

Subject Access Request
*/
type PostCustomersCustomerFidSarOK struct {
	Payload *models.PostCustomersCustomerFidSarOKBody
}

func (o *PostCustomersCustomerFidSarOK) Error() string {
	return fmt.Sprintf("[POST /customers/{customerFid}/sar][%d] postCustomersCustomerFidSarOK  %+v", 200, o.Payload)
}

func (o *PostCustomersCustomerFidSarOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostCustomersCustomerFidSarOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCustomersCustomerFidSarDefault creates a PostCustomersCustomerFidSarDefault with default headers values
func NewPostCustomersCustomerFidSarDefault(code int) *PostCustomersCustomerFidSarDefault {
	return &PostCustomersCustomerFidSarDefault{
		_statusCode: code,
	}
}

/*PostCustomersCustomerFidSarDefault handles this case with default header values.

Error
*/
type PostCustomersCustomerFidSarDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the post customers customer fid sar default response
func (o *PostCustomersCustomerFidSarDefault) Code() int {
	return o._statusCode
}

func (o *PostCustomersCustomerFidSarDefault) Error() string {
	return fmt.Sprintf("[POST /customers/{customerFid}/sar][%d] PostCustomersCustomerFidSar default  %+v", o._statusCode, o.Payload)
}

func (o *PostCustomersCustomerFidSarDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
