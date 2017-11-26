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

// PostCloudServiceNamePcaPcaServiceNameTasksReader is a Reader for the PostCloudServiceNamePcaPcaServiceNameTasks structure.
type PostCloudServiceNamePcaPcaServiceNameTasksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCloudServiceNamePcaPcaServiceNameTasksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostCloudServiceNamePcaPcaServiceNameTasksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostCloudServiceNamePcaPcaServiceNameTasksDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostCloudServiceNamePcaPcaServiceNameTasksOK creates a PostCloudServiceNamePcaPcaServiceNameTasksOK with default headers values
func NewPostCloudServiceNamePcaPcaServiceNameTasksOK() *PostCloudServiceNamePcaPcaServiceNameTasksOK {
	return &PostCloudServiceNamePcaPcaServiceNameTasksOK{}
}

/*PostCloudServiceNamePcaPcaServiceNameTasksOK handles this case with default header values.

description of 'pca.Task' response
*/
type PostCloudServiceNamePcaPcaServiceNameTasksOK struct {
	Payload *models.PcaTask
}

func (o *PostCloudServiceNamePcaPcaServiceNameTasksOK) Error() string {
	return fmt.Sprintf("[POST /cloud/{serviceName}/pca/{pcaServiceName}/tasks][%d] postCloudServiceNamePcaPcaServiceNameTasksOK  %+v", 200, o.Payload)
}

func (o *PostCloudServiceNamePcaPcaServiceNameTasksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PcaTask)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCloudServiceNamePcaPcaServiceNameTasksDefault creates a PostCloudServiceNamePcaPcaServiceNameTasksDefault with default headers values
func NewPostCloudServiceNamePcaPcaServiceNameTasksDefault(code int) *PostCloudServiceNamePcaPcaServiceNameTasksDefault {
	return &PostCloudServiceNamePcaPcaServiceNameTasksDefault{
		_statusCode: code,
	}
}

/*PostCloudServiceNamePcaPcaServiceNameTasksDefault handles this case with default header values.

Unexpected error
*/
type PostCloudServiceNamePcaPcaServiceNameTasksDefault struct {
	_statusCode int

	Payload *models.PostCloudServiceNamePcaPcaServiceNameTasksDefaultBody
}

// Code gets the status code for the post cloud service name pca pca service name tasks default response
func (o *PostCloudServiceNamePcaPcaServiceNameTasksDefault) Code() int {
	return o._statusCode
}

func (o *PostCloudServiceNamePcaPcaServiceNameTasksDefault) Error() string {
	return fmt.Sprintf("[POST /cloud/{serviceName}/pca/{pcaServiceName}/tasks][%d] PostCloudServiceNamePcaPcaServiceNameTasks default  %+v", o._statusCode, o.Payload)
}

func (o *PostCloudServiceNamePcaPcaServiceNameTasksDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostCloudServiceNamePcaPcaServiceNameTasksDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}