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

// NewDeleteCloudProjectServiceNameVolumeSnapshotSnapshotIDParams creates a new DeleteCloudProjectServiceNameVolumeSnapshotSnapshotIDParams object
// with the default values initialized.
func NewDeleteCloudProjectServiceNameVolumeSnapshotSnapshotIDParams() *DeleteCloudProjectServiceNameVolumeSnapshotSnapshotIDParams {
	var ()
	return &DeleteCloudProjectServiceNameVolumeSnapshotSnapshotIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCloudProjectServiceNameVolumeSnapshotSnapshotIDParamsWithTimeout creates a new DeleteCloudProjectServiceNameVolumeSnapshotSnapshotIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteCloudProjectServiceNameVolumeSnapshotSnapshotIDParamsWithTimeout(timeout time.Duration) *DeleteCloudProjectServiceNameVolumeSnapshotSnapshotIDParams {
	var ()
	return &DeleteCloudProjectServiceNameVolumeSnapshotSnapshotIDParams{

		timeout: timeout,
	}
}

// NewDeleteCloudProjectServiceNameVolumeSnapshotSnapshotIDParamsWithContext creates a new DeleteCloudProjectServiceNameVolumeSnapshotSnapshotIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteCloudProjectServiceNameVolumeSnapshotSnapshotIDParamsWithContext(ctx context.Context) *DeleteCloudProjectServiceNameVolumeSnapshotSnapshotIDParams {
	var ()
	return &DeleteCloudProjectServiceNameVolumeSnapshotSnapshotIDParams{

		Context: ctx,
	}
}

// NewDeleteCloudProjectServiceNameVolumeSnapshotSnapshotIDParamsWithHTTPClient creates a new DeleteCloudProjectServiceNameVolumeSnapshotSnapshotIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteCloudProjectServiceNameVolumeSnapshotSnapshotIDParamsWithHTTPClient(client *http.Client) *DeleteCloudProjectServiceNameVolumeSnapshotSnapshotIDParams {
	var ()
	return &DeleteCloudProjectServiceNameVolumeSnapshotSnapshotIDParams{
		HTTPClient: client,
	}
}

/*DeleteCloudProjectServiceNameVolumeSnapshotSnapshotIDParams contains all the parameters to send to the API endpoint
for the delete cloud project service name volume snapshot snapshot ID operation typically these are written to a http.Request
*/
type DeleteCloudProjectServiceNameVolumeSnapshotSnapshotIDParams struct {

	/*ServiceName*/
	ServiceName string
	/*SnapshotID*/
	SnapshotID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete cloud project service name volume snapshot snapshot ID params
func (o *DeleteCloudProjectServiceNameVolumeSnapshotSnapshotIDParams) WithTimeout(timeout time.Duration) *DeleteCloudProjectServiceNameVolumeSnapshotSnapshotIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete cloud project service name volume snapshot snapshot ID params
func (o *DeleteCloudProjectServiceNameVolumeSnapshotSnapshotIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete cloud project service name volume snapshot snapshot ID params
func (o *DeleteCloudProjectServiceNameVolumeSnapshotSnapshotIDParams) WithContext(ctx context.Context) *DeleteCloudProjectServiceNameVolumeSnapshotSnapshotIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete cloud project service name volume snapshot snapshot ID params
func (o *DeleteCloudProjectServiceNameVolumeSnapshotSnapshotIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete cloud project service name volume snapshot snapshot ID params
func (o *DeleteCloudProjectServiceNameVolumeSnapshotSnapshotIDParams) WithHTTPClient(client *http.Client) *DeleteCloudProjectServiceNameVolumeSnapshotSnapshotIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete cloud project service name volume snapshot snapshot ID params
func (o *DeleteCloudProjectServiceNameVolumeSnapshotSnapshotIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithServiceName adds the serviceName to the delete cloud project service name volume snapshot snapshot ID params
func (o *DeleteCloudProjectServiceNameVolumeSnapshotSnapshotIDParams) WithServiceName(serviceName string) *DeleteCloudProjectServiceNameVolumeSnapshotSnapshotIDParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the delete cloud project service name volume snapshot snapshot ID params
func (o *DeleteCloudProjectServiceNameVolumeSnapshotSnapshotIDParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WithSnapshotID adds the snapshotID to the delete cloud project service name volume snapshot snapshot ID params
func (o *DeleteCloudProjectServiceNameVolumeSnapshotSnapshotIDParams) WithSnapshotID(snapshotID string) *DeleteCloudProjectServiceNameVolumeSnapshotSnapshotIDParams {
	o.SetSnapshotID(snapshotID)
	return o
}

// SetSnapshotID adds the snapshotId to the delete cloud project service name volume snapshot snapshot ID params
func (o *DeleteCloudProjectServiceNameVolumeSnapshotSnapshotIDParams) SetSnapshotID(snapshotID string) {
	o.SnapshotID = snapshotID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCloudProjectServiceNameVolumeSnapshotSnapshotIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param serviceName
	if err := r.SetPathParam("serviceName", o.ServiceName); err != nil {
		return err
	}

	// path param snapshotId
	if err := r.SetPathParam("snapshotId", o.SnapshotID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}