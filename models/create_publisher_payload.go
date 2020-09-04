// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreatePublisherPayload Payload for creating an advertiser
//
// swagger:model CreatePublisherPayload
type CreatePublisherPayload struct {

	// FID of an Account Manager
	AccountManagerFid string `json:"accountManagerFid,omitempty"`

	// FID of a valid Brand
	BrandFid string `json:"brandFid,omitempty"`

	// Name of the company
	CompanyName string `json:"companyName,omitempty"`

	// Name of the person
	// Required: true
	ContactName *string `json:"contactName"`

	// Visible Display Name
	DisplayName string `json:"displayName,omitempty"`

	// Email Address
	Email string `json:"email,omitempty"`

	// FID of a valid Foundation
	FoundationFid string `json:"foundationFid,omitempty"`

	// Password
	// Required: true
	Password *string `json:"password"`

	// Phone Number
	Phone string `json:"phone,omitempty"`

	// type
	// Required: true
	Type AdvertiserType `json:"type"`

	// Username
	// Required: true
	Username *string `json:"username"`

	// URL
	Website string `json:"website,omitempty"`
}

// Validate validates this create publisher payload
func (m *CreatePublisherPayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContactName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreatePublisherPayload) validateContactName(formats strfmt.Registry) error {

	if err := validate.Required("contactName", "body", m.ContactName); err != nil {
		return err
	}

	return nil
}

func (m *CreatePublisherPayload) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password); err != nil {
		return err
	}

	return nil
}

func (m *CreatePublisherPayload) validateType(formats strfmt.Registry) error {

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		}
		return err
	}

	return nil
}

func (m *CreatePublisherPayload) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreatePublisherPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreatePublisherPayload) UnmarshalBinary(b []byte) error {
	var res CreatePublisherPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
