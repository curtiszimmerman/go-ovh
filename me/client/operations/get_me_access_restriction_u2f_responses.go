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

// GetMeAccessRestrictionU2fReader is a Reader for the GetMeAccessRestrictionU2f structure.
type GetMeAccessRestrictionU2fReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMeAccessRestrictionU2fReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMeAccessRestrictionU2fOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetMeAccessRestrictionU2fDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMeAccessRestrictionU2fOK creates a GetMeAccessRestrictionU2fOK with default headers values
func NewGetMeAccessRestrictionU2fOK() *GetMeAccessRestrictionU2fOK {
	return &GetMeAccessRestrictionU2fOK{}
}

/*GetMeAccessRestrictionU2fOK handles this case with default header values.

return value
*/
type GetMeAccessRestrictionU2fOK struct {
	Payload []int64
}

func (o *GetMeAccessRestrictionU2fOK) Error() string {
	return fmt.Sprintf("[GET /me/accessRestriction/u2f][%d] getMeAccessRestrictionU2fOK  %+v", 200, o.Payload)
}

func (o *GetMeAccessRestrictionU2fOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMeAccessRestrictionU2fDefault creates a GetMeAccessRestrictionU2fDefault with default headers values
func NewGetMeAccessRestrictionU2fDefault(code int) *GetMeAccessRestrictionU2fDefault {
	return &GetMeAccessRestrictionU2fDefault{
		_statusCode: code,
	}
}

/*GetMeAccessRestrictionU2fDefault handles this case with default header values.

Unexpected error
*/
type GetMeAccessRestrictionU2fDefault struct {
	_statusCode int

	Payload *models.GetMeAccessRestrictionU2fDefaultBody
}

// Code gets the status code for the get me access restriction u2f default response
func (o *GetMeAccessRestrictionU2fDefault) Code() int {
	return o._statusCode
}

func (o *GetMeAccessRestrictionU2fDefault) Error() string {
	return fmt.Sprintf("[GET /me/accessRestriction/u2f][%d] GetMeAccessRestrictionU2f default  %+v", o._statusCode, o.Payload)
}

func (o *GetMeAccessRestrictionU2fDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetMeAccessRestrictionU2fDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
