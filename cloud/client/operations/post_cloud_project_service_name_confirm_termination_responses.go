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

// PostCloudProjectServiceNameConfirmTerminationReader is a Reader for the PostCloudProjectServiceNameConfirmTermination structure.
type PostCloudProjectServiceNameConfirmTerminationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCloudProjectServiceNameConfirmTerminationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostCloudProjectServiceNameConfirmTerminationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostCloudProjectServiceNameConfirmTerminationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostCloudProjectServiceNameConfirmTerminationOK creates a PostCloudProjectServiceNameConfirmTerminationOK with default headers values
func NewPostCloudProjectServiceNameConfirmTerminationOK() *PostCloudProjectServiceNameConfirmTerminationOK {
	return &PostCloudProjectServiceNameConfirmTerminationOK{}
}

/*PostCloudProjectServiceNameConfirmTerminationOK handles this case with default header values.

return value
*/
type PostCloudProjectServiceNameConfirmTerminationOK struct {
	Payload string
}

func (o *PostCloudProjectServiceNameConfirmTerminationOK) Error() string {
	return fmt.Sprintf("[POST /cloud/project/{serviceName}/confirmTermination][%d] postCloudProjectServiceNameConfirmTerminationOK  %+v", 200, o.Payload)
}

func (o *PostCloudProjectServiceNameConfirmTerminationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCloudProjectServiceNameConfirmTerminationDefault creates a PostCloudProjectServiceNameConfirmTerminationDefault with default headers values
func NewPostCloudProjectServiceNameConfirmTerminationDefault(code int) *PostCloudProjectServiceNameConfirmTerminationDefault {
	return &PostCloudProjectServiceNameConfirmTerminationDefault{
		_statusCode: code,
	}
}

/*PostCloudProjectServiceNameConfirmTerminationDefault handles this case with default header values.

Unexpected error
*/
type PostCloudProjectServiceNameConfirmTerminationDefault struct {
	_statusCode int

	Payload *models.PostCloudProjectServiceNameConfirmTerminationDefaultBody
}

// Code gets the status code for the post cloud project service name confirm termination default response
func (o *PostCloudProjectServiceNameConfirmTerminationDefault) Code() int {
	return o._statusCode
}

func (o *PostCloudProjectServiceNameConfirmTerminationDefault) Error() string {
	return fmt.Sprintf("[POST /cloud/project/{serviceName}/confirmTermination][%d] PostCloudProjectServiceNameConfirmTermination default  %+v", o._statusCode, o.Payload)
}

func (o *PostCloudProjectServiceNameConfirmTerminationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostCloudProjectServiceNameConfirmTerminationDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}