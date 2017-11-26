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

// NewGetMeOrderOrderIDDetailsOrderDetailIDParams creates a new GetMeOrderOrderIDDetailsOrderDetailIDParams object
// with the default values initialized.
func NewGetMeOrderOrderIDDetailsOrderDetailIDParams() *GetMeOrderOrderIDDetailsOrderDetailIDParams {
	var ()
	return &GetMeOrderOrderIDDetailsOrderDetailIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMeOrderOrderIDDetailsOrderDetailIDParamsWithTimeout creates a new GetMeOrderOrderIDDetailsOrderDetailIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMeOrderOrderIDDetailsOrderDetailIDParamsWithTimeout(timeout time.Duration) *GetMeOrderOrderIDDetailsOrderDetailIDParams {
	var ()
	return &GetMeOrderOrderIDDetailsOrderDetailIDParams{

		timeout: timeout,
	}
}

// NewGetMeOrderOrderIDDetailsOrderDetailIDParamsWithContext creates a new GetMeOrderOrderIDDetailsOrderDetailIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMeOrderOrderIDDetailsOrderDetailIDParamsWithContext(ctx context.Context) *GetMeOrderOrderIDDetailsOrderDetailIDParams {
	var ()
	return &GetMeOrderOrderIDDetailsOrderDetailIDParams{

		Context: ctx,
	}
}

// NewGetMeOrderOrderIDDetailsOrderDetailIDParamsWithHTTPClient creates a new GetMeOrderOrderIDDetailsOrderDetailIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMeOrderOrderIDDetailsOrderDetailIDParamsWithHTTPClient(client *http.Client) *GetMeOrderOrderIDDetailsOrderDetailIDParams {
	var ()
	return &GetMeOrderOrderIDDetailsOrderDetailIDParams{
		HTTPClient: client,
	}
}

/*GetMeOrderOrderIDDetailsOrderDetailIDParams contains all the parameters to send to the API endpoint
for the get me order order ID details order detail ID operation typically these are written to a http.Request
*/
type GetMeOrderOrderIDDetailsOrderDetailIDParams struct {

	/*OrderDetailID*/
	OrderDetailID int64
	/*OrderID*/
	OrderID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get me order order ID details order detail ID params
func (o *GetMeOrderOrderIDDetailsOrderDetailIDParams) WithTimeout(timeout time.Duration) *GetMeOrderOrderIDDetailsOrderDetailIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get me order order ID details order detail ID params
func (o *GetMeOrderOrderIDDetailsOrderDetailIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get me order order ID details order detail ID params
func (o *GetMeOrderOrderIDDetailsOrderDetailIDParams) WithContext(ctx context.Context) *GetMeOrderOrderIDDetailsOrderDetailIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get me order order ID details order detail ID params
func (o *GetMeOrderOrderIDDetailsOrderDetailIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get me order order ID details order detail ID params
func (o *GetMeOrderOrderIDDetailsOrderDetailIDParams) WithHTTPClient(client *http.Client) *GetMeOrderOrderIDDetailsOrderDetailIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get me order order ID details order detail ID params
func (o *GetMeOrderOrderIDDetailsOrderDetailIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrderDetailID adds the orderDetailID to the get me order order ID details order detail ID params
func (o *GetMeOrderOrderIDDetailsOrderDetailIDParams) WithOrderDetailID(orderDetailID int64) *GetMeOrderOrderIDDetailsOrderDetailIDParams {
	o.SetOrderDetailID(orderDetailID)
	return o
}

// SetOrderDetailID adds the orderDetailId to the get me order order ID details order detail ID params
func (o *GetMeOrderOrderIDDetailsOrderDetailIDParams) SetOrderDetailID(orderDetailID int64) {
	o.OrderDetailID = orderDetailID
}

// WithOrderID adds the orderID to the get me order order ID details order detail ID params
func (o *GetMeOrderOrderIDDetailsOrderDetailIDParams) WithOrderID(orderID int64) *GetMeOrderOrderIDDetailsOrderDetailIDParams {
	o.SetOrderID(orderID)
	return o
}

// SetOrderID adds the orderId to the get me order order ID details order detail ID params
func (o *GetMeOrderOrderIDDetailsOrderDetailIDParams) SetOrderID(orderID int64) {
	o.OrderID = orderID
}

// WriteToRequest writes these params to a swagger request
func (o *GetMeOrderOrderIDDetailsOrderDetailIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param orderDetailId
	if err := r.SetPathParam("orderDetailId", swag.FormatInt64(o.OrderDetailID)); err != nil {
		return err
	}

	// path param orderId
	if err := r.SetPathParam("orderId", swag.FormatInt64(o.OrderID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
