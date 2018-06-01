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

// PaymentMode Payment Mode
// swagger:model PaymentMode
type PaymentMode string

const (

	// PaymentModeAutomatic captures enum value "automatic"
	PaymentModeAutomatic PaymentMode = "automatic"

	// PaymentModeRequest captures enum value "request"
	PaymentModeRequest PaymentMode = "request"
)

// for schema
var paymentModeEnum []interface{}

func init() {
	var res []PaymentMode
	if err := json.Unmarshal([]byte(`["automatic","request"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		paymentModeEnum = append(paymentModeEnum, v)
	}
}

func (m PaymentMode) validatePaymentModeEnum(path, location string, value PaymentMode) error {
	if err := validate.Enum(path, location, value, paymentModeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this payment mode
func (m PaymentMode) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validatePaymentModeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
