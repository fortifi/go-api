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

// InvoiceCreditNotePayload invoice credit note payload
//
// swagger:model InvoiceCreditNotePayload
type InvoiceCreditNotePayload struct {

	// amount
	// Required: true
	Amount *float32 `json:"amount"`

	// Charge Request FID
	ChargeRequestFid string `json:"chargeRequestFid,omitempty"`

	// credit amount type
	// Required: true
	CreditAmountType *CreditAmountType `json:"creditAmountType"`

	// Currency
	Currency string `json:"currency,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// display name
	DisplayName string `json:"displayName,omitempty"`

	// Reason FID
	ReasonFid string `json:"reasonFid,omitempty"`
}

// Validate validates this invoice credit note payload
func (m *InvoiceCreditNotePayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreditAmountType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InvoiceCreditNotePayload) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", m.Amount); err != nil {
		return err
	}

	return nil
}

func (m *InvoiceCreditNotePayload) validateCreditAmountType(formats strfmt.Registry) error {

	if err := validate.Required("creditAmountType", "body", m.CreditAmountType); err != nil {
		return err
	}

	if err := validate.Required("creditAmountType", "body", m.CreditAmountType); err != nil {
		return err
	}

	if m.CreditAmountType != nil {
		if err := m.CreditAmountType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("creditAmountType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("creditAmountType")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this invoice credit note payload based on the context it is used
func (m *InvoiceCreditNotePayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreditAmountType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InvoiceCreditNotePayload) contextValidateCreditAmountType(ctx context.Context, formats strfmt.Registry) error {

	if m.CreditAmountType != nil {

		if err := m.CreditAmountType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("creditAmountType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("creditAmountType")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InvoiceCreditNotePayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InvoiceCreditNotePayload) UnmarshalBinary(b []byte) error {
	var res InvoiceCreditNotePayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
