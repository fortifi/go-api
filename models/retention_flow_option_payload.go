// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RetentionFlowOptionPayload retention flow option payload
//
// swagger:model RetentionFlowOptionPayload
type RetentionFlowOptionPayload struct {

	// key
	Key string `json:"key,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this retention flow option payload
func (m *RetentionFlowOptionPayload) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this retention flow option payload based on context it is used
func (m *RetentionFlowOptionPayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RetentionFlowOptionPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RetentionFlowOptionPayload) UnmarshalBinary(b []byte) error {
	var res RetentionFlowOptionPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
