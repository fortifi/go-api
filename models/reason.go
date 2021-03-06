// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Reason Reason
//
// swagger:model Reason
type Reason struct {
	Entity

	// FID for the reason group
	GroupFid string `json:"groupFid,omitempty"`

	// used count
	UsedCount int64 `json:"usedCount,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *Reason) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 Entity
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Entity = aO0

	// AO1
	var dataAO1 struct {
		GroupFid string `json:"groupFid,omitempty"`

		UsedCount int64 `json:"usedCount,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.GroupFid = dataAO1.GroupFid

	m.UsedCount = dataAO1.UsedCount

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m Reason) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.Entity)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		GroupFid string `json:"groupFid,omitempty"`

		UsedCount int64 `json:"usedCount,omitempty"`
	}

	dataAO1.GroupFid = m.GroupFid

	dataAO1.UsedCount = m.UsedCount

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this reason
func (m *Reason) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Entity
	if err := m.Entity.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this reason based on the context it is used
func (m *Reason) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Entity
	if err := m.Entity.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Reason) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Reason) UnmarshalBinary(b []byte) error {
	var res Reason
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
