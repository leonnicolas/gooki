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

// AccountUserResourcePostPostReader is a Reader for the AccountUserResourcePostPost structure.
type AccountUserResourcePostPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AccountUserResourcePostPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAccountUserResourcePostPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewAccountUserResourcePostPostNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAccountUserResourcePostPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAccountUserResourcePostPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewAccountUserResourcePostPostConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /account/user/{accountUserId}] AccountUserResource_post_post", response, response.Code())
	}
}

// NewAccountUserResourcePostPostOK creates a AccountUserResourcePostPostOK with default headers values
func NewAccountUserResourcePostPostOK() *AccountUserResourcePostPostOK {
	return &AccountUserResourcePostPostOK{}
}

/*
AccountUserResourcePostPostOK describes a response with status code 200, with default header values.

successful operation
*/
type AccountUserResourcePostPostOK struct {
	Payload *models.AccountUser
}

// IsSuccess returns true when this account user resource post post o k response has a 2xx status code
func (o *AccountUserResourcePostPostOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this account user resource post post o k response has a 3xx status code
func (o *AccountUserResourcePostPostOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account user resource post post o k response has a 4xx status code
func (o *AccountUserResourcePostPostOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this account user resource post post o k response has a 5xx status code
func (o *AccountUserResourcePostPostOK) IsServerError() bool {
	return false
}

// IsCode returns true when this account user resource post post o k response a status code equal to that given
func (o *AccountUserResourcePostPostOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the account user resource post post o k response
func (o *AccountUserResourcePostPostOK) Code() int {
	return 200
}

func (o *AccountUserResourcePostPostOK) Error() string {
	return fmt.Sprintf("[POST /account/user/{accountUserId}][%d] accountUserResourcePostPostOK  %+v", 200, o.Payload)
}

func (o *AccountUserResourcePostPostOK) String() string {
	return fmt.Sprintf("[POST /account/user/{accountUserId}][%d] accountUserResourcePostPostOK  %+v", 200, o.Payload)
}

func (o *AccountUserResourcePostPostOK) GetPayload() *models.AccountUser {
	return o.Payload
}

func (o *AccountUserResourcePostPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccountUser)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountUserResourcePostPostNoContent creates a AccountUserResourcePostPostNoContent with default headers values
func NewAccountUserResourcePostPostNoContent() *AccountUserResourcePostPostNoContent {
	return &AccountUserResourcePostPostNoContent{}
}

/*
AccountUserResourcePostPostNoContent describes a response with status code 204, with default header values.

Ok
*/
type AccountUserResourcePostPostNoContent struct {
}

// IsSuccess returns true when this account user resource post post no content response has a 2xx status code
func (o *AccountUserResourcePostPostNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this account user resource post post no content response has a 3xx status code
func (o *AccountUserResourcePostPostNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account user resource post post no content response has a 4xx status code
func (o *AccountUserResourcePostPostNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this account user resource post post no content response has a 5xx status code
func (o *AccountUserResourcePostPostNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this account user resource post post no content response a status code equal to that given
func (o *AccountUserResourcePostPostNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the account user resource post post no content response
func (o *AccountUserResourcePostPostNoContent) Code() int {
	return 204
}

func (o *AccountUserResourcePostPostNoContent) Error() string {
	return fmt.Sprintf("[POST /account/user/{accountUserId}][%d] accountUserResourcePostPostNoContent ", 204)
}

func (o *AccountUserResourcePostPostNoContent) String() string {
	return fmt.Sprintf("[POST /account/user/{accountUserId}][%d] accountUserResourcePostPostNoContent ", 204)
}

func (o *AccountUserResourcePostPostNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAccountUserResourcePostPostBadRequest creates a AccountUserResourcePostPostBadRequest with default headers values
func NewAccountUserResourcePostPostBadRequest() *AccountUserResourcePostPostBadRequest {
	return &AccountUserResourcePostPostBadRequest{}
}

/*
AccountUserResourcePostPostBadRequest describes a response with status code 400, with default header values.

Invalid E-Mail address or name supplied
*/
type AccountUserResourcePostPostBadRequest struct {
}

// IsSuccess returns true when this account user resource post post bad request response has a 2xx status code
func (o *AccountUserResourcePostPostBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this account user resource post post bad request response has a 3xx status code
func (o *AccountUserResourcePostPostBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account user resource post post bad request response has a 4xx status code
func (o *AccountUserResourcePostPostBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this account user resource post post bad request response has a 5xx status code
func (o *AccountUserResourcePostPostBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this account user resource post post bad request response a status code equal to that given
func (o *AccountUserResourcePostPostBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the account user resource post post bad request response
func (o *AccountUserResourcePostPostBadRequest) Code() int {
	return 400
}

func (o *AccountUserResourcePostPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /account/user/{accountUserId}][%d] accountUserResourcePostPostBadRequest ", 400)
}

func (o *AccountUserResourcePostPostBadRequest) String() string {
	return fmt.Sprintf("[POST /account/user/{accountUserId}][%d] accountUserResourcePostPostBadRequest ", 400)
}

func (o *AccountUserResourcePostPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAccountUserResourcePostPostUnauthorized creates a AccountUserResourcePostPostUnauthorized with default headers values
func NewAccountUserResourcePostPostUnauthorized() *AccountUserResourcePostPostUnauthorized {
	return &AccountUserResourcePostPostUnauthorized{}
}

/*
AccountUserResourcePostPostUnauthorized describes a response with status code 401, with default header values.

Not authorized
*/
type AccountUserResourcePostPostUnauthorized struct {
}

// IsSuccess returns true when this account user resource post post unauthorized response has a 2xx status code
func (o *AccountUserResourcePostPostUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this account user resource post post unauthorized response has a 3xx status code
func (o *AccountUserResourcePostPostUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account user resource post post unauthorized response has a 4xx status code
func (o *AccountUserResourcePostPostUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this account user resource post post unauthorized response has a 5xx status code
func (o *AccountUserResourcePostPostUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this account user resource post post unauthorized response a status code equal to that given
func (o *AccountUserResourcePostPostUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the account user resource post post unauthorized response
func (o *AccountUserResourcePostPostUnauthorized) Code() int {
	return 401
}

func (o *AccountUserResourcePostPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /account/user/{accountUserId}][%d] accountUserResourcePostPostUnauthorized ", 401)
}

func (o *AccountUserResourcePostPostUnauthorized) String() string {
	return fmt.Sprintf("[POST /account/user/{accountUserId}][%d] accountUserResourcePostPostUnauthorized ", 401)
}

func (o *AccountUserResourcePostPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAccountUserResourcePostPostConflict creates a AccountUserResourcePostPostConflict with default headers values
func NewAccountUserResourcePostPostConflict() *AccountUserResourcePostPostConflict {
	return &AccountUserResourcePostPostConflict{}
}

/*
AccountUserResourcePostPostConflict describes a response with status code 409, with default header values.

E-Mail address already exists
*/
type AccountUserResourcePostPostConflict struct {
}

// IsSuccess returns true when this account user resource post post conflict response has a 2xx status code
func (o *AccountUserResourcePostPostConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this account user resource post post conflict response has a 3xx status code
func (o *AccountUserResourcePostPostConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account user resource post post conflict response has a 4xx status code
func (o *AccountUserResourcePostPostConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this account user resource post post conflict response has a 5xx status code
func (o *AccountUserResourcePostPostConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this account user resource post post conflict response a status code equal to that given
func (o *AccountUserResourcePostPostConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the account user resource post post conflict response
func (o *AccountUserResourcePostPostConflict) Code() int {
	return 409
}

func (o *AccountUserResourcePostPostConflict) Error() string {
	return fmt.Sprintf("[POST /account/user/{accountUserId}][%d] accountUserResourcePostPostConflict ", 409)
}

func (o *AccountUserResourcePostPostConflict) String() string {
	return fmt.Sprintf("[POST /account/user/{accountUserId}][%d] accountUserResourcePostPostConflict ", 409)
}

func (o *AccountUserResourcePostPostConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
