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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetIPIPTaskTaskIDParams creates a new GetIPIPTaskTaskIDParams object
// with the default values initialized.
func NewGetIPIPTaskTaskIDParams() *GetIPIPTaskTaskIDParams {
	var ()
	return &GetIPIPTaskTaskIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetIPIPTaskTaskIDParamsWithTimeout creates a new GetIPIPTaskTaskIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetIPIPTaskTaskIDParamsWithTimeout(timeout time.Duration) *GetIPIPTaskTaskIDParams {
	var ()
	return &GetIPIPTaskTaskIDParams{

		timeout: timeout,
	}
}

// NewGetIPIPTaskTaskIDParamsWithContext creates a new GetIPIPTaskTaskIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetIPIPTaskTaskIDParamsWithContext(ctx context.Context) *GetIPIPTaskTaskIDParams {
	var ()
	return &GetIPIPTaskTaskIDParams{

		Context: ctx,
	}
}

// NewGetIPIPTaskTaskIDParamsWithHTTPClient creates a new GetIPIPTaskTaskIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetIPIPTaskTaskIDParamsWithHTTPClient(client *http.Client) *GetIPIPTaskTaskIDParams {
	var ()
	return &GetIPIPTaskTaskIDParams{
		HTTPClient: client,
	}
}

/*GetIPIPTaskTaskIDParams contains all the parameters to send to the API endpoint
for the get IP IP task task ID operation typically these are written to a http.Request
*/
type GetIPIPTaskTaskIDParams struct {

	/*IP*/
	IP string
	/*TaskID*/
	TaskID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get IP IP task task ID params
func (o *GetIPIPTaskTaskIDParams) WithTimeout(timeout time.Duration) *GetIPIPTaskTaskIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get IP IP task task ID params
func (o *GetIPIPTaskTaskIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get IP IP task task ID params
func (o *GetIPIPTaskTaskIDParams) WithContext(ctx context.Context) *GetIPIPTaskTaskIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get IP IP task task ID params
func (o *GetIPIPTaskTaskIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get IP IP task task ID params
func (o *GetIPIPTaskTaskIDParams) WithHTTPClient(client *http.Client) *GetIPIPTaskTaskIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get IP IP task task ID params
func (o *GetIPIPTaskTaskIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIP adds the ip to the get IP IP task task ID params
func (o *GetIPIPTaskTaskIDParams) WithIP(ip string) *GetIPIPTaskTaskIDParams {
	o.SetIP(ip)
	return o
}

// SetIP adds the ip to the get IP IP task task ID params
func (o *GetIPIPTaskTaskIDParams) SetIP(ip string) {
	o.IP = ip
}

// WithTaskID adds the taskID to the get IP IP task task ID params
func (o *GetIPIPTaskTaskIDParams) WithTaskID(taskID int64) *GetIPIPTaskTaskIDParams {
	o.SetTaskID(taskID)
	return o
}

// SetTaskID adds the taskId to the get IP IP task task ID params
func (o *GetIPIPTaskTaskIDParams) SetTaskID(taskID int64) {
	o.TaskID = taskID
}

// WriteToRequest writes these params to a swagger request
func (o *GetIPIPTaskTaskIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ip
	if err := r.SetPathParam("ip", o.IP); err != nil {
		return err
	}

	// path param taskId
	if err := r.SetPathParam("taskId", swag.FormatInt64(o.TaskID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}