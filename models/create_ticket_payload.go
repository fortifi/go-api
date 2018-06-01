// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// CreateTicketPayload create ticket payload
// swagger:model CreateTicketPayload
type CreateTicketPayload struct {

	// brand fid
	BrandFid string `json:"brandFid,omitempty"`

	// customer email
	CustomerEmail string `json:"customerEmail,omitempty"`

	// customer fid
	CustomerFid string `json:"customerFid,omitempty"`

	// customer name
	CustomerName string `json:"customerName,omitempty"`

	// department email
	DepartmentEmail string `json:"departmentEmail,omitempty"`

	// department fid
	DepartmentFid string `json:"departmentFid,omitempty"`

	// department name
	DepartmentName string `json:"departmentName,omitempty"`

	// html body
	HTMLBody string `json:"htmlBody,omitempty"`

	// queue fid
	QueueFid string `json:"queueFid,omitempty"`

	// subject
	Subject string `json:"subject,omitempty"`

	// text body
	TextBody string `json:"textBody,omitempty"`

	// ticket type
	TicketType string `json:"ticketType,omitempty"`
}

// Validate validates this create ticket payload
func (m *CreateTicketPayload) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateTicketPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateTicketPayload) UnmarshalBinary(b []byte) error {
	var res CreateTicketPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
