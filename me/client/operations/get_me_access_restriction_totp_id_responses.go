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

// GetMeAccessRestrictionTotpIDReader is a Reader for the GetMeAccessRestrictionTotpID structure.
type GetMeAccessRestrictionTotpIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMeAccessRestrictionTotpIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMeAccessRestrictionTotpIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetMeAccessRestrictionTotpIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMeAccessRestrictionTotpIDOK creates a GetMeAccessRestrictionTotpIDOK with default headers values
func NewGetMeAccessRestrictionTotpIDOK() *GetMeAccessRestrictionTotpIDOK {
	return &GetMeAccessRestrictionTotpIDOK{}
}

/*GetMeAccessRestrictionTotpIDOK handles this case with default header values.

description of 'nichandle.AccessRestriction.TOTPAccount' response
*/
type GetMeAccessRestrictionTotpIDOK struct {
	Payload *models.NichandleAccessRestrictionTOTPAccount
}

func (o *GetMeAccessRestrictionTotpIDOK) Error() string {
	return fmt.Sprintf("[GET /me/accessRestriction/totp/{id}][%d] getMeAccessRestrictionTotpIdOK  %+v", 200, o.Payload)
}

func (o *GetMeAccessRestrictionTotpIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NichandleAccessRestrictionTOTPAccount)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMeAccessRestrictionTotpIDDefault creates a GetMeAccessRestrictionTotpIDDefault with default headers values
func NewGetMeAccessRestrictionTotpIDDefault(code int) *GetMeAccessRestrictionTotpIDDefault {
	return &GetMeAccessRestrictionTotpIDDefault{
		_statusCode: code,
	}
}

/*GetMeAccessRestrictionTotpIDDefault handles this case with default header values.

Unexpected error
*/
type GetMeAccessRestrictionTotpIDDefault struct {
	_statusCode int

	Payload *models.GetMeAccessRestrictionTotpIDDefaultBody
}

// Code gets the status code for the get me access restriction totp ID default response
func (o *GetMeAccessRestrictionTotpIDDefault) Code() int {
	return o._statusCode
}

func (o *GetMeAccessRestrictionTotpIDDefault) Error() string {
	return fmt.Sprintf("[GET /me/accessRestriction/totp/{id}][%d] GetMeAccessRestrictionTotpID default  %+v", o._statusCode, o.Payload)
}

func (o *GetMeAccessRestrictionTotpIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetMeAccessRestrictionTotpIDDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
