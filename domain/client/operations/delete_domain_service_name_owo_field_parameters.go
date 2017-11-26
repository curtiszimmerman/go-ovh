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

// NewDeleteDomainServiceNameOwoFieldParams creates a new DeleteDomainServiceNameOwoFieldParams object
// with the default values initialized.
func NewDeleteDomainServiceNameOwoFieldParams() *DeleteDomainServiceNameOwoFieldParams {
	var ()
	return &DeleteDomainServiceNameOwoFieldParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDomainServiceNameOwoFieldParamsWithTimeout creates a new DeleteDomainServiceNameOwoFieldParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteDomainServiceNameOwoFieldParamsWithTimeout(timeout time.Duration) *DeleteDomainServiceNameOwoFieldParams {
	var ()
	return &DeleteDomainServiceNameOwoFieldParams{

		timeout: timeout,
	}
}

// NewDeleteDomainServiceNameOwoFieldParamsWithContext creates a new DeleteDomainServiceNameOwoFieldParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteDomainServiceNameOwoFieldParamsWithContext(ctx context.Context) *DeleteDomainServiceNameOwoFieldParams {
	var ()
	return &DeleteDomainServiceNameOwoFieldParams{

		Context: ctx,
	}
}

// NewDeleteDomainServiceNameOwoFieldParamsWithHTTPClient creates a new DeleteDomainServiceNameOwoFieldParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteDomainServiceNameOwoFieldParamsWithHTTPClient(client *http.Client) *DeleteDomainServiceNameOwoFieldParams {
	var ()
	return &DeleteDomainServiceNameOwoFieldParams{
		HTTPClient: client,
	}
}

/*DeleteDomainServiceNameOwoFieldParams contains all the parameters to send to the API endpoint
for the delete domain service name owo field operation typically these are written to a http.Request
*/
type DeleteDomainServiceNameOwoFieldParams struct {

	/*Field*/
	Field string
	/*ServiceName*/
	ServiceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete domain service name owo field params
func (o *DeleteDomainServiceNameOwoFieldParams) WithTimeout(timeout time.Duration) *DeleteDomainServiceNameOwoFieldParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete domain service name owo field params
func (o *DeleteDomainServiceNameOwoFieldParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete domain service name owo field params
func (o *DeleteDomainServiceNameOwoFieldParams) WithContext(ctx context.Context) *DeleteDomainServiceNameOwoFieldParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete domain service name owo field params
func (o *DeleteDomainServiceNameOwoFieldParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete domain service name owo field params
func (o *DeleteDomainServiceNameOwoFieldParams) WithHTTPClient(client *http.Client) *DeleteDomainServiceNameOwoFieldParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete domain service name owo field params
func (o *DeleteDomainServiceNameOwoFieldParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithField adds the field to the delete domain service name owo field params
func (o *DeleteDomainServiceNameOwoFieldParams) WithField(field string) *DeleteDomainServiceNameOwoFieldParams {
	o.SetField(field)
	return o
}

// SetField adds the field to the delete domain service name owo field params
func (o *DeleteDomainServiceNameOwoFieldParams) SetField(field string) {
	o.Field = field
}

// WithServiceName adds the serviceName to the delete domain service name owo field params
func (o *DeleteDomainServiceNameOwoFieldParams) WithServiceName(serviceName string) *DeleteDomainServiceNameOwoFieldParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the delete domain service name owo field params
func (o *DeleteDomainServiceNameOwoFieldParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDomainServiceNameOwoFieldParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param field
	if err := r.SetPathParam("field", o.Field); err != nil {
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