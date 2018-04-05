// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ConfigItem config item
// swagger:model ConfigItem
type ConfigItem struct {
	Entity

	ConfigItemAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ConfigItem) UnmarshalJSON(raw []byte) error {

	var aO0 Entity
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Entity = aO0

	var aO1 ConfigItemAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.ConfigItemAllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ConfigItem) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	aO0, err := swag.WriteJSON(m.Entity)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.ConfigItemAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this config item
func (m *ConfigItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.Entity.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.ConfigItemAllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ConfigItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigItem) UnmarshalBinary(b []byte) error {
	var res ConfigItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
