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

// GetMeAPILogsSelfLogIDReader is a Reader for the GetMeAPILogsSelfLogID structure.
type GetMeAPILogsSelfLogIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMeAPILogsSelfLogIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMeAPILogsSelfLogIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetMeAPILogsSelfLogIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMeAPILogsSelfLogIDOK creates a GetMeAPILogsSelfLogIDOK with default headers values
func NewGetMeAPILogsSelfLogIDOK() *GetMeAPILogsSelfLogIDOK {
	return &GetMeAPILogsSelfLogIDOK{}
}

/*GetMeAPILogsSelfLogIDOK handles this case with default header values.

description of 'api.Log' response
*/
type GetMeAPILogsSelfLogIDOK struct {
	Payload *models.APILog
}

func (o *GetMeAPILogsSelfLogIDOK) Error() string {
	return fmt.Sprintf("[GET /me/api/logs/self/{logId}][%d] getMeApiLogsSelfLogIdOK  %+v", 200, o.Payload)
}

func (o *GetMeAPILogsSelfLogIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APILog)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMeAPILogsSelfLogIDDefault creates a GetMeAPILogsSelfLogIDDefault with default headers values
func NewGetMeAPILogsSelfLogIDDefault(code int) *GetMeAPILogsSelfLogIDDefault {
	return &GetMeAPILogsSelfLogIDDefault{
		_statusCode: code,
	}
}

/*GetMeAPILogsSelfLogIDDefault handles this case with default header values.

Unexpected error
*/
type GetMeAPILogsSelfLogIDDefault struct {
	_statusCode int

	Payload *models.GetMeAPILogsSelfLogIDDefaultBody
}

// Code gets the status code for the get me API logs self log ID default response
func (o *GetMeAPILogsSelfLogIDDefault) Code() int {
	return o._statusCode
}

func (o *GetMeAPILogsSelfLogIDDefault) Error() string {
	return fmt.Sprintf("[GET /me/api/logs/self/{logId}][%d] GetMeAPILogsSelfLogID default  %+v", o._statusCode, o.Payload)
}

func (o *GetMeAPILogsSelfLogIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetMeAPILogsSelfLogIDDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
