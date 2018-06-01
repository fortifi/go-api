// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ReasonGroups Reason Groups
// swagger:model ReasonGroups
type ReasonGroups struct {

	// reason groups
	ReasonGroups []*ReasonGroup `json:"reasonGroups"`
}

// Validate validates this reason groups
func (m *ReasonGroups) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReasonGroups(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReasonGroups) validateReasonGroups(formats strfmt.Registry) error {

	if swag.IsZero(m.ReasonGroups) { // not required
		return nil
	}

	for i := 0; i < len(m.ReasonGroups); i++ {
		if swag.IsZero(m.ReasonGroups[i]) { // not required
			continue
		}

		if m.ReasonGroups[i] != nil {
			if err := m.ReasonGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("reasonGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReasonGroups) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReasonGroups) UnmarshalBinary(b []byte) error {
	var res ReasonGroups
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
