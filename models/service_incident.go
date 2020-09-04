// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ServiceIncident Service Incident
//
// swagger:model ServiceIncident
type ServiceIncident struct {
	Entity

	// date created
	DateCreated int64 `json:"dateCreated,omitempty"`

	// date modified
	DateModified int64 `json:"dateModified,omitempty"`

	// date state changed
	DateStateChanged int64 `json:"dateStateChanged,omitempty"`

	// service fid
	ServiceFid string `json:"serviceFid,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// time ended
	TimeEnded int64 `json:"timeEnded,omitempty"`

	// time started
	TimeStarted int64 `json:"timeStarted,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ServiceIncident) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 Entity
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Entity = aO0

	// AO1
	var dataAO1 struct {
		DateCreated int64 `json:"dateCreated,omitempty"`

		DateModified int64 `json:"dateModified,omitempty"`

		DateStateChanged int64 `json:"dateStateChanged,omitempty"`

		ServiceFid string `json:"serviceFid,omitempty"`

		Status string `json:"status,omitempty"`

		TimeEnded int64 `json:"timeEnded,omitempty"`

		TimeStarted int64 `json:"timeStarted,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.DateCreated = dataAO1.DateCreated

	m.DateModified = dataAO1.DateModified

	m.DateStateChanged = dataAO1.DateStateChanged

	m.ServiceFid = dataAO1.ServiceFid

	m.Status = dataAO1.Status

	m.TimeEnded = dataAO1.TimeEnded

	m.TimeStarted = dataAO1.TimeStarted

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ServiceIncident) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.Entity)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		DateCreated int64 `json:"dateCreated,omitempty"`

		DateModified int64 `json:"dateModified,omitempty"`

		DateStateChanged int64 `json:"dateStateChanged,omitempty"`

		ServiceFid string `json:"serviceFid,omitempty"`

		Status string `json:"status,omitempty"`

		TimeEnded int64 `json:"timeEnded,omitempty"`

		TimeStarted int64 `json:"timeStarted,omitempty"`
	}

	dataAO1.DateCreated = m.DateCreated

	dataAO1.DateModified = m.DateModified

	dataAO1.DateStateChanged = m.DateStateChanged

	dataAO1.ServiceFid = m.ServiceFid

	dataAO1.Status = m.Status

	dataAO1.TimeEnded = m.TimeEnded

	dataAO1.TimeStarted = m.TimeStarted

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this service incident
func (m *ServiceIncident) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Entity
	if err := m.Entity.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ServiceIncident) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceIncident) UnmarshalBinary(b []byte) error {
	var res ServiceIncident
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
