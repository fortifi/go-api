// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OrderAllOf1 order all of1
// swagger:model orderAllOf1
type OrderAllOf1 struct {

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

	// fraud fid
	FraudFid string `json:"fraudFid,omitempty"`

	// invoice fid
	InvoiceFid string `json:"invoiceFid,omitempty"`

	// last payment fid
	LastPaymentFid string `json:"lastPaymentFid,omitempty"`

	// offer fid
	OfferFid string `json:"offerFid,omitempty"`

	// order hash
	OrderHash string `json:"orderHash,omitempty"`

	// payment account fid
	PaymentAccountFid string `json:"paymentAccountFid,omitempty"`

	// payment service type
	PaymentServiceType string `json:"paymentServiceType,omitempty"`

	// queue fid
	QueueFid string `json:"queueFid,omitempty"`

	// setup amount
	SetupAmount float32 `json:"setupAmount,omitempty"`

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

// Validate validates this order all of1
func (m *OrderAllOf1) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OrderAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrderAllOf1) UnmarshalBinary(b []byte) error {
	var res OrderAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
