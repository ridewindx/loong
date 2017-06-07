package access_token

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGenerateAccessTokenParams creates a new GenerateAccessTokenParams object
// with the default values initialized.
func NewGenerateAccessTokenParams() GenerateAccessTokenParams {
	var ()
	return GenerateAccessTokenParams{}
}

// GenerateAccessTokenParams contains all the bound params for the generate access token operation
// typically these are obtained from a http.Request
//
// swagger:parameters generateAccessToken
type GenerateAccessTokenParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*Full url path to the file that the token should be generated for example, https://api.box.com/2.0/files/{file_id}
	  In: formData
	*/
	ActorToken *string
	/*Full url path to the file that the token should be generated for example, https://api.box.com/2.0/files/{file_id}
	  In: formData
	*/
	ActorTokenType *string
	/*The client ID of the application requesting authentication.
	  Required: true
	  In: formData
	*/
	ClientID string
	/*The client secret of the application requesting authentication.
	  Required: true
	  In: formData
	*/
	ClientSecret string
	/*The authorization code returned by Box in response to a successfull authentication request.
	  In: formData
	*/
	Code *string
	/*
	  Required: true
	  In: formData
	*/
	GrantType string
	/*The refresh_token that is used to get the new access_token
	  In: formData
	*/
	RefreshToken *string
	/*Full url path to the file that the token should be generated for example, https://api.box.com/2.0/files/{file_id}
	  In: formData
	*/
	Resource *string
	/*This is the primary/secondary application token you will exchange for the file token.
	  In: formData
	*/
	SubjectToken *string
	/*The subject token type is urn:ietf:params:oauth:token-type:access_token
	  In: formData
	*/
	SubjectTokenType *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *GenerateAccessTokenParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	if err := r.ParseMultipartForm(32 << 20); err != nil {
		if err != http.ErrNotMultipart {
			return err
		} else if err := r.ParseForm(); err != nil {
			return err
		}
	}
	fds := runtime.Values(r.Form)

	fdActorToken, fdhkActorToken, _ := fds.GetOK("actor_token")
	if err := o.bindActorToken(fdActorToken, fdhkActorToken, route.Formats); err != nil {
		res = append(res, err)
	}

	fdActorTokenType, fdhkActorTokenType, _ := fds.GetOK("actor_token_type")
	if err := o.bindActorTokenType(fdActorTokenType, fdhkActorTokenType, route.Formats); err != nil {
		res = append(res, err)
	}

	fdClientID, fdhkClientID, _ := fds.GetOK("client_id")
	if err := o.bindClientID(fdClientID, fdhkClientID, route.Formats); err != nil {
		res = append(res, err)
	}

	fdClientSecret, fdhkClientSecret, _ := fds.GetOK("client_secret")
	if err := o.bindClientSecret(fdClientSecret, fdhkClientSecret, route.Formats); err != nil {
		res = append(res, err)
	}

	fdCode, fdhkCode, _ := fds.GetOK("code")
	if err := o.bindCode(fdCode, fdhkCode, route.Formats); err != nil {
		res = append(res, err)
	}

	fdGrantType, fdhkGrantType, _ := fds.GetOK("grant_type")
	if err := o.bindGrantType(fdGrantType, fdhkGrantType, route.Formats); err != nil {
		res = append(res, err)
	}

	fdRefreshToken, fdhkRefreshToken, _ := fds.GetOK("refresh_token")
	if err := o.bindRefreshToken(fdRefreshToken, fdhkRefreshToken, route.Formats); err != nil {
		res = append(res, err)
	}

	fdResource, fdhkResource, _ := fds.GetOK("resource")
	if err := o.bindResource(fdResource, fdhkResource, route.Formats); err != nil {
		res = append(res, err)
	}

	fdSubjectToken, fdhkSubjectToken, _ := fds.GetOK("subject_token")
	if err := o.bindSubjectToken(fdSubjectToken, fdhkSubjectToken, route.Formats); err != nil {
		res = append(res, err)
	}

	fdSubjectTokenType, fdhkSubjectTokenType, _ := fds.GetOK("subject_token_type")
	if err := o.bindSubjectTokenType(fdSubjectTokenType, fdhkSubjectTokenType, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GenerateAccessTokenParams) bindActorToken(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.ActorToken = &raw

	return nil
}

func (o *GenerateAccessTokenParams) bindActorTokenType(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.ActorTokenType = &raw

	if err := o.validateActorTokenType(formats); err != nil {
		return err
	}

	return nil
}

func (o *GenerateAccessTokenParams) validateActorTokenType(formats strfmt.Registry) error {

	if err := validate.Enum("actor_token_type", "formData", *o.ActorTokenType, []interface{}{"urn:ietf:params:oauth:token-type:id_token"}); err != nil {
		return err
	}

	return nil
}

func (o *GenerateAccessTokenParams) bindClientID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("client_id", "formData")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if err := validate.RequiredString("client_id", "formData", raw); err != nil {
		return err
	}

	o.ClientID = raw

	return nil
}

func (o *GenerateAccessTokenParams) bindClientSecret(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("client_secret", "formData")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if err := validate.RequiredString("client_secret", "formData", raw); err != nil {
		return err
	}

	o.ClientSecret = raw

	return nil
}

func (o *GenerateAccessTokenParams) bindCode(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Code = &raw

	return nil
}

func (o *GenerateAccessTokenParams) bindGrantType(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("grant_type", "formData")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if err := validate.RequiredString("grant_type", "formData", raw); err != nil {
		return err
	}

	o.GrantType = raw

	return nil
}

func (o *GenerateAccessTokenParams) bindRefreshToken(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.RefreshToken = &raw

	return nil
}

func (o *GenerateAccessTokenParams) bindResource(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Resource = &raw

	return nil
}

func (o *GenerateAccessTokenParams) bindSubjectToken(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.SubjectToken = &raw

	return nil
}

func (o *GenerateAccessTokenParams) bindSubjectTokenType(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.SubjectTokenType = &raw

	return nil
}