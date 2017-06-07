package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// UploadSessionEndpoints URLs for all other possible calls to this session.
// swagger:model UploadSessionEndpoints
type UploadSessionEndpoints struct {

	// The URL for abort API.
	Abort string `json:"abort,omitempty"`

	// The URL for commit API.
	Commit string `json:"commit,omitempty"`

	// The URL for list parts API.
	ListParts string `json:"list_parts,omitempty"`

	// The URL for log event API.
	LogEvent string `json:"log_event,omitempty"`

	// The URL for status API.
	Status string `json:"status,omitempty"`

	// The URL for upload part API.
	UploadPart string `json:"upload_part,omitempty"`
}

// Validate validates this upload session endpoints
func (m *UploadSessionEndpoints) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *UploadSessionEndpoints) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UploadSessionEndpoints) UnmarshalBinary(b []byte) error {
	var res UploadSessionEndpoints
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
