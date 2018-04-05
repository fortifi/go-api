// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OrderProductQuantityPayload order product quantity payload
// swagger:model OrderProductQuantityPayload
type OrderProductQuantityPayload struct {

	// price fid
	// Required: true
	PriceFid PriceFid `json:"priceFid"`

	// quantity
	Quantity *float64 `json:"quantity,omitempty"`
}

// Validate validates this order product quantity payload
func (m *OrderProductQuantityPayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePriceFid(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrderProductQuantityPayload) validatePriceFid(formats strfmt.Registry) error {

	if err := m.PriceFid.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("priceFid")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OrderProductQuantityPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrderProductQuantityPayload) UnmarshalBinary(b []byte) error {
	var res OrderProductQuantityPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}