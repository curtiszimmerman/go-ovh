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
)

// NewGetDomainDataSmdSmdIDParams creates a new GetDomainDataSmdSmdIDParams object
// with the default values initialized.
func NewGetDomainDataSmdSmdIDParams() *GetDomainDataSmdSmdIDParams {
	var ()
	return &GetDomainDataSmdSmdIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDomainDataSmdSmdIDParamsWithTimeout creates a new GetDomainDataSmdSmdIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDomainDataSmdSmdIDParamsWithTimeout(timeout time.Duration) *GetDomainDataSmdSmdIDParams {
	var ()
	return &GetDomainDataSmdSmdIDParams{

		timeout: timeout,
	}
}

// NewGetDomainDataSmdSmdIDParamsWithContext creates a new GetDomainDataSmdSmdIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDomainDataSmdSmdIDParamsWithContext(ctx context.Context) *GetDomainDataSmdSmdIDParams {
	var ()
	return &GetDomainDataSmdSmdIDParams{

		Context: ctx,
	}
}

// NewGetDomainDataSmdSmdIDParamsWithHTTPClient creates a new GetDomainDataSmdSmdIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDomainDataSmdSmdIDParamsWithHTTPClient(client *http.Client) *GetDomainDataSmdSmdIDParams {
	var ()
	return &GetDomainDataSmdSmdIDParams{
		HTTPClient: client,
	}
}

/*GetDomainDataSmdSmdIDParams contains all the parameters to send to the API endpoint
for the get domain data smd smd ID operation typically these are written to a http.Request
*/
type GetDomainDataSmdSmdIDParams struct {

	/*SmdID*/
	SmdID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get domain data smd smd ID params
func (o *GetDomainDataSmdSmdIDParams) WithTimeout(timeout time.Duration) *GetDomainDataSmdSmdIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get domain data smd smd ID params
func (o *GetDomainDataSmdSmdIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get domain data smd smd ID params
func (o *GetDomainDataSmdSmdIDParams) WithContext(ctx context.Context) *GetDomainDataSmdSmdIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get domain data smd smd ID params
func (o *GetDomainDataSmdSmdIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get domain data smd smd ID params
func (o *GetDomainDataSmdSmdIDParams) WithHTTPClient(client *http.Client) *GetDomainDataSmdSmdIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get domain data smd smd ID params
func (o *GetDomainDataSmdSmdIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSmdID adds the smdID to the get domain data smd smd ID params
func (o *GetDomainDataSmdSmdIDParams) WithSmdID(smdID int64) *GetDomainDataSmdSmdIDParams {
	o.SetSmdID(smdID)
	return o
}

// SetSmdID adds the smdId to the get domain data smd smd ID params
func (o *GetDomainDataSmdSmdIDParams) SetSmdID(smdID int64) {
	o.SmdID = smdID
}

// WriteToRequest writes these params to a swagger request
func (o *GetDomainDataSmdSmdIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param smdId
	if err := r.SetPathParam("smdId", swag.FormatInt64(o.SmdID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}