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

// CloudInterfaceInstanceIP IP
// swagger:model cloud.InterfaceInstance.IP
type CloudInterfaceInstanceIP struct {

	// IP address
	// Required: true
	// Read Only: true
	IPAddress string `json:"ipAddress"`

	// subnetwork
	// Required: true
	Subnetwork *CloudAPIResource `json:"subnetwork"`
}

// Validate validates this cloud interface instance IP
func (m *CloudInterfaceInstanceIP) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIPAddress(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSubnetwork(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudInterfaceInstanceIP) validateIPAddress(formats strfmt.Registry) error {

	if err := validate.RequiredString("ipAddress", "body", string(m.IPAddress)); err != nil {
		return err
	}

	return nil
}

func (m *CloudInterfaceInstanceIP) validateSubnetwork(formats strfmt.Registry) error {

	if err := validate.Required("subnetwork", "body", m.Subnetwork); err != nil {
		return err
	}

	if m.Subnetwork != nil {

		if err := m.Subnetwork.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subnetwork")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloudInterfaceInstanceIP) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudInterfaceInstanceIP) UnmarshalBinary(b []byte) error {
	var res CloudInterfaceInstanceIP
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}