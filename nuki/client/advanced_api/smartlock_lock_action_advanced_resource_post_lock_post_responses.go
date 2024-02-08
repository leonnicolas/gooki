// Code generated by go-swagger; DO NOT EDIT.

package advanced_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/leonnicolas/gooki/nuki/models"
)

// SmartlockLockActionAdvancedResourcePostLockPostReader is a Reader for the SmartlockLockActionAdvancedResourcePostLockPost structure.
type SmartlockLockActionAdvancedResourcePostLockPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SmartlockLockActionAdvancedResourcePostLockPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSmartlockLockActionAdvancedResourcePostLockPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSmartlockLockActionAdvancedResourcePostLockPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSmartlockLockActionAdvancedResourcePostLockPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewSmartlockLockActionAdvancedResourcePostLockPostMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /smartlock/{smartlockId}/action/lock/advanced] SmartlockLockActionAdvancedResource_postLock_post", response, response.Code())
	}
}

// NewSmartlockLockActionAdvancedResourcePostLockPostOK creates a SmartlockLockActionAdvancedResourcePostLockPostOK with default headers values
func NewSmartlockLockActionAdvancedResourcePostLockPostOK() *SmartlockLockActionAdvancedResourcePostLockPostOK {
	return &SmartlockLockActionAdvancedResourcePostLockPostOK{}
}

/*
SmartlockLockActionAdvancedResourcePostLockPostOK describes a response with status code 200, with default header values.

Ok
*/
type SmartlockLockActionAdvancedResourcePostLockPostOK struct {
	Payload *models.AdvancedConfirmationResponse
}

