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

// NewGetDomainRulesParams creates a new GetDomainRulesParams object
// with the default values initialized.
func NewGetDomainRulesParams() *GetDomainRulesParams {
	var ()
	return &GetDomainRulesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDomainRulesParamsWithTimeout creates a new GetDomainRulesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDomainRulesParamsWithTimeout(timeout time.Duration) *GetDomainRulesParams {
	var ()
	return &GetDomainRulesParams{

		timeout: timeout,
	}
}

// NewGetDomainRulesParamsWithContext creates a new GetDomainRulesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDomainRulesParamsWithContext(ctx context.Context) *GetDomainRulesParams {
	var ()
	return &GetDomainRulesParams{

		Context: ctx,
	}
}

// NewGetDomainRulesParamsWithHTTPClient creates a new GetDomainRulesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDomainRulesParamsWithHTTPClient(client *http.Client) *GetDomainRulesParams {
	var ()
	return &GetDomainRulesParams{
		HTTPClient: client,
	}
}

/*GetDomainRulesParams contains all the parameters to send to the API endpoint
for the get domain rules operation typically these are written to a http.Request
*/
type GetDomainRulesParams struct {

	/*CartID*/
	CartID string
	/*ItemID*/
	ItemID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get domain rules params
func (o *GetDomainRulesParams) WithTimeout(timeout time.Duration) *GetDomainRulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get domain rules params
func (o *GetDomainRulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get domain rules params
func (o *GetDomainRulesParams) WithContext(ctx context.Context) *GetDomainRulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get domain rules params
func (o *GetDomainRulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get domain rules params
func (o *GetDomainRulesParams) WithHTTPClient(client *http.Client) *GetDomainRulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get domain rules params
func (o *GetDomainRulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCartID adds the cartID to the get domain rules params
func (o *GetDomainRulesParams) WithCartID(cartID string) *GetDomainRulesParams {
	o.SetCartID(cartID)
	return o
}

// SetCartID adds the cartId to the get domain rules params
func (o *GetDomainRulesParams) SetCartID(cartID string) {
	o.CartID = cartID
}

// WithItemID adds the itemID to the get domain rules params
func (o *GetDomainRulesParams) WithItemID(itemID int64) *GetDomainRulesParams {
	o.SetItemID(itemID)
	return o
}

// SetItemID adds the itemId to the get domain rules params
func (o *GetDomainRulesParams) SetItemID(itemID int64) {
	o.ItemID = itemID
}

// WriteToRequest writes these params to a swagger request
func (o *GetDomainRulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param cartId
	qrCartID := o.CartID
	qCartID := qrCartID
	if qCartID != "" {
		if err := r.SetQueryParam("cartId", qCartID); err != nil {
			return err
		}
	}

	// query param itemId
	qrItemID := o.ItemID
	qItemID := swag.FormatInt64(qrItemID)
	if qItemID != "" {
		if err := r.SetQueryParam("itemId", qItemID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}