// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// AdvertisersAllOf1 advertisers all of1
// swagger:model advertisersAllOf1
type AdvertisersAllOf1 struct {

	// advertisers
	Advertisers []*Advertiser `json:"advertisers"`
}

// Validate validates this advertisers all of1
func (m *AdvertisersAllOf1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdvertisers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AdvertisersAllOf1) validateAdvertisers(formats strfmt.Registry) error {

	if swag.IsZero(m.Advertisers) { // not required
		return nil
	}

	for i := 0; i < len(m.Advertisers); i++ {
		if swag.IsZero(m.Advertisers[i]) { // not required
			continue
		}

		if m.Advertisers[i] != nil {
			if err := m.Advertisers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("advertisers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AdvertisersAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AdvertisersAllOf1) UnmarshalBinary(b []byte) error {
	var res AdvertisersAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
