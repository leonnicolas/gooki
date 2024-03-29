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

// Response response
//
// swagger:model Response
type Response struct {

	// access control allow credentials
	AccessControlAllowCredentials bool `json:"accessControlAllowCredentials,omitempty"`

	// access control allow headers
	// Unique: true
	AccessControlAllowHeaders []string `json:"accessControlAllowHeaders"`

	// access control allow methods
	// Unique: true
	AccessControlAllowMethods []*Method `json:"accessControlAllowMethods"`

	// access control allow origin
	AccessControlAllowOrigin string `json:"accessControlAllowOrigin,omitempty"`

	// access control expose headers
	// Unique: true
	AccessControlExposeHeaders []string `json:"accessControlExposeHeaders"`

	// access control max age
	AccessControlMaxAge int32 `json:"accessControlMaxAge,omitempty"`

	// age
	Age int32 `json:"age,omitempty"`

	// allowed methods
	// Unique: true
	AllowedMethods []*Method `json:"allowedMethods"`

	// attributes
	Attributes map[string]interface{} `json:"attributes,omitempty"`

	// authentication info
	AuthenticationInfo *AuthenticationInfo `json:"authenticationInfo,omitempty"`

	// auto committing
	AutoCommitting bool `json:"autoCommitting,omitempty"`

	// cache directives
	CacheDirectives []*CacheDirective `json:"cacheDirectives"`

	// challenge requests
	ChallengeRequests []*ChallengeRequest `json:"challengeRequests"`

	// committed
	Committed bool `json:"committed,omitempty"`

	// confidential
	Confidential bool `json:"confidential,omitempty"`

	// cookie settings
	CookieSettings []*CookieSetting `json:"cookieSettings"`

	// date
	// Format: date-time
	Date strfmt.DateTime `json:"date,omitempty"`

	// dimensions
	// Unique: true
	Dimensions []string `json:"dimensions"`

	// entity
	Entity *Representation `json:"entity,omitempty"`

	// entity as text
	EntityAsText string `json:"entityAsText,omitempty"`

	// entity available
	EntityAvailable bool `json:"entityAvailable,omitempty"`

	// final
	Final bool `json:"final,omitempty"`

	// headers
	Headers []*Header `json:"headers"`

	// location ref
	LocationRef *Reference `json:"locationRef,omitempty"`

	// on error
	OnError Uniform `json:"onError,omitempty"`

	// on sent
	OnSent Uniform `json:"onSent,omitempty"`

	// provisional
	Provisional bool `json:"provisional,omitempty"`

	// proxy challenge requests
	ProxyChallengeRequests []*ChallengeRequest `json:"proxyChallengeRequests"`

	// recipients info
	RecipientsInfo []*RecipientInfo `json:"recipientsInfo"`

	// request
	Request *Request `json:"request,omitempty"`

	// retry after
	// Format: date-time
	RetryAfter strfmt.DateTime `json:"retryAfter,omitempty"`

	// server info
	ServerInfo *ServerInfo `json:"serverInfo,omitempty"`

	// status
	Status *Status `json:"status,omitempty"`

	// warnings
	Warnings []*Warning `json:"warnings"`
}

