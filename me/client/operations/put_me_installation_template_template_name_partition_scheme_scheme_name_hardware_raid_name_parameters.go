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

	"github.com/appscode/go-ovh/me/models"
)

// NewPutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameParams creates a new PutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameParams object
// with the default values initialized.
func NewPutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameParams() *PutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameParams {
	var ()
	return &PutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameParamsWithTimeout creates a new PutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameParamsWithTimeout(timeout time.Duration) *PutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameParams {
	var ()
	return &PutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameParams{

		timeout: timeout,
	}
}

// NewPutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameParamsWithContext creates a new PutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameParamsWithContext(ctx context.Context) *PutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameParams {
	var ()
	return &PutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameParams{

		Context: ctx,
	}
}

// NewPutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameParamsWithHTTPClient creates a new PutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameParamsWithHTTPClient(client *http.Client) *PutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameParams {
	var ()
	return &PutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameParams{
		HTTPClient: client,
	}
}

/*PutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameParams contains all the parameters to send to the API endpoint
for the put me installation template template name partition scheme scheme name hardware raid name operation typically these are written to a http.Request
*/
type PutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameParams struct {

	/*MeInstallationTemplatePartitionSchemeHardwareRaidPut*/
	MeInstallationTemplatePartitionSchemeHardwareRaidPut *models.DedicatedInstallationTemplateHardwareRaid
	/*Name*/
	Name string
	/*SchemeName*/
	SchemeName string
	/*TemplateName*/
	TemplateName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put me installation template template name partition scheme scheme name hardware raid name params
func (o *PutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameParams) WithTimeout(timeout time.Duration) *PutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put me installation template template name partition scheme scheme name hardware raid name params
func (o *PutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put me installation template template name partition scheme scheme name hardware raid name params
func (o *PutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameParams) WithContext(ctx context.Context) *PutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put me installation template template name partition scheme scheme name hardware raid name params
func (o *PutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put me installation template template name partition scheme scheme name hardware raid name params
func (o *PutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameParams) WithHTTPClient(client *http.Client) *PutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put me installation template template name partition scheme scheme name hardware raid name params
func (o *PutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMeInstallationTemplatePartitionSchemeHardwareRaidPut adds the meInstallationTemplatePartitionSchemeHardwareRaidPut to the put me installation template template name partition scheme scheme name hardware raid name params
func (o *PutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameParams) WithMeInstallationTemplatePartitionSchemeHardwareRaidPut(meInstallationTemplatePartitionSchemeHardwareRaidPut *models.DedicatedInstallationTemplateHardwareRaid) *PutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameParams {
	o.SetMeInstallationTemplatePartitionSchemeHardwareRaidPut(meInstallationTemplatePartitionSchemeHardwareRaidPut)
	return o
}

// SetMeInstallationTemplatePartitionSchemeHardwareRaidPut adds the meInstallationTemplatePartitionSchemeHardwareRaidPut to the put me installation template template name partition scheme scheme name hardware raid name params
func (o *PutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameParams) SetMeInstallationTemplatePartitionSchemeHardwareRaidPut(meInstallationTemplatePartitionSchemeHardwareRaidPut *models.DedicatedInstallationTemplateHardwareRaid) {
	o.MeInstallationTemplatePartitionSchemeHardwareRaidPut = meInstallationTemplatePartitionSchemeHardwareRaidPut
}

// WithName adds the name to the put me installation template template name partition scheme scheme name hardware raid name params
func (o *PutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameParams) WithName(name string) *PutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the put me installation template template name partition scheme scheme name hardware raid name params
func (o *PutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameParams) SetName(name string) {
	o.Name = name
}

// WithSchemeName adds the schemeName to the put me installation template template name partition scheme scheme name hardware raid name params
func (o *PutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameParams) WithSchemeName(schemeName string) *PutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameParams {
	o.SetSchemeName(schemeName)
	return o
}

// SetSchemeName adds the schemeName to the put me installation template template name partition scheme scheme name hardware raid name params
func (o *PutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameParams) SetSchemeName(schemeName string) {
	o.SchemeName = schemeName
}

// WithTemplateName adds the templateName to the put me installation template template name partition scheme scheme name hardware raid name params
func (o *PutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameParams) WithTemplateName(templateName string) *PutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameParams {
	o.SetTemplateName(templateName)
	return o
}

// SetTemplateName adds the templateName to the put me installation template template name partition scheme scheme name hardware raid name params
func (o *PutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameParams) SetTemplateName(templateName string) {
	o.TemplateName = templateName
}

// WriteToRequest writes these params to a swagger request
func (o *PutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.MeInstallationTemplatePartitionSchemeHardwareRaidPut != nil {
		if err := r.SetBodyParam(o.MeInstallationTemplatePartitionSchemeHardwareRaidPut); err != nil {
			return err
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param schemeName
	if err := r.SetPathParam("schemeName", o.SchemeName); err != nil {
		return err
	}

	// path param templateName
	if err := r.SetPathParam("templateName", o.TemplateName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}