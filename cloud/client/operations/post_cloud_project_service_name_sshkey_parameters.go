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

// NewPostCloudProjectServiceNameSshkeyParams creates a new PostCloudProjectServiceNameSshkeyParams object
// with the default values initialized.
func NewPostCloudProjectServiceNameSshkeyParams() *PostCloudProjectServiceNameSshkeyParams {
	var ()
	return &PostCloudProjectServiceNameSshkeyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostCloudProjectServiceNameSshkeyParamsWithTimeout creates a new PostCloudProjectServiceNameSshkeyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostCloudProjectServiceNameSshkeyParamsWithTimeout(timeout time.Duration) *PostCloudProjectServiceNameSshkeyParams {
	var ()
	return &PostCloudProjectServiceNameSshkeyParams{

		timeout: timeout,
	}
}

// NewPostCloudProjectServiceNameSshkeyParamsWithContext creates a new PostCloudProjectServiceNameSshkeyParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostCloudProjectServiceNameSshkeyParamsWithContext(ctx context.Context) *PostCloudProjectServiceNameSshkeyParams {
	var ()
	return &PostCloudProjectServiceNameSshkeyParams{

		Context: ctx,
	}
}

// NewPostCloudProjectServiceNameSshkeyParamsWithHTTPClient creates a new PostCloudProjectServiceNameSshkeyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostCloudProjectServiceNameSshkeyParamsWithHTTPClient(client *http.Client) *PostCloudProjectServiceNameSshkeyParams {
	var ()
	return &PostCloudProjectServiceNameSshkeyParams{
		HTTPClient: client,
	}
}

/*PostCloudProjectServiceNameSshkeyParams contains all the parameters to send to the API endpoint
for the post cloud project service name sshkey operation typically these are written to a http.Request
*/
type PostCloudProjectServiceNameSshkeyParams struct {

	/*CloudProjectSshkeyPost*/
	CloudProjectSshkeyPost *models.PostCloudProjectServiceNameSshkeyParamsBody
	/*ServiceName*/
	ServiceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post cloud project service name sshkey params
func (o *PostCloudProjectServiceNameSshkeyParams) WithTimeout(timeout time.Duration) *PostCloudProjectServiceNameSshkeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post cloud project service name sshkey params
func (o *PostCloudProjectServiceNameSshkeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post cloud project service name sshkey params
func (o *PostCloudProjectServiceNameSshkeyParams) WithContext(ctx context.Context) *PostCloudProjectServiceNameSshkeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post cloud project service name sshkey params
func (o *PostCloudProjectServiceNameSshkeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post cloud project service name sshkey params
func (o *PostCloudProjectServiceNameSshkeyParams) WithHTTPClient(client *http.Client) *PostCloudProjectServiceNameSshkeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post cloud project service name sshkey params
func (o *PostCloudProjectServiceNameSshkeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudProjectSshkeyPost adds the cloudProjectSshkeyPost to the post cloud project service name sshkey params
func (o *PostCloudProjectServiceNameSshkeyParams) WithCloudProjectSshkeyPost(cloudProjectSshkeyPost *models.PostCloudProjectServiceNameSshkeyParamsBody) *PostCloudProjectServiceNameSshkeyParams {
	o.SetCloudProjectSshkeyPost(cloudProjectSshkeyPost)
	return o
}

// SetCloudProjectSshkeyPost adds the cloudProjectSshkeyPost to the post cloud project service name sshkey params
func (o *PostCloudProjectServiceNameSshkeyParams) SetCloudProjectSshkeyPost(cloudProjectSshkeyPost *models.PostCloudProjectServiceNameSshkeyParamsBody) {
	o.CloudProjectSshkeyPost = cloudProjectSshkeyPost
}

// WithServiceName adds the serviceName to the post cloud project service name sshkey params
func (o *PostCloudProjectServiceNameSshkeyParams) WithServiceName(serviceName string) *PostCloudProjectServiceNameSshkeyParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the post cloud project service name sshkey params
func (o *PostCloudProjectServiceNameSshkeyParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WriteToRequest writes these params to a swagger request
func (o *PostCloudProjectServiceNameSshkeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CloudProjectSshkeyPost != nil {
		if err := r.SetBodyParam(o.CloudProjectSshkeyPost); err != nil {
			return err
		}
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