package supercharged_upload

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetUploadSessionParams creates a new GetUploadSessionParams object
// with the default values initialized.
func NewGetUploadSessionParams() GetUploadSessionParams {
	var ()
	return GetUploadSessionParams{}
}

// GetUploadSessionParams contains all the bound params for the get upload session operation
// typically these are obtained from a http.Request
//
// swagger:parameters getUploadSession
type GetUploadSessionParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*
	  Required: true
	  In: path
	*/
	SESSIONID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *GetUploadSessionParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	rSESSIONID, rhkSESSIONID, _ := route.Params.GetOK("SESSION_ID")
	if err := o.bindSESSIONID(rSESSIONID, rhkSESSIONID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetUploadSessionParams) bindSESSIONID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.SESSIONID = raw

	return nil
}
