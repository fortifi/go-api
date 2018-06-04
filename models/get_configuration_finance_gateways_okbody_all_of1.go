// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetConfigurationFinanceGatewaysOKBodyAllOf1 get configuration finance gateways o k body all of1
// swagger:model getConfigurationFinanceGatewaysOKBodyAllOf1
type GetConfigurationFinanceGatewaysOKBodyAllOf1 struct {

	// data
	Data *PaymentGateways `json:"data,omitempty"`
}

// Validate validates this get configuration finance gateways o k body all of1
func (m *GetConfigurationFinanceGatewaysOKBodyAllOf1) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *GetConfigurationFinanceGatewaysOKBodyAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetConfigurationFinanceGatewaysOKBodyAllOf1) UnmarshalBinary(b []byte) error {
	var res GetConfigurationFinanceGatewaysOKBodyAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
