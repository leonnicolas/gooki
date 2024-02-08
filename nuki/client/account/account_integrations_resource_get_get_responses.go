// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/leonnicolas/gooki/nuki/models"
)

// AccountIntegrationsResourceGetGetReader is a Reader for the AccountIntegrationsResourceGetGet structure.
type AccountIntegrationsResourceGetGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AccountIntegrationsResourceGetGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAccountIntegrationsResourceGetGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAccountIntegrationsResourceGetGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /account/integration] AccountIntegrationsResource_get_get", response, response.Code())
	}
}

// NewAccountIntegrationsResourceGetGetOK creates a AccountIntegrationsResourceGetGetOK with default headers values
func NewAccountIntegrationsResourceGetGetOK() *AccountIntegrationsResourceGetGetOK {
	return &AccountIntegrationsResourceGetGetOK{}
}

/*
AccountIntegrationsResourceGetGetOK describes a response with status code 200, with default header values.

successful operation
*/
type AccountIntegrationsResourceGetGetOK struct {
	Payload []*models.AccountIntegration
}

// IsSuccess returns true when this account integrations resource get get o k response has a 2xx status code
func (o *AccountIntegrationsResourceGetGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this account integrations resource get get o k response has a 3xx status code
func (o *AccountIntegrationsResourceGetGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account integrations resource get get o k response has a 4xx status code
func (o *AccountIntegrationsResourceGetGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this account integrations resource get get o k response has a 5xx status code
func (o *AccountIntegrationsResourceGetGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this account integrations resource get get o k response a status code equal to that given
func (o *AccountIntegrationsResourceGetGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the account integrations resource get get o k response
func (o *AccountIntegrationsResourceGetGetOK) Code() int {
	return 200
}

func (o *AccountIntegrationsResourceGetGetOK) Error() string {
	return fmt.Sprintf("[GET /account/integration][%d] accountIntegrationsResourceGetGetOK  %+v", 200, o.Payload)
}

func (o *AccountIntegrationsResourceGetGetOK) String() string {
	return fmt.Sprintf("[GET /account/integration][%d] accountIntegrationsResourceGetGetOK  %+v", 200, o.Payload)
}

func (o *AccountIntegrationsResourceGetGetOK) GetPayload() []*models.AccountIntegration {
	return o.Payload
}

func (o *AccountIntegrationsResourceGetGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountIntegrationsResourceGetGetUnauthorized creates a AccountIntegrationsResourceGetGetUnauthorized with default headers values
func NewAccountIntegrationsResourceGetGetUnauthorized() *AccountIntegrationsResourceGetGetUnauthorized {
	return &AccountIntegrationsResourceGetGetUnauthorized{}
}

/*
AccountIntegrationsResourceGetGetUnauthorized describes a response with status code 401, with default header values.

Not authorized
*/
type AccountIntegrationsResourceGetGetUnauthorized struct {
}

// IsSuccess returns true when this account integrations resource get get unauthorized response has a 2xx status code
func (o *AccountIntegrationsResourceGetGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this account integrations resource get get unauthorized response has a 3xx status code
func (o *AccountIntegrationsResourceGetGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account integrations resource get get unauthorized response has a 4xx status code
func (o *AccountIntegrationsResourceGetGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this account integrations resource get get unauthorized response has a 5xx status code
func (o *AccountIntegrationsResourceGetGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this account integrations resource get get unauthorized response a status code equal to that given
func (o *AccountIntegrationsResourceGetGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the account integrations resource get get unauthorized response
func (o *AccountIntegrationsResourceGetGetUnauthorized) Code() int {
	return 401
}

func (o *AccountIntegrationsResourceGetGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /account/integration][%d] accountIntegrationsResourceGetGetUnauthorized ", 401)
}

func (o *AccountIntegrationsResourceGetGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /account/integration][%d] accountIntegrationsResourceGetGetUnauthorized ", 401)
}

func (o *AccountIntegrationsResourceGetGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}