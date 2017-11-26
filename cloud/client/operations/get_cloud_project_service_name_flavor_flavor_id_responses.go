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

// GetCloudProjectServiceNameFlavorFlavorIDReader is a Reader for the GetCloudProjectServiceNameFlavorFlavorID structure.
type GetCloudProjectServiceNameFlavorFlavorIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCloudProjectServiceNameFlavorFlavorIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCloudProjectServiceNameFlavorFlavorIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetCloudProjectServiceNameFlavorFlavorIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCloudProjectServiceNameFlavorFlavorIDOK creates a GetCloudProjectServiceNameFlavorFlavorIDOK with default headers values
func NewGetCloudProjectServiceNameFlavorFlavorIDOK() *GetCloudProjectServiceNameFlavorFlavorIDOK {
	return &GetCloudProjectServiceNameFlavorFlavorIDOK{}
}

/*GetCloudProjectServiceNameFlavorFlavorIDOK handles this case with default header values.

description of 'cloud.Flavor.Flavor' response
*/
type GetCloudProjectServiceNameFlavorFlavorIDOK struct {
	Payload *models.CloudFlavorFlavor
}

func (o *GetCloudProjectServiceNameFlavorFlavorIDOK) Error() string {
	return fmt.Sprintf("[GET /cloud/project/{serviceName}/flavor/{flavorId}][%d] getCloudProjectServiceNameFlavorFlavorIdOK  %+v", 200, o.Payload)
}

func (o *GetCloudProjectServiceNameFlavorFlavorIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudFlavorFlavor)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCloudProjectServiceNameFlavorFlavorIDDefault creates a GetCloudProjectServiceNameFlavorFlavorIDDefault with default headers values
func NewGetCloudProjectServiceNameFlavorFlavorIDDefault(code int) *GetCloudProjectServiceNameFlavorFlavorIDDefault {
	return &GetCloudProjectServiceNameFlavorFlavorIDDefault{
		_statusCode: code,
	}
}

/*GetCloudProjectServiceNameFlavorFlavorIDDefault handles this case with default header values.

Unexpected error
*/
type GetCloudProjectServiceNameFlavorFlavorIDDefault struct {
	_statusCode int

	Payload *models.GetCloudProjectServiceNameFlavorFlavorIDDefaultBody
}

// Code gets the status code for the get cloud project service name flavor flavor ID default response
func (o *GetCloudProjectServiceNameFlavorFlavorIDDefault) Code() int {
	return o._statusCode
}

func (o *GetCloudProjectServiceNameFlavorFlavorIDDefault) Error() string {
	return fmt.Sprintf("[GET /cloud/project/{serviceName}/flavor/{flavorId}][%d] GetCloudProjectServiceNameFlavorFlavorID default  %+v", o._statusCode, o.Payload)
}

func (o *GetCloudProjectServiceNameFlavorFlavorIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCloudProjectServiceNameFlavorFlavorIDDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}