// IsSuccess returns true when this smartlock lock action advanced resource post lock post o k response has a 2xx status code
func (o *SmartlockLockActionAdvancedResourcePostLockPostOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this smartlock lock action advanced resource post lock post o k response has a 3xx status code
func (o *SmartlockLockActionAdvancedResourcePostLockPostOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this smartlock lock action advanced resource post lock post o k response has a 4xx status code
func (o *SmartlockLockActionAdvancedResourcePostLockPostOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this smartlock lock action advanced resource post lock post o k response has a 5xx status code
func (o *SmartlockLockActionAdvancedResourcePostLockPostOK) IsServerError() bool {
	return false
}

// IsCode returns true when this smartlock lock action advanced resource post lock post o k response a status code equal to that given
func (o *SmartlockLockActionAdvancedResourcePostLockPostOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the smartlock lock action advanced resource post lock post o k response
func (o *SmartlockLockActionAdvancedResourcePostLockPostOK) Code() int {
	return 200
}

func (o *SmartlockLockActionAdvancedResourcePostLockPostOK) Error() string {
	return fmt.Sprintf("[POST /smartlock/{smartlockId}/action/lock/advanced][%d] smartlockLockActionAdvancedResourcePostLockPostOK  %+v", 200, o.Payload)
}

func (o *SmartlockLockActionAdvancedResourcePostLockPostOK) String() string {
	return fmt.Sprintf("[POST /smartlock/{smartlockId}/action/lock/advanced][%d] smartlockLockActionAdvancedResourcePostLockPostOK  %+v", 200, o.Payload)
}

func (o *SmartlockLockActionAdvancedResourcePostLockPostOK) GetPayload() *models.AdvancedConfirmationResponse {
	return o.Payload
}

func (o *SmartlockLockActionAdvancedResourcePostLockPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AdvancedConfirmationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSmartlockLockActionAdvancedResourcePostLockPostBadRequest creates a SmartlockLockActionAdvancedResourcePostLockPostBadRequest with default headers values
func NewSmartlockLockActionAdvancedResourcePostLockPostBadRequest() *SmartlockLockActionAdvancedResourcePostLockPostBadRequest {
	return &SmartlockLockActionAdvancedResourcePostLockPostBadRequest{}
}

/*
SmartlockLockActionAdvancedResourcePostLockPostBadRequest describes a response with status code 400, with default header values.

Bad Parameter
*/
type SmartlockLockActionAdvancedResourcePostLockPostBadRequest struct {
}

// IsSuccess returns true when this smartlock lock action advanced resource post lock post bad request response has a 2xx status code
func (o *SmartlockLockActionAdvancedResourcePostLockPostBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this smartlock lock action advanced resource post lock post bad request response has a 3xx status code
func (o *SmartlockLockActionAdvancedResourcePostLockPostBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this smartlock lock action advanced resource post lock post bad request response has a 4xx status code
func (o *SmartlockLockActionAdvancedResourcePostLockPostBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this smartlock lock action advanced resource post lock post bad request response has a 5xx status code
func (o *SmartlockLockActionAdvancedResourcePostLockPostBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this smartlock lock action advanced resource post lock post bad request response a status code equal to that given
func (o *SmartlockLockActionAdvancedResourcePostLockPostBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the smartlock lock action advanced resource post lock post bad request response
func (o *SmartlockLockActionAdvancedResourcePostLockPostBadRequest) Code() int {
	return 400
}

func (o *SmartlockLockActionAdvancedResourcePostLockPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /smartlock/{smartlockId}/action/lock/advanced][%d] smartlockLockActionAdvancedResourcePostLockPostBadRequest ", 400)
}

func (o *SmartlockLockActionAdvancedResourcePostLockPostBadRequest) String() string {
	return fmt.Sprintf("[POST /smartlock/{smartlockId}/action/lock/advanced][%d] smartlockLockActionAdvancedResourcePostLockPostBadRequest ", 400)
}

func (o *SmartlockLockActionAdvancedResourcePostLockPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSmartlockLockActionAdvancedResourcePostLockPostUnauthorized creates a SmartlockLockActionAdvancedResourcePostLockPostUnauthorized with default headers values
func NewSmartlockLockActionAdvancedResourcePostLockPostUnauthorized() *SmartlockLockActionAdvancedResourcePostLockPostUnauthorized {
	return &SmartlockLockActionAdvancedResourcePostLockPostUnauthorized{}
}

/*
SmartlockLockActionAdvancedResourcePostLockPostUnauthorized describes a response with status code 401, with default header values.

Not authorized
*/
type SmartlockLockActionAdvancedResourcePostLockPostUnauthorized struct {
}

// IsSuccess returns true when this smartlock lock action advanced resource post lock post unauthorized response has a 2xx status code
func (o *SmartlockLockActionAdvancedResourcePostLockPostUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this smartlock lock action advanced resource post lock post unauthorized response has a 3xx status code
func (o *SmartlockLockActionAdvancedResourcePostLockPostUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this smartlock lock action advanced resource post lock post unauthorized response has a 4xx status code
func (o *SmartlockLockActionAdvancedResourcePostLockPostUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this smartlock lock action advanced resource post lock post unauthorized response has a 5xx status code
func (o *SmartlockLockActionAdvancedResourcePostLockPostUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this smartlock lock action advanced resource post lock post unauthorized response a status code equal to that given
func (o *SmartlockLockActionAdvancedResourcePostLockPostUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the smartlock lock action advanced resource post lock post unauthorized response
func (o *SmartlockLockActionAdvancedResourcePostLockPostUnauthorized) Code() int {
	return 401
}

func (o *SmartlockLockActionAdvancedResourcePostLockPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /smartlock/{smartlockId}/action/lock/advanced][%d] smartlockLockActionAdvancedResourcePostLockPostUnauthorized ", 401)
}

func (o *SmartlockLockActionAdvancedResourcePostLockPostUnauthorized) String() string {
	return fmt.Sprintf("[POST /smartlock/{smartlockId}/action/lock/advanced][%d] smartlockLockActionAdvancedResourcePostLockPostUnauthorized ", 401)
}

func (o *SmartlockLockActionAdvancedResourcePostLockPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSmartlockLockActionAdvancedResourcePostLockPostMethodNotAllowed creates a SmartlockLockActionAdvancedResourcePostLockPostMethodNotAllowed with default headers values
func NewSmartlockLockActionAdvancedResourcePostLockPostMethodNotAllowed() *SmartlockLockActionAdvancedResourcePostLockPostMethodNotAllowed {
	return &SmartlockLockActionAdvancedResourcePostLockPostMethodNotAllowed{}
}

/*
SmartlockLockActionAdvancedResourcePostLockPostMethodNotAllowed describes a response with status code 405, with default header values.

Not allowed
*/
type SmartlockLockActionAdvancedResourcePostLockPostMethodNotAllowed struct {
}

// IsSuccess returns true when this smartlock lock action advanced resource post lock post method not allowed response has a 2xx status code
func (o *SmartlockLockActionAdvancedResourcePostLockPostMethodNotAllowed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this smartlock lock action advanced resource post lock post method not allowed response has a 3xx status code
func (o *SmartlockLockActionAdvancedResourcePostLockPostMethodNotAllowed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this smartlock lock action advanced resource post lock post method not allowed response has a 4xx status code
func (o *SmartlockLockActionAdvancedResourcePostLockPostMethodNotAllowed) IsClientError() bool {
	return true
}

// IsServerError returns true when this smartlock lock action advanced resource post lock post method not allowed response has a 5xx status code
func (o *SmartlockLockActionAdvancedResourcePostLockPostMethodNotAllowed) IsServerError() bool {
	return false
}

// IsCode returns true when this smartlock lock action advanced resource post lock post method not allowed response a status code equal to that given
func (o *SmartlockLockActionAdvancedResourcePostLockPostMethodNotAllowed) IsCode(code int) bool {
	return code == 405
}

// Code gets the status code for the smartlock lock action advanced resource post lock post method not allowed response
func (o *SmartlockLockActionAdvancedResourcePostLockPostMethodNotAllowed) Code() int {
	return 405
}

func (o *SmartlockLockActionAdvancedResourcePostLockPostMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /smartlock/{smartlockId}/action/lock/advanced][%d] smartlockLockActionAdvancedResourcePostLockPostMethodNotAllowed ", 405)
}

func (o *SmartlockLockActionAdvancedResourcePostLockPostMethodNotAllowed) String() string {
	return fmt.Sprintf("[POST /smartlock/{smartlockId}/action/lock/advanced][%d] smartlockLockActionAdvancedResourcePostLockPostMethodNotAllowed ", 405)
}

func (o *SmartlockLockActionAdvancedResourcePostLockPostMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}