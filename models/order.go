// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Order An Order
//
// swagger:model Order
type Order struct {
	Entity

	// amount
	Amount float32 `json:"amount,omitempty"`

	// amount paid
	AmountPaid float32 `json:"amountPaid,omitempty"`

	// approved by fid
	ApprovedByFid string `json:"approvedByFid,omitempty"`

	// authorize Id
	AuthorizeID string `json:"authorizeId,omitempty"`

	// brand fid
	BrandFid string `json:"brandFid,omitempty"`

	// charge Id
	ChargeID string `json:"chargeId,omitempty"`

	// country
	Country string `json:"country,omitempty"`

	// coupon code
	CouponCode string `json:"couponCode,omitempty"`

	// currency
	Currency string `json:"currency,omitempty"`

	// customer fid
	CustomerFid string `json:"customerFid,omitempty"`

	// discount amount
	DiscountAmount float32 `json:"discountAmount,omitempty"`

	// discount type
	DiscountType string `json:"discountType,omitempty"`

	// external reference
	ExternalReference string `json:"externalReference,omitempty"`

	// fraud fid
	FraudFid string `json:"fraudFid,omitempty"`

	// has addon
	HasAddon bool `json:"hasAddon,omitempty"`

	// has modify
	HasModify bool `json:"hasModify,omitempty"`

	// has product
	HasProduct bool `json:"hasProduct,omitempty"`

	// has subscription
	HasSubscription bool `json:"hasSubscription,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// invoice fid
	InvoiceFid string `json:"invoiceFid,omitempty"`

	// last payment fid
	LastPaymentFid string `json:"lastPaymentFid,omitempty"`

	// offer fid
	OfferFid string `json:"offerFid,omitempty"`

	// order hash
	OrderHash string `json:"orderHash,omitempty"`

	// order type
	// Enum: [initial purchase cancel prerenew]
	OrderType string `json:"orderType,omitempty"`

	// payment account fid
	PaymentAccountFid string `json:"paymentAccountFid,omitempty"`

	// payment service type
	PaymentServiceType string `json:"paymentServiceType,omitempty"`

	// queue fid
	QueueFid string `json:"queueFid,omitempty"`

	// setup amount
	SetupAmount float32 `json:"setupAmount,omitempty"`

	// setup discount amount
	SetupDiscountAmount float32 `json:"setupDiscountAmount,omitempty"`

	// state
	State int32 `json:"state,omitempty"`

	// tax amount
	TaxAmount float32 `json:"taxAmount,omitempty"`

	// total amount
	TotalAmount float32 `json:"totalAmount,omitempty"`

	// user agent
	UserAgent string `json:"userAgent,omitempty"`

	// user Ip
	UserIP string `json:"userIp,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *Order) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 Entity
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Entity = aO0

	// AO1
	var dataAO1 struct {
		Amount float32 `json:"amount,omitempty"`

		AmountPaid float32 `json:"amountPaid,omitempty"`

		ApprovedByFid string `json:"approvedByFid,omitempty"`

		AuthorizeID string `json:"authorizeId,omitempty"`

		BrandFid string `json:"brandFid,omitempty"`

		ChargeID string `json:"chargeId,omitempty"`

		Country string `json:"country,omitempty"`

		CouponCode string `json:"couponCode,omitempty"`

		Currency string `json:"currency,omitempty"`

		CustomerFid string `json:"customerFid,omitempty"`

		DiscountAmount float32 `json:"discountAmount,omitempty"`

		DiscountType string `json:"discountType,omitempty"`

		ExternalReference string `json:"externalReference,omitempty"`

		FraudFid string `json:"fraudFid,omitempty"`

		HasAddon bool `json:"hasAddon,omitempty"`

		HasModify bool `json:"hasModify,omitempty"`

		HasProduct bool `json:"hasProduct,omitempty"`

		HasSubscription bool `json:"hasSubscription,omitempty"`

		ID string `json:"id,omitempty"`

		InvoiceFid string `json:"invoiceFid,omitempty"`

		LastPaymentFid string `json:"lastPaymentFid,omitempty"`

		OfferFid string `json:"offerFid,omitempty"`

		OrderHash string `json:"orderHash,omitempty"`

		OrderType string `json:"orderType,omitempty"`

		PaymentAccountFid string `json:"paymentAccountFid,omitempty"`

		PaymentServiceType string `json:"paymentServiceType,omitempty"`

		QueueFid string `json:"queueFid,omitempty"`

		SetupAmount float32 `json:"setupAmount,omitempty"`

		SetupDiscountAmount float32 `json:"setupDiscountAmount,omitempty"`

		State int32 `json:"state,omitempty"`

		TaxAmount float32 `json:"taxAmount,omitempty"`

		TotalAmount float32 `json:"totalAmount,omitempty"`

		UserAgent string `json:"userAgent,omitempty"`

		UserIP string `json:"userIp,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Amount = dataAO1.Amount

	m.AmountPaid = dataAO1.AmountPaid

	m.ApprovedByFid = dataAO1.ApprovedByFid

	m.AuthorizeID = dataAO1.AuthorizeID

	m.BrandFid = dataAO1.BrandFid

	m.ChargeID = dataAO1.ChargeID

	m.Country = dataAO1.Country

	m.CouponCode = dataAO1.CouponCode

	m.Currency = dataAO1.Currency

	m.CustomerFid = dataAO1.CustomerFid

	m.DiscountAmount = dataAO1.DiscountAmount

	m.DiscountType = dataAO1.DiscountType

	m.ExternalReference = dataAO1.ExternalReference

	m.FraudFid = dataAO1.FraudFid

	m.HasAddon = dataAO1.HasAddon

	m.HasModify = dataAO1.HasModify

	m.HasProduct = dataAO1.HasProduct

	m.HasSubscription = dataAO1.HasSubscription

	m.ID = dataAO1.ID

	m.InvoiceFid = dataAO1.InvoiceFid

	m.LastPaymentFid = dataAO1.LastPaymentFid

	m.OfferFid = dataAO1.OfferFid

	m.OrderHash = dataAO1.OrderHash

	m.OrderType = dataAO1.OrderType

	m.PaymentAccountFid = dataAO1.PaymentAccountFid

	m.PaymentServiceType = dataAO1.PaymentServiceType

	m.QueueFid = dataAO1.QueueFid

	m.SetupAmount = dataAO1.SetupAmount

	m.SetupDiscountAmount = dataAO1.SetupDiscountAmount

	m.State = dataAO1.State

	m.TaxAmount = dataAO1.TaxAmount

	m.TotalAmount = dataAO1.TotalAmount

	m.UserAgent = dataAO1.UserAgent

	m.UserIP = dataAO1.UserIP

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m Order) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.Entity)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Amount float32 `json:"amount,omitempty"`

		AmountPaid float32 `json:"amountPaid,omitempty"`

		ApprovedByFid string `json:"approvedByFid,omitempty"`

		AuthorizeID string `json:"authorizeId,omitempty"`

		BrandFid string `json:"brandFid,omitempty"`

		ChargeID string `json:"chargeId,omitempty"`

		Country string `json:"country,omitempty"`

		CouponCode string `json:"couponCode,omitempty"`

		Currency string `json:"currency,omitempty"`

		CustomerFid string `json:"customerFid,omitempty"`

		DiscountAmount float32 `json:"discountAmount,omitempty"`

		DiscountType string `json:"discountType,omitempty"`

		ExternalReference string `json:"externalReference,omitempty"`

		FraudFid string `json:"fraudFid,omitempty"`

		HasAddon bool `json:"hasAddon,omitempty"`

		HasModify bool `json:"hasModify,omitempty"`

		HasProduct bool `json:"hasProduct,omitempty"`

		HasSubscription bool `json:"hasSubscription,omitempty"`

		ID string `json:"id,omitempty"`

		InvoiceFid string `json:"invoiceFid,omitempty"`

		LastPaymentFid string `json:"lastPaymentFid,omitempty"`

		OfferFid string `json:"offerFid,omitempty"`

		OrderHash string `json:"orderHash,omitempty"`

		OrderType string `json:"orderType,omitempty"`

		PaymentAccountFid string `json:"paymentAccountFid,omitempty"`

		PaymentServiceType string `json:"paymentServiceType,omitempty"`

		QueueFid string `json:"queueFid,omitempty"`

		SetupAmount float32 `json:"setupAmount,omitempty"`

		SetupDiscountAmount float32 `json:"setupDiscountAmount,omitempty"`

		State int32 `json:"state,omitempty"`

		TaxAmount float32 `json:"taxAmount,omitempty"`

		TotalAmount float32 `json:"totalAmount,omitempty"`

		UserAgent string `json:"userAgent,omitempty"`

		UserIP string `json:"userIp,omitempty"`
	}

	dataAO1.Amount = m.Amount

	dataAO1.AmountPaid = m.AmountPaid

	dataAO1.ApprovedByFid = m.ApprovedByFid

	dataAO1.AuthorizeID = m.AuthorizeID

	dataAO1.BrandFid = m.BrandFid

	dataAO1.ChargeID = m.ChargeID

	dataAO1.Country = m.Country

	dataAO1.CouponCode = m.CouponCode

	dataAO1.Currency = m.Currency

	dataAO1.CustomerFid = m.CustomerFid

	dataAO1.DiscountAmount = m.DiscountAmount

	dataAO1.DiscountType = m.DiscountType

	dataAO1.ExternalReference = m.ExternalReference

	dataAO1.FraudFid = m.FraudFid

	dataAO1.HasAddon = m.HasAddon

	dataAO1.HasModify = m.HasModify

	dataAO1.HasProduct = m.HasProduct

	dataAO1.HasSubscription = m.HasSubscription

	dataAO1.ID = m.ID

	dataAO1.InvoiceFid = m.InvoiceFid

	dataAO1.LastPaymentFid = m.LastPaymentFid

	dataAO1.OfferFid = m.OfferFid

	dataAO1.OrderHash = m.OrderHash

	dataAO1.OrderType = m.OrderType

	dataAO1.PaymentAccountFid = m.PaymentAccountFid

	dataAO1.PaymentServiceType = m.PaymentServiceType

	dataAO1.QueueFid = m.QueueFid

	dataAO1.SetupAmount = m.SetupAmount

	dataAO1.SetupDiscountAmount = m.SetupDiscountAmount

	dataAO1.State = m.State

	dataAO1.TaxAmount = m.TaxAmount

	dataAO1.TotalAmount = m.TotalAmount

	dataAO1.UserAgent = m.UserAgent

	dataAO1.UserIP = m.UserIP

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this order
func (m *Order) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Entity
	if err := m.Entity.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrderType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var orderTypeOrderTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["initial","purchase","cancel","prerenew"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		orderTypeOrderTypePropEnum = append(orderTypeOrderTypePropEnum, v)
	}
}

// property enum
func (m *Order) validateOrderTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, orderTypeOrderTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Order) validateOrderType(formats strfmt.Registry) error {

	if swag.IsZero(m.OrderType) { // not required
		return nil
	}

	// value enum
	if err := m.validateOrderTypeEnum("orderType", "body", m.OrderType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this order based on the context it is used
func (m *Order) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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
func (m *Order) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Order) UnmarshalBinary(b []byte) error {
	var res Order
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
