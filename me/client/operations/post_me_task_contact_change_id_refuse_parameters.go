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

// NewPostMeTaskContactChangeIDRefuseParams creates a new PostMeTaskContactChangeIDRefuseParams object
// with the default values initialized.
func NewPostMeTaskContactChangeIDRefuseParams() *PostMeTaskContactChangeIDRefuseParams {
	var ()
	return &PostMeTaskContactChangeIDRefuseParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostMeTaskContactChangeIDRefuseParamsWithTimeout creates a new PostMeTaskContactChangeIDRefuseParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostMeTaskContactChangeIDRefuseParamsWithTimeout(timeout time.Duration) *PostMeTaskContactChangeIDRefuseParams {
	var ()
	return &PostMeTaskContactChangeIDRefuseParams{

		timeout: timeout,
	}
}

// NewPostMeTaskContactChangeIDRefuseParamsWithContext creates a new PostMeTaskContactChangeIDRefuseParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostMeTaskContactChangeIDRefuseParamsWithContext(ctx context.Context) *PostMeTaskContactChangeIDRefuseParams {
	var ()
	return &PostMeTaskContactChangeIDRefuseParams{

		Context: ctx,
	}
}

// NewPostMeTaskContactChangeIDRefuseParamsWithHTTPClient creates a new PostMeTaskContactChangeIDRefuseParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostMeTaskContactChangeIDRefuseParamsWithHTTPClient(client *http.Client) *PostMeTaskContactChangeIDRefuseParams {
	var ()
	return &PostMeTaskContactChangeIDRefuseParams{
		HTTPClient: client,
	}
}

/*PostMeTaskContactChangeIDRefuseParams contains all the parameters to send to the API endpoint
for the post me task contact change ID refuse operation typically these are written to a http.Request
*/
type PostMeTaskContactChangeIDRefuseParams struct {

	/*ID*/
	ID int64
	/*MeTaskContactChangeRefusePost*/
	MeTaskContactChangeRefusePost *models.PostMeTaskContactChangeIDRefuseParamsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post me task contact change ID refuse params
func (o *PostMeTaskContactChangeIDRefuseParams) WithTimeout(timeout time.Duration) *PostMeTaskContactChangeIDRefuseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post me task contact change ID refuse params
func (o *PostMeTaskContactChangeIDRefuseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post me task contact change ID refuse params
func (o *PostMeTaskContactChangeIDRefuseParams) WithContext(ctx context.Context) *PostMeTaskContactChangeIDRefuseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post me task contact change ID refuse params
func (o *PostMeTaskContactChangeIDRefuseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post me task contact change ID refuse params
func (o *PostMeTaskContactChangeIDRefuseParams) WithHTTPClient(client *http.Client) *PostMeTaskContactChangeIDRefuseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post me task contact change ID refuse params
func (o *PostMeTaskContactChangeIDRefuseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the post me task contact change ID refuse params
func (o *PostMeTaskContactChangeIDRefuseParams) WithID(id int64) *PostMeTaskContactChangeIDRefuseParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post me task contact change ID refuse params
func (o *PostMeTaskContactChangeIDRefuseParams) SetID(id int64) {
	o.ID = id
}

// WithMeTaskContactChangeRefusePost adds the meTaskContactChangeRefusePost to the post me task contact change ID refuse params
func (o *PostMeTaskContactChangeIDRefuseParams) WithMeTaskContactChangeRefusePost(meTaskContactChangeRefusePost *models.PostMeTaskContactChangeIDRefuseParamsBody) *PostMeTaskContactChangeIDRefuseParams {
	o.SetMeTaskContactChangeRefusePost(meTaskContactChangeRefusePost)
	return o
}

// SetMeTaskContactChangeRefusePost adds the meTaskContactChangeRefusePost to the post me task contact change ID refuse params
func (o *PostMeTaskContactChangeIDRefuseParams) SetMeTaskContactChangeRefusePost(meTaskContactChangeRefusePost *models.PostMeTaskContactChangeIDRefuseParamsBody) {
	o.MeTaskContactChangeRefusePost = meTaskContactChangeRefusePost
}

// WriteToRequest writes these params to a swagger request
func (o *PostMeTaskContactChangeIDRefuseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if o.MeTaskContactChangeRefusePost != nil {
		if err := r.SetBodyParam(o.MeTaskContactChangeRefusePost); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
