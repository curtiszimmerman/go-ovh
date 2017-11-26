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

// NewGetMeOvhAccountOvhAccountIDMovementsParams creates a new GetMeOvhAccountOvhAccountIDMovementsParams object
// with the default values initialized.
func NewGetMeOvhAccountOvhAccountIDMovementsParams() *GetMeOvhAccountOvhAccountIDMovementsParams {
	var ()
	return &GetMeOvhAccountOvhAccountIDMovementsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMeOvhAccountOvhAccountIDMovementsParamsWithTimeout creates a new GetMeOvhAccountOvhAccountIDMovementsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMeOvhAccountOvhAccountIDMovementsParamsWithTimeout(timeout time.Duration) *GetMeOvhAccountOvhAccountIDMovementsParams {
	var ()
	return &GetMeOvhAccountOvhAccountIDMovementsParams{

		timeout: timeout,
	}
}

// NewGetMeOvhAccountOvhAccountIDMovementsParamsWithContext creates a new GetMeOvhAccountOvhAccountIDMovementsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMeOvhAccountOvhAccountIDMovementsParamsWithContext(ctx context.Context) *GetMeOvhAccountOvhAccountIDMovementsParams {
	var ()
	return &GetMeOvhAccountOvhAccountIDMovementsParams{

		Context: ctx,
	}
}

// NewGetMeOvhAccountOvhAccountIDMovementsParamsWithHTTPClient creates a new GetMeOvhAccountOvhAccountIDMovementsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMeOvhAccountOvhAccountIDMovementsParamsWithHTTPClient(client *http.Client) *GetMeOvhAccountOvhAccountIDMovementsParams {
	var ()
	return &GetMeOvhAccountOvhAccountIDMovementsParams{
		HTTPClient: client,
	}
}

/*GetMeOvhAccountOvhAccountIDMovementsParams contains all the parameters to send to the API endpoint
for the get me ovh account ovh account ID movements operation typically these are written to a http.Request
*/
type GetMeOvhAccountOvhAccountIDMovementsParams struct {

	/*DateFrom*/
	DateFrom *strfmt.DateTime
	/*DateTo*/
	DateTo *strfmt.DateTime
	/*OvhAccountID*/
	OvhAccountID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get me ovh account ovh account ID movements params
func (o *GetMeOvhAccountOvhAccountIDMovementsParams) WithTimeout(timeout time.Duration) *GetMeOvhAccountOvhAccountIDMovementsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get me ovh account ovh account ID movements params
func (o *GetMeOvhAccountOvhAccountIDMovementsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get me ovh account ovh account ID movements params
func (o *GetMeOvhAccountOvhAccountIDMovementsParams) WithContext(ctx context.Context) *GetMeOvhAccountOvhAccountIDMovementsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get me ovh account ovh account ID movements params
func (o *GetMeOvhAccountOvhAccountIDMovementsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get me ovh account ovh account ID movements params
func (o *GetMeOvhAccountOvhAccountIDMovementsParams) WithHTTPClient(client *http.Client) *GetMeOvhAccountOvhAccountIDMovementsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get me ovh account ovh account ID movements params
func (o *GetMeOvhAccountOvhAccountIDMovementsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDateFrom adds the dateFrom to the get me ovh account ovh account ID movements params
func (o *GetMeOvhAccountOvhAccountIDMovementsParams) WithDateFrom(dateFrom *strfmt.DateTime) *GetMeOvhAccountOvhAccountIDMovementsParams {
	o.SetDateFrom(dateFrom)
	return o
}

// SetDateFrom adds the dateFrom to the get me ovh account ovh account ID movements params
func (o *GetMeOvhAccountOvhAccountIDMovementsParams) SetDateFrom(dateFrom *strfmt.DateTime) {
	o.DateFrom = dateFrom
}

// WithDateTo adds the dateTo to the get me ovh account ovh account ID movements params
func (o *GetMeOvhAccountOvhAccountIDMovementsParams) WithDateTo(dateTo *strfmt.DateTime) *GetMeOvhAccountOvhAccountIDMovementsParams {
	o.SetDateTo(dateTo)
	return o
}

// SetDateTo adds the dateTo to the get me ovh account ovh account ID movements params
func (o *GetMeOvhAccountOvhAccountIDMovementsParams) SetDateTo(dateTo *strfmt.DateTime) {
	o.DateTo = dateTo
}

// WithOvhAccountID adds the ovhAccountID to the get me ovh account ovh account ID movements params
func (o *GetMeOvhAccountOvhAccountIDMovementsParams) WithOvhAccountID(ovhAccountID string) *GetMeOvhAccountOvhAccountIDMovementsParams {
	o.SetOvhAccountID(ovhAccountID)
	return o
}

// SetOvhAccountID adds the ovhAccountId to the get me ovh account ovh account ID movements params
func (o *GetMeOvhAccountOvhAccountIDMovementsParams) SetOvhAccountID(ovhAccountID string) {
	o.OvhAccountID = ovhAccountID
}

// WriteToRequest writes these params to a swagger request
func (o *GetMeOvhAccountOvhAccountIDMovementsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DateFrom != nil {

		// query param date.from
		var qrDateFrom strfmt.DateTime
		if o.DateFrom != nil {
			qrDateFrom = *o.DateFrom
		}
		qDateFrom := qrDateFrom.String()
		if qDateFrom != "" {
			if err := r.SetQueryParam("date.from", qDateFrom); err != nil {
				return err
			}
		}

	}

	if o.DateTo != nil {

		// query param date.to
		var qrDateTo strfmt.DateTime
		if o.DateTo != nil {
			qrDateTo = *o.DateTo
		}
		qDateTo := qrDateTo.String()
		if qDateTo != "" {
			if err := r.SetQueryParam("date.to", qDateTo); err != nil {
				return err
			}
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