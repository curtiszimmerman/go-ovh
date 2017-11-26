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

// IPLBTaskTask IP Load Balancing Operations
// swagger:model iplb.Task.Task
type IPLBTaskTask struct {

	// Operation type
	// Required: true
	// Read Only: true
	Action string `json:"action"`

	// Creation date of your operation
	// Required: true
	// Read Only: true
	CreationDate strfmt.DateTime `json:"creationDate"`

	// Done date of your operation
	// Read Only: true
	DoneDate strfmt.DateTime `json:"doneDate,omitempty"`

	// Id of the operation
	// Required: true
	// Read Only: true
	ID int64 `json:"id"`

	// Operation progress percentage
	// Required: true
	// Read Only: true
	Progress int64 `json:"progress"`

	// Current status of your operation
	// Required: true
	// Read Only: true
	Status string `json:"status"`

	// Zone of your Load Balancer which are updated by current Task
	// Required: true
	// Read Only: true
	Zones []string `json:"zones"`
}

// Validate validates this iplb task task
func (m *IPLBTaskTask) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCreationDate(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProgress(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateZones(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var iplbTaskTaskTypeActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["deleteIplb","deployIplb","install","installIplb","installZone","orderFreeCertificate","orderPaidCertificate","orderSsl","refreshIplb","releaseIplb","releaseIplbZone","reopenIplb","suspendIplb","suspendZone","switchToIplbNextGenerationApi","vrackAttach","vrackDetach"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		iplbTaskTaskTypeActionPropEnum = append(iplbTaskTaskTypeActionPropEnum, v)
	}
}

const (
	// IPLBTaskTaskActionDeleteIPLB captures enum value "deleteIplb"
	IPLBTaskTaskActionDeleteIPLB string = "deleteIplb"
	// IPLBTaskTaskActionDeployIPLB captures enum value "deployIplb"
	IPLBTaskTaskActionDeployIPLB string = "deployIplb"
	// IPLBTaskTaskActionInstall captures enum value "install"
	IPLBTaskTaskActionInstall string = "install"
	// IPLBTaskTaskActionInstallIPLB captures enum value "installIplb"
	IPLBTaskTaskActionInstallIPLB string = "installIplb"
	// IPLBTaskTaskActionInstallZone captures enum value "installZone"
	IPLBTaskTaskActionInstallZone string = "installZone"
	// IPLBTaskTaskActionOrderFreeCertificate captures enum value "orderFreeCertificate"
	IPLBTaskTaskActionOrderFreeCertificate string = "orderFreeCertificate"
	// IPLBTaskTaskActionOrderPaidCertificate captures enum value "orderPaidCertificate"
	IPLBTaskTaskActionOrderPaidCertificate string = "orderPaidCertificate"
	// IPLBTaskTaskActionOrderSsl captures enum value "orderSsl"
	IPLBTaskTaskActionOrderSsl string = "orderSsl"
	// IPLBTaskTaskActionRefreshIPLB captures enum value "refreshIplb"
	IPLBTaskTaskActionRefreshIPLB string = "refreshIplb"
	// IPLBTaskTaskActionReleaseIPLB captures enum value "releaseIplb"
	IPLBTaskTaskActionReleaseIPLB string = "releaseIplb"
	// IPLBTaskTaskActionReleaseIPLBZone captures enum value "releaseIplbZone"
	IPLBTaskTaskActionReleaseIPLBZone string = "releaseIplbZone"
	// IPLBTaskTaskActionReopenIPLB captures enum value "reopenIplb"
	IPLBTaskTaskActionReopenIPLB string = "reopenIplb"
	// IPLBTaskTaskActionSuspendIPLB captures enum value "suspendIplb"
	IPLBTaskTaskActionSuspendIPLB string = "suspendIplb"
	// IPLBTaskTaskActionSuspendZone captures enum value "suspendZone"
	IPLBTaskTaskActionSuspendZone string = "suspendZone"
	// IPLBTaskTaskActionSwitchToIPLBNextGenerationAPI captures enum value "switchToIplbNextGenerationApi"
	IPLBTaskTaskActionSwitchToIPLBNextGenerationAPI string = "switchToIplbNextGenerationApi"
	// IPLBTaskTaskActionVrackAttach captures enum value "vrackAttach"
	IPLBTaskTaskActionVrackAttach string = "vrackAttach"
	// IPLBTaskTaskActionVrackDetach captures enum value "vrackDetach"
	IPLBTaskTaskActionVrackDetach string = "vrackDetach"
)

// prop value enum
func (m *IPLBTaskTask) validateActionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, iplbTaskTaskTypeActionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *IPLBTaskTask) validateAction(formats strfmt.Registry) error {

	if err := validate.RequiredString("action", "body", string(m.Action)); err != nil {
		return err
	}

	// value enum
	if err := m.validateActionEnum("action", "body", m.Action); err != nil {
		return err
	}

	return nil
}

func (m *IPLBTaskTask) validateCreationDate(formats strfmt.Registry) error {

	if err := validate.Required("creationDate", "body", strfmt.DateTime(m.CreationDate)); err != nil {
		return err
	}

	if err := validate.FormatOf("creationDate", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *IPLBTaskTask) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *IPLBTaskTask) validateProgress(formats strfmt.Registry) error {

	if err := validate.Required("progress", "body", int64(m.Progress)); err != nil {
		return err
	}

	return nil
}

var iplbTaskTaskTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["blocked","cancelled","doing","done","error","todo"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		iplbTaskTaskTypeStatusPropEnum = append(iplbTaskTaskTypeStatusPropEnum, v)
	}
}

const (
	// IPLBTaskTaskStatusBlocked captures enum value "blocked"
	IPLBTaskTaskStatusBlocked string = "blocked"
	// IPLBTaskTaskStatusCancelled captures enum value "cancelled"
	IPLBTaskTaskStatusCancelled string = "cancelled"
	// IPLBTaskTaskStatusDoing captures enum value "doing"
	IPLBTaskTaskStatusDoing string = "doing"
	// IPLBTaskTaskStatusDone captures enum value "done"
	IPLBTaskTaskStatusDone string = "done"
	// IPLBTaskTaskStatusError captures enum value "error"
	IPLBTaskTaskStatusError string = "error"
	// IPLBTaskTaskStatusTodo captures enum value "todo"
	IPLBTaskTaskStatusTodo string = "todo"
)

// prop value enum
func (m *IPLBTaskTask) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, iplbTaskTaskTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *IPLBTaskTask) validateStatus(formats strfmt.Registry) error {

	if err := validate.RequiredString("status", "body", string(m.Status)); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *IPLBTaskTask) validateZones(formats strfmt.Registry) error {

	if err := validate.Required("zones", "body", m.Zones); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IPLBTaskTask) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IPLBTaskTask) UnmarshalBinary(b []byte) error {
	var res IPLBTaskTask
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}