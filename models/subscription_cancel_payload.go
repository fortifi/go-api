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

// SubscriptionCancelPayload subscription cancel payload
//
// swagger:model SubscriptionCancelPayload
type SubscriptionCancelPayload struct {

	// cancel at next renewal
	CancelAtNextRenewal bool `json:"cancelAtNextRenewal,omitempty"`

	// Reason FID
	ReasonFid string `json:"reasonFid,omitempty"`

	// subscription refund type
	// Required: true
	SubscriptionRefundType *SubscriptionRefundType `json:"subscriptionRefundType"`
}

// Validate validates this subscription cancel payload
func (m *SubscriptionCancelPayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSubscriptionRefundType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SubscriptionCancelPayload) validateSubscriptionRefundType(formats strfmt.Registry) error {

	if err := validate.Required("subscriptionRefundType", "body", m.SubscriptionRefundType); err != nil {
		return err
	}

	if err := validate.Required("subscriptionRefundType", "body", m.SubscriptionRefundType); err != nil {
		return err
	}

	if m.SubscriptionRefundType != nil {
		if err := m.SubscriptionRefundType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subscriptionRefundType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("subscriptionRefundType")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this subscription cancel payload based on the context it is used
func (m *SubscriptionCancelPayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSubscriptionRefundType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SubscriptionCancelPayload) contextValidateSubscriptionRefundType(ctx context.Context, formats strfmt.Registry) error {

	if m.SubscriptionRefundType != nil {

		if err := m.SubscriptionRefundType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subscriptionRefundType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("subscriptionRefundType")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SubscriptionCancelPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubscriptionCancelPayload) UnmarshalBinary(b []byte) error {
	var res SubscriptionCancelPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
