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

// Workflow workflow
// swagger:model Workflow
type Workflow struct {
	WorkflowSummary

	// jobs
	Jobs []*Job `json:"jobs"`

	// output
	Output string `json:"output,omitempty"`

	// status reason
	StatusReason string `json:"statusReason,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *Workflow) UnmarshalJSON(raw []byte) error {

	var aO0 WorkflowSummary
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.WorkflowSummary = aO0

	var data struct {
		Jobs []*Job `json:"jobs,omitempty"`

		Output string `json:"output,omitempty"`

		StatusReason string `json:"statusReason,omitempty"`
	}
	if err := swag.ReadJSON(raw, &data); err != nil {
		return err
	}

	m.Jobs = data.Jobs

	m.Output = data.Output

	m.StatusReason = data.StatusReason

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m Workflow) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	aO0, err := swag.WriteJSON(m.WorkflowSummary)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var data struct {
		Jobs []*Job `json:"jobs,omitempty"`

		Output string `json:"output,omitempty"`

		StatusReason string `json:"statusReason,omitempty"`
	}

	data.Jobs = m.Jobs

	data.Output = m.Output

	data.StatusReason = m.StatusReason

	jsonData, err := swag.WriteJSON(data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, jsonData)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this workflow
func (m *Workflow) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.WorkflowSummary.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJobs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Workflow) validateJobs(formats strfmt.Registry) error {

	if swag.IsZero(m.Jobs) { // not required
		return nil
	}

	for i := 0; i < len(m.Jobs); i++ {

		if swag.IsZero(m.Jobs[i]) { // not required
			continue
		}

		if m.Jobs[i] != nil {

			if err := m.Jobs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("jobs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Workflow) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Workflow) UnmarshalBinary(b []byte) error {
	var res Workflow
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
