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

// NewPostMeDepositDepositIDPaidBillsBillIDDebtPayParams creates a new PostMeDepositDepositIDPaidBillsBillIDDebtPayParams object
// with the default values initialized.
func NewPostMeDepositDepositIDPaidBillsBillIDDebtPayParams() *PostMeDepositDepositIDPaidBillsBillIDDebtPayParams {
	var ()
	return &PostMeDepositDepositIDPaidBillsBillIDDebtPayParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostMeDepositDepositIDPaidBillsBillIDDebtPayParamsWithTimeout creates a new PostMeDepositDepositIDPaidBillsBillIDDebtPayParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostMeDepositDepositIDPaidBillsBillIDDebtPayParamsWithTimeout(timeout time.Duration) *PostMeDepositDepositIDPaidBillsBillIDDebtPayParams {
	var ()
	return &PostMeDepositDepositIDPaidBillsBillIDDebtPayParams{

		timeout: timeout,
	}
}

// NewPostMeDepositDepositIDPaidBillsBillIDDebtPayParamsWithContext creates a new PostMeDepositDepositIDPaidBillsBillIDDebtPayParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostMeDepositDepositIDPaidBillsBillIDDebtPayParamsWithContext(ctx context.Context) *PostMeDepositDepositIDPaidBillsBillIDDebtPayParams {
	var ()
	return &PostMeDepositDepositIDPaidBillsBillIDDebtPayParams{

		Context: ctx,
	}
}

// NewPostMeDepositDepositIDPaidBillsBillIDDebtPayParamsWithHTTPClient creates a new PostMeDepositDepositIDPaidBillsBillIDDebtPayParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostMeDepositDepositIDPaidBillsBillIDDebtPayParamsWithHTTPClient(client *http.Client) *PostMeDepositDepositIDPaidBillsBillIDDebtPayParams {
	var ()
	return &PostMeDepositDepositIDPaidBillsBillIDDebtPayParams{
		HTTPClient: client,
	}
}

/*PostMeDepositDepositIDPaidBillsBillIDDebtPayParams contains all the parameters to send to the API endpoint
for the post me deposit deposit ID paid bills bill ID debt pay operation typically these are written to a http.Request
*/
type PostMeDepositDepositIDPaidBillsBillIDDebtPayParams struct {

	/*BillID*/
	BillID string
	/*DepositID*/
	DepositID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post me deposit deposit ID paid bills bill ID debt pay params
func (o *PostMeDepositDepositIDPaidBillsBillIDDebtPayParams) WithTimeout(timeout time.Duration) *PostMeDepositDepositIDPaidBillsBillIDDebtPayParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post me deposit deposit ID paid bills bill ID debt pay params
func (o *PostMeDepositDepositIDPaidBillsBillIDDebtPayParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post me deposit deposit ID paid bills bill ID debt pay params
func (o *PostMeDepositDepositIDPaidBillsBillIDDebtPayParams) WithContext(ctx context.Context) *PostMeDepositDepositIDPaidBillsBillIDDebtPayParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post me deposit deposit ID paid bills bill ID debt pay params
func (o *PostMeDepositDepositIDPaidBillsBillIDDebtPayParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post me deposit deposit ID paid bills bill ID debt pay params
func (o *PostMeDepositDepositIDPaidBillsBillIDDebtPayParams) WithHTTPClient(client *http.Client) *PostMeDepositDepositIDPaidBillsBillIDDebtPayParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post me deposit deposit ID paid bills bill ID debt pay params
func (o *PostMeDepositDepositIDPaidBillsBillIDDebtPayParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBillID adds the billID to the post me deposit deposit ID paid bills bill ID debt pay params
func (o *PostMeDepositDepositIDPaidBillsBillIDDebtPayParams) WithBillID(billID string) *PostMeDepositDepositIDPaidBillsBillIDDebtPayParams {
	o.SetBillID(billID)
	return o
}

// SetBillID adds the billId to the post me deposit deposit ID paid bills bill ID debt pay params
func (o *PostMeDepositDepositIDPaidBillsBillIDDebtPayParams) SetBillID(billID string) {
	o.BillID = billID
}

// WithDepositID adds the depositID to the post me deposit deposit ID paid bills bill ID debt pay params
func (o *PostMeDepositDepositIDPaidBillsBillIDDebtPayParams) WithDepositID(depositID string) *PostMeDepositDepositIDPaidBillsBillIDDebtPayParams {
	o.SetDepositID(depositID)
	return o
}

// SetDepositID adds the depositId to the post me deposit deposit ID paid bills bill ID debt pay params
func (o *PostMeDepositDepositIDPaidBillsBillIDDebtPayParams) SetDepositID(depositID string) {
	o.DepositID = depositID
}

// WriteToRequest writes these params to a swagger request
func (o *PostMeDepositDepositIDPaidBillsBillIDDebtPayParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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