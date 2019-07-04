// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/artprocessors/nsm-microservice-golang-userdata/restapi/models"
)

// GetUserOKCode is the HTTP code returned for type GetUserOK
const GetUserOKCode int = 200

/*GetUserOK Ok

swagger:response getUserOK
*/
type GetUserOK struct {

	/*
	  In: Body
	*/
	Payload *models.User `json:"body,omitempty"`
}

// NewGetUserOK creates GetUserOK with default headers values
func NewGetUserOK() *GetUserOK {

	return &GetUserOK{}
}

// WithPayload adds the payload to the get user o k response
func (o *GetUserOK) WithPayload(payload *models.User) *GetUserOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user o k response
func (o *GetUserOK) SetPayload(payload *models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUserBadRequestCode is the HTTP code returned for type GetUserBadRequest
const GetUserBadRequestCode int = 400

/*GetUserBadRequest Invalid request

swagger:response getUserBadRequest
*/
type GetUserBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewGetUserBadRequest creates GetUserBadRequest with default headers values
func NewGetUserBadRequest() *GetUserBadRequest {

	return &GetUserBadRequest{}
}

// WithPayload adds the payload to the get user bad request response
func (o *GetUserBadRequest) WithPayload(payload *models.APIError) *GetUserBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user bad request response
func (o *GetUserBadRequest) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUserNotFoundCode is the HTTP code returned for type GetUserNotFound
const GetUserNotFoundCode int = 404

/*GetUserNotFound Not found

swagger:response getUserNotFound
*/
type GetUserNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewGetUserNotFound creates GetUserNotFound with default headers values
func NewGetUserNotFound() *GetUserNotFound {

	return &GetUserNotFound{}
}

// WithPayload adds the payload to the get user not found response
func (o *GetUserNotFound) WithPayload(payload *models.APIError) *GetUserNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user not found response
func (o *GetUserNotFound) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUserInternalServerErrorCode is the HTTP code returned for type GetUserInternalServerError
const GetUserInternalServerErrorCode int = 500

/*GetUserInternalServerError Internal server error

swagger:response getUserInternalServerError
*/
type GetUserInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewGetUserInternalServerError creates GetUserInternalServerError with default headers values
func NewGetUserInternalServerError() *GetUserInternalServerError {

	return &GetUserInternalServerError{}
}

// WithPayload adds the payload to the get user internal server error response
func (o *GetUserInternalServerError) WithPayload(payload *models.APIError) *GetUserInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user internal server error response
func (o *GetUserInternalServerError) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
