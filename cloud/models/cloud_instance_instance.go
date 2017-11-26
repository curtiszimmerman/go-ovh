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

// CloudInstanceInstance Instance
// swagger:model cloud.Instance.Instance
type CloudInstanceInstance struct {

	// Instance creation date
	// Required: true
	// Read Only: true
	Created strfmt.DateTime `json:"created"`

	// Instance flavor id
	// Required: true
	// Read Only: true
	FlavorID string `json:"flavorId"`

	// Instance id
	// Required: true
	// Read Only: true
	ID string `json:"id"`

	// Instance image id
	// Required: true
	// Read Only: true
	ImageID string `json:"imageId"`

	// ip addresses
	// Required: true
	IPAddresses CloudInstanceInstanceIPAddresses `json:"ipAddresses"`

	// monthly billing
	MonthlyBilling *CloudInstanceMonthlyBilling `json:"monthlyBilling,omitempty"`

	// Instance name
	// Required: true
	// Read Only: true
	Name string `json:"name"`

	// Instance id
	// Required: true
	// Read Only: true
	Region string `json:"region"`

	// Instance ssh key id
	// Read Only: true
	SSHKeyID string `json:"sshKeyId,omitempty"`

	// Instance status
	// Required: true
	// Read Only: true
	Status string `json:"status"`
}

// Validate validates this cloud instance instance
func (m *CloudInstanceInstance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateFlavorID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateImageID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateIPAddresses(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMonthlyBilling(formats); err != nil {
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

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudInstanceInstance) validateCreated(formats strfmt.Registry) error {

	if err := validate.Required("created", "body", strfmt.DateTime(m.Created)); err != nil {
		return err
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CloudInstanceInstance) validateFlavorID(formats strfmt.Registry) error {

	if err := validate.RequiredString("flavorId", "body", string(m.FlavorID)); err != nil {
		return err
	}

	return nil
}

func (m *CloudInstanceInstance) validateID(formats strfmt.Registry) error {

	if err := validate.RequiredString("id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *CloudInstanceInstance) validateImageID(formats strfmt.Registry) error {

	if err := validate.RequiredString("imageId", "body", string(m.ImageID)); err != nil {
		return err
	}

	return nil
}

func (m *CloudInstanceInstance) validateIPAddresses(formats strfmt.Registry) error {

	if err := validate.Required("ipAddresses", "body", m.IPAddresses); err != nil {
		return err
	}

	if err := m.IPAddresses.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ipAddresses")
		}
		return err
	}

	return nil
}

func (m *CloudInstanceInstance) validateMonthlyBilling(formats strfmt.Registry) error {

	if swag.IsZero(m.MonthlyBilling) { // not required
		return nil
	}

	if m.MonthlyBilling != nil {

		if err := m.MonthlyBilling.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("monthlyBilling")
			}
			return err
		}
	}

	return nil
}

func (m *CloudInstanceInstance) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}

func (m *CloudInstanceInstance) validateRegion(formats strfmt.Registry) error {

	if err := validate.RequiredString("region", "body", string(m.Region)); err != nil {
		return err
	}

	return nil
}

var cloudInstanceInstanceTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ACTIVE","BUILDING","DELETED","ERROR","HARD_REBOOT","PASSWORD","PAUSED","REBOOT","REBUILD","RESCUED","RESIZED","REVERT_RESIZE","SOFT_DELETED","STOPPED","SUSPENDED","UNKNOWN","VERIFY_RESIZE","MIGRATING","RESIZE","BUILD","SHUTOFF","RESCUE","SHELVED","SHELVED_OFFLOADED","RESCUING","UNRESCUING","SNAPSHOTTING","RESUMING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cloudInstanceInstanceTypeStatusPropEnum = append(cloudInstanceInstanceTypeStatusPropEnum, v)
	}
}

