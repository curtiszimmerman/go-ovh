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

// NewGetVpsServiceNameTemplatesIDSoftwareSoftwareIDParams creates a new GetVpsServiceNameTemplatesIDSoftwareSoftwareIDParams object
// with the default values initialized.
func NewGetVpsServiceNameTemplatesIDSoftwareSoftwareIDParams() *GetVpsServiceNameTemplatesIDSoftwareSoftwareIDParams {
	var ()
	return &GetVpsServiceNameTemplatesIDSoftwareSoftwareIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetVpsServiceNameTemplatesIDSoftwareSoftwareIDParamsWithTimeout creates a new GetVpsServiceNameTemplatesIDSoftwareSoftwareIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetVpsServiceNameTemplatesIDSoftwareSoftwareIDParamsWithTimeout(timeout time.Duration) *GetVpsServiceNameTemplatesIDSoftwareSoftwareIDParams {
	var ()
	return &GetVpsServiceNameTemplatesIDSoftwareSoftwareIDParams{

		timeout: timeout,
	}
}

// NewGetVpsServiceNameTemplatesIDSoftwareSoftwareIDParamsWithContext creates a new GetVpsServiceNameTemplatesIDSoftwareSoftwareIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetVpsServiceNameTemplatesIDSoftwareSoftwareIDParamsWithContext(ctx context.Context) *GetVpsServiceNameTemplatesIDSoftwareSoftwareIDParams {
	var ()
	return &GetVpsServiceNameTemplatesIDSoftwareSoftwareIDParams{

		Context: ctx,
	}
}

// NewGetVpsServiceNameTemplatesIDSoftwareSoftwareIDParamsWithHTTPClient creates a new GetVpsServiceNameTemplatesIDSoftwareSoftwareIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetVpsServiceNameTemplatesIDSoftwareSoftwareIDParamsWithHTTPClient(client *http.Client) *GetVpsServiceNameTemplatesIDSoftwareSoftwareIDParams {
	var ()
	return &GetVpsServiceNameTemplatesIDSoftwareSoftwareIDParams{
		HTTPClient: client,
	}
}

/*GetVpsServiceNameTemplatesIDSoftwareSoftwareIDParams contains all the parameters to send to the API endpoint
for the get vps service name templates ID software software ID operation typically these are written to a http.Request
*/
type GetVpsServiceNameTemplatesIDSoftwareSoftwareIDParams struct {

	/*ID*/
	ID int64
	/*ServiceName*/
	ServiceName string
	/*SoftwareID*/
	SoftwareID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get vps service name templates ID software software ID params
func (o *GetVpsServiceNameTemplatesIDSoftwareSoftwareIDParams) WithTimeout(timeout time.Duration) *GetVpsServiceNameTemplatesIDSoftwareSoftwareIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get vps service name templates ID software software ID params
func (o *GetVpsServiceNameTemplatesIDSoftwareSoftwareIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get vps service name templates ID software software ID params
func (o *GetVpsServiceNameTemplatesIDSoftwareSoftwareIDParams) WithContext(ctx context.Context) *GetVpsServiceNameTemplatesIDSoftwareSoftwareIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get vps service name templates ID software software ID params
func (o *GetVpsServiceNameTemplatesIDSoftwareSoftwareIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get vps service name templates ID software software ID params
func (o *GetVpsServiceNameTemplatesIDSoftwareSoftwareIDParams) WithHTTPClient(client *http.Client) *GetVpsServiceNameTemplatesIDSoftwareSoftwareIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get vps service name templates ID software software ID params
func (o *GetVpsServiceNameTemplatesIDSoftwareSoftwareIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get vps service name templates ID software software ID params
func (o *GetVpsServiceNameTemplatesIDSoftwareSoftwareIDParams) WithID(id int64) *GetVpsServiceNameTemplatesIDSoftwareSoftwareIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get vps service name templates ID software software ID params
func (o *GetVpsServiceNameTemplatesIDSoftwareSoftwareIDParams) SetID(id int64) {
	o.ID = id
}

// WithServiceName adds the serviceName to the get vps service name templates ID software software ID params
func (o *GetVpsServiceNameTemplatesIDSoftwareSoftwareIDParams) WithServiceName(serviceName string) *GetVpsServiceNameTemplatesIDSoftwareSoftwareIDParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the get vps service name templates ID software software ID params
func (o *GetVpsServiceNameTemplatesIDSoftwareSoftwareIDParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WithSoftwareID adds the softwareID to the get vps service name templates ID software software ID params
func (o *GetVpsServiceNameTemplatesIDSoftwareSoftwareIDParams) WithSoftwareID(softwareID int64) *GetVpsServiceNameTemplatesIDSoftwareSoftwareIDParams {
	o.SetSoftwareID(softwareID)
	return o
}

// SetSoftwareID adds the softwareId to the get vps service name templates ID software software ID params
func (o *GetVpsServiceNameTemplatesIDSoftwareSoftwareIDParams) SetSoftwareID(softwareID int64) {
	o.SoftwareID = softwareID
}

// WriteToRequest writes these params to a swagger request
func (o *GetVpsServiceNameTemplatesIDSoftwareSoftwareIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	// path param serviceName
	if err := r.SetPathParam("serviceName", o.ServiceName); err != nil {
		return err
	}

	// path param softwareId
	if err := r.SetPathParam("softwareId", swag.FormatInt64(o.SoftwareID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
