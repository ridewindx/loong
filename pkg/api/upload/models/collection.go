package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Collection Collection Object
//
// Collections contain information about the items contained inside of them, including files and folders. The only collection available currently is a “Favorites” collection. The contents of the collection are discovered in a similar way in which the contents of a folder are discovered.
// swagger:model Collection
type Collection struct {
	Reference

	// The type of the collection. This is used to determine the proper visual treatment for Box-internally created collections. Initially only “favorites” collection-type will be supported.
	CollectionType string `json:"collection_type,omitempty"`

	// The name of this collection. The only collection currently available is named “Favorites”
	Name string `json:"name,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *Collection) UnmarshalJSON(raw []byte) error {

	var aO0 Reference
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Reference = aO0

	var data struct {
		CollectionType string `json:"collection_type,omitempty"`

		Name string `json:"name,omitempty"`
	}
	if err := swag.ReadJSON(raw, &data); err != nil {
		return err
	}

	m.CollectionType = data.CollectionType

	m.Name = data.Name

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m Collection) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	aO0, err := swag.WriteJSON(m.Reference)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var data struct {
		CollectionType string `json:"collection_type,omitempty"`

		Name string `json:"name,omitempty"`
	}

	data.CollectionType = m.CollectionType

	data.Name = m.Name

	jsonData, err := swag.WriteJSON(data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, jsonData)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this collection
func (m *Collection) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.Reference.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCollectionType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var collectionTypeCollectionTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["favorites"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		collectionTypeCollectionTypePropEnum = append(collectionTypeCollectionTypePropEnum, v)
	}
}

// property enum
func (m *Collection) validateCollectionTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, collectionTypeCollectionTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Collection) validateCollectionType(formats strfmt.Registry) error {

	if swag.IsZero(m.CollectionType) { // not required
		return nil
	}

	// value enum
	if err := m.validateCollectionTypeEnum("collection_type", "body", m.CollectionType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Collection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Collection) UnmarshalBinary(b []byte) error {
	var res Collection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}