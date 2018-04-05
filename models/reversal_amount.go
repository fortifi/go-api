// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
)

// ReversalAmount Amount of revene refunded to the client from the original transaction / chargeback amount
// swagger:model reversalAmount
type ReversalAmount float32

// Validate validates this reversal amount
func (m ReversalAmount) Validate(formats strfmt.Registry) error {
	return nil
}
