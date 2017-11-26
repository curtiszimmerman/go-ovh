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

// PostCloudProjectServiceNameNetworkPrivateNetworkIDSubnetParamsBody post cloud project service name network private network Id subnet params body
// swagger:model postCloudProjectServiceNameNetworkPrivateNetworkIdSubnetParamsBody
type PostCloudProjectServiceNameNetworkPrivateNetworkIDSubnetParamsBody struct {

	// dhcp
	// Required: true
	Dhcp bool `json:"dhcp"`

	// end
	// Required: true
	End *string `json:"end"`

	// network
	// Required: true
	Network *string `json:"network"`

	// no gateway
	// Required: true
	NoGateway bool `json:"noGateway"`

	// region
	// Required: true
	Region *string `json:"region"`

	// start
	// Required: true
	Start *string `json:"start"`
}

// Validate validates this post cloud project service name network private network Id subnet params body
func (m *PostCloudProjectServiceNameNetworkPrivateNetworkIDSubnetParamsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDhcp(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEnd(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNetwork(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNoGateway(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRegion(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStart(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostCloudProjectServiceNameNetworkPrivateNetworkIDSubnetParamsBody) validateDhcp(formats strfmt.Registry) error {

	if err := validate.Required("dhcp", "body", bool(m.Dhcp)); err != nil {
		return err
	}

	return nil
}

func (m *PostCloudProjectServiceNameNetworkPrivateNetworkIDSubnetParamsBody) validateEnd(formats strfmt.Registry) error {

	if err := validate.Required("end", "body", m.End); err != nil {
		return err
	}

	return nil
}

func (m *PostCloudProjectServiceNameNetworkPrivateNetworkIDSubnetParamsBody) validateNetwork(formats strfmt.Registry) error {

	if err := validate.Required("network", "body", m.Network); err != nil {
		return err
	}

	return nil
}

func (m *PostCloudProjectServiceNameNetworkPrivateNetworkIDSubnetParamsBody) validateNoGateway(formats strfmt.Registry) error {

	if err := validate.Required("noGateway", "body", bool(m.NoGateway)); err != nil {
		return err
	}

	return nil
}

func (m *PostCloudProjectServiceNameNetworkPrivateNetworkIDSubnetParamsBody) validateRegion(formats strfmt.Registry) error {

	if err := validate.Required("region", "body", m.Region); err != nil {
		return err
	}

	return nil
}

func (m *PostCloudProjectServiceNameNetworkPrivateNetworkIDSubnetParamsBody) validateStart(formats strfmt.Registry) error {

	if err := validate.Required("start", "body", m.Start); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostCloudProjectServiceNameNetworkPrivateNetworkIDSubnetParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostCloudProjectServiceNameNetworkPrivateNetworkIDSubnetParamsBody) UnmarshalBinary(b []byte) error {
	var res PostCloudProjectServiceNameNetworkPrivateNetworkIDSubnetParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}