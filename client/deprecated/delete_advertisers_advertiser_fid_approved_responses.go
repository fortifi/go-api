// Code generated by go-swagger; DO NOT EDIT.

package deprecated

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fortifi/go-api/models"
)

// DeleteAdvertisersAdvertiserFidApprovedReader is a Reader for the DeleteAdvertisersAdvertiserFidApproved structure.
type DeleteAdvertisersAdvertiserFidApprovedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAdvertisersAdvertiserFidApprovedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAdvertisersAdvertiserFidApprovedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteAdvertisersAdvertiserFidApprovedDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteAdvertisersAdvertiserFidApprovedOK creates a DeleteAdvertisersAdvertiserFidApprovedOK with default headers values
func NewDeleteAdvertisersAdvertiserFidApprovedOK() *DeleteAdvertisersAdvertiserFidApprovedOK {
	return &DeleteAdvertisersAdvertiserFidApprovedOK{}
}

/*
DeleteAdvertisersAdvertiserFidApprovedOK describes a response with status code 200, with default header values.

Advertiser no longer approved
*/
type DeleteAdvertisersAdvertiserFidApprovedOK struct {
}

// IsSuccess returns true when this delete advertisers advertiser fid approved o k response has a 2xx status code
func (o *DeleteAdvertisersAdvertiserFidApprovedOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete advertisers advertiser fid approved o k response has a 3xx status code
func (o *DeleteAdvertisersAdvertiserFidApprovedOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete advertisers advertiser fid approved o k response has a 4xx status code
func (o *DeleteAdvertisersAdvertiserFidApprovedOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete advertisers advertiser fid approved o k response has a 5xx status code
func (o *DeleteAdvertisersAdvertiserFidApprovedOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete advertisers advertiser fid approved o k response a status code equal to that given
func (o *DeleteAdvertisersAdvertiserFidApprovedOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete advertisers advertiser fid approved o k response
func (o *DeleteAdvertisersAdvertiserFidApprovedOK) Code() int {
	return 200
}

func (o *DeleteAdvertisersAdvertiserFidApprovedOK) Error() string {
	return fmt.Sprintf("[DELETE /advertisers/{advertiserFid}/approved][%d] deleteAdvertisersAdvertiserFidApprovedOK ", 200)
}

func (o *DeleteAdvertisersAdvertiserFidApprovedOK) String() string {
	return fmt.Sprintf("[DELETE /advertisers/{advertiserFid}/approved][%d] deleteAdvertisersAdvertiserFidApprovedOK ", 200)
}

func (o *DeleteAdvertisersAdvertiserFidApprovedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAdvertisersAdvertiserFidApprovedDefault creates a DeleteAdvertisersAdvertiserFidApprovedDefault with default headers values
func NewDeleteAdvertisersAdvertiserFidApprovedDefault(code int) *DeleteAdvertisersAdvertiserFidApprovedDefault {
	return &DeleteAdvertisersAdvertiserFidApprovedDefault{
		_statusCode: code,
	}
}

/*
DeleteAdvertisersAdvertiserFidApprovedDefault describes a response with status code -1, with default header values.

Error
*/
type DeleteAdvertisersAdvertiserFidApprovedDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// IsSuccess returns true when this delete advertisers advertiser fid approved default response has a 2xx status code
func (o *DeleteAdvertisersAdvertiserFidApprovedDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete advertisers advertiser fid approved default response has a 3xx status code
func (o *DeleteAdvertisersAdvertiserFidApprovedDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete advertisers advertiser fid approved default response has a 4xx status code
func (o *DeleteAdvertisersAdvertiserFidApprovedDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete advertisers advertiser fid approved default response has a 5xx status code
func (o *DeleteAdvertisersAdvertiserFidApprovedDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete advertisers advertiser fid approved default response a status code equal to that given
func (o *DeleteAdvertisersAdvertiserFidApprovedDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete advertisers advertiser fid approved default response
func (o *DeleteAdvertisersAdvertiserFidApprovedDefault) Code() int {
	return o._statusCode
}

func (o *DeleteAdvertisersAdvertiserFidApprovedDefault) Error() string {
	return fmt.Sprintf("[DELETE /advertisers/{advertiserFid}/approved][%d] DeleteAdvertisersAdvertiserFidApproved default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteAdvertisersAdvertiserFidApprovedDefault) String() string {
	return fmt.Sprintf("[DELETE /advertisers/{advertiserFid}/approved][%d] DeleteAdvertisersAdvertiserFidApproved default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteAdvertisersAdvertiserFidApprovedDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *DeleteAdvertisersAdvertiserFidApprovedDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
