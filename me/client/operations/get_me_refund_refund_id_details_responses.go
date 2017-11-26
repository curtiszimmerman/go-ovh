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

// GetMeRefundRefundIDDetailsReader is a Reader for the GetMeRefundRefundIDDetails structure.
type GetMeRefundRefundIDDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMeRefundRefundIDDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMeRefundRefundIDDetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetMeRefundRefundIDDetailsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMeRefundRefundIDDetailsOK creates a GetMeRefundRefundIDDetailsOK with default headers values
func NewGetMeRefundRefundIDDetailsOK() *GetMeRefundRefundIDDetailsOK {
	return &GetMeRefundRefundIDDetailsOK{}
}

/*GetMeRefundRefundIDDetailsOK handles this case with default header values.

return value
*/
type GetMeRefundRefundIDDetailsOK struct {
	Payload []string
}

func (o *GetMeRefundRefundIDDetailsOK) Error() string {
	return fmt.Sprintf("[GET /me/refund/{refundId}/details][%d] getMeRefundRefundIdDetailsOK  %+v", 200, o.Payload)
}

func (o *GetMeRefundRefundIDDetailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMeRefundRefundIDDetailsDefault creates a GetMeRefundRefundIDDetailsDefault with default headers values
func NewGetMeRefundRefundIDDetailsDefault(code int) *GetMeRefundRefundIDDetailsDefault {
	return &GetMeRefundRefundIDDetailsDefault{
		_statusCode: code,
	}
}

/*GetMeRefundRefundIDDetailsDefault handles this case with default header values.

Unexpected error
*/
type GetMeRefundRefundIDDetailsDefault struct {
	_statusCode int

	Payload *models.GetMeRefundRefundIDDetailsDefaultBody
}

// Code gets the status code for the get me refund refund ID details default response
func (o *GetMeRefundRefundIDDetailsDefault) Code() int {
	return o._statusCode
}

func (o *GetMeRefundRefundIDDetailsDefault) Error() string {
	return fmt.Sprintf("[GET /me/refund/{refundId}/details][%d] GetMeRefundRefundIDDetails default  %+v", o._statusCode, o.Payload)
}

func (o *GetMeRefundRefundIDDetailsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetMeRefundRefundIDDetailsDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
