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

	"github.com/appscode/go-ovh/iploadbalancing/models"
)

// NewPostIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams creates a new PostIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams object
// with the default values initialized.
func NewPostIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams() *PostIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams {
	var ()
	return &PostIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostIPLoadbalancingServiceNameHTTPFarmFarmIDServerParamsWithTimeout creates a new PostIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostIPLoadbalancingServiceNameHTTPFarmFarmIDServerParamsWithTimeout(timeout time.Duration) *PostIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams {
	var ()
	return &PostIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams{

		timeout: timeout,
	}
}

// NewPostIPLoadbalancingServiceNameHTTPFarmFarmIDServerParamsWithContext creates a new PostIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostIPLoadbalancingServiceNameHTTPFarmFarmIDServerParamsWithContext(ctx context.Context) *PostIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams {
	var ()
	return &PostIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams{

		Context: ctx,
	}
}

// NewPostIPLoadbalancingServiceNameHTTPFarmFarmIDServerParamsWithHTTPClient creates a new PostIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostIPLoadbalancingServiceNameHTTPFarmFarmIDServerParamsWithHTTPClient(client *http.Client) *PostIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams {
	var ()
	return &PostIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams{
		HTTPClient: client,
	}
}

/*PostIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams contains all the parameters to send to the API endpoint
for the post IP loadbalancing service name HTTP farm farm ID server operation typically these are written to a http.Request
*/
type PostIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams struct {

	/*FarmID*/
	FarmID int64
	/*IPLBHTTPFarmServerPost*/
	IPLBHTTPFarmServerPost *models.PostIPLoadbalancingServiceNameHTTPFarmFarmIDServerParamsBody
	/*ServiceName*/
	ServiceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post IP loadbalancing service name HTTP farm farm ID server params
func (o *PostIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams) WithTimeout(timeout time.Duration) *PostIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post IP loadbalancing service name HTTP farm farm ID server params
func (o *PostIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post IP loadbalancing service name HTTP farm farm ID server params
func (o *PostIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams) WithContext(ctx context.Context) *PostIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post IP loadbalancing service name HTTP farm farm ID server params
func (o *PostIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post IP loadbalancing service name HTTP farm farm ID server params
func (o *PostIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams) WithHTTPClient(client *http.Client) *PostIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post IP loadbalancing service name HTTP farm farm ID server params
func (o *PostIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFarmID adds the farmID to the post IP loadbalancing service name HTTP farm farm ID server params
func (o *PostIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams) WithFarmID(farmID int64) *PostIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams {
	o.SetFarmID(farmID)
	return o
}

// SetFarmID adds the farmId to the post IP loadbalancing service name HTTP farm farm ID server params
func (o *PostIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams) SetFarmID(farmID int64) {
	o.FarmID = farmID
}

// WithIPLBHTTPFarmServerPost adds the iPLBHTTPFarmServerPost to the post IP loadbalancing service name HTTP farm farm ID server params
func (o *PostIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams) WithIPLBHTTPFarmServerPost(iPLBHTTPFarmServerPost *models.PostIPLoadbalancingServiceNameHTTPFarmFarmIDServerParamsBody) *PostIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams {
	o.SetIPLBHTTPFarmServerPost(iPLBHTTPFarmServerPost)
	return o
}

// SetIPLBHTTPFarmServerPost adds the iplbHttpFarmServerPost to the post IP loadbalancing service name HTTP farm farm ID server params
func (o *PostIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams) SetIPLBHTTPFarmServerPost(iPLBHTTPFarmServerPost *models.PostIPLoadbalancingServiceNameHTTPFarmFarmIDServerParamsBody) {
	o.IPLBHTTPFarmServerPost = iPLBHTTPFarmServerPost
}

// WithServiceName adds the serviceName to the post IP loadbalancing service name HTTP farm farm ID server params
func (o *PostIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams) WithServiceName(serviceName string) *PostIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the post IP loadbalancing service name HTTP farm farm ID server params
func (o *PostIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WriteToRequest writes these params to a swagger request
func (o *PostIPLoadbalancingServiceNameHTTPFarmFarmIDServerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param farmId
	if err := r.SetPathParam("farmId", swag.FormatInt64(o.FarmID)); err != nil {
		return err
	}

	if o.IPLBHTTPFarmServerPost != nil {
		if err := r.SetBodyParam(o.IPLBHTTPFarmServerPost); err != nil {
			return err
		}
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