// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/stephram/nsm-userdata-service/restapi/models"
)

// PostUserOKCode is the HTTP code returned for type PostUserOK
const PostUserOKCode int = 200

/*PostUserOK Ok

swagger:response postUserOK
*/
type PostUserOK struct {

	/*
	  In: Body
	*/
	Payload *models.User `json:"body,omitempty"`
}

// NewPostUserOK creates PostUserOK with default headers values
func NewPostUserOK() *PostUserOK {

	return &PostUserOK{}
}

// WithPayload adds the payload to the post user o k response
func (o *PostUserOK) WithPayload(payload *models.User) *PostUserOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post user o k response
func (o *PostUserOK) SetPayload(payload *models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostUserCreatedCode is the HTTP code returned for type PostUserCreated
const PostUserCreatedCode int = 201

/*PostUserCreated Created

swagger:response postUserCreated
*/
type PostUserCreated struct {

	/*
	  In: Body
	*/
	Payload *models.User `json:"body,omitempty"`
}

// NewPostUserCreated creates PostUserCreated with default headers values
func NewPostUserCreated() *PostUserCreated {

	return &PostUserCreated{}
}

// WithPayload adds the payload to the post user created response
func (o *PostUserCreated) WithPayload(payload *models.User) *PostUserCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post user created response
func (o *PostUserCreated) SetPayload(payload *models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostUserCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostUserBadRequestCode is the HTTP code returned for type PostUserBadRequest
const PostUserBadRequestCode int = 400

/*PostUserBadRequest Invalid request

swagger:response postUserBadRequest
*/
type PostUserBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewPostUserBadRequest creates PostUserBadRequest with default headers values
func NewPostUserBadRequest() *PostUserBadRequest {

	return &PostUserBadRequest{}
}

// WithPayload adds the payload to the post user bad request response
func (o *PostUserBadRequest) WithPayload(payload *models.APIError) *PostUserBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post user bad request response
func (o *PostUserBadRequest) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostUserBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostUserNotFoundCode is the HTTP code returned for type PostUserNotFound
const PostUserNotFoundCode int = 404

/*PostUserNotFound Not found

swagger:response postUserNotFound
*/
type PostUserNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewPostUserNotFound creates PostUserNotFound with default headers values
func NewPostUserNotFound() *PostUserNotFound {

	return &PostUserNotFound{}
}

// WithPayload adds the payload to the post user not found response
func (o *PostUserNotFound) WithPayload(payload *models.APIError) *PostUserNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post user not found response
func (o *PostUserNotFound) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostUserNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostUserInternalServerErrorCode is the HTTP code returned for type PostUserInternalServerError
const PostUserInternalServerErrorCode int = 500

/*PostUserInternalServerError Internal server error

swagger:response postUserInternalServerError
*/
type PostUserInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewPostUserInternalServerError creates PostUserInternalServerError with default headers values
func NewPostUserInternalServerError() *PostUserInternalServerError {

	return &PostUserInternalServerError{}
}

// WithPayload adds the payload to the post user internal server error response
func (o *PostUserInternalServerError) WithPayload(payload *models.APIError) *PostUserInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post user internal server error response
func (o *PostUserInternalServerError) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostUserInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
