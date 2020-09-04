// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PropertyCounter A counter type property
//
// swagger:model PropertyCounter
type PropertyCounter struct {

	// Property name
	Key string `json:"key,omitempty"`

	// Counter value
	Value float64 `json:"value,omitempty"`
}

// Validate validates this property counter
func (m *PropertyCounter) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PropertyCounter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PropertyCounter) UnmarshalBinary(b []byte) error {
	var res PropertyCounter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
