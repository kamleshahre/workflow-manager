// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// FailureInfo failure info
// swagger:model FailureInfo
type FailureInfo struct {

	// reason
	Reason string `json:"reason,omitempty"`

	// resource
	Resource string `json:"resource,omitempty"`

	// state
	State string `json:"state,omitempty"`
}

// Validate validates this failure info
func (m *FailureInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *FailureInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FailureInfo) UnmarshalBinary(b []byte) error {
	var res FailureInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
