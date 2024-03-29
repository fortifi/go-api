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

// OrderProduct Generic Response
//
// swagger:model OrderProduct
type OrderProduct struct {
	Entity

	// currency
	Currency string `json:"currency,omitempty"`

	// Interval in ISO 8601 standard
	Cycle string `json:"cycle,omitempty"`

	// cycle exact
	CycleExact string `json:"cycleExact,omitempty"`

	// cycle term
	CycleTerm int32 `json:"cycleTerm,omitempty"`

	// cycle type
	CycleType CycleTermType `json:"cycleType,omitempty"`

	// discount amount
	DiscountAmount float32 `json:"discountAmount,omitempty"`

	// identity
	Identity string `json:"identity,omitempty"`

	// Custom initial term end date (Time in ISO 8601 standard e.g 2015-12-05T13:11:59Z)
	InitialTermEndDate string `json:"initialTermEndDate,omitempty"`

	// Custom initial term start date (Time in ISO 8601 standard e.g 2015-12-05T13:11:59Z)
	InitialTermStartDate string `json:"initialTermStartDate,omitempty"`

	// offer fid
	OfferFid string `json:"offerFid,omitempty"`

	// parent fid
	ParentFid string `json:"parentFid,omitempty"`

	// price
	Price float32 `json:"price,omitempty"`

	// price fid
	PriceFid string `json:"priceFid,omitempty"`

	// product fid
	ProductFid string `json:"productFid,omitempty"`

	// purchase fid
	PurchaseFid string `json:"purchaseFid,omitempty"`

	// quantity
	Quantity int64 `json:"quantity,omitempty"`

	// Time in ISO 8601 standard with optional fractions of a second e.g 2015-12-05T13:11:59.888Z
	// Format: date-time
	RenewalDate strfmt.DateTime `json:"renewalDate,omitempty"`

	// setup discount amount
	SetupDiscountAmount float32 `json:"setupDiscountAmount,omitempty"`

	// setup fee
	SetupFee float32 `json:"setupFee,omitempty"`

	// sku fid
	SkuFid string `json:"skuFid,omitempty"`

	// tax amount
	TaxAmount float32 `json:"taxAmount,omitempty"`

	// total amount
	TotalAmount float32 `json:"totalAmount,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *OrderProduct) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 Entity
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Entity = aO0

	// AO1
	var dataAO1 struct {
		Currency string `json:"currency,omitempty"`

		Cycle string `json:"cycle,omitempty"`

		CycleExact string `json:"cycleExact,omitempty"`

		CycleTerm int32 `json:"cycleTerm,omitempty"`

		CycleType CycleTermType `json:"cycleType,omitempty"`

		DiscountAmount float32 `json:"discountAmount,omitempty"`

		Identity string `json:"identity,omitempty"`

		InitialTermEndDate string `json:"initialTermEndDate,omitempty"`

		InitialTermStartDate string `json:"initialTermStartDate,omitempty"`

		OfferFid string `json:"offerFid,omitempty"`

		ParentFid string `json:"parentFid,omitempty"`

		Price float32 `json:"price,omitempty"`

		PriceFid string `json:"priceFid,omitempty"`

		ProductFid string `json:"productFid,omitempty"`

		PurchaseFid string `json:"purchaseFid,omitempty"`

		Quantity int64 `json:"quantity,omitempty"`

		RenewalDate strfmt.DateTime `json:"renewalDate,omitempty"`

		SetupDiscountAmount float32 `json:"setupDiscountAmount,omitempty"`

		SetupFee float32 `json:"setupFee,omitempty"`

		SkuFid string `json:"skuFid,omitempty"`

		TaxAmount float32 `json:"taxAmount,omitempty"`

		TotalAmount float32 `json:"totalAmount,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Currency = dataAO1.Currency

	m.Cycle = dataAO1.Cycle

	m.CycleExact = dataAO1.CycleExact

	m.CycleTerm = dataAO1.CycleTerm

	m.CycleType = dataAO1.CycleType

	m.DiscountAmount = dataAO1.DiscountAmount

	m.Identity = dataAO1.Identity

	m.InitialTermEndDate = dataAO1.InitialTermEndDate

	m.InitialTermStartDate = dataAO1.InitialTermStartDate

	m.OfferFid = dataAO1.OfferFid

	m.ParentFid = dataAO1.ParentFid

	m.Price = dataAO1.Price

	m.PriceFid = dataAO1.PriceFid

	m.ProductFid = dataAO1.ProductFid

	m.PurchaseFid = dataAO1.PurchaseFid

	m.Quantity = dataAO1.Quantity

	m.RenewalDate = dataAO1.RenewalDate

	m.SetupDiscountAmount = dataAO1.SetupDiscountAmount

	m.SetupFee = dataAO1.SetupFee

	m.SkuFid = dataAO1.SkuFid

	m.TaxAmount = dataAO1.TaxAmount

	m.TotalAmount = dataAO1.TotalAmount

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m OrderProduct) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.Entity)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Currency string `json:"currency,omitempty"`

		Cycle string `json:"cycle,omitempty"`

		CycleExact string `json:"cycleExact,omitempty"`

		CycleTerm int32 `json:"cycleTerm,omitempty"`

		CycleType CycleTermType `json:"cycleType,omitempty"`

		DiscountAmount float32 `json:"discountAmount,omitempty"`

		Identity string `json:"identity,omitempty"`

		InitialTermEndDate string `json:"initialTermEndDate,omitempty"`

		InitialTermStartDate string `json:"initialTermStartDate,omitempty"`

		OfferFid string `json:"offerFid,omitempty"`

		ParentFid string `json:"parentFid,omitempty"`

		Price float32 `json:"price,omitempty"`

		PriceFid string `json:"priceFid,omitempty"`

		ProductFid string `json:"productFid,omitempty"`

		PurchaseFid string `json:"purchaseFid,omitempty"`

		Quantity int64 `json:"quantity,omitempty"`

		RenewalDate strfmt.DateTime `json:"renewalDate,omitempty"`

		SetupDiscountAmount float32 `json:"setupDiscountAmount,omitempty"`

		SetupFee float32 `json:"setupFee,omitempty"`

		SkuFid string `json:"skuFid,omitempty"`

		TaxAmount float32 `json:"taxAmount,omitempty"`

		TotalAmount float32 `json:"totalAmount,omitempty"`
	}

	dataAO1.Currency = m.Currency

	dataAO1.Cycle = m.Cycle

	dataAO1.CycleExact = m.CycleExact

	dataAO1.CycleTerm = m.CycleTerm

	dataAO1.CycleType = m.CycleType

	dataAO1.DiscountAmount = m.DiscountAmount

	dataAO1.Identity = m.Identity

	dataAO1.InitialTermEndDate = m.InitialTermEndDate

	dataAO1.InitialTermStartDate = m.InitialTermStartDate

	dataAO1.OfferFid = m.OfferFid

	dataAO1.ParentFid = m.ParentFid

	dataAO1.Price = m.Price

	dataAO1.PriceFid = m.PriceFid

	dataAO1.ProductFid = m.ProductFid

	dataAO1.PurchaseFid = m.PurchaseFid

	dataAO1.Quantity = m.Quantity

	dataAO1.RenewalDate = m.RenewalDate

	dataAO1.SetupDiscountAmount = m.SetupDiscountAmount

	dataAO1.SetupFee = m.SetupFee

	dataAO1.SkuFid = m.SkuFid

	dataAO1.TaxAmount = m.TaxAmount

	dataAO1.TotalAmount = m.TotalAmount

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this order product
func (m *OrderProduct) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Entity
	if err := m.Entity.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCycleType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRenewalDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrderProduct) validateCycleType(formats strfmt.Registry) error {

	if swag.IsZero(m.CycleType) { // not required
		return nil
	}

	if err := m.CycleType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("cycleType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("cycleType")
		}
		return err
	}

	return nil
}

func (m *OrderProduct) validateRenewalDate(formats strfmt.Registry) error {

	if swag.IsZero(m.RenewalDate) { // not required
		return nil
	}

	if err := validate.FormatOf("renewalDate", "body", "date-time", m.RenewalDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this order product based on the context it is used
func (m *OrderProduct) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Entity
	if err := m.Entity.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCycleType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrderProduct) contextValidateCycleType(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.CycleType) { // not required
		return nil
	}

	if err := m.CycleType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("cycleType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("cycleType")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OrderProduct) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrderProduct) UnmarshalBinary(b []byte) error {
	var res OrderProduct
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
