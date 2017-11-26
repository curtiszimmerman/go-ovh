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

// DeleteMeAPIApplicationApplicationIDReader is a Reader for the DeleteMeAPIApplicationApplicationID structure.
type DeleteMeAPIApplicationApplicationIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteMeAPIApplicationApplicationIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteMeAPIApplicationApplicationIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteMeAPIApplicationApplicationIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteMeAPIApplicationApplicationIDOK creates a DeleteMeAPIApplicationApplicationIDOK with default headers values
func NewDeleteMeAPIApplicationApplicationIDOK() *DeleteMeAPIApplicationApplicationIDOK {
	return &DeleteMeAPIApplicationApplicationIDOK{}
}

/*DeleteMeAPIApplicationApplicationIDOK handles this case with default header values.

return 'void'
*/
type DeleteMeAPIApplicationApplicationIDOK struct {
}

func (o *DeleteMeAPIApplicationApplicationIDOK) Error() string {
	return fmt.Sprintf("[DELETE /me/api/application/{applicationId}][%d] deleteMeApiApplicationApplicationIdOK ", 200)
}

func (o *DeleteMeAPIApplicationApplicationIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteMeAPIApplicationApplicationIDDefault creates a DeleteMeAPIApplicationApplicationIDDefault with default headers values
func NewDeleteMeAPIApplicationApplicationIDDefault(code int) *DeleteMeAPIApplicationApplicationIDDefault {
	return &DeleteMeAPIApplicationApplicationIDDefault{
		_statusCode: code,
	}
}

/*DeleteMeAPIApplicationApplicationIDDefault handles this case with default header values.

Unexpected error
*/
type DeleteMeAPIApplicationApplicationIDDefault struct {
	_statusCode int

	Payload *models.DeleteMeAPIApplicationApplicationIDDefaultBody
}

// Code gets the status code for the delete me API application application ID default response
func (o *DeleteMeAPIApplicationApplicationIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteMeAPIApplicationApplicationIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /me/api/application/{applicationId}][%d] DeleteMeAPIApplicationApplicationID default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteMeAPIApplicationApplicationIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteMeAPIApplicationApplicationIDDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
