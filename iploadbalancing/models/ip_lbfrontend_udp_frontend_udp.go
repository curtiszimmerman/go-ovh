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

// IPLBFrontendUDPFrontendUDP Frontend UDP
// swagger:model iplb.FrontendUdp.FrontendUdp
type IPLBFrontendUDPFrontendUDP struct {

	// Only attach frontend on these ip. No restriction if null
	DedicatedIPFO []string `json:"dedicatedIpfo"`

	// Default UDP Farm of your frontend
	DefaultFarmID int64 `json:"defaultFarmId,omitempty"`

	// Disable frontend
	Disabled bool `json:"disabled,omitempty"`

	// Human readable name for your frontend, this field is for you
	DisplayName string `json:"displayName,omitempty"`

	// Id of your frontend
	// Required: true
	// Read Only: true
	FrontendID int64 `json:"frontendId"`

	// Listening port(s) on the server
	Port string `json:"port,omitempty"`

	// Zone of you frontend
	Zone string `json:"zone,omitempty"`
}

// Validate validates this iplb frontend Udp frontend Udp
func (m *IPLBFrontendUDPFrontendUDP) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDedicatedIPFO(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateFrontendID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IPLBFrontendUDPFrontendUDP) validateDedicatedIPFO(formats strfmt.Registry) error {

	if swag.IsZero(m.DedicatedIPFO) { // not required
		return nil
	}

	return nil
}

func (m *IPLBFrontendUDPFrontendUDP) validateFrontendID(formats strfmt.Registry) error {

	if err := validate.Required("frontendId", "body", int64(m.FrontendID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IPLBFrontendUDPFrontendUDP) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IPLBFrontendUDPFrontendUDP) UnmarshalBinary(b []byte) error {
	var res IPLBFrontendUDPFrontendUDP
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}