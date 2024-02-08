// Code generated by go-swagger; DO NOT EDIT.

package smartlock_log

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new smartlock log API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for smartlock log API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	SmartlockLogsResourceGetGet(params *SmartlockLogsResourceGetGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SmartlockLogsResourceGetGetOK, error)

	SmartlocksLogsResourceGetGet(params *SmartlocksLogsResourceGetGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SmartlocksLogsResourceGetGetOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
SmartlockLogsResourceGetGet gets a list of smartlock logs
*/
func (a *Client) SmartlockLogsResourceGetGet(params *SmartlockLogsResourceGetGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SmartlockLogsResourceGetGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSmartlockLogsResourceGetGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SmartlockLogsResource_get_get",
		Method:             "GET",
		PathPattern:        "/smartlock/{smartlockId}/log",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SmartlockLogsResourceGetGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SmartlockLogsResourceGetGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SmartlockLogsResource_get_get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SmartlocksLogsResourceGetGet gets a list of smartlock logs for all of your smartlocks
*/
func (a *Client) SmartlocksLogsResourceGetGet(params *SmartlocksLogsResourceGetGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SmartlocksLogsResourceGetGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSmartlocksLogsResourceGetGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SmartlocksLogsResource_get_get",
		Method:             "GET",
		PathPattern:        "/smartlock/log",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SmartlocksLogsResourceGetGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SmartlocksLogsResourceGetGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SmartlocksLogsResource_get_get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
