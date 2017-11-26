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

// NewGetMeDebtAccountDebtDebtIDOperationOperationIDAssociatedObjectParams creates a new GetMeDebtAccountDebtDebtIDOperationOperationIDAssociatedObjectParams object
// with the default values initialized.
func NewGetMeDebtAccountDebtDebtIDOperationOperationIDAssociatedObjectParams() *GetMeDebtAccountDebtDebtIDOperationOperationIDAssociatedObjectParams {
	var ()
	return &GetMeDebtAccountDebtDebtIDOperationOperationIDAssociatedObjectParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMeDebtAccountDebtDebtIDOperationOperationIDAssociatedObjectParamsWithTimeout creates a new GetMeDebtAccountDebtDebtIDOperationOperationIDAssociatedObjectParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMeDebtAccountDebtDebtIDOperationOperationIDAssociatedObjectParamsWithTimeout(timeout time.Duration) *GetMeDebtAccountDebtDebtIDOperationOperationIDAssociatedObjectParams {
	var ()
	return &GetMeDebtAccountDebtDebtIDOperationOperationIDAssociatedObjectParams{

		timeout: timeout,
	}
}

// NewGetMeDebtAccountDebtDebtIDOperationOperationIDAssociatedObjectParamsWithContext creates a new GetMeDebtAccountDebtDebtIDOperationOperationIDAssociatedObjectParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMeDebtAccountDebtDebtIDOperationOperationIDAssociatedObjectParamsWithContext(ctx context.Context) *GetMeDebtAccountDebtDebtIDOperationOperationIDAssociatedObjectParams {
	var ()
	return &GetMeDebtAccountDebtDebtIDOperationOperationIDAssociatedObjectParams{

		Context: ctx,
	}
}

// NewGetMeDebtAccountDebtDebtIDOperationOperationIDAssociatedObjectParamsWithHTTPClient creates a new GetMeDebtAccountDebtDebtIDOperationOperationIDAssociatedObjectParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMeDebtAccountDebtDebtIDOperationOperationIDAssociatedObjectParamsWithHTTPClient(client *http.Client) *GetMeDebtAccountDebtDebtIDOperationOperationIDAssociatedObjectParams {
	var ()
	return &GetMeDebtAccountDebtDebtIDOperationOperationIDAssociatedObjectParams{
		HTTPClient: client,
	}
}

/*GetMeDebtAccountDebtDebtIDOperationOperationIDAssociatedObjectParams contains all the parameters to send to the API endpoint
for the get me debt account debt debt ID operation operation ID associated object operation typically these are written to a http.Request
*/
type GetMeDebtAccountDebtDebtIDOperationOperationIDAssociatedObjectParams struct {

	/*DebtID*/
	DebtID int64
	/*OperationID*/
	OperationID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get me debt account debt debt ID operation operation ID associated object params
func (o *GetMeDebtAccountDebtDebtIDOperationOperationIDAssociatedObjectParams) WithTimeout(timeout time.Duration) *GetMeDebtAccountDebtDebtIDOperationOperationIDAssociatedObjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get me debt account debt debt ID operation operation ID associated object params
func (o *GetMeDebtAccountDebtDebtIDOperationOperationIDAssociatedObjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get me debt account debt debt ID operation operation ID associated object params
func (o *GetMeDebtAccountDebtDebtIDOperationOperationIDAssociatedObjectParams) WithContext(ctx context.Context) *GetMeDebtAccountDebtDebtIDOperationOperationIDAssociatedObjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get me debt account debt debt ID operation operation ID associated object params
func (o *GetMeDebtAccountDebtDebtIDOperationOperationIDAssociatedObjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get me debt account debt debt ID operation operation ID associated object params
func (o *GetMeDebtAccountDebtDebtIDOperationOperationIDAssociatedObjectParams) WithHTTPClient(client *http.Client) *GetMeDebtAccountDebtDebtIDOperationOperationIDAssociatedObjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get me debt account debt debt ID operation operation ID associated object params
func (o *GetMeDebtAccountDebtDebtIDOperationOperationIDAssociatedObjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDebtID adds the debtID to the get me debt account debt debt ID operation operation ID associated object params
func (o *GetMeDebtAccountDebtDebtIDOperationOperationIDAssociatedObjectParams) WithDebtID(debtID int64) *GetMeDebtAccountDebtDebtIDOperationOperationIDAssociatedObjectParams {
	o.SetDebtID(debtID)
	return o
}

// SetDebtID adds the debtId to the get me debt account debt debt ID operation operation ID associated object params
func (o *GetMeDebtAccountDebtDebtIDOperationOperationIDAssociatedObjectParams) SetDebtID(debtID int64) {
	o.DebtID = debtID
}

// WithOperationID adds the operationID to the get me debt account debt debt ID operation operation ID associated object params
func (o *GetMeDebtAccountDebtDebtIDOperationOperationIDAssociatedObjectParams) WithOperationID(operationID int64) *GetMeDebtAccountDebtDebtIDOperationOperationIDAssociatedObjectParams {
	o.SetOperationID(operationID)
	return o
}

// SetOperationID adds the operationId to the get me debt account debt debt ID operation operation ID associated object params
func (o *GetMeDebtAccountDebtDebtIDOperationOperationIDAssociatedObjectParams) SetOperationID(operationID int64) {
	o.OperationID = operationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetMeDebtAccountDebtDebtIDOperationOperationIDAssociatedObjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param debtId
	if err := r.SetPathParam("debtId", swag.FormatInt64(o.DebtID)); err != nil {
		return err
	}

	// path param operationId
	if err := r.SetPathParam("operationId", swag.FormatInt64(o.OperationID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}