// Code generated by go-swagger; DO NOT EDIT.

package marketing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/fortifi/go-api/models"
)

// PutAdvertisersAdvertiserFidApprovedReader is a Reader for the PutAdvertisersAdvertiserFidApproved structure.
type PutAdvertisersAdvertiserFidApprovedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutAdvertisersAdvertiserFidApprovedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutAdvertisersAdvertiserFidApprovedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPutAdvertisersAdvertiserFidApprovedDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutAdvertisersAdvertiserFidApprovedOK creates a PutAdvertisersAdvertiserFidApprovedOK with default headers values
func NewPutAdvertisersAdvertiserFidApprovedOK() *PutAdvertisersAdvertiserFidApprovedOK {
	return &PutAdvertisersAdvertiserFidApprovedOK{}
}

/*PutAdvertisersAdvertiserFidApprovedOK handles this case with default header values.

Approved Advertiser
*/
type PutAdvertisersAdvertiserFidApprovedOK struct {
}

func (o *PutAdvertisersAdvertiserFidApprovedOK) Error() string {
	return fmt.Sprintf("[PUT /advertisers/{advertiserFid}/approved][%d] putAdvertisersAdvertiserFidApprovedOK ", 200)
}

func (o *PutAdvertisersAdvertiserFidApprovedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutAdvertisersAdvertiserFidApprovedDefault creates a PutAdvertisersAdvertiserFidApprovedDefault with default headers values
func NewPutAdvertisersAdvertiserFidApprovedDefault(code int) *PutAdvertisersAdvertiserFidApprovedDefault {
	return &PutAdvertisersAdvertiserFidApprovedDefault{
		_statusCode: code,
	}
}

/*PutAdvertisersAdvertiserFidApprovedDefault handles this case with default header values.

Error
*/
type PutAdvertisersAdvertiserFidApprovedDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the put advertisers advertiser fid approved default response
func (o *PutAdvertisersAdvertiserFidApprovedDefault) Code() int {
	return o._statusCode
}

func (o *PutAdvertisersAdvertiserFidApprovedDefault) Error() string {
	return fmt.Sprintf("[PUT /advertisers/{advertiserFid}/approved][%d] PutAdvertisersAdvertiserFidApproved default  %+v", o._statusCode, o.Payload)
}

func (o *PutAdvertisersAdvertiserFidApprovedDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
