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

// NewPostMeAccessRestrictionBackupCodeParams creates a new PostMeAccessRestrictionBackupCodeParams object
// with the default values initialized.
func NewPostMeAccessRestrictionBackupCodeParams() *PostMeAccessRestrictionBackupCodeParams {

	return &PostMeAccessRestrictionBackupCodeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostMeAccessRestrictionBackupCodeParamsWithTimeout creates a new PostMeAccessRestrictionBackupCodeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostMeAccessRestrictionBackupCodeParamsWithTimeout(timeout time.Duration) *PostMeAccessRestrictionBackupCodeParams {

	return &PostMeAccessRestrictionBackupCodeParams{

		timeout: timeout,
	}
}

// NewPostMeAccessRestrictionBackupCodeParamsWithContext creates a new PostMeAccessRestrictionBackupCodeParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostMeAccessRestrictionBackupCodeParamsWithContext(ctx context.Context) *PostMeAccessRestrictionBackupCodeParams {

	return &PostMeAccessRestrictionBackupCodeParams{

		Context: ctx,
	}
}

// NewPostMeAccessRestrictionBackupCodeParamsWithHTTPClient creates a new PostMeAccessRestrictionBackupCodeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostMeAccessRestrictionBackupCodeParamsWithHTTPClient(client *http.Client) *PostMeAccessRestrictionBackupCodeParams {

	return &PostMeAccessRestrictionBackupCodeParams{
		HTTPClient: client,
	}
}

/*PostMeAccessRestrictionBackupCodeParams contains all the parameters to send to the API endpoint
for the post me access restriction backup code operation typically these are written to a http.Request
*/
type PostMeAccessRestrictionBackupCodeParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post me access restriction backup code params
func (o *PostMeAccessRestrictionBackupCodeParams) WithTimeout(timeout time.Duration) *PostMeAccessRestrictionBackupCodeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post me access restriction backup code params
func (o *PostMeAccessRestrictionBackupCodeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post me access restriction backup code params
func (o *PostMeAccessRestrictionBackupCodeParams) WithContext(ctx context.Context) *PostMeAccessRestrictionBackupCodeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post me access restriction backup code params
func (o *PostMeAccessRestrictionBackupCodeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post me access restriction backup code params
func (o *PostMeAccessRestrictionBackupCodeParams) WithHTTPClient(client *http.Client) *PostMeAccessRestrictionBackupCodeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post me access restriction backup code params
func (o *PostMeAccessRestrictionBackupCodeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *PostMeAccessRestrictionBackupCodeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
