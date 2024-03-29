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

// Orders List of orders
//
// swagger:model Orders
type Orders struct {

	// orders
	Orders []*Order `json:"orders"`
}

// Validate validates this orders
func (m *Orders) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrders(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Orders) validateOrders(formats strfmt.Registry) error {
	if swag.IsZero(m.Orders) { // not required
		return nil
	}

	for i := 0; i < len(m.Orders); i++ {
		if swag.IsZero(m.Orders[i]) { // not required
			continue
		}

		if m.Orders[i] != nil {
			if err := m.Orders[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("orders" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("orders" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this orders based on the context it is used
func (m *Orders) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOrders(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Orders) contextValidateOrders(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Orders); i++ {

		if m.Orders[i] != nil {

			if swag.IsZero(m.Orders[i]) { // not required
				return nil
			}

			if err := m.Orders[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("orders" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("orders" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Orders) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Orders) UnmarshalBinary(b []byte) error {
	var res Orders
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
