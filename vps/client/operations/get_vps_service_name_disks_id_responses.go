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

	"github.com/appscode/go-ovh/vps/models"
)

// GetVpsServiceNameDisksIDReader is a Reader for the GetVpsServiceNameDisksID structure.
type GetVpsServiceNameDisksIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVpsServiceNameDisksIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetVpsServiceNameDisksIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetVpsServiceNameDisksIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetVpsServiceNameDisksIDOK creates a GetVpsServiceNameDisksIDOK with default headers values
func NewGetVpsServiceNameDisksIDOK() *GetVpsServiceNameDisksIDOK {
	return &GetVpsServiceNameDisksIDOK{}
}

/*GetVpsServiceNameDisksIDOK handles this case with default header values.

description of 'vps.Disk' response
*/
type GetVpsServiceNameDisksIDOK struct {
	Payload *models.VpsDisk
}

func (o *GetVpsServiceNameDisksIDOK) Error() string {
	return fmt.Sprintf("[GET /vps/{serviceName}/disks/{id}][%d] getVpsServiceNameDisksIdOK  %+v", 200, o.Payload)
}

func (o *GetVpsServiceNameDisksIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VpsDisk)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVpsServiceNameDisksIDDefault creates a GetVpsServiceNameDisksIDDefault with default headers values
func NewGetVpsServiceNameDisksIDDefault(code int) *GetVpsServiceNameDisksIDDefault {
	return &GetVpsServiceNameDisksIDDefault{
		_statusCode: code,
	}
}

/*GetVpsServiceNameDisksIDDefault handles this case with default header values.

Unexpected error
*/
type GetVpsServiceNameDisksIDDefault struct {
	_statusCode int

	Payload *models.GetVpsServiceNameDisksIDDefaultBody
}

// Code gets the status code for the get vps service name disks ID default response
func (o *GetVpsServiceNameDisksIDDefault) Code() int {
	return o._statusCode
}

func (o *GetVpsServiceNameDisksIDDefault) Error() string {
	return fmt.Sprintf("[GET /vps/{serviceName}/disks/{id}][%d] GetVpsServiceNameDisksID default  %+v", o._statusCode, o.Payload)
}

func (o *GetVpsServiceNameDisksIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetVpsServiceNameDisksIDDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}