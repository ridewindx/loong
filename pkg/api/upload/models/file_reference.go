package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// FileReference It references either a file or a folder
// swagger:model FileReference
type FileReference struct {
	ItemReference

	// The sha1 hash of this file.
	Sha1 string `json:"sha1,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *FileReference) UnmarshalJSON(raw []byte) error {

	var aO0 ItemReference
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ItemReference = aO0

	var data struct {
		Sha1 string `json:"sha1,omitempty"`
	}
	if err := swag.ReadJSON(raw, &data); err != nil {
		return err
	}

	m.Sha1 = data.Sha1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m FileReference) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	aO0, err := swag.WriteJSON(m.ItemReference)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var data struct {
		Sha1 string `json:"sha1,omitempty"`
	}

	data.Sha1 = m.Sha1

	jsonData, err := swag.WriteJSON(data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, jsonData)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this file reference
func (m *FileReference) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.ItemReference.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *FileReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FileReference) UnmarshalBinary(b []byte) error {
	var res FileReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
