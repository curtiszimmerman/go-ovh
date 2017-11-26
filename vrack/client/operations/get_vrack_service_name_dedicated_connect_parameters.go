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

// NewGetVrackServiceNameDedicatedConnectParams creates a new GetVrackServiceNameDedicatedConnectParams object
// with the default values initialized.
func NewGetVrackServiceNameDedicatedConnectParams() *GetVrackServiceNameDedicatedConnectParams {
	var ()
	return &GetVrackServiceNameDedicatedConnectParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetVrackServiceNameDedicatedConnectParamsWithTimeout creates a new GetVrackServiceNameDedicatedConnectParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetVrackServiceNameDedicatedConnectParamsWithTimeout(timeout time.Duration) *GetVrackServiceNameDedicatedConnectParams {
	var ()
	return &GetVrackServiceNameDedicatedConnectParams{

		timeout: timeout,
	}
}

// NewGetVrackServiceNameDedicatedConnectParamsWithContext creates a new GetVrackServiceNameDedicatedConnectParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetVrackServiceNameDedicatedConnectParamsWithContext(ctx context.Context) *GetVrackServiceNameDedicatedConnectParams {
	var ()
	return &GetVrackServiceNameDedicatedConnectParams{

		Context: ctx,
	}
}

// NewGetVrackServiceNameDedicatedConnectParamsWithHTTPClient creates a new GetVrackServiceNameDedicatedConnectParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetVrackServiceNameDedicatedConnectParamsWithHTTPClient(client *http.Client) *GetVrackServiceNameDedicatedConnectParams {
	var ()
	return &GetVrackServiceNameDedicatedConnectParams{
		HTTPClient: client,
	}
}

/*GetVrackServiceNameDedicatedConnectParams contains all the parameters to send to the API endpoint
for the get vrack service name dedicated connect operation typically these are written to a http.Request
*/
type GetVrackServiceNameDedicatedConnectParams struct {

	/*ServiceName*/
	ServiceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get vrack service name dedicated connect params
func (o *GetVrackServiceNameDedicatedConnectParams) WithTimeout(timeout time.Duration) *GetVrackServiceNameDedicatedConnectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get vrack service name dedicated connect params
func (o *GetVrackServiceNameDedicatedConnectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get vrack service name dedicated connect params
func (o *GetVrackServiceNameDedicatedConnectParams) WithContext(ctx context.Context) *GetVrackServiceNameDedicatedConnectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get vrack service name dedicated connect params
func (o *GetVrackServiceNameDedicatedConnectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get vrack service name dedicated connect params
func (o *GetVrackServiceNameDedicatedConnectParams) WithHTTPClient(client *http.Client) *GetVrackServiceNameDedicatedConnectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get vrack service name dedicated connect params
func (o *GetVrackServiceNameDedicatedConnectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithServiceName adds the serviceName to the get vrack service name dedicated connect params
func (o *GetVrackServiceNameDedicatedConnectParams) WithServiceName(serviceName string) *GetVrackServiceNameDedicatedConnectParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the get vrack service name dedicated connect params
func (o *GetVrackServiceNameDedicatedConnectParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WriteToRequest writes these params to a swagger request
func (o *GetVrackServiceNameDedicatedConnectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param serviceName
	if err := r.SetPathParam("serviceName", o.ServiceName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}