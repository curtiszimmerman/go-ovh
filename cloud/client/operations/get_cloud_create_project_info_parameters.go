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

// NewGetCloudCreateProjectInfoParams creates a new GetCloudCreateProjectInfoParams object
// with the default values initialized.
func NewGetCloudCreateProjectInfoParams() *GetCloudCreateProjectInfoParams {

	return &GetCloudCreateProjectInfoParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCloudCreateProjectInfoParamsWithTimeout creates a new GetCloudCreateProjectInfoParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCloudCreateProjectInfoParamsWithTimeout(timeout time.Duration) *GetCloudCreateProjectInfoParams {

	return &GetCloudCreateProjectInfoParams{

		timeout: timeout,
	}
}

// NewGetCloudCreateProjectInfoParamsWithContext creates a new GetCloudCreateProjectInfoParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCloudCreateProjectInfoParamsWithContext(ctx context.Context) *GetCloudCreateProjectInfoParams {

	return &GetCloudCreateProjectInfoParams{

		Context: ctx,
	}
}

// NewGetCloudCreateProjectInfoParamsWithHTTPClient creates a new GetCloudCreateProjectInfoParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCloudCreateProjectInfoParamsWithHTTPClient(client *http.Client) *GetCloudCreateProjectInfoParams {

	return &GetCloudCreateProjectInfoParams{
		HTTPClient: client,
	}
}

/*GetCloudCreateProjectInfoParams contains all the parameters to send to the API endpoint
for the get cloud create project info operation typically these are written to a http.Request
*/
type GetCloudCreateProjectInfoParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get cloud create project info params
func (o *GetCloudCreateProjectInfoParams) WithTimeout(timeout time.Duration) *GetCloudCreateProjectInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cloud create project info params
func (o *GetCloudCreateProjectInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cloud create project info params
func (o *GetCloudCreateProjectInfoParams) WithContext(ctx context.Context) *GetCloudCreateProjectInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cloud create project info params
func (o *GetCloudCreateProjectInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cloud create project info params
func (o *GetCloudCreateProjectInfoParams) WithHTTPClient(client *http.Client) *GetCloudCreateProjectInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cloud create project info params
func (o *GetCloudCreateProjectInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetCloudCreateProjectInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}