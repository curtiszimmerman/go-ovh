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

// NewDeleteVpsServiceNameBackupftpAccessIPBlockParams creates a new DeleteVpsServiceNameBackupftpAccessIPBlockParams object
// with the default values initialized.
func NewDeleteVpsServiceNameBackupftpAccessIPBlockParams() *DeleteVpsServiceNameBackupftpAccessIPBlockParams {
	var ()
	return &DeleteVpsServiceNameBackupftpAccessIPBlockParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteVpsServiceNameBackupftpAccessIPBlockParamsWithTimeout creates a new DeleteVpsServiceNameBackupftpAccessIPBlockParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteVpsServiceNameBackupftpAccessIPBlockParamsWithTimeout(timeout time.Duration) *DeleteVpsServiceNameBackupftpAccessIPBlockParams {
	var ()
	return &DeleteVpsServiceNameBackupftpAccessIPBlockParams{

		timeout: timeout,
	}
}

// NewDeleteVpsServiceNameBackupftpAccessIPBlockParamsWithContext creates a new DeleteVpsServiceNameBackupftpAccessIPBlockParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteVpsServiceNameBackupftpAccessIPBlockParamsWithContext(ctx context.Context) *DeleteVpsServiceNameBackupftpAccessIPBlockParams {
	var ()
	return &DeleteVpsServiceNameBackupftpAccessIPBlockParams{

		Context: ctx,
	}
}

// NewDeleteVpsServiceNameBackupftpAccessIPBlockParamsWithHTTPClient creates a new DeleteVpsServiceNameBackupftpAccessIPBlockParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteVpsServiceNameBackupftpAccessIPBlockParamsWithHTTPClient(client *http.Client) *DeleteVpsServiceNameBackupftpAccessIPBlockParams {
	var ()
	return &DeleteVpsServiceNameBackupftpAccessIPBlockParams{
		HTTPClient: client,
	}
}

/*DeleteVpsServiceNameBackupftpAccessIPBlockParams contains all the parameters to send to the API endpoint
for the delete vps service name backupftp access IP block operation typically these are written to a http.Request
*/
type DeleteVpsServiceNameBackupftpAccessIPBlockParams struct {

	/*IPBlock*/
	IPBlock string
	/*ServiceName*/
	ServiceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete vps service name backupftp access IP block params
func (o *DeleteVpsServiceNameBackupftpAccessIPBlockParams) WithTimeout(timeout time.Duration) *DeleteVpsServiceNameBackupftpAccessIPBlockParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete vps service name backupftp access IP block params
func (o *DeleteVpsServiceNameBackupftpAccessIPBlockParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete vps service name backupftp access IP block params
func (o *DeleteVpsServiceNameBackupftpAccessIPBlockParams) WithContext(ctx context.Context) *DeleteVpsServiceNameBackupftpAccessIPBlockParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete vps service name backupftp access IP block params
func (o *DeleteVpsServiceNameBackupftpAccessIPBlockParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete vps service name backupftp access IP block params
func (o *DeleteVpsServiceNameBackupftpAccessIPBlockParams) WithHTTPClient(client *http.Client) *DeleteVpsServiceNameBackupftpAccessIPBlockParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete vps service name backupftp access IP block params
func (o *DeleteVpsServiceNameBackupftpAccessIPBlockParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIPBlock adds the iPBlock to the delete vps service name backupftp access IP block params
func (o *DeleteVpsServiceNameBackupftpAccessIPBlockParams) WithIPBlock(iPBlock string) *DeleteVpsServiceNameBackupftpAccessIPBlockParams {
	o.SetIPBlock(iPBlock)
	return o
}

// SetIPBlock adds the ipBlock to the delete vps service name backupftp access IP block params
func (o *DeleteVpsServiceNameBackupftpAccessIPBlockParams) SetIPBlock(iPBlock string) {
	o.IPBlock = iPBlock
}

// WithServiceName adds the serviceName to the delete vps service name backupftp access IP block params
func (o *DeleteVpsServiceNameBackupftpAccessIPBlockParams) WithServiceName(serviceName string) *DeleteVpsServiceNameBackupftpAccessIPBlockParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the delete vps service name backupftp access IP block params
func (o *DeleteVpsServiceNameBackupftpAccessIPBlockParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteVpsServiceNameBackupftpAccessIPBlockParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}