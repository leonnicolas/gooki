// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/leonnicolas/gooki/nuki/models"
)

// ServiceResourceGetGetReader is a Reader for the ServiceResourceGetGet structure.
type ServiceResourceGetGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceResourceGetGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServiceResourceGetGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewServiceResourceGetGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /service/{serviceId}] ServiceResource_get_get", response, response.Code())
	}
}

// NewServiceResourceGetGetOK creates a ServiceResourceGetGetOK with default headers values
func NewServiceResourceGetGetOK() *ServiceResourceGetGetOK {
	return &ServiceResourceGetGetOK{}
}

/*
ServiceResourceGetGetOK describes a response with status code 200, with default header values.

successful operation
*/
type ServiceResourceGetGetOK struct {
	Payload *models.Service
}

// IsSuccess returns true when this service resource get get o k response has a 2xx status code
func (o *ServiceResourceGetGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this service resource get get o k response has a 3xx status code
func (o *ServiceResourceGetGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service resource get get o k response has a 4xx status code
func (o *ServiceResourceGetGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this service resource get get o k response has a 5xx status code
func (o *ServiceResourceGetGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this service resource get get o k response a status code equal to that given
func (o *ServiceResourceGetGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the service resource get get o k response
func (o *ServiceResourceGetGetOK) Code() int {
	return 200
}

func (o *ServiceResourceGetGetOK) Error() string {
	return fmt.Sprintf("[GET /service/{serviceId}][%d] serviceResourceGetGetOK  %+v", 200, o.Payload)
}

func (o *ServiceResourceGetGetOK) String() string {
	return fmt.Sprintf("[GET /service/{serviceId}][%d] serviceResourceGetGetOK  %+v", 200, o.Payload)
}

func (o *ServiceResourceGetGetOK) GetPayload() *models.Service {
	return o.Payload
}

func (o *ServiceResourceGetGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Service)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceResourceGetGetUnauthorized creates a ServiceResourceGetGetUnauthorized with default headers values
func NewServiceResourceGetGetUnauthorized() *ServiceResourceGetGetUnauthorized {
	return &ServiceResourceGetGetUnauthorized{}
}

/*
ServiceResourceGetGetUnauthorized describes a response with status code 401, with default header values.

Not authorized
*/
type ServiceResourceGetGetUnauthorized struct {
}

// IsSuccess returns true when this service resource get get unauthorized response has a 2xx status code
func (o *ServiceResourceGetGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service resource get get unauthorized response has a 3xx status code
func (o *ServiceResourceGetGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service resource get get unauthorized response has a 4xx status code
func (o *ServiceResourceGetGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this service resource get get unauthorized response has a 5xx status code
func (o *ServiceResourceGetGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this service resource get get unauthorized response a status code equal to that given
func (o *ServiceResourceGetGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the service resource get get unauthorized response
func (o *ServiceResourceGetGetUnauthorized) Code() int {
	return 401
}

func (o *ServiceResourceGetGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /service/{serviceId}][%d] serviceResourceGetGetUnauthorized ", 401)
}

func (o *ServiceResourceGetGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /service/{serviceId}][%d] serviceResourceGetGetUnauthorized ", 401)
}

func (o *ServiceResourceGetGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}