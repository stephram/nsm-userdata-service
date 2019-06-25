// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetAvatarForTokenIDParams creates a new GetAvatarForTokenIDParams object
// no default values defined in spec.
func NewGetAvatarForTokenIDParams() GetAvatarForTokenIDParams {

	return GetAvatarForTokenIDParams{}
}

// GetAvatarForTokenIDParams contains all the bound params for the get avatar for token ID operation
// typically these are obtained from a http.Request
//
// swagger:parameters getAvatarForTokenID
type GetAvatarForTokenIDParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	TokenID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetAvatarForTokenIDParams() beforehand.
func (o *GetAvatarForTokenIDParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rTokenID, rhkTokenID, _ := route.Params.GetOK("tokenID")
	if err := o.bindTokenID(rTokenID, rhkTokenID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindTokenID binds and validates parameter TokenID from path.
func (o *GetAvatarForTokenIDParams) bindTokenID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.TokenID = raw

	return nil
}
