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

// NewPutMeOvhAccountOvhAccountIDParams creates a new PutMeOvhAccountOvhAccountIDParams object
// with the default values initialized.
func NewPutMeOvhAccountOvhAccountIDParams() *PutMeOvhAccountOvhAccountIDParams {
	var ()
	return &PutMeOvhAccountOvhAccountIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutMeOvhAccountOvhAccountIDParamsWithTimeout creates a new PutMeOvhAccountOvhAccountIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutMeOvhAccountOvhAccountIDParamsWithTimeout(timeout time.Duration) *PutMeOvhAccountOvhAccountIDParams {
	var ()
	return &PutMeOvhAccountOvhAccountIDParams{

		timeout: timeout,
	}
}

// NewPutMeOvhAccountOvhAccountIDParamsWithContext creates a new PutMeOvhAccountOvhAccountIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutMeOvhAccountOvhAccountIDParamsWithContext(ctx context.Context) *PutMeOvhAccountOvhAccountIDParams {
	var ()
	return &PutMeOvhAccountOvhAccountIDParams{

		Context: ctx,
	}
}

// NewPutMeOvhAccountOvhAccountIDParamsWithHTTPClient creates a new PutMeOvhAccountOvhAccountIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutMeOvhAccountOvhAccountIDParamsWithHTTPClient(client *http.Client) *PutMeOvhAccountOvhAccountIDParams {
	var ()
	return &PutMeOvhAccountOvhAccountIDParams{
		HTTPClient: client,
	}
}

/*PutMeOvhAccountOvhAccountIDParams contains all the parameters to send to the API endpoint
for the put me ovh account ovh account ID operation typically these are written to a http.Request
*/
type PutMeOvhAccountOvhAccountIDParams struct {

	/*MeOvhAccountPut*/
	MeOvhAccountPut *models.BillingOvhAccount
	/*OvhAccountID*/
	OvhAccountID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put me ovh account ovh account ID params
func (o *PutMeOvhAccountOvhAccountIDParams) WithTimeout(timeout time.Duration) *PutMeOvhAccountOvhAccountIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put me ovh account ovh account ID params
func (o *PutMeOvhAccountOvhAccountIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put me ovh account ovh account ID params
func (o *PutMeOvhAccountOvhAccountIDParams) WithContext(ctx context.Context) *PutMeOvhAccountOvhAccountIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put me ovh account ovh account ID params
func (o *PutMeOvhAccountOvhAccountIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put me ovh account ovh account ID params
func (o *PutMeOvhAccountOvhAccountIDParams) WithHTTPClient(client *http.Client) *PutMeOvhAccountOvhAccountIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put me ovh account ovh account ID params
func (o *PutMeOvhAccountOvhAccountIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMeOvhAccountPut adds the meOvhAccountPut to the put me ovh account ovh account ID params
func (o *PutMeOvhAccountOvhAccountIDParams) WithMeOvhAccountPut(meOvhAccountPut *models.BillingOvhAccount) *PutMeOvhAccountOvhAccountIDParams {
	o.SetMeOvhAccountPut(meOvhAccountPut)
	return o
}

// SetMeOvhAccountPut adds the meOvhAccountPut to the put me ovh account ovh account ID params
func (o *PutMeOvhAccountOvhAccountIDParams) SetMeOvhAccountPut(meOvhAccountPut *models.BillingOvhAccount) {
	o.MeOvhAccountPut = meOvhAccountPut
}

// WithOvhAccountID adds the ovhAccountID to the put me ovh account ovh account ID params
func (o *PutMeOvhAccountOvhAccountIDParams) WithOvhAccountID(ovhAccountID string) *PutMeOvhAccountOvhAccountIDParams {
	o.SetOvhAccountID(ovhAccountID)
	return o
}

// SetOvhAccountID adds the ovhAccountId to the put me ovh account ovh account ID params
func (o *PutMeOvhAccountOvhAccountIDParams) SetOvhAccountID(ovhAccountID string) {
	o.OvhAccountID = ovhAccountID
}

// WriteToRequest writes these params to a swagger request
func (o *PutMeOvhAccountOvhAccountIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.MeOvhAccountPut != nil {
		if err := r.SetBodyParam(o.MeOvhAccountPut); err != nil {
			return err
		}
	}

	// path param ovhAccountId
	if err := r.SetPathParam("ovhAccountId", o.OvhAccountID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
