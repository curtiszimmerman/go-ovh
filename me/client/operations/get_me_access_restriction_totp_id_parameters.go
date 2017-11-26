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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetMeAccessRestrictionTotpIDParams creates a new GetMeAccessRestrictionTotpIDParams object
// with the default values initialized.
func NewGetMeAccessRestrictionTotpIDParams() *GetMeAccessRestrictionTotpIDParams {
	var ()
	return &GetMeAccessRestrictionTotpIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMeAccessRestrictionTotpIDParamsWithTimeout creates a new GetMeAccessRestrictionTotpIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMeAccessRestrictionTotpIDParamsWithTimeout(timeout time.Duration) *GetMeAccessRestrictionTotpIDParams {
	var ()
	return &GetMeAccessRestrictionTotpIDParams{

		timeout: timeout,
	}
}

// NewGetMeAccessRestrictionTotpIDParamsWithContext creates a new GetMeAccessRestrictionTotpIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMeAccessRestrictionTotpIDParamsWithContext(ctx context.Context) *GetMeAccessRestrictionTotpIDParams {
	var ()
	return &GetMeAccessRestrictionTotpIDParams{

		Context: ctx,
	}
}

// NewGetMeAccessRestrictionTotpIDParamsWithHTTPClient creates a new GetMeAccessRestrictionTotpIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMeAccessRestrictionTotpIDParamsWithHTTPClient(client *http.Client) *GetMeAccessRestrictionTotpIDParams {
	var ()
	return &GetMeAccessRestrictionTotpIDParams{
		HTTPClient: client,
	}
}

/*GetMeAccessRestrictionTotpIDParams contains all the parameters to send to the API endpoint
for the get me access restriction totp ID operation typically these are written to a http.Request
*/
type GetMeAccessRestrictionTotpIDParams struct {

	/*ID*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get me access restriction totp ID params
func (o *GetMeAccessRestrictionTotpIDParams) WithTimeout(timeout time.Duration) *GetMeAccessRestrictionTotpIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get me access restriction totp ID params
func (o *GetMeAccessRestrictionTotpIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get me access restriction totp ID params
func (o *GetMeAccessRestrictionTotpIDParams) WithContext(ctx context.Context) *GetMeAccessRestrictionTotpIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get me access restriction totp ID params
func (o *GetMeAccessRestrictionTotpIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get me access restriction totp ID params
func (o *GetMeAccessRestrictionTotpIDParams) WithHTTPClient(client *http.Client) *GetMeAccessRestrictionTotpIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get me access restriction totp ID params
func (o *GetMeAccessRestrictionTotpIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get me access restriction totp ID params
func (o *GetMeAccessRestrictionTotpIDParams) WithID(id int64) *GetMeAccessRestrictionTotpIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get me access restriction totp ID params
func (o *GetMeAccessRestrictionTotpIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetMeAccessRestrictionTotpIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
