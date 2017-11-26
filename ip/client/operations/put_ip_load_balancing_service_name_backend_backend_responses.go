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

// PutIPLoadBalancingServiceNameBackendBackendReader is a Reader for the PutIPLoadBalancingServiceNameBackendBackend structure.
type PutIPLoadBalancingServiceNameBackendBackendReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutIPLoadBalancingServiceNameBackendBackendReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutIPLoadBalancingServiceNameBackendBackendOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPutIPLoadBalancingServiceNameBackendBackendDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutIPLoadBalancingServiceNameBackendBackendOK creates a PutIPLoadBalancingServiceNameBackendBackendOK with default headers values
func NewPutIPLoadBalancingServiceNameBackendBackendOK() *PutIPLoadBalancingServiceNameBackendBackendOK {
	return &PutIPLoadBalancingServiceNameBackendBackendOK{}
}

/*PutIPLoadBalancingServiceNameBackendBackendOK handles this case with default header values.

return 'void'
*/
type PutIPLoadBalancingServiceNameBackendBackendOK struct {
}

func (o *PutIPLoadBalancingServiceNameBackendBackendOK) Error() string {
	return fmt.Sprintf("[PUT /ip/loadBalancing/{serviceName}/backend/{backend}][%d] putIpLoadBalancingServiceNameBackendBackendOK ", 200)
}

func (o *PutIPLoadBalancingServiceNameBackendBackendOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutIPLoadBalancingServiceNameBackendBackendDefault creates a PutIPLoadBalancingServiceNameBackendBackendDefault with default headers values
func NewPutIPLoadBalancingServiceNameBackendBackendDefault(code int) *PutIPLoadBalancingServiceNameBackendBackendDefault {
	return &PutIPLoadBalancingServiceNameBackendBackendDefault{
		_statusCode: code,
	}
}

/*PutIPLoadBalancingServiceNameBackendBackendDefault handles this case with default header values.

Unexpected error
*/
type PutIPLoadBalancingServiceNameBackendBackendDefault struct {
	_statusCode int

	Payload *models.PutIPLoadBalancingServiceNameBackendBackendDefaultBody
}

// Code gets the status code for the put IP load balancing service name backend backend default response
func (o *PutIPLoadBalancingServiceNameBackendBackendDefault) Code() int {
	return o._statusCode
}

func (o *PutIPLoadBalancingServiceNameBackendBackendDefault) Error() string {
	return fmt.Sprintf("[PUT /ip/loadBalancing/{serviceName}/backend/{backend}][%d] PutIPLoadBalancingServiceNameBackendBackend default  %+v", o._statusCode, o.Payload)
}

func (o *PutIPLoadBalancingServiceNameBackendBackendDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PutIPLoadBalancingServiceNameBackendBackendDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}