package supercharged_upload

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	"github.com/ridewindx/loong/pkg/api/upload/models"
)

// CommitUploadSessionCreatedCode is the HTTP code returned for type CommitUploadSessionCreated
const CommitUploadSessionCreatedCode int = 201

/*CommitUploadSessionCreated The file has been committed successfully

swagger:response commitUploadSessionCreated
*/
type CommitUploadSessionCreated struct {

	/*
	  In: Body
	*/
	Payload *models.FileList `json:"body,omitempty"`
}

// NewCommitUploadSessionCreated creates CommitUploadSessionCreated with default headers values
func NewCommitUploadSessionCreated() *CommitUploadSessionCreated {
	return &CommitUploadSessionCreated{}
}

// WithPayload adds the payload to the commit upload session created response
func (o *CommitUploadSessionCreated) WithPayload(payload *models.FileList) *CommitUploadSessionCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the commit upload session created response
func (o *CommitUploadSessionCreated) SetPayload(payload *models.FileList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CommitUploadSessionCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CommitUploadSessionAcceptedCode is the HTTP code returned for type CommitUploadSessionAccepted
const CommitUploadSessionAcceptedCode int = 202

/*CommitUploadSessionAccepted All parts have been uploaded but not yet processed. Use Get Upload Session API to get more information about the progress of processing the parts. Retry commit afterwards

swagger:response commitUploadSessionAccepted
*/
type CommitUploadSessionAccepted struct {
	/*Indicates the number of seconds the client should wait before attempting their commit request again.
	  Required: true
	*/
	RetryAfter int64 `json:"Retry-After"`
}

// NewCommitUploadSessionAccepted creates CommitUploadSessionAccepted with default headers values
func NewCommitUploadSessionAccepted() *CommitUploadSessionAccepted {
	return &CommitUploadSessionAccepted{}
}

// WithRetryAfter adds the retryAfter to the commit upload session accepted response
func (o *CommitUploadSessionAccepted) WithRetryAfter(retryAfter int64) *CommitUploadSessionAccepted {
	o.RetryAfter = retryAfter
	return o
}

// SetRetryAfter sets the retryAfter to the commit upload session accepted response
func (o *CommitUploadSessionAccepted) SetRetryAfter(retryAfter int64) {
	o.RetryAfter = retryAfter
}

// WriteResponse to the client
func (o *CommitUploadSessionAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Retry-After

	retryAfter := swag.FormatInt64(o.RetryAfter)
	if retryAfter != "" {
		rw.Header().Set("Retry-After", retryAfter)
	}

	rw.WriteHeader(202)
}

// CommitUploadSessionBadRequestCode is the HTTP code returned for type CommitUploadSessionBadRequest
const CommitUploadSessionBadRequestCode int = 400

/*CommitUploadSessionBadRequest Bad request. See response body for details.

swagger:response commitUploadSessionBadRequest
*/
type CommitUploadSessionBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCommitUploadSessionBadRequest creates CommitUploadSessionBadRequest with default headers values
func NewCommitUploadSessionBadRequest() *CommitUploadSessionBadRequest {
	return &CommitUploadSessionBadRequest{}
}

// WithPayload adds the payload to the commit upload session bad request response
func (o *CommitUploadSessionBadRequest) WithPayload(payload *models.Error) *CommitUploadSessionBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the commit upload session bad request response
func (o *CommitUploadSessionBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CommitUploadSessionBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CommitUploadSessionConflictCode is the HTTP code returned for type CommitUploadSessionConflict
const CommitUploadSessionConflictCode int = 409

/*CommitUploadSessionConflict There is already a file with the same name in the target folder.

swagger:response commitUploadSessionConflict
*/
type CommitUploadSessionConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCommitUploadSessionConflict creates CommitUploadSessionConflict with default headers values
func NewCommitUploadSessionConflict() *CommitUploadSessionConflict {
	return &CommitUploadSessionConflict{}
}

// WithPayload adds the payload to the commit upload session conflict response
func (o *CommitUploadSessionConflict) WithPayload(payload *models.Error) *CommitUploadSessionConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the commit upload session conflict response
func (o *CommitUploadSessionConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CommitUploadSessionConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CommitUploadSessionPreconditionFailedCode is the HTTP code returned for type CommitUploadSessionPreconditionFailed
const CommitUploadSessionPreconditionFailedCode int = 412

/*CommitUploadSessionPreconditionFailed The If-Match condition failed.

swagger:response commitUploadSessionPreconditionFailed
*/
type CommitUploadSessionPreconditionFailed struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCommitUploadSessionPreconditionFailed creates CommitUploadSessionPreconditionFailed with default headers values
func NewCommitUploadSessionPreconditionFailed() *CommitUploadSessionPreconditionFailed {
	return &CommitUploadSessionPreconditionFailed{}
}

// WithPayload adds the payload to the commit upload session precondition failed response
func (o *CommitUploadSessionPreconditionFailed) WithPayload(payload *models.Error) *CommitUploadSessionPreconditionFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the commit upload session precondition failed response
func (o *CommitUploadSessionPreconditionFailed) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CommitUploadSessionPreconditionFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(412)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CommitUploadSessionRequestEntityTooLargeCode is the HTTP code returned for type CommitUploadSessionRequestEntityTooLarge
const CommitUploadSessionRequestEntityTooLargeCode int = 413

/*CommitUploadSessionRequestEntityTooLarge Request entity too large. See response body for details.

swagger:response commitUploadSessionRequestEntityTooLarge
*/
type CommitUploadSessionRequestEntityTooLarge struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCommitUploadSessionRequestEntityTooLarge creates CommitUploadSessionRequestEntityTooLarge with default headers values
func NewCommitUploadSessionRequestEntityTooLarge() *CommitUploadSessionRequestEntityTooLarge {
	return &CommitUploadSessionRequestEntityTooLarge{}
}

// WithPayload adds the payload to the commit upload session request entity too large response
func (o *CommitUploadSessionRequestEntityTooLarge) WithPayload(payload *models.Error) *CommitUploadSessionRequestEntityTooLarge {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the commit upload session request entity too large response
func (o *CommitUploadSessionRequestEntityTooLarge) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CommitUploadSessionRequestEntityTooLarge) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(413)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*CommitUploadSessionDefault Part upload error

swagger:response commitUploadSessionDefault
*/
type CommitUploadSessionDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCommitUploadSessionDefault creates CommitUploadSessionDefault with default headers values
func NewCommitUploadSessionDefault(code int) *CommitUploadSessionDefault {
	if code <= 0 {
		code = 500
	}

	return &CommitUploadSessionDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the commit upload session default response
func (o *CommitUploadSessionDefault) WithStatusCode(code int) *CommitUploadSessionDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the commit upload session default response
func (o *CommitUploadSessionDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the commit upload session default response
func (o *CommitUploadSessionDefault) WithPayload(payload *models.Error) *CommitUploadSessionDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the commit upload session default response
func (o *CommitUploadSessionDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CommitUploadSessionDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
