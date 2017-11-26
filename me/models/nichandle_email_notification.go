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

// NichandleEmailNotification Email notification
// swagger:model nichandle.EmailNotification
type NichandleEmailNotification struct {

	// This email body
	// Required: true
	// Read Only: true
	Body string `json:"body"`

	// This email date
	// Required: true
	// Read Only: true
	Date strfmt.DateTime `json:"date"`

	// This email Id
	// Required: true
	// Read Only: true
	ID int64 `json:"id"`

	// This email subject
	// Required: true
	// Read Only: true
	Subject string `json:"subject"`
}

// Validate validates this nichandle email notification
func (m *NichandleEmailNotification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBody(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDate(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSubject(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NichandleEmailNotification) validateBody(formats strfmt.Registry) error {

	if err := validate.RequiredString("body", "body", string(m.Body)); err != nil {
		return err
	}

	return nil
}

func (m *NichandleEmailNotification) validateDate(formats strfmt.Registry) error {

	if err := validate.Required("date", "body", strfmt.DateTime(m.Date)); err != nil {
		return err
	}

	if err := validate.FormatOf("date", "body", "date-time", m.Date.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NichandleEmailNotification) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *NichandleEmailNotification) validateSubject(formats strfmt.Registry) error {

	if err := validate.RequiredString("subject", "body", string(m.Subject)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NichandleEmailNotification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NichandleEmailNotification) UnmarshalBinary(b []byte) error {
	var res NichandleEmailNotification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
