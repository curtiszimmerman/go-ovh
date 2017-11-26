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

// PostMeDebtAccountDebtDebtIDPayReader is a Reader for the PostMeDebtAccountDebtDebtIDPay structure.
type PostMeDebtAccountDebtDebtIDPayReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostMeDebtAccountDebtDebtIDPayReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostMeDebtAccountDebtDebtIDPayOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostMeDebtAccountDebtDebtIDPayDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostMeDebtAccountDebtDebtIDPayOK creates a PostMeDebtAccountDebtDebtIDPayOK with default headers values
func NewPostMeDebtAccountDebtDebtIDPayOK() *PostMeDebtAccountDebtDebtIDPayOK {
	return &PostMeDebtAccountDebtDebtIDPayOK{}
}

/*PostMeDebtAccountDebtDebtIDPayOK handles this case with default header values.

description of 'billing.Order' response
*/
type PostMeDebtAccountDebtDebtIDPayOK struct {
	Payload *models.BillingOrder
}

func (o *PostMeDebtAccountDebtDebtIDPayOK) Error() string {
	return fmt.Sprintf("[POST /me/debtAccount/debt/{debtId}/pay][%d] postMeDebtAccountDebtDebtIdPayOK  %+v", 200, o.Payload)
}

func (o *PostMeDebtAccountDebtDebtIDPayOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BillingOrder)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostMeDebtAccountDebtDebtIDPayDefault creates a PostMeDebtAccountDebtDebtIDPayDefault with default headers values
func NewPostMeDebtAccountDebtDebtIDPayDefault(code int) *PostMeDebtAccountDebtDebtIDPayDefault {
	return &PostMeDebtAccountDebtDebtIDPayDefault{
		_statusCode: code,
	}
}

/*PostMeDebtAccountDebtDebtIDPayDefault handles this case with default header values.

Unexpected error
*/
type PostMeDebtAccountDebtDebtIDPayDefault struct {
	_statusCode int

	Payload *models.PostMeDebtAccountDebtDebtIDPayDefaultBody
}

// Code gets the status code for the post me debt account debt debt ID pay default response
func (o *PostMeDebtAccountDebtDebtIDPayDefault) Code() int {
	return o._statusCode
}

func (o *PostMeDebtAccountDebtDebtIDPayDefault) Error() string {
	return fmt.Sprintf("[POST /me/debtAccount/debt/{debtId}/pay][%d] PostMeDebtAccountDebtDebtIDPay default  %+v", o._statusCode, o.Payload)
}

func (o *PostMeDebtAccountDebtDebtIDPayDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostMeDebtAccountDebtDebtIDPayDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
