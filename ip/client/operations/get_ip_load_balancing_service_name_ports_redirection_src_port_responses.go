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

// GetIPLoadBalancingServiceNamePortsRedirectionSrcPortReader is a Reader for the GetIPLoadBalancingServiceNamePortsRedirectionSrcPort structure.
type GetIPLoadBalancingServiceNamePortsRedirectionSrcPortReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIPLoadBalancingServiceNamePortsRedirectionSrcPortReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetIPLoadBalancingServiceNamePortsRedirectionSrcPortOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetIPLoadBalancingServiceNamePortsRedirectionSrcPortDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetIPLoadBalancingServiceNamePortsRedirectionSrcPortOK creates a GetIPLoadBalancingServiceNamePortsRedirectionSrcPortOK with default headers values
func NewGetIPLoadBalancingServiceNamePortsRedirectionSrcPortOK() *GetIPLoadBalancingServiceNamePortsRedirectionSrcPortOK {
	return &GetIPLoadBalancingServiceNamePortsRedirectionSrcPortOK{}
}

/*GetIPLoadBalancingServiceNamePortsRedirectionSrcPortOK handles this case with default header values.

description of 'ip.LoadBalancingIp.LoadBalancingPort' response
*/
type GetIPLoadBalancingServiceNamePortsRedirectionSrcPortOK struct {
	Payload *models.IPLoadBalancingIPLoadBalancingPort
}

func (o *GetIPLoadBalancingServiceNamePortsRedirectionSrcPortOK) Error() string {
	return fmt.Sprintf("[GET /ip/loadBalancing/{serviceName}/portsRedirection/{srcPort}][%d] getIpLoadBalancingServiceNamePortsRedirectionSrcPortOK  %+v", 200, o.Payload)
}

func (o *GetIPLoadBalancingServiceNamePortsRedirectionSrcPortOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IPLoadBalancingIPLoadBalancingPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIPLoadBalancingServiceNamePortsRedirectionSrcPortDefault creates a GetIPLoadBalancingServiceNamePortsRedirectionSrcPortDefault with default headers values
func NewGetIPLoadBalancingServiceNamePortsRedirectionSrcPortDefault(code int) *GetIPLoadBalancingServiceNamePortsRedirectionSrcPortDefault {
	return &GetIPLoadBalancingServiceNamePortsRedirectionSrcPortDefault{
		_statusCode: code,
	}
}

/*GetIPLoadBalancingServiceNamePortsRedirectionSrcPortDefault handles this case with default header values.

Unexpected error
*/
type GetIPLoadBalancingServiceNamePortsRedirectionSrcPortDefault struct {
	_statusCode int

	Payload *models.GetIPLoadBalancingServiceNamePortsRedirectionSrcPortDefaultBody
}

// Code gets the status code for the get IP load balancing service name ports redirection src port default response
func (o *GetIPLoadBalancingServiceNamePortsRedirectionSrcPortDefault) Code() int {
	return o._statusCode
}

func (o *GetIPLoadBalancingServiceNamePortsRedirectionSrcPortDefault) Error() string {
	return fmt.Sprintf("[GET /ip/loadBalancing/{serviceName}/portsRedirection/{srcPort}][%d] GetIPLoadBalancingServiceNamePortsRedirectionSrcPort default  %+v", o._statusCode, o.Payload)
}

func (o *GetIPLoadBalancingServiceNamePortsRedirectionSrcPortDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetIPLoadBalancingServiceNamePortsRedirectionSrcPortDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}