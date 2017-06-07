package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// FilePermissions The permissions that the current user has on the file
// swagger:model FilePermissions
type FilePermissions struct {

	// can download
	CanDownload bool `json:"can_download,omitempty"`

	// can invite collaborator
	CanInviteCollaborator bool `json:"can_invite_collaborator,omitempty"`

	// can preview
	CanPreview bool `json:"can_preview,omitempty"`

	// can rename
	CanRename bool `json:"can_rename,omitempty"`

	// can set share access
	CanSetShareAccess bool `json:"can_set_share_access,omitempty"`

	// can share
	CanShare bool `json:"can_share,omitempty"`

	// can upload
	CanUpload bool `json:"can_upload,omitempty"`

	// cand delete
	CandDelete bool `json:"cand_delete,omitempty"`
}

// Validate validates this file permissions
func (m *FilePermissions) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *FilePermissions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FilePermissions) UnmarshalBinary(b []byte) error {
	var res FilePermissions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
