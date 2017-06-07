package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Pagination pagination
// swagger:model Pagination
type Pagination struct {

	// limit
	Limit int64 `json:"limit,omitempty"`

	// offset
	Offset int64 `json:"offset,omitempty"`

	// order
	Order []*PaginationOrderItems0 `json:"order"`

	// total count
	TotalCount int64 `json:"total_count,omitempty"`
}

// Validate validates this pagination
func (m *Pagination) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrder(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Pagination) validateOrder(formats strfmt.Registry) error {

	if swag.IsZero(m.Order) { // not required
		return nil
	}

	for i := 0; i < len(m.Order); i++ {

		if swag.IsZero(m.Order[i]) { // not required
			continue
		}

		if m.Order[i] != nil {

			if err := m.Order[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("order" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Pagination) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Pagination) UnmarshalBinary(b []byte) error {
	var res Pagination
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PaginationOrderItems0 pagination order items0
// swagger:model PaginationOrderItems0
type PaginationOrderItems0 struct {

	// by
	By string `json:"by,omitempty"`

	// direction
	Direction string `json:"direction,omitempty"`
}

// Validate validates this pagination order items0
func (m *PaginationOrderItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PaginationOrderItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaginationOrderItems0) UnmarshalBinary(b []byte) error {
	var res PaginationOrderItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}