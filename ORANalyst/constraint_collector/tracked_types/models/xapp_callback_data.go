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

	"github.com/go-openapi/swag"
)

// XappCallbackData xapp callback data
// swagger:model xapp-callback-data
type XappCallbackData struct {

	// event
	Event string `json:"event,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// version
	Version int64 `json:"version,omitempty"`

	// x apps
	XApps string `json:"xApps,omitempty"`
}

// Validate validates this xapp callback data
func (m *XappCallbackData) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *XappCallbackData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *XappCallbackData) UnmarshalBinary(b []byte) error {
	var res XappCallbackData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
