package supercharged_upload

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// CommitUploadSessionHandlerFunc turns a function with the right signature into a commit upload session handler
type CommitUploadSessionHandlerFunc func(CommitUploadSessionParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn CommitUploadSessionHandlerFunc) Handle(params CommitUploadSessionParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// CommitUploadSessionHandler interface for that can handle valid commit upload session params
type CommitUploadSessionHandler interface {
	Handle(CommitUploadSessionParams, interface{}) middleware.Responder
}

// NewCommitUploadSession creates a new http.Handler for the commit upload session operation
func NewCommitUploadSession(ctx *middleware.Context, handler CommitUploadSessionHandler) *CommitUploadSession {
	return &CommitUploadSession{Context: ctx, Handler: handler}
}

/*CommitUploadSession swagger:route POST /files/upload_sessions/{SESSION_ID}/commit SuperchargedUpload commitUploadSession

Commit Upload

Create a Box file comprised of the uploaded parts.

*/
type CommitUploadSession struct {
	Context *middleware.Context
	Handler CommitUploadSessionHandler
}

func (o *CommitUploadSession) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewCommitUploadSessionParams()

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
