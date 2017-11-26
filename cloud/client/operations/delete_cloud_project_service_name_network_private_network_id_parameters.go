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

// NewDeleteCloudProjectServiceNameNetworkPrivateNetworkIDParams creates a new DeleteCloudProjectServiceNameNetworkPrivateNetworkIDParams object
// with the default values initialized.
func NewDeleteCloudProjectServiceNameNetworkPrivateNetworkIDParams() *DeleteCloudProjectServiceNameNetworkPrivateNetworkIDParams {
	var ()
	return &DeleteCloudProjectServiceNameNetworkPrivateNetworkIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCloudProjectServiceNameNetworkPrivateNetworkIDParamsWithTimeout creates a new DeleteCloudProjectServiceNameNetworkPrivateNetworkIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteCloudProjectServiceNameNetworkPrivateNetworkIDParamsWithTimeout(timeout time.Duration) *DeleteCloudProjectServiceNameNetworkPrivateNetworkIDParams {
	var ()
	return &DeleteCloudProjectServiceNameNetworkPrivateNetworkIDParams{

		timeout: timeout,
	}
}

// NewDeleteCloudProjectServiceNameNetworkPrivateNetworkIDParamsWithContext creates a new DeleteCloudProjectServiceNameNetworkPrivateNetworkIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteCloudProjectServiceNameNetworkPrivateNetworkIDParamsWithContext(ctx context.Context) *DeleteCloudProjectServiceNameNetworkPrivateNetworkIDParams {
	var ()
	return &DeleteCloudProjectServiceNameNetworkPrivateNetworkIDParams{

		Context: ctx,
	}
}

// NewDeleteCloudProjectServiceNameNetworkPrivateNetworkIDParamsWithHTTPClient creates a new DeleteCloudProjectServiceNameNetworkPrivateNetworkIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteCloudProjectServiceNameNetworkPrivateNetworkIDParamsWithHTTPClient(client *http.Client) *DeleteCloudProjectServiceNameNetworkPrivateNetworkIDParams {
	var ()
	return &DeleteCloudProjectServiceNameNetworkPrivateNetworkIDParams{
		HTTPClient: client,
	}
}

/*DeleteCloudProjectServiceNameNetworkPrivateNetworkIDParams contains all the parameters to send to the API endpoint
for the delete cloud project service name network private network ID operation typically these are written to a http.Request
*/
type DeleteCloudProjectServiceNameNetworkPrivateNetworkIDParams struct {

	/*NetworkID*/
	NetworkID string
	/*ServiceName*/
	ServiceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete cloud project service name network private network ID params
func (o *DeleteCloudProjectServiceNameNetworkPrivateNetworkIDParams) WithTimeout(timeout time.Duration) *DeleteCloudProjectServiceNameNetworkPrivateNetworkIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete cloud project service name network private network ID params
func (o *DeleteCloudProjectServiceNameNetworkPrivateNetworkIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete cloud project service name network private network ID params
func (o *DeleteCloudProjectServiceNameNetworkPrivateNetworkIDParams) WithContext(ctx context.Context) *DeleteCloudProjectServiceNameNetworkPrivateNetworkIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete cloud project service name network private network ID params
func (o *DeleteCloudProjectServiceNameNetworkPrivateNetworkIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete cloud project service name network private network ID params
func (o *DeleteCloudProjectServiceNameNetworkPrivateNetworkIDParams) WithHTTPClient(client *http.Client) *DeleteCloudProjectServiceNameNetworkPrivateNetworkIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete cloud project service name network private network ID params
func (o *DeleteCloudProjectServiceNameNetworkPrivateNetworkIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the delete cloud project service name network private network ID params
func (o *DeleteCloudProjectServiceNameNetworkPrivateNetworkIDParams) WithNetworkID(networkID string) *DeleteCloudProjectServiceNameNetworkPrivateNetworkIDParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the delete cloud project service name network private network ID params
func (o *DeleteCloudProjectServiceNameNetworkPrivateNetworkIDParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithServiceName adds the serviceName to the delete cloud project service name network private network ID params
func (o *DeleteCloudProjectServiceNameNetworkPrivateNetworkIDParams) WithServiceName(serviceName string) *DeleteCloudProjectServiceNameNetworkPrivateNetworkIDParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the delete cloud project service name network private network ID params
func (o *DeleteCloudProjectServiceNameNetworkPrivateNetworkIDParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCloudProjectServiceNameNetworkPrivateNetworkIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
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