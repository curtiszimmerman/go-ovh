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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/appscode/go-ovh/me/models"
)

// PostMeAccessRestrictionSmsIDDisableReader is a Reader for the PostMeAccessRestrictionSmsIDDisable structure.
type PostMeAccessRestrictionSmsIDDisableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostMeAccessRestrictionSmsIDDisableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostMeAccessRestrictionSmsIDDisableOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostMeAccessRestrictionSmsIDDisableDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostMeAccessRestrictionSmsIDDisableOK creates a PostMeAccessRestrictionSmsIDDisableOK with default headers values
func NewPostMeAccessRestrictionSmsIDDisableOK() *PostMeAccessRestrictionSmsIDDisableOK {
	return &PostMeAccessRestrictionSmsIDDisableOK{}
}

/*PostMeAccessRestrictionSmsIDDisableOK handles this case with default header values.

return 'void'
*/
type PostMeAccessRestrictionSmsIDDisableOK struct {
}

func (o *PostMeAccessRestrictionSmsIDDisableOK) Error() string {
	return fmt.Sprintf("[POST /me/accessRestriction/sms/{id}/disable][%d] postMeAccessRestrictionSmsIdDisableOK ", 200)
}

func (o *PostMeAccessRestrictionSmsIDDisableOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostMeAccessRestrictionSmsIDDisableDefault creates a PostMeAccessRestrictionSmsIDDisableDefault with default headers values
func NewPostMeAccessRestrictionSmsIDDisableDefault(code int) *PostMeAccessRestrictionSmsIDDisableDefault {
	return &PostMeAccessRestrictionSmsIDDisableDefault{
		_statusCode: code,
	}
}

/*PostMeAccessRestrictionSmsIDDisableDefault handles this case with default header values.

Unexpected error
*/
type PostMeAccessRestrictionSmsIDDisableDefault struct {
	_statusCode int

	Payload *models.PostMeAccessRestrictionSmsIDDisableDefaultBody
}

// Code gets the status code for the post me access restriction sms ID disable default response
func (o *PostMeAccessRestrictionSmsIDDisableDefault) Code() int {
	return o._statusCode
}

func (o *PostMeAccessRestrictionSmsIDDisableDefault) Error() string {
	return fmt.Sprintf("[POST /me/accessRestriction/sms/{id}/disable][%d] PostMeAccessRestrictionSmsIDDisable default  %+v", o._statusCode, o.Payload)
}

func (o *PostMeAccessRestrictionSmsIDDisableDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostMeAccessRestrictionSmsIDDisableDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}