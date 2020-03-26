// Copyright 2018-2020 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by go-swagger; DO NOT EDIT.

package runs_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// InvalidateRunReader is a Reader for the InvalidateRun structure.
type InvalidateRunReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InvalidateRunReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInvalidateRunOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewInvalidateRunNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewInvalidateRunForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewInvalidateRunNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewInvalidateRunDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewInvalidateRunOK creates a InvalidateRunOK with default headers values
func NewInvalidateRunOK() *InvalidateRunOK {
	return &InvalidateRunOK{}
}

/*InvalidateRunOK handles this case with default header values.

A successful response.
*/
type InvalidateRunOK struct {
}

func (o *InvalidateRunOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/runs/{uuid}/invalidate][%d] invalidateRunOK ", 200)
}

func (o *InvalidateRunOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInvalidateRunNoContent creates a InvalidateRunNoContent with default headers values
func NewInvalidateRunNoContent() *InvalidateRunNoContent {
	return &InvalidateRunNoContent{}
}

/*InvalidateRunNoContent handles this case with default header values.

No content.
*/
type InvalidateRunNoContent struct {
	Payload interface{}
}

func (o *InvalidateRunNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/runs/{uuid}/invalidate][%d] invalidateRunNoContent  %+v", 204, o.Payload)
}

func (o *InvalidateRunNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *InvalidateRunNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInvalidateRunForbidden creates a InvalidateRunForbidden with default headers values
func NewInvalidateRunForbidden() *InvalidateRunForbidden {
	return &InvalidateRunForbidden{}
}

/*InvalidateRunForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type InvalidateRunForbidden struct {
	Payload interface{}
}

func (o *InvalidateRunForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/runs/{uuid}/invalidate][%d] invalidateRunForbidden  %+v", 403, o.Payload)
}

func (o *InvalidateRunForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *InvalidateRunForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInvalidateRunNotFound creates a InvalidateRunNotFound with default headers values
func NewInvalidateRunNotFound() *InvalidateRunNotFound {
	return &InvalidateRunNotFound{}
}

/*InvalidateRunNotFound handles this case with default header values.

Resource does not exist.
*/
type InvalidateRunNotFound struct {
	Payload interface{}
}

func (o *InvalidateRunNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/runs/{uuid}/invalidate][%d] invalidateRunNotFound  %+v", 404, o.Payload)
}

func (o *InvalidateRunNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *InvalidateRunNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInvalidateRunDefault creates a InvalidateRunDefault with default headers values
func NewInvalidateRunDefault(code int) *InvalidateRunDefault {
	return &InvalidateRunDefault{
		_statusCode: code,
	}
}

/*InvalidateRunDefault handles this case with default header values.

An unexpected error response
*/
type InvalidateRunDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the invalidate run default response
func (o *InvalidateRunDefault) Code() int {
	return o._statusCode
}

func (o *InvalidateRunDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/runs/{uuid}/invalidate][%d] InvalidateRun default  %+v", o._statusCode, o.Payload)
}

func (o *InvalidateRunDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *InvalidateRunDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}