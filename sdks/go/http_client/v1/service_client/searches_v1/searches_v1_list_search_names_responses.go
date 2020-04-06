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

package searches_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// SearchesV1ListSearchNamesReader is a Reader for the SearchesV1ListSearchNames structure.
type SearchesV1ListSearchNamesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchesV1ListSearchNamesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchesV1ListSearchNamesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewSearchesV1ListSearchNamesNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewSearchesV1ListSearchNamesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSearchesV1ListSearchNamesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewSearchesV1ListSearchNamesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSearchesV1ListSearchNamesOK creates a SearchesV1ListSearchNamesOK with default headers values
func NewSearchesV1ListSearchNamesOK() *SearchesV1ListSearchNamesOK {
	return &SearchesV1ListSearchNamesOK{}
}

/*SearchesV1ListSearchNamesOK handles this case with default header values.

A successful response.
*/
type SearchesV1ListSearchNamesOK struct {
	Payload *service_model.V1ListSearchesResponse
}

func (o *SearchesV1ListSearchNamesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/searches/names][%d] searchesV1ListSearchNamesOK  %+v", 200, o.Payload)
}

func (o *SearchesV1ListSearchNamesOK) GetPayload() *service_model.V1ListSearchesResponse {
	return o.Payload
}

func (o *SearchesV1ListSearchNamesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ListSearchesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchesV1ListSearchNamesNoContent creates a SearchesV1ListSearchNamesNoContent with default headers values
func NewSearchesV1ListSearchNamesNoContent() *SearchesV1ListSearchNamesNoContent {
	return &SearchesV1ListSearchNamesNoContent{}
}

/*SearchesV1ListSearchNamesNoContent handles this case with default header values.

No content.
*/
type SearchesV1ListSearchNamesNoContent struct {
	Payload interface{}
}

func (o *SearchesV1ListSearchNamesNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/searches/names][%d] searchesV1ListSearchNamesNoContent  %+v", 204, o.Payload)
}

func (o *SearchesV1ListSearchNamesNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *SearchesV1ListSearchNamesNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchesV1ListSearchNamesForbidden creates a SearchesV1ListSearchNamesForbidden with default headers values
func NewSearchesV1ListSearchNamesForbidden() *SearchesV1ListSearchNamesForbidden {
	return &SearchesV1ListSearchNamesForbidden{}
}

/*SearchesV1ListSearchNamesForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type SearchesV1ListSearchNamesForbidden struct {
	Payload interface{}
}

func (o *SearchesV1ListSearchNamesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/searches/names][%d] searchesV1ListSearchNamesForbidden  %+v", 403, o.Payload)
}

func (o *SearchesV1ListSearchNamesForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *SearchesV1ListSearchNamesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchesV1ListSearchNamesNotFound creates a SearchesV1ListSearchNamesNotFound with default headers values
func NewSearchesV1ListSearchNamesNotFound() *SearchesV1ListSearchNamesNotFound {
	return &SearchesV1ListSearchNamesNotFound{}
}

/*SearchesV1ListSearchNamesNotFound handles this case with default header values.

Resource does not exist.
*/
type SearchesV1ListSearchNamesNotFound struct {
	Payload interface{}
}

func (o *SearchesV1ListSearchNamesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/searches/names][%d] searchesV1ListSearchNamesNotFound  %+v", 404, o.Payload)
}

func (o *SearchesV1ListSearchNamesNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *SearchesV1ListSearchNamesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchesV1ListSearchNamesDefault creates a SearchesV1ListSearchNamesDefault with default headers values
func NewSearchesV1ListSearchNamesDefault(code int) *SearchesV1ListSearchNamesDefault {
	return &SearchesV1ListSearchNamesDefault{
		_statusCode: code,
	}
}

/*SearchesV1ListSearchNamesDefault handles this case with default header values.

An unexpected error response
*/
type SearchesV1ListSearchNamesDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the searches v1 list search names default response
func (o *SearchesV1ListSearchNamesDefault) Code() int {
	return o._statusCode
}

func (o *SearchesV1ListSearchNamesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/searches/names][%d] SearchesV1_ListSearchNames default  %+v", o._statusCode, o.Payload)
}

func (o *SearchesV1ListSearchNamesDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *SearchesV1ListSearchNamesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}