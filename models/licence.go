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

// Licence Licence
//
// swagger:model Licence
type Licence struct {

	// Time in ISO 8601 standard with optional fractions of a second e.g 2015-12-05T13:11:59.888Z
	// Format: date-time
	AutoCancelDate strfmt.DateTime `json:"autoCancelDate,omitempty"`

	// Time in ISO 8601 standard with optional fractions of a second e.g 2015-12-05T13:11:59.888Z
	// Format: date-time
	AutoSuspendDate strfmt.DateTime `json:"autoSuspendDate,omitempty"`

	// customer fid
	CustomerFid string `json:"customerFid,omitempty"`

	// Interval in ISO 8601 standard
	Cycle string `json:"cycle,omitempty"`

	// cycle exact
	CycleExact string `json:"cycleExact,omitempty"`

	// cycle term
	CycleTerm string `json:"cycleTerm,omitempty"`

	// cycle type
	CycleType CycleTermType `json:"cycleType,omitempty"`

	// Time in ISO 8601 standard with optional fractions of a second e.g 2015-12-05T13:11:59.888Z
	// Format: date-time
	DateCreated strfmt.DateTime `json:"dateCreated,omitempty"`

	// Time in ISO 8601 standard with optional fractions of a second e.g 2015-12-05T13:11:59.888Z
	// Format: date-time
	EndDate strfmt.DateTime `json:"endDate,omitempty"`

	// Time in ISO 8601 standard with optional fractions of a second e.g 2015-12-05T13:11:59.888Z
	// Format: date-time
	LastRenewDate strfmt.DateTime `json:"lastRenewDate,omitempty"`

	// Time in ISO 8601 standard with optional fractions of a second e.g 2015-12-05T13:11:59.888Z
	// Format: date-time
	NextRenewDate strfmt.DateTime `json:"nextRenewDate,omitempty"`

	// purchase fid
	PurchaseFid string `json:"purchaseFid,omitempty"`

	// Time in ISO 8601 standard with optional fractions of a second e.g 2015-12-05T13:11:59.888Z
	// Format: date-time
	RenewDate strfmt.DateTime `json:"renewDate,omitempty"`

	// Time in ISO 8601 standard with optional fractions of a second e.g 2015-12-05T13:11:59.888Z
	// Format: date-time
	StartDate strfmt.DateTime `json:"startDate,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// Time in ISO 8601 standard with optional fractions of a second e.g 2015-12-05T13:11:59.888Z
	// Format: date-time
	TrialEndDate strfmt.DateTime `json:"trialEndDate,omitempty"`

	// Time in ISO 8601 standard with optional fractions of a second e.g 2015-12-05T13:11:59.888Z
	// Format: date-time
	TrialStartDate strfmt.DateTime `json:"trialStartDate,omitempty"`
}

// Validate validates this licence
func (m *Licence) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAutoCancelDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAutoSuspendDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCycleType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastRenewDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNextRenewDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRenewDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrialEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrialStartDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Licence) validateAutoCancelDate(formats strfmt.Registry) error {

	if swag.IsZero(m.AutoCancelDate) { // not required
		return nil
	}

	if err := validate.FormatOf("autoCancelDate", "body", "date-time", m.AutoCancelDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Licence) validateAutoSuspendDate(formats strfmt.Registry) error {

	if swag.IsZero(m.AutoSuspendDate) { // not required
		return nil
	}

	if err := validate.FormatOf("autoSuspendDate", "body", "date-time", m.AutoSuspendDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Licence) validateCycleType(formats strfmt.Registry) error {

	if swag.IsZero(m.CycleType) { // not required
		return nil
	}

	if err := m.CycleType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("cycleType")
		}
		return err
	}

	return nil
}

func (m *Licence) validateDateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.DateCreated) { // not required
		return nil
	}

	if err := validate.FormatOf("dateCreated", "body", "date-time", m.DateCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Licence) validateEndDate(formats strfmt.Registry) error {

	if swag.IsZero(m.EndDate) { // not required
		return nil
	}

	if err := validate.FormatOf("endDate", "body", "date-time", m.EndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Licence) validateLastRenewDate(formats strfmt.Registry) error {

	if swag.IsZero(m.LastRenewDate) { // not required
		return nil
	}

	if err := validate.FormatOf("lastRenewDate", "body", "date-time", m.LastRenewDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Licence) validateNextRenewDate(formats strfmt.Registry) error {

	if swag.IsZero(m.NextRenewDate) { // not required
		return nil
	}

	if err := validate.FormatOf("nextRenewDate", "body", "date-time", m.NextRenewDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Licence) validateRenewDate(formats strfmt.Registry) error {

	if swag.IsZero(m.RenewDate) { // not required
		return nil
	}

	if err := validate.FormatOf("renewDate", "body", "date-time", m.RenewDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Licence) validateStartDate(formats strfmt.Registry) error {

	if swag.IsZero(m.StartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("startDate", "body", "date-time", m.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Licence) validateTrialEndDate(formats strfmt.Registry) error {

	if swag.IsZero(m.TrialEndDate) { // not required
		return nil
	}

	if err := validate.FormatOf("trialEndDate", "body", "date-time", m.TrialEndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Licence) validateTrialStartDate(formats strfmt.Registry) error {

	if swag.IsZero(m.TrialStartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("trialStartDate", "body", "date-time", m.TrialStartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Licence) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Licence) UnmarshalBinary(b []byte) error {
	var res Licence
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
