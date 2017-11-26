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

// NewDeleteDomainZoneZoneNameDnssecParams creates a new DeleteDomainZoneZoneNameDnssecParams object
// with the default values initialized.
func NewDeleteDomainZoneZoneNameDnssecParams() *DeleteDomainZoneZoneNameDnssecParams {
	var ()
	return &DeleteDomainZoneZoneNameDnssecParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDomainZoneZoneNameDnssecParamsWithTimeout creates a new DeleteDomainZoneZoneNameDnssecParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteDomainZoneZoneNameDnssecParamsWithTimeout(timeout time.Duration) *DeleteDomainZoneZoneNameDnssecParams {
	var ()
	return &DeleteDomainZoneZoneNameDnssecParams{

		timeout: timeout,
	}
}

// NewDeleteDomainZoneZoneNameDnssecParamsWithContext creates a new DeleteDomainZoneZoneNameDnssecParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteDomainZoneZoneNameDnssecParamsWithContext(ctx context.Context) *DeleteDomainZoneZoneNameDnssecParams {
	var ()
	return &DeleteDomainZoneZoneNameDnssecParams{

		Context: ctx,
	}
}

// NewDeleteDomainZoneZoneNameDnssecParamsWithHTTPClient creates a new DeleteDomainZoneZoneNameDnssecParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteDomainZoneZoneNameDnssecParamsWithHTTPClient(client *http.Client) *DeleteDomainZoneZoneNameDnssecParams {
	var ()
	return &DeleteDomainZoneZoneNameDnssecParams{
		HTTPClient: client,
	}
}

/*DeleteDomainZoneZoneNameDnssecParams contains all the parameters to send to the API endpoint
for the delete domain zone zone name dnssec operation typically these are written to a http.Request
*/
type DeleteDomainZoneZoneNameDnssecParams struct {

	/*ZoneName*/
	ZoneName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete domain zone zone name dnssec params
func (o *DeleteDomainZoneZoneNameDnssecParams) WithTimeout(timeout time.Duration) *DeleteDomainZoneZoneNameDnssecParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete domain zone zone name dnssec params
func (o *DeleteDomainZoneZoneNameDnssecParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete domain zone zone name dnssec params
func (o *DeleteDomainZoneZoneNameDnssecParams) WithContext(ctx context.Context) *DeleteDomainZoneZoneNameDnssecParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete domain zone zone name dnssec params
func (o *DeleteDomainZoneZoneNameDnssecParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete domain zone zone name dnssec params
func (o *DeleteDomainZoneZoneNameDnssecParams) WithHTTPClient(client *http.Client) *DeleteDomainZoneZoneNameDnssecParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete domain zone zone name dnssec params
func (o *DeleteDomainZoneZoneNameDnssecParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithZoneName adds the zoneName to the delete domain zone zone name dnssec params
func (o *DeleteDomainZoneZoneNameDnssecParams) WithZoneName(zoneName string) *DeleteDomainZoneZoneNameDnssecParams {
	o.SetZoneName(zoneName)
	return o
}

// SetZoneName adds the zoneName to the delete domain zone zone name dnssec params
func (o *DeleteDomainZoneZoneNameDnssecParams) SetZoneName(zoneName string) {
	o.ZoneName = zoneName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDomainZoneZoneNameDnssecParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param zoneName
	if err := r.SetPathParam("zoneName", o.ZoneName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}