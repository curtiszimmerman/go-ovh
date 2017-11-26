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

	"github.com/appscode/go-ovh/ip/models"
)

// GetIPIPMitigationProfilesIPMitigationProfileReader is a Reader for the GetIPIPMitigationProfilesIPMitigationProfile structure.
type GetIPIPMitigationProfilesIPMitigationProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIPIPMitigationProfilesIPMitigationProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetIPIPMitigationProfilesIPMitigationProfileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetIPIPMitigationProfilesIPMitigationProfileDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetIPIPMitigationProfilesIPMitigationProfileOK creates a GetIPIPMitigationProfilesIPMitigationProfileOK with default headers values
func NewGetIPIPMitigationProfilesIPMitigationProfileOK() *GetIPIPMitigationProfilesIPMitigationProfileOK {
	return &GetIPIPMitigationProfilesIPMitigationProfileOK{}
}

/*GetIPIPMitigationProfilesIPMitigationProfileOK handles this case with default header values.

description of 'ip.MitigationProfile' response
*/
type GetIPIPMitigationProfilesIPMitigationProfileOK struct {
	Payload *models.IPMitigationProfile
}

func (o *GetIPIPMitigationProfilesIPMitigationProfileOK) Error() string {
	return fmt.Sprintf("[GET /ip/{ip}/mitigationProfiles/{ipMitigationProfile}][%d] getIpIpMitigationProfilesIpMitigationProfileOK  %+v", 200, o.Payload)
}

func (o *GetIPIPMitigationProfilesIPMitigationProfileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IPMitigationProfile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIPIPMitigationProfilesIPMitigationProfileDefault creates a GetIPIPMitigationProfilesIPMitigationProfileDefault with default headers values
func NewGetIPIPMitigationProfilesIPMitigationProfileDefault(code int) *GetIPIPMitigationProfilesIPMitigationProfileDefault {
	return &GetIPIPMitigationProfilesIPMitigationProfileDefault{
		_statusCode: code,
	}
}

/*GetIPIPMitigationProfilesIPMitigationProfileDefault handles this case with default header values.

Unexpected error
*/
type GetIPIPMitigationProfilesIPMitigationProfileDefault struct {
	_statusCode int

	Payload *models.GetIPIPMitigationProfilesIPMitigationProfileDefaultBody
}

// Code gets the status code for the get IP IP mitigation profiles IP mitigation profile default response
func (o *GetIPIPMitigationProfilesIPMitigationProfileDefault) Code() int {
	return o._statusCode
}

func (o *GetIPIPMitigationProfilesIPMitigationProfileDefault) Error() string {
	return fmt.Sprintf("[GET /ip/{ip}/mitigationProfiles/{ipMitigationProfile}][%d] GetIPIPMitigationProfilesIPMitigationProfile default  %+v", o._statusCode, o.Payload)
}

func (o *GetIPIPMitigationProfilesIPMitigationProfileDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetIPIPMitigationProfilesIPMitigationProfileDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}