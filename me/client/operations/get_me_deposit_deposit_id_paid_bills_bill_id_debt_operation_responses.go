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

// GetMeDepositDepositIDPaidBillsBillIDDebtOperationReader is a Reader for the GetMeDepositDepositIDPaidBillsBillIDDebtOperation structure.
type GetMeDepositDepositIDPaidBillsBillIDDebtOperationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMeDepositDepositIDPaidBillsBillIDDebtOperationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMeDepositDepositIDPaidBillsBillIDDebtOperationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetMeDepositDepositIDPaidBillsBillIDDebtOperationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMeDepositDepositIDPaidBillsBillIDDebtOperationOK creates a GetMeDepositDepositIDPaidBillsBillIDDebtOperationOK with default headers values
func NewGetMeDepositDepositIDPaidBillsBillIDDebtOperationOK() *GetMeDepositDepositIDPaidBillsBillIDDebtOperationOK {
	return &GetMeDepositDepositIDPaidBillsBillIDDebtOperationOK{}
}

/*GetMeDepositDepositIDPaidBillsBillIDDebtOperationOK handles this case with default header values.

return value
*/
type GetMeDepositDepositIDPaidBillsBillIDDebtOperationOK struct {
	Payload []int64
}

func (o *GetMeDepositDepositIDPaidBillsBillIDDebtOperationOK) Error() string {
	return fmt.Sprintf("[GET /me/deposit/{depositId}/paidBills/{billId}/debt/operation][%d] getMeDepositDepositIdPaidBillsBillIdDebtOperationOK  %+v", 200, o.Payload)
}

func (o *GetMeDepositDepositIDPaidBillsBillIDDebtOperationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMeDepositDepositIDPaidBillsBillIDDebtOperationDefault creates a GetMeDepositDepositIDPaidBillsBillIDDebtOperationDefault with default headers values
func NewGetMeDepositDepositIDPaidBillsBillIDDebtOperationDefault(code int) *GetMeDepositDepositIDPaidBillsBillIDDebtOperationDefault {
	return &GetMeDepositDepositIDPaidBillsBillIDDebtOperationDefault{
		_statusCode: code,
	}
}

/*GetMeDepositDepositIDPaidBillsBillIDDebtOperationDefault handles this case with default header values.

Unexpected error
*/
type GetMeDepositDepositIDPaidBillsBillIDDebtOperationDefault struct {
	_statusCode int

	Payload *models.GetMeDepositDepositIDPaidBillsBillIDDebtOperationDefaultBody
}

// Code gets the status code for the get me deposit deposit ID paid bills bill ID debt operation default response
func (o *GetMeDepositDepositIDPaidBillsBillIDDebtOperationDefault) Code() int {
	return o._statusCode
}

func (o *GetMeDepositDepositIDPaidBillsBillIDDebtOperationDefault) Error() string {
	return fmt.Sprintf("[GET /me/deposit/{depositId}/paidBills/{billId}/debt/operation][%d] GetMeDepositDepositIDPaidBillsBillIDDebtOperation default  %+v", o._statusCode, o.Payload)
}

func (o *GetMeDepositDepositIDPaidBillsBillIDDebtOperationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetMeDepositDepositIDPaidBillsBillIDDebtOperationDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}