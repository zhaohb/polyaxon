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

package teams_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// DeleteTeamMemberReader is a Reader for the DeleteTeamMember structure.
type DeleteTeamMemberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTeamMemberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteTeamMemberOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteTeamMemberNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteTeamMemberForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteTeamMemberNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteTeamMemberDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteTeamMemberOK creates a DeleteTeamMemberOK with default headers values
func NewDeleteTeamMemberOK() *DeleteTeamMemberOK {
	return &DeleteTeamMemberOK{}
}

/*DeleteTeamMemberOK handles this case with default header values.

A successful response.
*/
type DeleteTeamMemberOK struct {
	Payload interface{}
}

func (o *DeleteTeamMemberOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/teams/{team}/members/{user}][%d] deleteTeamMemberOK  %+v", 200, o.Payload)
}

func (o *DeleteTeamMemberOK) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteTeamMemberOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTeamMemberNoContent creates a DeleteTeamMemberNoContent with default headers values
func NewDeleteTeamMemberNoContent() *DeleteTeamMemberNoContent {
	return &DeleteTeamMemberNoContent{}
}

/*DeleteTeamMemberNoContent handles this case with default header values.

No content.
*/
type DeleteTeamMemberNoContent struct {
	Payload interface{}
}

func (o *DeleteTeamMemberNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/teams/{team}/members/{user}][%d] deleteTeamMemberNoContent  %+v", 204, o.Payload)
}

func (o *DeleteTeamMemberNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteTeamMemberNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTeamMemberForbidden creates a DeleteTeamMemberForbidden with default headers values
func NewDeleteTeamMemberForbidden() *DeleteTeamMemberForbidden {
	return &DeleteTeamMemberForbidden{}
}

/*DeleteTeamMemberForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type DeleteTeamMemberForbidden struct {
	Payload interface{}
}

func (o *DeleteTeamMemberForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/teams/{team}/members/{user}][%d] deleteTeamMemberForbidden  %+v", 403, o.Payload)
}

func (o *DeleteTeamMemberForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteTeamMemberForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTeamMemberNotFound creates a DeleteTeamMemberNotFound with default headers values
func NewDeleteTeamMemberNotFound() *DeleteTeamMemberNotFound {
	return &DeleteTeamMemberNotFound{}
}

/*DeleteTeamMemberNotFound handles this case with default header values.

Resource does not exist.
*/
type DeleteTeamMemberNotFound struct {
	Payload interface{}
}

func (o *DeleteTeamMemberNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/teams/{team}/members/{user}][%d] deleteTeamMemberNotFound  %+v", 404, o.Payload)
}

func (o *DeleteTeamMemberNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteTeamMemberNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTeamMemberDefault creates a DeleteTeamMemberDefault with default headers values
func NewDeleteTeamMemberDefault(code int) *DeleteTeamMemberDefault {
	return &DeleteTeamMemberDefault{
		_statusCode: code,
	}
}

/*DeleteTeamMemberDefault handles this case with default header values.

An unexpected error response
*/
type DeleteTeamMemberDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the delete team member default response
func (o *DeleteTeamMemberDefault) Code() int {
	return o._statusCode
}

func (o *DeleteTeamMemberDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/teams/{team}/members/{user}][%d] DeleteTeamMember default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteTeamMemberDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *DeleteTeamMemberDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}