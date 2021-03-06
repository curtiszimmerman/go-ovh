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

// IPIPTask IP tasks
// swagger:model ip.IpTask
type IPIPTask struct {

	// Details of this task
	// Read Only: true
	Comment string `json:"comment,omitempty"`

	// destination
	Destination *IPRoutedTo `json:"destination,omitempty"`

	// Completion date
	// Read Only: true
	DoneDate strfmt.DateTime `json:"doneDate,omitempty"`

	// Function name
	// Required: true
	// Read Only: true
	Function string `json:"function"`

	// last update
	// Read Only: true
	LastUpdate strfmt.DateTime `json:"lastUpdate,omitempty"`

	// Task Creation date
	// Required: true
	// Read Only: true
	StartDate strfmt.DateTime `json:"startDate"`

	// Task status
	// Required: true
	// Read Only: true
	Status string `json:"status"`

	// the id of the task
	// Required: true
	// Read Only: true
	TaskID int64 `json:"taskId"`
}

// Validate validates this ip Ip task
func (m *IPIPTask) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDestination(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateFunction(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStartDate(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTaskID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IPIPTask) validateDestination(formats strfmt.Registry) error {

	if swag.IsZero(m.Destination) { // not required
		return nil
	}

	if m.Destination != nil {

		if err := m.Destination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("destination")
			}
			return err
		}
	}

	return nil
}

var ipIpTaskTypeFunctionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["arinBlockReassign","changeRipeOrg","checkAndReleaseIp","genericMoveFloatingIp"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ipIpTaskTypeFunctionPropEnum = append(ipIpTaskTypeFunctionPropEnum, v)
	}
}

const (
	// IPIPTaskFunctionArinBlockReassign captures enum value "arinBlockReassign"
	IPIPTaskFunctionArinBlockReassign string = "arinBlockReassign"
	// IPIPTaskFunctionChangeRipeOrg captures enum value "changeRipeOrg"
	IPIPTaskFunctionChangeRipeOrg string = "changeRipeOrg"
	// IPIPTaskFunctionCheckAndReleaseIP captures enum value "checkAndReleaseIp"
	IPIPTaskFunctionCheckAndReleaseIP string = "checkAndReleaseIp"
	// IPIPTaskFunctionGenericMoveFloatingIP captures enum value "genericMoveFloatingIp"
	IPIPTaskFunctionGenericMoveFloatingIP string = "genericMoveFloatingIp"
)

// prop value enum
func (m *IPIPTask) validateFunctionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, ipIpTaskTypeFunctionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *IPIPTask) validateFunction(formats strfmt.Registry) error {

	if err := validate.RequiredString("function", "body", string(m.Function)); err != nil {
		return err
	}

	// value enum
	if err := m.validateFunctionEnum("function", "body", m.Function); err != nil {
		return err
	}

	return nil
}

func (m *IPIPTask) validateStartDate(formats strfmt.Registry) error {

	if err := validate.Required("startDate", "body", strfmt.DateTime(m.StartDate)); err != nil {
		return err
	}

	if err := validate.FormatOf("startDate", "body", "date-time", m.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

var ipIpTaskTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["cancelled","customerError","doing","done","init","ovhError","todo"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ipIpTaskTypeStatusPropEnum = append(ipIpTaskTypeStatusPropEnum, v)
	}
}

const (
	// IPIPTaskStatusCancelled captures enum value "cancelled"
	IPIPTaskStatusCancelled string = "cancelled"
	// IPIPTaskStatusCustomerError captures enum value "customerError"
	IPIPTaskStatusCustomerError string = "customerError"
	// IPIPTaskStatusDoing captures enum value "doing"
	IPIPTaskStatusDoing string = "doing"
	// IPIPTaskStatusDone captures enum value "done"
	IPIPTaskStatusDone string = "done"
	// IPIPTaskStatusInit captures enum value "init"
	IPIPTaskStatusInit string = "init"
	// IPIPTaskStatusOvhError captures enum value "ovhError"
	IPIPTaskStatusOvhError string = "ovhError"
	// IPIPTaskStatusTodo captures enum value "todo"
	IPIPTaskStatusTodo string = "todo"
)

// prop value enum
func (m *IPIPTask) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, ipIpTaskTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *IPIPTask) validateStatus(formats strfmt.Registry) error {

	if err := validate.RequiredString("status", "body", string(m.Status)); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *IPIPTask) validateTaskID(formats strfmt.Registry) error {

	if err := validate.Required("taskId", "body", int64(m.TaskID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IPIPTask) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IPIPTask) UnmarshalBinary(b []byte) error {
	var res IPIPTask
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
