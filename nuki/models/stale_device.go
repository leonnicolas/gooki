// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StaleDevice stale device
//
// swagger:model StaleDevice
type StaleDevice struct {

	// name
	Name string `json:"name,omitempty"`

	// read
	Read bool `json:"read,omitempty"`

	// smartlock Id
	SmartlockID int64 `json:"smartlockId,omitempty"`
}

// Validate validates this stale device
func (m *StaleDevice) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this stale device based on context it is used
func (m *StaleDevice) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StaleDevice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StaleDevice) UnmarshalBinary(b []byte) error {
	var res StaleDevice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
