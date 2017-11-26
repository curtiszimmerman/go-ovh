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

// GetMeDebtAccountDebtDebtIDOperationReader is a Reader for the GetMeDebtAccountDebtDebtIDOperation structure.
type GetMeDebtAccountDebtDebtIDOperationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMeDebtAccountDebtDebtIDOperationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMeDebtAccountDebtDebtIDOperationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetMeDebtAccountDebtDebtIDOperationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMeDebtAccountDebtDebtIDOperationOK creates a GetMeDebtAccountDebtDebtIDOperationOK with default headers values
func NewGetMeDebtAccountDebtDebtIDOperationOK() *GetMeDebtAccountDebtDebtIDOperationOK {
	return &GetMeDebtAccountDebtDebtIDOperationOK{}
}

/*GetMeDebtAccountDebtDebtIDOperationOK handles this case with default header values.

return value
*/
type GetMeDebtAccountDebtDebtIDOperationOK struct {
	Payload []int64
}

func (o *GetMeDebtAccountDebtDebtIDOperationOK) Error() string {
	return fmt.Sprintf("[GET /me/debtAccount/debt/{debtId}/operation][%d] getMeDebtAccountDebtDebtIdOperationOK  %+v", 200, o.Payload)
}

func (o *GetMeDebtAccountDebtDebtIDOperationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMeDebtAccountDebtDebtIDOperationDefault creates a GetMeDebtAccountDebtDebtIDOperationDefault with default headers values
func NewGetMeDebtAccountDebtDebtIDOperationDefault(code int) *GetMeDebtAccountDebtDebtIDOperationDefault {
	return &GetMeDebtAccountDebtDebtIDOperationDefault{
		_statusCode: code,
	}
}

/*GetMeDebtAccountDebtDebtIDOperationDefault handles this case with default header values.

Unexpected error
*/
type GetMeDebtAccountDebtDebtIDOperationDefault struct {
	_statusCode int

	Payload *models.GetMeDebtAccountDebtDebtIDOperationDefaultBody
}

// Code gets the status code for the get me debt account debt debt ID operation default response
func (o *GetMeDebtAccountDebtDebtIDOperationDefault) Code() int {
	return o._statusCode
}

func (o *GetMeDebtAccountDebtDebtIDOperationDefault) Error() string {
	return fmt.Sprintf("[GET /me/debtAccount/debt/{debtId}/operation][%d] GetMeDebtAccountDebtDebtIDOperation default  %+v", o._statusCode, o.Payload)
}

func (o *GetMeDebtAccountDebtDebtIDOperationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetMeDebtAccountDebtDebtIDOperationDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
