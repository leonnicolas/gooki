// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DecoderService decoder service
//
// swagger:model DecoderService
type DecoderService struct {

	// context
	Context *Context `json:"context,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// started
	Started bool `json:"started,omitempty"`

	// stopped
	Stopped bool `json:"stopped,omitempty"`
}

// Validate validates this decoder service
func (m *DecoderService) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContext(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DecoderService) validateContext(formats strfmt.Registry) error {
	if swag.IsZero(m.Context) { // not required
		return nil
	}

	if m.Context != nil {
		if err := m.Context.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("context")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("context")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this decoder service based on the context it is used
func (m *DecoderService) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateContext(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DecoderService) contextValidateContext(ctx context.Context, formats strfmt.Registry) error {

	if m.Context != nil {

		if swag.IsZero(m.Context) { // not required
			return nil
		}

		if err := m.Context.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("context")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("context")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DecoderService) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DecoderService) UnmarshalBinary(b []byte) error {
	var res DecoderService
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}