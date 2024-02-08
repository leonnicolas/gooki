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

// BulkWebConfigRequest bulk web config request
//
// swagger:model BulkWebConfigRequest
type BulkWebConfigRequest struct {

	// web config requests
	WebConfigRequests []*WebConfigRequest `json:"webConfigRequests"`
}

// Validate validates this bulk web config request
func (m *BulkWebConfigRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateWebConfigRequests(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BulkWebConfigRequest) validateWebConfigRequests(formats strfmt.Registry) error {
	if swag.IsZero(m.WebConfigRequests) { // not required
		return nil
	}

	for i := 0; i < len(m.WebConfigRequests); i++ {
		if swag.IsZero(m.WebConfigRequests[i]) { // not required
			continue
		}

		if m.WebConfigRequests[i] != nil {
			if err := m.WebConfigRequests[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("webConfigRequests" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("webConfigRequests" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this bulk web config request based on the context it is used
func (m *BulkWebConfigRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateWebConfigRequests(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BulkWebConfigRequest) contextValidateWebConfigRequests(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.WebConfigRequests); i++ {

		if m.WebConfigRequests[i] != nil {

			if swag.IsZero(m.WebConfigRequests[i]) { // not required
				return nil
			}

			if err := m.WebConfigRequests[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("webConfigRequests" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("webConfigRequests" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *BulkWebConfigRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BulkWebConfigRequest) UnmarshalBinary(b []byte) error {
	var res BulkWebConfigRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}