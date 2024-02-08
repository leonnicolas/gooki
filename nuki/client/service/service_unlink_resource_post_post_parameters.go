// Code generated by go-swagger; DO NOT EDIT.

package service

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
)

// NewServiceUnlinkResourcePostPostParams creates a new ServiceUnlinkResourcePostPostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewServiceUnlinkResourcePostPostParams() *ServiceUnlinkResourcePostPostParams {
	return &ServiceUnlinkResourcePostPostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewServiceUnlinkResourcePostPostParamsWithTimeout creates a new ServiceUnlinkResourcePostPostParams object
// with the ability to set a timeout on a request.
func NewServiceUnlinkResourcePostPostParamsWithTimeout(timeout time.Duration) *ServiceUnlinkResourcePostPostParams {
	return &ServiceUnlinkResourcePostPostParams{
		timeout: timeout,
	}
}

// NewServiceUnlinkResourcePostPostParamsWithContext creates a new ServiceUnlinkResourcePostPostParams object
// with the ability to set a context for a request.
func NewServiceUnlinkResourcePostPostParamsWithContext(ctx context.Context) *ServiceUnlinkResourcePostPostParams {
	return &ServiceUnlinkResourcePostPostParams{
		Context: ctx,
	}
}

// NewServiceUnlinkResourcePostPostParamsWithHTTPClient creates a new ServiceUnlinkResourcePostPostParams object
// with the ability to set a custom HTTPClient for a request.
func NewServiceUnlinkResourcePostPostParamsWithHTTPClient(client *http.Client) *ServiceUnlinkResourcePostPostParams {
	return &ServiceUnlinkResourcePostPostParams{
		HTTPClient: client,
	}
}

/*
ServiceUnlinkResourcePostPostParams contains all the parameters to send to the API endpoint

	for the service unlink resource post post operation.

	Typically these are written to a http.Request.
*/
type ServiceUnlinkResourcePostPostParams struct {

	/* ServiceID.

	   The service id
	*/
	ServiceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the service unlink resource post post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServiceUnlinkResourcePostPostParams) WithDefaults() *ServiceUnlinkResourcePostPostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the service unlink resource post post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServiceUnlinkResourcePostPostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the service unlink resource post post params
func (o *ServiceUnlinkResourcePostPostParams) WithTimeout(timeout time.Duration) *ServiceUnlinkResourcePostPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the service unlink resource post post params
func (o *ServiceUnlinkResourcePostPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the service unlink resource post post params
func (o *ServiceUnlinkResourcePostPostParams) WithContext(ctx context.Context) *ServiceUnlinkResourcePostPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the service unlink resource post post params
func (o *ServiceUnlinkResourcePostPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the service unlink resource post post params
func (o *ServiceUnlinkResourcePostPostParams) WithHTTPClient(client *http.Client) *ServiceUnlinkResourcePostPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the service unlink resource post post params
func (o *ServiceUnlinkResourcePostPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithServiceID adds the serviceID to the service unlink resource post post params
func (o *ServiceUnlinkResourcePostPostParams) WithServiceID(serviceID string) *ServiceUnlinkResourcePostPostParams {
	o.SetServiceID(serviceID)
	return o
}

// SetServiceID adds the serviceId to the service unlink resource post post params
func (o *ServiceUnlinkResourcePostPostParams) SetServiceID(serviceID string) {
	o.ServiceID = serviceID
}

// WriteToRequest writes these params to a swagger request
func (o *ServiceUnlinkResourcePostPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param serviceId
	if err := r.SetPathParam("serviceId", o.ServiceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
