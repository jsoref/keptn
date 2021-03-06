// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/keptn/keptn/configuration-service/models"
)

// PutProjectProjectNameNoContentCode is the HTTP code returned for type PutProjectProjectNameNoContent
const PutProjectProjectNameNoContentCode int = 204

/*PutProjectProjectNameNoContent Success. Project has been updated. Response does not have a body.

swagger:response putProjectProjectNameNoContent
*/
type PutProjectProjectNameNoContent struct {
}

// NewPutProjectProjectNameNoContent creates PutProjectProjectNameNoContent with default headers values
func NewPutProjectProjectNameNoContent() *PutProjectProjectNameNoContent {

	return &PutProjectProjectNameNoContent{}
}

// WriteResponse to the client
func (o *PutProjectProjectNameNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// PutProjectProjectNameBadRequestCode is the HTTP code returned for type PutProjectProjectNameBadRequest
const PutProjectProjectNameBadRequestCode int = 400

/*PutProjectProjectNameBadRequest Failed. Project could not be updated.

swagger:response putProjectProjectNameBadRequest
*/
type PutProjectProjectNameBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPutProjectProjectNameBadRequest creates PutProjectProjectNameBadRequest with default headers values
func NewPutProjectProjectNameBadRequest() *PutProjectProjectNameBadRequest {

	return &PutProjectProjectNameBadRequest{}
}

// WithPayload adds the payload to the put project project name bad request response
func (o *PutProjectProjectNameBadRequest) WithPayload(payload *models.Error) *PutProjectProjectNameBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put project project name bad request response
func (o *PutProjectProjectNameBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutProjectProjectNameBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PutProjectProjectNameDefault Error

swagger:response putProjectProjectNameDefault
*/
type PutProjectProjectNameDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPutProjectProjectNameDefault creates PutProjectProjectNameDefault with default headers values
func NewPutProjectProjectNameDefault(code int) *PutProjectProjectNameDefault {
	if code <= 0 {
		code = 500
	}

	return &PutProjectProjectNameDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the put project project name default response
func (o *PutProjectProjectNameDefault) WithStatusCode(code int) *PutProjectProjectNameDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the put project project name default response
func (o *PutProjectProjectNameDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the put project project name default response
func (o *PutProjectProjectNameDefault) WithPayload(payload *models.Error) *PutProjectProjectNameDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put project project name default response
func (o *PutProjectProjectNameDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutProjectProjectNameDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
