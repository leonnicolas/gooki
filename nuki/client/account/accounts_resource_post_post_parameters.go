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
	"github.com/go-openapi/swag"

	"github.com/leonnicolas/gooki/nuki/models"
)

// NewAccountsResourcePostPostParams creates a new AccountsResourcePostPostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAccountsResourcePostPostParams() *AccountsResourcePostPostParams {
	return &AccountsResourcePostPostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAccountsResourcePostPostParamsWithTimeout creates a new AccountsResourcePostPostParams object
// with the ability to set a timeout on a request.
func NewAccountsResourcePostPostParamsWithTimeout(timeout time.Duration) *AccountsResourcePostPostParams {
	return &AccountsResourcePostPostParams{
		timeout: timeout,
	}
}

// NewAccountsResourcePostPostParamsWithContext creates a new AccountsResourcePostPostParams object
// with the ability to set a context for a request.
func NewAccountsResourcePostPostParamsWithContext(ctx context.Context) *AccountsResourcePostPostParams {
	return &AccountsResourcePostPostParams{
		Context: ctx,
	}
}

// NewAccountsResourcePostPostParamsWithHTTPClient creates a new AccountsResourcePostPostParams object
// with the ability to set a custom HTTPClient for a request.
func NewAccountsResourcePostPostParamsWithHTTPClient(client *http.Client) *AccountsResourcePostPostParams {
	return &AccountsResourcePostPostParams{
		HTTPClient: client,
	}
}

/*
AccountsResourcePostPostParams contains all the parameters to send to the API endpoint

	for the accounts resource post post operation.

	Typically these are written to a http.Request.
*/
type AccountsResourcePostPostParams struct {

	/* Body.

	   Account update representation
	*/
	Body *models.AccountUpdate

	/* DeleteAPITokens.

	   If false existing API tokens are not deleted if the password is changed

	   Default: true
	*/
	DeleteAPITokens *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the accounts resource post post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AccountsResourcePostPostParams) WithDefaults() *AccountsResourcePostPostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the accounts resource post post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AccountsResourcePostPostParams) SetDefaults() {
	var (
		deleteAPITokensDefault = bool(true)
	)

	val := AccountsResourcePostPostParams{
		DeleteAPITokens: &deleteAPITokensDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the accounts resource post post params
func (o *AccountsResourcePostPostParams) WithTimeout(timeout time.Duration) *AccountsResourcePostPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the accounts resource post post params
func (o *AccountsResourcePostPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the accounts resource post post params
func (o *AccountsResourcePostPostParams) WithContext(ctx context.Context) *AccountsResourcePostPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the accounts resource post post params
func (o *AccountsResourcePostPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the accounts resource post post params
func (o *AccountsResourcePostPostParams) WithHTTPClient(client *http.Client) *AccountsResourcePostPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the accounts resource post post params
func (o *AccountsResourcePostPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the accounts resource post post params
func (o *AccountsResourcePostPostParams) WithBody(body *models.AccountUpdate) *AccountsResourcePostPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the accounts resource post post params
func (o *AccountsResourcePostPostParams) SetBody(body *models.AccountUpdate) {
	o.Body = body
}

// WithDeleteAPITokens adds the deleteAPITokens to the accounts resource post post params
func (o *AccountsResourcePostPostParams) WithDeleteAPITokens(deleteAPITokens *bool) *AccountsResourcePostPostParams {
	o.SetDeleteAPITokens(deleteAPITokens)
	return o
}

// SetDeleteAPITokens adds the deleteApiTokens to the accounts resource post post params
func (o *AccountsResourcePostPostParams) SetDeleteAPITokens(deleteAPITokens *bool) {
	o.DeleteAPITokens = deleteAPITokens
}

// WriteToRequest writes these params to a swagger request
func (o *AccountsResourcePostPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.DeleteAPITokens != nil {

		// query param deleteApiTokens
		var qrDeleteAPITokens bool

		if o.DeleteAPITokens != nil {
			qrDeleteAPITokens = *o.DeleteAPITokens
		}
		qDeleteAPITokens := swag.FormatBool(qrDeleteAPITokens)
		if qDeleteAPITokens != "" {

			if err := r.SetQueryParam("deleteApiTokens", qDeleteAPITokens); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}