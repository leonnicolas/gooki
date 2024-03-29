// Code generated by go-swagger; DO NOT EDIT.

package address_token

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// AddressTokenRedeemResourcePostPostReader is a Reader for the AddressTokenRedeemResourcePostPost structure.
type AddressTokenRedeemResourcePostPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddressTokenRedeemResourcePostPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewAddressTokenRedeemResourcePostPostNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddressTokenRedeemResourcePostPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAddressTokenRedeemResourcePostPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAddressTokenRedeemResourcePostPostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /address/token/{id}/redeem] AddressTokenRedeemResource_post_post", response, response.Code())
	}
}

// NewAddressTokenRedeemResourcePostPostNoContent creates a AddressTokenRedeemResourcePostPostNoContent with default headers values
func NewAddressTokenRedeemResourcePostPostNoContent() *AddressTokenRedeemResourcePostPostNoContent {
	return &AddressTokenRedeemResourcePostPostNoContent{}
}

/*
AddressTokenRedeemResourcePostPostNoContent describes a response with status code 204, with default header values.

Ok
*/
type AddressTokenRedeemResourcePostPostNoContent struct {
}

// IsSuccess returns true when this address token redeem resource post post no content response has a 2xx status code
func (o *AddressTokenRedeemResourcePostPostNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this address token redeem resource post post no content response has a 3xx status code
func (o *AddressTokenRedeemResourcePostPostNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this address token redeem resource post post no content response has a 4xx status code
func (o *AddressTokenRedeemResourcePostPostNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this address token redeem resource post post no content response has a 5xx status code
func (o *AddressTokenRedeemResourcePostPostNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this address token redeem resource post post no content response a status code equal to that given
func (o *AddressTokenRedeemResourcePostPostNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the address token redeem resource post post no content response
func (o *AddressTokenRedeemResourcePostPostNoContent) Code() int {
	return 204
}

func (o *AddressTokenRedeemResourcePostPostNoContent) Error() string {
	return fmt.Sprintf("[POST /address/token/{id}/redeem][%d] addressTokenRedeemResourcePostPostNoContent ", 204)
}

func (o *AddressTokenRedeemResourcePostPostNoContent) String() string {
	return fmt.Sprintf("[POST /address/token/{id}/redeem][%d] addressTokenRedeemResourcePostPostNoContent ", 204)
}

func (o *AddressTokenRedeemResourcePostPostNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddressTokenRedeemResourcePostPostBadRequest creates a AddressTokenRedeemResourcePostPostBadRequest with default headers values
func NewAddressTokenRedeemResourcePostPostBadRequest() *AddressTokenRedeemResourcePostPostBadRequest {
	return &AddressTokenRedeemResourcePostPostBadRequest{}
}

/*
AddressTokenRedeemResourcePostPostBadRequest describes a response with status code 400, with default header values.

Invalid parameter given
*/
type AddressTokenRedeemResourcePostPostBadRequest struct {
}

// IsSuccess returns true when this address token redeem resource post post bad request response has a 2xx status code
func (o *AddressTokenRedeemResourcePostPostBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this address token redeem resource post post bad request response has a 3xx status code
func (o *AddressTokenRedeemResourcePostPostBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this address token redeem resource post post bad request response has a 4xx status code
func (o *AddressTokenRedeemResourcePostPostBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this address token redeem resource post post bad request response has a 5xx status code
func (o *AddressTokenRedeemResourcePostPostBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this address token redeem resource post post bad request response a status code equal to that given
func (o *AddressTokenRedeemResourcePostPostBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the address token redeem resource post post bad request response
func (o *AddressTokenRedeemResourcePostPostBadRequest) Code() int {
	return 400
}

func (o *AddressTokenRedeemResourcePostPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /address/token/{id}/redeem][%d] addressTokenRedeemResourcePostPostBadRequest ", 400)
}

func (o *AddressTokenRedeemResourcePostPostBadRequest) String() string {
	return fmt.Sprintf("[POST /address/token/{id}/redeem][%d] addressTokenRedeemResourcePostPostBadRequest ", 400)
}

func (o *AddressTokenRedeemResourcePostPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddressTokenRedeemResourcePostPostUnauthorized creates a AddressTokenRedeemResourcePostPostUnauthorized with default headers values
func NewAddressTokenRedeemResourcePostPostUnauthorized() *AddressTokenRedeemResourcePostPostUnauthorized {
	return &AddressTokenRedeemResourcePostPostUnauthorized{}
}

/*
AddressTokenRedeemResourcePostPostUnauthorized describes a response with status code 401, with default header values.

Not authorized
*/
type AddressTokenRedeemResourcePostPostUnauthorized struct {
}

// IsSuccess returns true when this address token redeem resource post post unauthorized response has a 2xx status code
func (o *AddressTokenRedeemResourcePostPostUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this address token redeem resource post post unauthorized response has a 3xx status code
func (o *AddressTokenRedeemResourcePostPostUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this address token redeem resource post post unauthorized response has a 4xx status code
func (o *AddressTokenRedeemResourcePostPostUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this address token redeem resource post post unauthorized response has a 5xx status code
func (o *AddressTokenRedeemResourcePostPostUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this address token redeem resource post post unauthorized response a status code equal to that given
func (o *AddressTokenRedeemResourcePostPostUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the address token redeem resource post post unauthorized response
func (o *AddressTokenRedeemResourcePostPostUnauthorized) Code() int {
	return 401
}

func (o *AddressTokenRedeemResourcePostPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /address/token/{id}/redeem][%d] addressTokenRedeemResourcePostPostUnauthorized ", 401)
}

func (o *AddressTokenRedeemResourcePostPostUnauthorized) String() string {
	return fmt.Sprintf("[POST /address/token/{id}/redeem][%d] addressTokenRedeemResourcePostPostUnauthorized ", 401)
}

func (o *AddressTokenRedeemResourcePostPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddressTokenRedeemResourcePostPostNotFound creates a AddressTokenRedeemResourcePostPostNotFound with default headers values
func NewAddressTokenRedeemResourcePostPostNotFound() *AddressTokenRedeemResourcePostPostNotFound {
	return &AddressTokenRedeemResourcePostPostNotFound{}
}

/*
AddressTokenRedeemResourcePostPostNotFound describes a response with status code 404, with default header values.

Token not found
*/
type AddressTokenRedeemResourcePostPostNotFound struct {
}

// IsSuccess returns true when this address token redeem resource post post not found response has a 2xx status code
func (o *AddressTokenRedeemResourcePostPostNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this address token redeem resource post post not found response has a 3xx status code
func (o *AddressTokenRedeemResourcePostPostNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this address token redeem resource post post not found response has a 4xx status code
func (o *AddressTokenRedeemResourcePostPostNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this address token redeem resource post post not found response has a 5xx status code
func (o *AddressTokenRedeemResourcePostPostNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this address token redeem resource post post not found response a status code equal to that given
func (o *AddressTokenRedeemResourcePostPostNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the address token redeem resource post post not found response
func (o *AddressTokenRedeemResourcePostPostNotFound) Code() int {
	return 404
}

func (o *AddressTokenRedeemResourcePostPostNotFound) Error() string {
	return fmt.Sprintf("[POST /address/token/{id}/redeem][%d] addressTokenRedeemResourcePostPostNotFound ", 404)
}

func (o *AddressTokenRedeemResourcePostPostNotFound) String() string {
	return fmt.Sprintf("[POST /address/token/{id}/redeem][%d] addressTokenRedeemResourcePostPostNotFound ", 404)
}

func (o *AddressTokenRedeemResourcePostPostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