const (
	// CloudInstanceInstanceStatusACTIVE captures enum value "ACTIVE"
	CloudInstanceInstanceStatusACTIVE string = "ACTIVE"
	// CloudInstanceInstanceStatusBUILDING captures enum value "BUILDING"
	CloudInstanceInstanceStatusBUILDING string = "BUILDING"
	// CloudInstanceInstanceStatusDELETED captures enum value "DELETED"
	CloudInstanceInstanceStatusDELETED string = "DELETED"
	// CloudInstanceInstanceStatusERROR captures enum value "ERROR"
	CloudInstanceInstanceStatusERROR string = "ERROR"
	// CloudInstanceInstanceStatusHARDREBOOT captures enum value "HARD_REBOOT"
	CloudInstanceInstanceStatusHARDREBOOT string = "HARD_REBOOT"
	// CloudInstanceInstanceStatusPASSWORD captures enum value "PASSWORD"
	CloudInstanceInstanceStatusPASSWORD string = "PASSWORD"
	// CloudInstanceInstanceStatusPAUSED captures enum value "PAUSED"
	CloudInstanceInstanceStatusPAUSED string = "PAUSED"
	// CloudInstanceInstanceStatusREBOOT captures enum value "REBOOT"
	CloudInstanceInstanceStatusREBOOT string = "REBOOT"
	// CloudInstanceInstanceStatusREBUILD captures enum value "REBUILD"
	CloudInstanceInstanceStatusREBUILD string = "REBUILD"
	// CloudInstanceInstanceStatusRESCUED captures enum value "RESCUED"
	CloudInstanceInstanceStatusRESCUED string = "RESCUED"
	// CloudInstanceInstanceStatusRESIZED captures enum value "RESIZED"
	CloudInstanceInstanceStatusRESIZED string = "RESIZED"
	// CloudInstanceInstanceStatusREVERTRESIZE captures enum value "REVERT_RESIZE"
	CloudInstanceInstanceStatusREVERTRESIZE string = "REVERT_RESIZE"
	// CloudInstanceInstanceStatusSOFTDELETED captures enum value "SOFT_DELETED"
	CloudInstanceInstanceStatusSOFTDELETED string = "SOFT_DELETED"
	// CloudInstanceInstanceStatusSTOPPED captures enum value "STOPPED"
	CloudInstanceInstanceStatusSTOPPED string = "STOPPED"
	// CloudInstanceInstanceStatusSUSPENDED captures enum value "SUSPENDED"
	CloudInstanceInstanceStatusSUSPENDED string = "SUSPENDED"
	// CloudInstanceInstanceStatusUNKNOWN captures enum value "UNKNOWN"
	CloudInstanceInstanceStatusUNKNOWN string = "UNKNOWN"
	// CloudInstanceInstanceStatusVERIFYRESIZE captures enum value "VERIFY_RESIZE"
	CloudInstanceInstanceStatusVERIFYRESIZE string = "VERIFY_RESIZE"
	// CloudInstanceInstanceStatusMIGRATING captures enum value "MIGRATING"
	CloudInstanceInstanceStatusMIGRATING string = "MIGRATING"
	// CloudInstanceInstanceStatusRESIZE captures enum value "RESIZE"
	CloudInstanceInstanceStatusRESIZE string = "RESIZE"
	// CloudInstanceInstanceStatusBUILD captures enum value "BUILD"
	CloudInstanceInstanceStatusBUILD string = "BUILD"
	// CloudInstanceInstanceStatusSHUTOFF captures enum value "SHUTOFF"
	CloudInstanceInstanceStatusSHUTOFF string = "SHUTOFF"
	// CloudInstanceInstanceStatusRESCUE captures enum value "RESCUE"
	CloudInstanceInstanceStatusRESCUE string = "RESCUE"
	// CloudInstanceInstanceStatusSHELVED captures enum value "SHELVED"
	CloudInstanceInstanceStatusSHELVED string = "SHELVED"
	// CloudInstanceInstanceStatusSHELVEDOFFLOADED captures enum value "SHELVED_OFFLOADED"
	CloudInstanceInstanceStatusSHELVEDOFFLOADED string = "SHELVED_OFFLOADED"
	// CloudInstanceInstanceStatusRESCUING captures enum value "RESCUING"
	CloudInstanceInstanceStatusRESCUING string = "RESCUING"
	// CloudInstanceInstanceStatusUNRESCUING captures enum value "UNRESCUING"
	CloudInstanceInstanceStatusUNRESCUING string = "UNRESCUING"
	// CloudInstanceInstanceStatusSNAPSHOTTING captures enum value "SNAPSHOTTING"
	CloudInstanceInstanceStatusSNAPSHOTTING string = "SNAPSHOTTING"
	// CloudInstanceInstanceStatusRESUMING captures enum value "RESUMING"
	CloudInstanceInstanceStatusRESUMING string = "RESUMING"
)

// prop value enum
func (m *CloudInstanceInstance) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, cloudInstanceInstanceTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CloudInstanceInstance) validateStatus(formats strfmt.Registry) error {

	if err := validate.RequiredString("status", "body", string(m.Status)); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloudInstanceInstance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudInstanceInstance) UnmarshalBinary(b []byte) error {
	var res CloudInstanceInstance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}