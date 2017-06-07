package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AccessToken Access token response
// swagger:model AccessToken
type AccessToken struct {

	// access token
	AccessToken string `json:"access_token,omitempty"`

	// expires in
	ExpiresIn int64 `json:"expires_in,omitempty"`

	// issued token type
	IssuedTokenType string `json:"issued_token_type,omitempty"`

	// refresh token
	RefreshToken string `json:"refresh_token,omitempty"`

	// restricted to
	RestrictedTo []*RestrictedTo `json:"restricted_to"`

	// token type
	TokenType string `json:"token_type,omitempty"`
}

// Validate validates this access token
func (m *AccessToken) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRestrictedTo(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTokenType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccessToken) validateRestrictedTo(formats strfmt.Registry) error {

	if swag.IsZero(m.RestrictedTo) { // not required
		return nil
	}

	for i := 0; i < len(m.RestrictedTo); i++ {

		if swag.IsZero(m.RestrictedTo[i]) { // not required
			continue
		}

		if m.RestrictedTo[i] != nil {

			if err := m.RestrictedTo[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("restricted_to" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var accessTokenTypeTokenTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["bearer"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accessTokenTypeTokenTypePropEnum = append(accessTokenTypeTokenTypePropEnum, v)
	}
}

const (
	// AccessTokenTokenTypeBearer captures enum value "bearer"
	AccessTokenTokenTypeBearer string = "bearer"
)

// prop value enum
func (m *AccessToken) validateTokenTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, accessTokenTypeTokenTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AccessToken) validateTokenType(formats strfmt.Registry) error {

	if swag.IsZero(m.TokenType) { // not required
		return nil
	}

	// value enum
	if err := m.validateTokenTypeEnum("token_type", "body", m.TokenType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccessToken) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccessToken) UnmarshalBinary(b []byte) error {
	var res AccessToken
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}