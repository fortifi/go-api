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

// Payments Payments
//
// swagger:model Payments
type Payments struct {

	// payments
	Payments []*Payment `json:"payments"`
}

// Validate validates this payments
func (m *Payments) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePayments(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Payments) validatePayments(formats strfmt.Registry) error {
	if swag.IsZero(m.Payments) { // not required
		return nil
	}

	for i := 0; i < len(m.Payments); i++ {
		if swag.IsZero(m.Payments[i]) { // not required
			continue
		}

		if m.Payments[i] != nil {
			if err := m.Payments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("payments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("payments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this payments based on the context it is used
func (m *Payments) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePayments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Payments) contextValidatePayments(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Payments); i++ {

		if m.Payments[i] != nil {

			if swag.IsZero(m.Payments[i]) { // not required
				return nil
			}

			if err := m.Payments[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("payments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("payments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Payments) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Payments) UnmarshalBinary(b []byte) error {
	var res Payments
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
