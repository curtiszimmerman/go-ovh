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

// NewPostMeAccessRestrictionTotpParams creates a new PostMeAccessRestrictionTotpParams object
// with the default values initialized.
func NewPostMeAccessRestrictionTotpParams() *PostMeAccessRestrictionTotpParams {

	return &PostMeAccessRestrictionTotpParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostMeAccessRestrictionTotpParamsWithTimeout creates a new PostMeAccessRestrictionTotpParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostMeAccessRestrictionTotpParamsWithTimeout(timeout time.Duration) *PostMeAccessRestrictionTotpParams {

	return &PostMeAccessRestrictionTotpParams{

		timeout: timeout,
	}
}

// NewPostMeAccessRestrictionTotpParamsWithContext creates a new PostMeAccessRestrictionTotpParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostMeAccessRestrictionTotpParamsWithContext(ctx context.Context) *PostMeAccessRestrictionTotpParams {

	return &PostMeAccessRestrictionTotpParams{

		Context: ctx,
	}
}

// NewPostMeAccessRestrictionTotpParamsWithHTTPClient creates a new PostMeAccessRestrictionTotpParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostMeAccessRestrictionTotpParamsWithHTTPClient(client *http.Client) *PostMeAccessRestrictionTotpParams {

	return &PostMeAccessRestrictionTotpParams{
		HTTPClient: client,
	}
}

/*PostMeAccessRestrictionTotpParams contains all the parameters to send to the API endpoint
for the post me access restriction totp operation typically these are written to a http.Request
*/
type PostMeAccessRestrictionTotpParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post me access restriction totp params
func (o *PostMeAccessRestrictionTotpParams) WithTimeout(timeout time.Duration) *PostMeAccessRestrictionTotpParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post me access restriction totp params
func (o *PostMeAccessRestrictionTotpParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post me access restriction totp params
func (o *PostMeAccessRestrictionTotpParams) WithContext(ctx context.Context) *PostMeAccessRestrictionTotpParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post me access restriction totp params
func (o *PostMeAccessRestrictionTotpParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post me access restriction totp params
func (o *PostMeAccessRestrictionTotpParams) WithHTTPClient(client *http.Client) *PostMeAccessRestrictionTotpParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post me access restriction totp params
func (o *PostMeAccessRestrictionTotpParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *PostMeAccessRestrictionTotpParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
