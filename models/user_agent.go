// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
)

// UserAgent User Agent of the visitors browser 'HTTP_USER_AGENT'
// swagger:model userAgent
type UserAgent string

// Validate validates this user agent
func (m UserAgent) Validate(formats strfmt.Registry) error {
	return nil
}
