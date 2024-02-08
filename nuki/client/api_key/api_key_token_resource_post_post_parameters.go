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

// NewAPIKeyTokenResourcePostPostParams creates a new APIKeyTokenResourcePostPostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAPIKeyTokenResourcePostPostParams() *APIKeyTokenResourcePostPostParams {
	return &APIKeyTokenResourcePostPostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAPIKeyTokenResourcePostPostParamsWithTimeout creates a new APIKeyTokenResourcePostPostParams object
// with the ability to set a timeout on a request.
func NewAPIKeyTokenResourcePostPostParamsWithTimeout(timeout time.Duration) *APIKeyTokenResourcePostPostParams {
	return &APIKeyTokenResourcePostPostParams{
		timeout: timeout,
	}
}

// NewAPIKeyTokenResourcePostPostParamsWithContext creates a new APIKeyTokenResourcePostPostParams object
// with the ability to set a context for a request.
func NewAPIKeyTokenResourcePostPostParamsWithContext(ctx context.Context) *APIKeyTokenResourcePostPostParams {
	return &APIKeyTokenResourcePostPostParams{
		Context: ctx,
	}
}

// NewAPIKeyTokenResourcePostPostParamsWithHTTPClient creates a new APIKeyTokenResourcePostPostParams object
// with the ability to set a custom HTTPClient for a request.
func NewAPIKeyTokenResourcePostPostParamsWithHTTPClient(client *http.Client) *APIKeyTokenResourcePostPostParams {
	return &APIKeyTokenResourcePostPostParams{
		HTTPClient: client,
	}
}

/*
APIKeyTokenResourcePostPostParams contains all the parameters to send to the API endpoint

	for the Api key token resource post post operation.

	Typically these are written to a http.Request.
*/
type APIKeyTokenResourcePostPostParams struct {

	/* APIKeyID.

	   The api key id
	*/
	APIKeyID int64

	/* Body.

	   Api key token update representation
	*/
	Body *models.APIKeyTokenUpdate

	/* ID.

	   The api key token id
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the Api key token resource post post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *APIKeyTokenResourcePostPostParams) WithDefaults() *APIKeyTokenResourcePostPostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the Api key token resource post post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *APIKeyTokenResourcePostPostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the Api key token resource post post params
func (o *APIKeyTokenResourcePostPostParams) WithTimeout(timeout time.Duration) *APIKeyTokenResourcePostPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the Api key token resource post post params
func (o *APIKeyTokenResourcePostPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the Api key token resource post post params
func (o *APIKeyTokenResourcePostPostParams) WithContext(ctx context.Context) *APIKeyTokenResourcePostPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the Api key token resource post post params
func (o *APIKeyTokenResourcePostPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the Api key token resource post post params
func (o *APIKeyTokenResourcePostPostParams) WithHTTPClient(client *http.Client) *APIKeyTokenResourcePostPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the Api key token resource post post params
func (o *APIKeyTokenResourcePostPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIKeyID adds the aPIKeyID to the Api key token resource post post params
func (o *APIKeyTokenResourcePostPostParams) WithAPIKeyID(aPIKeyID int64) *APIKeyTokenResourcePostPostParams {
	o.SetAPIKeyID(aPIKeyID)
	return o
}

// SetAPIKeyID adds the apiKeyId to the Api key token resource post post params
func (o *APIKeyTokenResourcePostPostParams) SetAPIKeyID(aPIKeyID int64) {
	o.APIKeyID = aPIKeyID
}

// WithBody adds the body to the Api key token resource post post params
func (o *APIKeyTokenResourcePostPostParams) WithBody(body *models.APIKeyTokenUpdate) *APIKeyTokenResourcePostPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the Api key token resource post post params
func (o *APIKeyTokenResourcePostPostParams) SetBody(body *models.APIKeyTokenUpdate) {
	o.Body = body
}

// WithID adds the id to the Api key token resource post post params
func (o *APIKeyTokenResourcePostPostParams) WithID(id string) *APIKeyTokenResourcePostPostParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the Api key token resource post post params
func (o *APIKeyTokenResourcePostPostParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *APIKeyTokenResourcePostPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
