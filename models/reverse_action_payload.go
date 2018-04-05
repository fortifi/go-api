// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ReverseActionPayload reverse action payload
// swagger:model ReverseActionPayload
type ReverseActionPayload struct {

	// client Ip
	ClientIP ClientIP `json:"clientIp,omitempty"`

	// encoding
	Encoding Encoding `json:"encoding,omitempty"`

	// event Id
	EventID EventID `json:"eventId,omitempty"`

	// external reference
	ExternalReference VisitorExternalReference `json:"externalReference,omitempty"`

	// language
	Language Language `json:"language,omitempty"`

	// meta data
	MetaData MetaData `json:"metaData"`

	// reason
	Reason ReversalReason `json:"reason,omitempty"`

	// reversal amount
	ReversalAmount ReversalAmount `json:"reversalAmount,omitempty"`

	// reversal Id
	ReversalID ReversalID `json:"reversalId,omitempty"`

	// source transaction Id
	SourceTransactionID SourceTransactionID `json:"sourceTransactionId,omitempty"`

	// time
	Time IsoTime `json:"time,omitempty"`

	// user agent
	UserAgent UserAgent `json:"userAgent,omitempty"`
}

// Validate validates this reverse action payload
func (m *ReverseActionPayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReason(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReverseActionPayload) validateReason(formats strfmt.Registry) error {

	if swag.IsZero(m.Reason) { // not required
		return nil
	}

	if err := m.Reason.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("reason")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReverseActionPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReverseActionPayload) UnmarshalBinary(b []byte) error {
	var res ReverseActionPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}