// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FindTransactionPayload find transaction payload
//
// swagger:model FindTransactionPayload
type FindTransactionPayload struct {

	// amount
	Amount string `json:"amount,omitempty"`

	// company fid
	CompanyFid string `json:"companyFid,omitempty"`

	// country
	Country string `json:"country,omitempty"`

	// currency
	Currency string `json:"currency,omitempty"`

	// date
	Date string `json:"date,omitempty"`

	// email address
	EmailAddress string `json:"emailAddress,omitempty"`

	// first6
	First6 string `json:"first6,omitempty"`

	// first name
	FirstName string `json:"firstName,omitempty"`

	// last4
	Last4 string `json:"last4,omitempty"`

	// last name
	LastName string `json:"lastName,omitempty"`

	// name on card
	NameOnCard string `json:"nameOnCard,omitempty"`

	// payment method
	PaymentMethod string `json:"paymentMethod,omitempty"`

	// paypal email
	PaypalEmail string `json:"paypalEmail,omitempty"`

	// post code
	PostCode string `json:"postCode,omitempty"`
}

// Validate validates this find transaction payload
func (m *FindTransactionPayload) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this find transaction payload based on context it is used
func (m *FindTransactionPayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FindTransactionPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FindTransactionPayload) UnmarshalBinary(b []byte) error {
	var res FindTransactionPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
