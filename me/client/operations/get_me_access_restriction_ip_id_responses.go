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

// GetMeAccessRestrictionIPIDReader is a Reader for the GetMeAccessRestrictionIPID structure.
type GetMeAccessRestrictionIPIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMeAccessRestrictionIPIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMeAccessRestrictionIPIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetMeAccessRestrictionIPIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMeAccessRestrictionIPIDOK creates a GetMeAccessRestrictionIPIDOK with default headers values
func NewGetMeAccessRestrictionIPIDOK() *GetMeAccessRestrictionIPIDOK {
	return &GetMeAccessRestrictionIPIDOK{}
}

/*GetMeAccessRestrictionIPIDOK handles this case with default header values.

description of 'nichandle.IpRestriction' response
*/
type GetMeAccessRestrictionIPIDOK struct {
	Payload *models.NichandleIPRestriction
}

func (o *GetMeAccessRestrictionIPIDOK) Error() string {
	return fmt.Sprintf("[GET /me/accessRestriction/ip/{id}][%d] getMeAccessRestrictionIpIdOK  %+v", 200, o.Payload)
}

func (o *GetMeAccessRestrictionIPIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NichandleIPRestriction)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMeAccessRestrictionIPIDDefault creates a GetMeAccessRestrictionIPIDDefault with default headers values
func NewGetMeAccessRestrictionIPIDDefault(code int) *GetMeAccessRestrictionIPIDDefault {
	return &GetMeAccessRestrictionIPIDDefault{
		_statusCode: code,
	}
}

/*GetMeAccessRestrictionIPIDDefault handles this case with default header values.

Unexpected error
*/
type GetMeAccessRestrictionIPIDDefault struct {
	_statusCode int

	Payload *models.GetMeAccessRestrictionIPIDDefaultBody
}

// Code gets the status code for the get me access restriction IP ID default response
func (o *GetMeAccessRestrictionIPIDDefault) Code() int {
	return o._statusCode
}

func (o *GetMeAccessRestrictionIPIDDefault) Error() string {
	return fmt.Sprintf("[GET /me/accessRestriction/ip/{id}][%d] GetMeAccessRestrictionIPID default  %+v", o._statusCode, o.Payload)
}

func (o *GetMeAccessRestrictionIPIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetMeAccessRestrictionIPIDDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
