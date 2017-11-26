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

// VpsDatacenter Information about a datacenter of a VPS Virtual Machine
// swagger:model vps.Datacenter
type VpsDatacenter struct {

	// long name
	// Required: true
	// Read Only: true
	LongName string `json:"longName"`

	// name
	// Required: true
	// Read Only: true
	Name string `json:"name"`
}

// Validate validates this vps datacenter
func (m *VpsDatacenter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLongName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VpsDatacenter) validateLongName(formats strfmt.Registry) error {

	if err := validate.RequiredString("longName", "body", string(m.LongName)); err != nil {
		return err
	}

	return nil
}

func (m *VpsDatacenter) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VpsDatacenter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VpsDatacenter) UnmarshalBinary(b []byte) error {
	var res VpsDatacenter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}