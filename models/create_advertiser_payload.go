// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateAdvertiserPayload Payload for creating an advertiser
// swagger:model CreateAdvertiserPayload
type CreateAdvertiserPayload struct {

	// account manager fid
	AccountManagerFid AccountManagerFid `json:"accountManagerFid,omitempty"`

	// brand fid
	BrandFid BrandFid `json:"brandFid,omitempty"`

	// Name of the company
	CompanyName string `json:"companyName,omitempty"`

	// Name of the person
	// Required: true
	ContactName *string `json:"contactName"`

	// display name
	DisplayName DisplayName `json:"displayName,omitempty"`

	// email
	Email Email `json:"email,omitempty"`

	// foundation fid
	FoundationFid FoundationFid `json:"foundationFid,omitempty"`

	// password
	// Required: true
	Password Password `json:"password"`

	// phone
	Phone PhoneNumber `json:"phone,omitempty"`

	// type
	// Required: true
	Type AdvertiserType `json:"type"`

	// username
	// Required: true
	Username Username `json:"username"`

	// website
	Website URL `json:"website,omitempty"`
}

// Validate validates this create advertiser payload
func (m *CreateAdvertiserPayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContactName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateAdvertiserPayload) validateContactName(formats strfmt.Registry) error {

	if err := validate.Required("contactName", "body", m.ContactName); err != nil {
		return err
	}

	return nil
}

func (m *CreateAdvertiserPayload) validatePassword(formats strfmt.Registry) error {

	if err := m.Password.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("password")
		}
		return err
	}

	return nil
}

func (m *CreateAdvertiserPayload) validateType(formats strfmt.Registry) error {

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		}
		return err
	}

	return nil
}

func (m *CreateAdvertiserPayload) validateUsername(formats strfmt.Registry) error {

	if err := m.Username.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("username")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateAdvertiserPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateAdvertiserPayload) UnmarshalBinary(b []byte) error {
	var res CreateAdvertiserPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}