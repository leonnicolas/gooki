// Code generated by go-swagger; DO NOT EDIT.

package smartlock

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

// NewSmartlockBulkWebConfigResourcePostPostParams creates a new SmartlockBulkWebConfigResourcePostPostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSmartlockBulkWebConfigResourcePostPostParams() *SmartlockBulkWebConfigResourcePostPostParams {
	return &SmartlockBulkWebConfigResourcePostPostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSmartlockBulkWebConfigResourcePostPostParamsWithTimeout creates a new SmartlockBulkWebConfigResourcePostPostParams object
// with the ability to set a timeout on a request.
func NewSmartlockBulkWebConfigResourcePostPostParamsWithTimeout(timeout time.Duration) *SmartlockBulkWebConfigResourcePostPostParams {
	return &SmartlockBulkWebConfigResourcePostPostParams{
		timeout: timeout,
	}
}

// NewSmartlockBulkWebConfigResourcePostPostParamsWithContext creates a new SmartlockBulkWebConfigResourcePostPostParams object
// with the ability to set a context for a request.
func NewSmartlockBulkWebConfigResourcePostPostParamsWithContext(ctx context.Context) *SmartlockBulkWebConfigResourcePostPostParams {
	return &SmartlockBulkWebConfigResourcePostPostParams{
		Context: ctx,
	}
}

// NewSmartlockBulkWebConfigResourcePostPostParamsWithHTTPClient creates a new SmartlockBulkWebConfigResourcePostPostParams object
// with the ability to set a custom HTTPClient for a request.
func NewSmartlockBulkWebConfigResourcePostPostParamsWithHTTPClient(client *http.Client) *SmartlockBulkWebConfigResourcePostPostParams {
	return &SmartlockBulkWebConfigResourcePostPostParams{
		HTTPClient: client,
	}
}

/*
SmartlockBulkWebConfigResourcePostPostParams contains all the parameters to send to the API endpoint

	for the smartlock bulk web config resource post post operation.

	Typically these are written to a http.Request.
*/
type SmartlockBulkWebConfigResourcePostPostParams struct {

	/* Body.

	   Smartlocks web config update representation
	*/
	Body *models.BulkWebConfigRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the smartlock bulk web config resource post post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SmartlockBulkWebConfigResourcePostPostParams) WithDefaults() *SmartlockBulkWebConfigResourcePostPostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the smartlock bulk web config resource post post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SmartlockBulkWebConfigResourcePostPostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the smartlock bulk web config resource post post params
func (o *SmartlockBulkWebConfigResourcePostPostParams) WithTimeout(timeout time.Duration) *SmartlockBulkWebConfigResourcePostPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the smartlock bulk web config resource post post params
func (o *SmartlockBulkWebConfigResourcePostPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the smartlock bulk web config resource post post params
func (o *SmartlockBulkWebConfigResourcePostPostParams) WithContext(ctx context.Context) *SmartlockBulkWebConfigResourcePostPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the smartlock bulk web config resource post post params
func (o *SmartlockBulkWebConfigResourcePostPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the smartlock bulk web config resource post post params
func (o *SmartlockBulkWebConfigResourcePostPostParams) WithHTTPClient(client *http.Client) *SmartlockBulkWebConfigResourcePostPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the smartlock bulk web config resource post post params
func (o *SmartlockBulkWebConfigResourcePostPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the smartlock bulk web config resource post post params
func (o *SmartlockBulkWebConfigResourcePostPostParams) WithBody(body *models.BulkWebConfigRequest) *SmartlockBulkWebConfigResourcePostPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the smartlock bulk web config resource post post params
func (o *SmartlockBulkWebConfigResourcePostPostParams) SetBody(body *models.BulkWebConfigRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SmartlockBulkWebConfigResourcePostPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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