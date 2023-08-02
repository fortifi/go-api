// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CustomerInteraction customer interaction
//
// swagger:model CustomerInteraction
type CustomerInteraction struct {

	// department fid
	DepartmentFid string `json:"departmentFid,omitempty"`

	// disconnect time
	DisconnectTime string `json:"disconnectTime,omitempty"`

	// fid
	Fid string `json:"fid,omitempty"`

	// initiated time
	InitiatedTime string `json:"initiatedTime,omitempty"`

	// queue fid
	QueueFid string `json:"queueFid,omitempty"`

	// topic
	Topic string `json:"topic,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this customer interaction
func (m *CustomerInteraction) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this customer interaction based on context it is used
func (m *CustomerInteraction) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CustomerInteraction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomerInteraction) UnmarshalBinary(b []byte) error {
	var res CustomerInteraction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}