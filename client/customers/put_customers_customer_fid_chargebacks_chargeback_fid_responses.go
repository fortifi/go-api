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

// PutCustomersCustomerFidChargebacksChargebackFidReader is a Reader for the PutCustomersCustomerFidChargebacksChargebackFid structure.
type PutCustomersCustomerFidChargebacksChargebackFidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutCustomersCustomerFidChargebacksChargebackFidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutCustomersCustomerFidChargebacksChargebackFidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPutCustomersCustomerFidChargebacksChargebackFidBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutCustomersCustomerFidChargebacksChargebackFidOK creates a PutCustomersCustomerFidChargebacksChargebackFidOK with default headers values
func NewPutCustomersCustomerFidChargebacksChargebackFidOK() *PutCustomersCustomerFidChargebacksChargebackFidOK {
	return &PutCustomersCustomerFidChargebacksChargebackFidOK{}
}

/*PutCustomersCustomerFidChargebacksChargebackFidOK handles this case with default header values.

Chargeback Actioned
*/
type PutCustomersCustomerFidChargebacksChargebackFidOK struct {
	Payload *models.BoolMessage
}

func (o *PutCustomersCustomerFidChargebacksChargebackFidOK) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/chargebacks/{chargebackFid}][%d] putCustomersCustomerFidChargebacksChargebackFidOK  %+v", 200, o.Payload)
}

func (o *PutCustomersCustomerFidChargebacksChargebackFidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BoolMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutCustomersCustomerFidChargebacksChargebackFidBadRequest creates a PutCustomersCustomerFidChargebacksChargebackFidBadRequest with default headers values
func NewPutCustomersCustomerFidChargebacksChargebackFidBadRequest() *PutCustomersCustomerFidChargebacksChargebackFidBadRequest {
	return &PutCustomersCustomerFidChargebacksChargebackFidBadRequest{}
}

/*PutCustomersCustomerFidChargebacksChargebackFidBadRequest handles this case with default header values.

Invalid Payload
*/
type PutCustomersCustomerFidChargebacksChargebackFidBadRequest struct {
}

func (o *PutCustomersCustomerFidChargebacksChargebackFidBadRequest) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/chargebacks/{chargebackFid}][%d] putCustomersCustomerFidChargebacksChargebackFidBadRequest ", 400)
}

func (o *PutCustomersCustomerFidChargebacksChargebackFidBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}