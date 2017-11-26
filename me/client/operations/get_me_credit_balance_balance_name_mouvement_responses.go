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

// GetMeCreditBalanceBalanceNameMouvementReader is a Reader for the GetMeCreditBalanceBalanceNameMouvement structure.
type GetMeCreditBalanceBalanceNameMouvementReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMeCreditBalanceBalanceNameMouvementReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMeCreditBalanceBalanceNameMouvementOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetMeCreditBalanceBalanceNameMouvementDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMeCreditBalanceBalanceNameMouvementOK creates a GetMeCreditBalanceBalanceNameMouvementOK with default headers values
func NewGetMeCreditBalanceBalanceNameMouvementOK() *GetMeCreditBalanceBalanceNameMouvementOK {
	return &GetMeCreditBalanceBalanceNameMouvementOK{}
}

/*GetMeCreditBalanceBalanceNameMouvementOK handles this case with default header values.

return value
*/
type GetMeCreditBalanceBalanceNameMouvementOK struct {
	Payload []int64
}

func (o *GetMeCreditBalanceBalanceNameMouvementOK) Error() string {
	return fmt.Sprintf("[GET /me/credit/balance/{balanceName}/mouvement][%d] getMeCreditBalanceBalanceNameMouvementOK  %+v", 200, o.Payload)
}

func (o *GetMeCreditBalanceBalanceNameMouvementOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMeCreditBalanceBalanceNameMouvementDefault creates a GetMeCreditBalanceBalanceNameMouvementDefault with default headers values
func NewGetMeCreditBalanceBalanceNameMouvementDefault(code int) *GetMeCreditBalanceBalanceNameMouvementDefault {
	return &GetMeCreditBalanceBalanceNameMouvementDefault{
		_statusCode: code,
	}
}

/*GetMeCreditBalanceBalanceNameMouvementDefault handles this case with default header values.

Unexpected error
*/
type GetMeCreditBalanceBalanceNameMouvementDefault struct {
	_statusCode int

	Payload *models.GetMeCreditBalanceBalanceNameMouvementDefaultBody
}

// Code gets the status code for the get me credit balance balance name mouvement default response
func (o *GetMeCreditBalanceBalanceNameMouvementDefault) Code() int {
	return o._statusCode
}

func (o *GetMeCreditBalanceBalanceNameMouvementDefault) Error() string {
	return fmt.Sprintf("[GET /me/credit/balance/{balanceName}/mouvement][%d] GetMeCreditBalanceBalanceNameMouvement default  %+v", o._statusCode, o.Payload)
}

func (o *GetMeCreditBalanceBalanceNameMouvementDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetMeCreditBalanceBalanceNameMouvementDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
