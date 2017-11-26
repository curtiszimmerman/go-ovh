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

// PutCloudServiceNamePcaPcaServiceNameServiceInfosReader is a Reader for the PutCloudServiceNamePcaPcaServiceNameServiceInfos structure.
type PutCloudServiceNamePcaPcaServiceNameServiceInfosReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutCloudServiceNamePcaPcaServiceNameServiceInfosReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutCloudServiceNamePcaPcaServiceNameServiceInfosOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPutCloudServiceNamePcaPcaServiceNameServiceInfosDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutCloudServiceNamePcaPcaServiceNameServiceInfosOK creates a PutCloudServiceNamePcaPcaServiceNameServiceInfosOK with default headers values
func NewPutCloudServiceNamePcaPcaServiceNameServiceInfosOK() *PutCloudServiceNamePcaPcaServiceNameServiceInfosOK {
	return &PutCloudServiceNamePcaPcaServiceNameServiceInfosOK{}
}

/*PutCloudServiceNamePcaPcaServiceNameServiceInfosOK handles this case with default header values.

return 'void'
*/
type PutCloudServiceNamePcaPcaServiceNameServiceInfosOK struct {
}

func (o *PutCloudServiceNamePcaPcaServiceNameServiceInfosOK) Error() string {
	return fmt.Sprintf("[PUT /cloud/{serviceName}/pca/{pcaServiceName}/serviceInfos][%d] putCloudServiceNamePcaPcaServiceNameServiceInfosOK ", 200)
}

func (o *PutCloudServiceNamePcaPcaServiceNameServiceInfosOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCloudServiceNamePcaPcaServiceNameServiceInfosDefault creates a PutCloudServiceNamePcaPcaServiceNameServiceInfosDefault with default headers values
func NewPutCloudServiceNamePcaPcaServiceNameServiceInfosDefault(code int) *PutCloudServiceNamePcaPcaServiceNameServiceInfosDefault {
	return &PutCloudServiceNamePcaPcaServiceNameServiceInfosDefault{
		_statusCode: code,
	}
}

/*PutCloudServiceNamePcaPcaServiceNameServiceInfosDefault handles this case with default header values.

Unexpected error
*/
type PutCloudServiceNamePcaPcaServiceNameServiceInfosDefault struct {
	_statusCode int

	Payload *models.PutCloudServiceNamePcaPcaServiceNameServiceInfosDefaultBody
}

// Code gets the status code for the put cloud service name pca pca service name service infos default response
func (o *PutCloudServiceNamePcaPcaServiceNameServiceInfosDefault) Code() int {
	return o._statusCode
}

func (o *PutCloudServiceNamePcaPcaServiceNameServiceInfosDefault) Error() string {
	return fmt.Sprintf("[PUT /cloud/{serviceName}/pca/{pcaServiceName}/serviceInfos][%d] PutCloudServiceNamePcaPcaServiceNameServiceInfos default  %+v", o._statusCode, o.Payload)
}

func (o *PutCloudServiceNamePcaPcaServiceNameServiceInfosDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PutCloudServiceNamePcaPcaServiceNameServiceInfosDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}