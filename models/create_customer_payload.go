// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CreateCustomerPayload Payload for creating a customer
// swagger:model CreateCustomerPayload
type CreateCustomerPayload struct {

	// account manager fid
	AccountManagerFid AccountManagerFid `json:"accountManagerFid,omitempty"`

	// account status
	AccountStatus CustomerAccountStatus `json:"accountStatus,omitempty"`

	// account type
	AccountType CustomerAccountType `json:"accountType,omitempty"`

	// brand fid
	// Required: true
	BrandFid BrandFid `json:"brandFid"`

	// client Ip
	ClientIP ClientIP `json:"clientIp,omitempty"`

	// currency
	Currency Currency `json:"currency,omitempty"`

	// email
	Email Email `json:"email,omitempty"`

	// external reference
	ExternalReference ExternalReference `json:"externalReference,omitempty"`

	// first name
	FirstName FirstName `json:"firstName,omitempty"`

	// last name
	LastName LastName `json:"lastName,omitempty"`

	// lifecycle
	Lifecycle CustomerLifecycle `json:"lifecycle,omitempty"`

	// phone number
	PhoneNumber PhoneNumber `json:"phoneNumber,omitempty"`

	// subscription type
	SubscriptionType CustomerSubscriptionType `json:"subscriptionType,omitempty"`

	// time
	Time IsoTime `json:"time,omitempty"`

	// visitor Id
	VisitorID VisitorID `json:"visitorId,omitempty"`
}

// Validate validates this create customer payload
func (m *CreateCustomerPayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateAccountType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateBrandFid(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLifecycle(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSubscriptionType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateCustomerPayload) validateAccountStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.AccountStatus) { // not required
		return nil
	}

	if err := m.AccountStatus.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("accountStatus")
		}
		return err
	}

	return nil
}

func (m *CreateCustomerPayload) validateAccountType(formats strfmt.Registry) error {

	if swag.IsZero(m.AccountType) { // not required
		return nil
	}

	if err := m.AccountType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("accountType")
		}
		return err
	}

	return nil
}

func (m *CreateCustomerPayload) validateBrandFid(formats strfmt.Registry) error {

	if err := m.BrandFid.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("brandFid")
		}
		return err
	}

	return nil
}

func (m *CreateCustomerPayload) validateLifecycle(formats strfmt.Registry) error {

	if swag.IsZero(m.Lifecycle) { // not required
		return nil
	}

	if err := m.Lifecycle.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("lifecycle")
		}
		return err
	}

	return nil
}

func (m *CreateCustomerPayload) validateSubscriptionType(formats strfmt.Registry) error {

	if swag.IsZero(m.SubscriptionType) { // not required
		return nil
	}

	if err := m.SubscriptionType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("subscriptionType")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateCustomerPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateCustomerPayload) UnmarshalBinary(b []byte) error {
	var res CreateCustomerPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}