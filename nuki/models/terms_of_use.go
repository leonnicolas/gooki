// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TermsOfUse terms of use
//
// swagger:model TermsOfUse
type TermsOfUse struct {

	// acceptance date
	// Format: date-time
	AcceptanceDate strfmt.DateTime `json:"acceptanceDate,omitempty"`

	// publish date
	// Format: date-time
	PublishDate strfmt.DateTime `json:"publishDate,omitempty"`

	// state
	// Enum: [Accepted Inactive]
	State string `json:"state,omitempty"`
}

// Validate validates this terms of use
func (m *TermsOfUse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAcceptanceDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublishDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TermsOfUse) validateAcceptanceDate(formats strfmt.Registry) error {
	if swag.IsZero(m.AcceptanceDate) { // not required
		return nil
	}

	if err := validate.FormatOf("acceptanceDate", "body", "date-time", m.AcceptanceDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TermsOfUse) validatePublishDate(formats strfmt.Registry) error {
	if swag.IsZero(m.PublishDate) { // not required
		return nil
	}

	if err := validate.FormatOf("publishDate", "body", "date-time", m.PublishDate.String(), formats); err != nil {
		return err
	}

	return nil
}

var termsOfUseTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Accepted","Inactive"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		termsOfUseTypeStatePropEnum = append(termsOfUseTypeStatePropEnum, v)
	}
}

const (

	// TermsOfUseStateAccepted captures enum value "Accepted"
	TermsOfUseStateAccepted string = "Accepted"

	// TermsOfUseStateInactive captures enum value "Inactive"
	TermsOfUseStateInactive string = "Inactive"
)

// prop value enum
func (m *TermsOfUse) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, termsOfUseTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TermsOfUse) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this terms of use based on context it is used
func (m *TermsOfUse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TermsOfUse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TermsOfUse) UnmarshalBinary(b []byte) error {
	var res TermsOfUse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
