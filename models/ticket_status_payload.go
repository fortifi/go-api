// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TicketStatusPayload Set the ticket status
// swagger:model TicketStatusPayload
type TicketStatusPayload struct {

	// ticket status
	TicketStatus TicketStatus `json:"ticketStatus,omitempty"`
}

// Validate validates this ticket status payload
func (m *TicketStatusPayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTicketStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TicketStatusPayload) validateTicketStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.TicketStatus) { // not required
		return nil
	}

	if err := m.TicketStatus.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ticketStatus")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TicketStatusPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TicketStatusPayload) UnmarshalBinary(b []byte) error {
	var res TicketStatusPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
