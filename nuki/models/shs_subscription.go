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

// ShsSubscription shs subscription
//
// swagger:model ShsSubscription
type ShsSubscription struct {

	// creation date
	// Format: date-time
	CreationDate strfmt.DateTime `json:"creationDate,omitempty"`

	// expiration date
	// Format: date-time
	ExpirationDate strfmt.DateTime `json:"expirationDate,omitempty"`

	// grace period warning email sent
	GracePeriodWarningEmailSent bool `json:"gracePeriodWarningEmailSent,omitempty"`

	// is grace period warning dismissed
	IsGracePeriodWarningDismissed bool `json:"isGracePeriodWarningDismissed,omitempty"`

	// is in grace period
	IsInGracePeriod bool `json:"isInGracePeriod,omitempty"`

	// shs subscription type
	// Enum: [BUSINESS STANDARD BUSINESS_PLUS API_ONLY]
	ShsSubscriptionType string `json:"shsSubscriptionType,omitempty"`

	// state
	// Enum: [ACTIVE INACTIVE CANCELLED EXPIRED ON_HOLD PENDING PENDING_CANCEL]
	State string `json:"state,omitempty"`

	// type
	// Enum: [B2C B2B]
	Type string `json:"type,omitempty"`

	// update date
	// Format: date-time
	UpdateDate strfmt.DateTime `json:"updateDate,omitempty"`
}

// Validate validates this shs subscription
func (m *ShsSubscription) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpirationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShsSubscriptionType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
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

func (m *ShsSubscription) validateCreationDate(formats strfmt.Registry) error {
	if swag.IsZero(m.CreationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("creationDate", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ShsSubscription) validateExpirationDate(formats strfmt.Registry) error {
	if swag.IsZero(m.ExpirationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("expirationDate", "body", "date-time", m.ExpirationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

var shsSubscriptionTypeShsSubscriptionTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["BUSINESS","STANDARD","BUSINESS_PLUS","API_ONLY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		shsSubscriptionTypeShsSubscriptionTypePropEnum = append(shsSubscriptionTypeShsSubscriptionTypePropEnum, v)
	}
}

const (

	// ShsSubscriptionShsSubscriptionTypeBUSINESS captures enum value "BUSINESS"
	ShsSubscriptionShsSubscriptionTypeBUSINESS string = "BUSINESS"

	// ShsSubscriptionShsSubscriptionTypeSTANDARD captures enum value "STANDARD"
	ShsSubscriptionShsSubscriptionTypeSTANDARD string = "STANDARD"

	// ShsSubscriptionShsSubscriptionTypeBUSINESSPLUS captures enum value "BUSINESS_PLUS"
	ShsSubscriptionShsSubscriptionTypeBUSINESSPLUS string = "BUSINESS_PLUS"

	// ShsSubscriptionShsSubscriptionTypeAPIONLY captures enum value "API_ONLY"
	ShsSubscriptionShsSubscriptionTypeAPIONLY string = "API_ONLY"
)

// prop value enum
func (m *ShsSubscription) validateShsSubscriptionTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, shsSubscriptionTypeShsSubscriptionTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ShsSubscription) validateShsSubscriptionType(formats strfmt.Registry) error {
	if swag.IsZero(m.ShsSubscriptionType) { // not required
		return nil
	}

	// value enum
	if err := m.validateShsSubscriptionTypeEnum("shsSubscriptionType", "body", m.ShsSubscriptionType); err != nil {
		return err
	}

	return nil
}

var shsSubscriptionTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ACTIVE","INACTIVE","CANCELLED","EXPIRED","ON_HOLD","PENDING","PENDING_CANCEL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		shsSubscriptionTypeStatePropEnum = append(shsSubscriptionTypeStatePropEnum, v)
	}
}

const (

	// ShsSubscriptionStateACTIVE captures enum value "ACTIVE"
	ShsSubscriptionStateACTIVE string = "ACTIVE"

	// ShsSubscriptionStateINACTIVE captures enum value "INACTIVE"
	ShsSubscriptionStateINACTIVE string = "INACTIVE"

	// ShsSubscriptionStateCANCELLED captures enum value "CANCELLED"
	ShsSubscriptionStateCANCELLED string = "CANCELLED"

	// ShsSubscriptionStateEXPIRED captures enum value "EXPIRED"
	ShsSubscriptionStateEXPIRED string = "EXPIRED"

	// ShsSubscriptionStateONHOLD captures enum value "ON_HOLD"
	ShsSubscriptionStateONHOLD string = "ON_HOLD"

	// ShsSubscriptionStatePENDING captures enum value "PENDING"
	ShsSubscriptionStatePENDING string = "PENDING"

	// ShsSubscriptionStatePENDINGCANCEL captures enum value "PENDING_CANCEL"
	ShsSubscriptionStatePENDINGCANCEL string = "PENDING_CANCEL"
)

// prop value enum
func (m *ShsSubscription) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, shsSubscriptionTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ShsSubscription) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

var shsSubscriptionTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["B2C","B2B"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		shsSubscriptionTypeTypePropEnum = append(shsSubscriptionTypeTypePropEnum, v)
	}
}

const (

	// ShsSubscriptionTypeB2C captures enum value "B2C"
	ShsSubscriptionTypeB2C string = "B2C"

	// ShsSubscriptionTypeB2B captures enum value "B2B"
	ShsSubscriptionTypeB2B string = "B2B"
)

// prop value enum
func (m *ShsSubscription) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, shsSubscriptionTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ShsSubscription) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *ShsSubscription) validateUpdateDate(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdateDate) { // not required
		return nil
	}

	if err := validate.FormatOf("updateDate", "body", "date-time", m.UpdateDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this shs subscription based on context it is used
func (m *ShsSubscription) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ShsSubscription) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ShsSubscription) UnmarshalBinary(b []byte) error {
	var res ShsSubscription
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}