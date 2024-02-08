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
)

// NewAccountIntegrationsResourceDeleteDeleteParams creates a new AccountIntegrationsResourceDeleteDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAccountIntegrationsResourceDeleteDeleteParams() *AccountIntegrationsResourceDeleteDeleteParams {
	return &AccountIntegrationsResourceDeleteDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAccountIntegrationsResourceDeleteDeleteParamsWithTimeout creates a new AccountIntegrationsResourceDeleteDeleteParams object
// with the ability to set a timeout on a request.
func NewAccountIntegrationsResourceDeleteDeleteParamsWithTimeout(timeout time.Duration) *AccountIntegrationsResourceDeleteDeleteParams {
	return &AccountIntegrationsResourceDeleteDeleteParams{
		timeout: timeout,
	}
}

// NewAccountIntegrationsResourceDeleteDeleteParamsWithContext creates a new AccountIntegrationsResourceDeleteDeleteParams object
// with the ability to set a context for a request.
func NewAccountIntegrationsResourceDeleteDeleteParamsWithContext(ctx context.Context) *AccountIntegrationsResourceDeleteDeleteParams {
	return &AccountIntegrationsResourceDeleteDeleteParams{
		Context: ctx,
	}
}

// NewAccountIntegrationsResourceDeleteDeleteParamsWithHTTPClient creates a new AccountIntegrationsResourceDeleteDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewAccountIntegrationsResourceDeleteDeleteParamsWithHTTPClient(client *http.Client) *AccountIntegrationsResourceDeleteDeleteParams {
	return &AccountIntegrationsResourceDeleteDeleteParams{
		HTTPClient: client,
	}
}

/*
AccountIntegrationsResourceDeleteDeleteParams contains all the parameters to send to the API endpoint

	for the account integrations resource delete delete operation.

	Typically these are written to a http.Request.
*/
type AccountIntegrationsResourceDeleteDeleteParams struct {

	/* APIKeyID.

	   The api key id to delete (this also removes all tokens if no specific tokenId is given)
	*/
	APIKeyID *int64

	/* TokenID.

	   The token id if a specific token has to be deleted but not the full api key
	*/
	TokenID *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the account integrations resource delete delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AccountIntegrationsResourceDeleteDeleteParams) WithDefaults() *AccountIntegrationsResourceDeleteDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the account integrations resource delete delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AccountIntegrationsResourceDeleteDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the account integrations resource delete delete params
func (o *AccountIntegrationsResourceDeleteDeleteParams) WithTimeout(timeout time.Duration) *AccountIntegrationsResourceDeleteDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the account integrations resource delete delete params
func (o *AccountIntegrationsResourceDeleteDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the account integrations resource delete delete params
func (o *AccountIntegrationsResourceDeleteDeleteParams) WithContext(ctx context.Context) *AccountIntegrationsResourceDeleteDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the account integrations resource delete delete params
func (o *AccountIntegrationsResourceDeleteDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the account integrations resource delete delete params
func (o *AccountIntegrationsResourceDeleteDeleteParams) WithHTTPClient(client *http.Client) *AccountIntegrationsResourceDeleteDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the account integrations resource delete delete params
func (o *AccountIntegrationsResourceDeleteDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIKeyID adds the aPIKeyID to the account integrations resource delete delete params
func (o *AccountIntegrationsResourceDeleteDeleteParams) WithAPIKeyID(aPIKeyID *int64) *AccountIntegrationsResourceDeleteDeleteParams {
	o.SetAPIKeyID(aPIKeyID)
	return o
}

// SetAPIKeyID adds the apiKeyId to the account integrations resource delete delete params
func (o *AccountIntegrationsResourceDeleteDeleteParams) SetAPIKeyID(aPIKeyID *int64) {
	o.APIKeyID = aPIKeyID
}

// WithTokenID adds the tokenID to the account integrations resource delete delete params
func (o *AccountIntegrationsResourceDeleteDeleteParams) WithTokenID(tokenID *int64) *AccountIntegrationsResourceDeleteDeleteParams {
	o.SetTokenID(tokenID)
	return o
}

// SetTokenID adds the tokenId to the account integrations resource delete delete params
func (o *AccountIntegrationsResourceDeleteDeleteParams) SetTokenID(tokenID *int64) {
	o.TokenID = tokenID
}

// WriteToRequest writes these params to a swagger request
func (o *AccountIntegrationsResourceDeleteDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.APIKeyID != nil {

		// query param apiKeyId
		var qrAPIKeyID int64

		if o.APIKeyID != nil {
			qrAPIKeyID = *o.APIKeyID
		}
		qAPIKeyID := swag.FormatInt64(qrAPIKeyID)
		if qAPIKeyID != "" {

			if err := r.SetQueryParam("apiKeyId", qAPIKeyID); err != nil {
				return err
			}
		}
	}

	if o.TokenID != nil {

		// query param tokenId
		var qrTokenID int64

		if o.TokenID != nil {
			qrTokenID = *o.TokenID
		}
		qTokenID := swag.FormatInt64(qrTokenID)
		if qTokenID != "" {

			if err := r.SetQueryParam("tokenId", qTokenID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}