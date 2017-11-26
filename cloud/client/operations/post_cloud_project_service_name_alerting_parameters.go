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

// NewPostCloudProjectServiceNameAlertingParams creates a new PostCloudProjectServiceNameAlertingParams object
// with the default values initialized.
func NewPostCloudProjectServiceNameAlertingParams() *PostCloudProjectServiceNameAlertingParams {
	var ()
	return &PostCloudProjectServiceNameAlertingParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostCloudProjectServiceNameAlertingParamsWithTimeout creates a new PostCloudProjectServiceNameAlertingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostCloudProjectServiceNameAlertingParamsWithTimeout(timeout time.Duration) *PostCloudProjectServiceNameAlertingParams {
	var ()
	return &PostCloudProjectServiceNameAlertingParams{

		timeout: timeout,
	}
}

// NewPostCloudProjectServiceNameAlertingParamsWithContext creates a new PostCloudProjectServiceNameAlertingParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostCloudProjectServiceNameAlertingParamsWithContext(ctx context.Context) *PostCloudProjectServiceNameAlertingParams {
	var ()
	return &PostCloudProjectServiceNameAlertingParams{

		Context: ctx,
	}
}

// NewPostCloudProjectServiceNameAlertingParamsWithHTTPClient creates a new PostCloudProjectServiceNameAlertingParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostCloudProjectServiceNameAlertingParamsWithHTTPClient(client *http.Client) *PostCloudProjectServiceNameAlertingParams {
	var ()
	return &PostCloudProjectServiceNameAlertingParams{
		HTTPClient: client,
	}
}

/*PostCloudProjectServiceNameAlertingParams contains all the parameters to send to the API endpoint
for the post cloud project service name alerting operation typically these are written to a http.Request
*/
type PostCloudProjectServiceNameAlertingParams struct {

	/*CloudProjectAlertingPost*/
	CloudProjectAlertingPost *models.PostCloudProjectServiceNameAlertingParamsBody
	/*ServiceName*/
	ServiceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post cloud project service name alerting params
func (o *PostCloudProjectServiceNameAlertingParams) WithTimeout(timeout time.Duration) *PostCloudProjectServiceNameAlertingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post cloud project service name alerting params
func (o *PostCloudProjectServiceNameAlertingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post cloud project service name alerting params
func (o *PostCloudProjectServiceNameAlertingParams) WithContext(ctx context.Context) *PostCloudProjectServiceNameAlertingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post cloud project service name alerting params
func (o *PostCloudProjectServiceNameAlertingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post cloud project service name alerting params
func (o *PostCloudProjectServiceNameAlertingParams) WithHTTPClient(client *http.Client) *PostCloudProjectServiceNameAlertingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post cloud project service name alerting params
func (o *PostCloudProjectServiceNameAlertingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudProjectAlertingPost adds the cloudProjectAlertingPost to the post cloud project service name alerting params
func (o *PostCloudProjectServiceNameAlertingParams) WithCloudProjectAlertingPost(cloudProjectAlertingPost *models.PostCloudProjectServiceNameAlertingParamsBody) *PostCloudProjectServiceNameAlertingParams {
	o.SetCloudProjectAlertingPost(cloudProjectAlertingPost)
	return o
}

// SetCloudProjectAlertingPost adds the cloudProjectAlertingPost to the post cloud project service name alerting params
func (o *PostCloudProjectServiceNameAlertingParams) SetCloudProjectAlertingPost(cloudProjectAlertingPost *models.PostCloudProjectServiceNameAlertingParamsBody) {
	o.CloudProjectAlertingPost = cloudProjectAlertingPost
}

// WithServiceName adds the serviceName to the post cloud project service name alerting params
func (o *PostCloudProjectServiceNameAlertingParams) WithServiceName(serviceName string) *PostCloudProjectServiceNameAlertingParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the post cloud project service name alerting params
func (o *PostCloudProjectServiceNameAlertingParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WriteToRequest writes these params to a swagger request
func (o *PostCloudProjectServiceNameAlertingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CloudProjectAlertingPost != nil {
		if err := r.SetBodyParam(o.CloudProjectAlertingPost); err != nil {
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