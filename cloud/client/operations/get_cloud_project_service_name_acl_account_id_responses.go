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

	"github.com/appscode/go-ovh/cloud/models"
)

// GetCloudProjectServiceNameACLAccountIDReader is a Reader for the GetCloudProjectServiceNameACLAccountID structure.
type GetCloudProjectServiceNameACLAccountIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCloudProjectServiceNameACLAccountIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCloudProjectServiceNameACLAccountIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetCloudProjectServiceNameACLAccountIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCloudProjectServiceNameACLAccountIDOK creates a GetCloudProjectServiceNameACLAccountIDOK with default headers values
func NewGetCloudProjectServiceNameACLAccountIDOK() *GetCloudProjectServiceNameACLAccountIDOK {
	return &GetCloudProjectServiceNameACLAccountIDOK{}
}

/*GetCloudProjectServiceNameACLAccountIDOK handles this case with default header values.

description of 'cloud.Acl' response
*/
type GetCloudProjectServiceNameACLAccountIDOK struct {
	Payload *models.CloudACL
}

func (o *GetCloudProjectServiceNameACLAccountIDOK) Error() string {
	return fmt.Sprintf("[GET /cloud/project/{serviceName}/acl/{accountId}][%d] getCloudProjectServiceNameAclAccountIdOK  %+v", 200, o.Payload)
}

func (o *GetCloudProjectServiceNameACLAccountIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudACL)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCloudProjectServiceNameACLAccountIDDefault creates a GetCloudProjectServiceNameACLAccountIDDefault with default headers values
func NewGetCloudProjectServiceNameACLAccountIDDefault(code int) *GetCloudProjectServiceNameACLAccountIDDefault {
	return &GetCloudProjectServiceNameACLAccountIDDefault{
		_statusCode: code,
	}
}

/*GetCloudProjectServiceNameACLAccountIDDefault handles this case with default header values.

Unexpected error
*/
type GetCloudProjectServiceNameACLAccountIDDefault struct {
	_statusCode int

	Payload *models.GetCloudProjectServiceNameACLAccountIDDefaultBody
}

// Code gets the status code for the get cloud project service name ACL account ID default response
func (o *GetCloudProjectServiceNameACLAccountIDDefault) Code() int {
	return o._statusCode
}

func (o *GetCloudProjectServiceNameACLAccountIDDefault) Error() string {
	return fmt.Sprintf("[GET /cloud/project/{serviceName}/acl/{accountId}][%d] GetCloudProjectServiceNameACLAccountID default  %+v", o._statusCode, o.Payload)
}

func (o *GetCloudProjectServiceNameACLAccountIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCloudProjectServiceNameACLAccountIDDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
