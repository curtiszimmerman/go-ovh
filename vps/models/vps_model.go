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
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// VpsModel A structure describing characteristics of a VPS model
// swagger:model vps.Model
type VpsModel struct {

	// available options
	AvailableOptions []string `json:"availableOptions"`

	// datacenter
	Datacenter []string `json:"datacenter"`

	// disk
	Disk int64 `json:"disk,omitempty"`

	// maximum additionnal Ip
	MaximumAdditionnalIP int64 `json:"maximumAdditionnalIp,omitempty"`

	// memory
	Memory int64 `json:"memory,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// offer
	Offer string `json:"offer,omitempty"`

	// vcore
	Vcore int64 `json:"vcore,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this vps model
func (m *VpsModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailableOptions(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDatacenter(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var vpsModelAvailableOptionsItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["additionalDisk","automatedBackup","cpanel","ftpbackup","plesk","snapshot","veeam","windows"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vpsModelAvailableOptionsItemsEnum = append(vpsModelAvailableOptionsItemsEnum, v)
	}
}

func (m *VpsModel) validateAvailableOptionsItemsEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, vpsModelAvailableOptionsItemsEnum); err != nil {
		return err
	}
	return nil
}

func (m *VpsModel) validateAvailableOptions(formats strfmt.Registry) error {

	if swag.IsZero(m.AvailableOptions) { // not required
		return nil
	}

	for i := 0; i < len(m.AvailableOptions); i++ {

		// value enum
		if err := m.validateAvailableOptionsItemsEnum("availableOptions"+"."+strconv.Itoa(i), "body", m.AvailableOptions[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *VpsModel) validateDatacenter(formats strfmt.Registry) error {

	if swag.IsZero(m.Datacenter) { // not required
		return nil
	}

	return nil
}

var vpsModelTypeVersionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["2013v1","2014v1","2015v1","2017v1","2017v2"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vpsModelTypeVersionPropEnum = append(vpsModelTypeVersionPropEnum, v)
	}
}

const (
	// VpsModelVersionNr2013v1 captures enum value "2013v1"
	VpsModelVersionNr2013v1 string = "2013v1"
	// VpsModelVersionNr2014v1 captures enum value "2014v1"
	VpsModelVersionNr2014v1 string = "2014v1"
	// VpsModelVersionNr2015v1 captures enum value "2015v1"
	VpsModelVersionNr2015v1 string = "2015v1"
	// VpsModelVersionNr2017v1 captures enum value "2017v1"
	VpsModelVersionNr2017v1 string = "2017v1"
	// VpsModelVersionNr2017v2 captures enum value "2017v2"
	VpsModelVersionNr2017v2 string = "2017v2"
)

// prop value enum
func (m *VpsModel) validateVersionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, vpsModelTypeVersionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *VpsModel) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	// value enum
	if err := m.validateVersionEnum("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VpsModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VpsModel) UnmarshalBinary(b []byte) error {
	var res VpsModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}