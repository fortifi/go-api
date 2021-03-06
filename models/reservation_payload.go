// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ReservationPayload Data for a reservation
//
// swagger:model ReservationPayload
type ReservationPayload struct {

	// Time in seconds to hold this reservation for
	Duration int32 `json:"duration,omitempty"`

	// Related Fids
	RelatedFids []string `json:"relatedFids"`
}

// Validate validates this reservation payload
func (m *ReservationPayload) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this reservation payload based on context it is used
func (m *ReservationPayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ReservationPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReservationPayload) UnmarshalBinary(b []byte) error {
	var res ReservationPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
