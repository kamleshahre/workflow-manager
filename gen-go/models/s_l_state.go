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

// SLState s l state
// swagger:model SLState

type SLState struct {

	// catch
	Catch []*SLCatcher `json:"Catch,omitempty"`

	// cause
	Cause string `json:"Cause,omitempty"`

	// choices
	Choices []*SLChoice `json:"Choices,omitempty"`

	// comment
	Comment string `json:"Comment,omitempty"`

	// default
	Default string `json:"Default,omitempty"`

	// end
	End bool `json:"End,omitempty"`

	// error
	Error string `json:"Error,omitempty"`

	// heartbeat seconds
	HeartbeatSeconds int64 `json:"HeartbeatSeconds,omitempty"`

	// input path
	InputPath string `json:"InputPath,omitempty"`

	// next
	Next string `json:"Next,omitempty"`

	// output path
	OutputPath string `json:"OutputPath,omitempty"`

	// resource
	Resource string `json:"Resource,omitempty"`

	// result
	Result string `json:"Result,omitempty"`

	// result path
	ResultPath string `json:"ResultPath,omitempty"`

	// retry
	Retry []*SLRetrier `json:"Retry,omitempty"`

	// timeout seconds
	TimeoutSeconds int64 `json:"TimeoutSeconds,omitempty"`

	// type
	Type SLStateType `json:"Type,omitempty"`
}

/* polymorph SLState Catch false */

/* polymorph SLState Cause false */

/* polymorph SLState Choices false */

/* polymorph SLState Comment false */

/* polymorph SLState Default false */

/* polymorph SLState End false */

/* polymorph SLState Error false */

/* polymorph SLState HeartbeatSeconds false */

/* polymorph SLState InputPath false */

/* polymorph SLState Next false */

/* polymorph SLState OutputPath false */

/* polymorph SLState Resource false */

/* polymorph SLState Result false */

/* polymorph SLState ResultPath false */

/* polymorph SLState Retry false */

/* polymorph SLState TimeoutSeconds false */

/* polymorph SLState Type false */

// Validate validates this s l state
func (m *SLState) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCatch(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateChoices(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRetry(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SLState) validateCatch(formats strfmt.Registry) error {

	if swag.IsZero(m.Catch) { // not required
		return nil
	}

	for i := 0; i < len(m.Catch); i++ {

		if swag.IsZero(m.Catch[i]) { // not required
			continue
		}

		if m.Catch[i] != nil {

			if err := m.Catch[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Catch" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SLState) validateChoices(formats strfmt.Registry) error {

	if swag.IsZero(m.Choices) { // not required
		return nil
	}

	for i := 0; i < len(m.Choices); i++ {

		if swag.IsZero(m.Choices[i]) { // not required
			continue
		}

		if m.Choices[i] != nil {

			if err := m.Choices[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Choices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SLState) validateRetry(formats strfmt.Registry) error {

	if swag.IsZero(m.Retry) { // not required
		return nil
	}

	for i := 0; i < len(m.Retry); i++ {

		if swag.IsZero(m.Retry[i]) { // not required
			continue
		}

		if m.Retry[i] != nil {

			if err := m.Retry[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Retry" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SLState) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SLState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SLState) UnmarshalBinary(b []byte) error {
	var res SLState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
