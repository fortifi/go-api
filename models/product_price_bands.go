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

// ProductPriceBands Product Prices
//
// swagger:model ProductPriceBands
type ProductPriceBands struct {

	// pricebands
	Pricebands []*ProductPriceBand `json:"pricebands"`
}

// Validate validates this product price bands
func (m *ProductPriceBands) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePricebands(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProductPriceBands) validatePricebands(formats strfmt.Registry) error {
	if swag.IsZero(m.Pricebands) { // not required
		return nil
	}

	for i := 0; i < len(m.Pricebands); i++ {
		if swag.IsZero(m.Pricebands[i]) { // not required
			continue
		}

		if m.Pricebands[i] != nil {
			if err := m.Pricebands[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pricebands" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("pricebands" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this product price bands based on the context it is used
func (m *ProductPriceBands) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePricebands(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProductPriceBands) contextValidatePricebands(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Pricebands); i++ {

		if m.Pricebands[i] != nil {

			if swag.IsZero(m.Pricebands[i]) { // not required
				return nil
			}

			if err := m.Pricebands[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pricebands" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("pricebands" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProductPriceBands) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProductPriceBands) UnmarshalBinary(b []byte) error {
	var res ProductPriceBands
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
