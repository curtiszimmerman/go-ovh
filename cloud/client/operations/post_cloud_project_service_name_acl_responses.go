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

// PostCloudProjectServiceNameACLReader is a Reader for the PostCloudProjectServiceNameACL structure.
type PostCloudProjectServiceNameACLReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCloudProjectServiceNameACLReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostCloudProjectServiceNameACLOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostCloudProjectServiceNameACLDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostCloudProjectServiceNameACLOK creates a PostCloudProjectServiceNameACLOK with default headers values
func NewPostCloudProjectServiceNameACLOK() *PostCloudProjectServiceNameACLOK {
	return &PostCloudProjectServiceNameACLOK{}
}

/*PostCloudProjectServiceNameACLOK handles this case with default header values.

description of 'cloud.Acl' response
*/
type PostCloudProjectServiceNameACLOK struct {
	Payload *models.CloudACL
}

func (o *PostCloudProjectServiceNameACLOK) Error() string {
	return fmt.Sprintf("[POST /cloud/project/{serviceName}/acl][%d] postCloudProjectServiceNameAclOK  %+v", 200, o.Payload)
}

func (o *PostCloudProjectServiceNameACLOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudACL)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCloudProjectServiceNameACLDefault creates a PostCloudProjectServiceNameACLDefault with default headers values
func NewPostCloudProjectServiceNameACLDefault(code int) *PostCloudProjectServiceNameACLDefault {
	return &PostCloudProjectServiceNameACLDefault{
		_statusCode: code,
	}
}

/*PostCloudProjectServiceNameACLDefault handles this case with default header values.

Unexpected error
*/
type PostCloudProjectServiceNameACLDefault struct {
	_statusCode int

	Payload *models.PostCloudProjectServiceNameACLDefaultBody
}

// Code gets the status code for the post cloud project service name ACL default response
func (o *PostCloudProjectServiceNameACLDefault) Code() int {
	return o._statusCode
}

func (o *PostCloudProjectServiceNameACLDefault) Error() string {
	return fmt.Sprintf("[POST /cloud/project/{serviceName}/acl][%d] PostCloudProjectServiceNameACL default  %+v", o._statusCode, o.Payload)
}

func (o *PostCloudProjectServiceNameACLDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostCloudProjectServiceNameACLDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}