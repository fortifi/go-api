// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// VisitorReversalAction visitor reversal action
// swagger:model VisitorReversalAction
type VisitorReversalAction struct {

	// advertiser fid
	AdvertiserFid string `json:"advertiserFid,omitempty"`

	// campaign fid
	CampaignFid string `json:"campaignFid,omitempty"`

	// commission currency
	CommissionCurrency string `json:"commissionCurrency,omitempty"`

	// event Id
	EventID string `json:"eventId,omitempty"`

	// fee charged
	FeeCharged float32 `json:"feeCharged,omitempty"`

	// reversed commission
	ReversedCommission float32 `json:"reversedCommission,omitempty"`

	// visitor Id
	VisitorID string `json:"visitorId,omitempty"`
}

// Validate validates this visitor reversal action
func (m *VisitorReversalAction) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *VisitorReversalAction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VisitorReversalAction) UnmarshalBinary(b []byte) error {
	var res VisitorReversalAction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
