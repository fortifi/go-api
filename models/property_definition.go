// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PropertyDefinition Definition of a property
// swagger:model PropertyDefinition
type PropertyDefinition struct {
	Entity

	// group fid
	GroupFid string `json:"groupFid,omitempty"`

	// key
	Key string `json:"key,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// validation config
	ValidationConfig string `json:"validationConfig,omitempty"`

	// validation type
	// Enum: [regex prerequisite]
	ValidationType string `json:"validationType,omitempty"`

	// value type
	ValueType string `json:"valueType,omitempty"`

	// values
	Values []string `json:"values"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *PropertyDefinition) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 Entity
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Entity = aO0

	// AO1
	var dataAO1 struct {
		GroupFid string `json:"groupFid,omitempty"`

		Key string `json:"key,omitempty"`

		Type string `json:"type,omitempty"`

		ValidationConfig string `json:"validationConfig,omitempty"`

		ValidationType string `json:"validationType,omitempty"`

		ValueType string `json:"valueType,omitempty"`

		Values []string `json:"values"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.GroupFid = dataAO1.GroupFid

	m.Key = dataAO1.Key

	m.Type = dataAO1.Type

	m.ValidationConfig = dataAO1.ValidationConfig

	m.ValidationType = dataAO1.ValidationType

	m.ValueType = dataAO1.ValueType

	m.Values = dataAO1.Values

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m PropertyDefinition) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.Entity)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		GroupFid string `json:"groupFid,omitempty"`

		Key string `json:"key,omitempty"`

		Type string `json:"type,omitempty"`

		ValidationConfig string `json:"validationConfig,omitempty"`

		ValidationType string `json:"validationType,omitempty"`

		ValueType string `json:"valueType,omitempty"`

		Values []string `json:"values"`
	}

	dataAO1.GroupFid = m.GroupFid

	dataAO1.Key = m.Key

	dataAO1.Type = m.Type

	dataAO1.ValidationConfig = m.ValidationConfig

	dataAO1.ValidationType = m.ValidationType

	dataAO1.ValueType = m.ValueType

	dataAO1.Values = m.Values

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this property definition
func (m *PropertyDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Entity
	if err := m.Entity.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidationType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var propertyDefinitionTypeValidationTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["regex","prerequisite"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		propertyDefinitionTypeValidationTypePropEnum = append(propertyDefinitionTypeValidationTypePropEnum, v)
	}
}

// property enum
func (m *PropertyDefinition) validateValidationTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, propertyDefinitionTypeValidationTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PropertyDefinition) validateValidationType(formats strfmt.Registry) error {

	if swag.IsZero(m.ValidationType) { // not required
		return nil
	}

	// value enum
	if err := m.validateValidationTypeEnum("validationType", "body", m.ValidationType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PropertyDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PropertyDefinition) UnmarshalBinary(b []byte) error {
	var res PropertyDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
