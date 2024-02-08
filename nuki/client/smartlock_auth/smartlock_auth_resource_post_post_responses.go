// Code generated by go-swagger; DO NOT EDIT.

package smartlock_auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SmartlockAuthResourcePostPostReader is a Reader for the SmartlockAuthResourcePostPost structure.
type SmartlockAuthResourcePostPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SmartlockAuthResourcePostPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewSmartlockAuthResourcePostPostNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSmartlockAuthResourcePostPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSmartlockAuthResourcePostPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSmartlockAuthResourcePostPostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewSmartlockAuthResourcePostPostConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 423:
		result := NewSmartlockAuthResourcePostPostLocked()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /smartlock/{smartlockId}/auth/{id}] SmartlockAuthResource_post_post", response, response.Code())
	}
}

// NewSmartlockAuthResourcePostPostNoContent creates a SmartlockAuthResourcePostPostNoContent with default headers values
func NewSmartlockAuthResourcePostPostNoContent() *SmartlockAuthResourcePostPostNoContent {
	return &SmartlockAuthResourcePostPostNoContent{}
}

/*
SmartlockAuthResourcePostPostNoContent describes a response with status code 204, with default header values.

Ok
*/
type SmartlockAuthResourcePostPostNoContent struct {
}

// IsSuccess returns true when this smartlock auth resource post post no content response has a 2xx status code
func (o *SmartlockAuthResourcePostPostNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this smartlock auth resource post post no content response has a 3xx status code
func (o *SmartlockAuthResourcePostPostNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this smartlock auth resource post post no content response has a 4xx status code
func (o *SmartlockAuthResourcePostPostNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this smartlock auth resource post post no content response has a 5xx status code
func (o *SmartlockAuthResourcePostPostNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this smartlock auth resource post post no content response a status code equal to that given
func (o *SmartlockAuthResourcePostPostNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the smartlock auth resource post post no content response
func (o *SmartlockAuthResourcePostPostNoContent) Code() int {
	return 204
}

func (o *SmartlockAuthResourcePostPostNoContent) Error() string {
	return fmt.Sprintf("[POST /smartlock/{smartlockId}/auth/{id}][%d] smartlockAuthResourcePostPostNoContent ", 204)
}

func (o *SmartlockAuthResourcePostPostNoContent) String() string {
	return fmt.Sprintf("[POST /smartlock/{smartlockId}/auth/{id}][%d] smartlockAuthResourcePostPostNoContent ", 204)
}

func (o *SmartlockAuthResourcePostPostNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSmartlockAuthResourcePostPostBadRequest creates a SmartlockAuthResourcePostPostBadRequest with default headers values
func NewSmartlockAuthResourcePostPostBadRequest() *SmartlockAuthResourcePostPostBadRequest {
	return &SmartlockAuthResourcePostPostBadRequest{}
}

/*
SmartlockAuthResourcePostPostBadRequest describes a response with status code 400, with default header values.

Bad parameter
*/
type SmartlockAuthResourcePostPostBadRequest struct {
}

// IsSuccess returns true when this smartlock auth resource post post bad request response has a 2xx status code
func (o *SmartlockAuthResourcePostPostBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this smartlock auth resource post post bad request response has a 3xx status code
func (o *SmartlockAuthResourcePostPostBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this smartlock auth resource post post bad request response has a 4xx status code
func (o *SmartlockAuthResourcePostPostBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this smartlock auth resource post post bad request response has a 5xx status code
func (o *SmartlockAuthResourcePostPostBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this smartlock auth resource post post bad request response a status code equal to that given
func (o *SmartlockAuthResourcePostPostBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the smartlock auth resource post post bad request response
func (o *SmartlockAuthResourcePostPostBadRequest) Code() int {
	return 400
}

func (o *SmartlockAuthResourcePostPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /smartlock/{smartlockId}/auth/{id}][%d] smartlockAuthResourcePostPostBadRequest ", 400)
}

func (o *SmartlockAuthResourcePostPostBadRequest) String() string {
	return fmt.Sprintf("[POST /smartlock/{smartlockId}/auth/{id}][%d] smartlockAuthResourcePostPostBadRequest ", 400)
}

func (o *SmartlockAuthResourcePostPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSmartlockAuthResourcePostPostUnauthorized creates a SmartlockAuthResourcePostPostUnauthorized with default headers values
func NewSmartlockAuthResourcePostPostUnauthorized() *SmartlockAuthResourcePostPostUnauthorized {
	return &SmartlockAuthResourcePostPostUnauthorized{}
}

/*
SmartlockAuthResourcePostPostUnauthorized describes a response with status code 401, with default header values.

Not authorized
*/
type SmartlockAuthResourcePostPostUnauthorized struct {
}

// IsSuccess returns true when this smartlock auth resource post post unauthorized response has a 2xx status code
func (o *SmartlockAuthResourcePostPostUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this smartlock auth resource post post unauthorized response has a 3xx status code
func (o *SmartlockAuthResourcePostPostUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this smartlock auth resource post post unauthorized response has a 4xx status code
func (o *SmartlockAuthResourcePostPostUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this smartlock auth resource post post unauthorized response has a 5xx status code
func (o *SmartlockAuthResourcePostPostUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this smartlock auth resource post post unauthorized response a status code equal to that given
func (o *SmartlockAuthResourcePostPostUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the smartlock auth resource post post unauthorized response
func (o *SmartlockAuthResourcePostPostUnauthorized) Code() int {
	return 401
}

func (o *SmartlockAuthResourcePostPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /smartlock/{smartlockId}/auth/{id}][%d] smartlockAuthResourcePostPostUnauthorized ", 401)
}

func (o *SmartlockAuthResourcePostPostUnauthorized) String() string {
	return fmt.Sprintf("[POST /smartlock/{smartlockId}/auth/{id}][%d] smartlockAuthResourcePostPostUnauthorized ", 401)
}

func (o *SmartlockAuthResourcePostPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSmartlockAuthResourcePostPostForbidden creates a SmartlockAuthResourcePostPostForbidden with default headers values
func NewSmartlockAuthResourcePostPostForbidden() *SmartlockAuthResourcePostPostForbidden {
	return &SmartlockAuthResourcePostPostForbidden{}
}

/*
SmartlockAuthResourcePostPostForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type SmartlockAuthResourcePostPostForbidden struct {
}

// IsSuccess returns true when this smartlock auth resource post post forbidden response has a 2xx status code
func (o *SmartlockAuthResourcePostPostForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this smartlock auth resource post post forbidden response has a 3xx status code
func (o *SmartlockAuthResourcePostPostForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this smartlock auth resource post post forbidden response has a 4xx status code
func (o *SmartlockAuthResourcePostPostForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this smartlock auth resource post post forbidden response has a 5xx status code
func (o *SmartlockAuthResourcePostPostForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this smartlock auth resource post post forbidden response a status code equal to that given
func (o *SmartlockAuthResourcePostPostForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the smartlock auth resource post post forbidden response
func (o *SmartlockAuthResourcePostPostForbidden) Code() int {
	return 403
}

func (o *SmartlockAuthResourcePostPostForbidden) Error() string {
	return fmt.Sprintf("[POST /smartlock/{smartlockId}/auth/{id}][%d] smartlockAuthResourcePostPostForbidden ", 403)
}

func (o *SmartlockAuthResourcePostPostForbidden) String() string {
	return fmt.Sprintf("[POST /smartlock/{smartlockId}/auth/{id}][%d] smartlockAuthResourcePostPostForbidden ", 403)
}

func (o *SmartlockAuthResourcePostPostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSmartlockAuthResourcePostPostConflict creates a SmartlockAuthResourcePostPostConflict with default headers values
func NewSmartlockAuthResourcePostPostConflict() *SmartlockAuthResourcePostPostConflict {
	return &SmartlockAuthResourcePostPostConflict{}
}

/*
SmartlockAuthResourcePostPostConflict describes a response with status code 409, with default header values.

Parameter conflicts
*/
type SmartlockAuthResourcePostPostConflict struct {
}

// IsSuccess returns true when this smartlock auth resource post post conflict response has a 2xx status code
func (o *SmartlockAuthResourcePostPostConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this smartlock auth resource post post conflict response has a 3xx status code
func (o *SmartlockAuthResourcePostPostConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this smartlock auth resource post post conflict response has a 4xx status code
func (o *SmartlockAuthResourcePostPostConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this smartlock auth resource post post conflict response has a 5xx status code
func (o *SmartlockAuthResourcePostPostConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this smartlock auth resource post post conflict response a status code equal to that given
func (o *SmartlockAuthResourcePostPostConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the smartlock auth resource post post conflict response
func (o *SmartlockAuthResourcePostPostConflict) Code() int {
	return 409
}

func (o *SmartlockAuthResourcePostPostConflict) Error() string {
	return fmt.Sprintf("[POST /smartlock/{smartlockId}/auth/{id}][%d] smartlockAuthResourcePostPostConflict ", 409)
}

func (o *SmartlockAuthResourcePostPostConflict) String() string {
	return fmt.Sprintf("[POST /smartlock/{smartlockId}/auth/{id}][%d] smartlockAuthResourcePostPostConflict ", 409)
}

func (o *SmartlockAuthResourcePostPostConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSmartlockAuthResourcePostPostLocked creates a SmartlockAuthResourcePostPostLocked with default headers values
func NewSmartlockAuthResourcePostPostLocked() *SmartlockAuthResourcePostPostLocked {
	return &SmartlockAuthResourcePostPostLocked{}
}

/*
SmartlockAuthResourcePostPostLocked describes a response with status code 423, with default header values.

Locked
*/
type SmartlockAuthResourcePostPostLocked struct {
}

// IsSuccess returns true when this smartlock auth resource post post locked response has a 2xx status code
func (o *SmartlockAuthResourcePostPostLocked) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this smartlock auth resource post post locked response has a 3xx status code
func (o *SmartlockAuthResourcePostPostLocked) IsRedirect() bool {
	return false
}

// IsClientError returns true when this smartlock auth resource post post locked response has a 4xx status code
func (o *SmartlockAuthResourcePostPostLocked) IsClientError() bool {
	return true
}

// IsServerError returns true when this smartlock auth resource post post locked response has a 5xx status code
func (o *SmartlockAuthResourcePostPostLocked) IsServerError() bool {
	return false
}

// IsCode returns true when this smartlock auth resource post post locked response a status code equal to that given
func (o *SmartlockAuthResourcePostPostLocked) IsCode(code int) bool {
	return code == 423
}

// Code gets the status code for the smartlock auth resource post post locked response
func (o *SmartlockAuthResourcePostPostLocked) Code() int {
	return 423
}

func (o *SmartlockAuthResourcePostPostLocked) Error() string {
	return fmt.Sprintf("[POST /smartlock/{smartlockId}/auth/{id}][%d] smartlockAuthResourcePostPostLocked ", 423)
}

func (o *SmartlockAuthResourcePostPostLocked) String() string {
	return fmt.Sprintf("[POST /smartlock/{smartlockId}/auth/{id}][%d] smartlockAuthResourcePostPostLocked ", 423)
}

func (o *SmartlockAuthResourcePostPostLocked) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
