// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AddInteractionMessagePayload add interaction message payload
//
// swagger:model AddInteractionMessagePayload
type AddInteractionMessagePayload struct {

	// action type
	ActionType string `json:"actionType,omitempty"`

	// language
	Language string `json:"language,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// IP of the message sender
	Source string `json:"source,omitempty"`
}

// Validate validates this add interaction message payload
func (m *AddInteractionMessagePayload) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add interaction message payload based on context it is used
func (m *AddInteractionMessagePayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AddInteractionMessagePayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AddInteractionMessagePayload) UnmarshalBinary(b []byte) error {
	var res AddInteractionMessagePayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
