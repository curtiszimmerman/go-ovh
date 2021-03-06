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

// NewPostCloudProjectServiceNameIPLoadbalancingParams creates a new PostCloudProjectServiceNameIPLoadbalancingParams object
// with the default values initialized.
func NewPostCloudProjectServiceNameIPLoadbalancingParams() *PostCloudProjectServiceNameIPLoadbalancingParams {
	var ()
	return &PostCloudProjectServiceNameIPLoadbalancingParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostCloudProjectServiceNameIPLoadbalancingParamsWithTimeout creates a new PostCloudProjectServiceNameIPLoadbalancingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostCloudProjectServiceNameIPLoadbalancingParamsWithTimeout(timeout time.Duration) *PostCloudProjectServiceNameIPLoadbalancingParams {
	var ()
	return &PostCloudProjectServiceNameIPLoadbalancingParams{

		timeout: timeout,
	}
}

// NewPostCloudProjectServiceNameIPLoadbalancingParamsWithContext creates a new PostCloudProjectServiceNameIPLoadbalancingParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostCloudProjectServiceNameIPLoadbalancingParamsWithContext(ctx context.Context) *PostCloudProjectServiceNameIPLoadbalancingParams {
	var ()
	return &PostCloudProjectServiceNameIPLoadbalancingParams{

		Context: ctx,
	}
}

// NewPostCloudProjectServiceNameIPLoadbalancingParamsWithHTTPClient creates a new PostCloudProjectServiceNameIPLoadbalancingParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostCloudProjectServiceNameIPLoadbalancingParamsWithHTTPClient(client *http.Client) *PostCloudProjectServiceNameIPLoadbalancingParams {
	var ()
	return &PostCloudProjectServiceNameIPLoadbalancingParams{
		HTTPClient: client,
	}
}

/*PostCloudProjectServiceNameIPLoadbalancingParams contains all the parameters to send to the API endpoint
for the post cloud project service name IP loadbalancing operation typically these are written to a http.Request
*/
type PostCloudProjectServiceNameIPLoadbalancingParams struct {

	/*CloudProjectIPLBPost*/
	CloudProjectIPLBPost *models.PostCloudProjectServiceNameIPLoadbalancingParamsBody
	/*ServiceName*/
	ServiceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post cloud project service name IP loadbalancing params
func (o *PostCloudProjectServiceNameIPLoadbalancingParams) WithTimeout(timeout time.Duration) *PostCloudProjectServiceNameIPLoadbalancingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post cloud project service name IP loadbalancing params
func (o *PostCloudProjectServiceNameIPLoadbalancingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post cloud project service name IP loadbalancing params
func (o *PostCloudProjectServiceNameIPLoadbalancingParams) WithContext(ctx context.Context) *PostCloudProjectServiceNameIPLoadbalancingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post cloud project service name IP loadbalancing params
func (o *PostCloudProjectServiceNameIPLoadbalancingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post cloud project service name IP loadbalancing params
func (o *PostCloudProjectServiceNameIPLoadbalancingParams) WithHTTPClient(client *http.Client) *PostCloudProjectServiceNameIPLoadbalancingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post cloud project service name IP loadbalancing params
func (o *PostCloudProjectServiceNameIPLoadbalancingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudProjectIPLBPost adds the cloudProjectIPLBPost to the post cloud project service name IP loadbalancing params
func (o *PostCloudProjectServiceNameIPLoadbalancingParams) WithCloudProjectIPLBPost(cloudProjectIPLBPost *models.PostCloudProjectServiceNameIPLoadbalancingParamsBody) *PostCloudProjectServiceNameIPLoadbalancingParams {
	o.SetCloudProjectIPLBPost(cloudProjectIPLBPost)
	return o
}

// SetCloudProjectIPLBPost adds the cloudProjectIplbPost to the post cloud project service name IP loadbalancing params
func (o *PostCloudProjectServiceNameIPLoadbalancingParams) SetCloudProjectIPLBPost(cloudProjectIPLBPost *models.PostCloudProjectServiceNameIPLoadbalancingParamsBody) {
	o.CloudProjectIPLBPost = cloudProjectIPLBPost
}

// WithServiceName adds the serviceName to the post cloud project service name IP loadbalancing params
func (o *PostCloudProjectServiceNameIPLoadbalancingParams) WithServiceName(serviceName string) *PostCloudProjectServiceNameIPLoadbalancingParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the post cloud project service name IP loadbalancing params
func (o *PostCloudProjectServiceNameIPLoadbalancingParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WriteToRequest writes these params to a swagger request
func (o *PostCloudProjectServiceNameIPLoadbalancingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CloudProjectIPLBPost != nil {
		if err := r.SetBodyParam(o.CloudProjectIPLBPost); err != nil {
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
