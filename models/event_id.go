// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
)

// EventID If known, the Fortifi event ID for this visitors action
// swagger:model eventId
type EventID string

// Validate validates this event Id
func (m EventID) Validate(formats strfmt.Registry) error {
	return nil
}