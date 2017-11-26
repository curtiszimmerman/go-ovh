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

// GetCloudCreateProjectInfoReader is a Reader for the GetCloudCreateProjectInfo structure.
type GetCloudCreateProjectInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCloudCreateProjectInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCloudCreateProjectInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetCloudCreateProjectInfoDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCloudCreateProjectInfoOK creates a GetCloudCreateProjectInfoOK with default headers values
func NewGetCloudCreateProjectInfoOK() *GetCloudCreateProjectInfoOK {
	return &GetCloudCreateProjectInfoOK{}
}

/*GetCloudCreateProjectInfoOK handles this case with default header values.

description of 'cloud.Project.NewProjectInfo' response
*/
type GetCloudCreateProjectInfoOK struct {
	Payload *models.CloudProjectNewProjectInfo
}

func (o *GetCloudCreateProjectInfoOK) Error() string {
	return fmt.Sprintf("[GET /cloud/createProjectInfo][%d] getCloudCreateProjectInfoOK  %+v", 200, o.Payload)
}

func (o *GetCloudCreateProjectInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudProjectNewProjectInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCloudCreateProjectInfoDefault creates a GetCloudCreateProjectInfoDefault with default headers values
func NewGetCloudCreateProjectInfoDefault(code int) *GetCloudCreateProjectInfoDefault {
	return &GetCloudCreateProjectInfoDefault{
		_statusCode: code,
	}
}

/*GetCloudCreateProjectInfoDefault handles this case with default header values.

Unexpected error
*/
type GetCloudCreateProjectInfoDefault struct {
	_statusCode int

	Payload *models.GetCloudCreateProjectInfoDefaultBody
}

// Code gets the status code for the get cloud create project info default response
func (o *GetCloudCreateProjectInfoDefault) Code() int {
	return o._statusCode
}

func (o *GetCloudCreateProjectInfoDefault) Error() string {
	return fmt.Sprintf("[GET /cloud/createProjectInfo][%d] GetCloudCreateProjectInfo default  %+v", o._statusCode, o.Payload)
}

func (o *GetCloudCreateProjectInfoDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCloudCreateProjectInfoDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}