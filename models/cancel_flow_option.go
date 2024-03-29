// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CancelFlowOption cancel flow option
//
// swagger:model CancelFlowOption
type CancelFlowOption struct {

	// display name
	DisplayName string `json:"displayName,omitempty"`

	// key
	Key string `json:"key,omitempty"`

	// type
	// Enum: [flag value]
	Type string `json:"type,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this cancel flow option
func (m *CancelFlowOption) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var cancelFlowOptionTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["flag","value"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cancelFlowOptionTypeTypePropEnum = append(cancelFlowOptionTypeTypePropEnum, v)
	}
}

const (

	// CancelFlowOptionTypeFlag captures enum value "flag"
	CancelFlowOptionTypeFlag string = "flag"

	// CancelFlowOptionTypeValue captures enum value "value"
	CancelFlowOptionTypeValue string = "value"
)

// prop value enum
func (m *CancelFlowOption) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, cancelFlowOptionTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CancelFlowOption) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this cancel flow option based on context it is used
func (m *CancelFlowOption) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CancelFlowOption) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CancelFlowOption) UnmarshalBinary(b []byte) error {
	var res CancelFlowOption
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
