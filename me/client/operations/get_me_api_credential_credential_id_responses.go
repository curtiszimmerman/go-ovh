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

// GetMeAPICredentialCredentialIDReader is a Reader for the GetMeAPICredentialCredentialID structure.
type GetMeAPICredentialCredentialIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMeAPICredentialCredentialIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMeAPICredentialCredentialIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetMeAPICredentialCredentialIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMeAPICredentialCredentialIDOK creates a GetMeAPICredentialCredentialIDOK with default headers values
func NewGetMeAPICredentialCredentialIDOK() *GetMeAPICredentialCredentialIDOK {
	return &GetMeAPICredentialCredentialIDOK{}
}

/*GetMeAPICredentialCredentialIDOK handles this case with default header values.

description of 'api.Credential' response
*/
type GetMeAPICredentialCredentialIDOK struct {
	Payload *models.APICredential
}

func (o *GetMeAPICredentialCredentialIDOK) Error() string {
	return fmt.Sprintf("[GET /me/api/credential/{credentialId}][%d] getMeApiCredentialCredentialIdOK  %+v", 200, o.Payload)
}

func (o *GetMeAPICredentialCredentialIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APICredential)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMeAPICredentialCredentialIDDefault creates a GetMeAPICredentialCredentialIDDefault with default headers values
func NewGetMeAPICredentialCredentialIDDefault(code int) *GetMeAPICredentialCredentialIDDefault {
	return &GetMeAPICredentialCredentialIDDefault{
		_statusCode: code,
	}
}

/*GetMeAPICredentialCredentialIDDefault handles this case with default header values.

Unexpected error
*/
type GetMeAPICredentialCredentialIDDefault struct {
	_statusCode int

	Payload *models.GetMeAPICredentialCredentialIDDefaultBody
}

// Code gets the status code for the get me API credential credential ID default response
func (o *GetMeAPICredentialCredentialIDDefault) Code() int {
	return o._statusCode
}

func (o *GetMeAPICredentialCredentialIDDefault) Error() string {
	return fmt.Sprintf("[GET /me/api/credential/{credentialId}][%d] GetMeAPICredentialCredentialID default  %+v", o._statusCode, o.Payload)
}

func (o *GetMeAPICredentialCredentialIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetMeAPICredentialCredentialIDDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
