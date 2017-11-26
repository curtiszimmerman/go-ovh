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

	"github.com/appscode/go-ovh/vrack/models"
)

// DeleteVrackServiceNameIPIPReader is a Reader for the DeleteVrackServiceNameIPIP structure.
type DeleteVrackServiceNameIPIPReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteVrackServiceNameIPIPReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteVrackServiceNameIPIPOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteVrackServiceNameIPIPDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteVrackServiceNameIPIPOK creates a DeleteVrackServiceNameIPIPOK with default headers values
func NewDeleteVrackServiceNameIPIPOK() *DeleteVrackServiceNameIPIPOK {
	return &DeleteVrackServiceNameIPIPOK{}
}

/*DeleteVrackServiceNameIPIPOK handles this case with default header values.

description of 'vrack.Task' response
*/
type DeleteVrackServiceNameIPIPOK struct {
	Payload *models.VrackTask
}

func (o *DeleteVrackServiceNameIPIPOK) Error() string {
	return fmt.Sprintf("[DELETE /vrack/{serviceName}/ip/{ip}][%d] deleteVrackServiceNameIpIpOK  %+v", 200, o.Payload)
}

func (o *DeleteVrackServiceNameIPIPOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VrackTask)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVrackServiceNameIPIPDefault creates a DeleteVrackServiceNameIPIPDefault with default headers values
func NewDeleteVrackServiceNameIPIPDefault(code int) *DeleteVrackServiceNameIPIPDefault {
	return &DeleteVrackServiceNameIPIPDefault{
		_statusCode: code,
	}
}

/*DeleteVrackServiceNameIPIPDefault handles this case with default header values.

Unexpected error
*/
type DeleteVrackServiceNameIPIPDefault struct {
	_statusCode int

	Payload *models.DeleteVrackServiceNameIPIPDefaultBody
}

// Code gets the status code for the delete vrack service name IP IP default response
func (o *DeleteVrackServiceNameIPIPDefault) Code() int {
	return o._statusCode
}

func (o *DeleteVrackServiceNameIPIPDefault) Error() string {
	return fmt.Sprintf("[DELETE /vrack/{serviceName}/ip/{ip}][%d] DeleteVrackServiceNameIPIP default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteVrackServiceNameIPIPDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteVrackServiceNameIPIPDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}