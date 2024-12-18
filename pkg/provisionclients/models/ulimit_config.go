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

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UlimitConfig ulimit config
//
// swagger:model UlimitConfig
type UlimitConfig struct {

	// max locked memory
	MaxLockedMemory *string `json:"maxLockedMemory,omitempty"`

	// no file
	NoFile *string `json:"noFile,omitempty"`
}

// Validate validates this ulimit config
func (m *UlimitConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ulimit config based on context it is used
func (m *UlimitConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UlimitConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UlimitConfig) UnmarshalBinary(b []byte) error {
	var res UlimitConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
