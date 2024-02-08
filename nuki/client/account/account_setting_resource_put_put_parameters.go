// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/leonnicolas/gooki/nuki/models"
)

// NewAccountSettingResourcePutPutParams creates a new AccountSettingResourcePutPutParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAccountSettingResourcePutPutParams() *AccountSettingResourcePutPutParams {
	return &AccountSettingResourcePutPutParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAccountSettingResourcePutPutParamsWithTimeout creates a new AccountSettingResourcePutPutParams object
// with the ability to set a timeout on a request.
func NewAccountSettingResourcePutPutParamsWithTimeout(timeout time.Duration) *AccountSettingResourcePutPutParams {
	return &AccountSettingResourcePutPutParams{
		timeout: timeout,
	}
}

// NewAccountSettingResourcePutPutParamsWithContext creates a new AccountSettingResourcePutPutParams object
// with the ability to set a context for a request.
func NewAccountSettingResourcePutPutParamsWithContext(ctx context.Context) *AccountSettingResourcePutPutParams {
	return &AccountSettingResourcePutPutParams{
		Context: ctx,
	}
}

// NewAccountSettingResourcePutPutParamsWithHTTPClient creates a new AccountSettingResourcePutPutParams object
// with the ability to set a custom HTTPClient for a request.
func NewAccountSettingResourcePutPutParamsWithHTTPClient(client *http.Client) *AccountSettingResourcePutPutParams {
	return &AccountSettingResourcePutPutParams{
		HTTPClient: client,
	}
}

/*
AccountSettingResourcePutPutParams contains all the parameters to send to the API endpoint

	for the account setting resource put put operation.

	Typically these are written to a http.Request.
*/
type AccountSettingResourcePutPutParams struct {

	/* Body.

	   Account setting representation
	*/
	Body *models.AccountSetting

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the account setting resource put put params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AccountSettingResourcePutPutParams) WithDefaults() *AccountSettingResourcePutPutParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the account setting resource put put params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AccountSettingResourcePutPutParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the account setting resource put put params
func (o *AccountSettingResourcePutPutParams) WithTimeout(timeout time.Duration) *AccountSettingResourcePutPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the account setting resource put put params
func (o *AccountSettingResourcePutPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the account setting resource put put params
func (o *AccountSettingResourcePutPutParams) WithContext(ctx context.Context) *AccountSettingResourcePutPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the account setting resource put put params
func (o *AccountSettingResourcePutPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the account setting resource put put params
func (o *AccountSettingResourcePutPutParams) WithHTTPClient(client *http.Client) *AccountSettingResourcePutPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the account setting resource put put params
func (o *AccountSettingResourcePutPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the account setting resource put put params
func (o *AccountSettingResourcePutPutParams) WithBody(body *models.AccountSetting) *AccountSettingResourcePutPutParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the account setting resource put put params
func (o *AccountSettingResourcePutPutParams) SetBody(body *models.AccountSetting) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AccountSettingResourcePutPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
