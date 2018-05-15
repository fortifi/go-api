// Code generated by go-swagger; DO NOT EDIT.

package service_status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/fortifi/go-api/models"
)

// GetServicesServiceFidIncidentsReader is a Reader for the GetServicesServiceFidIncidents structure.
type GetServicesServiceFidIncidentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServicesServiceFidIncidentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetServicesServiceFidIncidentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetServicesServiceFidIncidentsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetServicesServiceFidIncidentsOK creates a GetServicesServiceFidIncidentsOK with default headers values
func NewGetServicesServiceFidIncidentsOK() *GetServicesServiceFidIncidentsOK {
	return &GetServicesServiceFidIncidentsOK{}
}

/*GetServicesServiceFidIncidentsOK handles this case with default header values.

Service incidents
*/
type GetServicesServiceFidIncidentsOK struct {
	Payload *models.GetServicesServiceFidIncidentsOKBody
}

func (o *GetServicesServiceFidIncidentsOK) Error() string {
	return fmt.Sprintf("[GET /services/{serviceFid}/incidents][%d] getServicesServiceFidIncidentsOK  %+v", 200, o.Payload)
}

func (o *GetServicesServiceFidIncidentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetServicesServiceFidIncidentsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServicesServiceFidIncidentsDefault creates a GetServicesServiceFidIncidentsDefault with default headers values
func NewGetServicesServiceFidIncidentsDefault(code int) *GetServicesServiceFidIncidentsDefault {
	return &GetServicesServiceFidIncidentsDefault{
		_statusCode: code,
	}
}

/*GetServicesServiceFidIncidentsDefault handles this case with default header values.

Error
*/
type GetServicesServiceFidIncidentsDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the get services service fid incidents default response
func (o *GetServicesServiceFidIncidentsDefault) Code() int {
	return o._statusCode
}

func (o *GetServicesServiceFidIncidentsDefault) Error() string {
	return fmt.Sprintf("[GET /services/{serviceFid}/incidents][%d] GetServicesServiceFidIncidents default  %+v", o._statusCode, o.Payload)
}

func (o *GetServicesServiceFidIncidentsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
