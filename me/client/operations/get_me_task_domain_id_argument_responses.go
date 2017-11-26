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

// GetMeTaskDomainIDArgumentReader is a Reader for the GetMeTaskDomainIDArgument structure.
type GetMeTaskDomainIDArgumentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMeTaskDomainIDArgumentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMeTaskDomainIDArgumentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetMeTaskDomainIDArgumentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMeTaskDomainIDArgumentOK creates a GetMeTaskDomainIDArgumentOK with default headers values
func NewGetMeTaskDomainIDArgumentOK() *GetMeTaskDomainIDArgumentOK {
	return &GetMeTaskDomainIDArgumentOK{}
}

/*GetMeTaskDomainIDArgumentOK handles this case with default header values.

return value
*/
type GetMeTaskDomainIDArgumentOK struct {
	Payload []string
}

func (o *GetMeTaskDomainIDArgumentOK) Error() string {
	return fmt.Sprintf("[GET /me/task/domain/{id}/argument][%d] getMeTaskDomainIdArgumentOK  %+v", 200, o.Payload)
}

func (o *GetMeTaskDomainIDArgumentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMeTaskDomainIDArgumentDefault creates a GetMeTaskDomainIDArgumentDefault with default headers values
func NewGetMeTaskDomainIDArgumentDefault(code int) *GetMeTaskDomainIDArgumentDefault {
	return &GetMeTaskDomainIDArgumentDefault{
		_statusCode: code,
	}
}

/*GetMeTaskDomainIDArgumentDefault handles this case with default header values.

Unexpected error
*/
type GetMeTaskDomainIDArgumentDefault struct {
	_statusCode int

	Payload *models.GetMeTaskDomainIDArgumentDefaultBody
}

// Code gets the status code for the get me task domain ID argument default response
func (o *GetMeTaskDomainIDArgumentDefault) Code() int {
	return o._statusCode
}

func (o *GetMeTaskDomainIDArgumentDefault) Error() string {
	return fmt.Sprintf("[GET /me/task/domain/{id}/argument][%d] GetMeTaskDomainIDArgument default  %+v", o._statusCode, o.Payload)
}

func (o *GetMeTaskDomainIDArgumentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetMeTaskDomainIDArgumentDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
