// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ConfirmOrderWithNewCardPayload confirm order with new card payload
// swagger:model ConfirmOrderWithNewCardPayload
type ConfirmOrderWithNewCardPayload struct {

	// address city
	AddressCity CardAddressCity `json:"addressCity,omitempty"`

	// address country
	AddressCountry CardAddressCountry `json:"addressCountry,omitempty"`

	// address fid
	AddressFid CardAddressFid `json:"addressFid,omitempty"`

	// address line1
	AddressLine1 CardAddressLine1 `json:"addressLine1,omitempty"`

	// address line2
	AddressLine2 CardAddressLine2 `json:"addressLine2,omitempty"`

	// address line3
	AddressLine3 CardAddressLine3 `json:"addressLine3,omitempty"`

	// address postal
	AddressPostal CardAddressPostal `json:"addressPostal,omitempty"`

	// address state
	AddressState CardAddressState `json:"addressState,omitempty"`

	// card display name
	CardDisplayName DisplayName `json:"cardDisplayName,omitempty"`

	// card holder
	CardHolder CardCardHolder `json:"cardHolder,omitempty"`

	// card type
	CardType CardType `json:"cardType,omitempty"`

	// encrypted card number
	EncryptedCardNumber CardEncryptedNumber `json:"encryptedCardNumber,omitempty"`

	// expiry month
	ExpiryMonth CardExpiryMonth `json:"expiryMonth,omitempty"`

	// expiry year
	ExpiryYear CardExpiryYear `json:"expiryYear,omitempty"`

	// issue
	Issue CardIssue `json:"issue,omitempty"`

	// last4
	Last4 CardLast4 `json:"last4,omitempty"`

	// payment service fid
	PaymentServiceFid PaymentServiceFid `json:"paymentServiceFid,omitempty"`

	// start month
	StartMonth CardStartMonth `json:"startMonth,omitempty"`

	// start year
	StartYear CardStartYear `json:"startYear,omitempty"`
}

// Validate validates this confirm order with new card payload
func (m *ConfirmOrderWithNewCardPayload) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ConfirmOrderWithNewCardPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfirmOrderWithNewCardPayload) UnmarshalBinary(b []byte) error {
	var res ConfirmOrderWithNewCardPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
