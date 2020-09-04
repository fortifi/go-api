// Code generated by go-swagger; DO NOT EDIT.

package marketing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/fortifi/go-api/models"
)

// PutPublishersPublisherFidRestoreReader is a Reader for the PutPublishersPublisherFidRestore structure.
type PutPublishersPublisherFidRestoreReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutPublishersPublisherFidRestoreReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutPublishersPublisherFidRestoreOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutPublishersPublisherFidRestoreDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutPublishersPublisherFidRestoreOK creates a PutPublishersPublisherFidRestoreOK with default headers values
func NewPutPublishersPublisherFidRestoreOK() *PutPublishersPublisherFidRestoreOK {
	return &PutPublishersPublisherFidRestoreOK{}
}

/*PutPublishersPublisherFidRestoreOK handles this case with default header values.

Publisher restored
*/
type PutPublishersPublisherFidRestoreOK struct {
}

func (o *PutPublishersPublisherFidRestoreOK) Error() string {
	return fmt.Sprintf("[PUT /publishers/{publisherFid}/restore][%d] putPublishersPublisherFidRestoreOK ", 200)
}

func (o *PutPublishersPublisherFidRestoreOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutPublishersPublisherFidRestoreDefault creates a PutPublishersPublisherFidRestoreDefault with default headers values
func NewPutPublishersPublisherFidRestoreDefault(code int) *PutPublishersPublisherFidRestoreDefault {
	return &PutPublishersPublisherFidRestoreDefault{
		_statusCode: code,
	}
}

/*PutPublishersPublisherFidRestoreDefault handles this case with default header values.

Error
*/
type PutPublishersPublisherFidRestoreDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the put publishers publisher fid restore default response
func (o *PutPublishersPublisherFidRestoreDefault) Code() int {
	return o._statusCode
}

func (o *PutPublishersPublisherFidRestoreDefault) Error() string {
	return fmt.Sprintf("[PUT /publishers/{publisherFid}/restore][%d] PutPublishersPublisherFidRestore default  %+v", o._statusCode, o.Payload)
}

func (o *PutPublishersPublisherFidRestoreDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PutPublishersPublisherFidRestoreDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}