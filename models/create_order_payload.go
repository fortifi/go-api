// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CreateOrderPayload create order payload
// swagger:model CreateOrderPayload
type CreateOrderPayload struct {

	// FID of a valid Brand
	BrandFid string `json:"brandFid,omitempty"`

	// IP Address of the visitor
	ClientIP string `json:"clientIp,omitempty"`

	// FID for the customer placing the order
	CustomerFid string `json:"customerFid,omitempty"`

	// Offer FIDs to apply to the order
	OfferFids []string `json:"offerFids"`

	// FID for the payment account you wish to charge the customer through
	PaymentAccountFid string `json:"paymentAccountFid,omitempty"`

	// Product price FIDs to add
	ProductPriceFids []string `json:"productPriceFids"`

	// type
	Type CreateOrderType `json:"type,omitempty"`

	// User Agent of the visitors browser 'HTTP_USER_AGENT'
	UserAgent string `json:"userAgent,omitempty"`
}

// Validate validates this create order payload
func (m *CreateOrderPayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateOrderPayload) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateOrderPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateOrderPayload) UnmarshalBinary(b []byte) error {
	var res CreateOrderPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
