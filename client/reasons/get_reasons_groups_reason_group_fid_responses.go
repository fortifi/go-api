// Code generated by go-swagger; DO NOT EDIT.

package reasons

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/fortifi/go-api/models"
)

// GetReasonsGroupsReasonGroupFidReader is a Reader for the GetReasonsGroupsReasonGroupFid structure.
type GetReasonsGroupsReasonGroupFidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReasonsGroupsReasonGroupFidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetReasonsGroupsReasonGroupFidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetReasonsGroupsReasonGroupFidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetReasonsGroupsReasonGroupFidOK creates a GetReasonsGroupsReasonGroupFidOK with default headers values
func NewGetReasonsGroupsReasonGroupFidOK() *GetReasonsGroupsReasonGroupFidOK {
	return &GetReasonsGroupsReasonGroupFidOK{}
}

/*GetReasonsGroupsReasonGroupFidOK handles this case with default header values.

Reason Group retrieved
*/
type GetReasonsGroupsReasonGroupFidOK struct {
	Payload *models.Reasons
}

func (o *GetReasonsGroupsReasonGroupFidOK) Error() string {
	return fmt.Sprintf("[GET /reasons/groups/{reasonGroupFid}][%d] getReasonsGroupsReasonGroupFidOK  %+v", 200, o.Payload)
}

func (o *GetReasonsGroupsReasonGroupFidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Reasons)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReasonsGroupsReasonGroupFidNotFound creates a GetReasonsGroupsReasonGroupFidNotFound with default headers values
func NewGetReasonsGroupsReasonGroupFidNotFound() *GetReasonsGroupsReasonGroupFidNotFound {
	return &GetReasonsGroupsReasonGroupFidNotFound{}
}

/*GetReasonsGroupsReasonGroupFidNotFound handles this case with default header values.

reasonGroupFid not found
*/
type GetReasonsGroupsReasonGroupFidNotFound struct {
}

func (o *GetReasonsGroupsReasonGroupFidNotFound) Error() string {
	return fmt.Sprintf("[GET /reasons/groups/{reasonGroupFid}][%d] getReasonsGroupsReasonGroupFidNotFound ", 404)
}

func (o *GetReasonsGroupsReasonGroupFidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
