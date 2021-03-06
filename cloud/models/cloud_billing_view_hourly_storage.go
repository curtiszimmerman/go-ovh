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

// CloudBillingViewHourlyStorage HourlyStorage
// swagger:model cloud.BillingView.HourlyStorage
type CloudBillingViewHourlyStorage struct {

	// incoming bandwidth
	IncomingBandwidth *CloudBillingViewBandwidthStorage `json:"incomingBandwidth,omitempty"`

	// outgoing bandwidth
	OutgoingBandwidth *CloudBillingViewBandwidthStorage `json:"outgoingBandwidth,omitempty"`

	// Region
	// Required: true
	// Read Only: true
	Region string `json:"region"`

	// stored
	Stored *CloudBillingViewStoredStorage `json:"stored,omitempty"`

	// Total price
	// Required: true
	// Read Only: true
	TotalPrice float64 `json:"totalPrice"`

	// Storage type
	// Required: true
	// Read Only: true
	Type string `json:"type"`
}

// Validate validates this cloud billing view hourly storage
func (m *CloudBillingViewHourlyStorage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIncomingBandwidth(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOutgoingBandwidth(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRegion(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStored(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTotalPrice(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudBillingViewHourlyStorage) validateIncomingBandwidth(formats strfmt.Registry) error {

	if swag.IsZero(m.IncomingBandwidth) { // not required
		return nil
	}

	if m.IncomingBandwidth != nil {

		if err := m.IncomingBandwidth.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("incomingBandwidth")
			}
			return err
		}
	}

	return nil
}

func (m *CloudBillingViewHourlyStorage) validateOutgoingBandwidth(formats strfmt.Registry) error {

	if swag.IsZero(m.OutgoingBandwidth) { // not required
		return nil
	}

	if m.OutgoingBandwidth != nil {

		if err := m.OutgoingBandwidth.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("outgoingBandwidth")
			}
			return err
		}
	}

	return nil
}

func (m *CloudBillingViewHourlyStorage) validateRegion(formats strfmt.Registry) error {

	if err := validate.RequiredString("region", "body", string(m.Region)); err != nil {
		return err
	}

	return nil
}

func (m *CloudBillingViewHourlyStorage) validateStored(formats strfmt.Registry) error {

	if swag.IsZero(m.Stored) { // not required
		return nil
	}

	if m.Stored != nil {

		if err := m.Stored.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stored")
			}
			return err
		}
	}

	return nil
}

func (m *CloudBillingViewHourlyStorage) validateTotalPrice(formats strfmt.Registry) error {

	if err := validate.Required("totalPrice", "body", float64(m.TotalPrice)); err != nil {
		return err
	}

	return nil
}

var cloudBillingViewHourlyStorageTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["pcs","pca"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cloudBillingViewHourlyStorageTypeTypePropEnum = append(cloudBillingViewHourlyStorageTypeTypePropEnum, v)
	}
}

const (
	// CloudBillingViewHourlyStorageTypePcs captures enum value "pcs"
	CloudBillingViewHourlyStorageTypePcs string = "pcs"
	// CloudBillingViewHourlyStorageTypePca captures enum value "pca"
	CloudBillingViewHourlyStorageTypePca string = "pca"
)

// prop value enum
func (m *CloudBillingViewHourlyStorage) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, cloudBillingViewHourlyStorageTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CloudBillingViewHourlyStorage) validateType(formats strfmt.Registry) error {

	if err := validate.RequiredString("type", "body", string(m.Type)); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloudBillingViewHourlyStorage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudBillingViewHourlyStorage) UnmarshalBinary(b []byte) error {
	var res CloudBillingViewHourlyStorage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
