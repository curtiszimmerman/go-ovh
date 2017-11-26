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

// NewPostCloudProjectServiceNameInstanceInstanceIDVncParams creates a new PostCloudProjectServiceNameInstanceInstanceIDVncParams object
// with the default values initialized.
func NewPostCloudProjectServiceNameInstanceInstanceIDVncParams() *PostCloudProjectServiceNameInstanceInstanceIDVncParams {
	var ()
	return &PostCloudProjectServiceNameInstanceInstanceIDVncParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostCloudProjectServiceNameInstanceInstanceIDVncParamsWithTimeout creates a new PostCloudProjectServiceNameInstanceInstanceIDVncParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostCloudProjectServiceNameInstanceInstanceIDVncParamsWithTimeout(timeout time.Duration) *PostCloudProjectServiceNameInstanceInstanceIDVncParams {
	var ()
	return &PostCloudProjectServiceNameInstanceInstanceIDVncParams{

		timeout: timeout,
	}
}

// NewPostCloudProjectServiceNameInstanceInstanceIDVncParamsWithContext creates a new PostCloudProjectServiceNameInstanceInstanceIDVncParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostCloudProjectServiceNameInstanceInstanceIDVncParamsWithContext(ctx context.Context) *PostCloudProjectServiceNameInstanceInstanceIDVncParams {
	var ()
	return &PostCloudProjectServiceNameInstanceInstanceIDVncParams{

		Context: ctx,
	}
}

// NewPostCloudProjectServiceNameInstanceInstanceIDVncParamsWithHTTPClient creates a new PostCloudProjectServiceNameInstanceInstanceIDVncParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostCloudProjectServiceNameInstanceInstanceIDVncParamsWithHTTPClient(client *http.Client) *PostCloudProjectServiceNameInstanceInstanceIDVncParams {
	var ()
	return &PostCloudProjectServiceNameInstanceInstanceIDVncParams{
		HTTPClient: client,
	}
}

/*PostCloudProjectServiceNameInstanceInstanceIDVncParams contains all the parameters to send to the API endpoint
for the post cloud project service name instance instance ID vnc operation typically these are written to a http.Request
*/
type PostCloudProjectServiceNameInstanceInstanceIDVncParams struct {

	/*InstanceID*/
	InstanceID string
	/*ServiceName*/
	ServiceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post cloud project service name instance instance ID vnc params
func (o *PostCloudProjectServiceNameInstanceInstanceIDVncParams) WithTimeout(timeout time.Duration) *PostCloudProjectServiceNameInstanceInstanceIDVncParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post cloud project service name instance instance ID vnc params
func (o *PostCloudProjectServiceNameInstanceInstanceIDVncParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post cloud project service name instance instance ID vnc params
func (o *PostCloudProjectServiceNameInstanceInstanceIDVncParams) WithContext(ctx context.Context) *PostCloudProjectServiceNameInstanceInstanceIDVncParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post cloud project service name instance instance ID vnc params
func (o *PostCloudProjectServiceNameInstanceInstanceIDVncParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post cloud project service name instance instance ID vnc params
func (o *PostCloudProjectServiceNameInstanceInstanceIDVncParams) WithHTTPClient(client *http.Client) *PostCloudProjectServiceNameInstanceInstanceIDVncParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post cloud project service name instance instance ID vnc params
func (o *PostCloudProjectServiceNameInstanceInstanceIDVncParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInstanceID adds the instanceID to the post cloud project service name instance instance ID vnc params
func (o *PostCloudProjectServiceNameInstanceInstanceIDVncParams) WithInstanceID(instanceID string) *PostCloudProjectServiceNameInstanceInstanceIDVncParams {
	o.SetInstanceID(instanceID)
	return o
}

// SetInstanceID adds the instanceId to the post cloud project service name instance instance ID vnc params
func (o *PostCloudProjectServiceNameInstanceInstanceIDVncParams) SetInstanceID(instanceID string) {
	o.InstanceID = instanceID
}

// WithServiceName adds the serviceName to the post cloud project service name instance instance ID vnc params
func (o *PostCloudProjectServiceNameInstanceInstanceIDVncParams) WithServiceName(serviceName string) *PostCloudProjectServiceNameInstanceInstanceIDVncParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the post cloud project service name instance instance ID vnc params
func (o *PostCloudProjectServiceNameInstanceInstanceIDVncParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WriteToRequest writes these params to a swagger request
func (o *PostCloudProjectServiceNameInstanceInstanceIDVncParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param instanceId
	if err := r.SetPathParam("instanceId", o.InstanceID); err != nil {
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