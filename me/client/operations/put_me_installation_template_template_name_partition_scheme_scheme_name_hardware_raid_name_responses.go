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

// PutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameReader is a Reader for the PutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidName structure.
type PutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameOK creates a PutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameOK with default headers values
func NewPutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameOK() *PutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameOK {
	return &PutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameOK{}
}

/*PutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameOK handles this case with default header values.

return 'void'
*/
type PutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameOK struct {
}

func (o *PutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameOK) Error() string {
	return fmt.Sprintf("[PUT /me/installationTemplate/{templateName}/partitionScheme/{schemeName}/hardwareRaid/{name}][%d] putMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameOK ", 200)
}

func (o *PutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameDefault creates a PutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameDefault with default headers values
func NewPutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameDefault(code int) *PutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameDefault {
	return &PutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameDefault{
		_statusCode: code,
	}
}

/*PutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameDefault handles this case with default header values.

Unexpected error
*/
type PutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameDefault struct {
	_statusCode int

	Payload *models.PutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameDefaultBody
}

// Code gets the status code for the put me installation template template name partition scheme scheme name hardware raid name default response
func (o *PutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameDefault) Code() int {
	return o._statusCode
}

func (o *PutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameDefault) Error() string {
	return fmt.Sprintf("[PUT /me/installationTemplate/{templateName}/partitionScheme/{schemeName}/hardwareRaid/{name}][%d] PutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidName default  %+v", o._statusCode, o.Payload)
}

func (o *PutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PutMeInstallationTemplateTemplateNamePartitionSchemeSchemeNameHardwareRaidNameDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
