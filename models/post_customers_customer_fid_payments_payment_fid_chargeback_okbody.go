// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PostCustomersCustomerFidPaymentsPaymentFidChargebackOKBody post customers customer fid payments payment fid chargeback o k body
// swagger:model postCustomersCustomerFidPaymentsPaymentFidChargebackOKBody
type PostCustomersCustomerFidPaymentsPaymentFidChargebackOKBody struct {
	Envelope

	PostCustomersCustomerFidPaymentsPaymentFidChargebackOKBodyAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *PostCustomersCustomerFidPaymentsPaymentFidChargebackOKBody) UnmarshalJSON(raw []byte) error {

	var aO0 Envelope
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Envelope = aO0

	var aO1 PostCustomersCustomerFidPaymentsPaymentFidChargebackOKBodyAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.PostCustomersCustomerFidPaymentsPaymentFidChargebackOKBodyAllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m PostCustomersCustomerFidPaymentsPaymentFidChargebackOKBody) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	aO0, err := swag.WriteJSON(m.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.PostCustomersCustomerFidPaymentsPaymentFidChargebackOKBodyAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post customers customer fid payments payment fid chargeback o k body
func (m *PostCustomersCustomerFidPaymentsPaymentFidChargebackOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.Envelope.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.PostCustomersCustomerFidPaymentsPaymentFidChargebackOKBodyAllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PostCustomersCustomerFidPaymentsPaymentFidChargebackOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostCustomersCustomerFidPaymentsPaymentFidChargebackOKBody) UnmarshalBinary(b []byte) error {
	var res PostCustomersCustomerFidPaymentsPaymentFidChargebackOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}