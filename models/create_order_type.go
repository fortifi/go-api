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

// CreateOrderType Order Type (Must be initial to Auth Cards)
// swagger:model CreateOrderType
type CreateOrderType string

const (

	// CreateOrderTypeInitial captures enum value "initial"
	CreateOrderTypeInitial CreateOrderType = "initial"

	// CreateOrderTypePurchase captures enum value "purchase"
	CreateOrderTypePurchase CreateOrderType = "purchase"
)

// for schema
var createOrderTypeEnum []interface{}

func init() {
	var res []CreateOrderType
	if err := json.Unmarshal([]byte(`["initial","purchase"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createOrderTypeEnum = append(createOrderTypeEnum, v)
	}
}

func (m CreateOrderType) validateCreateOrderTypeEnum(path, location string, value CreateOrderType) error {
	if err := validate.Enum(path, location, value, createOrderTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this create order type
func (m CreateOrderType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateCreateOrderTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
