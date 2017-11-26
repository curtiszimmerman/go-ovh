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

// NewPostMeMailingListSubscribeParams creates a new PostMeMailingListSubscribeParams object
// with the default values initialized.
func NewPostMeMailingListSubscribeParams() *PostMeMailingListSubscribeParams {
	var ()
	return &PostMeMailingListSubscribeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostMeMailingListSubscribeParamsWithTimeout creates a new PostMeMailingListSubscribeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostMeMailingListSubscribeParamsWithTimeout(timeout time.Duration) *PostMeMailingListSubscribeParams {
	var ()
	return &PostMeMailingListSubscribeParams{

		timeout: timeout,
	}
}

// NewPostMeMailingListSubscribeParamsWithContext creates a new PostMeMailingListSubscribeParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostMeMailingListSubscribeParamsWithContext(ctx context.Context) *PostMeMailingListSubscribeParams {
	var ()
	return &PostMeMailingListSubscribeParams{

		Context: ctx,
	}
}

// NewPostMeMailingListSubscribeParamsWithHTTPClient creates a new PostMeMailingListSubscribeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostMeMailingListSubscribeParamsWithHTTPClient(client *http.Client) *PostMeMailingListSubscribeParams {
	var ()
	return &PostMeMailingListSubscribeParams{
		HTTPClient: client,
	}
}

/*PostMeMailingListSubscribeParams contains all the parameters to send to the API endpoint
for the post me mailing list subscribe operation typically these are written to a http.Request
*/
type PostMeMailingListSubscribeParams struct {

	/*MeMailingListSubscribePost*/
	MeMailingListSubscribePost *models.PostMeMailingListSubscribeParamsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post me mailing list subscribe params
func (o *PostMeMailingListSubscribeParams) WithTimeout(timeout time.Duration) *PostMeMailingListSubscribeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post me mailing list subscribe params
func (o *PostMeMailingListSubscribeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post me mailing list subscribe params
func (o *PostMeMailingListSubscribeParams) WithContext(ctx context.Context) *PostMeMailingListSubscribeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post me mailing list subscribe params
func (o *PostMeMailingListSubscribeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post me mailing list subscribe params
func (o *PostMeMailingListSubscribeParams) WithHTTPClient(client *http.Client) *PostMeMailingListSubscribeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post me mailing list subscribe params
func (o *PostMeMailingListSubscribeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMeMailingListSubscribePost adds the meMailingListSubscribePost to the post me mailing list subscribe params
func (o *PostMeMailingListSubscribeParams) WithMeMailingListSubscribePost(meMailingListSubscribePost *models.PostMeMailingListSubscribeParamsBody) *PostMeMailingListSubscribeParams {
	o.SetMeMailingListSubscribePost(meMailingListSubscribePost)
	return o
}

// SetMeMailingListSubscribePost adds the meMailingListSubscribePost to the post me mailing list subscribe params
func (o *PostMeMailingListSubscribeParams) SetMeMailingListSubscribePost(meMailingListSubscribePost *models.PostMeMailingListSubscribeParamsBody) {
	o.MeMailingListSubscribePost = meMailingListSubscribePost
}

// WriteToRequest writes these params to a swagger request
func (o *PostMeMailingListSubscribeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.MeMailingListSubscribePost != nil {
		if err := r.SetBodyParam(o.MeMailingListSubscribePost); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
