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

	"github.com/appscode/go-ovh/domain/models"
)

// NewPostDomainServiceNameActivateZoneParams creates a new PostDomainServiceNameActivateZoneParams object
// with the default values initialized.
func NewPostDomainServiceNameActivateZoneParams() *PostDomainServiceNameActivateZoneParams {
	var ()
	return &PostDomainServiceNameActivateZoneParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostDomainServiceNameActivateZoneParamsWithTimeout creates a new PostDomainServiceNameActivateZoneParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostDomainServiceNameActivateZoneParamsWithTimeout(timeout time.Duration) *PostDomainServiceNameActivateZoneParams {
	var ()
	return &PostDomainServiceNameActivateZoneParams{

		timeout: timeout,
	}
}

// NewPostDomainServiceNameActivateZoneParamsWithContext creates a new PostDomainServiceNameActivateZoneParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostDomainServiceNameActivateZoneParamsWithContext(ctx context.Context) *PostDomainServiceNameActivateZoneParams {
	var ()
	return &PostDomainServiceNameActivateZoneParams{

		Context: ctx,
	}
}

// NewPostDomainServiceNameActivateZoneParamsWithHTTPClient creates a new PostDomainServiceNameActivateZoneParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostDomainServiceNameActivateZoneParamsWithHTTPClient(client *http.Client) *PostDomainServiceNameActivateZoneParams {
	var ()
	return &PostDomainServiceNameActivateZoneParams{
		HTTPClient: client,
	}
}

/*PostDomainServiceNameActivateZoneParams contains all the parameters to send to the API endpoint
for the post domain service name activate zone operation typically these are written to a http.Request
*/
type PostDomainServiceNameActivateZoneParams struct {

	/*DomainActivateZonePost*/
	DomainActivateZonePost *models.PostDomainServiceNameActivateZoneParamsBody
	/*ServiceName*/
	ServiceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post domain service name activate zone params
func (o *PostDomainServiceNameActivateZoneParams) WithTimeout(timeout time.Duration) *PostDomainServiceNameActivateZoneParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post domain service name activate zone params
func (o *PostDomainServiceNameActivateZoneParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post domain service name activate zone params
func (o *PostDomainServiceNameActivateZoneParams) WithContext(ctx context.Context) *PostDomainServiceNameActivateZoneParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post domain service name activate zone params
func (o *PostDomainServiceNameActivateZoneParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post domain service name activate zone params
func (o *PostDomainServiceNameActivateZoneParams) WithHTTPClient(client *http.Client) *PostDomainServiceNameActivateZoneParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post domain service name activate zone params
func (o *PostDomainServiceNameActivateZoneParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomainActivateZonePost adds the domainActivateZonePost to the post domain service name activate zone params
func (o *PostDomainServiceNameActivateZoneParams) WithDomainActivateZonePost(domainActivateZonePost *models.PostDomainServiceNameActivateZoneParamsBody) *PostDomainServiceNameActivateZoneParams {
	o.SetDomainActivateZonePost(domainActivateZonePost)
	return o
}

// SetDomainActivateZonePost adds the domainActivateZonePost to the post domain service name activate zone params
func (o *PostDomainServiceNameActivateZoneParams) SetDomainActivateZonePost(domainActivateZonePost *models.PostDomainServiceNameActivateZoneParamsBody) {
	o.DomainActivateZonePost = domainActivateZonePost
}

// WithServiceName adds the serviceName to the post domain service name activate zone params
func (o *PostDomainServiceNameActivateZoneParams) WithServiceName(serviceName string) *PostDomainServiceNameActivateZoneParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the post domain service name activate zone params
func (o *PostDomainServiceNameActivateZoneParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WriteToRequest writes these params to a swagger request
func (o *PostDomainServiceNameActivateZoneParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DomainActivateZonePost != nil {
		if err := r.SetBodyParam(o.DomainActivateZonePost); err != nil {
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
