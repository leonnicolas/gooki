// Code generated by go-swagger; DO NOT EDIT.

package advanced_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new advanced api API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for advanced api API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DecentralWebhookResourceDeleteDelete(params *DecentralWebhookResourceDeleteDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DecentralWebhookResourceDeleteDeleteNoContent, error)

	DecentralWebhooksResourceGetGet(params *DecentralWebhooksResourceGetGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DecentralWebhooksResourceGetGetOK, error)

	DecentralWebhooksResourcePutPut(params *DecentralWebhooksResourcePutPutParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DecentralWebhooksResourcePutPutOK, error)

	SmartlockActionAdvancedResourceActionPost(params *SmartlockActionAdvancedResourceActionPostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SmartlockActionAdvancedResourceActionPostOK, error)

	SmartlockAuthsAdvancedResourcePutPut(params *SmartlockAuthsAdvancedResourcePutPutParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SmartlockAuthsAdvancedResourcePutPutOK, error)

	SmartlockLockActionAdvancedResourcePostLockPost(params *SmartlockLockActionAdvancedResourcePostLockPostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SmartlockLockActionAdvancedResourcePostLockPostOK, error)

	SmartlockUnlockActionAdvancedResourcePostLockPost(params *SmartlockUnlockActionAdvancedResourcePostLockPostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SmartlockUnlockActionAdvancedResourcePostLockPostOK, error)

	WebhookLogsResourceGetGet(params *WebhookLogsResourceGetGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*WebhookLogsResourceGetGetOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DecentralWebhookResourceDeleteDelete unregisters a decentral webhook
*/
func (a *Client) DecentralWebhookResourceDeleteDelete(params *DecentralWebhookResourceDeleteDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DecentralWebhookResourceDeleteDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDecentralWebhookResourceDeleteDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DecentralWebhookResource_delete_delete",
		Method:             "DELETE",
		PathPattern:        "/api/decentralWebhook/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DecentralWebhookResourceDeleteDeleteReader{formats: a.formats},
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
	success, ok := result.(*DecentralWebhookResourceDeleteDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DecentralWebhookResource_delete_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DecentralWebhooksResourceGetGet gets all registered decentral webhooks
*/
func (a *Client) DecentralWebhooksResourceGetGet(params *DecentralWebhooksResourceGetGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DecentralWebhooksResourceGetGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDecentralWebhooksResourceGetGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DecentralWebhooksResource_get_get",
		Method:             "GET",
		PathPattern:        "/api/decentralWebhook",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DecentralWebhooksResourceGetGetReader{formats: a.formats},
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
	success, ok := result.(*DecentralWebhooksResourceGetGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DecentralWebhooksResource_get_get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DecentralWebhooksResourcePutPut creates decentral webhook
*/
func (a *Client) DecentralWebhooksResourcePutPut(params *DecentralWebhooksResourcePutPutParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DecentralWebhooksResourcePutPutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDecentralWebhooksResourcePutPutParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DecentralWebhooksResource_put_put",
		Method:             "PUT",
		PathPattern:        "/api/decentralWebhook",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DecentralWebhooksResourcePutPutReader{formats: a.formats},
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
	success, ok := result.(*DecentralWebhooksResourcePutPutOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DecentralWebhooksResource_put_put: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SmartlockActionAdvancedResourceActionPost smartlocks action with callback
*/
func (a *Client) SmartlockActionAdvancedResourceActionPost(params *SmartlockActionAdvancedResourceActionPostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SmartlockActionAdvancedResourceActionPostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSmartlockActionAdvancedResourceActionPostParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SmartlockActionAdvancedResource_action_post",
		Method:             "POST",
		PathPattern:        "/smartlock/{smartlockId}/action/advanced",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SmartlockActionAdvancedResourceActionPostReader{formats: a.formats},
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
	success, ok := result.(*SmartlockActionAdvancedResourceActionPostOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SmartlockActionAdvancedResource_action_post: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SmartlockAuthsAdvancedResourcePutPut creates asynchronous smartlock authorizations
*/
func (a *Client) SmartlockAuthsAdvancedResourcePutPut(params *SmartlockAuthsAdvancedResourcePutPutParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SmartlockAuthsAdvancedResourcePutPutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSmartlockAuthsAdvancedResourcePutPutParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SmartlockAuthsAdvancedResource_put_put",
		Method:             "PUT",
		PathPattern:        "/smartlock/auth/advanced",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SmartlockAuthsAdvancedResourcePutPutReader{formats: a.formats},
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
	success, ok := result.(*SmartlockAuthsAdvancedResourcePutPutOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SmartlockAuthsAdvancedResource_put_put: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SmartlockLockActionAdvancedResourcePostLockPost locks a smartlock
*/
func (a *Client) SmartlockLockActionAdvancedResourcePostLockPost(params *SmartlockLockActionAdvancedResourcePostLockPostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SmartlockLockActionAdvancedResourcePostLockPostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSmartlockLockActionAdvancedResourcePostLockPostParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SmartlockLockActionAdvancedResource_postLock_post",
		Method:             "POST",
		PathPattern:        "/smartlock/{smartlockId}/action/lock/advanced",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SmartlockLockActionAdvancedResourcePostLockPostReader{formats: a.formats},
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
	success, ok := result.(*SmartlockLockActionAdvancedResourcePostLockPostOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SmartlockLockActionAdvancedResource_postLock_post: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SmartlockUnlockActionAdvancedResourcePostLockPost unlocks a smartlock
*/
func (a *Client) SmartlockUnlockActionAdvancedResourcePostLockPost(params *SmartlockUnlockActionAdvancedResourcePostLockPostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SmartlockUnlockActionAdvancedResourcePostLockPostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSmartlockUnlockActionAdvancedResourcePostLockPostParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SmartlockUnlockActionAdvancedResource_postLock_post",
		Method:             "POST",
		PathPattern:        "/smartlock/{smartlockId}/action/unlock/advanced",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SmartlockUnlockActionAdvancedResourcePostLockPostReader{formats: a.formats},
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
	success, ok := result.(*SmartlockUnlockActionAdvancedResourcePostLockPostOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SmartlockUnlockActionAdvancedResource_postLock_post: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
WebhookLogsResourceGetGet gets a list of webhook logs descending
*/
func (a *Client) WebhookLogsResourceGetGet(params *WebhookLogsResourceGetGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*WebhookLogsResourceGetGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWebhookLogsResourceGetGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "WebhookLogsResource_get_get",
		Method:             "GET",
		PathPattern:        "/api/key/{apiKeyId}/webhook/logs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &WebhookLogsResourceGetGetReader{formats: a.formats},
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
	success, ok := result.(*WebhookLogsResourceGetGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for WebhookLogsResource_get_get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
