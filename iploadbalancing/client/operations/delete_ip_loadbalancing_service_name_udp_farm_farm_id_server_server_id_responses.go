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

// DeleteIPLoadbalancingServiceNameUDPFarmFarmIDServerServerIDReader is a Reader for the DeleteIPLoadbalancingServiceNameUDPFarmFarmIDServerServerID structure.
type DeleteIPLoadbalancingServiceNameUDPFarmFarmIDServerServerIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteIPLoadbalancingServiceNameUDPFarmFarmIDServerServerIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteIPLoadbalancingServiceNameUDPFarmFarmIDServerServerIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteIPLoadbalancingServiceNameUDPFarmFarmIDServerServerIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteIPLoadbalancingServiceNameUDPFarmFarmIDServerServerIDOK creates a DeleteIPLoadbalancingServiceNameUDPFarmFarmIDServerServerIDOK with default headers values
func NewDeleteIPLoadbalancingServiceNameUDPFarmFarmIDServerServerIDOK() *DeleteIPLoadbalancingServiceNameUDPFarmFarmIDServerServerIDOK {
	return &DeleteIPLoadbalancingServiceNameUDPFarmFarmIDServerServerIDOK{}
}

/*DeleteIPLoadbalancingServiceNameUDPFarmFarmIDServerServerIDOK handles this case with default header values.

return 'void'
*/
type DeleteIPLoadbalancingServiceNameUDPFarmFarmIDServerServerIDOK struct {
}

func (o *DeleteIPLoadbalancingServiceNameUDPFarmFarmIDServerServerIDOK) Error() string {
	return fmt.Sprintf("[DELETE /ipLoadbalancing/{serviceName}/udp/farm/{farmId}/server/{serverId}][%d] deleteIpLoadbalancingServiceNameUdpFarmFarmIdServerServerIdOK ", 200)
}

func (o *DeleteIPLoadbalancingServiceNameUDPFarmFarmIDServerServerIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteIPLoadbalancingServiceNameUDPFarmFarmIDServerServerIDDefault creates a DeleteIPLoadbalancingServiceNameUDPFarmFarmIDServerServerIDDefault with default headers values
func NewDeleteIPLoadbalancingServiceNameUDPFarmFarmIDServerServerIDDefault(code int) *DeleteIPLoadbalancingServiceNameUDPFarmFarmIDServerServerIDDefault {
	return &DeleteIPLoadbalancingServiceNameUDPFarmFarmIDServerServerIDDefault{
		_statusCode: code,
	}
}

/*DeleteIPLoadbalancingServiceNameUDPFarmFarmIDServerServerIDDefault handles this case with default header values.

Unexpected error
*/
type DeleteIPLoadbalancingServiceNameUDPFarmFarmIDServerServerIDDefault struct {
	_statusCode int

	Payload *models.DeleteIPLoadbalancingServiceNameUDPFarmFarmIDServerServerIDDefaultBody
}

// Code gets the status code for the delete IP loadbalancing service name UDP farm farm ID server server ID default response
func (o *DeleteIPLoadbalancingServiceNameUDPFarmFarmIDServerServerIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteIPLoadbalancingServiceNameUDPFarmFarmIDServerServerIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /ipLoadbalancing/{serviceName}/udp/farm/{farmId}/server/{serverId}][%d] DeleteIPLoadbalancingServiceNameUDPFarmFarmIDServerServerID default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteIPLoadbalancingServiceNameUDPFarmFarmIDServerServerIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteIPLoadbalancingServiceNameUDPFarmFarmIDServerServerIDDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}