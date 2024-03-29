// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// AccountOtpResourcePostPostReader is a Reader for the AccountOtpResourcePostPost structure.
type AccountOtpResourcePostPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AccountOtpResourcePostPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewAccountOtpResourcePostPostNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAccountOtpResourcePostPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAccountOtpResourcePostPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewAccountOtpResourcePostPostTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /account/otp] AccountOtpResource_post_post", response, response.Code())
	}
}

// NewAccountOtpResourcePostPostNoContent creates a AccountOtpResourcePostPostNoContent with default headers values
func NewAccountOtpResourcePostPostNoContent() *AccountOtpResourcePostPostNoContent {
	return &AccountOtpResourcePostPostNoContent{}
}

/*
AccountOtpResourcePostPostNoContent describes a response with status code 204, with default header values.

Ok
*/
type AccountOtpResourcePostPostNoContent struct {
}

// IsSuccess returns true when this account otp resource post post no content response has a 2xx status code
func (o *AccountOtpResourcePostPostNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this account otp resource post post no content response has a 3xx status code
func (o *AccountOtpResourcePostPostNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account otp resource post post no content response has a 4xx status code
func (o *AccountOtpResourcePostPostNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this account otp resource post post no content response has a 5xx status code
func (o *AccountOtpResourcePostPostNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this account otp resource post post no content response a status code equal to that given
func (o *AccountOtpResourcePostPostNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the account otp resource post post no content response
func (o *AccountOtpResourcePostPostNoContent) Code() int {
	return 204
}

func (o *AccountOtpResourcePostPostNoContent) Error() string {
	return fmt.Sprintf("[POST /account/otp][%d] accountOtpResourcePostPostNoContent ", 204)
}

func (o *AccountOtpResourcePostPostNoContent) String() string {
	return fmt.Sprintf("[POST /account/otp][%d] accountOtpResourcePostPostNoContent ", 204)
}

func (o *AccountOtpResourcePostPostNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAccountOtpResourcePostPostBadRequest creates a AccountOtpResourcePostPostBadRequest with default headers values
func NewAccountOtpResourcePostPostBadRequest() *AccountOtpResourcePostPostBadRequest {
	return &AccountOtpResourcePostPostBadRequest{}
}

/*
AccountOtpResourcePostPostBadRequest describes a response with status code 400, with default header values.

One time password empty
*/
type AccountOtpResourcePostPostBadRequest struct {
}

// IsSuccess returns true when this account otp resource post post bad request response has a 2xx status code
func (o *AccountOtpResourcePostPostBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this account otp resource post post bad request response has a 3xx status code
func (o *AccountOtpResourcePostPostBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account otp resource post post bad request response has a 4xx status code
func (o *AccountOtpResourcePostPostBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this account otp resource post post bad request response has a 5xx status code
func (o *AccountOtpResourcePostPostBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this account otp resource post post bad request response a status code equal to that given
func (o *AccountOtpResourcePostPostBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the account otp resource post post bad request response
func (o *AccountOtpResourcePostPostBadRequest) Code() int {
	return 400
}

func (o *AccountOtpResourcePostPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /account/otp][%d] accountOtpResourcePostPostBadRequest ", 400)
}

func (o *AccountOtpResourcePostPostBadRequest) String() string {
	return fmt.Sprintf("[POST /account/otp][%d] accountOtpResourcePostPostBadRequest ", 400)
}

func (o *AccountOtpResourcePostPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAccountOtpResourcePostPostUnauthorized creates a AccountOtpResourcePostPostUnauthorized with default headers values
func NewAccountOtpResourcePostPostUnauthorized() *AccountOtpResourcePostPostUnauthorized {
	return &AccountOtpResourcePostPostUnauthorized{}
}

/*
AccountOtpResourcePostPostUnauthorized describes a response with status code 401, with default header values.

Not authorized or one time password wrong
*/
type AccountOtpResourcePostPostUnauthorized struct {
}

// IsSuccess returns true when this account otp resource post post unauthorized response has a 2xx status code
func (o *AccountOtpResourcePostPostUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this account otp resource post post unauthorized response has a 3xx status code
func (o *AccountOtpResourcePostPostUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account otp resource post post unauthorized response has a 4xx status code
func (o *AccountOtpResourcePostPostUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this account otp resource post post unauthorized response has a 5xx status code
func (o *AccountOtpResourcePostPostUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this account otp resource post post unauthorized response a status code equal to that given
func (o *AccountOtpResourcePostPostUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the account otp resource post post unauthorized response
func (o *AccountOtpResourcePostPostUnauthorized) Code() int {
	return 401
}

func (o *AccountOtpResourcePostPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /account/otp][%d] accountOtpResourcePostPostUnauthorized ", 401)
}

func (o *AccountOtpResourcePostPostUnauthorized) String() string {
	return fmt.Sprintf("[POST /account/otp][%d] accountOtpResourcePostPostUnauthorized ", 401)
}

func (o *AccountOtpResourcePostPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAccountOtpResourcePostPostTooManyRequests creates a AccountOtpResourcePostPostTooManyRequests with default headers values
func NewAccountOtpResourcePostPostTooManyRequests() *AccountOtpResourcePostPostTooManyRequests {
	return &AccountOtpResourcePostPostTooManyRequests{}
}

/*
AccountOtpResourcePostPostTooManyRequests describes a response with status code 429, with default header values.

Too many failed attempts
*/
type AccountOtpResourcePostPostTooManyRequests struct {
}

// IsSuccess returns true when this account otp resource post post too many requests response has a 2xx status code
func (o *AccountOtpResourcePostPostTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this account otp resource post post too many requests response has a 3xx status code
func (o *AccountOtpResourcePostPostTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account otp resource post post too many requests response has a 4xx status code
func (o *AccountOtpResourcePostPostTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this account otp resource post post too many requests response has a 5xx status code
func (o *AccountOtpResourcePostPostTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this account otp resource post post too many requests response a status code equal to that given
func (o *AccountOtpResourcePostPostTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the account otp resource post post too many requests response
func (o *AccountOtpResourcePostPostTooManyRequests) Code() int {
	return 429
}

func (o *AccountOtpResourcePostPostTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /account/otp][%d] accountOtpResourcePostPostTooManyRequests ", 429)
}

func (o *AccountOtpResourcePostPostTooManyRequests) String() string {
	return fmt.Sprintf("[POST /account/otp][%d] accountOtpResourcePostPostTooManyRequests ", 429)
}

func (o *AccountOtpResourcePostPostTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
