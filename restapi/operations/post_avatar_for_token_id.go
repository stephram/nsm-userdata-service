// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostAvatarForTokenIDHandlerFunc turns a function with the right signature into a post avatar for token ID handler
type PostAvatarForTokenIDHandlerFunc func(PostAvatarForTokenIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostAvatarForTokenIDHandlerFunc) Handle(params PostAvatarForTokenIDParams) middleware.Responder {
	return fn(params)
}

// PostAvatarForTokenIDHandler interface for that can handle valid post avatar for token ID params
type PostAvatarForTokenIDHandler interface {
	Handle(PostAvatarForTokenIDParams) middleware.Responder
}

// NewPostAvatarForTokenID creates a new http.Handler for the post avatar for token ID operation
func NewPostAvatarForTokenID(ctx *middleware.Context, handler PostAvatarForTokenIDHandler) *PostAvatarForTokenID {
	return &PostAvatarForTokenID{Context: ctx, Handler: handler}
}

/*PostAvatarForTokenID swagger:route POST /userdata/v1/avatar/{tokenID} postAvatarForTokenId

Store the Avatar associated with the tokenID

*/
type PostAvatarForTokenID struct {
	Context *middleware.Context
	Handler PostAvatarForTokenIDHandler
}

func (o *PostAvatarForTokenID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostAvatarForTokenIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
