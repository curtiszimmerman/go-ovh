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

// NewPostMeAutorenewParams creates a new PostMeAutorenewParams object
// with the default values initialized.
func NewPostMeAutorenewParams() *PostMeAutorenewParams {
	var ()
	return &PostMeAutorenewParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostMeAutorenewParamsWithTimeout creates a new PostMeAutorenewParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostMeAutorenewParamsWithTimeout(timeout time.Duration) *PostMeAutorenewParams {
	var ()
	return &PostMeAutorenewParams{

		timeout: timeout,
	}
}

// NewPostMeAutorenewParamsWithContext creates a new PostMeAutorenewParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostMeAutorenewParamsWithContext(ctx context.Context) *PostMeAutorenewParams {
	var ()
	return &PostMeAutorenewParams{

		Context: ctx,
	}
}

// NewPostMeAutorenewParamsWithHTTPClient creates a new PostMeAutorenewParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostMeAutorenewParamsWithHTTPClient(client *http.Client) *PostMeAutorenewParams {
	var ()
	return &PostMeAutorenewParams{
		HTTPClient: client,
	}
}

/*PostMeAutorenewParams contains all the parameters to send to the API endpoint
for the post me autorenew operation typically these are written to a http.Request
*/
type PostMeAutorenewParams struct {

	/*MeAutorenewPost*/
	MeAutorenewPost *models.PostMeAutorenewParamsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post me autorenew params
func (o *PostMeAutorenewParams) WithTimeout(timeout time.Duration) *PostMeAutorenewParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post me autorenew params
func (o *PostMeAutorenewParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post me autorenew params
func (o *PostMeAutorenewParams) WithContext(ctx context.Context) *PostMeAutorenewParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post me autorenew params
func (o *PostMeAutorenewParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post me autorenew params
func (o *PostMeAutorenewParams) WithHTTPClient(client *http.Client) *PostMeAutorenewParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post me autorenew params
func (o *PostMeAutorenewParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMeAutorenewPost adds the meAutorenewPost to the post me autorenew params
func (o *PostMeAutorenewParams) WithMeAutorenewPost(meAutorenewPost *models.PostMeAutorenewParamsBody) *PostMeAutorenewParams {
	o.SetMeAutorenewPost(meAutorenewPost)
	return o
}

// SetMeAutorenewPost adds the meAutorenewPost to the post me autorenew params
func (o *PostMeAutorenewParams) SetMeAutorenewPost(meAutorenewPost *models.PostMeAutorenewParamsBody) {
	o.MeAutorenewPost = meAutorenewPost
}

// WriteToRequest writes these params to a swagger request
func (o *PostMeAutorenewParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.MeAutorenewPost != nil {
		if err := r.SetBodyParam(o.MeAutorenewPost); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
