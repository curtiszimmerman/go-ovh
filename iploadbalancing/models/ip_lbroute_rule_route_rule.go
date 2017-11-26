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

// IPLBRouteRuleRouteRule Rule of a route
// swagger:model iplb.RouteRule.RouteRule
type IPLBRouteRuleRouteRule struct {

	// Name of the field to match like "protocol" or "host". See "/ipLoadbalancing/{serviceName}/route/availableRules" for a list of available rules
	Field string `json:"field,omitempty"`

	// Matching operator. Not all operators are available for all fields. See "/availableRules"
	Match string `json:"match,omitempty"`

	// Invert the matching operator effect
	Negate bool `json:"negate,omitempty"`

	// Value to match against this match. Interpretation if this field depends on the match and field
	Pattern string `json:"pattern,omitempty"`

	// Id of your rule
	// Required: true
	// Read Only: true
	RuleID int64 `json:"ruleId"`

	// Name of sub-field, if applicable. This may be a Cookie or Header name for instance
	SubField string `json:"subField,omitempty"`
}

// Validate validates this iplb route rule route rule
func (m *IPLBRouteRuleRouteRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMatch(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRuleID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var iplbRouteRuleRouteRuleTypeMatchPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["contains","endswith","exists","in","internal","is","matches","startswith"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		iplbRouteRuleRouteRuleTypeMatchPropEnum = append(iplbRouteRuleRouteRuleTypeMatchPropEnum, v)
	}
}

const (
	// IPLBRouteRuleRouteRuleMatchContains captures enum value "contains"
	IPLBRouteRuleRouteRuleMatchContains string = "contains"
	// IPLBRouteRuleRouteRuleMatchEndswith captures enum value "endswith"
	IPLBRouteRuleRouteRuleMatchEndswith string = "endswith"
	// IPLBRouteRuleRouteRuleMatchExists captures enum value "exists"
	IPLBRouteRuleRouteRuleMatchExists string = "exists"
	// IPLBRouteRuleRouteRuleMatchIn captures enum value "in"
	IPLBRouteRuleRouteRuleMatchIn string = "in"
	// IPLBRouteRuleRouteRuleMatchInternal captures enum value "internal"
	IPLBRouteRuleRouteRuleMatchInternal string = "internal"
	// IPLBRouteRuleRouteRuleMatchIs captures enum value "is"
	IPLBRouteRuleRouteRuleMatchIs string = "is"
	// IPLBRouteRuleRouteRuleMatchMatches captures enum value "matches"
	IPLBRouteRuleRouteRuleMatchMatches string = "matches"
	// IPLBRouteRuleRouteRuleMatchStartswith captures enum value "startswith"
	IPLBRouteRuleRouteRuleMatchStartswith string = "startswith"
)

// prop value enum
func (m *IPLBRouteRuleRouteRule) validateMatchEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, iplbRouteRuleRouteRuleTypeMatchPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *IPLBRouteRuleRouteRule) validateMatch(formats strfmt.Registry) error {

	if swag.IsZero(m.Match) { // not required
		return nil
	}

	// value enum
	if err := m.validateMatchEnum("match", "body", m.Match); err != nil {
		return err
	}

	return nil
}

func (m *IPLBRouteRuleRouteRule) validateRuleID(formats strfmt.Registry) error {

	if err := validate.Required("ruleId", "body", int64(m.RuleID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IPLBRouteRuleRouteRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IPLBRouteRuleRouteRule) UnmarshalBinary(b []byte) error {
	var res IPLBRouteRuleRouteRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}