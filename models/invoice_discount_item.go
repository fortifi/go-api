// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// InvoiceDiscountItem Generic Response
//
// swagger:model InvoiceDiscountItem
type InvoiceDiscountItem struct {
	Entity

	// amount
	Amount float32 `json:"amount,omitempty"`

	// product fid
	ProductFid string `json:"productFid,omitempty"`

	// purchase fid
	PurchaseFid string `json:"purchaseFid,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *InvoiceDiscountItem) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 Entity
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Entity = aO0

	// AO1
	var dataAO1 struct {
		Amount float32 `json:"amount,omitempty"`

		ProductFid string `json:"productFid,omitempty"`

		PurchaseFid string `json:"purchaseFid,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Amount = dataAO1.Amount

	m.ProductFid = dataAO1.ProductFid

	m.PurchaseFid = dataAO1.PurchaseFid

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m InvoiceDiscountItem) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.Entity)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Amount float32 `json:"amount,omitempty"`

		ProductFid string `json:"productFid,omitempty"`

		PurchaseFid string `json:"purchaseFid,omitempty"`
	}

	dataAO1.Amount = m.Amount

	dataAO1.ProductFid = m.ProductFid

	dataAO1.PurchaseFid = m.PurchaseFid

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this invoice discount item
func (m *InvoiceDiscountItem) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Entity
	if err := m.Entity.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this invoice discount item based on the context it is used
func (m *InvoiceDiscountItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Entity
	if err := m.Entity.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *InvoiceDiscountItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InvoiceDiscountItem) UnmarshalBinary(b []byte) error {
	var res InvoiceDiscountItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
