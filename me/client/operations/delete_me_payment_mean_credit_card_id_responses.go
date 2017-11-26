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

// DeleteMePaymentMeanCreditCardIDReader is a Reader for the DeleteMePaymentMeanCreditCardID structure.
type DeleteMePaymentMeanCreditCardIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteMePaymentMeanCreditCardIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteMePaymentMeanCreditCardIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteMePaymentMeanCreditCardIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteMePaymentMeanCreditCardIDOK creates a DeleteMePaymentMeanCreditCardIDOK with default headers values
func NewDeleteMePaymentMeanCreditCardIDOK() *DeleteMePaymentMeanCreditCardIDOK {
	return &DeleteMePaymentMeanCreditCardIDOK{}
}

/*DeleteMePaymentMeanCreditCardIDOK handles this case with default header values.

return 'void'
*/
type DeleteMePaymentMeanCreditCardIDOK struct {
}

func (o *DeleteMePaymentMeanCreditCardIDOK) Error() string {
	return fmt.Sprintf("[DELETE /me/paymentMean/creditCard/{id}][%d] deleteMePaymentMeanCreditCardIdOK ", 200)
}

func (o *DeleteMePaymentMeanCreditCardIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteMePaymentMeanCreditCardIDDefault creates a DeleteMePaymentMeanCreditCardIDDefault with default headers values
func NewDeleteMePaymentMeanCreditCardIDDefault(code int) *DeleteMePaymentMeanCreditCardIDDefault {
	return &DeleteMePaymentMeanCreditCardIDDefault{
		_statusCode: code,
	}
}

/*DeleteMePaymentMeanCreditCardIDDefault handles this case with default header values.

Unexpected error
*/
type DeleteMePaymentMeanCreditCardIDDefault struct {
	_statusCode int

	Payload *models.DeleteMePaymentMeanCreditCardIDDefaultBody
}

// Code gets the status code for the delete me payment mean credit card ID default response
func (o *DeleteMePaymentMeanCreditCardIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteMePaymentMeanCreditCardIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /me/paymentMean/creditCard/{id}][%d] DeleteMePaymentMeanCreditCardID default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteMePaymentMeanCreditCardIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteMePaymentMeanCreditCardIDDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}