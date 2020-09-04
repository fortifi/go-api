// Code generated by go-swagger; DO NOT EDIT.

package orders

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fortifi/go-api/models"
)

// PutOrdersOrderFidFinalizeReader is a Reader for the PutOrdersOrderFidFinalize structure.
type PutOrdersOrderFidFinalizeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutOrdersOrderFidFinalizeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutOrdersOrderFidFinalizeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutOrdersOrderFidFinalizeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutOrdersOrderFidFinalizeOK creates a PutOrdersOrderFidFinalizeOK with default headers values
func NewPutOrdersOrderFidFinalizeOK() *PutOrdersOrderFidFinalizeOK {
	return &PutOrdersOrderFidFinalizeOK{}
}

/*PutOrdersOrderFidFinalizeOK handles this case with default header values.

Order Finalized
*/
type PutOrdersOrderFidFinalizeOK struct {
}

func (o *PutOrdersOrderFidFinalizeOK) Error() string {
	return fmt.Sprintf("[PUT /orders/{orderFid}/finalize][%d] putOrdersOrderFidFinalizeOK ", 200)
}

func (o *PutOrdersOrderFidFinalizeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutOrdersOrderFidFinalizeDefault creates a PutOrdersOrderFidFinalizeDefault with default headers values
func NewPutOrdersOrderFidFinalizeDefault(code int) *PutOrdersOrderFidFinalizeDefault {
	return &PutOrdersOrderFidFinalizeDefault{
		_statusCode: code,
	}
}

/*PutOrdersOrderFidFinalizeDefault handles this case with default header values.

Error
*/
type PutOrdersOrderFidFinalizeDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the put orders order fid finalize default response
func (o *PutOrdersOrderFidFinalizeDefault) Code() int {
	return o._statusCode
}

func (o *PutOrdersOrderFidFinalizeDefault) Error() string {
	return fmt.Sprintf("[PUT /orders/{orderFid}/finalize][%d] PutOrdersOrderFidFinalize default  %+v", o._statusCode, o.Payload)
}

func (o *PutOrdersOrderFidFinalizeDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PutOrdersOrderFidFinalizeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
