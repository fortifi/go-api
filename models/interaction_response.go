// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// InteractionResponse interaction response
//
// swagger:model InteractionResponse
type InteractionResponse struct {

	// agent name
	AgentName string `json:"agentName,omitempty"`

	// feedback
	Feedback float64 `json:"feedback,omitempty"`

	// fid
	Fid string `json:"fid,omitempty"`

	// provider
	Provider string `json:"provider,omitempty"`

	// provider reference
	ProviderReference string `json:"providerReference,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// topic
	Topic string `json:"topic,omitempty"`
}

// Validate validates this interaction response
func (m *InteractionResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this interaction response based on context it is used
func (m *InteractionResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InteractionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InteractionResponse) UnmarshalBinary(b []byte) error {
	var res InteractionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
