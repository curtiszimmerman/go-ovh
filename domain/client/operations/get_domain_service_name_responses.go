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

// GetDomainServiceNameReader is a Reader for the GetDomainServiceName structure.
type GetDomainServiceNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDomainServiceNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDomainServiceNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetDomainServiceNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDomainServiceNameOK creates a GetDomainServiceNameOK with default headers values
func NewGetDomainServiceNameOK() *GetDomainServiceNameOK {
	return &GetDomainServiceNameOK{}
}

/*GetDomainServiceNameOK handles this case with default header values.

description of 'domain.Domain' response
*/
type GetDomainServiceNameOK struct {
	Payload *models.DomainDomain
}

func (o *GetDomainServiceNameOK) Error() string {
	return fmt.Sprintf("[GET /domain/{serviceName}][%d] getDomainServiceNameOK  %+v", 200, o.Payload)
}

func (o *GetDomainServiceNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DomainDomain)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDomainServiceNameDefault creates a GetDomainServiceNameDefault with default headers values
func NewGetDomainServiceNameDefault(code int) *GetDomainServiceNameDefault {
	return &GetDomainServiceNameDefault{
		_statusCode: code,
	}
}

/*GetDomainServiceNameDefault handles this case with default header values.

Unexpected error
*/
type GetDomainServiceNameDefault struct {
	_statusCode int

	Payload *models.GetDomainServiceNameDefaultBody
}

// Code gets the status code for the get domain service name default response
func (o *GetDomainServiceNameDefault) Code() int {
	return o._statusCode
}

func (o *GetDomainServiceNameDefault) Error() string {
	return fmt.Sprintf("[GET /domain/{serviceName}][%d] GetDomainServiceName default  %+v", o._statusCode, o.Payload)
}

func (o *GetDomainServiceNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetDomainServiceNameDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}