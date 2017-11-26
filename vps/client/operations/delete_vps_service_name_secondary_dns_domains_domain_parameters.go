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

// NewDeleteVpsServiceNameSecondaryDNSDomainsDomainParams creates a new DeleteVpsServiceNameSecondaryDNSDomainsDomainParams object
// with the default values initialized.
func NewDeleteVpsServiceNameSecondaryDNSDomainsDomainParams() *DeleteVpsServiceNameSecondaryDNSDomainsDomainParams {
	var ()
	return &DeleteVpsServiceNameSecondaryDNSDomainsDomainParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteVpsServiceNameSecondaryDNSDomainsDomainParamsWithTimeout creates a new DeleteVpsServiceNameSecondaryDNSDomainsDomainParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteVpsServiceNameSecondaryDNSDomainsDomainParamsWithTimeout(timeout time.Duration) *DeleteVpsServiceNameSecondaryDNSDomainsDomainParams {
	var ()
	return &DeleteVpsServiceNameSecondaryDNSDomainsDomainParams{

		timeout: timeout,
	}
}

// NewDeleteVpsServiceNameSecondaryDNSDomainsDomainParamsWithContext creates a new DeleteVpsServiceNameSecondaryDNSDomainsDomainParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteVpsServiceNameSecondaryDNSDomainsDomainParamsWithContext(ctx context.Context) *DeleteVpsServiceNameSecondaryDNSDomainsDomainParams {
	var ()
	return &DeleteVpsServiceNameSecondaryDNSDomainsDomainParams{

		Context: ctx,
	}
}

// NewDeleteVpsServiceNameSecondaryDNSDomainsDomainParamsWithHTTPClient creates a new DeleteVpsServiceNameSecondaryDNSDomainsDomainParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteVpsServiceNameSecondaryDNSDomainsDomainParamsWithHTTPClient(client *http.Client) *DeleteVpsServiceNameSecondaryDNSDomainsDomainParams {
	var ()
	return &DeleteVpsServiceNameSecondaryDNSDomainsDomainParams{
		HTTPClient: client,
	}
}

/*DeleteVpsServiceNameSecondaryDNSDomainsDomainParams contains all the parameters to send to the API endpoint
for the delete vps service name secondary DNS domains domain operation typically these are written to a http.Request
*/
type DeleteVpsServiceNameSecondaryDNSDomainsDomainParams struct {

	/*Domain*/
	Domain string
	/*ServiceName*/
	ServiceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete vps service name secondary DNS domains domain params
func (o *DeleteVpsServiceNameSecondaryDNSDomainsDomainParams) WithTimeout(timeout time.Duration) *DeleteVpsServiceNameSecondaryDNSDomainsDomainParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete vps service name secondary DNS domains domain params
func (o *DeleteVpsServiceNameSecondaryDNSDomainsDomainParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete vps service name secondary DNS domains domain params
func (o *DeleteVpsServiceNameSecondaryDNSDomainsDomainParams) WithContext(ctx context.Context) *DeleteVpsServiceNameSecondaryDNSDomainsDomainParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete vps service name secondary DNS domains domain params
func (o *DeleteVpsServiceNameSecondaryDNSDomainsDomainParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete vps service name secondary DNS domains domain params
func (o *DeleteVpsServiceNameSecondaryDNSDomainsDomainParams) WithHTTPClient(client *http.Client) *DeleteVpsServiceNameSecondaryDNSDomainsDomainParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete vps service name secondary DNS domains domain params
func (o *DeleteVpsServiceNameSecondaryDNSDomainsDomainParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomain adds the domain to the delete vps service name secondary DNS domains domain params
func (o *DeleteVpsServiceNameSecondaryDNSDomainsDomainParams) WithDomain(domain string) *DeleteVpsServiceNameSecondaryDNSDomainsDomainParams {
	o.SetDomain(domain)
	return o
}

// SetDomain adds the domain to the delete vps service name secondary DNS domains domain params
func (o *DeleteVpsServiceNameSecondaryDNSDomainsDomainParams) SetDomain(domain string) {
	o.Domain = domain
}

// WithServiceName adds the serviceName to the delete vps service name secondary DNS domains domain params
func (o *DeleteVpsServiceNameSecondaryDNSDomainsDomainParams) WithServiceName(serviceName string) *DeleteVpsServiceNameSecondaryDNSDomainsDomainParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the delete vps service name secondary DNS domains domain params
func (o *DeleteVpsServiceNameSecondaryDNSDomainsDomainParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteVpsServiceNameSecondaryDNSDomainsDomainParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param domain
	if err := r.SetPathParam("domain", o.Domain); err != nil {
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