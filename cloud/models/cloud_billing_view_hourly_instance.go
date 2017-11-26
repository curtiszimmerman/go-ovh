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

// CloudBillingViewHourlyInstance HourlyInstance
// swagger:model cloud.BillingView.HourlyInstance
type CloudBillingViewHourlyInstance struct {

	// details
	// Required: true
	Details CloudBillingViewHourlyInstanceDetails `json:"details"`

	// quantity
	// Required: true
	Quantity *CloudBillingViewQuantity `json:"quantity"`

	// Instance reference
	// Required: true
	// Read Only: true
	Reference string `json:"reference"`

	// Instance region
	// Required: true
	// Read Only: true
	Region string `json:"region"`

	// Total price
	// Required: true
	// Read Only: true
	TotalPrice float64 `json:"totalPrice"`
}

// Validate validates this cloud billing view hourly instance
func (m *CloudBillingViewHourlyInstance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDetails(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateQuantity(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateReference(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRegion(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTotalPrice(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudBillingViewHourlyInstance) validateDetails(formats strfmt.Registry) error {

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

func (m *CloudBillingViewHourlyInstance) validateQuantity(formats strfmt.Registry) error {

	if err := validate.Required("quantity", "body", m.Quantity); err != nil {
		return err
	}

	if m.Quantity != nil {

		if err := m.Quantity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("quantity")
			}
			return err
		}
	}

	return nil
}

func (m *CloudBillingViewHourlyInstance) validateReference(formats strfmt.Registry) error {

	if err := validate.RequiredString("reference", "body", string(m.Reference)); err != nil {
		return err
	}

	return nil
}

func (m *CloudBillingViewHourlyInstance) validateRegion(formats strfmt.Registry) error {

	if err := validate.RequiredString("region", "body", string(m.Region)); err != nil {
		return err
	}

	return nil
}

func (m *CloudBillingViewHourlyInstance) validateTotalPrice(formats strfmt.Registry) error {

	if err := validate.Required("totalPrice", "body", float64(m.TotalPrice)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloudBillingViewHourlyInstance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudBillingViewHourlyInstance) UnmarshalBinary(b []byte) error {
	var res CloudBillingViewHourlyInstance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}