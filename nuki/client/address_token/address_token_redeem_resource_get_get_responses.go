// Code generated by go-swagger; DO NOT EDIT.

package address_token

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/leonnicolas/gooki/nuki/models"
)

// AddressTokenRedeemResourceGetGetReader is a Reader for the AddressTokenRedeemResourceGetGet structure.
type AddressTokenRedeemResourceGetGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddressTokenRedeemResourceGetGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddressTokenRedeemResourceGetGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAddressTokenRedeemResourceGetGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAddressTokenRedeemResourceGetGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /address/token/{id}/redeem] AddressTokenRedeemResource_get_get", response, response.Code())
	}
}

// NewAddressTokenRedeemResourceGetGetOK creates a AddressTokenRedeemResourceGetGetOK with default headers values
func NewAddressTokenRedeemResourceGetGetOK() *AddressTokenRedeemResourceGetGetOK {
	return &AddressTokenRedeemResourceGetGetOK{}
}

/*
AddressTokenRedeemResourceGetGetOK describes a response with status code 200, with default header values.

successful operation
*/
type AddressTokenRedeemResourceGetGetOK struct {
	Payload *models.AddressToken
}

// IsSuccess returns true when this address token redeem resource get get o k response has a 2xx status code
func (o *AddressTokenRedeemResourceGetGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this address token redeem resource get get o k response has a 3xx status code
func (o *AddressTokenRedeemResourceGetGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this address token redeem resource get get o k response has a 4xx status code
func (o *AddressTokenRedeemResourceGetGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this address token redeem resource get get o k response has a 5xx status code
func (o *AddressTokenRedeemResourceGetGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this address token redeem resource get get o k response a status code equal to that given
func (o *AddressTokenRedeemResourceGetGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the address token redeem resource get get o k response
func (o *AddressTokenRedeemResourceGetGetOK) Code() int {
	return 200
}

func (o *AddressTokenRedeemResourceGetGetOK) Error() string {
	return fmt.Sprintf("[GET /address/token/{id}/redeem][%d] addressTokenRedeemResourceGetGetOK  %+v", 200, o.Payload)
}

func (o *AddressTokenRedeemResourceGetGetOK) String() string {
	return fmt.Sprintf("[GET /address/token/{id}/redeem][%d] addressTokenRedeemResourceGetGetOK  %+v", 200, o.Payload)
}

func (o *AddressTokenRedeemResourceGetGetOK) GetPayload() *models.AddressToken {
	return o.Payload
}

func (o *AddressTokenRedeemResourceGetGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AddressToken)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddressTokenRedeemResourceGetGetUnauthorized creates a AddressTokenRedeemResourceGetGetUnauthorized with default headers values
func NewAddressTokenRedeemResourceGetGetUnauthorized() *AddressTokenRedeemResourceGetGetUnauthorized {
	return &AddressTokenRedeemResourceGetGetUnauthorized{}
}

/*
AddressTokenRedeemResourceGetGetUnauthorized describes a response with status code 401, with default header values.

Not authorized
*/
type AddressTokenRedeemResourceGetGetUnauthorized struct {
}

// IsSuccess returns true when this address token redeem resource get get unauthorized response has a 2xx status code
func (o *AddressTokenRedeemResourceGetGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this address token redeem resource get get unauthorized response has a 3xx status code
func (o *AddressTokenRedeemResourceGetGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this address token redeem resource get get unauthorized response has a 4xx status code
func (o *AddressTokenRedeemResourceGetGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this address token redeem resource get get unauthorized response has a 5xx status code
func (o *AddressTokenRedeemResourceGetGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this address token redeem resource get get unauthorized response a status code equal to that given
func (o *AddressTokenRedeemResourceGetGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the address token redeem resource get get unauthorized response
func (o *AddressTokenRedeemResourceGetGetUnauthorized) Code() int {
	return 401
}

func (o *AddressTokenRedeemResourceGetGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /address/token/{id}/redeem][%d] addressTokenRedeemResourceGetGetUnauthorized ", 401)
}

func (o *AddressTokenRedeemResourceGetGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /address/token/{id}/redeem][%d] addressTokenRedeemResourceGetGetUnauthorized ", 401)
}

func (o *AddressTokenRedeemResourceGetGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddressTokenRedeemResourceGetGetNotFound creates a AddressTokenRedeemResourceGetGetNotFound with default headers values
func NewAddressTokenRedeemResourceGetGetNotFound() *AddressTokenRedeemResourceGetGetNotFound {
	return &AddressTokenRedeemResourceGetGetNotFound{}
}

/*
AddressTokenRedeemResourceGetGetNotFound describes a response with status code 404, with default header values.

Token not found
*/
type AddressTokenRedeemResourceGetGetNotFound struct {
}

// IsSuccess returns true when this address token redeem resource get get not found response has a 2xx status code
func (o *AddressTokenRedeemResourceGetGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this address token redeem resource get get not found response has a 3xx status code
func (o *AddressTokenRedeemResourceGetGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this address token redeem resource get get not found response has a 4xx status code
func (o *AddressTokenRedeemResourceGetGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this address token redeem resource get get not found response has a 5xx status code
func (o *AddressTokenRedeemResourceGetGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this address token redeem resource get get not found response a status code equal to that given
func (o *AddressTokenRedeemResourceGetGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the address token redeem resource get get not found response
func (o *AddressTokenRedeemResourceGetGetNotFound) Code() int {
	return 404
}

func (o *AddressTokenRedeemResourceGetGetNotFound) Error() string {
	return fmt.Sprintf("[GET /address/token/{id}/redeem][%d] addressTokenRedeemResourceGetGetNotFound ", 404)
}

func (o *AddressTokenRedeemResourceGetGetNotFound) String() string {
	return fmt.Sprintf("[GET /address/token/{id}/redeem][%d] addressTokenRedeemResourceGetGetNotFound ", 404)
}

func (o *AddressTokenRedeemResourceGetGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
