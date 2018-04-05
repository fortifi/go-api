// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// AdvertiserAllOf1 advertiser all of1
// swagger:model advertiserAllOf1
type AdvertiserAllOf1 struct {

	// accepted terms
	AcceptedTerms string `json:"acceptedTerms,omitempty"`

	// account manager fid
	AccountManagerFid string `json:"accountManagerFid,omitempty"`

	// approved
	Approved bool `json:"approved,omitempty"`

	// company fid
	CompanyFid string `json:"companyFid,omitempty"`

	// company name
	CompanyName string `json:"companyName,omitempty"`

	// contact name
	ContactName string `json:"contactName,omitempty"`

	// display name
	DisplayName string `json:"displayName,omitempty"`

	// email
	Email string `json:"email,omitempty"`

	// email fid
	EmailFid string `json:"emailFid,omitempty"`

	// foundation fid
	FoundationFid string `json:"foundationFid,omitempty"`

	// is disabled
	IsDisabled string `json:"isDisabled,omitempty"`

	// payout type
	PayoutType string `json:"payoutType,omitempty"`

	// phone
	Phone string `json:"phone,omitempty"`

	// phone fid
	PhoneFid string `json:"phoneFid,omitempty"`

	// suspended
	Suspended string `json:"suspended,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// website
	Website string `json:"website,omitempty"`
}

// Validate validates this advertiser all of1
func (m *AdvertiserAllOf1) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *AdvertiserAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AdvertiserAllOf1) UnmarshalBinary(b []byte) error {
	var res AdvertiserAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
