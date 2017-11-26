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

	"github.com/appscode/go-ovh/iploadbalancing/models"
)

// NewPostIPLoadbalancingServiceNameHTTPFrontendParams creates a new PostIPLoadbalancingServiceNameHTTPFrontendParams object
// with the default values initialized.
func NewPostIPLoadbalancingServiceNameHTTPFrontendParams() *PostIPLoadbalancingServiceNameHTTPFrontendParams {
	var ()
	return &PostIPLoadbalancingServiceNameHTTPFrontendParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostIPLoadbalancingServiceNameHTTPFrontendParamsWithTimeout creates a new PostIPLoadbalancingServiceNameHTTPFrontendParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostIPLoadbalancingServiceNameHTTPFrontendParamsWithTimeout(timeout time.Duration) *PostIPLoadbalancingServiceNameHTTPFrontendParams {
	var ()
	return &PostIPLoadbalancingServiceNameHTTPFrontendParams{

		timeout: timeout,
	}
}

// NewPostIPLoadbalancingServiceNameHTTPFrontendParamsWithContext creates a new PostIPLoadbalancingServiceNameHTTPFrontendParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostIPLoadbalancingServiceNameHTTPFrontendParamsWithContext(ctx context.Context) *PostIPLoadbalancingServiceNameHTTPFrontendParams {
	var ()
	return &PostIPLoadbalancingServiceNameHTTPFrontendParams{

		Context: ctx,
	}
}

// NewPostIPLoadbalancingServiceNameHTTPFrontendParamsWithHTTPClient creates a new PostIPLoadbalancingServiceNameHTTPFrontendParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostIPLoadbalancingServiceNameHTTPFrontendParamsWithHTTPClient(client *http.Client) *PostIPLoadbalancingServiceNameHTTPFrontendParams {
	var ()
	return &PostIPLoadbalancingServiceNameHTTPFrontendParams{
		HTTPClient: client,
	}
}

/*PostIPLoadbalancingServiceNameHTTPFrontendParams contains all the parameters to send to the API endpoint
for the post IP loadbalancing service name HTTP frontend operation typically these are written to a http.Request
*/
type PostIPLoadbalancingServiceNameHTTPFrontendParams struct {

	/*IPLBHTTPFrontendPost*/
	IPLBHTTPFrontendPost *models.PostIPLoadbalancingServiceNameHTTPFrontendParamsBody
	/*ServiceName*/
	ServiceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post IP loadbalancing service name HTTP frontend params
func (o *PostIPLoadbalancingServiceNameHTTPFrontendParams) WithTimeout(timeout time.Duration) *PostIPLoadbalancingServiceNameHTTPFrontendParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post IP loadbalancing service name HTTP frontend params
func (o *PostIPLoadbalancingServiceNameHTTPFrontendParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post IP loadbalancing service name HTTP frontend params
func (o *PostIPLoadbalancingServiceNameHTTPFrontendParams) WithContext(ctx context.Context) *PostIPLoadbalancingServiceNameHTTPFrontendParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post IP loadbalancing service name HTTP frontend params
func (o *PostIPLoadbalancingServiceNameHTTPFrontendParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post IP loadbalancing service name HTTP frontend params
func (o *PostIPLoadbalancingServiceNameHTTPFrontendParams) WithHTTPClient(client *http.Client) *PostIPLoadbalancingServiceNameHTTPFrontendParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post IP loadbalancing service name HTTP frontend params
func (o *PostIPLoadbalancingServiceNameHTTPFrontendParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIPLBHTTPFrontendPost adds the iPLBHTTPFrontendPost to the post IP loadbalancing service name HTTP frontend params
func (o *PostIPLoadbalancingServiceNameHTTPFrontendParams) WithIPLBHTTPFrontendPost(iPLBHTTPFrontendPost *models.PostIPLoadbalancingServiceNameHTTPFrontendParamsBody) *PostIPLoadbalancingServiceNameHTTPFrontendParams {
	o.SetIPLBHTTPFrontendPost(iPLBHTTPFrontendPost)
	return o
}

// SetIPLBHTTPFrontendPost adds the iplbHttpFrontendPost to the post IP loadbalancing service name HTTP frontend params
func (o *PostIPLoadbalancingServiceNameHTTPFrontendParams) SetIPLBHTTPFrontendPost(iPLBHTTPFrontendPost *models.PostIPLoadbalancingServiceNameHTTPFrontendParamsBody) {
	o.IPLBHTTPFrontendPost = iPLBHTTPFrontendPost
}

// WithServiceName adds the serviceName to the post IP loadbalancing service name HTTP frontend params
func (o *PostIPLoadbalancingServiceNameHTTPFrontendParams) WithServiceName(serviceName string) *PostIPLoadbalancingServiceNameHTTPFrontendParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the post IP loadbalancing service name HTTP frontend params
func (o *PostIPLoadbalancingServiceNameHTTPFrontendParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WriteToRequest writes these params to a swagger request
func (o *PostIPLoadbalancingServiceNameHTTPFrontendParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IPLBHTTPFrontendPost != nil {
		if err := r.SetBodyParam(o.IPLBHTTPFrontendPost); err != nil {
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