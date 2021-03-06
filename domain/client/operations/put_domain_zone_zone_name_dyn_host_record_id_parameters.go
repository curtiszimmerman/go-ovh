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

	"github.com/appscode/go-ovh/domain/models"
)

// NewPutDomainZoneZoneNameDynHostRecordIDParams creates a new PutDomainZoneZoneNameDynHostRecordIDParams object
// with the default values initialized.
func NewPutDomainZoneZoneNameDynHostRecordIDParams() *PutDomainZoneZoneNameDynHostRecordIDParams {
	var ()
	return &PutDomainZoneZoneNameDynHostRecordIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutDomainZoneZoneNameDynHostRecordIDParamsWithTimeout creates a new PutDomainZoneZoneNameDynHostRecordIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutDomainZoneZoneNameDynHostRecordIDParamsWithTimeout(timeout time.Duration) *PutDomainZoneZoneNameDynHostRecordIDParams {
	var ()
	return &PutDomainZoneZoneNameDynHostRecordIDParams{

		timeout: timeout,
	}
}

// NewPutDomainZoneZoneNameDynHostRecordIDParamsWithContext creates a new PutDomainZoneZoneNameDynHostRecordIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutDomainZoneZoneNameDynHostRecordIDParamsWithContext(ctx context.Context) *PutDomainZoneZoneNameDynHostRecordIDParams {
	var ()
	return &PutDomainZoneZoneNameDynHostRecordIDParams{

		Context: ctx,
	}
}

// NewPutDomainZoneZoneNameDynHostRecordIDParamsWithHTTPClient creates a new PutDomainZoneZoneNameDynHostRecordIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutDomainZoneZoneNameDynHostRecordIDParamsWithHTTPClient(client *http.Client) *PutDomainZoneZoneNameDynHostRecordIDParams {
	var ()
	return &PutDomainZoneZoneNameDynHostRecordIDParams{
		HTTPClient: client,
	}
}

/*PutDomainZoneZoneNameDynHostRecordIDParams contains all the parameters to send to the API endpoint
for the put domain zone zone name dyn host record ID operation typically these are written to a http.Request
*/
type PutDomainZoneZoneNameDynHostRecordIDParams struct {

	/*DomainZoneDynHostRecordPut*/
	DomainZoneDynHostRecordPut *models.DomainZoneDynHostRecord
	/*ID*/
	ID int64
	/*ZoneName*/
	ZoneName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put domain zone zone name dyn host record ID params
func (o *PutDomainZoneZoneNameDynHostRecordIDParams) WithTimeout(timeout time.Duration) *PutDomainZoneZoneNameDynHostRecordIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put domain zone zone name dyn host record ID params
func (o *PutDomainZoneZoneNameDynHostRecordIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put domain zone zone name dyn host record ID params
func (o *PutDomainZoneZoneNameDynHostRecordIDParams) WithContext(ctx context.Context) *PutDomainZoneZoneNameDynHostRecordIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put domain zone zone name dyn host record ID params
func (o *PutDomainZoneZoneNameDynHostRecordIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put domain zone zone name dyn host record ID params
func (o *PutDomainZoneZoneNameDynHostRecordIDParams) WithHTTPClient(client *http.Client) *PutDomainZoneZoneNameDynHostRecordIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put domain zone zone name dyn host record ID params
func (o *PutDomainZoneZoneNameDynHostRecordIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomainZoneDynHostRecordPut adds the domainZoneDynHostRecordPut to the put domain zone zone name dyn host record ID params
func (o *PutDomainZoneZoneNameDynHostRecordIDParams) WithDomainZoneDynHostRecordPut(domainZoneDynHostRecordPut *models.DomainZoneDynHostRecord) *PutDomainZoneZoneNameDynHostRecordIDParams {
	o.SetDomainZoneDynHostRecordPut(domainZoneDynHostRecordPut)
	return o
}

// SetDomainZoneDynHostRecordPut adds the domainZoneDynHostRecordPut to the put domain zone zone name dyn host record ID params
func (o *PutDomainZoneZoneNameDynHostRecordIDParams) SetDomainZoneDynHostRecordPut(domainZoneDynHostRecordPut *models.DomainZoneDynHostRecord) {
	o.DomainZoneDynHostRecordPut = domainZoneDynHostRecordPut
}

// WithID adds the id to the put domain zone zone name dyn host record ID params
func (o *PutDomainZoneZoneNameDynHostRecordIDParams) WithID(id int64) *PutDomainZoneZoneNameDynHostRecordIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put domain zone zone name dyn host record ID params
func (o *PutDomainZoneZoneNameDynHostRecordIDParams) SetID(id int64) {
	o.ID = id
}

// WithZoneName adds the zoneName to the put domain zone zone name dyn host record ID params
func (o *PutDomainZoneZoneNameDynHostRecordIDParams) WithZoneName(zoneName string) *PutDomainZoneZoneNameDynHostRecordIDParams {
	o.SetZoneName(zoneName)
	return o
}

// SetZoneName adds the zoneName to the put domain zone zone name dyn host record ID params
func (o *PutDomainZoneZoneNameDynHostRecordIDParams) SetZoneName(zoneName string) {
	o.ZoneName = zoneName
}

// WriteToRequest writes these params to a swagger request
func (o *PutDomainZoneZoneNameDynHostRecordIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DomainZoneDynHostRecordPut != nil {
		if err := r.SetBodyParam(o.DomainZoneDynHostRecordPut); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	// path param zoneName
	if err := r.SetPathParam("zoneName", o.ZoneName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
