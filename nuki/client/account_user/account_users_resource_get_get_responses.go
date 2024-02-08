// Code generated by go-swagger; DO NOT EDIT.

package account_user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/leonnicolas/gooki/nuki/models"
)

// AccountUsersResourceGetGetReader is a Reader for the AccountUsersResourceGetGet structure.
type AccountUsersResourceGetGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AccountUsersResourceGetGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAccountUsersResourceGetGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAccountUsersResourceGetGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /account/user] AccountUsersResource_get_get", response, response.Code())
	}
}

// NewAccountUsersResourceGetGetOK creates a AccountUsersResourceGetGetOK with default headers values
func NewAccountUsersResourceGetGetOK() *AccountUsersResourceGetGetOK {
	return &AccountUsersResourceGetGetOK{}
}

/*
AccountUsersResourceGetGetOK describes a response with status code 200, with default header values.

successful operation
*/
type AccountUsersResourceGetGetOK struct {
	Payload []*models.AccountUser
}

// IsSuccess returns true when this account users resource get get o k response has a 2xx status code
func (o *AccountUsersResourceGetGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this account users resource get get o k response has a 3xx status code
func (o *AccountUsersResourceGetGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account users resource get get o k response has a 4xx status code
func (o *AccountUsersResourceGetGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this account users resource get get o k response has a 5xx status code
func (o *AccountUsersResourceGetGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this account users resource get get o k response a status code equal to that given
func (o *AccountUsersResourceGetGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the account users resource get get o k response
func (o *AccountUsersResourceGetGetOK) Code() int {
	return 200
}

func (o *AccountUsersResourceGetGetOK) Error() string {
	return fmt.Sprintf("[GET /account/user][%d] accountUsersResourceGetGetOK  %+v", 200, o.Payload)
}

func (o *AccountUsersResourceGetGetOK) String() string {
	return fmt.Sprintf("[GET /account/user][%d] accountUsersResourceGetGetOK  %+v", 200, o.Payload)
}

func (o *AccountUsersResourceGetGetOK) GetPayload() []*models.AccountUser {
	return o.Payload
}

func (o *AccountUsersResourceGetGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountUsersResourceGetGetUnauthorized creates a AccountUsersResourceGetGetUnauthorized with default headers values
func NewAccountUsersResourceGetGetUnauthorized() *AccountUsersResourceGetGetUnauthorized {
	return &AccountUsersResourceGetGetUnauthorized{}
}

/*
AccountUsersResourceGetGetUnauthorized describes a response with status code 401, with default header values.

Not authorized
*/
type AccountUsersResourceGetGetUnauthorized struct {
}

// IsSuccess returns true when this account users resource get get unauthorized response has a 2xx status code
func (o *AccountUsersResourceGetGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this account users resource get get unauthorized response has a 3xx status code
func (o *AccountUsersResourceGetGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account users resource get get unauthorized response has a 4xx status code
func (o *AccountUsersResourceGetGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this account users resource get get unauthorized response has a 5xx status code
func (o *AccountUsersResourceGetGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this account users resource get get unauthorized response a status code equal to that given
func (o *AccountUsersResourceGetGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the account users resource get get unauthorized response
func (o *AccountUsersResourceGetGetUnauthorized) Code() int {
	return 401
}

func (o *AccountUsersResourceGetGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /account/user][%d] accountUsersResourceGetGetUnauthorized ", 401)
}

func (o *AccountUsersResourceGetGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /account/user][%d] accountUsersResourceGetGetUnauthorized ", 401)
}

func (o *AccountUsersResourceGetGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
