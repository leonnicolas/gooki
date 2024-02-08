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

// NewAPIKeyAdvancedResourcePostPostParams creates a new APIKeyAdvancedResourcePostPostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAPIKeyAdvancedResourcePostPostParams() *APIKeyAdvancedResourcePostPostParams {
	return &APIKeyAdvancedResourcePostPostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAPIKeyAdvancedResourcePostPostParamsWithTimeout creates a new APIKeyAdvancedResourcePostPostParams object
// with the ability to set a timeout on a request.
func NewAPIKeyAdvancedResourcePostPostParamsWithTimeout(timeout time.Duration) *APIKeyAdvancedResourcePostPostParams {
	return &APIKeyAdvancedResourcePostPostParams{
		timeout: timeout,
	}
}

// NewAPIKeyAdvancedResourcePostPostParamsWithContext creates a new APIKeyAdvancedResourcePostPostParams object
// with the ability to set a context for a request.
func NewAPIKeyAdvancedResourcePostPostParamsWithContext(ctx context.Context) *APIKeyAdvancedResourcePostPostParams {
	return &APIKeyAdvancedResourcePostPostParams{
		Context: ctx,
	}
}

// NewAPIKeyAdvancedResourcePostPostParamsWithHTTPClient creates a new APIKeyAdvancedResourcePostPostParams object
// with the ability to set a custom HTTPClient for a request.
func NewAPIKeyAdvancedResourcePostPostParamsWithHTTPClient(client *http.Client) *APIKeyAdvancedResourcePostPostParams {
	return &APIKeyAdvancedResourcePostPostParams{
		HTTPClient: client,
	}
}

/*
APIKeyAdvancedResourcePostPostParams contains all the parameters to send to the API endpoint

	for the Api key advanced resource post post operation.

	Typically these are written to a http.Request.
*/
type APIKeyAdvancedResourcePostPostParams struct {

	/* APIKeyID.

	   The api key id
	*/
	APIKeyID int64

	/* Body.

	   Update for advaced api key representation
	*/
	Body *models.AdvancedAPIKeyUpdate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the Api key advanced resource post post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *APIKeyAdvancedResourcePostPostParams) WithDefaults() *APIKeyAdvancedResourcePostPostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the Api key advanced resource post post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *APIKeyAdvancedResourcePostPostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the Api key advanced resource post post params
func (o *APIKeyAdvancedResourcePostPostParams) WithTimeout(timeout time.Duration) *APIKeyAdvancedResourcePostPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the Api key advanced resource post post params
func (o *APIKeyAdvancedResourcePostPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the Api key advanced resource post post params
func (o *APIKeyAdvancedResourcePostPostParams) WithContext(ctx context.Context) *APIKeyAdvancedResourcePostPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the Api key advanced resource post post params
func (o *APIKeyAdvancedResourcePostPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the Api key advanced resource post post params
func (o *APIKeyAdvancedResourcePostPostParams) WithHTTPClient(client *http.Client) *APIKeyAdvancedResourcePostPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the Api key advanced resource post post params
func (o *APIKeyAdvancedResourcePostPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIKeyID adds the aPIKeyID to the Api key advanced resource post post params
func (o *APIKeyAdvancedResourcePostPostParams) WithAPIKeyID(aPIKeyID int64) *APIKeyAdvancedResourcePostPostParams {
	o.SetAPIKeyID(aPIKeyID)
	return o
}

// SetAPIKeyID adds the apiKeyId to the Api key advanced resource post post params
func (o *APIKeyAdvancedResourcePostPostParams) SetAPIKeyID(aPIKeyID int64) {
	o.APIKeyID = aPIKeyID
}

// WithBody adds the body to the Api key advanced resource post post params
func (o *APIKeyAdvancedResourcePostPostParams) WithBody(body *models.AdvancedAPIKeyUpdate) *APIKeyAdvancedResourcePostPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the Api key advanced resource post post params
func (o *APIKeyAdvancedResourcePostPostParams) SetBody(body *models.AdvancedAPIKeyUpdate) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *APIKeyAdvancedResourcePostPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
