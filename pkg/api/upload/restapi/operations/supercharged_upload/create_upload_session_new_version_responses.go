package supercharged_upload

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/ridewindx/loong/pkg/api/upload/models"
)

// CreateUploadSessionNewVersionCreatedCode is the HTTP code returned for type CreateUploadSessionNewVersionCreated
const CreateUploadSessionNewVersionCreatedCode int = 201

/*CreateUploadSessionNewVersionCreated The upload session creation would be successful

swagger:response createUploadSessionNewVersionCreated
*/
type CreateUploadSessionNewVersionCreated struct {

	/*
	  In: Body
	*/
	Payload *models.UploadSession `json:"body,omitempty"`
}

// NewCreateUploadSessionNewVersionCreated creates CreateUploadSessionNewVersionCreated with default headers values
func NewCreateUploadSessionNewVersionCreated() *CreateUploadSessionNewVersionCreated {
	return &CreateUploadSessionNewVersionCreated{}
}

// WithPayload adds the payload to the create upload session new version created response
func (o *CreateUploadSessionNewVersionCreated) WithPayload(payload *models.UploadSession) *CreateUploadSessionNewVersionCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create upload session new version created response
func (o *CreateUploadSessionNewVersionCreated) SetPayload(payload *models.UploadSession) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateUploadSessionNewVersionCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateUploadSessionNewVersionBadRequestCode is the HTTP code returned for type CreateUploadSessionNewVersionBadRequest
const CreateUploadSessionNewVersionBadRequestCode int = 400

/*CreateUploadSessionNewVersionBadRequest Bad request. See response body for details.

swagger:response createUploadSessionNewVersionBadRequest
*/
type CreateUploadSessionNewVersionBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateUploadSessionNewVersionBadRequest creates CreateUploadSessionNewVersionBadRequest with default headers values
func NewCreateUploadSessionNewVersionBadRequest() *CreateUploadSessionNewVersionBadRequest {
	return &CreateUploadSessionNewVersionBadRequest{}
}

// WithPayload adds the payload to the create upload session new version bad request response
func (o *CreateUploadSessionNewVersionBadRequest) WithPayload(payload *models.Error) *CreateUploadSessionNewVersionBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create upload session new version bad request response
func (o *CreateUploadSessionNewVersionBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateUploadSessionNewVersionBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateUploadSessionNewVersionConflictCode is the HTTP code returned for type CreateUploadSessionNewVersionConflict
const CreateUploadSessionNewVersionConflictCode int = 409

/*CreateUploadSessionNewVersionConflict Name collision occurs. See response body for details.

swagger:response createUploadSessionNewVersionConflict
*/
type CreateUploadSessionNewVersionConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateUploadSessionNewVersionConflict creates CreateUploadSessionNewVersionConflict with default headers values
func NewCreateUploadSessionNewVersionConflict() *CreateUploadSessionNewVersionConflict {
	return &CreateUploadSessionNewVersionConflict{}
}

// WithPayload adds the payload to the create upload session new version conflict response
func (o *CreateUploadSessionNewVersionConflict) WithPayload(payload *models.Error) *CreateUploadSessionNewVersionConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create upload session new version conflict response
func (o *CreateUploadSessionNewVersionConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateUploadSessionNewVersionConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateUploadSessionNewVersionPreconditionFailedCode is the HTTP code returned for type CreateUploadSessionNewVersionPreconditionFailed
const CreateUploadSessionNewVersionPreconditionFailedCode int = 412

/*CreateUploadSessionNewVersionPreconditionFailed Preconditioned failed. See response body for details.

swagger:response createUploadSessionNewVersionPreconditionFailed
*/
type CreateUploadSessionNewVersionPreconditionFailed struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateUploadSessionNewVersionPreconditionFailed creates CreateUploadSessionNewVersionPreconditionFailed with default headers values
func NewCreateUploadSessionNewVersionPreconditionFailed() *CreateUploadSessionNewVersionPreconditionFailed {
	return &CreateUploadSessionNewVersionPreconditionFailed{}
}

// WithPayload adds the payload to the create upload session new version precondition failed response
func (o *CreateUploadSessionNewVersionPreconditionFailed) WithPayload(payload *models.Error) *CreateUploadSessionNewVersionPreconditionFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create upload session new version precondition failed response
func (o *CreateUploadSessionNewVersionPreconditionFailed) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateUploadSessionNewVersionPreconditionFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(412)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateUploadSessionNewVersionRequestEntityTooLargeCode is the HTTP code returned for type CreateUploadSessionNewVersionRequestEntityTooLarge
const CreateUploadSessionNewVersionRequestEntityTooLargeCode int = 413

/*CreateUploadSessionNewVersionRequestEntityTooLarge Request entity too large. See response body for details.

swagger:response createUploadSessionNewVersionRequestEntityTooLarge
*/
type CreateUploadSessionNewVersionRequestEntityTooLarge struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateUploadSessionNewVersionRequestEntityTooLarge creates CreateUploadSessionNewVersionRequestEntityTooLarge with default headers values
func NewCreateUploadSessionNewVersionRequestEntityTooLarge() *CreateUploadSessionNewVersionRequestEntityTooLarge {
	return &CreateUploadSessionNewVersionRequestEntityTooLarge{}
}

// WithPayload adds the payload to the create upload session new version request entity too large response
func (o *CreateUploadSessionNewVersionRequestEntityTooLarge) WithPayload(payload *models.Error) *CreateUploadSessionNewVersionRequestEntityTooLarge {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create upload session new version request entity too large response
func (o *CreateUploadSessionNewVersionRequestEntityTooLarge) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateUploadSessionNewVersionRequestEntityTooLarge) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(413)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*CreateUploadSessionNewVersionDefault File upload error

swagger:response createUploadSessionNewVersionDefault
*/
type CreateUploadSessionNewVersionDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateUploadSessionNewVersionDefault creates CreateUploadSessionNewVersionDefault with default headers values
func NewCreateUploadSessionNewVersionDefault(code int) *CreateUploadSessionNewVersionDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateUploadSessionNewVersionDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create upload session new version default response
func (o *CreateUploadSessionNewVersionDefault) WithStatusCode(code int) *CreateUploadSessionNewVersionDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create upload session new version default response
func (o *CreateUploadSessionNewVersionDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the create upload session new version default response
func (o *CreateUploadSessionNewVersionDefault) WithPayload(payload *models.Error) *CreateUploadSessionNewVersionDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create upload session new version default response
func (o *CreateUploadSessionNewVersionDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateUploadSessionNewVersionDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
