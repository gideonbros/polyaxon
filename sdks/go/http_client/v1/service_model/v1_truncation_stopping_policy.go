// Copyright 2018-2021 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by go-swagger; DO NOT EDIT.

package service_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1TruncationStoppingPolicy Early stopping with truncation stopping, this policy stops a percentage of
// all running runs at every evaluation.
//
// swagger:model v1TruncationStoppingPolicy
type V1TruncationStoppingPolicy struct {

	// Interval/Frequency for applying the policy.
	EvaluationInterval int32 `json:"evaluationInterval,omitempty"`

	// Flag to include or not the succeeded runs in the calculation
	IncludeSucceeded bool `json:"includeSucceeded,omitempty"`

	// Kind of this stopping policy, should be equal to "truncation"
	Kind *string `json:"kind,omitempty"`

	// Min interval (e.g steps) before starting the process
	MinInterval int32 `json:"minInterval,omitempty"`

	// Min samples runs succeeded before starting the process
	MinSamples int32 `json:"minSamples,omitempty"`

	// The percentage of runs to stop, at each evaluation interval.
	// e.g. 1 - 99.
	Percent int32 `json:"percent,omitempty"`
}

// Validate validates this v1 truncation stopping policy
func (m *V1TruncationStoppingPolicy) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 truncation stopping policy based on context it is used
func (m *V1TruncationStoppingPolicy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1TruncationStoppingPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1TruncationStoppingPolicy) UnmarshalBinary(b []byte) error {
	var res V1TruncationStoppingPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
