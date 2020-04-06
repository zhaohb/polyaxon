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

package organizations_v1

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

// NewOrganizationsV1CreateOrganizationMemberParams creates a new OrganizationsV1CreateOrganizationMemberParams object
// with the default values initialized.
func NewOrganizationsV1CreateOrganizationMemberParams() *OrganizationsV1CreateOrganizationMemberParams {
	var ()
	return &OrganizationsV1CreateOrganizationMemberParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOrganizationsV1CreateOrganizationMemberParamsWithTimeout creates a new OrganizationsV1CreateOrganizationMemberParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOrganizationsV1CreateOrganizationMemberParamsWithTimeout(timeout time.Duration) *OrganizationsV1CreateOrganizationMemberParams {
	var ()
	return &OrganizationsV1CreateOrganizationMemberParams{

		timeout: timeout,
	}
}

// NewOrganizationsV1CreateOrganizationMemberParamsWithContext creates a new OrganizationsV1CreateOrganizationMemberParams object
// with the default values initialized, and the ability to set a context for a request
func NewOrganizationsV1CreateOrganizationMemberParamsWithContext(ctx context.Context) *OrganizationsV1CreateOrganizationMemberParams {
	var ()
	return &OrganizationsV1CreateOrganizationMemberParams{

		Context: ctx,
	}
}

// NewOrganizationsV1CreateOrganizationMemberParamsWithHTTPClient creates a new OrganizationsV1CreateOrganizationMemberParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewOrganizationsV1CreateOrganizationMemberParamsWithHTTPClient(client *http.Client) *OrganizationsV1CreateOrganizationMemberParams {
	var ()
	return &OrganizationsV1CreateOrganizationMemberParams{
		HTTPClient: client,
	}
}

/*OrganizationsV1CreateOrganizationMemberParams contains all the parameters to send to the API endpoint
for the organizations v1 create organization member operation typically these are written to a http.Request
*/
type OrganizationsV1CreateOrganizationMemberParams struct {

	/*Body
	  Organization body

	*/
	Body *service_model.V1OrganizationMember
	/*Owner
	  Owner of the namespace

	*/
	Owner string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the organizations v1 create organization member params
func (o *OrganizationsV1CreateOrganizationMemberParams) WithTimeout(timeout time.Duration) *OrganizationsV1CreateOrganizationMemberParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the organizations v1 create organization member params
func (o *OrganizationsV1CreateOrganizationMemberParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the organizations v1 create organization member params
func (o *OrganizationsV1CreateOrganizationMemberParams) WithContext(ctx context.Context) *OrganizationsV1CreateOrganizationMemberParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the organizations v1 create organization member params
func (o *OrganizationsV1CreateOrganizationMemberParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the organizations v1 create organization member params
func (o *OrganizationsV1CreateOrganizationMemberParams) WithHTTPClient(client *http.Client) *OrganizationsV1CreateOrganizationMemberParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the organizations v1 create organization member params
func (o *OrganizationsV1CreateOrganizationMemberParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the organizations v1 create organization member params
func (o *OrganizationsV1CreateOrganizationMemberParams) WithBody(body *service_model.V1OrganizationMember) *OrganizationsV1CreateOrganizationMemberParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the organizations v1 create organization member params
func (o *OrganizationsV1CreateOrganizationMemberParams) SetBody(body *service_model.V1OrganizationMember) {
	o.Body = body
}

// WithOwner adds the owner to the organizations v1 create organization member params
func (o *OrganizationsV1CreateOrganizationMemberParams) WithOwner(owner string) *OrganizationsV1CreateOrganizationMemberParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the organizations v1 create organization member params
func (o *OrganizationsV1CreateOrganizationMemberParams) SetOwner(owner string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *OrganizationsV1CreateOrganizationMemberParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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