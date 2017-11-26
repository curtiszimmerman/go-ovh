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

	"github.com/appscode/go-ovh/me/models"
)

// NewPostMeAccessRestrictionTotpIDEnableParams creates a new PostMeAccessRestrictionTotpIDEnableParams object
// with the default values initialized.
func NewPostMeAccessRestrictionTotpIDEnableParams() *PostMeAccessRestrictionTotpIDEnableParams {
	var ()
	return &PostMeAccessRestrictionTotpIDEnableParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostMeAccessRestrictionTotpIDEnableParamsWithTimeout creates a new PostMeAccessRestrictionTotpIDEnableParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostMeAccessRestrictionTotpIDEnableParamsWithTimeout(timeout time.Duration) *PostMeAccessRestrictionTotpIDEnableParams {
	var ()
	return &PostMeAccessRestrictionTotpIDEnableParams{

		timeout: timeout,
	}
}

// NewPostMeAccessRestrictionTotpIDEnableParamsWithContext creates a new PostMeAccessRestrictionTotpIDEnableParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostMeAccessRestrictionTotpIDEnableParamsWithContext(ctx context.Context) *PostMeAccessRestrictionTotpIDEnableParams {
	var ()
	return &PostMeAccessRestrictionTotpIDEnableParams{

		Context: ctx,
	}
}

// NewPostMeAccessRestrictionTotpIDEnableParamsWithHTTPClient creates a new PostMeAccessRestrictionTotpIDEnableParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostMeAccessRestrictionTotpIDEnableParamsWithHTTPClient(client *http.Client) *PostMeAccessRestrictionTotpIDEnableParams {
	var ()
	return &PostMeAccessRestrictionTotpIDEnableParams{
		HTTPClient: client,
	}
}

/*PostMeAccessRestrictionTotpIDEnableParams contains all the parameters to send to the API endpoint
for the post me access restriction totp ID enable operation typically these are written to a http.Request
*/
type PostMeAccessRestrictionTotpIDEnableParams struct {

	/*ID*/
	ID int64
	/*MeAccessRestrictionTotpEnablePost*/
	MeAccessRestrictionTotpEnablePost *models.PostMeAccessRestrictionTotpIDEnableParamsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post me access restriction totp ID enable params
func (o *PostMeAccessRestrictionTotpIDEnableParams) WithTimeout(timeout time.Duration) *PostMeAccessRestrictionTotpIDEnableParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post me access restriction totp ID enable params
func (o *PostMeAccessRestrictionTotpIDEnableParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post me access restriction totp ID enable params
func (o *PostMeAccessRestrictionTotpIDEnableParams) WithContext(ctx context.Context) *PostMeAccessRestrictionTotpIDEnableParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post me access restriction totp ID enable params
func (o *PostMeAccessRestrictionTotpIDEnableParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post me access restriction totp ID enable params
func (o *PostMeAccessRestrictionTotpIDEnableParams) WithHTTPClient(client *http.Client) *PostMeAccessRestrictionTotpIDEnableParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post me access restriction totp ID enable params
func (o *PostMeAccessRestrictionTotpIDEnableParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the post me access restriction totp ID enable params
func (o *PostMeAccessRestrictionTotpIDEnableParams) WithID(id int64) *PostMeAccessRestrictionTotpIDEnableParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post me access restriction totp ID enable params
func (o *PostMeAccessRestrictionTotpIDEnableParams) SetID(id int64) {
	o.ID = id
}

// WithMeAccessRestrictionTotpEnablePost adds the meAccessRestrictionTotpEnablePost to the post me access restriction totp ID enable params
func (o *PostMeAccessRestrictionTotpIDEnableParams) WithMeAccessRestrictionTotpEnablePost(meAccessRestrictionTotpEnablePost *models.PostMeAccessRestrictionTotpIDEnableParamsBody) *PostMeAccessRestrictionTotpIDEnableParams {
	o.SetMeAccessRestrictionTotpEnablePost(meAccessRestrictionTotpEnablePost)
	return o
}

// SetMeAccessRestrictionTotpEnablePost adds the meAccessRestrictionTotpEnablePost to the post me access restriction totp ID enable params
func (o *PostMeAccessRestrictionTotpIDEnableParams) SetMeAccessRestrictionTotpEnablePost(meAccessRestrictionTotpEnablePost *models.PostMeAccessRestrictionTotpIDEnableParamsBody) {
	o.MeAccessRestrictionTotpEnablePost = meAccessRestrictionTotpEnablePost
}

// WriteToRequest writes these params to a swagger request
func (o *PostMeAccessRestrictionTotpIDEnableParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if o.MeAccessRestrictionTotpEnablePost != nil {
		if err := r.SetBodyParam(o.MeAccessRestrictionTotpEnablePost); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
