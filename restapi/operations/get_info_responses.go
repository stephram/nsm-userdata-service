// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/stephram/nsm-userdata-service/restapi/models"
)

// GetInfoOKCode is the HTTP code returned for type GetInfoOK
const GetInfoOKCode int = 200

/*GetInfoOK OK

swagger:response getInfoOK
*/
type GetInfoOK struct {

	/*
	  In: Body
	*/
	Payload *models.InfoResponse `json:"body,omitempty"`
}

// NewGetInfoOK creates GetInfoOK with default headers values
func NewGetInfoOK() *GetInfoOK {

	return &GetInfoOK{}
}

// WithPayload adds the payload to the get info o k response
func (o *GetInfoOK) WithPayload(payload *models.InfoResponse) *GetInfoOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get info o k response
func (o *GetInfoOK) SetPayload(payload *models.InfoResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetInfoOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetInfoInternalServerErrorCode is the HTTP code returned for type GetInfoInternalServerError
const GetInfoInternalServerErrorCode int = 500

/*GetInfoInternalServerError Internal server error

swagger:response getInfoInternalServerError
*/
type GetInfoInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewGetInfoInternalServerError creates GetInfoInternalServerError with default headers values
func NewGetInfoInternalServerError() *GetInfoInternalServerError {

	return &GetInfoInternalServerError{}
}

// WithPayload adds the payload to the get info internal server error response
func (o *GetInfoInternalServerError) WithPayload(payload *models.APIError) *GetInfoInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get info internal server error response
func (o *GetInfoInternalServerError) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetInfoInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
