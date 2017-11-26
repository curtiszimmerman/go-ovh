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

// GetCloudProjectServiceNameImageReader is a Reader for the GetCloudProjectServiceNameImage structure.
type GetCloudProjectServiceNameImageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCloudProjectServiceNameImageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCloudProjectServiceNameImageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetCloudProjectServiceNameImageDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCloudProjectServiceNameImageOK creates a GetCloudProjectServiceNameImageOK with default headers values
func NewGetCloudProjectServiceNameImageOK() *GetCloudProjectServiceNameImageOK {
	return &GetCloudProjectServiceNameImageOK{}
}

/*GetCloudProjectServiceNameImageOK handles this case with default header values.

return value
*/
type GetCloudProjectServiceNameImageOK struct {
	Payload models.GetCloudProjectServiceNameImageOKBody
}

func (o *GetCloudProjectServiceNameImageOK) Error() string {
	return fmt.Sprintf("[GET /cloud/project/{serviceName}/image][%d] getCloudProjectServiceNameImageOK  %+v", 200, o.Payload)
}

func (o *GetCloudProjectServiceNameImageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCloudProjectServiceNameImageDefault creates a GetCloudProjectServiceNameImageDefault with default headers values
func NewGetCloudProjectServiceNameImageDefault(code int) *GetCloudProjectServiceNameImageDefault {
	return &GetCloudProjectServiceNameImageDefault{
		_statusCode: code,
	}
}

/*GetCloudProjectServiceNameImageDefault handles this case with default header values.

Unexpected error
*/
type GetCloudProjectServiceNameImageDefault struct {
	_statusCode int

	Payload *models.GetCloudProjectServiceNameImageDefaultBody
}

// Code gets the status code for the get cloud project service name image default response
func (o *GetCloudProjectServiceNameImageDefault) Code() int {
	return o._statusCode
}

func (o *GetCloudProjectServiceNameImageDefault) Error() string {
	return fmt.Sprintf("[GET /cloud/project/{serviceName}/image][%d] GetCloudProjectServiceNameImage default  %+v", o._statusCode, o.Payload)
}

func (o *GetCloudProjectServiceNameImageDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCloudProjectServiceNameImageDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}