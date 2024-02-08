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

// Variant variant
//
// swagger:model Variant
type Variant struct {

	// character set
	CharacterSet *CharacterSet `json:"characterSet,omitempty"`

	// encodings
	Encodings []*Encoding `json:"encodings"`

	// languages
	Languages []*Language `json:"languages"`

	// location ref
	LocationRef *Reference `json:"locationRef,omitempty"`

	// media type
	MediaType *MediaType `json:"mediaType,omitempty"`
}

// Validate validates this variant
func (m *Variant) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCharacterSet(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEncodings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLanguages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocationRef(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMediaType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Variant) validateCharacterSet(formats strfmt.Registry) error {
	if swag.IsZero(m.CharacterSet) { // not required
		return nil
	}

	if m.CharacterSet != nil {
		if err := m.CharacterSet.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("characterSet")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("characterSet")
			}
			return err
		}
	}

	return nil
}

func (m *Variant) validateEncodings(formats strfmt.Registry) error {
	if swag.IsZero(m.Encodings) { // not required
		return nil
	}

	for i := 0; i < len(m.Encodings); i++ {
		if swag.IsZero(m.Encodings[i]) { // not required
			continue
		}

		if m.Encodings[i] != nil {
			if err := m.Encodings[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("encodings" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("encodings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Variant) validateLanguages(formats strfmt.Registry) error {
	if swag.IsZero(m.Languages) { // not required
		return nil
	}

	for i := 0; i < len(m.Languages); i++ {
		if swag.IsZero(m.Languages[i]) { // not required
			continue
		}

		if m.Languages[i] != nil {
			if err := m.Languages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("languages" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("languages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Variant) validateLocationRef(formats strfmt.Registry) error {
	if swag.IsZero(m.LocationRef) { // not required
		return nil
	}

	if m.LocationRef != nil {
		if err := m.LocationRef.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("locationRef")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("locationRef")
			}
			return err
		}
	}

	return nil
}

func (m *Variant) validateMediaType(formats strfmt.Registry) error {
	if swag.IsZero(m.MediaType) { // not required
		return nil
	}

	if m.MediaType != nil {
		if err := m.MediaType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mediaType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mediaType")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this variant based on the context it is used
func (m *Variant) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCharacterSet(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEncodings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLanguages(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLocationRef(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMediaType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Variant) contextValidateCharacterSet(ctx context.Context, formats strfmt.Registry) error {

	if m.CharacterSet != nil {

		if swag.IsZero(m.CharacterSet) { // not required
			return nil
		}

		if err := m.CharacterSet.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("characterSet")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("characterSet")
			}
			return err
		}
	}

	return nil
}

func (m *Variant) contextValidateEncodings(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Encodings); i++ {

		if m.Encodings[i] != nil {

			if swag.IsZero(m.Encodings[i]) { // not required
				return nil
			}

			if err := m.Encodings[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("encodings" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("encodings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Variant) contextValidateLanguages(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Languages); i++ {

		if m.Languages[i] != nil {

			if swag.IsZero(m.Languages[i]) { // not required
				return nil
			}

			if err := m.Languages[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("languages" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("languages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Variant) contextValidateLocationRef(ctx context.Context, formats strfmt.Registry) error {

	if m.LocationRef != nil {

		if swag.IsZero(m.LocationRef) { // not required
			return nil
		}

		if err := m.LocationRef.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("locationRef")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("locationRef")
			}
			return err
		}
	}

	return nil
}

func (m *Variant) contextValidateMediaType(ctx context.Context, formats strfmt.Registry) error {

	if m.MediaType != nil {

		if swag.IsZero(m.MediaType) { // not required
			return nil
		}

		if err := m.MediaType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mediaType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mediaType")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Variant) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Variant) UnmarshalBinary(b []byte) error {
	var res Variant
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}