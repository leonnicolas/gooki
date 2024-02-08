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

// MyAccount my account
//
// swagger:model MyAccount
type MyAccount struct {

	// The account id
	// Required: true
	AccountID *int32 `json:"accountId"`

	// api terms of use
	APITermsOfUse *TermsOfUse `json:"apiTermsOfUse,omitempty"`

	// b2b active
	B2bActive bool `json:"b2bActive,omitempty"`

	// The optional config
	Config *AccountConfig `json:"config,omitempty"`

	// The creation date
	// Required: true
	// Format: date-time
	CreationDate *strfmt.DateTime `json:"creationDate"`

	// Set, if your account is not a standard Nuki Web account
	// Read Only: true
	Descent *AccountDescent `json:"descent,omitempty"`

	// The email address
	// Required: true
	Email *string `json:"email"`

	// true, if the email is verified
	EmailVerified bool `json:"emailVerified,omitempty"`

	// The language code
	// Example: de
	Language string `json:"language,omitempty"`

	// The master account id if it's a sub account
	MasterAccountID int32 `json:"masterAccountId,omitempty"`

	// The name
	// Required: true
	Name *string `json:"name"`

	// The optional profile
	Profile *AccountProfile `json:"profile,omitempty"`

	// The rights bitmask if it's a sub account: 1 .. manage smartlock, 2 .. operate smartlock, 4 .. manage smartlock config, 8 .. manage smartlock authorizations, 16 .. view smartlock logs, 32 .. manage sub accounts, 64 .. create smartlocks
	// Maximum: 127
	// Minimum: 0
	Rights *int32 `json:"rights,omitempty"`

	// The secret base64 encoded
	Secret []strfmt.Base64 `json:"secret"`

	// subscription type of the account (b2b)
	// Enum: [BUSINESS STANDARD BUSINESS_PLUS API_ONLY]
	ShsSubscriptionType string `json:"shsSubscriptionType,omitempty"`

	// The type: 0 .. user, 1 .. company, 2 .. caretaker
	// Required: true
	// Maximum: 2
	// Minimum: 0
	Type *int32 `json:"type"`

	// The update date
	// Required: true
	// Format: date-time
	UpdateDate *strfmt.DateTime `json:"updateDate"`
}

// Validate validates this my account
func (m *MyAccount) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAPITermsOfUse(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProfile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRights(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShsSubscriptionType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MyAccount) validateAccountID(formats strfmt.Registry) error {

	if err := validate.Required("accountId", "body", m.AccountID); err != nil {
		return err
	}

	return nil
}

func (m *MyAccount) validateAPITermsOfUse(formats strfmt.Registry) error {
	if swag.IsZero(m.APITermsOfUse) { // not required
		return nil
	}

	if m.APITermsOfUse != nil {
		if err := m.APITermsOfUse.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("apiTermsOfUse")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("apiTermsOfUse")
			}
			return err
		}
	}

	return nil
}

func (m *MyAccount) validateConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.Config) { // not required
		return nil
	}

	if m.Config != nil {
		if err := m.Config.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("config")
			}
			return err
		}
	}

	return nil
}

func (m *MyAccount) validateCreationDate(formats strfmt.Registry) error {

	if err := validate.Required("creationDate", "body", m.CreationDate); err != nil {
		return err
	}

	if err := validate.FormatOf("creationDate", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MyAccount) validateDescent(formats strfmt.Registry) error {
	if swag.IsZero(m.Descent) { // not required
		return nil
	}

	if m.Descent != nil {
		if err := m.Descent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("descent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("descent")
			}
			return err
		}
	}

	return nil
}

func (m *MyAccount) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	return nil
}

func (m *MyAccount) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *MyAccount) validateProfile(formats strfmt.Registry) error {
	if swag.IsZero(m.Profile) { // not required
		return nil
	}

	if m.Profile != nil {
		if err := m.Profile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("profile")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("profile")
			}
			return err
		}
	}

	return nil
}

func (m *MyAccount) validateRights(formats strfmt.Registry) error {
	if swag.IsZero(m.Rights) { // not required
		return nil
	}

	if err := validate.MinimumInt("rights", "body", int64(*m.Rights), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("rights", "body", int64(*m.Rights), 127, false); err != nil {
		return err
	}

	return nil
}

var myAccountTypeShsSubscriptionTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["BUSINESS","STANDARD","BUSINESS_PLUS","API_ONLY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		myAccountTypeShsSubscriptionTypePropEnum = append(myAccountTypeShsSubscriptionTypePropEnum, v)
	}
}

const (

	// MyAccountShsSubscriptionTypeBUSINESS captures enum value "BUSINESS"
	MyAccountShsSubscriptionTypeBUSINESS string = "BUSINESS"

	// MyAccountShsSubscriptionTypeSTANDARD captures enum value "STANDARD"
	MyAccountShsSubscriptionTypeSTANDARD string = "STANDARD"

	// MyAccountShsSubscriptionTypeBUSINESSPLUS captures enum value "BUSINESS_PLUS"
	MyAccountShsSubscriptionTypeBUSINESSPLUS string = "BUSINESS_PLUS"

	// MyAccountShsSubscriptionTypeAPIONLY captures enum value "API_ONLY"
	MyAccountShsSubscriptionTypeAPIONLY string = "API_ONLY"
)

// prop value enum
func (m *MyAccount) validateShsSubscriptionTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, myAccountTypeShsSubscriptionTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MyAccount) validateShsSubscriptionType(formats strfmt.Registry) error {
	if swag.IsZero(m.ShsSubscriptionType) { // not required
		return nil
	}

	// value enum
	if err := m.validateShsSubscriptionTypeEnum("shsSubscriptionType", "body", m.ShsSubscriptionType); err != nil {
		return err
	}

	return nil
}

func (m *MyAccount) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if err := validate.MinimumInt("type", "body", int64(*m.Type), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("type", "body", int64(*m.Type), 2, false); err != nil {
		return err
	}

	return nil
}

func (m *MyAccount) validateUpdateDate(formats strfmt.Registry) error {

	if err := validate.Required("updateDate", "body", m.UpdateDate); err != nil {
		return err
	}

	if err := validate.FormatOf("updateDate", "body", "date-time", m.UpdateDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this my account based on the context it is used
func (m *MyAccount) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAPITermsOfUse(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDescent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProfile(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MyAccount) contextValidateAPITermsOfUse(ctx context.Context, formats strfmt.Registry) error {

	if m.APITermsOfUse != nil {

		if swag.IsZero(m.APITermsOfUse) { // not required
			return nil
		}

		if err := m.APITermsOfUse.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("apiTermsOfUse")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("apiTermsOfUse")
			}
			return err
		}
	}

	return nil
}

func (m *MyAccount) contextValidateConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.Config != nil {

		if swag.IsZero(m.Config) { // not required
			return nil
		}

		if err := m.Config.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("config")
			}
			return err
		}
	}

	return nil
}

func (m *MyAccount) contextValidateDescent(ctx context.Context, formats strfmt.Registry) error {

	if m.Descent != nil {

		if swag.IsZero(m.Descent) { // not required
			return nil
		}

		if err := m.Descent.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("descent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("descent")
			}
			return err
		}
	}

	return nil
}

func (m *MyAccount) contextValidateProfile(ctx context.Context, formats strfmt.Registry) error {

	if m.Profile != nil {

		if swag.IsZero(m.Profile) { // not required
			return nil
		}

		if err := m.Profile.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("profile")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("profile")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MyAccount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MyAccount) UnmarshalBinary(b []byte) error {
	var res MyAccount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
