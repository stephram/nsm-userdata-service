// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/stephram/nsm-userdata-service/restapi/models"
)

// GetGameOnResultsOKCode is the HTTP code returned for type GetGameOnResultsOK
const GetGameOnResultsOKCode int = 200

/*GetGameOnResultsOK Ok

swagger:response getGameOnResultsOK
*/
type GetGameOnResultsOK struct {

	/*
	  In: Body
	*/
	Payload *models.GameOnResults `json:"body,omitempty"`
}

// NewGetGameOnResultsOK creates GetGameOnResultsOK with default headers values
func NewGetGameOnResultsOK() *GetGameOnResultsOK {

	return &GetGameOnResultsOK{}
}

// WithPayload adds the payload to the get game on results o k response
func (o *GetGameOnResultsOK) WithPayload(payload *models.GameOnResults) *GetGameOnResultsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get game on results o k response
func (o *GetGameOnResultsOK) SetPayload(payload *models.GameOnResults) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetGameOnResultsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetGameOnResultsBadRequestCode is the HTTP code returned for type GetGameOnResultsBadRequest
const GetGameOnResultsBadRequestCode int = 400

/*GetGameOnResultsBadRequest Invalid request

swagger:response getGameOnResultsBadRequest
*/
type GetGameOnResultsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewGetGameOnResultsBadRequest creates GetGameOnResultsBadRequest with default headers values
func NewGetGameOnResultsBadRequest() *GetGameOnResultsBadRequest {

	return &GetGameOnResultsBadRequest{}
}

// WithPayload adds the payload to the get game on results bad request response
func (o *GetGameOnResultsBadRequest) WithPayload(payload *models.APIError) *GetGameOnResultsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get game on results bad request response
func (o *GetGameOnResultsBadRequest) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetGameOnResultsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetGameOnResultsNotFoundCode is the HTTP code returned for type GetGameOnResultsNotFound
const GetGameOnResultsNotFoundCode int = 404

/*GetGameOnResultsNotFound Not found

swagger:response getGameOnResultsNotFound
*/
type GetGameOnResultsNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewGetGameOnResultsNotFound creates GetGameOnResultsNotFound with default headers values
func NewGetGameOnResultsNotFound() *GetGameOnResultsNotFound {

	return &GetGameOnResultsNotFound{}
}

// WithPayload adds the payload to the get game on results not found response
func (o *GetGameOnResultsNotFound) WithPayload(payload *models.APIError) *GetGameOnResultsNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get game on results not found response
func (o *GetGameOnResultsNotFound) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetGameOnResultsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetGameOnResultsInternalServerErrorCode is the HTTP code returned for type GetGameOnResultsInternalServerError
const GetGameOnResultsInternalServerErrorCode int = 500

/*GetGameOnResultsInternalServerError Internal server error

swagger:response getGameOnResultsInternalServerError
*/
type GetGameOnResultsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewGetGameOnResultsInternalServerError creates GetGameOnResultsInternalServerError with default headers values
func NewGetGameOnResultsInternalServerError() *GetGameOnResultsInternalServerError {

	return &GetGameOnResultsInternalServerError{}
}

// WithPayload adds the payload to the get game on results internal server error response
func (o *GetGameOnResultsInternalServerError) WithPayload(payload *models.APIError) *GetGameOnResultsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get game on results internal server error response
func (o *GetGameOnResultsInternalServerError) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetGameOnResultsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
