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

// NewPostMeAccessRestrictionU2fIDDisableParams creates a new PostMeAccessRestrictionU2fIDDisableParams object
// with the default values initialized.
func NewPostMeAccessRestrictionU2fIDDisableParams() *PostMeAccessRestrictionU2fIDDisableParams {
	var ()
	return &PostMeAccessRestrictionU2fIDDisableParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostMeAccessRestrictionU2fIDDisableParamsWithTimeout creates a new PostMeAccessRestrictionU2fIDDisableParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostMeAccessRestrictionU2fIDDisableParamsWithTimeout(timeout time.Duration) *PostMeAccessRestrictionU2fIDDisableParams {
	var ()
	return &PostMeAccessRestrictionU2fIDDisableParams{

		timeout: timeout,
	}
}

// NewPostMeAccessRestrictionU2fIDDisableParamsWithContext creates a new PostMeAccessRestrictionU2fIDDisableParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostMeAccessRestrictionU2fIDDisableParamsWithContext(ctx context.Context) *PostMeAccessRestrictionU2fIDDisableParams {
	var ()
	return &PostMeAccessRestrictionU2fIDDisableParams{

		Context: ctx,
	}
}

// NewPostMeAccessRestrictionU2fIDDisableParamsWithHTTPClient creates a new PostMeAccessRestrictionU2fIDDisableParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostMeAccessRestrictionU2fIDDisableParamsWithHTTPClient(client *http.Client) *PostMeAccessRestrictionU2fIDDisableParams {
	var ()
	return &PostMeAccessRestrictionU2fIDDisableParams{
		HTTPClient: client,
	}
}

/*PostMeAccessRestrictionU2fIDDisableParams contains all the parameters to send to the API endpoint
for the post me access restriction u2f ID disable operation typically these are written to a http.Request
*/
type PostMeAccessRestrictionU2fIDDisableParams struct {

	/*ID*/
	ID int64
	/*MeAccessRestrictionU2fDisablePost*/
	MeAccessRestrictionU2fDisablePost *models.PostMeAccessRestrictionU2fIDDisableParamsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post me access restriction u2f ID disable params
func (o *PostMeAccessRestrictionU2fIDDisableParams) WithTimeout(timeout time.Duration) *PostMeAccessRestrictionU2fIDDisableParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post me access restriction u2f ID disable params
func (o *PostMeAccessRestrictionU2fIDDisableParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post me access restriction u2f ID disable params
func (o *PostMeAccessRestrictionU2fIDDisableParams) WithContext(ctx context.Context) *PostMeAccessRestrictionU2fIDDisableParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post me access restriction u2f ID disable params
func (o *PostMeAccessRestrictionU2fIDDisableParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post me access restriction u2f ID disable params
func (o *PostMeAccessRestrictionU2fIDDisableParams) WithHTTPClient(client *http.Client) *PostMeAccessRestrictionU2fIDDisableParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post me access restriction u2f ID disable params
func (o *PostMeAccessRestrictionU2fIDDisableParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the post me access restriction u2f ID disable params
func (o *PostMeAccessRestrictionU2fIDDisableParams) WithID(id int64) *PostMeAccessRestrictionU2fIDDisableParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post me access restriction u2f ID disable params
func (o *PostMeAccessRestrictionU2fIDDisableParams) SetID(id int64) {
	o.ID = id
}

// WithMeAccessRestrictionU2fDisablePost adds the meAccessRestrictionU2fDisablePost to the post me access restriction u2f ID disable params
func (o *PostMeAccessRestrictionU2fIDDisableParams) WithMeAccessRestrictionU2fDisablePost(meAccessRestrictionU2fDisablePost *models.PostMeAccessRestrictionU2fIDDisableParamsBody) *PostMeAccessRestrictionU2fIDDisableParams {
	o.SetMeAccessRestrictionU2fDisablePost(meAccessRestrictionU2fDisablePost)
	return o
}

// SetMeAccessRestrictionU2fDisablePost adds the meAccessRestrictionU2fDisablePost to the post me access restriction u2f ID disable params
func (o *PostMeAccessRestrictionU2fIDDisableParams) SetMeAccessRestrictionU2fDisablePost(meAccessRestrictionU2fDisablePost *models.PostMeAccessRestrictionU2fIDDisableParamsBody) {
	o.MeAccessRestrictionU2fDisablePost = meAccessRestrictionU2fDisablePost
}

// WriteToRequest writes these params to a swagger request
func (o *PostMeAccessRestrictionU2fIDDisableParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if o.MeAccessRestrictionU2fDisablePost != nil {
		if err := r.SetBodyParam(o.MeAccessRestrictionU2fDisablePost); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
