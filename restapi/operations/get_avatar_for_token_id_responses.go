// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/artprocessors/nsm-microservice-golang-userdata/restapi/models"
)

// GetAvatarForTokenIDOKCode is the HTTP code returned for type GetAvatarForTokenIDOK
const GetAvatarForTokenIDOKCode int = 200

/*GetAvatarForTokenIDOK Ok

swagger:response getAvatarForTokenIdOK
*/
type GetAvatarForTokenIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.Avatar `json:"body,omitempty"`
}

// NewGetAvatarForTokenIDOK creates GetAvatarForTokenIDOK with default headers values
func NewGetAvatarForTokenIDOK() *GetAvatarForTokenIDOK {

	return &GetAvatarForTokenIDOK{}
}

// WithPayload adds the payload to the get avatar for token Id o k response
func (o *GetAvatarForTokenIDOK) WithPayload(payload *models.Avatar) *GetAvatarForTokenIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get avatar for token Id o k response
func (o *GetAvatarForTokenIDOK) SetPayload(payload *models.Avatar) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAvatarForTokenIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAvatarForTokenIDBadRequestCode is the HTTP code returned for type GetAvatarForTokenIDBadRequest
const GetAvatarForTokenIDBadRequestCode int = 400

/*GetAvatarForTokenIDBadRequest Invalid request

swagger:response getAvatarForTokenIdBadRequest
*/
type GetAvatarForTokenIDBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewGetAvatarForTokenIDBadRequest creates GetAvatarForTokenIDBadRequest with default headers values
func NewGetAvatarForTokenIDBadRequest() *GetAvatarForTokenIDBadRequest {

	return &GetAvatarForTokenIDBadRequest{}
}

// WithPayload adds the payload to the get avatar for token Id bad request response
func (o *GetAvatarForTokenIDBadRequest) WithPayload(payload *models.APIError) *GetAvatarForTokenIDBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get avatar for token Id bad request response
func (o *GetAvatarForTokenIDBadRequest) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAvatarForTokenIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAvatarForTokenIDNotFoundCode is the HTTP code returned for type GetAvatarForTokenIDNotFound
const GetAvatarForTokenIDNotFoundCode int = 404

/*GetAvatarForTokenIDNotFound Not found

swagger:response getAvatarForTokenIdNotFound
*/
type GetAvatarForTokenIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewGetAvatarForTokenIDNotFound creates GetAvatarForTokenIDNotFound with default headers values
func NewGetAvatarForTokenIDNotFound() *GetAvatarForTokenIDNotFound {

	return &GetAvatarForTokenIDNotFound{}
}

// WithPayload adds the payload to the get avatar for token Id not found response
func (o *GetAvatarForTokenIDNotFound) WithPayload(payload *models.APIError) *GetAvatarForTokenIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get avatar for token Id not found response
func (o *GetAvatarForTokenIDNotFound) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAvatarForTokenIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAvatarForTokenIDInternalServerErrorCode is the HTTP code returned for type GetAvatarForTokenIDInternalServerError
const GetAvatarForTokenIDInternalServerErrorCode int = 500

/*GetAvatarForTokenIDInternalServerError Internal Server Error

swagger:response getAvatarForTokenIdInternalServerError
*/
type GetAvatarForTokenIDInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewGetAvatarForTokenIDInternalServerError creates GetAvatarForTokenIDInternalServerError with default headers values
func NewGetAvatarForTokenIDInternalServerError() *GetAvatarForTokenIDInternalServerError {

	return &GetAvatarForTokenIDInternalServerError{}
}

// WithPayload adds the payload to the get avatar for token Id internal server error response
func (o *GetAvatarForTokenIDInternalServerError) WithPayload(payload *models.APIError) *GetAvatarForTokenIDInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get avatar for token Id internal server error response
func (o *GetAvatarForTokenIDInternalServerError) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAvatarForTokenIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
