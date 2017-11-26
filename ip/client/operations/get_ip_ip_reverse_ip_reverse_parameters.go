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

// NewGetIPIPReverseIPReverseParams creates a new GetIPIPReverseIPReverseParams object
// with the default values initialized.
func NewGetIPIPReverseIPReverseParams() *GetIPIPReverseIPReverseParams {
	var ()
	return &GetIPIPReverseIPReverseParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetIPIPReverseIPReverseParamsWithTimeout creates a new GetIPIPReverseIPReverseParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetIPIPReverseIPReverseParamsWithTimeout(timeout time.Duration) *GetIPIPReverseIPReverseParams {
	var ()
	return &GetIPIPReverseIPReverseParams{

		timeout: timeout,
	}
}

// NewGetIPIPReverseIPReverseParamsWithContext creates a new GetIPIPReverseIPReverseParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetIPIPReverseIPReverseParamsWithContext(ctx context.Context) *GetIPIPReverseIPReverseParams {
	var ()
	return &GetIPIPReverseIPReverseParams{

		Context: ctx,
	}
}

// NewGetIPIPReverseIPReverseParamsWithHTTPClient creates a new GetIPIPReverseIPReverseParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetIPIPReverseIPReverseParamsWithHTTPClient(client *http.Client) *GetIPIPReverseIPReverseParams {
	var ()
	return &GetIPIPReverseIPReverseParams{
		HTTPClient: client,
	}
}

/*GetIPIPReverseIPReverseParams contains all the parameters to send to the API endpoint
for the get IP IP reverse IP reverse operation typically these are written to a http.Request
*/
type GetIPIPReverseIPReverseParams struct {

	/*IP*/
	IP string
	/*IPReverse*/
	IPReverse string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get IP IP reverse IP reverse params
func (o *GetIPIPReverseIPReverseParams) WithTimeout(timeout time.Duration) *GetIPIPReverseIPReverseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get IP IP reverse IP reverse params
func (o *GetIPIPReverseIPReverseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get IP IP reverse IP reverse params
func (o *GetIPIPReverseIPReverseParams) WithContext(ctx context.Context) *GetIPIPReverseIPReverseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get IP IP reverse IP reverse params
func (o *GetIPIPReverseIPReverseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get IP IP reverse IP reverse params
func (o *GetIPIPReverseIPReverseParams) WithHTTPClient(client *http.Client) *GetIPIPReverseIPReverseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get IP IP reverse IP reverse params
func (o *GetIPIPReverseIPReverseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIP adds the ip to the get IP IP reverse IP reverse params
func (o *GetIPIPReverseIPReverseParams) WithIP(ip string) *GetIPIPReverseIPReverseParams {
	o.SetIP(ip)
	return o
}

// SetIP adds the ip to the get IP IP reverse IP reverse params
func (o *GetIPIPReverseIPReverseParams) SetIP(ip string) {
	o.IP = ip
}

// WithIPReverse adds the iPReverse to the get IP IP reverse IP reverse params
func (o *GetIPIPReverseIPReverseParams) WithIPReverse(iPReverse string) *GetIPIPReverseIPReverseParams {
	o.SetIPReverse(iPReverse)
	return o
}

// SetIPReverse adds the ipReverse to the get IP IP reverse IP reverse params
func (o *GetIPIPReverseIPReverseParams) SetIPReverse(iPReverse string) {
	o.IPReverse = iPReverse
}

// WriteToRequest writes these params to a swagger request
func (o *GetIPIPReverseIPReverseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ip
	if err := r.SetPathParam("ip", o.IP); err != nil {
		return err
	}

	// path param ipReverse
	if err := r.SetPathParam("ipReverse", o.IPReverse); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}