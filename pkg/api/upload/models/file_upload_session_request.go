package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// FileUploadSessionRequest file upload session request
// swagger:model FileUploadSessionRequest
type FileUploadSessionRequest struct {

	// Name of new file
	FileName string `json:"file_name,omitempty"`

	// The total number of bytes in the file to be uploaded
	FileSize int64 `json:"file_size,omitempty"`

	// The ID of the folder that will contain the new file
	FolderID string `json:"folder_id,omitempty"`
}

// Validate validates this file upload session request
func (m *FileUploadSessionRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *FileUploadSessionRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FileUploadSessionRequest) UnmarshalBinary(b []byte) error {
	var res FileUploadSessionRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}