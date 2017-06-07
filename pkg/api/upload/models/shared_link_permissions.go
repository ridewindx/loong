package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SharedLinkPermissions shared link permissions
// swagger:model SharedLinkPermissions
type SharedLinkPermissions struct {

	// can download
	CanDownload bool `json:"can_download,omitempty"`

	// can preview
	CanPreview bool `json:"can_preview,omitempty"`
}

// Validate validates this shared link permissions
func (m *SharedLinkPermissions) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *SharedLinkPermissions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SharedLinkPermissions) UnmarshalBinary(b []byte) error {
	var res SharedLinkPermissions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
