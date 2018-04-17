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

// GetAdvertisersReader is a Reader for the GetAdvertisers structure.
type GetAdvertisersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAdvertisersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAdvertisersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAdvertisersOK creates a GetAdvertisersOK with default headers values
func NewGetAdvertisersOK() *GetAdvertisersOK {
	return &GetAdvertisersOK{}
}

/*GetAdvertisersOK handles this case with default header values.

List of advertisers
*/
type GetAdvertisersOK struct {
	Payload *models.Advertisers
}

func (o *GetAdvertisersOK) Error() string {
	return fmt.Sprintf("[GET /advertisers][%d] getAdvertisersOK  %+v", 200, o.Payload)
}

func (o *GetAdvertisersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Advertisers)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
