// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// VerifyOrderPayload verify order payload
// swagger:model VerifyOrderPayload
type VerifyOrderPayload struct {
	ConfirmOrderPayload

	VerifyOrderPayloadAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *VerifyOrderPayload) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ConfirmOrderPayload
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ConfirmOrderPayload = aO0

	// AO1
	var aO1 VerifyOrderPayloadAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.VerifyOrderPayloadAllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m VerifyOrderPayload) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.ConfirmOrderPayload)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.VerifyOrderPayloadAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this verify order payload
func (m *VerifyOrderPayload) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ConfirmOrderPayload
	if err := m.ConfirmOrderPayload.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with VerifyOrderPayloadAllOf1
	if err := m.VerifyOrderPayloadAllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *VerifyOrderPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VerifyOrderPayload) UnmarshalBinary(b []byte) error {
	var res VerifyOrderPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
