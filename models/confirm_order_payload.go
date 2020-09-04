// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConfirmOrderPayload confirm order payload
//
// swagger:model ConfirmOrderPayload
type ConfirmOrderPayload struct {

	// FID for the payment account you wish to charge the customer through
	PaymentAccountFid string `json:"paymentAccountFid,omitempty"`

	// FID for the payment service you wish to charge the customer through
	PaymentServiceFid string `json:"paymentServiceFid,omitempty"`

	// payment service processor
	PaymentServiceProcessor PaymentServiceProcessor `json:"paymentServiceProcessor,omitempty"`
}

// Validate validates this confirm order payload
func (m *ConfirmOrderPayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePaymentServiceProcessor(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfirmOrderPayload) validatePaymentServiceProcessor(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentServiceProcessor) { // not required
		return nil
	}

	if err := m.PaymentServiceProcessor.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("paymentServiceProcessor")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfirmOrderPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfirmOrderPayload) UnmarshalBinary(b []byte) error {
	var res ConfirmOrderPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
