// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OrderProductDisplayNamePayload order product display name payload
//
// swagger:model OrderProductDisplayNamePayload
type OrderProductDisplayNamePayload struct {

	// display name
	DisplayName string `json:"displayName,omitempty"`

	// Price FID to modify subscription with
	// Required: true
	PriceFid *string `json:"priceFid"`
}

// Validate validates this order product display name payload
func (m *OrderProductDisplayNamePayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePriceFid(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrderProductDisplayNamePayload) validatePriceFid(formats strfmt.Registry) error {

	if err := validate.Required("priceFid", "body", m.PriceFid); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this order product display name payload based on context it is used
func (m *OrderProductDisplayNamePayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OrderProductDisplayNamePayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrderProductDisplayNamePayload) UnmarshalBinary(b []byte) error {
	var res OrderProductDisplayNamePayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
