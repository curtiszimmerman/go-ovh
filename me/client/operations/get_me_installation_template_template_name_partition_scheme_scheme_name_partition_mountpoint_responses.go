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

// GetMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionMountpointReader is a Reader for the GetMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionMountpoint structure.
type GetMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionMountpointReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionMountpointReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionMountpointOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionMountpointDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionMountpointOK creates a GetMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionMountpointOK with default headers values
func NewGetMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionMountpointOK() *GetMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionMountpointOK {
	return &GetMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionMountpointOK{}
}

/*GetMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionMountpointOK handles this case with default header values.

description of 'dedicated.InstallationTemplate.TemplatePartitions' response
*/
type GetMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionMountpointOK struct {
	Payload *models.DedicatedInstallationTemplateTemplatePartitions
}

func (o *GetMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionMountpointOK) Error() string {
	return fmt.Sprintf("[GET /me/installationTemplate/{templateName}/partitionScheme/{schemeName}/partition/{mountpoint}][%d] getMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionMountpointOK  %+v", 200, o.Payload)
}

func (o *GetMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionMountpointOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DedicatedInstallationTemplateTemplatePartitions)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionMountpointDefault creates a GetMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionMountpointDefault with default headers values
func NewGetMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionMountpointDefault(code int) *GetMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionMountpointDefault {
	return &GetMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionMountpointDefault{
		_statusCode: code,
	}
}

/*GetMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionMountpointDefault handles this case with default header values.

Unexpected error
*/
type GetMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionMountpointDefault struct {
	_statusCode int

	Payload *models.GetMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionMountpointDefaultBody
}

// Code gets the status code for the get me installation template template name partition scheme scheme name partition mountpoint default response
func (o *GetMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionMountpointDefault) Code() int {
	return o._statusCode
}

func (o *GetMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionMountpointDefault) Error() string {
	return fmt.Sprintf("[GET /me/installationTemplate/{templateName}/partitionScheme/{schemeName}/partition/{mountpoint}][%d] GetMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionMountpoint default  %+v", o._statusCode, o.Payload)
}

func (o *GetMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionMountpointDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetMeInstallationTemplateTemplateNamePartitionSchemeSchemeNamePartitionMountpointDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}