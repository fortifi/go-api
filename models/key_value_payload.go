// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// KeyValuePayload Key Value Payload
//
// swagger:model keyValuePayload
type KeyValuePayload struct {

	// key
	Key string `json:"key,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this key value payload
func (m *KeyValuePayload) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *KeyValuePayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KeyValuePayload) UnmarshalBinary(b []byte) error {
	var res KeyValuePayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
