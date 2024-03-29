// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// AccountOtpResourceDeleteDeleteReader is a Reader for the AccountOtpResourceDeleteDelete structure.
type AccountOtpResourceDeleteDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AccountOtpResourceDeleteDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewAccountOtpResourceDeleteDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAccountOtpResourceDeleteDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /account/otp] AccountOtpResource_delete_delete", response, response.Code())
	}
}

// NewAccountOtpResourceDeleteDeleteNoContent creates a AccountOtpResourceDeleteDeleteNoContent with default headers values
func NewAccountOtpResourceDeleteDeleteNoContent() *AccountOtpResourceDeleteDeleteNoContent {
	return &AccountOtpResourceDeleteDeleteNoContent{}
}

/*
AccountOtpResourceDeleteDeleteNoContent describes a response with status code 204, with default header values.

Ok
*/
type AccountOtpResourceDeleteDeleteNoContent struct {
}

// IsSuccess returns true when this account otp resource delete delete no content response has a 2xx status code
func (o *AccountOtpResourceDeleteDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this account otp resource delete delete no content response has a 3xx status code
func (o *AccountOtpResourceDeleteDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account otp resource delete delete no content response has a 4xx status code
func (o *AccountOtpResourceDeleteDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this account otp resource delete delete no content response has a 5xx status code
func (o *AccountOtpResourceDeleteDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this account otp resource delete delete no content response a status code equal to that given
func (o *AccountOtpResourceDeleteDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the account otp resource delete delete no content response
func (o *AccountOtpResourceDeleteDeleteNoContent) Code() int {
	return 204
}

func (o *AccountOtpResourceDeleteDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /account/otp][%d] accountOtpResourceDeleteDeleteNoContent ", 204)
}

func (o *AccountOtpResourceDeleteDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /account/otp][%d] accountOtpResourceDeleteDeleteNoContent ", 204)
}

func (o *AccountOtpResourceDeleteDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAccountOtpResourceDeleteDeleteUnauthorized creates a AccountOtpResourceDeleteDeleteUnauthorized with default headers values
func NewAccountOtpResourceDeleteDeleteUnauthorized() *AccountOtpResourceDeleteDeleteUnauthorized {
	return &AccountOtpResourceDeleteDeleteUnauthorized{}
}

/*
AccountOtpResourceDeleteDeleteUnauthorized describes a response with status code 401, with default header values.

Not authorized
*/
type AccountOtpResourceDeleteDeleteUnauthorized struct {
}

// IsSuccess returns true when this account otp resource delete delete unauthorized response has a 2xx status code
func (o *AccountOtpResourceDeleteDeleteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this account otp resource delete delete unauthorized response has a 3xx status code
func (o *AccountOtpResourceDeleteDeleteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account otp resource delete delete unauthorized response has a 4xx status code
func (o *AccountOtpResourceDeleteDeleteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this account otp resource delete delete unauthorized response has a 5xx status code
func (o *AccountOtpResourceDeleteDeleteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this account otp resource delete delete unauthorized response a status code equal to that given
func (o *AccountOtpResourceDeleteDeleteUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the account otp resource delete delete unauthorized response
func (o *AccountOtpResourceDeleteDeleteUnauthorized) Code() int {
	return 401
}

func (o *AccountOtpResourceDeleteDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /account/otp][%d] accountOtpResourceDeleteDeleteUnauthorized ", 401)
}

func (o *AccountOtpResourceDeleteDeleteUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /account/otp][%d] accountOtpResourceDeleteDeleteUnauthorized ", 401)
}

func (o *AccountOtpResourceDeleteDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
