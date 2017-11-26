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

	"github.com/appscode/go-ovh/vps/models"
)

// NewPostVpsServiceNameChangeContactParams creates a new PostVpsServiceNameChangeContactParams object
// with the default values initialized.
func NewPostVpsServiceNameChangeContactParams() *PostVpsServiceNameChangeContactParams {
	var ()
	return &PostVpsServiceNameChangeContactParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostVpsServiceNameChangeContactParamsWithTimeout creates a new PostVpsServiceNameChangeContactParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostVpsServiceNameChangeContactParamsWithTimeout(timeout time.Duration) *PostVpsServiceNameChangeContactParams {
	var ()
	return &PostVpsServiceNameChangeContactParams{

		timeout: timeout,
	}
}

// NewPostVpsServiceNameChangeContactParamsWithContext creates a new PostVpsServiceNameChangeContactParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostVpsServiceNameChangeContactParamsWithContext(ctx context.Context) *PostVpsServiceNameChangeContactParams {
	var ()
	return &PostVpsServiceNameChangeContactParams{

		Context: ctx,
	}
}

// NewPostVpsServiceNameChangeContactParamsWithHTTPClient creates a new PostVpsServiceNameChangeContactParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostVpsServiceNameChangeContactParamsWithHTTPClient(client *http.Client) *PostVpsServiceNameChangeContactParams {
	var ()
	return &PostVpsServiceNameChangeContactParams{
		HTTPClient: client,
	}
}

/*PostVpsServiceNameChangeContactParams contains all the parameters to send to the API endpoint
for the post vps service name change contact operation typically these are written to a http.Request
*/
type PostVpsServiceNameChangeContactParams struct {

	/*ServiceName*/
	ServiceName string
	/*VpsChangeContactPost*/
	VpsChangeContactPost *models.PostVpsServiceNameChangeContactParamsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post vps service name change contact params
func (o *PostVpsServiceNameChangeContactParams) WithTimeout(timeout time.Duration) *PostVpsServiceNameChangeContactParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post vps service name change contact params
func (o *PostVpsServiceNameChangeContactParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post vps service name change contact params
func (o *PostVpsServiceNameChangeContactParams) WithContext(ctx context.Context) *PostVpsServiceNameChangeContactParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post vps service name change contact params
func (o *PostVpsServiceNameChangeContactParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post vps service name change contact params
func (o *PostVpsServiceNameChangeContactParams) WithHTTPClient(client *http.Client) *PostVpsServiceNameChangeContactParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post vps service name change contact params
func (o *PostVpsServiceNameChangeContactParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithServiceName adds the serviceName to the post vps service name change contact params
func (o *PostVpsServiceNameChangeContactParams) WithServiceName(serviceName string) *PostVpsServiceNameChangeContactParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the post vps service name change contact params
func (o *PostVpsServiceNameChangeContactParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WithVpsChangeContactPost adds the vpsChangeContactPost to the post vps service name change contact params
func (o *PostVpsServiceNameChangeContactParams) WithVpsChangeContactPost(vpsChangeContactPost *models.PostVpsServiceNameChangeContactParamsBody) *PostVpsServiceNameChangeContactParams {
	o.SetVpsChangeContactPost(vpsChangeContactPost)
	return o
}

// SetVpsChangeContactPost adds the vpsChangeContactPost to the post vps service name change contact params
func (o *PostVpsServiceNameChangeContactParams) SetVpsChangeContactPost(vpsChangeContactPost *models.PostVpsServiceNameChangeContactParamsBody) {
	o.VpsChangeContactPost = vpsChangeContactPost
}

// WriteToRequest writes these params to a swagger request
func (o *PostVpsServiceNameChangeContactParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param serviceName
	if err := r.SetPathParam("serviceName", o.ServiceName); err != nil {
		return err
	}

	if o.VpsChangeContactPost != nil {
		if err := r.SetBodyParam(o.VpsChangeContactPost); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}