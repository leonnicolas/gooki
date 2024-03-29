// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ServiceSyncResourcePostPostReader is a Reader for the ServiceSyncResourcePostPost structure.
type ServiceSyncResourcePostPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceSyncResourcePostPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewServiceSyncResourcePostPostNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewServiceSyncResourcePostPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewServiceSyncResourcePostPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /service/{serviceId}/sync] ServiceSyncResource_post_post", response, response.Code())
	}
}

// NewServiceSyncResourcePostPostNoContent creates a ServiceSyncResourcePostPostNoContent with default headers values
func NewServiceSyncResourcePostPostNoContent() *ServiceSyncResourcePostPostNoContent {
	return &ServiceSyncResourcePostPostNoContent{}
}

/*
ServiceSyncResourcePostPostNoContent describes a response with status code 204, with default header values.

Ok
*/
type ServiceSyncResourcePostPostNoContent struct {
}

// IsSuccess returns true when this service sync resource post post no content response has a 2xx status code
func (o *ServiceSyncResourcePostPostNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this service sync resource post post no content response has a 3xx status code
func (o *ServiceSyncResourcePostPostNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service sync resource post post no content response has a 4xx status code
func (o *ServiceSyncResourcePostPostNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this service sync resource post post no content response has a 5xx status code
func (o *ServiceSyncResourcePostPostNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this service sync resource post post no content response a status code equal to that given
func (o *ServiceSyncResourcePostPostNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the service sync resource post post no content response
func (o *ServiceSyncResourcePostPostNoContent) Code() int {
	return 204
}

func (o *ServiceSyncResourcePostPostNoContent) Error() string {
	return fmt.Sprintf("[POST /service/{serviceId}/sync][%d] serviceSyncResourcePostPostNoContent ", 204)
}

func (o *ServiceSyncResourcePostPostNoContent) String() string {
	return fmt.Sprintf("[POST /service/{serviceId}/sync][%d] serviceSyncResourcePostPostNoContent ", 204)
}

func (o *ServiceSyncResourcePostPostNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewServiceSyncResourcePostPostBadRequest creates a ServiceSyncResourcePostPostBadRequest with default headers values
func NewServiceSyncResourcePostPostBadRequest() *ServiceSyncResourcePostPostBadRequest {
	return &ServiceSyncResourcePostPostBadRequest{}
}

/*
ServiceSyncResourcePostPostBadRequest describes a response with status code 400, with default header values.

Bad Parameter
*/
type ServiceSyncResourcePostPostBadRequest struct {
}

// IsSuccess returns true when this service sync resource post post bad request response has a 2xx status code
func (o *ServiceSyncResourcePostPostBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service sync resource post post bad request response has a 3xx status code
func (o *ServiceSyncResourcePostPostBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service sync resource post post bad request response has a 4xx status code
func (o *ServiceSyncResourcePostPostBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this service sync resource post post bad request response has a 5xx status code
func (o *ServiceSyncResourcePostPostBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this service sync resource post post bad request response a status code equal to that given
func (o *ServiceSyncResourcePostPostBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the service sync resource post post bad request response
func (o *ServiceSyncResourcePostPostBadRequest) Code() int {
	return 400
}

func (o *ServiceSyncResourcePostPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /service/{serviceId}/sync][%d] serviceSyncResourcePostPostBadRequest ", 400)
}

func (o *ServiceSyncResourcePostPostBadRequest) String() string {
	return fmt.Sprintf("[POST /service/{serviceId}/sync][%d] serviceSyncResourcePostPostBadRequest ", 400)
}

func (o *ServiceSyncResourcePostPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewServiceSyncResourcePostPostUnauthorized creates a ServiceSyncResourcePostPostUnauthorized with default headers values
func NewServiceSyncResourcePostPostUnauthorized() *ServiceSyncResourcePostPostUnauthorized {
	return &ServiceSyncResourcePostPostUnauthorized{}
}

/*
ServiceSyncResourcePostPostUnauthorized describes a response with status code 401, with default header values.

Not authorized
*/
type ServiceSyncResourcePostPostUnauthorized struct {
}

// IsSuccess returns true when this service sync resource post post unauthorized response has a 2xx status code
func (o *ServiceSyncResourcePostPostUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service sync resource post post unauthorized response has a 3xx status code
func (o *ServiceSyncResourcePostPostUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service sync resource post post unauthorized response has a 4xx status code
func (o *ServiceSyncResourcePostPostUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this service sync resource post post unauthorized response has a 5xx status code
func (o *ServiceSyncResourcePostPostUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this service sync resource post post unauthorized response a status code equal to that given
func (o *ServiceSyncResourcePostPostUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the service sync resource post post unauthorized response
func (o *ServiceSyncResourcePostPostUnauthorized) Code() int {
	return 401
}

func (o *ServiceSyncResourcePostPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /service/{serviceId}/sync][%d] serviceSyncResourcePostPostUnauthorized ", 401)
}

func (o *ServiceSyncResourcePostPostUnauthorized) String() string {
	return fmt.Sprintf("[POST /service/{serviceId}/sync][%d] serviceSyncResourcePostPostUnauthorized ", 401)
}

func (o *ServiceSyncResourcePostPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
