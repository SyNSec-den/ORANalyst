// Code generated by go-swagger; DO NOT EDIT.

// ==================================================================================
// Unless otherwise specified, all software contained herein is licensed
// under the Apache License, Version 2.0 (the "Software License");
// you may not use this software except in compliance with the Software
// License. You may obtain a copy of the Software License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the Software License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the Software License for the specific language governing permissions
// and limitations under the Software License.
//
// ==================================================================================
//
// Unless otherwise specified, all documentation contained herein is licensed
// under the Creative Commons License, Attribution 4.0 Intl. (the
// "Documentation License"); you may not use this documentation except in
// compliance with the Documentation License. You may obtain a copy of the
// Documentation License at
//
// https://creativecommons.org/licenses/by/4.0/
//
// Unless required by applicable law or agreed to in writing, documentation
// distributed under the Documentation License is distributed on an "AS IS"
// BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the Documentation License for the specific language governing
// permissions and limitations under the Documentation License.
// ==================================================================================
//
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Endpoint endpoint
// swagger:model endpoint
type Endpoint struct {

	// end point fqdn
	EndPointFqdn string `json:"EndPointFqdn,omitempty"`

	// end point name
	EndPointName string `json:"EndPointName,omitempty"`

	// end point port
	// Maximum: 65535
	// Minimum: 0
	EndPointPort *uint16 `json:"EndPointPort,omitempty"`
}

// Validate validates this endpoint
func (m *Endpoint) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndPointPort(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Endpoint) validateEndPointPort(formats strfmt.Registry) error {

	if swag.IsZero(m.EndPointPort) { // not required
		return nil
	}

	if err := validate.MinimumInt("EndPointPort", "body", int64(*m.EndPointPort), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("EndPointPort", "body", int64(*m.EndPointPort), 65535, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Endpoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Endpoint) UnmarshalBinary(b []byte) error {
	var res Endpoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
