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

// DeleteMeAccessRestrictionTotpIDReader is a Reader for the DeleteMeAccessRestrictionTotpID structure.
type DeleteMeAccessRestrictionTotpIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteMeAccessRestrictionTotpIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteMeAccessRestrictionTotpIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteMeAccessRestrictionTotpIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteMeAccessRestrictionTotpIDOK creates a DeleteMeAccessRestrictionTotpIDOK with default headers values
func NewDeleteMeAccessRestrictionTotpIDOK() *DeleteMeAccessRestrictionTotpIDOK {
	return &DeleteMeAccessRestrictionTotpIDOK{}
}

/*DeleteMeAccessRestrictionTotpIDOK handles this case with default header values.

return 'void'
*/
type DeleteMeAccessRestrictionTotpIDOK struct {
}

func (o *DeleteMeAccessRestrictionTotpIDOK) Error() string {
	return fmt.Sprintf("[DELETE /me/accessRestriction/totp/{id}][%d] deleteMeAccessRestrictionTotpIdOK ", 200)
}

func (o *DeleteMeAccessRestrictionTotpIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteMeAccessRestrictionTotpIDDefault creates a DeleteMeAccessRestrictionTotpIDDefault with default headers values
func NewDeleteMeAccessRestrictionTotpIDDefault(code int) *DeleteMeAccessRestrictionTotpIDDefault {
	return &DeleteMeAccessRestrictionTotpIDDefault{
		_statusCode: code,
	}
}

/*DeleteMeAccessRestrictionTotpIDDefault handles this case with default header values.

Unexpected error
*/
type DeleteMeAccessRestrictionTotpIDDefault struct {
	_statusCode int

	Payload *models.DeleteMeAccessRestrictionTotpIDDefaultBody
}

// Code gets the status code for the delete me access restriction totp ID default response
func (o *DeleteMeAccessRestrictionTotpIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteMeAccessRestrictionTotpIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /me/accessRestriction/totp/{id}][%d] DeleteMeAccessRestrictionTotpID default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteMeAccessRestrictionTotpIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteMeAccessRestrictionTotpIDDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}