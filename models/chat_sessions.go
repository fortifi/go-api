// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ChatSessions List of chat sessions
//
// swagger:model ChatSessions
type ChatSessions struct {
	Pagination

	// chat sessions
	ChatSessions []*ChatSession `json:"chatSessions"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ChatSessions) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 Pagination
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Pagination = aO0

	// AO1
	var dataAO1 struct {
		ChatSessions []*ChatSession `json:"chatSessions"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.ChatSessions = dataAO1.ChatSessions

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ChatSessions) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.Pagination)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		ChatSessions []*ChatSession `json:"chatSessions"`
	}

	dataAO1.ChatSessions = m.ChatSessions

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this chat sessions
func (m *ChatSessions) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Pagination
	if err := m.Pagination.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChatSessions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChatSessions) validateChatSessions(formats strfmt.Registry) error {

	if swag.IsZero(m.ChatSessions) { // not required
		return nil
	}

	for i := 0; i < len(m.ChatSessions); i++ {
		if swag.IsZero(m.ChatSessions[i]) { // not required
			continue
		}

		if m.ChatSessions[i] != nil {
			if err := m.ChatSessions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("chatSessions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("chatSessions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this chat sessions based on the context it is used
func (m *ChatSessions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Pagination
	if err := m.Pagination.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateChatSessions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChatSessions) contextValidateChatSessions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ChatSessions); i++ {

		if m.ChatSessions[i] != nil {

			if swag.IsZero(m.ChatSessions[i]) { // not required
				return nil
			}

			if err := m.ChatSessions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("chatSessions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("chatSessions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ChatSessions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChatSessions) UnmarshalBinary(b []byte) error {
	var res ChatSessions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
