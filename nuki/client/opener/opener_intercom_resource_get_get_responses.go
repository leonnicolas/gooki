// Code generated by go-swagger; DO NOT EDIT.

package opener

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/leonnicolas/gooki/nuki/models"
)

// OpenerIntercomResourceGetGetReader is a Reader for the OpenerIntercomResourceGetGet structure.
type OpenerIntercomResourceGetGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OpenerIntercomResourceGetGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOpenerIntercomResourceGetGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /opener/intercom/{intercomId}] OpenerIntercomResource_get_get", response, response.Code())
	}
}

// NewOpenerIntercomResourceGetGetOK creates a OpenerIntercomResourceGetGetOK with default headers values
func NewOpenerIntercomResourceGetGetOK() *OpenerIntercomResourceGetGetOK {
	return &OpenerIntercomResourceGetGetOK{}
}

/*
OpenerIntercomResourceGetGetOK describes a response with status code 200, with default header values.

successful operation
*/
type OpenerIntercomResourceGetGetOK struct {
	Payload *models.OpenerIntercomModel
}

// IsSuccess returns true when this opener intercom resource get get o k response has a 2xx status code
func (o *OpenerIntercomResourceGetGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this opener intercom resource get get o k response has a 3xx status code
func (o *OpenerIntercomResourceGetGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this opener intercom resource get get o k response has a 4xx status code
func (o *OpenerIntercomResourceGetGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this opener intercom resource get get o k response has a 5xx status code
func (o *OpenerIntercomResourceGetGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this opener intercom resource get get o k response a status code equal to that given
func (o *OpenerIntercomResourceGetGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the opener intercom resource get get o k response
func (o *OpenerIntercomResourceGetGetOK) Code() int {
	return 200
}

func (o *OpenerIntercomResourceGetGetOK) Error() string {
	return fmt.Sprintf("[GET /opener/intercom/{intercomId}][%d] openerIntercomResourceGetGetOK  %+v", 200, o.Payload)
}

func (o *OpenerIntercomResourceGetGetOK) String() string {
	return fmt.Sprintf("[GET /opener/intercom/{intercomId}][%d] openerIntercomResourceGetGetOK  %+v", 200, o.Payload)
}

func (o *OpenerIntercomResourceGetGetOK) GetPayload() *models.OpenerIntercomModel {
	return o.Payload
}

func (o *OpenerIntercomResourceGetGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenerIntercomModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
