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
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CloudBillingViewUsedCredits UsedCredits
// swagger:model cloud.BillingView.UsedCredits
type CloudBillingViewUsedCredits struct {

	// details
	// Required: true
	Details CloudBillingViewUsedCreditsDetails `json:"details"`

	// Total credit that will be used to pay the bill
	// Required: true
	// Read Only: true
	TotalCredit float64 `json:"totalCredit"`
}

// Validate validates this cloud billing view used credits
func (m *CloudBillingViewUsedCredits) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDetails(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTotalCredit(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudBillingViewUsedCredits) validateDetails(formats strfmt.Registry) error {

	if err := validate.Required("details", "body", m.Details); err != nil {
		return err
	}

	if err := m.Details.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("details")
		}
		return err
	}

	return nil
}

func (m *CloudBillingViewUsedCredits) validateTotalCredit(formats strfmt.Registry) error {

	if err := validate.Required("totalCredit", "body", float64(m.TotalCredit)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloudBillingViewUsedCredits) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudBillingViewUsedCredits) UnmarshalBinary(b []byte) error {
	var res CloudBillingViewUsedCredits
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}