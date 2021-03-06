package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// VersionUploadSessionRequest version upload session request
// swagger:model VersionUploadSessionRequest
type VersionUploadSessionRequest struct {

	// Name of new file
	FileName string `json:"file_name,omitempty"`

	// The total number of bytes in the file to be uploaded
	FileSize int64 `json:"file_size,omitempty"`
}

// Validate validates this version upload session request
func (m *VersionUploadSessionRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *VersionUploadSessionRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VersionUploadSessionRequest) UnmarshalBinary(b []byte) error {
	var res VersionUploadSessionRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
