// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/stephram/nsm-userdata-service/restapi/models"
)

// NewPostGameOnResultsParams creates a new PostGameOnResultsParams object
// no default values defined in spec.
func NewPostGameOnResultsParams() PostGameOnResultsParams {

	return PostGameOnResultsParams{}
}

// PostGameOnResultsParams contains all the bound params for the post game on results operation
// typically these are obtained from a http.Request
//
// swagger:parameters postGameOnResults
type PostGameOnResultsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: header
	*/
	ExhibitID *string
	/*
	  Required: true
	  In: body
	*/
	GameOnResults *models.GameOnResults
	/*
	  Required: true
	  In: path
	*/
	TokenID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostGameOnResultsParams() beforehand.
func (o *PostGameOnResultsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := o.bindExhibitID(r.Header[http.CanonicalHeaderKey("exhibitId")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.GameOnResults
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("gameOnResults", "body"))
			} else {
				res = append(res, errors.NewParseError("gameOnResults", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.GameOnResults = &body
			}
		}
	} else {
		res = append(res, errors.Required("gameOnResults", "body"))
	}
	rTokenID, rhkTokenID, _ := route.Params.GetOK("tokenId")
	if err := o.bindTokenID(rTokenID, rhkTokenID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindExhibitID binds and validates parameter ExhibitID from header.
func (o *PostGameOnResultsParams) bindExhibitID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.ExhibitID = &raw

	return nil
}

// bindTokenID binds and validates parameter TokenID from path.
func (o *PostGameOnResultsParams) bindTokenID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.TokenID = raw

	return nil
}
