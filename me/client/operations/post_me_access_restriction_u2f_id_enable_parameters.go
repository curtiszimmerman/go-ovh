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

// NewPostMeAccessRestrictionU2fIDEnableParams creates a new PostMeAccessRestrictionU2fIDEnableParams object
// with the default values initialized.
func NewPostMeAccessRestrictionU2fIDEnableParams() *PostMeAccessRestrictionU2fIDEnableParams {
	var ()
	return &PostMeAccessRestrictionU2fIDEnableParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostMeAccessRestrictionU2fIDEnableParamsWithTimeout creates a new PostMeAccessRestrictionU2fIDEnableParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostMeAccessRestrictionU2fIDEnableParamsWithTimeout(timeout time.Duration) *PostMeAccessRestrictionU2fIDEnableParams {
	var ()
	return &PostMeAccessRestrictionU2fIDEnableParams{

		timeout: timeout,
	}
}

// NewPostMeAccessRestrictionU2fIDEnableParamsWithContext creates a new PostMeAccessRestrictionU2fIDEnableParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostMeAccessRestrictionU2fIDEnableParamsWithContext(ctx context.Context) *PostMeAccessRestrictionU2fIDEnableParams {
	var ()
	return &PostMeAccessRestrictionU2fIDEnableParams{

		Context: ctx,
	}
}

// NewPostMeAccessRestrictionU2fIDEnableParamsWithHTTPClient creates a new PostMeAccessRestrictionU2fIDEnableParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostMeAccessRestrictionU2fIDEnableParamsWithHTTPClient(client *http.Client) *PostMeAccessRestrictionU2fIDEnableParams {
	var ()
	return &PostMeAccessRestrictionU2fIDEnableParams{
		HTTPClient: client,
	}
}

/*PostMeAccessRestrictionU2fIDEnableParams contains all the parameters to send to the API endpoint
for the post me access restriction u2f ID enable operation typically these are written to a http.Request
*/
type PostMeAccessRestrictionU2fIDEnableParams struct {

	/*ID*/
	ID int64
	/*MeAccessRestrictionU2fEnablePost*/
	MeAccessRestrictionU2fEnablePost *models.PostMeAccessRestrictionU2fIDEnableParamsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post me access restriction u2f ID enable params
func (o *PostMeAccessRestrictionU2fIDEnableParams) WithTimeout(timeout time.Duration) *PostMeAccessRestrictionU2fIDEnableParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post me access restriction u2f ID enable params
func (o *PostMeAccessRestrictionU2fIDEnableParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post me access restriction u2f ID enable params
func (o *PostMeAccessRestrictionU2fIDEnableParams) WithContext(ctx context.Context) *PostMeAccessRestrictionU2fIDEnableParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post me access restriction u2f ID enable params
func (o *PostMeAccessRestrictionU2fIDEnableParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post me access restriction u2f ID enable params
func (o *PostMeAccessRestrictionU2fIDEnableParams) WithHTTPClient(client *http.Client) *PostMeAccessRestrictionU2fIDEnableParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post me access restriction u2f ID enable params
func (o *PostMeAccessRestrictionU2fIDEnableParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the post me access restriction u2f ID enable params
func (o *PostMeAccessRestrictionU2fIDEnableParams) WithID(id int64) *PostMeAccessRestrictionU2fIDEnableParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post me access restriction u2f ID enable params
func (o *PostMeAccessRestrictionU2fIDEnableParams) SetID(id int64) {
	o.ID = id
}

// WithMeAccessRestrictionU2fEnablePost adds the meAccessRestrictionU2fEnablePost to the post me access restriction u2f ID enable params
func (o *PostMeAccessRestrictionU2fIDEnableParams) WithMeAccessRestrictionU2fEnablePost(meAccessRestrictionU2fEnablePost *models.PostMeAccessRestrictionU2fIDEnableParamsBody) *PostMeAccessRestrictionU2fIDEnableParams {
	o.SetMeAccessRestrictionU2fEnablePost(meAccessRestrictionU2fEnablePost)
	return o
}

// SetMeAccessRestrictionU2fEnablePost adds the meAccessRestrictionU2fEnablePost to the post me access restriction u2f ID enable params
func (o *PostMeAccessRestrictionU2fIDEnableParams) SetMeAccessRestrictionU2fEnablePost(meAccessRestrictionU2fEnablePost *models.PostMeAccessRestrictionU2fIDEnableParamsBody) {
	o.MeAccessRestrictionU2fEnablePost = meAccessRestrictionU2fEnablePost
}

// WriteToRequest writes these params to a swagger request
func (o *PostMeAccessRestrictionU2fIDEnableParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if o.MeAccessRestrictionU2fEnablePost != nil {
		if err := r.SetBodyParam(o.MeAccessRestrictionU2fEnablePost); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}