// Code generated by go-swagger; DO NOT EDIT.

package opener

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

// NewOpenerBrandResourceGetGetParams creates a new OpenerBrandResourceGetGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOpenerBrandResourceGetGetParams() *OpenerBrandResourceGetGetParams {
	return &OpenerBrandResourceGetGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOpenerBrandResourceGetGetParamsWithTimeout creates a new OpenerBrandResourceGetGetParams object
// with the ability to set a timeout on a request.
func NewOpenerBrandResourceGetGetParamsWithTimeout(timeout time.Duration) *OpenerBrandResourceGetGetParams {
	return &OpenerBrandResourceGetGetParams{
		timeout: timeout,
	}
}

// NewOpenerBrandResourceGetGetParamsWithContext creates a new OpenerBrandResourceGetGetParams object
// with the ability to set a context for a request.
func NewOpenerBrandResourceGetGetParamsWithContext(ctx context.Context) *OpenerBrandResourceGetGetParams {
	return &OpenerBrandResourceGetGetParams{
		Context: ctx,
	}
}

// NewOpenerBrandResourceGetGetParamsWithHTTPClient creates a new OpenerBrandResourceGetGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewOpenerBrandResourceGetGetParamsWithHTTPClient(client *http.Client) *OpenerBrandResourceGetGetParams {
	return &OpenerBrandResourceGetGetParams{
		HTTPClient: client,
	}
}

/*
OpenerBrandResourceGetGetParams contains all the parameters to send to the API endpoint

	for the opener brand resource get get operation.

	Typically these are written to a http.Request.
*/
type OpenerBrandResourceGetGetParams struct {

	/* BrandID.

	   The brand ID
	*/
	BrandID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the opener brand resource get get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OpenerBrandResourceGetGetParams) WithDefaults() *OpenerBrandResourceGetGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the opener brand resource get get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OpenerBrandResourceGetGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the opener brand resource get get params
func (o *OpenerBrandResourceGetGetParams) WithTimeout(timeout time.Duration) *OpenerBrandResourceGetGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the opener brand resource get get params
func (o *OpenerBrandResourceGetGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the opener brand resource get get params
func (o *OpenerBrandResourceGetGetParams) WithContext(ctx context.Context) *OpenerBrandResourceGetGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the opener brand resource get get params
func (o *OpenerBrandResourceGetGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the opener brand resource get get params
func (o *OpenerBrandResourceGetGetParams) WithHTTPClient(client *http.Client) *OpenerBrandResourceGetGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the opener brand resource get get params
func (o *OpenerBrandResourceGetGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBrandID adds the brandID to the opener brand resource get get params
func (o *OpenerBrandResourceGetGetParams) WithBrandID(brandID int64) *OpenerBrandResourceGetGetParams {
	o.SetBrandID(brandID)
	return o
}

// SetBrandID adds the brandId to the opener brand resource get get params
func (o *OpenerBrandResourceGetGetParams) SetBrandID(brandID int64) {
	o.BrandID = brandID
}

// WriteToRequest writes these params to a swagger request
func (o *OpenerBrandResourceGetGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param brandId
	if err := r.SetPathParam("brandId", swag.FormatInt64(o.BrandID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
