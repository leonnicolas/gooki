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
)

// NewAPIKeyAdvancedReactivateResourcePostPostParams creates a new APIKeyAdvancedReactivateResourcePostPostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAPIKeyAdvancedReactivateResourcePostPostParams() *APIKeyAdvancedReactivateResourcePostPostParams {
	return &APIKeyAdvancedReactivateResourcePostPostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAPIKeyAdvancedReactivateResourcePostPostParamsWithTimeout creates a new APIKeyAdvancedReactivateResourcePostPostParams object
// with the ability to set a timeout on a request.
func NewAPIKeyAdvancedReactivateResourcePostPostParamsWithTimeout(timeout time.Duration) *APIKeyAdvancedReactivateResourcePostPostParams {
	return &APIKeyAdvancedReactivateResourcePostPostParams{
		timeout: timeout,
	}
}

// NewAPIKeyAdvancedReactivateResourcePostPostParamsWithContext creates a new APIKeyAdvancedReactivateResourcePostPostParams object
// with the ability to set a context for a request.
func NewAPIKeyAdvancedReactivateResourcePostPostParamsWithContext(ctx context.Context) *APIKeyAdvancedReactivateResourcePostPostParams {
	return &APIKeyAdvancedReactivateResourcePostPostParams{
		Context: ctx,
	}
}

// NewAPIKeyAdvancedReactivateResourcePostPostParamsWithHTTPClient creates a new APIKeyAdvancedReactivateResourcePostPostParams object
// with the ability to set a custom HTTPClient for a request.
func NewAPIKeyAdvancedReactivateResourcePostPostParamsWithHTTPClient(client *http.Client) *APIKeyAdvancedReactivateResourcePostPostParams {
	return &APIKeyAdvancedReactivateResourcePostPostParams{
		HTTPClient: client,
	}
}

/*
APIKeyAdvancedReactivateResourcePostPostParams contains all the parameters to send to the API endpoint

	for the Api key advanced reactivate resource post post operation.

	Typically these are written to a http.Request.
*/
type APIKeyAdvancedReactivateResourcePostPostParams struct {

	/* APIKeyID.

	   The api key id
	*/
	APIKeyID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the Api key advanced reactivate resource post post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *APIKeyAdvancedReactivateResourcePostPostParams) WithDefaults() *APIKeyAdvancedReactivateResourcePostPostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the Api key advanced reactivate resource post post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *APIKeyAdvancedReactivateResourcePostPostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the Api key advanced reactivate resource post post params
func (o *APIKeyAdvancedReactivateResourcePostPostParams) WithTimeout(timeout time.Duration) *APIKeyAdvancedReactivateResourcePostPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the Api key advanced reactivate resource post post params
func (o *APIKeyAdvancedReactivateResourcePostPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the Api key advanced reactivate resource post post params
func (o *APIKeyAdvancedReactivateResourcePostPostParams) WithContext(ctx context.Context) *APIKeyAdvancedReactivateResourcePostPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the Api key advanced reactivate resource post post params
func (o *APIKeyAdvancedReactivateResourcePostPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the Api key advanced reactivate resource post post params
func (o *APIKeyAdvancedReactivateResourcePostPostParams) WithHTTPClient(client *http.Client) *APIKeyAdvancedReactivateResourcePostPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the Api key advanced reactivate resource post post params
func (o *APIKeyAdvancedReactivateResourcePostPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIKeyID adds the aPIKeyID to the Api key advanced reactivate resource post post params
func (o *APIKeyAdvancedReactivateResourcePostPostParams) WithAPIKeyID(aPIKeyID int64) *APIKeyAdvancedReactivateResourcePostPostParams {
	o.SetAPIKeyID(aPIKeyID)
	return o
}

// SetAPIKeyID adds the apiKeyId to the Api key advanced reactivate resource post post params
func (o *APIKeyAdvancedReactivateResourcePostPostParams) SetAPIKeyID(aPIKeyID int64) {
	o.APIKeyID = aPIKeyID
}

// WriteToRequest writes these params to a swagger request
func (o *APIKeyAdvancedReactivateResourcePostPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param apiKeyId
	if err := r.SetPathParam("apiKeyId", swag.FormatInt64(o.APIKeyID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
