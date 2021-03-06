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

// CloudStorageContainer Container
// swagger:model cloud.Storage.Container
type CloudStorageContainer struct {

	// Storage id
	// Required: true
	// Read Only: true
	ID string `json:"id"`

	// Storage name
	// Required: true
	// Read Only: true
	Name string `json:"name"`

	// region
	// Required: true
	// Read Only: true
	Region string `json:"region"`

	// Total bytes stored
	// Required: true
	// Read Only: true
	StoredBytes int64 `json:"storedBytes"`

	// Total objects stored
	// Required: true
	// Read Only: true
	StoredObjects int64 `json:"storedObjects"`
}

// Validate validates this cloud storage container
func (m *CloudStorageContainer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRegion(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStoredBytes(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStoredObjects(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudStorageContainer) validateID(formats strfmt.Registry) error {

	if err := validate.RequiredString("id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *CloudStorageContainer) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}

func (m *CloudStorageContainer) validateRegion(formats strfmt.Registry) error {

	if err := validate.RequiredString("region", "body", string(m.Region)); err != nil {
		return err
	}

	return nil
}

func (m *CloudStorageContainer) validateStoredBytes(formats strfmt.Registry) error {

	if err := validate.Required("storedBytes", "body", int64(m.StoredBytes)); err != nil {
		return err
	}

	return nil
}

func (m *CloudStorageContainer) validateStoredObjects(formats strfmt.Registry) error {

	if err := validate.Required("storedObjects", "body", int64(m.StoredObjects)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloudStorageContainer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudStorageContainer) UnmarshalBinary(b []byte) error {
	var res CloudStorageContainer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
