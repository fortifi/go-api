// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// CycleTermType Term Type
// swagger:model cycleTermType
type CycleTermType string

const (
	// CycleTermTypeOnetime captures enum value "onetime"
	CycleTermTypeOnetime CycleTermType = "onetime"
	// CycleTermTypeLifetime captures enum value "lifetime"
	CycleTermTypeLifetime CycleTermType = "lifetime"
	// CycleTermTypeDay captures enum value "day"
	CycleTermTypeDay CycleTermType = "day"
	// CycleTermTypeWeek captures enum value "week"
	CycleTermTypeWeek CycleTermType = "week"
	// CycleTermTypeMonth captures enum value "month"
	CycleTermTypeMonth CycleTermType = "month"
	// CycleTermTypeYear captures enum value "year"
	CycleTermTypeYear CycleTermType = "year"
)

// for schema
var cycleTermTypeEnum []interface{}

func init() {
	var res []CycleTermType
	if err := json.Unmarshal([]byte(`["onetime","lifetime","day","week","month","year"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cycleTermTypeEnum = append(cycleTermTypeEnum, v)
	}
}

func (m CycleTermType) validateCycleTermTypeEnum(path, location string, value CycleTermType) error {
	if err := validate.Enum(path, location, value, cycleTermTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this cycle term type
func (m CycleTermType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateCycleTermTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
