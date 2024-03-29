// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// InteractionActionPayload interaction action payload
//
// swagger:model InteractionActionPayload
type InteractionActionPayload struct {

	// data
	Data interface{} `json:"data,omitempty"`

	// provider
	Provider string `json:"provider,omitempty"`

	// provider reference
	ProviderReference string `json:"providerReference,omitempty"`

	// timestamp
	Timestamp string `json:"timestamp,omitempty"`
}

// Validate validates this interaction action payload
func (m *InteractionActionPayload) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this interaction action payload based on context it is used
func (m *InteractionActionPayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InteractionActionPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InteractionActionPayload) UnmarshalBinary(b []byte) error {
	var res InteractionActionPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
