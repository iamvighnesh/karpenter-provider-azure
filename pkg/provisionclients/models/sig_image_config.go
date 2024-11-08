/*
Portions Copyright (c) Microsoft Corporation.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SigImageConfig sig image config
//
// swagger:model SigImageConfig
type SigImageConfig struct {

	// sig image config template
	SigImageConfigTemplate *SigImageConfigTemplate `json:"sigImageConfigTemplate,omitempty"`

	// subscription ID
	SubscriptionID *string `json:"subscriptionID,omitempty"`
}

// Validate validates this sig image config
func (m *SigImageConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSigImageConfigTemplate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SigImageConfig) validateSigImageConfigTemplate(formats strfmt.Registry) error {
	if swag.IsZero(m.SigImageConfigTemplate) { // not required
		return nil
	}

	if m.SigImageConfigTemplate != nil {
		if err := m.SigImageConfigTemplate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sigImageConfigTemplate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sigImageConfigTemplate")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this sig image config based on the context it is used
func (m *SigImageConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSigImageConfigTemplate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SigImageConfig) contextValidateSigImageConfigTemplate(ctx context.Context, formats strfmt.Registry) error {

	if m.SigImageConfigTemplate != nil {

		if swag.IsZero(m.SigImageConfigTemplate) { // not required
			return nil
		}

		if err := m.SigImageConfigTemplate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sigImageConfigTemplate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sigImageConfigTemplate")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SigImageConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SigImageConfig) UnmarshalBinary(b []byte) error {
	var res SigImageConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}