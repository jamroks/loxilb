// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/loxilb-io/loxilb/api/models"
)

// DeleteConfigBfdRemoteIPNoContentCode is the HTTP code returned for type DeleteConfigBfdRemoteIPNoContent
const DeleteConfigBfdRemoteIPNoContentCode int = 204

/*
DeleteConfigBfdRemoteIPNoContent OK

swagger:response deleteConfigBfdRemoteIpNoContent
*/
type DeleteConfigBfdRemoteIPNoContent struct {
}

// NewDeleteConfigBfdRemoteIPNoContent creates DeleteConfigBfdRemoteIPNoContent with default headers values
func NewDeleteConfigBfdRemoteIPNoContent() *DeleteConfigBfdRemoteIPNoContent {

	return &DeleteConfigBfdRemoteIPNoContent{}
}

// WriteResponse to the client
func (o *DeleteConfigBfdRemoteIPNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteConfigBfdRemoteIPBadRequestCode is the HTTP code returned for type DeleteConfigBfdRemoteIPBadRequest
const DeleteConfigBfdRemoteIPBadRequestCode int = 400

/*
DeleteConfigBfdRemoteIPBadRequest Malformed arguments for API call

swagger:response deleteConfigBfdRemoteIpBadRequest
*/
type DeleteConfigBfdRemoteIPBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigBfdRemoteIPBadRequest creates DeleteConfigBfdRemoteIPBadRequest with default headers values
func NewDeleteConfigBfdRemoteIPBadRequest() *DeleteConfigBfdRemoteIPBadRequest {

	return &DeleteConfigBfdRemoteIPBadRequest{}
}

// WithPayload adds the payload to the delete config bfd remote Ip bad request response
func (o *DeleteConfigBfdRemoteIPBadRequest) WithPayload(payload *models.Error) *DeleteConfigBfdRemoteIPBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config bfd remote Ip bad request response
func (o *DeleteConfigBfdRemoteIPBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigBfdRemoteIPBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteConfigBfdRemoteIPUnauthorizedCode is the HTTP code returned for type DeleteConfigBfdRemoteIPUnauthorized
const DeleteConfigBfdRemoteIPUnauthorizedCode int = 401

/*
DeleteConfigBfdRemoteIPUnauthorized Invalid authentication credentials

swagger:response deleteConfigBfdRemoteIpUnauthorized
*/
type DeleteConfigBfdRemoteIPUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigBfdRemoteIPUnauthorized creates DeleteConfigBfdRemoteIPUnauthorized with default headers values
func NewDeleteConfigBfdRemoteIPUnauthorized() *DeleteConfigBfdRemoteIPUnauthorized {

	return &DeleteConfigBfdRemoteIPUnauthorized{}
}

// WithPayload adds the payload to the delete config bfd remote Ip unauthorized response
func (o *DeleteConfigBfdRemoteIPUnauthorized) WithPayload(payload *models.Error) *DeleteConfigBfdRemoteIPUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config bfd remote Ip unauthorized response
func (o *DeleteConfigBfdRemoteIPUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigBfdRemoteIPUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteConfigBfdRemoteIPForbiddenCode is the HTTP code returned for type DeleteConfigBfdRemoteIPForbidden
const DeleteConfigBfdRemoteIPForbiddenCode int = 403

/*
DeleteConfigBfdRemoteIPForbidden Capacity insufficient

swagger:response deleteConfigBfdRemoteIpForbidden
*/
type DeleteConfigBfdRemoteIPForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigBfdRemoteIPForbidden creates DeleteConfigBfdRemoteIPForbidden with default headers values
func NewDeleteConfigBfdRemoteIPForbidden() *DeleteConfigBfdRemoteIPForbidden {

	return &DeleteConfigBfdRemoteIPForbidden{}
}

// WithPayload adds the payload to the delete config bfd remote Ip forbidden response
func (o *DeleteConfigBfdRemoteIPForbidden) WithPayload(payload *models.Error) *DeleteConfigBfdRemoteIPForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config bfd remote Ip forbidden response
func (o *DeleteConfigBfdRemoteIPForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigBfdRemoteIPForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteConfigBfdRemoteIPNotFoundCode is the HTTP code returned for type DeleteConfigBfdRemoteIPNotFound
const DeleteConfigBfdRemoteIPNotFoundCode int = 404

/*
DeleteConfigBfdRemoteIPNotFound Resource not found

swagger:response deleteConfigBfdRemoteIpNotFound
*/
type DeleteConfigBfdRemoteIPNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigBfdRemoteIPNotFound creates DeleteConfigBfdRemoteIPNotFound with default headers values
func NewDeleteConfigBfdRemoteIPNotFound() *DeleteConfigBfdRemoteIPNotFound {

	return &DeleteConfigBfdRemoteIPNotFound{}
}

// WithPayload adds the payload to the delete config bfd remote Ip not found response
func (o *DeleteConfigBfdRemoteIPNotFound) WithPayload(payload *models.Error) *DeleteConfigBfdRemoteIPNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config bfd remote Ip not found response
func (o *DeleteConfigBfdRemoteIPNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigBfdRemoteIPNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteConfigBfdRemoteIPConflictCode is the HTTP code returned for type DeleteConfigBfdRemoteIPConflict
const DeleteConfigBfdRemoteIPConflictCode int = 409

/*
DeleteConfigBfdRemoteIPConflict Resource Conflict. BFD session already exists

swagger:response deleteConfigBfdRemoteIpConflict
*/
type DeleteConfigBfdRemoteIPConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigBfdRemoteIPConflict creates DeleteConfigBfdRemoteIPConflict with default headers values
func NewDeleteConfigBfdRemoteIPConflict() *DeleteConfigBfdRemoteIPConflict {

	return &DeleteConfigBfdRemoteIPConflict{}
}

// WithPayload adds the payload to the delete config bfd remote Ip conflict response
func (o *DeleteConfigBfdRemoteIPConflict) WithPayload(payload *models.Error) *DeleteConfigBfdRemoteIPConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config bfd remote Ip conflict response
func (o *DeleteConfigBfdRemoteIPConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigBfdRemoteIPConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteConfigBfdRemoteIPInternalServerErrorCode is the HTTP code returned for type DeleteConfigBfdRemoteIPInternalServerError
const DeleteConfigBfdRemoteIPInternalServerErrorCode int = 500

/*
DeleteConfigBfdRemoteIPInternalServerError Internal service error

swagger:response deleteConfigBfdRemoteIpInternalServerError
*/
type DeleteConfigBfdRemoteIPInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigBfdRemoteIPInternalServerError creates DeleteConfigBfdRemoteIPInternalServerError with default headers values
func NewDeleteConfigBfdRemoteIPInternalServerError() *DeleteConfigBfdRemoteIPInternalServerError {

	return &DeleteConfigBfdRemoteIPInternalServerError{}
}

// WithPayload adds the payload to the delete config bfd remote Ip internal server error response
func (o *DeleteConfigBfdRemoteIPInternalServerError) WithPayload(payload *models.Error) *DeleteConfigBfdRemoteIPInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config bfd remote Ip internal server error response
func (o *DeleteConfigBfdRemoteIPInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigBfdRemoteIPInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteConfigBfdRemoteIPServiceUnavailableCode is the HTTP code returned for type DeleteConfigBfdRemoteIPServiceUnavailable
const DeleteConfigBfdRemoteIPServiceUnavailableCode int = 503

/*
DeleteConfigBfdRemoteIPServiceUnavailable Maintanence mode

swagger:response deleteConfigBfdRemoteIpServiceUnavailable
*/
type DeleteConfigBfdRemoteIPServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigBfdRemoteIPServiceUnavailable creates DeleteConfigBfdRemoteIPServiceUnavailable with default headers values
func NewDeleteConfigBfdRemoteIPServiceUnavailable() *DeleteConfigBfdRemoteIPServiceUnavailable {

	return &DeleteConfigBfdRemoteIPServiceUnavailable{}
}

// WithPayload adds the payload to the delete config bfd remote Ip service unavailable response
func (o *DeleteConfigBfdRemoteIPServiceUnavailable) WithPayload(payload *models.Error) *DeleteConfigBfdRemoteIPServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config bfd remote Ip service unavailable response
func (o *DeleteConfigBfdRemoteIPServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigBfdRemoteIPServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
