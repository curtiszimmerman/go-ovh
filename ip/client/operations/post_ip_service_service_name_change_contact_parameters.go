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

	"github.com/appscode/go-ovh/ip/models"
)

// NewPostIPServiceServiceNameChangeContactParams creates a new PostIPServiceServiceNameChangeContactParams object
// with the default values initialized.
func NewPostIPServiceServiceNameChangeContactParams() *PostIPServiceServiceNameChangeContactParams {
	var ()
	return &PostIPServiceServiceNameChangeContactParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostIPServiceServiceNameChangeContactParamsWithTimeout creates a new PostIPServiceServiceNameChangeContactParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostIPServiceServiceNameChangeContactParamsWithTimeout(timeout time.Duration) *PostIPServiceServiceNameChangeContactParams {
	var ()
	return &PostIPServiceServiceNameChangeContactParams{

		timeout: timeout,
	}
}

// NewPostIPServiceServiceNameChangeContactParamsWithContext creates a new PostIPServiceServiceNameChangeContactParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostIPServiceServiceNameChangeContactParamsWithContext(ctx context.Context) *PostIPServiceServiceNameChangeContactParams {
	var ()
	return &PostIPServiceServiceNameChangeContactParams{

		Context: ctx,
	}
}

// NewPostIPServiceServiceNameChangeContactParamsWithHTTPClient creates a new PostIPServiceServiceNameChangeContactParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostIPServiceServiceNameChangeContactParamsWithHTTPClient(client *http.Client) *PostIPServiceServiceNameChangeContactParams {
	var ()
	return &PostIPServiceServiceNameChangeContactParams{
		HTTPClient: client,
	}
}

/*PostIPServiceServiceNameChangeContactParams contains all the parameters to send to the API endpoint
for the post IP service service name change contact operation typically these are written to a http.Request
*/
type PostIPServiceServiceNameChangeContactParams struct {

	/*IPServiceChangeContactPost*/
	IPServiceChangeContactPost *models.PostIPServiceServiceNameChangeContactParamsBody
	/*ServiceName*/
	ServiceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post IP service service name change contact params
func (o *PostIPServiceServiceNameChangeContactParams) WithTimeout(timeout time.Duration) *PostIPServiceServiceNameChangeContactParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post IP service service name change contact params
func (o *PostIPServiceServiceNameChangeContactParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post IP service service name change contact params
func (o *PostIPServiceServiceNameChangeContactParams) WithContext(ctx context.Context) *PostIPServiceServiceNameChangeContactParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post IP service service name change contact params
func (o *PostIPServiceServiceNameChangeContactParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post IP service service name change contact params
func (o *PostIPServiceServiceNameChangeContactParams) WithHTTPClient(client *http.Client) *PostIPServiceServiceNameChangeContactParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post IP service service name change contact params
func (o *PostIPServiceServiceNameChangeContactParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIPServiceChangeContactPost adds the iPServiceChangeContactPost to the post IP service service name change contact params
func (o *PostIPServiceServiceNameChangeContactParams) WithIPServiceChangeContactPost(iPServiceChangeContactPost *models.PostIPServiceServiceNameChangeContactParamsBody) *PostIPServiceServiceNameChangeContactParams {
	o.SetIPServiceChangeContactPost(iPServiceChangeContactPost)
	return o
}

// SetIPServiceChangeContactPost adds the ipServiceChangeContactPost to the post IP service service name change contact params
func (o *PostIPServiceServiceNameChangeContactParams) SetIPServiceChangeContactPost(iPServiceChangeContactPost *models.PostIPServiceServiceNameChangeContactParamsBody) {
	o.IPServiceChangeContactPost = iPServiceChangeContactPost
}

// WithServiceName adds the serviceName to the post IP service service name change contact params
func (o *PostIPServiceServiceNameChangeContactParams) WithServiceName(serviceName string) *PostIPServiceServiceNameChangeContactParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the post IP service service name change contact params
func (o *PostIPServiceServiceNameChangeContactParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WriteToRequest writes these params to a swagger request
func (o *PostIPServiceServiceNameChangeContactParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IPServiceChangeContactPost != nil {
		if err := r.SetBodyParam(o.IPServiceChangeContactPost); err != nil {
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