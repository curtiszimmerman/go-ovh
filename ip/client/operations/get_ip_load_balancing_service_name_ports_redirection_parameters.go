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

// NewGetIPLoadBalancingServiceNamePortsRedirectionParams creates a new GetIPLoadBalancingServiceNamePortsRedirectionParams object
// with the default values initialized.
func NewGetIPLoadBalancingServiceNamePortsRedirectionParams() *GetIPLoadBalancingServiceNamePortsRedirectionParams {
	var ()
	return &GetIPLoadBalancingServiceNamePortsRedirectionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetIPLoadBalancingServiceNamePortsRedirectionParamsWithTimeout creates a new GetIPLoadBalancingServiceNamePortsRedirectionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetIPLoadBalancingServiceNamePortsRedirectionParamsWithTimeout(timeout time.Duration) *GetIPLoadBalancingServiceNamePortsRedirectionParams {
	var ()
	return &GetIPLoadBalancingServiceNamePortsRedirectionParams{

		timeout: timeout,
	}
}

// NewGetIPLoadBalancingServiceNamePortsRedirectionParamsWithContext creates a new GetIPLoadBalancingServiceNamePortsRedirectionParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetIPLoadBalancingServiceNamePortsRedirectionParamsWithContext(ctx context.Context) *GetIPLoadBalancingServiceNamePortsRedirectionParams {
	var ()
	return &GetIPLoadBalancingServiceNamePortsRedirectionParams{

		Context: ctx,
	}
}

// NewGetIPLoadBalancingServiceNamePortsRedirectionParamsWithHTTPClient creates a new GetIPLoadBalancingServiceNamePortsRedirectionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetIPLoadBalancingServiceNamePortsRedirectionParamsWithHTTPClient(client *http.Client) *GetIPLoadBalancingServiceNamePortsRedirectionParams {
	var ()
	return &GetIPLoadBalancingServiceNamePortsRedirectionParams{
		HTTPClient: client,
	}
}

/*GetIPLoadBalancingServiceNamePortsRedirectionParams contains all the parameters to send to the API endpoint
for the get IP load balancing service name ports redirection operation typically these are written to a http.Request
*/
type GetIPLoadBalancingServiceNamePortsRedirectionParams struct {

	/*ServiceName*/
	ServiceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get IP load balancing service name ports redirection params
func (o *GetIPLoadBalancingServiceNamePortsRedirectionParams) WithTimeout(timeout time.Duration) *GetIPLoadBalancingServiceNamePortsRedirectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get IP load balancing service name ports redirection params
func (o *GetIPLoadBalancingServiceNamePortsRedirectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get IP load balancing service name ports redirection params
func (o *GetIPLoadBalancingServiceNamePortsRedirectionParams) WithContext(ctx context.Context) *GetIPLoadBalancingServiceNamePortsRedirectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get IP load balancing service name ports redirection params
func (o *GetIPLoadBalancingServiceNamePortsRedirectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get IP load balancing service name ports redirection params
func (o *GetIPLoadBalancingServiceNamePortsRedirectionParams) WithHTTPClient(client *http.Client) *GetIPLoadBalancingServiceNamePortsRedirectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get IP load balancing service name ports redirection params
func (o *GetIPLoadBalancingServiceNamePortsRedirectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithServiceName adds the serviceName to the get IP load balancing service name ports redirection params
func (o *GetIPLoadBalancingServiceNamePortsRedirectionParams) WithServiceName(serviceName string) *GetIPLoadBalancingServiceNamePortsRedirectionParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the get IP load balancing service name ports redirection params
func (o *GetIPLoadBalancingServiceNamePortsRedirectionParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WriteToRequest writes these params to a swagger request
func (o *GetIPLoadBalancingServiceNamePortsRedirectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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