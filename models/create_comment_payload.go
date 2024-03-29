// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateCommentPayload Create a note/comment against a data-node
//
// swagger:model CreateCommentPayload
type CreateCommentPayload struct {

	// author fid
	AuthorFid string `json:"authorFid,omitempty"`

	// colour
	Colour string `json:"colour,omitempty"`

	// data node fid
	DataNodeFid string `json:"dataNodeFid,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// microtime
	Microtime string `json:"microtime,omitempty"`
}

// Validate validates this create comment payload
func (m *CreateCommentPayload) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create comment payload based on context it is used
func (m *CreateCommentPayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateCommentPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateCommentPayload) UnmarshalBinary(b []byte) error {
	var res CreateCommentPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
