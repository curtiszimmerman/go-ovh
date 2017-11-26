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

	"github.com/appscode/go-ovh/iploadbalancing/models"
)

// DeleteIPLoadbalancingServiceNameHTTPFarmFarmIDReader is a Reader for the DeleteIPLoadbalancingServiceNameHTTPFarmFarmID structure.
type DeleteIPLoadbalancingServiceNameHTTPFarmFarmIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteIPLoadbalancingServiceNameHTTPFarmFarmIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteIPLoadbalancingServiceNameHTTPFarmFarmIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteIPLoadbalancingServiceNameHTTPFarmFarmIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteIPLoadbalancingServiceNameHTTPFarmFarmIDOK creates a DeleteIPLoadbalancingServiceNameHTTPFarmFarmIDOK with default headers values
func NewDeleteIPLoadbalancingServiceNameHTTPFarmFarmIDOK() *DeleteIPLoadbalancingServiceNameHTTPFarmFarmIDOK {
	return &DeleteIPLoadbalancingServiceNameHTTPFarmFarmIDOK{}
}

/*DeleteIPLoadbalancingServiceNameHTTPFarmFarmIDOK handles this case with default header values.

return 'void'
*/
type DeleteIPLoadbalancingServiceNameHTTPFarmFarmIDOK struct {
}

func (o *DeleteIPLoadbalancingServiceNameHTTPFarmFarmIDOK) Error() string {
	return fmt.Sprintf("[DELETE /ipLoadbalancing/{serviceName}/http/farm/{farmId}][%d] deleteIpLoadbalancingServiceNameHttpFarmFarmIdOK ", 200)
}

func (o *DeleteIPLoadbalancingServiceNameHTTPFarmFarmIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteIPLoadbalancingServiceNameHTTPFarmFarmIDDefault creates a DeleteIPLoadbalancingServiceNameHTTPFarmFarmIDDefault with default headers values
func NewDeleteIPLoadbalancingServiceNameHTTPFarmFarmIDDefault(code int) *DeleteIPLoadbalancingServiceNameHTTPFarmFarmIDDefault {
	return &DeleteIPLoadbalancingServiceNameHTTPFarmFarmIDDefault{
		_statusCode: code,
	}
}

/*DeleteIPLoadbalancingServiceNameHTTPFarmFarmIDDefault handles this case with default header values.

Unexpected error
*/
type DeleteIPLoadbalancingServiceNameHTTPFarmFarmIDDefault struct {
	_statusCode int

	Payload *models.DeleteIPLoadbalancingServiceNameHTTPFarmFarmIDDefaultBody
}

// Code gets the status code for the delete IP loadbalancing service name HTTP farm farm ID default response
func (o *DeleteIPLoadbalancingServiceNameHTTPFarmFarmIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteIPLoadbalancingServiceNameHTTPFarmFarmIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /ipLoadbalancing/{serviceName}/http/farm/{farmId}][%d] DeleteIPLoadbalancingServiceNameHTTPFarmFarmID default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteIPLoadbalancingServiceNameHTTPFarmFarmIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteIPLoadbalancingServiceNameHTTPFarmFarmIDDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}