// Validate validates this response
func (m *Response) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessControlAllowHeaders(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccessControlAllowMethods(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccessControlExposeHeaders(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAllowedMethods(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthenticationInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCacheDirectives(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChallengeRequests(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCookieSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDimensions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHeaders(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocationRef(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProxyChallengeRequests(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecipientsInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequest(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRetryAfter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServerInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWarnings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Response) validateAccessControlAllowHeaders(formats strfmt.Registry) error {
	if swag.IsZero(m.AccessControlAllowHeaders) { // not required
		return nil
	}

	if err := validate.UniqueItems("accessControlAllowHeaders", "body", m.AccessControlAllowHeaders); err != nil {
		return err
	}

	return nil
}

func (m *Response) validateAccessControlAllowMethods(formats strfmt.Registry) error {
	if swag.IsZero(m.AccessControlAllowMethods) { // not required
		return nil
	}

	if err := validate.UniqueItems("accessControlAllowMethods", "body", m.AccessControlAllowMethods); err != nil {
		return err
	}

	for i := 0; i < len(m.AccessControlAllowMethods); i++ {
		if swag.IsZero(m.AccessControlAllowMethods[i]) { // not required
			continue
		}

		if m.AccessControlAllowMethods[i] != nil {
			if err := m.AccessControlAllowMethods[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("accessControlAllowMethods" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("accessControlAllowMethods" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Response) validateAccessControlExposeHeaders(formats strfmt.Registry) error {
	if swag.IsZero(m.AccessControlExposeHeaders) { // not required
		return nil
	}

	if err := validate.UniqueItems("accessControlExposeHeaders", "body", m.AccessControlExposeHeaders); err != nil {
		return err
	}

	return nil
}

func (m *Response) validateAllowedMethods(formats strfmt.Registry) error {
	if swag.IsZero(m.AllowedMethods) { // not required
		return nil
	}

	if err := validate.UniqueItems("allowedMethods", "body", m.AllowedMethods); err != nil {
		return err
	}

	for i := 0; i < len(m.AllowedMethods); i++ {
		if swag.IsZero(m.AllowedMethods[i]) { // not required
			continue
		}

		if m.AllowedMethods[i] != nil {
			if err := m.AllowedMethods[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("allowedMethods" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("allowedMethods" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Response) validateAuthenticationInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.AuthenticationInfo) { // not required
		return nil
	}

	if m.AuthenticationInfo != nil {
		if err := m.AuthenticationInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("authenticationInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("authenticationInfo")
			}
			return err
		}
	}

	return nil
}

func (m *Response) validateCacheDirectives(formats strfmt.Registry) error {
	if swag.IsZero(m.CacheDirectives) { // not required
		return nil
	}

	for i := 0; i < len(m.CacheDirectives); i++ {
		if swag.IsZero(m.CacheDirectives[i]) { // not required
			continue
		}

		if m.CacheDirectives[i] != nil {
			if err := m.CacheDirectives[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cacheDirectives" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("cacheDirectives" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Response) validateChallengeRequests(formats strfmt.Registry) error {
	if swag.IsZero(m.ChallengeRequests) { // not required
		return nil
	}

	for i := 0; i < len(m.ChallengeRequests); i++ {
		if swag.IsZero(m.ChallengeRequests[i]) { // not required
			continue
		}

		if m.ChallengeRequests[i] != nil {
			if err := m.ChallengeRequests[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("challengeRequests" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("challengeRequests" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Response) validateCookieSettings(formats strfmt.Registry) error {
	if swag.IsZero(m.CookieSettings) { // not required
		return nil
	}

	for i := 0; i < len(m.CookieSettings); i++ {
		if swag.IsZero(m.CookieSettings[i]) { // not required
			continue
		}

		if m.CookieSettings[i] != nil {
			if err := m.CookieSettings[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cookieSettings" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("cookieSettings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Response) validateDate(formats strfmt.Registry) error {
	if swag.IsZero(m.Date) { // not required
		return nil
	}

	if err := validate.FormatOf("date", "body", "date-time", m.Date.String(), formats); err != nil {
		return err
	}

	return nil
}

var responseDimensionsItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AUTHORIZATION","CHARACTER_SET","CLIENT_ADDRESS","CLIENT_AGENT","UNSPECIFIED","ENCODING","LANGUAGE","MEDIA_TYPE","TIME","ORIGIN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		responseDimensionsItemsEnum = append(responseDimensionsItemsEnum, v)
	}
}

func (m *Response) validateDimensionsItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, responseDimensionsItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Response) validateDimensions(formats strfmt.Registry) error {
	if swag.IsZero(m.Dimensions) { // not required
		return nil
	}

	if err := validate.UniqueItems("dimensions", "body", m.Dimensions); err != nil {
		return err
	}

	for i := 0; i < len(m.Dimensions); i++ {

		// value enum
		if err := m.validateDimensionsItemsEnum("dimensions"+"."+strconv.Itoa(i), "body", m.Dimensions[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *Response) validateEntity(formats strfmt.Registry) error {
	if swag.IsZero(m.Entity) { // not required
		return nil
	}

	if m.Entity != nil {
		if err := m.Entity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entity")
			}
			return err
		}
	}

	return nil
}

func (m *Response) validateHeaders(formats strfmt.Registry) error {
	if swag.IsZero(m.Headers) { // not required
		return nil
	}

	for i := 0; i < len(m.Headers); i++ {
		if swag.IsZero(m.Headers[i]) { // not required
			continue
		}

		if m.Headers[i] != nil {
			if err := m.Headers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("headers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("headers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Response) validateLocationRef(formats strfmt.Registry) error {
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

func (m *Response) validateProxyChallengeRequests(formats strfmt.Registry) error {
	if swag.IsZero(m.ProxyChallengeRequests) { // not required
		return nil
	}

	for i := 0; i < len(m.ProxyChallengeRequests); i++ {
		if swag.IsZero(m.ProxyChallengeRequests[i]) { // not required
			continue
		}

		if m.ProxyChallengeRequests[i] != nil {
			if err := m.ProxyChallengeRequests[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("proxyChallengeRequests" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("proxyChallengeRequests" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Response) validateRecipientsInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.RecipientsInfo) { // not required
		return nil
	}

	for i := 0; i < len(m.RecipientsInfo); i++ {
		if swag.IsZero(m.RecipientsInfo[i]) { // not required
			continue
		}

		if m.RecipientsInfo[i] != nil {
			if err := m.RecipientsInfo[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("recipientsInfo" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("recipientsInfo" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Response) validateRequest(formats strfmt.Registry) error {
	if swag.IsZero(m.Request) { // not required
		return nil
	}

	if m.Request != nil {
		if err := m.Request.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("request")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("request")
			}
			return err
		}
	}

	return nil
}

func (m *Response) validateRetryAfter(formats strfmt.Registry) error {
	if swag.IsZero(m.RetryAfter) { // not required
		return nil
	}

	if err := validate.FormatOf("retryAfter", "body", "date-time", m.RetryAfter.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Response) validateServerInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.ServerInfo) { // not required
		return nil
	}

	if m.ServerInfo != nil {
		if err := m.ServerInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("serverInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("serverInfo")
			}
			return err
		}
	}

	return nil
}

func (m *Response) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *Response) validateWarnings(formats strfmt.Registry) error {
	if swag.IsZero(m.Warnings) { // not required
		return nil
	}

	for i := 0; i < len(m.Warnings); i++ {
		if swag.IsZero(m.Warnings[i]) { // not required
			continue
		}

		if m.Warnings[i] != nil {
			if err := m.Warnings[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("warnings" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("warnings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this response based on the context it is used
func (m *Response) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccessControlAllowMethods(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAllowedMethods(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAuthenticationInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCacheDirectives(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateChallengeRequests(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCookieSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHeaders(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLocationRef(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProxyChallengeRequests(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRecipientsInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRequest(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServerInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWarnings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Response) contextValidateAccessControlAllowMethods(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AccessControlAllowMethods); i++ {

		if m.AccessControlAllowMethods[i] != nil {

			if swag.IsZero(m.AccessControlAllowMethods[i]) { // not required
				return nil
			}

			if err := m.AccessControlAllowMethods[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("accessControlAllowMethods" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("accessControlAllowMethods" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Response) contextValidateAllowedMethods(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AllowedMethods); i++ {

		if m.AllowedMethods[i] != nil {

			if swag.IsZero(m.AllowedMethods[i]) { // not required
				return nil
			}

			if err := m.AllowedMethods[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("allowedMethods" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("allowedMethods" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Response) contextValidateAuthenticationInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.AuthenticationInfo != nil {

		if swag.IsZero(m.AuthenticationInfo) { // not required
			return nil
		}

		if err := m.AuthenticationInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("authenticationInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("authenticationInfo")
			}
			return err
		}
	}

	return nil
}

func (m *Response) contextValidateCacheDirectives(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CacheDirectives); i++ {

		if m.CacheDirectives[i] != nil {

			if swag.IsZero(m.CacheDirectives[i]) { // not required
				return nil
			}

			if err := m.CacheDirectives[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cacheDirectives" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("cacheDirectives" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Response) contextValidateChallengeRequests(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ChallengeRequests); i++ {

		if m.ChallengeRequests[i] != nil {

			if swag.IsZero(m.ChallengeRequests[i]) { // not required
				return nil
			}

			if err := m.ChallengeRequests[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("challengeRequests" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("challengeRequests" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Response) contextValidateCookieSettings(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CookieSettings); i++ {

		if m.CookieSettings[i] != nil {

			if swag.IsZero(m.CookieSettings[i]) { // not required
				return nil
			}

			if err := m.CookieSettings[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cookieSettings" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("cookieSettings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Response) contextValidateEntity(ctx context.Context, formats strfmt.Registry) error {

	if m.Entity != nil {

		if swag.IsZero(m.Entity) { // not required
			return nil
		}

		if err := m.Entity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entity")
			}
			return err
		}
	}

	return nil
}

func (m *Response) contextValidateHeaders(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Headers); i++ {

		if m.Headers[i] != nil {

			if swag.IsZero(m.Headers[i]) { // not required
				return nil
			}

			if err := m.Headers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("headers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("headers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Response) contextValidateLocationRef(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Response) contextValidateProxyChallengeRequests(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ProxyChallengeRequests); i++ {

		if m.ProxyChallengeRequests[i] != nil {

			if swag.IsZero(m.ProxyChallengeRequests[i]) { // not required
				return nil
			}

			if err := m.ProxyChallengeRequests[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("proxyChallengeRequests" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("proxyChallengeRequests" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Response) contextValidateRecipientsInfo(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RecipientsInfo); i++ {

		if m.RecipientsInfo[i] != nil {

			if swag.IsZero(m.RecipientsInfo[i]) { // not required
				return nil
			}

			if err := m.RecipientsInfo[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("recipientsInfo" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("recipientsInfo" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Response) contextValidateRequest(ctx context.Context, formats strfmt.Registry) error {

	if m.Request != nil {

		if swag.IsZero(m.Request) { // not required
			return nil
		}

		if err := m.Request.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("request")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("request")
			}
			return err
		}
	}

	return nil
}

func (m *Response) contextValidateServerInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.ServerInfo != nil {

		if swag.IsZero(m.ServerInfo) { // not required
			return nil
		}

		if err := m.ServerInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("serverInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("serverInfo")
			}
			return err
		}
	}

	return nil
}

func (m *Response) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {

		if swag.IsZero(m.Status) { // not required
			return nil
		}

		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *Response) contextValidateWarnings(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Warnings); i++ {

		if m.Warnings[i] != nil {

			if swag.IsZero(m.Warnings[i]) { // not required
				return nil
			}

			if err := m.Warnings[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("warnings" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("warnings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Response) UnmarshalBinary(b []byte) error {
	var res Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
