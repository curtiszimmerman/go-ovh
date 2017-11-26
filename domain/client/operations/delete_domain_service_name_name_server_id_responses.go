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

	"github.com/appscode/go-ovh/domain/models"
)

// DeleteDomainServiceNameNameServerIDReader is a Reader for the DeleteDomainServiceNameNameServerID structure.
type DeleteDomainServiceNameNameServerIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDomainServiceNameNameServerIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteDomainServiceNameNameServerIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteDomainServiceNameNameServerIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteDomainServiceNameNameServerIDOK creates a DeleteDomainServiceNameNameServerIDOK with default headers values
func NewDeleteDomainServiceNameNameServerIDOK() *DeleteDomainServiceNameNameServerIDOK {
	return &DeleteDomainServiceNameNameServerIDOK{}
}

/*DeleteDomainServiceNameNameServerIDOK handles this case with default header values.

description of 'domain.Task' response
*/
type DeleteDomainServiceNameNameServerIDOK struct {
	Payload *models.DomainTask
}

func (o *DeleteDomainServiceNameNameServerIDOK) Error() string {
	return fmt.Sprintf("[DELETE /domain/{serviceName}/nameServer/{id}][%d] deleteDomainServiceNameNameServerIdOK  %+v", 200, o.Payload)
}

func (o *DeleteDomainServiceNameNameServerIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DomainTask)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDomainServiceNameNameServerIDDefault creates a DeleteDomainServiceNameNameServerIDDefault with default headers values
func NewDeleteDomainServiceNameNameServerIDDefault(code int) *DeleteDomainServiceNameNameServerIDDefault {
	return &DeleteDomainServiceNameNameServerIDDefault{
		_statusCode: code,
	}
}

/*DeleteDomainServiceNameNameServerIDDefault handles this case with default header values.

Unexpected error
*/
type DeleteDomainServiceNameNameServerIDDefault struct {
	_statusCode int

	Payload *models.DeleteDomainServiceNameNameServerIDDefaultBody
}

// Code gets the status code for the delete domain service name name server ID default response
func (o *DeleteDomainServiceNameNameServerIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteDomainServiceNameNameServerIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /domain/{serviceName}/nameServer/{id}][%d] DeleteDomainServiceNameNameServerID default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteDomainServiceNameNameServerIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteDomainServiceNameNameServerIDDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}