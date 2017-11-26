// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017 The go-ovh Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetMeIPOrganisationOrganisationIDParams creates a new GetMeIPOrganisationOrganisationIDParams object
// with the default values initialized.
func NewGetMeIPOrganisationOrganisationIDParams() *GetMeIPOrganisationOrganisationIDParams {
	var ()
	return &GetMeIPOrganisationOrganisationIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMeIPOrganisationOrganisationIDParamsWithTimeout creates a new GetMeIPOrganisationOrganisationIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMeIPOrganisationOrganisationIDParamsWithTimeout(timeout time.Duration) *GetMeIPOrganisationOrganisationIDParams {
	var ()
	return &GetMeIPOrganisationOrganisationIDParams{

		timeout: timeout,
	}
}

// NewGetMeIPOrganisationOrganisationIDParamsWithContext creates a new GetMeIPOrganisationOrganisationIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMeIPOrganisationOrganisationIDParamsWithContext(ctx context.Context) *GetMeIPOrganisationOrganisationIDParams {
	var ()
	return &GetMeIPOrganisationOrganisationIDParams{

		Context: ctx,
	}
}

// NewGetMeIPOrganisationOrganisationIDParamsWithHTTPClient creates a new GetMeIPOrganisationOrganisationIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMeIPOrganisationOrganisationIDParamsWithHTTPClient(client *http.Client) *GetMeIPOrganisationOrganisationIDParams {
	var ()
	return &GetMeIPOrganisationOrganisationIDParams{
		HTTPClient: client,
	}
}

/*GetMeIPOrganisationOrganisationIDParams contains all the parameters to send to the API endpoint
for the get me IP organisation organisation ID operation typically these are written to a http.Request
*/
type GetMeIPOrganisationOrganisationIDParams struct {

	/*OrganisationID*/
	OrganisationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get me IP organisation organisation ID params
func (o *GetMeIPOrganisationOrganisationIDParams) WithTimeout(timeout time.Duration) *GetMeIPOrganisationOrganisationIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get me IP organisation organisation ID params
func (o *GetMeIPOrganisationOrganisationIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get me IP organisation organisation ID params
func (o *GetMeIPOrganisationOrganisationIDParams) WithContext(ctx context.Context) *GetMeIPOrganisationOrganisationIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get me IP organisation organisation ID params
func (o *GetMeIPOrganisationOrganisationIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get me IP organisation organisation ID params
func (o *GetMeIPOrganisationOrganisationIDParams) WithHTTPClient(client *http.Client) *GetMeIPOrganisationOrganisationIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get me IP organisation organisation ID params
func (o *GetMeIPOrganisationOrganisationIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganisationID adds the organisationID to the get me IP organisation organisation ID params
func (o *GetMeIPOrganisationOrganisationIDParams) WithOrganisationID(organisationID string) *GetMeIPOrganisationOrganisationIDParams {
	o.SetOrganisationID(organisationID)
	return o
}

// SetOrganisationID adds the organisationId to the get me IP organisation organisation ID params
func (o *GetMeIPOrganisationOrganisationIDParams) SetOrganisationID(organisationID string) {
	o.OrganisationID = organisationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetMeIPOrganisationOrganisationIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param organisationId
	if err := r.SetPathParam("organisationId", o.OrganisationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
