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

	"github.com/appscode/go-ovh/cloud/models"
)

// NewPostCloudProjectServiceNameInstanceInstanceIDResizeParams creates a new PostCloudProjectServiceNameInstanceInstanceIDResizeParams object
// with the default values initialized.
func NewPostCloudProjectServiceNameInstanceInstanceIDResizeParams() *PostCloudProjectServiceNameInstanceInstanceIDResizeParams {
	var ()
	return &PostCloudProjectServiceNameInstanceInstanceIDResizeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostCloudProjectServiceNameInstanceInstanceIDResizeParamsWithTimeout creates a new PostCloudProjectServiceNameInstanceInstanceIDResizeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostCloudProjectServiceNameInstanceInstanceIDResizeParamsWithTimeout(timeout time.Duration) *PostCloudProjectServiceNameInstanceInstanceIDResizeParams {
	var ()
	return &PostCloudProjectServiceNameInstanceInstanceIDResizeParams{

		timeout: timeout,
	}
}

// NewPostCloudProjectServiceNameInstanceInstanceIDResizeParamsWithContext creates a new PostCloudProjectServiceNameInstanceInstanceIDResizeParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostCloudProjectServiceNameInstanceInstanceIDResizeParamsWithContext(ctx context.Context) *PostCloudProjectServiceNameInstanceInstanceIDResizeParams {
	var ()
	return &PostCloudProjectServiceNameInstanceInstanceIDResizeParams{

		Context: ctx,
	}
}

// NewPostCloudProjectServiceNameInstanceInstanceIDResizeParamsWithHTTPClient creates a new PostCloudProjectServiceNameInstanceInstanceIDResizeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostCloudProjectServiceNameInstanceInstanceIDResizeParamsWithHTTPClient(client *http.Client) *PostCloudProjectServiceNameInstanceInstanceIDResizeParams {
	var ()
	return &PostCloudProjectServiceNameInstanceInstanceIDResizeParams{
		HTTPClient: client,
	}
}

/*PostCloudProjectServiceNameInstanceInstanceIDResizeParams contains all the parameters to send to the API endpoint
for the post cloud project service name instance instance ID resize operation typically these are written to a http.Request
*/
type PostCloudProjectServiceNameInstanceInstanceIDResizeParams struct {

	/*CloudProjectInstanceResizePost*/
	CloudProjectInstanceResizePost *models.PostCloudProjectServiceNameInstanceInstanceIDResizeParamsBody
	/*InstanceID*/
	InstanceID string
	/*ServiceName*/
	ServiceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post cloud project service name instance instance ID resize params
func (o *PostCloudProjectServiceNameInstanceInstanceIDResizeParams) WithTimeout(timeout time.Duration) *PostCloudProjectServiceNameInstanceInstanceIDResizeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post cloud project service name instance instance ID resize params
func (o *PostCloudProjectServiceNameInstanceInstanceIDResizeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post cloud project service name instance instance ID resize params
func (o *PostCloudProjectServiceNameInstanceInstanceIDResizeParams) WithContext(ctx context.Context) *PostCloudProjectServiceNameInstanceInstanceIDResizeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post cloud project service name instance instance ID resize params
func (o *PostCloudProjectServiceNameInstanceInstanceIDResizeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post cloud project service name instance instance ID resize params
func (o *PostCloudProjectServiceNameInstanceInstanceIDResizeParams) WithHTTPClient(client *http.Client) *PostCloudProjectServiceNameInstanceInstanceIDResizeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post cloud project service name instance instance ID resize params
func (o *PostCloudProjectServiceNameInstanceInstanceIDResizeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudProjectInstanceResizePost adds the cloudProjectInstanceResizePost to the post cloud project service name instance instance ID resize params
func (o *PostCloudProjectServiceNameInstanceInstanceIDResizeParams) WithCloudProjectInstanceResizePost(cloudProjectInstanceResizePost *models.PostCloudProjectServiceNameInstanceInstanceIDResizeParamsBody) *PostCloudProjectServiceNameInstanceInstanceIDResizeParams {
	o.SetCloudProjectInstanceResizePost(cloudProjectInstanceResizePost)
	return o
}

// SetCloudProjectInstanceResizePost adds the cloudProjectInstanceResizePost to the post cloud project service name instance instance ID resize params
func (o *PostCloudProjectServiceNameInstanceInstanceIDResizeParams) SetCloudProjectInstanceResizePost(cloudProjectInstanceResizePost *models.PostCloudProjectServiceNameInstanceInstanceIDResizeParamsBody) {
	o.CloudProjectInstanceResizePost = cloudProjectInstanceResizePost
}

// WithInstanceID adds the instanceID to the post cloud project service name instance instance ID resize params
func (o *PostCloudProjectServiceNameInstanceInstanceIDResizeParams) WithInstanceID(instanceID string) *PostCloudProjectServiceNameInstanceInstanceIDResizeParams {
	o.SetInstanceID(instanceID)
	return o
}

// SetInstanceID adds the instanceId to the post cloud project service name instance instance ID resize params
func (o *PostCloudProjectServiceNameInstanceInstanceIDResizeParams) SetInstanceID(instanceID string) {
	o.InstanceID = instanceID
}

// WithServiceName adds the serviceName to the post cloud project service name instance instance ID resize params
func (o *PostCloudProjectServiceNameInstanceInstanceIDResizeParams) WithServiceName(serviceName string) *PostCloudProjectServiceNameInstanceInstanceIDResizeParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the post cloud project service name instance instance ID resize params
func (o *PostCloudProjectServiceNameInstanceInstanceIDResizeParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WriteToRequest writes these params to a swagger request
func (o *PostCloudProjectServiceNameInstanceInstanceIDResizeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CloudProjectInstanceResizePost != nil {
		if err := r.SetBodyParam(o.CloudProjectInstanceResizePost); err != nil {
			return err
		}
	}

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