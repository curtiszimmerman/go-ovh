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

// GetIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDReader is a Reader for the GetIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleID structure.
type GetIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDOK creates a GetIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDOK with default headers values
func NewGetIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDOK() *GetIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDOK {
	return &GetIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDOK{}
}

/*GetIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDOK handles this case with default header values.

description of 'iplb.RouteRule.RouteRule' response
*/
type GetIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDOK struct {
	Payload *models.IPLBRouteRuleRouteRule
}

func (o *GetIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDOK) Error() string {
	return fmt.Sprintf("[GET /ipLoadbalancing/{serviceName}/http/route/{routeId}/rule/{ruleId}][%d] getIpLoadbalancingServiceNameHttpRouteRouteIdRuleRuleIdOK  %+v", 200, o.Payload)
}

func (o *GetIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IPLBRouteRuleRouteRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDDefault creates a GetIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDDefault with default headers values
func NewGetIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDDefault(code int) *GetIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDDefault {
	return &GetIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDDefault{
		_statusCode: code,
	}
}

/*GetIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDDefault handles this case with default header values.

Unexpected error
*/
type GetIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDDefault struct {
	_statusCode int

	Payload *models.GetIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDDefaultBody
}

// Code gets the status code for the get IP loadbalancing service name HTTP route route ID rule rule ID default response
func (o *GetIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDDefault) Code() int {
	return o._statusCode
}

func (o *GetIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDDefault) Error() string {
	return fmt.Sprintf("[GET /ipLoadbalancing/{serviceName}/http/route/{routeId}/rule/{ruleId}][%d] GetIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleID default  %+v", o._statusCode, o.Payload)
}

func (o *GetIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetIPLoadbalancingServiceNameHTTPRouteRouteIDRuleRuleIDDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}