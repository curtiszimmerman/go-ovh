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

	"github.com/appscode/go-ovh/ip/models"
)

// PostIPLoadBalancingServiceNameSwitchToIPLBNextGenerationAPIReader is a Reader for the PostIPLoadBalancingServiceNameSwitchToIPLBNextGenerationAPI structure.
type PostIPLoadBalancingServiceNameSwitchToIPLBNextGenerationAPIReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostIPLoadBalancingServiceNameSwitchToIPLBNextGenerationAPIReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostIPLoadBalancingServiceNameSwitchToIPLBNextGenerationAPIOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostIPLoadBalancingServiceNameSwitchToIPLBNextGenerationAPIDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostIPLoadBalancingServiceNameSwitchToIPLBNextGenerationAPIOK creates a PostIPLoadBalancingServiceNameSwitchToIPLBNextGenerationAPIOK with default headers values
func NewPostIPLoadBalancingServiceNameSwitchToIPLBNextGenerationAPIOK() *PostIPLoadBalancingServiceNameSwitchToIPLBNextGenerationAPIOK {
	return &PostIPLoadBalancingServiceNameSwitchToIPLBNextGenerationAPIOK{}
}

/*PostIPLoadBalancingServiceNameSwitchToIPLBNextGenerationAPIOK handles this case with default header values.

description of 'iplb.Task.Task' response
*/
type PostIPLoadBalancingServiceNameSwitchToIPLBNextGenerationAPIOK struct {
	Payload *models.IPLBTaskTask
}

func (o *PostIPLoadBalancingServiceNameSwitchToIPLBNextGenerationAPIOK) Error() string {
	return fmt.Sprintf("[POST /ip/loadBalancing/{serviceName}/switchToIplbNextGenerationApi][%d] postIpLoadBalancingServiceNameSwitchToIpLBNextGenerationApiOK  %+v", 200, o.Payload)
}

func (o *PostIPLoadBalancingServiceNameSwitchToIPLBNextGenerationAPIOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IPLBTaskTask)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIPLoadBalancingServiceNameSwitchToIPLBNextGenerationAPIDefault creates a PostIPLoadBalancingServiceNameSwitchToIPLBNextGenerationAPIDefault with default headers values
func NewPostIPLoadBalancingServiceNameSwitchToIPLBNextGenerationAPIDefault(code int) *PostIPLoadBalancingServiceNameSwitchToIPLBNextGenerationAPIDefault {
	return &PostIPLoadBalancingServiceNameSwitchToIPLBNextGenerationAPIDefault{
		_statusCode: code,
	}
}

/*PostIPLoadBalancingServiceNameSwitchToIPLBNextGenerationAPIDefault handles this case with default header values.

Unexpected error
*/
type PostIPLoadBalancingServiceNameSwitchToIPLBNextGenerationAPIDefault struct {
	_statusCode int

	Payload *models.PostIPLoadBalancingServiceNameSwitchToIPLBNextGenerationAPIDefaultBody
}

// Code gets the status code for the post IP load balancing service name switch to IP l b next generation API default response
func (o *PostIPLoadBalancingServiceNameSwitchToIPLBNextGenerationAPIDefault) Code() int {
	return o._statusCode
}

func (o *PostIPLoadBalancingServiceNameSwitchToIPLBNextGenerationAPIDefault) Error() string {
	return fmt.Sprintf("[POST /ip/loadBalancing/{serviceName}/switchToIplbNextGenerationApi][%d] PostIPLoadBalancingServiceNameSwitchToIPLBNextGenerationAPI default  %+v", o._statusCode, o.Payload)
}

func (o *PostIPLoadBalancingServiceNameSwitchToIPLBNextGenerationAPIDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostIPLoadBalancingServiceNameSwitchToIPLBNextGenerationAPIDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
