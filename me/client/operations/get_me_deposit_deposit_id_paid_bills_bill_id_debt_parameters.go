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

// NewGetMeDepositDepositIDPaidBillsBillIDDebtParams creates a new GetMeDepositDepositIDPaidBillsBillIDDebtParams object
// with the default values initialized.
func NewGetMeDepositDepositIDPaidBillsBillIDDebtParams() *GetMeDepositDepositIDPaidBillsBillIDDebtParams {
	var ()
	return &GetMeDepositDepositIDPaidBillsBillIDDebtParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMeDepositDepositIDPaidBillsBillIDDebtParamsWithTimeout creates a new GetMeDepositDepositIDPaidBillsBillIDDebtParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMeDepositDepositIDPaidBillsBillIDDebtParamsWithTimeout(timeout time.Duration) *GetMeDepositDepositIDPaidBillsBillIDDebtParams {
	var ()
	return &GetMeDepositDepositIDPaidBillsBillIDDebtParams{

		timeout: timeout,
	}
}

// NewGetMeDepositDepositIDPaidBillsBillIDDebtParamsWithContext creates a new GetMeDepositDepositIDPaidBillsBillIDDebtParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMeDepositDepositIDPaidBillsBillIDDebtParamsWithContext(ctx context.Context) *GetMeDepositDepositIDPaidBillsBillIDDebtParams {
	var ()
	return &GetMeDepositDepositIDPaidBillsBillIDDebtParams{

		Context: ctx,
	}
}

// NewGetMeDepositDepositIDPaidBillsBillIDDebtParamsWithHTTPClient creates a new GetMeDepositDepositIDPaidBillsBillIDDebtParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMeDepositDepositIDPaidBillsBillIDDebtParamsWithHTTPClient(client *http.Client) *GetMeDepositDepositIDPaidBillsBillIDDebtParams {
	var ()
	return &GetMeDepositDepositIDPaidBillsBillIDDebtParams{
		HTTPClient: client,
	}
}

/*GetMeDepositDepositIDPaidBillsBillIDDebtParams contains all the parameters to send to the API endpoint
for the get me deposit deposit ID paid bills bill ID debt operation typically these are written to a http.Request
*/
type GetMeDepositDepositIDPaidBillsBillIDDebtParams struct {

	/*BillID*/
	BillID string
	/*DepositID*/
	DepositID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get me deposit deposit ID paid bills bill ID debt params
func (o *GetMeDepositDepositIDPaidBillsBillIDDebtParams) WithTimeout(timeout time.Duration) *GetMeDepositDepositIDPaidBillsBillIDDebtParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get me deposit deposit ID paid bills bill ID debt params
func (o *GetMeDepositDepositIDPaidBillsBillIDDebtParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get me deposit deposit ID paid bills bill ID debt params
func (o *GetMeDepositDepositIDPaidBillsBillIDDebtParams) WithContext(ctx context.Context) *GetMeDepositDepositIDPaidBillsBillIDDebtParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get me deposit deposit ID paid bills bill ID debt params
func (o *GetMeDepositDepositIDPaidBillsBillIDDebtParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get me deposit deposit ID paid bills bill ID debt params
func (o *GetMeDepositDepositIDPaidBillsBillIDDebtParams) WithHTTPClient(client *http.Client) *GetMeDepositDepositIDPaidBillsBillIDDebtParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get me deposit deposit ID paid bills bill ID debt params
func (o *GetMeDepositDepositIDPaidBillsBillIDDebtParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBillID adds the billID to the get me deposit deposit ID paid bills bill ID debt params
func (o *GetMeDepositDepositIDPaidBillsBillIDDebtParams) WithBillID(billID string) *GetMeDepositDepositIDPaidBillsBillIDDebtParams {
	o.SetBillID(billID)
	return o
}

// SetBillID adds the billId to the get me deposit deposit ID paid bills bill ID debt params
func (o *GetMeDepositDepositIDPaidBillsBillIDDebtParams) SetBillID(billID string) {
	o.BillID = billID
}

// WithDepositID adds the depositID to the get me deposit deposit ID paid bills bill ID debt params
func (o *GetMeDepositDepositIDPaidBillsBillIDDebtParams) WithDepositID(depositID string) *GetMeDepositDepositIDPaidBillsBillIDDebtParams {
	o.SetDepositID(depositID)
	return o
}

// SetDepositID adds the depositId to the get me deposit deposit ID paid bills bill ID debt params
func (o *GetMeDepositDepositIDPaidBillsBillIDDebtParams) SetDepositID(depositID string) {
	o.DepositID = depositID
}

// WriteToRequest writes these params to a swagger request
func (o *GetMeDepositDepositIDPaidBillsBillIDDebtParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param billId
	if err := r.SetPathParam("billId", o.BillID); err != nil {
		return err
	}

	// path param depositId
	if err := r.SetPathParam("depositId", o.DepositID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
