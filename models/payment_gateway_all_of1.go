// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// PaymentGatewayAllOf1 payment gateway all of1
// swagger:model paymentGatewayAllOf1
type PaymentGatewayAllOf1 struct {

	// brands
	Brands []string `json:"brands"`

	// card types
	CardTypes []string `json:"cardTypes"`

	// currencies
	Currencies []string `json:"currencies"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// external Id
	ExternalID string `json:"externalId,omitempty"`

	// provider
	Provider string `json:"provider,omitempty"`
}

// Validate validates this payment gateway all of1
func (m *PaymentGatewayAllOf1) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PaymentGatewayAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentGatewayAllOf1) UnmarshalBinary(b []byte) error {
	var res PaymentGatewayAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
