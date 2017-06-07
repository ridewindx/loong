package supercharged_upload

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// CreateUploadSessionHandlerFunc turns a function with the right signature into a create upload session handler
type CreateUploadSessionHandlerFunc func(CreateUploadSessionParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateUploadSessionHandlerFunc) Handle(params CreateUploadSessionParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// CreateUploadSessionHandler interface for that can handle valid create upload session params
type CreateUploadSessionHandler interface {
	Handle(CreateUploadSessionParams, interface{}) middleware.Responder
}

// NewCreateUploadSession creates a new http.Handler for the create upload session operation
func NewCreateUploadSession(ctx *middleware.Context, handler CreateUploadSessionHandler) *CreateUploadSession {
	return &CreateUploadSession{Context: ctx, Handler: handler}
}

/*CreateUploadSession swagger:route POST /files/upload_sessions SuperchargedUpload createUploadSession

Create File Upload Session

Use the Uploads API to create a new session to upload a new file.

*/
type CreateUploadSession struct {
	Context *middleware.Context
	Handler CreateUploadSessionHandler
}

func (o *CreateUploadSession) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewCreateUploadSessionParams()

	uprinc, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}