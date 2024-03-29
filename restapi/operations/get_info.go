// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetInfoHandlerFunc turns a function with the right signature into a get info handler
type GetInfoHandlerFunc func(GetInfoParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetInfoHandlerFunc) Handle(params GetInfoParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetInfoHandler interface for that can handle valid get info params
type GetInfoHandler interface {
	Handle(GetInfoParams, interface{}) middleware.Responder
}

// NewGetInfo creates a new http.Handler for the get info operation
func NewGetInfo(ctx *middleware.Context, handler GetInfoHandler) *GetInfo {
	return &GetInfo{Context: ctx, Handler: handler}
}

/*GetInfo swagger:route GET /userdata/v1/info getInfo

Information

*/
type GetInfo struct {
	Context *middleware.Context
	Handler GetInfoHandler
}

func (o *GetInfo) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetInfoParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
