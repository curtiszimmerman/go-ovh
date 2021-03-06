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

// GetCloudServiceNamePcaPcaServiceNameTasksTaskIDReader is a Reader for the GetCloudServiceNamePcaPcaServiceNameTasksTaskID structure.
type GetCloudServiceNamePcaPcaServiceNameTasksTaskIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCloudServiceNamePcaPcaServiceNameTasksTaskIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCloudServiceNamePcaPcaServiceNameTasksTaskIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetCloudServiceNamePcaPcaServiceNameTasksTaskIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCloudServiceNamePcaPcaServiceNameTasksTaskIDOK creates a GetCloudServiceNamePcaPcaServiceNameTasksTaskIDOK with default headers values
func NewGetCloudServiceNamePcaPcaServiceNameTasksTaskIDOK() *GetCloudServiceNamePcaPcaServiceNameTasksTaskIDOK {
	return &GetCloudServiceNamePcaPcaServiceNameTasksTaskIDOK{}
}

/*GetCloudServiceNamePcaPcaServiceNameTasksTaskIDOK handles this case with default header values.

description of 'pca.Task' response
*/
type GetCloudServiceNamePcaPcaServiceNameTasksTaskIDOK struct {
	Payload *models.PcaTask
}

func (o *GetCloudServiceNamePcaPcaServiceNameTasksTaskIDOK) Error() string {
	return fmt.Sprintf("[GET /cloud/{serviceName}/pca/{pcaServiceName}/tasks/{taskId}][%d] getCloudServiceNamePcaPcaServiceNameTasksTaskIdOK  %+v", 200, o.Payload)
}

func (o *GetCloudServiceNamePcaPcaServiceNameTasksTaskIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PcaTask)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCloudServiceNamePcaPcaServiceNameTasksTaskIDDefault creates a GetCloudServiceNamePcaPcaServiceNameTasksTaskIDDefault with default headers values
func NewGetCloudServiceNamePcaPcaServiceNameTasksTaskIDDefault(code int) *GetCloudServiceNamePcaPcaServiceNameTasksTaskIDDefault {
	return &GetCloudServiceNamePcaPcaServiceNameTasksTaskIDDefault{
		_statusCode: code,
	}
}

/*GetCloudServiceNamePcaPcaServiceNameTasksTaskIDDefault handles this case with default header values.

Unexpected error
*/
type GetCloudServiceNamePcaPcaServiceNameTasksTaskIDDefault struct {
	_statusCode int

	Payload *models.GetCloudServiceNamePcaPcaServiceNameTasksTaskIDDefaultBody
}

// Code gets the status code for the get cloud service name pca pca service name tasks task ID default response
func (o *GetCloudServiceNamePcaPcaServiceNameTasksTaskIDDefault) Code() int {
	return o._statusCode
}

func (o *GetCloudServiceNamePcaPcaServiceNameTasksTaskIDDefault) Error() string {
	return fmt.Sprintf("[GET /cloud/{serviceName}/pca/{pcaServiceName}/tasks/{taskId}][%d] GetCloudServiceNamePcaPcaServiceNameTasksTaskID default  %+v", o._statusCode, o.Payload)
}

func (o *GetCloudServiceNamePcaPcaServiceNameTasksTaskIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCloudServiceNamePcaPcaServiceNameTasksTaskIDDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
