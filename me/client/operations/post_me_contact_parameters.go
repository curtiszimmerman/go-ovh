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

	"github.com/appscode/go-ovh/me/models"
)

// NewPostMeContactParams creates a new PostMeContactParams object
// with the default values initialized.
func NewPostMeContactParams() *PostMeContactParams {
	var ()
	return &PostMeContactParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostMeContactParamsWithTimeout creates a new PostMeContactParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostMeContactParamsWithTimeout(timeout time.Duration) *PostMeContactParams {
	var ()
	return &PostMeContactParams{

		timeout: timeout,
	}
}

// NewPostMeContactParamsWithContext creates a new PostMeContactParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostMeContactParamsWithContext(ctx context.Context) *PostMeContactParams {
	var ()
	return &PostMeContactParams{

		Context: ctx,
	}
}

// NewPostMeContactParamsWithHTTPClient creates a new PostMeContactParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostMeContactParamsWithHTTPClient(client *http.Client) *PostMeContactParams {
	var ()
	return &PostMeContactParams{
		HTTPClient: client,
	}
}

/*PostMeContactParams contains all the parameters to send to the API endpoint
for the post me contact operation typically these are written to a http.Request
*/
type PostMeContactParams struct {

	/*MeContactPost*/
	MeContactPost *models.PostMeContactParamsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post me contact params
func (o *PostMeContactParams) WithTimeout(timeout time.Duration) *PostMeContactParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post me contact params
func (o *PostMeContactParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post me contact params
func (o *PostMeContactParams) WithContext(ctx context.Context) *PostMeContactParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post me contact params
func (o *PostMeContactParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post me contact params
func (o *PostMeContactParams) WithHTTPClient(client *http.Client) *PostMeContactParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post me contact params
func (o *PostMeContactParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMeContactPost adds the meContactPost to the post me contact params
func (o *PostMeContactParams) WithMeContactPost(meContactPost *models.PostMeContactParamsBody) *PostMeContactParams {
	o.SetMeContactPost(meContactPost)
	return o
}

// SetMeContactPost adds the meContactPost to the post me contact params
func (o *PostMeContactParams) SetMeContactPost(meContactPost *models.PostMeContactParamsBody) {
	o.MeContactPost = meContactPost
}

// WriteToRequest writes these params to a swagger request
func (o *PostMeContactParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.MeContactPost != nil {
		if err := r.SetBodyParam(o.MeContactPost); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
