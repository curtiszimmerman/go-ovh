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

// IPArpBlockedIP ARP blocked IP information
// swagger:model ip.ArpBlockedIp
type IPArpBlockedIP struct {

	// The last blocking date
	// Required: true
	// Read Only: true
	BlockedSince strfmt.DateTime `json:"blockedSince"`

	// your IP
	// Required: true
	// Read Only: true
	IPBlocked string `json:"ipBlocked"`

	// ARP logs
	// Read Only: true
	Logs string `json:"logs,omitempty"`

	// this IP address state
	// Required: true
	// Read Only: true
	State string `json:"state"`

	// Time (in seconds) remaining before you can request your IP to be unblocked
	// Required: true
	// Read Only: true
	Time int64 `json:"time"`
}

// Validate validates this ip arp blocked Ip
func (m *IPArpBlockedIP) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBlockedSince(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateIPBlocked(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTime(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IPArpBlockedIP) validateBlockedSince(formats strfmt.Registry) error {

	if err := validate.Required("blockedSince", "body", strfmt.DateTime(m.BlockedSince)); err != nil {
		return err
	}

	if err := validate.FormatOf("blockedSince", "body", "date-time", m.BlockedSince.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *IPArpBlockedIP) validateIPBlocked(formats strfmt.Registry) error {

	if err := validate.RequiredString("ipBlocked", "body", string(m.IPBlocked)); err != nil {
		return err
	}

	return nil
}

var ipArpBlockedIpTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["blocked","unblocking"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ipArpBlockedIpTypeStatePropEnum = append(ipArpBlockedIpTypeStatePropEnum, v)
	}
}

const (
	// IPArpBlockedIPStateBlocked captures enum value "blocked"
	IPArpBlockedIPStateBlocked string = "blocked"
	// IPArpBlockedIPStateUnblocking captures enum value "unblocking"
	IPArpBlockedIPStateUnblocking string = "unblocking"
)

// prop value enum
func (m *IPArpBlockedIP) validateStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, ipArpBlockedIpTypeStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *IPArpBlockedIP) validateState(formats strfmt.Registry) error {

	if err := validate.RequiredString("state", "body", string(m.State)); err != nil {
		return err
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *IPArpBlockedIP) validateTime(formats strfmt.Registry) error {

	if err := validate.Required("time", "body", int64(m.Time)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IPArpBlockedIP) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IPArpBlockedIP) UnmarshalBinary(b []byte) error {
	var res IPArpBlockedIP
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
