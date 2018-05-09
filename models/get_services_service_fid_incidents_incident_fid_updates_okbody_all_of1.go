// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetServicesServiceFidIncidentsIncidentFidUpdatesOKBodyAllOf1 get services service fid incidents incident fid updates o k body all of1
// swagger:model getServicesServiceFidIncidentsIncidentFidUpdatesOKBodyAllOf1
type GetServicesServiceFidIncidentsIncidentFidUpdatesOKBodyAllOf1 struct {

	// data
	Data *IncidentUpdates `json:"data,omitempty"`
}

// Validate validates this get services service fid incidents incident fid updates o k body all of1
func (m *GetServicesServiceFidIncidentsIncidentFidUpdatesOKBodyAllOf1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetServicesServiceFidIncidentsIncidentFidUpdatesOKBodyAllOf1) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	if m.Data != nil {

		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetServicesServiceFidIncidentsIncidentFidUpdatesOKBodyAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetServicesServiceFidIncidentsIncidentFidUpdatesOKBodyAllOf1) UnmarshalBinary(b []byte) error {
	var res GetServicesServiceFidIncidentsIncidentFidUpdatesOKBodyAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
