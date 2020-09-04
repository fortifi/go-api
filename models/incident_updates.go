// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IncidentUpdates Incident updates
//
// swagger:model IncidentUpdates
type IncidentUpdates struct {

	// incident updates
	IncidentUpdates []*IncidentUpdate `json:"incidentUpdates"`
}

// Validate validates this incident updates
func (m *IncidentUpdates) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIncidentUpdates(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IncidentUpdates) validateIncidentUpdates(formats strfmt.Registry) error {

	if swag.IsZero(m.IncidentUpdates) { // not required
		return nil
	}

	for i := 0; i < len(m.IncidentUpdates); i++ {
		if swag.IsZero(m.IncidentUpdates[i]) { // not required
			continue
		}

		if m.IncidentUpdates[i] != nil {
			if err := m.IncidentUpdates[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("incidentUpdates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IncidentUpdates) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IncidentUpdates) UnmarshalBinary(b []byte) error {
	var res IncidentUpdates
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
