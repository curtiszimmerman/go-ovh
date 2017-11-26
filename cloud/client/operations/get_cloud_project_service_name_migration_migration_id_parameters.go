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

// NewGetCloudProjectServiceNameMigrationMigrationIDParams creates a new GetCloudProjectServiceNameMigrationMigrationIDParams object
// with the default values initialized.
func NewGetCloudProjectServiceNameMigrationMigrationIDParams() *GetCloudProjectServiceNameMigrationMigrationIDParams {
	var ()
	return &GetCloudProjectServiceNameMigrationMigrationIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCloudProjectServiceNameMigrationMigrationIDParamsWithTimeout creates a new GetCloudProjectServiceNameMigrationMigrationIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCloudProjectServiceNameMigrationMigrationIDParamsWithTimeout(timeout time.Duration) *GetCloudProjectServiceNameMigrationMigrationIDParams {
	var ()
	return &GetCloudProjectServiceNameMigrationMigrationIDParams{

		timeout: timeout,
	}
}

// NewGetCloudProjectServiceNameMigrationMigrationIDParamsWithContext creates a new GetCloudProjectServiceNameMigrationMigrationIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCloudProjectServiceNameMigrationMigrationIDParamsWithContext(ctx context.Context) *GetCloudProjectServiceNameMigrationMigrationIDParams {
	var ()
	return &GetCloudProjectServiceNameMigrationMigrationIDParams{

		Context: ctx,
	}
}

// NewGetCloudProjectServiceNameMigrationMigrationIDParamsWithHTTPClient creates a new GetCloudProjectServiceNameMigrationMigrationIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCloudProjectServiceNameMigrationMigrationIDParamsWithHTTPClient(client *http.Client) *GetCloudProjectServiceNameMigrationMigrationIDParams {
	var ()
	return &GetCloudProjectServiceNameMigrationMigrationIDParams{
		HTTPClient: client,
	}
}

/*GetCloudProjectServiceNameMigrationMigrationIDParams contains all the parameters to send to the API endpoint
for the get cloud project service name migration migration ID operation typically these are written to a http.Request
*/
type GetCloudProjectServiceNameMigrationMigrationIDParams struct {

	/*MigrationID*/
	MigrationID string
	/*ServiceName*/
	ServiceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get cloud project service name migration migration ID params
func (o *GetCloudProjectServiceNameMigrationMigrationIDParams) WithTimeout(timeout time.Duration) *GetCloudProjectServiceNameMigrationMigrationIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cloud project service name migration migration ID params
func (o *GetCloudProjectServiceNameMigrationMigrationIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cloud project service name migration migration ID params
func (o *GetCloudProjectServiceNameMigrationMigrationIDParams) WithContext(ctx context.Context) *GetCloudProjectServiceNameMigrationMigrationIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cloud project service name migration migration ID params
func (o *GetCloudProjectServiceNameMigrationMigrationIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cloud project service name migration migration ID params
func (o *GetCloudProjectServiceNameMigrationMigrationIDParams) WithHTTPClient(client *http.Client) *GetCloudProjectServiceNameMigrationMigrationIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cloud project service name migration migration ID params
func (o *GetCloudProjectServiceNameMigrationMigrationIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMigrationID adds the migrationID to the get cloud project service name migration migration ID params
func (o *GetCloudProjectServiceNameMigrationMigrationIDParams) WithMigrationID(migrationID string) *GetCloudProjectServiceNameMigrationMigrationIDParams {
	o.SetMigrationID(migrationID)
	return o
}

// SetMigrationID adds the migrationId to the get cloud project service name migration migration ID params
func (o *GetCloudProjectServiceNameMigrationMigrationIDParams) SetMigrationID(migrationID string) {
	o.MigrationID = migrationID
}

// WithServiceName adds the serviceName to the get cloud project service name migration migration ID params
func (o *GetCloudProjectServiceNameMigrationMigrationIDParams) WithServiceName(serviceName string) *GetCloudProjectServiceNameMigrationMigrationIDParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the get cloud project service name migration migration ID params
func (o *GetCloudProjectServiceNameMigrationMigrationIDParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WriteToRequest writes these params to a swagger request
func (o *GetCloudProjectServiceNameMigrationMigrationIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param migrationId
	if err := r.SetPathParam("migrationId", o.MigrationID); err != nil {
		return err
	}

	// path param serviceName
	if err := r.SetPathParam("serviceName", o.ServiceName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}