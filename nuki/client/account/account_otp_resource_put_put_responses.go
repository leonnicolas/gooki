// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// AccountOtpResourcePutPutReader is a Reader for the AccountOtpResourcePutPut structure.
type AccountOtpResourcePutPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AccountOtpResourcePutPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAccountOtpResourcePutPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 405:
		result := NewAccountOtpResourcePutPutMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /account/otp] AccountOtpResource_put_put", response, response.Code())
	}
}

// NewAccountOtpResourcePutPutOK creates a AccountOtpResourcePutPutOK with default headers values
func NewAccountOtpResourcePutPutOK() *AccountOtpResourcePutPutOK {
	return &AccountOtpResourcePutPutOK{}
}

/*
AccountOtpResourcePutPutOK describes a response with status code 200, with default header values.

Ok
*/
type AccountOtpResourcePutPutOK struct {
	Payload string
}

// IsSuccess returns true when this account otp resource put put o k response has a 2xx status code
func (o *AccountOtpResourcePutPutOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this account otp resource put put o k response has a 3xx status code
func (o *AccountOtpResourcePutPutOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account otp resource put put o k response has a 4xx status code
func (o *AccountOtpResourcePutPutOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this account otp resource put put o k response has a 5xx status code
func (o *AccountOtpResourcePutPutOK) IsServerError() bool {
	return false
}

// IsCode returns true when this account otp resource put put o k response a status code equal to that given
func (o *AccountOtpResourcePutPutOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the account otp resource put put o k response
func (o *AccountOtpResourcePutPutOK) Code() int {
	return 200
}

func (o *AccountOtpResourcePutPutOK) Error() string {
	return fmt.Sprintf("[PUT /account/otp][%d] accountOtpResourcePutPutOK  %+v", 200, o.Payload)
}

func (o *AccountOtpResourcePutPutOK) String() string {
	return fmt.Sprintf("[PUT /account/otp][%d] accountOtpResourcePutPutOK  %+v", 200, o.Payload)
}

func (o *AccountOtpResourcePutPutOK) GetPayload() string {
	return o.Payload
}

func (o *AccountOtpResourcePutPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountOtpResourcePutPutMethodNotAllowed creates a AccountOtpResourcePutPutMethodNotAllowed with default headers values
func NewAccountOtpResourcePutPutMethodNotAllowed() *AccountOtpResourcePutPutMethodNotAllowed {
	return &AccountOtpResourcePutPutMethodNotAllowed{}
}

/*
AccountOtpResourcePutPutMethodNotAllowed describes a response with status code 405, with default header values.

One time password is already enabled
*/
type AccountOtpResourcePutPutMethodNotAllowed struct {
}

// IsSuccess returns true when this account otp resource put put method not allowed response has a 2xx status code
func (o *AccountOtpResourcePutPutMethodNotAllowed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this account otp resource put put method not allowed response has a 3xx status code
func (o *AccountOtpResourcePutPutMethodNotAllowed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account otp resource put put method not allowed response has a 4xx status code
func (o *AccountOtpResourcePutPutMethodNotAllowed) IsClientError() bool {
	return true
}

// IsServerError returns true when this account otp resource put put method not allowed response has a 5xx status code
func (o *AccountOtpResourcePutPutMethodNotAllowed) IsServerError() bool {
	return false
}

// IsCode returns true when this account otp resource put put method not allowed response a status code equal to that given
func (o *AccountOtpResourcePutPutMethodNotAllowed) IsCode(code int) bool {
	return code == 405
}

// Code gets the status code for the account otp resource put put method not allowed response
func (o *AccountOtpResourcePutPutMethodNotAllowed) Code() int {
	return 405
}

func (o *AccountOtpResourcePutPutMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PUT /account/otp][%d] accountOtpResourcePutPutMethodNotAllowed ", 405)
}

func (o *AccountOtpResourcePutPutMethodNotAllowed) String() string {
	return fmt.Sprintf("[PUT /account/otp][%d] accountOtpResourcePutPutMethodNotAllowed ", 405)
}

func (o *AccountOtpResourcePutPutMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
