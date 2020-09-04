// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PaymentMethodUpdatePayload Update payment method details
//
// swagger:model PaymentMethodUpdatePayload
type PaymentMethodUpdatePayload struct {

	// Name as appears on card/account
	AccountHolder string `json:"accountHolder,omitempty"`

	// Expiration Month
	ExpiryMonth int32 `json:"expiryMonth,omitempty"`

	// Expiration Year
	ExpiryYear int32 `json:"expiryYear,omitempty"`

	// Issue number
	Issue int32 `json:"issue,omitempty"`

	// Valid from month
	ValidMonth int32 `json:"validMonth,omitempty"`

	// Valid from year
	ValidYear int32 `json:"validYear,omitempty"`
}

// Validate validates this payment method update payload
func (m *PaymentMethodUpdatePayload) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PaymentMethodUpdatePayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentMethodUpdatePayload) UnmarshalBinary(b []byte) error {
	var res PaymentMethodUpdatePayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
