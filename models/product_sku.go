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

// ProductSku Product SKU
//
// swagger:model ProductSku
type ProductSku struct {
	Entity

	// price band fid
	PriceBandFid string `json:"priceBandFid,omitempty"`

	// product fid
	ProductFid string `json:"productFid,omitempty"`

	// resource pool fid
	ResourcePoolFid string `json:"resourcePoolFid,omitempty"`

	// sku
	Sku string `json:"sku,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ProductSku) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 Entity
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Entity = aO0

	// AO1
	var dataAO1 struct {
		PriceBandFid string `json:"priceBandFid,omitempty"`

		ProductFid string `json:"productFid,omitempty"`

		ResourcePoolFid string `json:"resourcePoolFid,omitempty"`

		Sku string `json:"sku,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.PriceBandFid = dataAO1.PriceBandFid

	m.ProductFid = dataAO1.ProductFid

	m.ResourcePoolFid = dataAO1.ResourcePoolFid

	m.Sku = dataAO1.Sku

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ProductSku) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.Entity)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		PriceBandFid string `json:"priceBandFid,omitempty"`

		ProductFid string `json:"productFid,omitempty"`

		ResourcePoolFid string `json:"resourcePoolFid,omitempty"`

		Sku string `json:"sku,omitempty"`
	}

	dataAO1.PriceBandFid = m.PriceBandFid

	dataAO1.ProductFid = m.ProductFid

	dataAO1.ResourcePoolFid = m.ResourcePoolFid

	dataAO1.Sku = m.Sku

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this product sku
func (m *ProductSku) Validate(formats strfmt.Registry) error {
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

// ContextValidate validate this product sku based on the context it is used
func (m *ProductSku) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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
func (m *ProductSku) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProductSku) UnmarshalBinary(b []byte) error {
	var res ProductSku
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
