// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/stephram/nsm-userdata-service/restapi/models"
)

// GetInteractionsOKCode is the HTTP code returned for type GetInteractionsOK
const GetInteractionsOKCode int = 200

/*GetInteractionsOK Ok

swagger:response getInteractionsOK
*/
type GetInteractionsOK struct {

	/*
	  In: Body
	*/
	Payload models.ConsoleInteractions `json:"body,omitempty"`
}

// NewGetInteractionsOK creates GetInteractionsOK with default headers values
func NewGetInteractionsOK() *GetInteractionsOK {

	return &GetInteractionsOK{}
}

// WithPayload adds the payload to the get interactions o k response
func (o *GetInteractionsOK) WithPayload(payload models.ConsoleInteractions) *GetInteractionsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get interactions o k response
func (o *GetInteractionsOK) SetPayload(payload models.ConsoleInteractions) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetInteractionsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.ConsoleInteractions{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetInteractionsBadRequestCode is the HTTP code returned for type GetInteractionsBadRequest
const GetInteractionsBadRequestCode int = 400

/*GetInteractionsBadRequest Invalid request

swagger:response getInteractionsBadRequest
*/
type GetInteractionsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewGetInteractionsBadRequest creates GetInteractionsBadRequest with default headers values
func NewGetInteractionsBadRequest() *GetInteractionsBadRequest {

	return &GetInteractionsBadRequest{}
}

// WithPayload adds the payload to the get interactions bad request response
func (o *GetInteractionsBadRequest) WithPayload(payload *models.APIError) *GetInteractionsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get interactions bad request response
func (o *GetInteractionsBadRequest) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetInteractionsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetInteractionsNotFoundCode is the HTTP code returned for type GetInteractionsNotFound
const GetInteractionsNotFoundCode int = 404

/*GetInteractionsNotFound Not found

swagger:response getInteractionsNotFound
*/
type GetInteractionsNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewGetInteractionsNotFound creates GetInteractionsNotFound with default headers values
func NewGetInteractionsNotFound() *GetInteractionsNotFound {

	return &GetInteractionsNotFound{}
}

// WithPayload adds the payload to the get interactions not found response
func (o *GetInteractionsNotFound) WithPayload(payload *models.APIError) *GetInteractionsNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get interactions not found response
func (o *GetInteractionsNotFound) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetInteractionsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetInteractionsInternalServerErrorCode is the HTTP code returned for type GetInteractionsInternalServerError
const GetInteractionsInternalServerErrorCode int = 500

/*GetInteractionsInternalServerError Internal server error

swagger:response getInteractionsInternalServerError
*/
type GetInteractionsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewGetInteractionsInternalServerError creates GetInteractionsInternalServerError with default headers values
func NewGetInteractionsInternalServerError() *GetInteractionsInternalServerError {

	return &GetInteractionsInternalServerError{}
}

// WithPayload adds the payload to the get interactions internal server error response
func (o *GetInteractionsInternalServerError) WithPayload(payload *models.APIError) *GetInteractionsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get interactions internal server error response
func (o *GetInteractionsInternalServerError) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetInteractionsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
