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

// PostMeDepositDepositIDPaidBillsBillIDDebtPayReader is a Reader for the PostMeDepositDepositIDPaidBillsBillIDDebtPay structure.
type PostMeDepositDepositIDPaidBillsBillIDDebtPayReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostMeDepositDepositIDPaidBillsBillIDDebtPayReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostMeDepositDepositIDPaidBillsBillIDDebtPayOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostMeDepositDepositIDPaidBillsBillIDDebtPayDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostMeDepositDepositIDPaidBillsBillIDDebtPayOK creates a PostMeDepositDepositIDPaidBillsBillIDDebtPayOK with default headers values
func NewPostMeDepositDepositIDPaidBillsBillIDDebtPayOK() *PostMeDepositDepositIDPaidBillsBillIDDebtPayOK {
	return &PostMeDepositDepositIDPaidBillsBillIDDebtPayOK{}
}

/*PostMeDepositDepositIDPaidBillsBillIDDebtPayOK handles this case with default header values.

description of 'billing.Order' response
*/
type PostMeDepositDepositIDPaidBillsBillIDDebtPayOK struct {
	Payload *models.BillingOrder
}

func (o *PostMeDepositDepositIDPaidBillsBillIDDebtPayOK) Error() string {
	return fmt.Sprintf("[POST /me/deposit/{depositId}/paidBills/{billId}/debt/pay][%d] postMeDepositDepositIdPaidBillsBillIdDebtPayOK  %+v", 200, o.Payload)
}

func (o *PostMeDepositDepositIDPaidBillsBillIDDebtPayOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BillingOrder)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostMeDepositDepositIDPaidBillsBillIDDebtPayDefault creates a PostMeDepositDepositIDPaidBillsBillIDDebtPayDefault with default headers values
func NewPostMeDepositDepositIDPaidBillsBillIDDebtPayDefault(code int) *PostMeDepositDepositIDPaidBillsBillIDDebtPayDefault {
	return &PostMeDepositDepositIDPaidBillsBillIDDebtPayDefault{
		_statusCode: code,
	}
}

/*PostMeDepositDepositIDPaidBillsBillIDDebtPayDefault handles this case with default header values.

Unexpected error
*/
type PostMeDepositDepositIDPaidBillsBillIDDebtPayDefault struct {
	_statusCode int

	Payload *models.PostMeDepositDepositIDPaidBillsBillIDDebtPayDefaultBody
}

// Code gets the status code for the post me deposit deposit ID paid bills bill ID debt pay default response
func (o *PostMeDepositDepositIDPaidBillsBillIDDebtPayDefault) Code() int {
	return o._statusCode
}

func (o *PostMeDepositDepositIDPaidBillsBillIDDebtPayDefault) Error() string {
	return fmt.Sprintf("[POST /me/deposit/{depositId}/paidBills/{billId}/debt/pay][%d] PostMeDepositDepositIDPaidBillsBillIDDebtPay default  %+v", o._statusCode, o.Payload)
}

func (o *PostMeDepositDepositIDPaidBillsBillIDDebtPayDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostMeDepositDepositIDPaidBillsBillIDDebtPayDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
