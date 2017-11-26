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

// DeleteIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDReader is a Reader for the DeleteIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleID structure.
type DeleteIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDOK creates a DeleteIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDOK with default headers values
func NewDeleteIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDOK() *DeleteIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDOK {
	return &DeleteIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDOK{}
}

/*DeleteIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDOK handles this case with default header values.

return 'void'
*/
type DeleteIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDOK struct {
}

func (o *DeleteIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDOK) Error() string {
	return fmt.Sprintf("[DELETE /ipLoadbalancing/{serviceName}/http/route/{routeId}/rule/{ruleId}][%d] deleteIpLoadbalancingServiceNameHttpRouteRouteIdRuleRuleIdOK ", 200)
}

func (o *DeleteIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDDefault creates a DeleteIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDDefault with default headers values
func NewDeleteIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDDefault(code int) *DeleteIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDDefault {
	return &DeleteIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDDefault{
		_statusCode: code,
	}
}

/*DeleteIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDDefault handles this case with default header values.

Unexpected error
*/
type DeleteIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDDefault struct {
	_statusCode int

	Payload *models.DeleteIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDDefaultBody
}

// Code gets the status code for the delete IP loadbalancing service name HTTP route route ID rule rule ID default response
func (o *DeleteIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /ipLoadbalancing/{serviceName}/http/route/{routeId}/rule/{ruleId}][%d] DeleteIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleID default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}