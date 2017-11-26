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

// NichandleDomainTaskArgument Domain operation argument
// swagger:model nichandle.DomainTaskArgument
type NichandleDomainTaskArgument struct {

	// List of accepted formats
	// Read Only: true
	AcceptedFormats []string `json:"acceptedFormats"`

	// List of accepted values
	// Read Only: true
	AcceptedValues []string `json:"acceptedValues"`

	// Description of the argument
	// Read Only: true
	Description string `json:"description,omitempty"`

	// List of impacted field names
	// Read Only: true
	Fields []string `json:"fields"`

	// Key of the argument
	// Required: true
	// Read Only: true
	Key string `json:"key"`

	// Maximum of the content length that you can send
	// Read Only: true
	MaximumSize int64 `json:"maximumSize,omitempty"`

	// Minimum of the content length that you can send
	// Read Only: true
	MinimumSize int64 `json:"minimumSize,omitempty"`

	// True if the argument is in read only
	// Required: true
	// Read Only: true
	ReadOnly bool `json:"readOnly"`

	// Template of the content
	// Read Only: true
	Template string `json:"template,omitempty"`

	// Type of the argument
	// Required: true
	// Read Only: true
	Type string `json:"type"`

	// Value of the argument
	Value string `json:"value,omitempty"`
}

// Validate validates this nichandle domain task argument
func (m *NichandleDomainTaskArgument) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAcceptedFormats(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateAcceptedValues(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateFields(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateKey(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateReadOnly(formats); err != nil {
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

var nichandleDomainTaskArgumentAcceptedFormatsItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["gif","jpeg","jpg","pdf","png"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		nichandleDomainTaskArgumentAcceptedFormatsItemsEnum = append(nichandleDomainTaskArgumentAcceptedFormatsItemsEnum, v)
	}
}

func (m *NichandleDomainTaskArgument) validateAcceptedFormatsItemsEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, nichandleDomainTaskArgumentAcceptedFormatsItemsEnum); err != nil {
		return err
	}
	return nil
}

func (m *NichandleDomainTaskArgument) validateAcceptedFormats(formats strfmt.Registry) error {

	if swag.IsZero(m.AcceptedFormats) { // not required
		return nil
	}

	for i := 0; i < len(m.AcceptedFormats); i++ {

		// value enum
		if err := m.validateAcceptedFormatsItemsEnum("acceptedFormats"+"."+strconv.Itoa(i), "body", m.AcceptedFormats[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *NichandleDomainTaskArgument) validateAcceptedValues(formats strfmt.Registry) error {

	if swag.IsZero(m.AcceptedValues) { // not required
		return nil
	}

	return nil
}

var nichandleDomainTaskArgumentFieldsItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["address.city","address.country","address.line1","address.line2","address.line3","address.otherDetails","address.province","address.zip","birthCity","birthCountry","birthDay","birthZip","cellPhone","companyNationalIdentificationNumber","email","fax","firstName","gender","language","lastName","legalForm","nationalIdentificationNumber","nationality","organisationName","organisationType","phone","spareEmail","vat"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		nichandleDomainTaskArgumentFieldsItemsEnum = append(nichandleDomainTaskArgumentFieldsItemsEnum, v)
	}
}

func (m *NichandleDomainTaskArgument) validateFieldsItemsEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, nichandleDomainTaskArgumentFieldsItemsEnum); err != nil {
		return err
	}
	return nil
}

func (m *NichandleDomainTaskArgument) validateFields(formats strfmt.Registry) error {

	if swag.IsZero(m.Fields) { // not required
		return nil
	}

	for i := 0; i < len(m.Fields); i++ {

		// value enum
		if err := m.validateFieldsItemsEnum("fields"+"."+strconv.Itoa(i), "body", m.Fields[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *NichandleDomainTaskArgument) validateKey(formats strfmt.Registry) error {

	if err := validate.RequiredString("key", "body", string(m.Key)); err != nil {
		return err
	}

	return nil
}

func (m *NichandleDomainTaskArgument) validateReadOnly(formats strfmt.Registry) error {

	if err := validate.Required("readOnly", "body", bool(m.ReadOnly)); err != nil {
		return err
	}

	return nil
}

func (m *NichandleDomainTaskArgument) validateType(formats strfmt.Registry) error {

	if err := validate.RequiredString("type", "body", string(m.Type)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NichandleDomainTaskArgument) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NichandleDomainTaskArgument) UnmarshalBinary(b []byte) error {
	var res NichandleDomainTaskArgument
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}