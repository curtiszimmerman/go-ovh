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

// GetMeOrderOrderIDDebtOperationOperationIDAssociatedObjectReader is a Reader for the GetMeOrderOrderIDDebtOperationOperationIDAssociatedObject structure.
type GetMeOrderOrderIDDebtOperationOperationIDAssociatedObjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMeOrderOrderIDDebtOperationOperationIDAssociatedObjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMeOrderOrderIDDebtOperationOperationIDAssociatedObjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetMeOrderOrderIDDebtOperationOperationIDAssociatedObjectDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMeOrderOrderIDDebtOperationOperationIDAssociatedObjectOK creates a GetMeOrderOrderIDDebtOperationOperationIDAssociatedObjectOK with default headers values
func NewGetMeOrderOrderIDDebtOperationOperationIDAssociatedObjectOK() *GetMeOrderOrderIDDebtOperationOperationIDAssociatedObjectOK {
	return &GetMeOrderOrderIDDebtOperationOperationIDAssociatedObjectOK{}
}

/*GetMeOrderOrderIDDebtOperationOperationIDAssociatedObjectOK handles this case with default header values.

description of 'debt.Entry.AssociatedObject' response
*/
type GetMeOrderOrderIDDebtOperationOperationIDAssociatedObjectOK struct {
	Payload *models.DebtEntryAssociatedObject
}

func (o *GetMeOrderOrderIDDebtOperationOperationIDAssociatedObjectOK) Error() string {
	return fmt.Sprintf("[GET /me/order/{orderId}/debt/operation/{operationId}/associatedObject][%d] getMeOrderOrderIdDebtOperationOperationIdAssociatedObjectOK  %+v", 200, o.Payload)
}

func (o *GetMeOrderOrderIDDebtOperationOperationIDAssociatedObjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DebtEntryAssociatedObject)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMeOrderOrderIDDebtOperationOperationIDAssociatedObjectDefault creates a GetMeOrderOrderIDDebtOperationOperationIDAssociatedObjectDefault with default headers values
func NewGetMeOrderOrderIDDebtOperationOperationIDAssociatedObjectDefault(code int) *GetMeOrderOrderIDDebtOperationOperationIDAssociatedObjectDefault {
	return &GetMeOrderOrderIDDebtOperationOperationIDAssociatedObjectDefault{
		_statusCode: code,
	}
}

/*GetMeOrderOrderIDDebtOperationOperationIDAssociatedObjectDefault handles this case with default header values.

Unexpected error
*/
type GetMeOrderOrderIDDebtOperationOperationIDAssociatedObjectDefault struct {
	_statusCode int

	Payload *models.GetMeOrderOrderIDDebtOperationOperationIDAssociatedObjectDefaultBody
}

// Code gets the status code for the get me order order ID debt operation operation ID associated object default response
func (o *GetMeOrderOrderIDDebtOperationOperationIDAssociatedObjectDefault) Code() int {
	return o._statusCode
}

func (o *GetMeOrderOrderIDDebtOperationOperationIDAssociatedObjectDefault) Error() string {
	return fmt.Sprintf("[GET /me/order/{orderId}/debt/operation/{operationId}/associatedObject][%d] GetMeOrderOrderIDDebtOperationOperationIDAssociatedObject default  %+v", o._statusCode, o.Payload)
}

func (o *GetMeOrderOrderIDDebtOperationOperationIDAssociatedObjectDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetMeOrderOrderIDDebtOperationOperationIDAssociatedObjectDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}