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

// GetMeBillBillIDReader is a Reader for the GetMeBillBillID structure.
type GetMeBillBillIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMeBillBillIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMeBillBillIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetMeBillBillIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMeBillBillIDOK creates a GetMeBillBillIDOK with default headers values
func NewGetMeBillBillIDOK() *GetMeBillBillIDOK {
	return &GetMeBillBillIDOK{}
}

/*GetMeBillBillIDOK handles this case with default header values.

description of 'billing.Bill' response
*/
type GetMeBillBillIDOK struct {
	Payload *models.BillingBill
}

func (o *GetMeBillBillIDOK) Error() string {
	return fmt.Sprintf("[GET /me/bill/{billId}][%d] getMeBillBillIdOK  %+v", 200, o.Payload)
}

func (o *GetMeBillBillIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BillingBill)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMeBillBillIDDefault creates a GetMeBillBillIDDefault with default headers values
func NewGetMeBillBillIDDefault(code int) *GetMeBillBillIDDefault {
	return &GetMeBillBillIDDefault{
		_statusCode: code,
	}
}

/*GetMeBillBillIDDefault handles this case with default header values.

Unexpected error
*/
type GetMeBillBillIDDefault struct {
	_statusCode int

	Payload *models.GetMeBillBillIDDefaultBody
}

// Code gets the status code for the get me bill bill ID default response
func (o *GetMeBillBillIDDefault) Code() int {
	return o._statusCode
}

func (o *GetMeBillBillIDDefault) Error() string {
	return fmt.Sprintf("[GET /me/bill/{billId}][%d] GetMeBillBillID default  %+v", o._statusCode, o.Payload)
}

func (o *GetMeBillBillIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetMeBillBillIDDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
