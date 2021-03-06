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

	"github.com/appscode/go-ovh/vps/models"
)

// NewPutVpsServiceNameBackupftpAccessIPBlockParams creates a new PutVpsServiceNameBackupftpAccessIPBlockParams object
// with the default values initialized.
func NewPutVpsServiceNameBackupftpAccessIPBlockParams() *PutVpsServiceNameBackupftpAccessIPBlockParams {
	var ()
	return &PutVpsServiceNameBackupftpAccessIPBlockParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutVpsServiceNameBackupftpAccessIPBlockParamsWithTimeout creates a new PutVpsServiceNameBackupftpAccessIPBlockParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutVpsServiceNameBackupftpAccessIPBlockParamsWithTimeout(timeout time.Duration) *PutVpsServiceNameBackupftpAccessIPBlockParams {
	var ()
	return &PutVpsServiceNameBackupftpAccessIPBlockParams{

		timeout: timeout,
	}
}

// NewPutVpsServiceNameBackupftpAccessIPBlockParamsWithContext creates a new PutVpsServiceNameBackupftpAccessIPBlockParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutVpsServiceNameBackupftpAccessIPBlockParamsWithContext(ctx context.Context) *PutVpsServiceNameBackupftpAccessIPBlockParams {
	var ()
	return &PutVpsServiceNameBackupftpAccessIPBlockParams{

		Context: ctx,
	}
}

// NewPutVpsServiceNameBackupftpAccessIPBlockParamsWithHTTPClient creates a new PutVpsServiceNameBackupftpAccessIPBlockParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutVpsServiceNameBackupftpAccessIPBlockParamsWithHTTPClient(client *http.Client) *PutVpsServiceNameBackupftpAccessIPBlockParams {
	var ()
	return &PutVpsServiceNameBackupftpAccessIPBlockParams{
		HTTPClient: client,
	}
}

/*PutVpsServiceNameBackupftpAccessIPBlockParams contains all the parameters to send to the API endpoint
for the put vps service name backupftp access IP block operation typically these are written to a http.Request
*/
type PutVpsServiceNameBackupftpAccessIPBlockParams struct {

	/*IPBlock*/
	IPBlock string
	/*ServiceName*/
	ServiceName string
	/*VpsBackupftpAccessPut*/
	VpsBackupftpAccessPut *models.DedicatedServerBackupFtpACL

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put vps service name backupftp access IP block params
func (o *PutVpsServiceNameBackupftpAccessIPBlockParams) WithTimeout(timeout time.Duration) *PutVpsServiceNameBackupftpAccessIPBlockParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put vps service name backupftp access IP block params
func (o *PutVpsServiceNameBackupftpAccessIPBlockParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put vps service name backupftp access IP block params
func (o *PutVpsServiceNameBackupftpAccessIPBlockParams) WithContext(ctx context.Context) *PutVpsServiceNameBackupftpAccessIPBlockParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put vps service name backupftp access IP block params
func (o *PutVpsServiceNameBackupftpAccessIPBlockParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put vps service name backupftp access IP block params
func (o *PutVpsServiceNameBackupftpAccessIPBlockParams) WithHTTPClient(client *http.Client) *PutVpsServiceNameBackupftpAccessIPBlockParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put vps service name backupftp access IP block params
func (o *PutVpsServiceNameBackupftpAccessIPBlockParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIPBlock adds the iPBlock to the put vps service name backupftp access IP block params
func (o *PutVpsServiceNameBackupftpAccessIPBlockParams) WithIPBlock(iPBlock string) *PutVpsServiceNameBackupftpAccessIPBlockParams {
	o.SetIPBlock(iPBlock)
	return o
}

// SetIPBlock adds the ipBlock to the put vps service name backupftp access IP block params
func (o *PutVpsServiceNameBackupftpAccessIPBlockParams) SetIPBlock(iPBlock string) {
	o.IPBlock = iPBlock
}

// WithServiceName adds the serviceName to the put vps service name backupftp access IP block params
func (o *PutVpsServiceNameBackupftpAccessIPBlockParams) WithServiceName(serviceName string) *PutVpsServiceNameBackupftpAccessIPBlockParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the put vps service name backupftp access IP block params
func (o *PutVpsServiceNameBackupftpAccessIPBlockParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WithVpsBackupftpAccessPut adds the vpsBackupftpAccessPut to the put vps service name backupftp access IP block params
func (o *PutVpsServiceNameBackupftpAccessIPBlockParams) WithVpsBackupftpAccessPut(vpsBackupftpAccessPut *models.DedicatedServerBackupFtpACL) *PutVpsServiceNameBackupftpAccessIPBlockParams {
	o.SetVpsBackupftpAccessPut(vpsBackupftpAccessPut)
	return o
}

// SetVpsBackupftpAccessPut adds the vpsBackupftpAccessPut to the put vps service name backupftp access IP block params
func (o *PutVpsServiceNameBackupftpAccessIPBlockParams) SetVpsBackupftpAccessPut(vpsBackupftpAccessPut *models.DedicatedServerBackupFtpACL) {
	o.VpsBackupftpAccessPut = vpsBackupftpAccessPut
}

// WriteToRequest writes these params to a swagger request
func (o *PutVpsServiceNameBackupftpAccessIPBlockParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ipBlock
	if err := r.SetPathParam("ipBlock", o.IPBlock); err != nil {
		return err
	}

	// path param serviceName
	if err := r.SetPathParam("serviceName", o.ServiceName); err != nil {
		return err
	}

	if o.VpsBackupftpAccessPut != nil {
		if err := r.SetBodyParam(o.VpsBackupftpAccessPut); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
