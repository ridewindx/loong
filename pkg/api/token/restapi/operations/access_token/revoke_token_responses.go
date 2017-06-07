package access_token

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/ridewindx/loong/pkg/api/token/models"
)

// RevokeTokenOKCode is the HTTP code returned for type RevokeTokenOK
const RevokeTokenOKCode int = 200

/*RevokeTokenOK Revoken token is successful

swagger:response revokeTokenOK
*/
type RevokeTokenOK struct {
}

// NewRevokeTokenOK creates RevokeTokenOK with default headers values
func NewRevokeTokenOK() *RevokeTokenOK {
	return &RevokeTokenOK{}
}

// WriteResponse to the client
func (o *RevokeTokenOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
}

/*RevokeTokenDefault Revoke token error

swagger:response revokeTokenDefault
*/
type RevokeTokenDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.OAuthError `json:"body,omitempty"`
}

// NewRevokeTokenDefault creates RevokeTokenDefault with default headers values
func NewRevokeTokenDefault(code int) *RevokeTokenDefault {
	if code <= 0 {
		code = 500
	}

	return &RevokeTokenDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the revoke token default response
func (o *RevokeTokenDefault) WithStatusCode(code int) *RevokeTokenDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the revoke token default response
func (o *RevokeTokenDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the revoke token default response
func (o *RevokeTokenDefault) WithPayload(payload *models.OAuthError) *RevokeTokenDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the revoke token default response
func (o *RevokeTokenDefault) SetPayload(payload *models.OAuthError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RevokeTokenDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}