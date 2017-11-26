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

// NewGetDomainZoneZoneNameRecordParams creates a new GetDomainZoneZoneNameRecordParams object
// with the default values initialized.
func NewGetDomainZoneZoneNameRecordParams() *GetDomainZoneZoneNameRecordParams {
	var ()
	return &GetDomainZoneZoneNameRecordParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDomainZoneZoneNameRecordParamsWithTimeout creates a new GetDomainZoneZoneNameRecordParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDomainZoneZoneNameRecordParamsWithTimeout(timeout time.Duration) *GetDomainZoneZoneNameRecordParams {
	var ()
	return &GetDomainZoneZoneNameRecordParams{

		timeout: timeout,
	}
}

// NewGetDomainZoneZoneNameRecordParamsWithContext creates a new GetDomainZoneZoneNameRecordParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDomainZoneZoneNameRecordParamsWithContext(ctx context.Context) *GetDomainZoneZoneNameRecordParams {
	var ()
	return &GetDomainZoneZoneNameRecordParams{

		Context: ctx,
	}
}

// NewGetDomainZoneZoneNameRecordParamsWithHTTPClient creates a new GetDomainZoneZoneNameRecordParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDomainZoneZoneNameRecordParamsWithHTTPClient(client *http.Client) *GetDomainZoneZoneNameRecordParams {
	var ()
	return &GetDomainZoneZoneNameRecordParams{
		HTTPClient: client,
	}
}

/*GetDomainZoneZoneNameRecordParams contains all the parameters to send to the API endpoint
for the get domain zone zone name record operation typically these are written to a http.Request
*/
type GetDomainZoneZoneNameRecordParams struct {

	/*FieldType*/
	FieldType *string
	/*SubDomain*/
	SubDomain *string
	/*ZoneName*/
	ZoneName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get domain zone zone name record params
func (o *GetDomainZoneZoneNameRecordParams) WithTimeout(timeout time.Duration) *GetDomainZoneZoneNameRecordParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get domain zone zone name record params
func (o *GetDomainZoneZoneNameRecordParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get domain zone zone name record params
func (o *GetDomainZoneZoneNameRecordParams) WithContext(ctx context.Context) *GetDomainZoneZoneNameRecordParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get domain zone zone name record params
func (o *GetDomainZoneZoneNameRecordParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get domain zone zone name record params
func (o *GetDomainZoneZoneNameRecordParams) WithHTTPClient(client *http.Client) *GetDomainZoneZoneNameRecordParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get domain zone zone name record params
func (o *GetDomainZoneZoneNameRecordParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFieldType adds the fieldType to the get domain zone zone name record params
func (o *GetDomainZoneZoneNameRecordParams) WithFieldType(fieldType *string) *GetDomainZoneZoneNameRecordParams {
	o.SetFieldType(fieldType)
	return o
}

// SetFieldType adds the fieldType to the get domain zone zone name record params
func (o *GetDomainZoneZoneNameRecordParams) SetFieldType(fieldType *string) {
	o.FieldType = fieldType
}

// WithSubDomain adds the subDomain to the get domain zone zone name record params
func (o *GetDomainZoneZoneNameRecordParams) WithSubDomain(subDomain *string) *GetDomainZoneZoneNameRecordParams {
	o.SetSubDomain(subDomain)
	return o
}

// SetSubDomain adds the subDomain to the get domain zone zone name record params
func (o *GetDomainZoneZoneNameRecordParams) SetSubDomain(subDomain *string) {
	o.SubDomain = subDomain
}

// WithZoneName adds the zoneName to the get domain zone zone name record params
func (o *GetDomainZoneZoneNameRecordParams) WithZoneName(zoneName string) *GetDomainZoneZoneNameRecordParams {
	o.SetZoneName(zoneName)
	return o
}

// SetZoneName adds the zoneName to the get domain zone zone name record params
func (o *GetDomainZoneZoneNameRecordParams) SetZoneName(zoneName string) {
	o.ZoneName = zoneName
}

// WriteToRequest writes these params to a swagger request
func (o *GetDomainZoneZoneNameRecordParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FieldType != nil {

		// query param fieldType
		var qrFieldType string
		if o.FieldType != nil {
			qrFieldType = *o.FieldType
		}
		qFieldType := qrFieldType
		if qFieldType != "" {
			if err := r.SetQueryParam("fieldType", qFieldType); err != nil {
				return err
			}
		}

	}

	if o.SubDomain != nil {

		// query param subDomain
		var qrSubDomain string
		if o.SubDomain != nil {
			qrSubDomain = *o.SubDomain
		}
		qSubDomain := qrSubDomain
		if qSubDomain != "" {
			if err := r.SetQueryParam("subDomain", qSubDomain); err != nil {
				return err
			}
		}

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