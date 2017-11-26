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

// GetCloudServiceNamePcaPcaServiceNameUsageReader is a Reader for the GetCloudServiceNamePcaPcaServiceNameUsage structure.
type GetCloudServiceNamePcaPcaServiceNameUsageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCloudServiceNamePcaPcaServiceNameUsageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCloudServiceNamePcaPcaServiceNameUsageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetCloudServiceNamePcaPcaServiceNameUsageDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCloudServiceNamePcaPcaServiceNameUsageOK creates a GetCloudServiceNamePcaPcaServiceNameUsageOK with default headers values
func NewGetCloudServiceNamePcaPcaServiceNameUsageOK() *GetCloudServiceNamePcaPcaServiceNameUsageOK {
	return &GetCloudServiceNamePcaPcaServiceNameUsageOK{}
}

/*GetCloudServiceNamePcaPcaServiceNameUsageOK handles this case with default header values.

return value
*/
type GetCloudServiceNamePcaPcaServiceNameUsageOK struct {
	Payload int64
}

func (o *GetCloudServiceNamePcaPcaServiceNameUsageOK) Error() string {
	return fmt.Sprintf("[GET /cloud/{serviceName}/pca/{pcaServiceName}/usage][%d] getCloudServiceNamePcaPcaServiceNameUsageOK  %+v", 200, o.Payload)
}

func (o *GetCloudServiceNamePcaPcaServiceNameUsageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCloudServiceNamePcaPcaServiceNameUsageDefault creates a GetCloudServiceNamePcaPcaServiceNameUsageDefault with default headers values
func NewGetCloudServiceNamePcaPcaServiceNameUsageDefault(code int) *GetCloudServiceNamePcaPcaServiceNameUsageDefault {
	return &GetCloudServiceNamePcaPcaServiceNameUsageDefault{
		_statusCode: code,
	}
}

/*GetCloudServiceNamePcaPcaServiceNameUsageDefault handles this case with default header values.

Unexpected error
*/
type GetCloudServiceNamePcaPcaServiceNameUsageDefault struct {
	_statusCode int

	Payload *models.GetCloudServiceNamePcaPcaServiceNameUsageDefaultBody
}

// Code gets the status code for the get cloud service name pca pca service name usage default response
func (o *GetCloudServiceNamePcaPcaServiceNameUsageDefault) Code() int {
	return o._statusCode
}

func (o *GetCloudServiceNamePcaPcaServiceNameUsageDefault) Error() string {
	return fmt.Sprintf("[GET /cloud/{serviceName}/pca/{pcaServiceName}/usage][%d] GetCloudServiceNamePcaPcaServiceNameUsage default  %+v", o._statusCode, o.Payload)
}

func (o *GetCloudServiceNamePcaPcaServiceNameUsageDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCloudServiceNamePcaPcaServiceNameUsageDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}