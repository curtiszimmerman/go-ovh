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

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BillingPaymentMeanValidation A validation required to add a payment mean
// swagger:model billing.PaymentMeanValidation
type BillingPaymentMeanValidation struct {

	// id
	ID int64 `json:"id,omitempty"`

	// submit Url
	SubmitURL string `json:"submitUrl,omitempty"`

	// url
	URL string `json:"url,omitempty"`

	// validation type
	ValidationType string `json:"validationType,omitempty"`
}

// Validate validates this billing payment mean validation
func (m *BillingPaymentMeanValidation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateValidationType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var billingPaymentMeanValidationTypeValidationTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["creditAccount","documentToSend","simpleValidation"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		billingPaymentMeanValidationTypeValidationTypePropEnum = append(billingPaymentMeanValidationTypeValidationTypePropEnum, v)
	}
}

const (
	// BillingPaymentMeanValidationValidationTypeCreditAccount captures enum value "creditAccount"
	BillingPaymentMeanValidationValidationTypeCreditAccount string = "creditAccount"
	// BillingPaymentMeanValidationValidationTypeDocumentToSend captures enum value "documentToSend"
	BillingPaymentMeanValidationValidationTypeDocumentToSend string = "documentToSend"
	// BillingPaymentMeanValidationValidationTypeSimpleValidation captures enum value "simpleValidation"
	BillingPaymentMeanValidationValidationTypeSimpleValidation string = "simpleValidation"
)

// prop value enum
func (m *BillingPaymentMeanValidation) validateValidationTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, billingPaymentMeanValidationTypeValidationTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *BillingPaymentMeanValidation) validateValidationType(formats strfmt.Registry) error {

	if swag.IsZero(m.ValidationType) { // not required
		return nil
	}

	// value enum
	if err := m.validateValidationTypeEnum("validationType", "body", m.ValidationType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BillingPaymentMeanValidation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BillingPaymentMeanValidation) UnmarshalBinary(b []byte) error {
	var res BillingPaymentMeanValidation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
