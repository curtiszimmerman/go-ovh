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

// NewGetCloudServiceNamePcaPcaServiceNameSessionsParams creates a new GetCloudServiceNamePcaPcaServiceNameSessionsParams object
// with the default values initialized.
func NewGetCloudServiceNamePcaPcaServiceNameSessionsParams() *GetCloudServiceNamePcaPcaServiceNameSessionsParams {
	var ()
	return &GetCloudServiceNamePcaPcaServiceNameSessionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCloudServiceNamePcaPcaServiceNameSessionsParamsWithTimeout creates a new GetCloudServiceNamePcaPcaServiceNameSessionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCloudServiceNamePcaPcaServiceNameSessionsParamsWithTimeout(timeout time.Duration) *GetCloudServiceNamePcaPcaServiceNameSessionsParams {
	var ()
	return &GetCloudServiceNamePcaPcaServiceNameSessionsParams{

		timeout: timeout,
	}
}

// NewGetCloudServiceNamePcaPcaServiceNameSessionsParamsWithContext creates a new GetCloudServiceNamePcaPcaServiceNameSessionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCloudServiceNamePcaPcaServiceNameSessionsParamsWithContext(ctx context.Context) *GetCloudServiceNamePcaPcaServiceNameSessionsParams {
	var ()
	return &GetCloudServiceNamePcaPcaServiceNameSessionsParams{

		Context: ctx,
	}
}

// NewGetCloudServiceNamePcaPcaServiceNameSessionsParamsWithHTTPClient creates a new GetCloudServiceNamePcaPcaServiceNameSessionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCloudServiceNamePcaPcaServiceNameSessionsParamsWithHTTPClient(client *http.Client) *GetCloudServiceNamePcaPcaServiceNameSessionsParams {
	var ()
	return &GetCloudServiceNamePcaPcaServiceNameSessionsParams{
		HTTPClient: client,
	}
}

/*GetCloudServiceNamePcaPcaServiceNameSessionsParams contains all the parameters to send to the API endpoint
for the get cloud service name pca pca service name sessions operation typically these are written to a http.Request
*/
type GetCloudServiceNamePcaPcaServiceNameSessionsParams struct {

	/*Name*/
	Name *string
	/*PcaServiceName*/
	PcaServiceName string
	/*ServiceName*/
	ServiceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get cloud service name pca pca service name sessions params
func (o *GetCloudServiceNamePcaPcaServiceNameSessionsParams) WithTimeout(timeout time.Duration) *GetCloudServiceNamePcaPcaServiceNameSessionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cloud service name pca pca service name sessions params
func (o *GetCloudServiceNamePcaPcaServiceNameSessionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cloud service name pca pca service name sessions params
func (o *GetCloudServiceNamePcaPcaServiceNameSessionsParams) WithContext(ctx context.Context) *GetCloudServiceNamePcaPcaServiceNameSessionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cloud service name pca pca service name sessions params
func (o *GetCloudServiceNamePcaPcaServiceNameSessionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cloud service name pca pca service name sessions params
func (o *GetCloudServiceNamePcaPcaServiceNameSessionsParams) WithHTTPClient(client *http.Client) *GetCloudServiceNamePcaPcaServiceNameSessionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cloud service name pca pca service name sessions params
func (o *GetCloudServiceNamePcaPcaServiceNameSessionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get cloud service name pca pca service name sessions params
func (o *GetCloudServiceNamePcaPcaServiceNameSessionsParams) WithName(name *string) *GetCloudServiceNamePcaPcaServiceNameSessionsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get cloud service name pca pca service name sessions params
func (o *GetCloudServiceNamePcaPcaServiceNameSessionsParams) SetName(name *string) {
	o.Name = name
}

// WithPcaServiceName adds the pcaServiceName to the get cloud service name pca pca service name sessions params
func (o *GetCloudServiceNamePcaPcaServiceNameSessionsParams) WithPcaServiceName(pcaServiceName string) *GetCloudServiceNamePcaPcaServiceNameSessionsParams {
	o.SetPcaServiceName(pcaServiceName)
	return o
}

// SetPcaServiceName adds the pcaServiceName to the get cloud service name pca pca service name sessions params
func (o *GetCloudServiceNamePcaPcaServiceNameSessionsParams) SetPcaServiceName(pcaServiceName string) {
	o.PcaServiceName = pcaServiceName
}

// WithServiceName adds the serviceName to the get cloud service name pca pca service name sessions params
func (o *GetCloudServiceNamePcaPcaServiceNameSessionsParams) WithServiceName(serviceName string) *GetCloudServiceNamePcaPcaServiceNameSessionsParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the get cloud service name pca pca service name sessions params
func (o *GetCloudServiceNamePcaPcaServiceNameSessionsParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WriteToRequest writes these params to a swagger request
func (o *GetCloudServiceNamePcaPcaServiceNameSessionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Name != nil {

		// query param name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}

	}

	// path param pcaServiceName
	if err := r.SetPathParam("pcaServiceName", o.PcaServiceName); err != nil {
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