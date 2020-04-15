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

package hub_components_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// NewUpdateHubComponentParams creates a new UpdateHubComponentParams object
// with the default values initialized.
func NewUpdateHubComponentParams() *UpdateHubComponentParams {
	var ()
	return &UpdateHubComponentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateHubComponentParamsWithTimeout creates a new UpdateHubComponentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateHubComponentParamsWithTimeout(timeout time.Duration) *UpdateHubComponentParams {
	var ()
	return &UpdateHubComponentParams{

		timeout: timeout,
	}
}

// NewUpdateHubComponentParamsWithContext creates a new UpdateHubComponentParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateHubComponentParamsWithContext(ctx context.Context) *UpdateHubComponentParams {
	var ()
	return &UpdateHubComponentParams{

		Context: ctx,
	}
}

// NewUpdateHubComponentParamsWithHTTPClient creates a new UpdateHubComponentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateHubComponentParamsWithHTTPClient(client *http.Client) *UpdateHubComponentParams {
	var ()
	return &UpdateHubComponentParams{
		HTTPClient: client,
	}
}

/*UpdateHubComponentParams contains all the parameters to send to the API endpoint
for the update hub component operation typically these are written to a http.Request
*/
type UpdateHubComponentParams struct {

	/*Body
	  Component body

	*/
	Body *service_model.V1HubComponent
	/*ComponentUUID
	  UUID

	*/
	ComponentUUID string
	/*Owner
	  Owner of the namespace

	*/
	Owner string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update hub component params
func (o *UpdateHubComponentParams) WithTimeout(timeout time.Duration) *UpdateHubComponentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update hub component params
func (o *UpdateHubComponentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update hub component params
func (o *UpdateHubComponentParams) WithContext(ctx context.Context) *UpdateHubComponentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update hub component params
func (o *UpdateHubComponentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update hub component params
func (o *UpdateHubComponentParams) WithHTTPClient(client *http.Client) *UpdateHubComponentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update hub component params
func (o *UpdateHubComponentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update hub component params
func (o *UpdateHubComponentParams) WithBody(body *service_model.V1HubComponent) *UpdateHubComponentParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update hub component params
func (o *UpdateHubComponentParams) SetBody(body *service_model.V1HubComponent) {
	o.Body = body
}

// WithComponentUUID adds the componentUUID to the update hub component params
func (o *UpdateHubComponentParams) WithComponentUUID(componentUUID string) *UpdateHubComponentParams {
	o.SetComponentUUID(componentUUID)
	return o
}

// SetComponentUUID adds the componentUuid to the update hub component params
func (o *UpdateHubComponentParams) SetComponentUUID(componentUUID string) {
	o.ComponentUUID = componentUUID
}

// WithOwner adds the owner to the update hub component params
func (o *UpdateHubComponentParams) WithOwner(owner string) *UpdateHubComponentParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the update hub component params
func (o *UpdateHubComponentParams) SetOwner(owner string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateHubComponentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param component.uuid
	if err := r.SetPathParam("component.uuid", o.ComponentUUID); err != nil {
		return err
	}

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}