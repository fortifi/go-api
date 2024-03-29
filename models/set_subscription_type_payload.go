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

// SetSubscriptionTypePayload Subscription Type
//
// swagger:model SetSubscriptionTypePayload
type SetSubscriptionTypePayload struct {

	// subscription type
	SubscriptionType CustomerSubscriptionType `json:"subscriptionType,omitempty"`
}

// Validate validates this set subscription type payload
func (m *SetSubscriptionTypePayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSubscriptionType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SetSubscriptionTypePayload) validateSubscriptionType(formats strfmt.Registry) error {
	if swag.IsZero(m.SubscriptionType) { // not required
		return nil
	}

	if err := m.SubscriptionType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("subscriptionType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("subscriptionType")
		}
		return err
	}

	return nil
}

// ContextValidate validate this set subscription type payload based on the context it is used
func (m *SetSubscriptionTypePayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSubscriptionType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SetSubscriptionTypePayload) contextValidateSubscriptionType(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.SubscriptionType) { // not required
		return nil
	}

	if err := m.SubscriptionType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("subscriptionType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("subscriptionType")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SetSubscriptionTypePayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SetSubscriptionTypePayload) UnmarshalBinary(b []byte) error {
	var res SetSubscriptionTypePayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
