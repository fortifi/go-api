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

// CustomerLifecycle Customer Lifecycle Stage
// swagger:model customerLifecycle
type CustomerLifecycle string

const (
	// CustomerLifecycleProspect captures enum value "prospect"
	CustomerLifecycleProspect CustomerLifecycle = "prospect"
	// CustomerLifecycleCustomer captures enum value "customer"
	CustomerLifecycleCustomer CustomerLifecycle = "customer"
	// CustomerLifecycleClosed captures enum value "closed"
	CustomerLifecycleClosed CustomerLifecycle = "closed"
)

// for schema
var customerLifecycleEnum []interface{}

func init() {
	var res []CustomerLifecycle
	if err := json.Unmarshal([]byte(`["prospect","customer","closed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		customerLifecycleEnum = append(customerLifecycleEnum, v)
	}
}

func (m CustomerLifecycle) validateCustomerLifecycleEnum(path, location string, value CustomerLifecycle) error {
	if err := validate.Enum(path, location, value, customerLifecycleEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this customer lifecycle
func (m CustomerLifecycle) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateCustomerLifecycleEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}