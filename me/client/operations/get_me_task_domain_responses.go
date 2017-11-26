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

	"github.com/appscode/go-ovh/me/models"
)

// GetMeTaskDomainReader is a Reader for the GetMeTaskDomain structure.
type GetMeTaskDomainReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMeTaskDomainReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMeTaskDomainOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetMeTaskDomainDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMeTaskDomainOK creates a GetMeTaskDomainOK with default headers values
func NewGetMeTaskDomainOK() *GetMeTaskDomainOK {
	return &GetMeTaskDomainOK{}
}

/*GetMeTaskDomainOK handles this case with default header values.

return value
*/
type GetMeTaskDomainOK struct {
	Payload []int64
}

func (o *GetMeTaskDomainOK) Error() string {
	return fmt.Sprintf("[GET /me/task/domain][%d] getMeTaskDomainOK  %+v", 200, o.Payload)
}

func (o *GetMeTaskDomainOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMeTaskDomainDefault creates a GetMeTaskDomainDefault with default headers values
func NewGetMeTaskDomainDefault(code int) *GetMeTaskDomainDefault {
	return &GetMeTaskDomainDefault{
		_statusCode: code,
	}
}

/*GetMeTaskDomainDefault handles this case with default header values.

Unexpected error
*/
type GetMeTaskDomainDefault struct {
	_statusCode int

	Payload *models.GetMeTaskDomainDefaultBody
}

// Code gets the status code for the get me task domain default response
func (o *GetMeTaskDomainDefault) Code() int {
	return o._statusCode
}

func (o *GetMeTaskDomainDefault) Error() string {
	return fmt.Sprintf("[GET /me/task/domain][%d] GetMeTaskDomain default  %+v", o._statusCode, o.Payload)
}

func (o *GetMeTaskDomainDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetMeTaskDomainDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
