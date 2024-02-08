// Code generated by go-swagger; DO NOT EDIT.

package smartlock_auth

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

// NewSmartlocksAuthsResourceGetGetParams creates a new SmartlocksAuthsResourceGetGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSmartlocksAuthsResourceGetGetParams() *SmartlocksAuthsResourceGetGetParams {
	return &SmartlocksAuthsResourceGetGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSmartlocksAuthsResourceGetGetParamsWithTimeout creates a new SmartlocksAuthsResourceGetGetParams object
// with the ability to set a timeout on a request.
func NewSmartlocksAuthsResourceGetGetParamsWithTimeout(timeout time.Duration) *SmartlocksAuthsResourceGetGetParams {
	return &SmartlocksAuthsResourceGetGetParams{
		timeout: timeout,
	}
}

// NewSmartlocksAuthsResourceGetGetParamsWithContext creates a new SmartlocksAuthsResourceGetGetParams object
// with the ability to set a context for a request.
func NewSmartlocksAuthsResourceGetGetParamsWithContext(ctx context.Context) *SmartlocksAuthsResourceGetGetParams {
	return &SmartlocksAuthsResourceGetGetParams{
		Context: ctx,
	}
}

// NewSmartlocksAuthsResourceGetGetParamsWithHTTPClient creates a new SmartlocksAuthsResourceGetGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewSmartlocksAuthsResourceGetGetParamsWithHTTPClient(client *http.Client) *SmartlocksAuthsResourceGetGetParams {
	return &SmartlocksAuthsResourceGetGetParams{
		HTTPClient: client,
	}
}

/*
SmartlocksAuthsResourceGetGetParams contains all the parameters to send to the API endpoint

	for the smartlocks auths resource get get operation.

	Typically these are written to a http.Request.
*/
type SmartlocksAuthsResourceGetGetParams struct {

	/* AccountUserID.

	   Filter for account users:  set to a positive number will filter for authorizations with this specific accountUserId, set to a negative number will filter without set accountUserId
	*/
	AccountUserID *int64

	/* Types.

	   Filter for authorization's types (comma-separated eg: 0,2,3)
	*/
	Types *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the smartlocks auths resource get get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SmartlocksAuthsResourceGetGetParams) WithDefaults() *SmartlocksAuthsResourceGetGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the smartlocks auths resource get get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SmartlocksAuthsResourceGetGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the smartlocks auths resource get get params
func (o *SmartlocksAuthsResourceGetGetParams) WithTimeout(timeout time.Duration) *SmartlocksAuthsResourceGetGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the smartlocks auths resource get get params
func (o *SmartlocksAuthsResourceGetGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the smartlocks auths resource get get params
func (o *SmartlocksAuthsResourceGetGetParams) WithContext(ctx context.Context) *SmartlocksAuthsResourceGetGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the smartlocks auths resource get get params
func (o *SmartlocksAuthsResourceGetGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the smartlocks auths resource get get params
func (o *SmartlocksAuthsResourceGetGetParams) WithHTTPClient(client *http.Client) *SmartlocksAuthsResourceGetGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the smartlocks auths resource get get params
func (o *SmartlocksAuthsResourceGetGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountUserID adds the accountUserID to the smartlocks auths resource get get params
func (o *SmartlocksAuthsResourceGetGetParams) WithAccountUserID(accountUserID *int64) *SmartlocksAuthsResourceGetGetParams {
	o.SetAccountUserID(accountUserID)
	return o
}

// SetAccountUserID adds the accountUserId to the smartlocks auths resource get get params
func (o *SmartlocksAuthsResourceGetGetParams) SetAccountUserID(accountUserID *int64) {
	o.AccountUserID = accountUserID
}

// WithTypes adds the types to the smartlocks auths resource get get params
func (o *SmartlocksAuthsResourceGetGetParams) WithTypes(types *string) *SmartlocksAuthsResourceGetGetParams {
	o.SetTypes(types)
	return o
}

// SetTypes adds the types to the smartlocks auths resource get get params
func (o *SmartlocksAuthsResourceGetGetParams) SetTypes(types *string) {
	o.Types = types
}

// WriteToRequest writes these params to a swagger request
func (o *SmartlocksAuthsResourceGetGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AccountUserID != nil {

		// query param accountUserId
		var qrAccountUserID int64

		if o.AccountUserID != nil {
			qrAccountUserID = *o.AccountUserID
		}
		qAccountUserID := swag.FormatInt64(qrAccountUserID)
		if qAccountUserID != "" {

			if err := r.SetQueryParam("accountUserId", qAccountUserID); err != nil {
				return err
			}
		}
	}

	if o.Types != nil {

		// query param types
		var qrTypes string

		if o.Types != nil {
			qrTypes = *o.Types
		}
		qTypes := qrTypes
		if qTypes != "" {

			if err := r.SetQueryParam("types", qTypes); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
