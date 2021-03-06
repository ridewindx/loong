package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Part part
// swagger:model Part
type Part struct {

	// offset
	Offset int64 `json:"offset,omitempty"`

	// part id
	PartID string `json:"part_id,omitempty"`

	// sha1
	Sha1 string `json:"sha1,omitempty"`

	// size
	Size int64 `json:"size,omitempty"`
}

// Validate validates this part
func (m *Part) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Part) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Part) UnmarshalBinary(b []byte) error {
	var res Part
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
