// Code generated by go-swagger; DO NOT EDIT.

package api_key

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

// NewAPIKeyTokensResourcePutPutParams creates a new APIKeyTokensResourcePutPutParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAPIKeyTokensResourcePutPutParams() *APIKeyTokensResourcePutPutParams {
	return &APIKeyTokensResourcePutPutParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAPIKeyTokensResourcePutPutParamsWithTimeout creates a new APIKeyTokensResourcePutPutParams object
// with the ability to set a timeout on a request.
func NewAPIKeyTokensResourcePutPutParamsWithTimeout(timeout time.Duration) *APIKeyTokensResourcePutPutParams {
	return &APIKeyTokensResourcePutPutParams{
		timeout: timeout,
	}
}

// NewAPIKeyTokensResourcePutPutParamsWithContext creates a new APIKeyTokensResourcePutPutParams object
// with the ability to set a context for a request.
func NewAPIKeyTokensResourcePutPutParamsWithContext(ctx context.Context) *APIKeyTokensResourcePutPutParams {
	return &APIKeyTokensResourcePutPutParams{
		Context: ctx,
	}
}

// NewAPIKeyTokensResourcePutPutParamsWithHTTPClient creates a new APIKeyTokensResourcePutPutParams object
// with the ability to set a custom HTTPClient for a request.
func NewAPIKeyTokensResourcePutPutParamsWithHTTPClient(client *http.Client) *APIKeyTokensResourcePutPutParams {
	return &APIKeyTokensResourcePutPutParams{
		HTTPClient: client,
	}
}

/*
APIKeyTokensResourcePutPutParams contains all the parameters to send to the API endpoint

	for the Api key tokens resource put put operation.

	Typically these are written to a http.Request.
*/
type APIKeyTokensResourcePutPutParams struct {

	/* APIKeyID.

	   The api key id
	*/
	APIKeyID int64

	/* Body.

	   Api key token create representation
	*/
	Body *models.APIKeyTokenCreate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the Api key tokens resource put put params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *APIKeyTokensResourcePutPutParams) WithDefaults() *APIKeyTokensResourcePutPutParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the Api key tokens resource put put params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *APIKeyTokensResourcePutPutParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the Api key tokens resource put put params
func (o *APIKeyTokensResourcePutPutParams) WithTimeout(timeout time.Duration) *APIKeyTokensResourcePutPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the Api key tokens resource put put params
func (o *APIKeyTokensResourcePutPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the Api key tokens resource put put params
func (o *APIKeyTokensResourcePutPutParams) WithContext(ctx context.Context) *APIKeyTokensResourcePutPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the Api key tokens resource put put params
func (o *APIKeyTokensResourcePutPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the Api key tokens resource put put params
func (o *APIKeyTokensResourcePutPutParams) WithHTTPClient(client *http.Client) *APIKeyTokensResourcePutPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the Api key tokens resource put put params
func (o *APIKeyTokensResourcePutPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIKeyID adds the aPIKeyID to the Api key tokens resource put put params
func (o *APIKeyTokensResourcePutPutParams) WithAPIKeyID(aPIKeyID int64) *APIKeyTokensResourcePutPutParams {
	o.SetAPIKeyID(aPIKeyID)
	return o
}

// SetAPIKeyID adds the apiKeyId to the Api key tokens resource put put params
func (o *APIKeyTokensResourcePutPutParams) SetAPIKeyID(aPIKeyID int64) {
	o.APIKeyID = aPIKeyID
}

// WithBody adds the body to the Api key tokens resource put put params
func (o *APIKeyTokensResourcePutPutParams) WithBody(body *models.APIKeyTokenCreate) *APIKeyTokensResourcePutPutParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the Api key tokens resource put put params
func (o *APIKeyTokensResourcePutPutParams) SetBody(body *models.APIKeyTokenCreate) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *APIKeyTokensResourcePutPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param apiKeyId
	if err := r.SetPathParam("apiKeyId", swag.FormatInt64(o.APIKeyID)); err != nil {
		return err
	}
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
