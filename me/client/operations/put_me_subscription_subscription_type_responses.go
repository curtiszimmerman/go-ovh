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

// PutMeSubscriptionSubscriptionTypeReader is a Reader for the PutMeSubscriptionSubscriptionType structure.
type PutMeSubscriptionSubscriptionTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutMeSubscriptionSubscriptionTypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutMeSubscriptionSubscriptionTypeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPutMeSubscriptionSubscriptionTypeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutMeSubscriptionSubscriptionTypeOK creates a PutMeSubscriptionSubscriptionTypeOK with default headers values
func NewPutMeSubscriptionSubscriptionTypeOK() *PutMeSubscriptionSubscriptionTypeOK {
	return &PutMeSubscriptionSubscriptionTypeOK{}
}

/*PutMeSubscriptionSubscriptionTypeOK handles this case with default header values.

return 'void'
*/
type PutMeSubscriptionSubscriptionTypeOK struct {
}

func (o *PutMeSubscriptionSubscriptionTypeOK) Error() string {
	return fmt.Sprintf("[PUT /me/subscription/{subscriptionType}][%d] putMeSubscriptionSubscriptionTypeOK ", 200)
}

func (o *PutMeSubscriptionSubscriptionTypeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutMeSubscriptionSubscriptionTypeDefault creates a PutMeSubscriptionSubscriptionTypeDefault with default headers values
func NewPutMeSubscriptionSubscriptionTypeDefault(code int) *PutMeSubscriptionSubscriptionTypeDefault {
	return &PutMeSubscriptionSubscriptionTypeDefault{
		_statusCode: code,
	}
}

/*PutMeSubscriptionSubscriptionTypeDefault handles this case with default header values.

Unexpected error
*/
type PutMeSubscriptionSubscriptionTypeDefault struct {
	_statusCode int

	Payload *models.PutMeSubscriptionSubscriptionTypeDefaultBody
}

// Code gets the status code for the put me subscription subscription type default response
func (o *PutMeSubscriptionSubscriptionTypeDefault) Code() int {
	return o._statusCode
}

func (o *PutMeSubscriptionSubscriptionTypeDefault) Error() string {
	return fmt.Sprintf("[PUT /me/subscription/{subscriptionType}][%d] PutMeSubscriptionSubscriptionType default  %+v", o._statusCode, o.Payload)
}

func (o *PutMeSubscriptionSubscriptionTypeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PutMeSubscriptionSubscriptionTypeDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
