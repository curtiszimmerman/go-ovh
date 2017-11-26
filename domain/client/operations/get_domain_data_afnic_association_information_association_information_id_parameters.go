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

// NewGetDomainDataAfnicAssociationInformationAssociationInformationIDParams creates a new GetDomainDataAfnicAssociationInformationAssociationInformationIDParams object
// with the default values initialized.
func NewGetDomainDataAfnicAssociationInformationAssociationInformationIDParams() *GetDomainDataAfnicAssociationInformationAssociationInformationIDParams {
	var ()
	return &GetDomainDataAfnicAssociationInformationAssociationInformationIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDomainDataAfnicAssociationInformationAssociationInformationIDParamsWithTimeout creates a new GetDomainDataAfnicAssociationInformationAssociationInformationIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDomainDataAfnicAssociationInformationAssociationInformationIDParamsWithTimeout(timeout time.Duration) *GetDomainDataAfnicAssociationInformationAssociationInformationIDParams {
	var ()
	return &GetDomainDataAfnicAssociationInformationAssociationInformationIDParams{

		timeout: timeout,
	}
}

// NewGetDomainDataAfnicAssociationInformationAssociationInformationIDParamsWithContext creates a new GetDomainDataAfnicAssociationInformationAssociationInformationIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDomainDataAfnicAssociationInformationAssociationInformationIDParamsWithContext(ctx context.Context) *GetDomainDataAfnicAssociationInformationAssociationInformationIDParams {
	var ()
	return &GetDomainDataAfnicAssociationInformationAssociationInformationIDParams{

		Context: ctx,
	}
}

// NewGetDomainDataAfnicAssociationInformationAssociationInformationIDParamsWithHTTPClient creates a new GetDomainDataAfnicAssociationInformationAssociationInformationIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDomainDataAfnicAssociationInformationAssociationInformationIDParamsWithHTTPClient(client *http.Client) *GetDomainDataAfnicAssociationInformationAssociationInformationIDParams {
	var ()
	return &GetDomainDataAfnicAssociationInformationAssociationInformationIDParams{
		HTTPClient: client,
	}
}

/*GetDomainDataAfnicAssociationInformationAssociationInformationIDParams contains all the parameters to send to the API endpoint
for the get domain data afnic association information association information ID operation typically these are written to a http.Request
*/
type GetDomainDataAfnicAssociationInformationAssociationInformationIDParams struct {

	/*AssociationInformationID*/
	AssociationInformationID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get domain data afnic association information association information ID params
func (o *GetDomainDataAfnicAssociationInformationAssociationInformationIDParams) WithTimeout(timeout time.Duration) *GetDomainDataAfnicAssociationInformationAssociationInformationIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get domain data afnic association information association information ID params
func (o *GetDomainDataAfnicAssociationInformationAssociationInformationIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get domain data afnic association information association information ID params
func (o *GetDomainDataAfnicAssociationInformationAssociationInformationIDParams) WithContext(ctx context.Context) *GetDomainDataAfnicAssociationInformationAssociationInformationIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get domain data afnic association information association information ID params
func (o *GetDomainDataAfnicAssociationInformationAssociationInformationIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get domain data afnic association information association information ID params
func (o *GetDomainDataAfnicAssociationInformationAssociationInformationIDParams) WithHTTPClient(client *http.Client) *GetDomainDataAfnicAssociationInformationAssociationInformationIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get domain data afnic association information association information ID params
func (o *GetDomainDataAfnicAssociationInformationAssociationInformationIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAssociationInformationID adds the associationInformationID to the get domain data afnic association information association information ID params
func (o *GetDomainDataAfnicAssociationInformationAssociationInformationIDParams) WithAssociationInformationID(associationInformationID int64) *GetDomainDataAfnicAssociationInformationAssociationInformationIDParams {
	o.SetAssociationInformationID(associationInformationID)
	return o
}

// SetAssociationInformationID adds the associationInformationId to the get domain data afnic association information association information ID params
func (o *GetDomainDataAfnicAssociationInformationAssociationInformationIDParams) SetAssociationInformationID(associationInformationID int64) {
	o.AssociationInformationID = associationInformationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetDomainDataAfnicAssociationInformationAssociationInformationIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param associationInformationId
	if err := r.SetPathParam("associationInformationId", swag.FormatInt64(o.AssociationInformationID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}