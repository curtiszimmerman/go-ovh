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

// PostMePaymentMeanCreditCardReader is a Reader for the PostMePaymentMeanCreditCard structure.
type PostMePaymentMeanCreditCardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostMePaymentMeanCreditCardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostMePaymentMeanCreditCardOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostMePaymentMeanCreditCardDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostMePaymentMeanCreditCardOK creates a PostMePaymentMeanCreditCardOK with default headers values
func NewPostMePaymentMeanCreditCardOK() *PostMePaymentMeanCreditCardOK {
	return &PostMePaymentMeanCreditCardOK{}
}

/*PostMePaymentMeanCreditCardOK handles this case with default header values.

description of 'billing.PaymentMeanValidation' response
*/
type PostMePaymentMeanCreditCardOK struct {
	Payload *models.BillingPaymentMeanValidation
}

func (o *PostMePaymentMeanCreditCardOK) Error() string {
	return fmt.Sprintf("[POST /me/paymentMean/creditCard][%d] postMePaymentMeanCreditCardOK  %+v", 200, o.Payload)
}

func (o *PostMePaymentMeanCreditCardOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BillingPaymentMeanValidation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostMePaymentMeanCreditCardDefault creates a PostMePaymentMeanCreditCardDefault with default headers values
func NewPostMePaymentMeanCreditCardDefault(code int) *PostMePaymentMeanCreditCardDefault {
	return &PostMePaymentMeanCreditCardDefault{
		_statusCode: code,
	}
}

/*PostMePaymentMeanCreditCardDefault handles this case with default header values.

Unexpected error
*/
type PostMePaymentMeanCreditCardDefault struct {
	_statusCode int

	Payload *models.PostMePaymentMeanCreditCardDefaultBody
}

// Code gets the status code for the post me payment mean credit card default response
func (o *PostMePaymentMeanCreditCardDefault) Code() int {
	return o._statusCode
}

func (o *PostMePaymentMeanCreditCardDefault) Error() string {
	return fmt.Sprintf("[POST /me/paymentMean/creditCard][%d] PostMePaymentMeanCreditCard default  %+v", o._statusCode, o.Payload)
}

func (o *PostMePaymentMeanCreditCardDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostMePaymentMeanCreditCardDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
