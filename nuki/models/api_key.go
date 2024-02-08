// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// APIKey Api key
//
// swagger:model ApiKey
type APIKey struct {

	// The account id
	// Required: true
	AccountID *int32 `json:"accountId"`

	// The api key
	APIKey string `json:"apiKey,omitempty"`

	// The id
	// Required: true
	APIKeyID *int32 `json:"apiKeyId"`

	// The creation date
	// Required: true
	// Format: date-time
	CreationDate *strfmt.DateTime `json:"creationDate"`

	// The description
	Description string `json:"description,omitempty"`

	// The redirect uris
	RedirectUris []string `json:"redirectUris"`
}

// Validate validates this Api key
func (m *APIKey) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAPIKeyID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreationDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIKey) validateAccountID(formats strfmt.Registry) error {

	if err := validate.Required("accountId", "body", m.AccountID); err != nil {
		return err
	}

	return nil
}

func (m *APIKey) validateAPIKeyID(formats strfmt.Registry) error {

	if err := validate.Required("apiKeyId", "body", m.APIKeyID); err != nil {
		return err
	}

	return nil
}

func (m *APIKey) validateCreationDate(formats strfmt.Registry) error {

	if err := validate.Required("creationDate", "body", m.CreationDate); err != nil {
		return err
	}

	if err := validate.FormatOf("creationDate", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this Api key based on context it is used
func (m *APIKey) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *APIKey) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIKey) UnmarshalBinary(b []byte) error {
	var res APIKey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
