// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Product Generic Response
//
// swagger:model Product
type Product struct {
	Entity

	// allow quantity
	AllowQuantity int32 `json:"allowQuantity,omitempty"`

	// default price
	DefaultPrice *ProductPrice `json:"defaultPrice,omitempty"`

	// default price fid
	DefaultPriceFid string `json:"defaultPriceFid,omitempty"`

	// manager type
	ManagerType string `json:"managerType,omitempty"`

	// max quantity
	MaxQuantity int64 `json:"maxQuantity,omitempty"`

	// parent fid
	ParentFid string `json:"parentFid,omitempty"`

	// product type
	ProductType int32 `json:"productType,omitempty"`

	// statement description
	StatementDescription string `json:"statementDescription,omitempty"`

	// tax group fid
	TaxGroupFid string `json:"taxGroupFid,omitempty"`

	// taxable
	Taxable int64 `json:"taxable,omitempty"`

	// trial days
	TrialDays int32 `json:"trialDays,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *Product) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 Entity
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Entity = aO0

	// AO1
	var dataAO1 struct {
		AllowQuantity int32 `json:"allowQuantity,omitempty"`

		DefaultPrice *ProductPrice `json:"defaultPrice,omitempty"`

		DefaultPriceFid string `json:"defaultPriceFid,omitempty"`

		ManagerType string `json:"managerType,omitempty"`

		MaxQuantity int64 `json:"maxQuantity,omitempty"`

		ParentFid string `json:"parentFid,omitempty"`

		ProductType int32 `json:"productType,omitempty"`

		StatementDescription string `json:"statementDescription,omitempty"`

		TaxGroupFid string `json:"taxGroupFid,omitempty"`

		Taxable int64 `json:"taxable,omitempty"`

		TrialDays int32 `json:"trialDays,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.AllowQuantity = dataAO1.AllowQuantity

	m.DefaultPrice = dataAO1.DefaultPrice

	m.DefaultPriceFid = dataAO1.DefaultPriceFid

	m.ManagerType = dataAO1.ManagerType

	m.MaxQuantity = dataAO1.MaxQuantity

	m.ParentFid = dataAO1.ParentFid

	m.ProductType = dataAO1.ProductType

	m.StatementDescription = dataAO1.StatementDescription

	m.TaxGroupFid = dataAO1.TaxGroupFid

	m.Taxable = dataAO1.Taxable

	m.TrialDays = dataAO1.TrialDays

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m Product) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.Entity)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		AllowQuantity int32 `json:"allowQuantity,omitempty"`

		DefaultPrice *ProductPrice `json:"defaultPrice,omitempty"`

		DefaultPriceFid string `json:"defaultPriceFid,omitempty"`

		ManagerType string `json:"managerType,omitempty"`

		MaxQuantity int64 `json:"maxQuantity,omitempty"`

		ParentFid string `json:"parentFid,omitempty"`

		ProductType int32 `json:"productType,omitempty"`

		StatementDescription string `json:"statementDescription,omitempty"`

		TaxGroupFid string `json:"taxGroupFid,omitempty"`

		Taxable int64 `json:"taxable,omitempty"`

		TrialDays int32 `json:"trialDays,omitempty"`
	}

	dataAO1.AllowQuantity = m.AllowQuantity

	dataAO1.DefaultPrice = m.DefaultPrice

	dataAO1.DefaultPriceFid = m.DefaultPriceFid

	dataAO1.ManagerType = m.ManagerType

	dataAO1.MaxQuantity = m.MaxQuantity

	dataAO1.ParentFid = m.ParentFid

	dataAO1.ProductType = m.ProductType

	dataAO1.StatementDescription = m.StatementDescription

	dataAO1.TaxGroupFid = m.TaxGroupFid

	dataAO1.Taxable = m.Taxable

	dataAO1.TrialDays = m.TrialDays

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this product
func (m *Product) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Entity
	if err := m.Entity.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefaultPrice(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Product) validateDefaultPrice(formats strfmt.Registry) error {

	if swag.IsZero(m.DefaultPrice) { // not required
		return nil
	}

	if m.DefaultPrice != nil {
		if err := m.DefaultPrice.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("defaultPrice")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Product) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Product) UnmarshalBinary(b []byte) error {
	var res Product
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
