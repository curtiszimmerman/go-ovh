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

// NewGetVrackServiceNameIPLoadbalancingParams creates a new GetVrackServiceNameIPLoadbalancingParams object
// with the default values initialized.
func NewGetVrackServiceNameIPLoadbalancingParams() *GetVrackServiceNameIPLoadbalancingParams {
	var ()
	return &GetVrackServiceNameIPLoadbalancingParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetVrackServiceNameIPLoadbalancingParamsWithTimeout creates a new GetVrackServiceNameIPLoadbalancingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetVrackServiceNameIPLoadbalancingParamsWithTimeout(timeout time.Duration) *GetVrackServiceNameIPLoadbalancingParams {
	var ()
	return &GetVrackServiceNameIPLoadbalancingParams{

		timeout: timeout,
	}
}

// NewGetVrackServiceNameIPLoadbalancingParamsWithContext creates a new GetVrackServiceNameIPLoadbalancingParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetVrackServiceNameIPLoadbalancingParamsWithContext(ctx context.Context) *GetVrackServiceNameIPLoadbalancingParams {
	var ()
	return &GetVrackServiceNameIPLoadbalancingParams{

		Context: ctx,
	}
}

// NewGetVrackServiceNameIPLoadbalancingParamsWithHTTPClient creates a new GetVrackServiceNameIPLoadbalancingParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetVrackServiceNameIPLoadbalancingParamsWithHTTPClient(client *http.Client) *GetVrackServiceNameIPLoadbalancingParams {
	var ()
	return &GetVrackServiceNameIPLoadbalancingParams{
		HTTPClient: client,
	}
}

/*GetVrackServiceNameIPLoadbalancingParams contains all the parameters to send to the API endpoint
for the get vrack service name IP loadbalancing operation typically these are written to a http.Request
*/
type GetVrackServiceNameIPLoadbalancingParams struct {

	/*ServiceName*/
	ServiceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get vrack service name IP loadbalancing params
func (o *GetVrackServiceNameIPLoadbalancingParams) WithTimeout(timeout time.Duration) *GetVrackServiceNameIPLoadbalancingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get vrack service name IP loadbalancing params
func (o *GetVrackServiceNameIPLoadbalancingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get vrack service name IP loadbalancing params
func (o *GetVrackServiceNameIPLoadbalancingParams) WithContext(ctx context.Context) *GetVrackServiceNameIPLoadbalancingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get vrack service name IP loadbalancing params
func (o *GetVrackServiceNameIPLoadbalancingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get vrack service name IP loadbalancing params
func (o *GetVrackServiceNameIPLoadbalancingParams) WithHTTPClient(client *http.Client) *GetVrackServiceNameIPLoadbalancingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get vrack service name IP loadbalancing params
func (o *GetVrackServiceNameIPLoadbalancingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithServiceName adds the serviceName to the get vrack service name IP loadbalancing params
func (o *GetVrackServiceNameIPLoadbalancingParams) WithServiceName(serviceName string) *GetVrackServiceNameIPLoadbalancingParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the get vrack service name IP loadbalancing params
func (o *GetVrackServiceNameIPLoadbalancingParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WriteToRequest writes these params to a swagger request
func (o *GetVrackServiceNameIPLoadbalancingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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