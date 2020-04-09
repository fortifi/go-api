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

// PutCustomersCustomerFidPaymentsPaymentFidRefundReader is a Reader for the PutCustomersCustomerFidPaymentsPaymentFidRefund structure.
type PutCustomersCustomerFidPaymentsPaymentFidRefundReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutCustomersCustomerFidPaymentsPaymentFidRefundReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutCustomersCustomerFidPaymentsPaymentFidRefundOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutCustomersCustomerFidPaymentsPaymentFidRefundDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutCustomersCustomerFidPaymentsPaymentFidRefundOK creates a PutCustomersCustomerFidPaymentsPaymentFidRefundOK with default headers values
func NewPutCustomersCustomerFidPaymentsPaymentFidRefundOK() *PutCustomersCustomerFidPaymentsPaymentFidRefundOK {
	return &PutCustomersCustomerFidPaymentsPaymentFidRefundOK{}
}

/*PutCustomersCustomerFidPaymentsPaymentFidRefundOK handles this case with default header values.

Payment Refunded
*/
type PutCustomersCustomerFidPaymentsPaymentFidRefundOK struct {
}

func (o *PutCustomersCustomerFidPaymentsPaymentFidRefundOK) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/payments/{paymentFid}/refund][%d] putCustomersCustomerFidPaymentsPaymentFidRefundOK ", 200)
}

func (o *PutCustomersCustomerFidPaymentsPaymentFidRefundOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCustomersCustomerFidPaymentsPaymentFidRefundDefault creates a PutCustomersCustomerFidPaymentsPaymentFidRefundDefault with default headers values
func NewPutCustomersCustomerFidPaymentsPaymentFidRefundDefault(code int) *PutCustomersCustomerFidPaymentsPaymentFidRefundDefault {
	return &PutCustomersCustomerFidPaymentsPaymentFidRefundDefault{
		_statusCode: code,
	}
}

/*PutCustomersCustomerFidPaymentsPaymentFidRefundDefault handles this case with default header values.

Error
*/
type PutCustomersCustomerFidPaymentsPaymentFidRefundDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the put customers customer fid payments payment fid refund default response
func (o *PutCustomersCustomerFidPaymentsPaymentFidRefundDefault) Code() int {
	return o._statusCode
}

func (o *PutCustomersCustomerFidPaymentsPaymentFidRefundDefault) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/payments/{paymentFid}/refund][%d] PutCustomersCustomerFidPaymentsPaymentFidRefund default  %+v", o._statusCode, o.Payload)
}

func (o *PutCustomersCustomerFidPaymentsPaymentFidRefundDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PutCustomersCustomerFidPaymentsPaymentFidRefundDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}