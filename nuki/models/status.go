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

// Status status
//
// swagger:model Status
type Status struct {

	// client error
	ClientError bool `json:"clientError,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// connector error
	ConnectorError bool `json:"connectorError,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// error
	Error bool `json:"error,omitempty"`

	// global error
	GlobalError bool `json:"globalError,omitempty"`

	// informational
	Informational bool `json:"informational,omitempty"`

	// reason phrase
	ReasonPhrase string `json:"reasonPhrase,omitempty"`

	// recoverable error
	RecoverableError bool `json:"recoverableError,omitempty"`

	// redirection
	Redirection bool `json:"redirection,omitempty"`

	// server error
	ServerError bool `json:"serverError,omitempty"`

	// success
	Success bool `json:"success,omitempty"`

	// throwable
	Throwable *Throwable `json:"throwable,omitempty"`

	// uri
	URI string `json:"uri,omitempty"`
}

// Validate validates this status
func (m *Status) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateThrowable(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Status) validateThrowable(formats strfmt.Registry) error {
	if swag.IsZero(m.Throwable) { // not required
		return nil
	}

	if m.Throwable != nil {
		if err := m.Throwable.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("throwable")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("throwable")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this status based on the context it is used
func (m *Status) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateThrowable(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Status) contextValidateThrowable(ctx context.Context, formats strfmt.Registry) error {

	if m.Throwable != nil {

		if swag.IsZero(m.Throwable) { // not required
			return nil
		}

		if err := m.Throwable.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("throwable")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("throwable")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Status) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Status) UnmarshalBinary(b []byte) error {
	var res Status
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
