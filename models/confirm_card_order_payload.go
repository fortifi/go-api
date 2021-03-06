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

// ConfirmCardOrderPayload confirm card order payload
//
// swagger:model ConfirmCardOrderPayload
type ConfirmCardOrderPayload struct {
	ConfirmOrderPayload

	// cvv
	Cvv string `json:"cvv,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ConfirmCardOrderPayload) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ConfirmOrderPayload
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ConfirmOrderPayload = aO0

	// AO1
	var dataAO1 struct {
		Cvv string `json:"cvv,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Cvv = dataAO1.Cvv

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ConfirmCardOrderPayload) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.ConfirmOrderPayload)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Cvv string `json:"cvv,omitempty"`
	}

	dataAO1.Cvv = m.Cvv

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this confirm card order payload
func (m *ConfirmCardOrderPayload) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ConfirmOrderPayload
	if err := m.ConfirmOrderPayload.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this confirm card order payload based on the context it is used
func (m *ConfirmCardOrderPayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ConfirmOrderPayload
	if err := m.ConfirmOrderPayload.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ConfirmCardOrderPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfirmCardOrderPayload) UnmarshalBinary(b []byte) error {
	var res ConfirmCardOrderPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
