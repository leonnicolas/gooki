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

// Role role
//
// swagger:model Role
type Role struct {

	// application
	Application *Application `json:"application,omitempty"`

	// child roles
	ChildRoles []*Role `json:"childRoles"`

	// description
	Description string `json:"description,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this role
func (m *Role) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApplication(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChildRoles(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Role) validateApplication(formats strfmt.Registry) error {
	if swag.IsZero(m.Application) { // not required
		return nil
	}

	if m.Application != nil {
		if err := m.Application.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("application")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("application")
			}
			return err
		}
	}

	return nil
}

func (m *Role) validateChildRoles(formats strfmt.Registry) error {
	if swag.IsZero(m.ChildRoles) { // not required
		return nil
	}

	for i := 0; i < len(m.ChildRoles); i++ {
		if swag.IsZero(m.ChildRoles[i]) { // not required
			continue
		}

		if m.ChildRoles[i] != nil {
			if err := m.ChildRoles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("childRoles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("childRoles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this role based on the context it is used
func (m *Role) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateApplication(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateChildRoles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Role) contextValidateApplication(ctx context.Context, formats strfmt.Registry) error {

	if m.Application != nil {

		if swag.IsZero(m.Application) { // not required
			return nil
		}

		if err := m.Application.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("application")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("application")
			}
			return err
		}
	}

	return nil
}

func (m *Role) contextValidateChildRoles(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ChildRoles); i++ {

		if m.ChildRoles[i] != nil {

			if swag.IsZero(m.ChildRoles[i]) { // not required
				return nil
			}

			if err := m.ChildRoles[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("childRoles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("childRoles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Role) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Role) UnmarshalBinary(b []byte) error {
	var res Role
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}