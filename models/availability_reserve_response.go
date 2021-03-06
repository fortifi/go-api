// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AvailabilityReserveResponse Reservation response
//
// swagger:model AvailabilityReserveResponse
type AvailabilityReserveResponse struct {

	// reserved
	Reserved bool `json:"reserved,omitempty"`
}

// Validate validates this availability reserve response
func (m *AvailabilityReserveResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this availability reserve response based on context it is used
func (m *AvailabilityReserveResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AvailabilityReserveResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AvailabilityReserveResponse) UnmarshalBinary(b []byte) error {
	var res AvailabilityReserveResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
