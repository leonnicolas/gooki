// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EncoderService encoder service
//
// swagger:model EncoderService
type EncoderService struct {

	// accepted media types
	AcceptedMediaTypes []*MediaType `json:"acceptedMediaTypes"`

	// context
	Context *Context `json:"context,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// ignored media types
	IgnoredMediaTypes []*MediaType `json:"ignoredMediaTypes"`

	// minimum size
	MinimumSize int64 `json:"minimumSize,omitempty"`

	// started
	Started bool `json:"started,omitempty"`

	// stopped
	Stopped bool `json:"stopped,omitempty"`
}

// Validate validates this encoder service
func (m *EncoderService) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAcceptedMediaTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIgnoredMediaTypes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EncoderService) validateAcceptedMediaTypes(formats strfmt.Registry) error {
	if swag.IsZero(m.AcceptedMediaTypes) { // not required
		return nil
	}

	for i := 0; i < len(m.AcceptedMediaTypes); i++ {
		if swag.IsZero(m.AcceptedMediaTypes[i]) { // not required
			continue
		}

		if m.AcceptedMediaTypes[i] != nil {
			if err := m.AcceptedMediaTypes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("acceptedMediaTypes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("acceptedMediaTypes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EncoderService) validateContext(formats strfmt.Registry) error {
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

func (m *EncoderService) validateIgnoredMediaTypes(formats strfmt.Registry) error {
	if swag.IsZero(m.IgnoredMediaTypes) { // not required
		return nil
	}

	for i := 0; i < len(m.IgnoredMediaTypes); i++ {
		if swag.IsZero(m.IgnoredMediaTypes[i]) { // not required
			continue
		}

		if m.IgnoredMediaTypes[i] != nil {
			if err := m.IgnoredMediaTypes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ignoredMediaTypes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ignoredMediaTypes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this encoder service based on the context it is used
func (m *EncoderService) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAcceptedMediaTypes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateContext(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIgnoredMediaTypes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EncoderService) contextValidateAcceptedMediaTypes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AcceptedMediaTypes); i++ {

		if m.AcceptedMediaTypes[i] != nil {

			if swag.IsZero(m.AcceptedMediaTypes[i]) { // not required
				return nil
			}

			if err := m.AcceptedMediaTypes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("acceptedMediaTypes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("acceptedMediaTypes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EncoderService) contextValidateContext(ctx context.Context, formats strfmt.Registry) error {

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

func (m *EncoderService) contextValidateIgnoredMediaTypes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.IgnoredMediaTypes); i++ {

		if m.IgnoredMediaTypes[i] != nil {

			if swag.IsZero(m.IgnoredMediaTypes[i]) { // not required
				return nil
			}

			if err := m.IgnoredMediaTypes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ignoredMediaTypes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ignoredMediaTypes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *EncoderService) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EncoderService) UnmarshalBinary(b []byte) error {
	var res EncoderService
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
