// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PaymentCards Generic Response
// swagger:model PaymentCards
type PaymentCards struct {

	// cards
	Cards []*PaymentCard `json:"cards"`
}

// Validate validates this payment cards
func (m *PaymentCards) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCards(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentCards) validateCards(formats strfmt.Registry) error {

	if swag.IsZero(m.Cards) { // not required
		return nil
	}

	for i := 0; i < len(m.Cards); i++ {
		if swag.IsZero(m.Cards[i]) { // not required
			continue
		}

		if m.Cards[i] != nil {
			if err := m.Cards[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cards" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentCards) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentCards) UnmarshalBinary(b []byte) error {
	var res PaymentCards
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
