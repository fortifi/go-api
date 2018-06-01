// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// InvoiceItem Generic Response
// swagger:model InvoiceItem
type InvoiceItem struct {
	Entity

	InvoiceItemAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *InvoiceItem) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 Entity
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Entity = aO0

	// AO1
	var aO1 InvoiceItemAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.InvoiceItemAllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m InvoiceItem) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.Entity)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.InvoiceItemAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this invoice item
func (m *InvoiceItem) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Entity
	if err := m.Entity.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with InvoiceItemAllOf1
	if err := m.InvoiceItemAllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *InvoiceItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InvoiceItem) UnmarshalBinary(b []byte) error {
	var res InvoiceItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
