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

// OrderProductPayload Product information to add to order
//
// swagger:model OrderProductPayload
type OrderProductPayload struct {

	// Interval in ISO 8601 standard
	Cycle string `json:"cycle,omitempty"`

	// Cycle Term
	CycleTerm int32 `json:"cycleTerm,omitempty"`

	// cycle type
	CycleType CycleTermType `json:"cycleType,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// display name
	DisplayName string `json:"displayName,omitempty"`

	// identity
	Identity string `json:"identity,omitempty"`

	// Custom initial term end date (Time in ISO 8601 standard e.g 2015-12-05T13:11:59Z)
	InitialTermEndDate string `json:"initialTermEndDate,omitempty"`

	// Custom initial term start date (Time in ISO 8601 standard e.g 2015-12-05T13:11:59Z)
	InitialTermStartDate string `json:"initialTermStartDate,omitempty"`

	// Offer FID to apply to product
	OfferFid string `json:"offerFid,omitempty"`

	// Product Parent FID or reference
	Parent string `json:"parent,omitempty"`

	// Product Price FID
	PriceFid string `json:"priceFid,omitempty"`

	// Product Fid
	ProductFid string `json:"productFid,omitempty"`

	// properties
	Properties *PropertyBulkSetPayload `json:"properties,omitempty"`

	// quantity
	Quantity int64 `json:"quantity,omitempty"`

	// reservation app
	ReservationApp string `json:"reservationApp,omitempty"`

	// reservation key
	ReservationKey string `json:"reservationKey,omitempty"`

	// Product SKU
	Sku string `json:"sku,omitempty"`

	// Reference for use with parent field
	TransportReference string `json:"transportReference,omitempty"`
}

// Validate validates this order product payload
func (m *OrderProductPayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCycleType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProperties(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrderProductPayload) validateCycleType(formats strfmt.Registry) error {
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

func (m *OrderProductPayload) validateProperties(formats strfmt.Registry) error {
	if swag.IsZero(m.Properties) { // not required
		return nil
	}

	if m.Properties != nil {
		if err := m.Properties.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("properties")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("properties")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this order product payload based on the context it is used
func (m *OrderProductPayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCycleType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProperties(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrderProductPayload) contextValidateCycleType(ctx context.Context, formats strfmt.Registry) error {

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

func (m *OrderProductPayload) contextValidateProperties(ctx context.Context, formats strfmt.Registry) error {

	if m.Properties != nil {

		if swag.IsZero(m.Properties) { // not required
			return nil
		}

		if err := m.Properties.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("properties")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("properties")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OrderProductPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrderProductPayload) UnmarshalBinary(b []byte) error {
	var res OrderProductPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
