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

// GetCloudProjectServiceNameIPLoadbalancingReader is a Reader for the GetCloudProjectServiceNameIPLoadbalancing structure.
type GetCloudProjectServiceNameIPLoadbalancingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCloudProjectServiceNameIPLoadbalancingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCloudProjectServiceNameIPLoadbalancingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetCloudProjectServiceNameIPLoadbalancingDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCloudProjectServiceNameIPLoadbalancingOK creates a GetCloudProjectServiceNameIPLoadbalancingOK with default headers values
func NewGetCloudProjectServiceNameIPLoadbalancingOK() *GetCloudProjectServiceNameIPLoadbalancingOK {
	return &GetCloudProjectServiceNameIPLoadbalancingOK{}
}

/*GetCloudProjectServiceNameIPLoadbalancingOK handles this case with default header values.

return value
*/
type GetCloudProjectServiceNameIPLoadbalancingOK struct {
	Payload []string
}

func (o *GetCloudProjectServiceNameIPLoadbalancingOK) Error() string {
	return fmt.Sprintf("[GET /cloud/project/{serviceName}/ipLoadbalancing][%d] getCloudProjectServiceNameIpLoadbalancingOK  %+v", 200, o.Payload)
}

func (o *GetCloudProjectServiceNameIPLoadbalancingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCloudProjectServiceNameIPLoadbalancingDefault creates a GetCloudProjectServiceNameIPLoadbalancingDefault with default headers values
func NewGetCloudProjectServiceNameIPLoadbalancingDefault(code int) *GetCloudProjectServiceNameIPLoadbalancingDefault {
	return &GetCloudProjectServiceNameIPLoadbalancingDefault{
		_statusCode: code,
	}
}

/*GetCloudProjectServiceNameIPLoadbalancingDefault handles this case with default header values.

Unexpected error
*/
type GetCloudProjectServiceNameIPLoadbalancingDefault struct {
	_statusCode int

	Payload *models.GetCloudProjectServiceNameIPLoadbalancingDefaultBody
}

// Code gets the status code for the get cloud project service name IP loadbalancing default response
func (o *GetCloudProjectServiceNameIPLoadbalancingDefault) Code() int {
	return o._statusCode
}

func (o *GetCloudProjectServiceNameIPLoadbalancingDefault) Error() string {
	return fmt.Sprintf("[GET /cloud/project/{serviceName}/ipLoadbalancing][%d] GetCloudProjectServiceNameIPLoadbalancing default  %+v", o._statusCode, o.Payload)
}

func (o *GetCloudProjectServiceNameIPLoadbalancingDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCloudProjectServiceNameIPLoadbalancingDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
