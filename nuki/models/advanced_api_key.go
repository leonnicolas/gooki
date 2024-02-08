// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AdvancedAPIKey advanced Api key
//
// swagger:model AdvancedApiKey
type AdvancedAPIKey struct {

	// The country of the headquarter or the country where you are mainly doing business in
	// Required: true
	Country *string `json:"country"`

	// The creation date
	// Required: true
	// Format: date-time
	CreationDate *strfmt.DateTime `json:"creationDate"`

	// Describe the services and/or products you offer to your customers and how your customers would use Nuki devices in their processes
	// Required: true
	Description *string `json:"description"`

	// An email address where we can contact you for checks on your application
	// Required: true
	Email *string `json:"email"`

	// The name of the company for which you apply for access
	// Required: true
	Name *string `json:"name"`

	// Whether the advanced API key is restricted
	// Required: true
	Restricted *bool `json:"restricted"`

	// The client secret, visible if application is approved (status >= 'TESTING')
	// Required: true
	Secret *string `json:"secret"`

	// The application status
	// Required: true
	// Enum: [INACTIVE APPLIED TESTING ACTIVE]
	Status *string `json:"status"`

	// The application type
	// Required: true
	// Enum: [ONLY_SECRET SHORT_RENTAL HEALTHCARE SMART_HOME OTHER]
	Type *string `json:"type"`

	// The update date
	// Required: true
	// Format: date-time
	UpdateDate *strfmt.DateTime `json:"updateDate"`

	// A website where we can find more information on the company and its business model
	// Required: true
	URL *string `json:"url"`

	// The features to trigger webhooks, for all types except 'ONLY_SECRET'
	// Required: true
	// Unique: true
	WebhookFeatures []string `json:"webhookFeatures"`

	// The status of the webhook posting automation
	// Read Only: true
	// Enum: [ACTIVE DEACTIVATED]
	WebhookStatus string `json:"webhookStatus,omitempty"`

	// The URL where our webhooks should point to
	// Required: true
	WebhookURL *string `json:"webhookUrl"`
}

// Validate validates this advanced Api key
func (m *AdvancedAPIKey) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCountry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRestricted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecret(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWebhookFeatures(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWebhookStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWebhookURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AdvancedAPIKey) validateCountry(formats strfmt.Registry) error {

	if err := validate.Required("country", "body", m.Country); err != nil {
		return err
	}

	return nil
}

func (m *AdvancedAPIKey) validateCreationDate(formats strfmt.Registry) error {

	if err := validate.Required("creationDate", "body", m.CreationDate); err != nil {
		return err
	}

	if err := validate.FormatOf("creationDate", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AdvancedAPIKey) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *AdvancedAPIKey) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	return nil
}

func (m *AdvancedAPIKey) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *AdvancedAPIKey) validateRestricted(formats strfmt.Registry) error {

	if err := validate.Required("restricted", "body", m.Restricted); err != nil {
		return err
	}

	return nil
}

func (m *AdvancedAPIKey) validateSecret(formats strfmt.Registry) error {

	if err := validate.Required("secret", "body", m.Secret); err != nil {
		return err
	}

	return nil
}

var advancedApiKeyTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["INACTIVE","APPLIED","TESTING","ACTIVE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		advancedApiKeyTypeStatusPropEnum = append(advancedApiKeyTypeStatusPropEnum, v)
	}
}

const (

	// AdvancedAPIKeyStatusINACTIVE captures enum value "INACTIVE"
	AdvancedAPIKeyStatusINACTIVE string = "INACTIVE"

	// AdvancedAPIKeyStatusAPPLIED captures enum value "APPLIED"
	AdvancedAPIKeyStatusAPPLIED string = "APPLIED"

	// AdvancedAPIKeyStatusTESTING captures enum value "TESTING"
	AdvancedAPIKeyStatusTESTING string = "TESTING"

	// AdvancedAPIKeyStatusACTIVE captures enum value "ACTIVE"
	AdvancedAPIKeyStatusACTIVE string = "ACTIVE"
)

// prop value enum
func (m *AdvancedAPIKey) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, advancedApiKeyTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AdvancedAPIKey) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

var advancedApiKeyTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ONLY_SECRET","SHORT_RENTAL","HEALTHCARE","SMART_HOME","OTHER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		advancedApiKeyTypeTypePropEnum = append(advancedApiKeyTypeTypePropEnum, v)
	}
}

const (

	// AdvancedAPIKeyTypeONLYSECRET captures enum value "ONLY_SECRET"
	AdvancedAPIKeyTypeONLYSECRET string = "ONLY_SECRET"

	// AdvancedAPIKeyTypeSHORTRENTAL captures enum value "SHORT_RENTAL"
	AdvancedAPIKeyTypeSHORTRENTAL string = "SHORT_RENTAL"

	// AdvancedAPIKeyTypeHEALTHCARE captures enum value "HEALTHCARE"
	AdvancedAPIKeyTypeHEALTHCARE string = "HEALTHCARE"

	// AdvancedAPIKeyTypeSMARTHOME captures enum value "SMART_HOME"
	AdvancedAPIKeyTypeSMARTHOME string = "SMART_HOME"

	// AdvancedAPIKeyTypeOTHER captures enum value "OTHER"
	AdvancedAPIKeyTypeOTHER string = "OTHER"
)

// prop value enum
func (m *AdvancedAPIKey) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, advancedApiKeyTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AdvancedAPIKey) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

func (m *AdvancedAPIKey) validateUpdateDate(formats strfmt.Registry) error {

	if err := validate.Required("updateDate", "body", m.UpdateDate); err != nil {
		return err
	}

	if err := validate.FormatOf("updateDate", "body", "date-time", m.UpdateDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AdvancedAPIKey) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("url", "body", m.URL); err != nil {
		return err
	}

	return nil
}

var advancedApiKeyWebhookFeaturesItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DEVICE_STATUS","DEVICE_MASTERDATA","DEVICE_CONFIG","DEVICE_LOGS","DEVICE_AUTHS","ACCOUNT_USER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		advancedApiKeyWebhookFeaturesItemsEnum = append(advancedApiKeyWebhookFeaturesItemsEnum, v)
	}
}

func (m *AdvancedAPIKey) validateWebhookFeaturesItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, advancedApiKeyWebhookFeaturesItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AdvancedAPIKey) validateWebhookFeatures(formats strfmt.Registry) error {

	if err := validate.Required("webhookFeatures", "body", m.WebhookFeatures); err != nil {
		return err
	}

	if err := validate.UniqueItems("webhookFeatures", "body", m.WebhookFeatures); err != nil {
		return err
	}

	for i := 0; i < len(m.WebhookFeatures); i++ {

		// value enum
		if err := m.validateWebhookFeaturesItemsEnum("webhookFeatures"+"."+strconv.Itoa(i), "body", m.WebhookFeatures[i]); err != nil {
			return err
		}

	}

	return nil
}

var advancedApiKeyTypeWebhookStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ACTIVE","DEACTIVATED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		advancedApiKeyTypeWebhookStatusPropEnum = append(advancedApiKeyTypeWebhookStatusPropEnum, v)
	}
}

const (

	// AdvancedAPIKeyWebhookStatusACTIVE captures enum value "ACTIVE"
	AdvancedAPIKeyWebhookStatusACTIVE string = "ACTIVE"

	// AdvancedAPIKeyWebhookStatusDEACTIVATED captures enum value "DEACTIVATED"
	AdvancedAPIKeyWebhookStatusDEACTIVATED string = "DEACTIVATED"
)

// prop value enum
func (m *AdvancedAPIKey) validateWebhookStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, advancedApiKeyTypeWebhookStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AdvancedAPIKey) validateWebhookStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.WebhookStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateWebhookStatusEnum("webhookStatus", "body", m.WebhookStatus); err != nil {
		return err
	}

	return nil
}

func (m *AdvancedAPIKey) validateWebhookURL(formats strfmt.Registry) error {

	if err := validate.Required("webhookUrl", "body", m.WebhookURL); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this advanced Api key based on the context it is used
func (m *AdvancedAPIKey) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateWebhookStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AdvancedAPIKey) contextValidateWebhookStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "webhookStatus", "body", string(m.WebhookStatus)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AdvancedAPIKey) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AdvancedAPIKey) UnmarshalBinary(b []byte) error {
	var res AdvancedAPIKey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
