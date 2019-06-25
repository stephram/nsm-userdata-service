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

// NewGetInteractionsParams creates a new GetInteractionsParams object
// no default values defined in spec.
func NewGetInteractionsParams() GetInteractionsParams {

	return GetInteractionsParams{}
}

// GetInteractionsParams contains all the bound params for the get interactions operation
// typically these are obtained from a http.Request
//
// swagger:parameters getInteractions
type GetInteractionsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The unique and static ID of the exhibit
	  Required: true
	  In: path
	*/
	ExhibitID string
	/*The unique tokenID of the user
	  Required: true
	  In: path
	*/
	TokenID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetInteractionsParams() beforehand.
func (o *GetInteractionsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rExhibitID, rhkExhibitID, _ := route.Params.GetOK("exhibitID")
	if err := o.bindExhibitID(rExhibitID, rhkExhibitID, route.Formats); err != nil {
		res = append(res, err)
	}

	rTokenID, rhkTokenID, _ := route.Params.GetOK("tokenID")
	if err := o.bindTokenID(rTokenID, rhkTokenID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindExhibitID binds and validates parameter ExhibitID from path.
func (o *GetInteractionsParams) bindExhibitID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ExhibitID = raw

	return nil
}

// bindTokenID binds and validates parameter TokenID from path.
func (o *GetInteractionsParams) bindTokenID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.TokenID = raw

	return nil
}