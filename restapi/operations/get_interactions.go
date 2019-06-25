// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetInteractionsHandlerFunc turns a function with the right signature into a get interactions handler
type GetInteractionsHandlerFunc func(GetInteractionsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetInteractionsHandlerFunc) Handle(params GetInteractionsParams) middleware.Responder {
	return fn(params)
}

// GetInteractionsHandler interface for that can handle valid get interactions params
type GetInteractionsHandler interface {
	Handle(GetInteractionsParams) middleware.Responder
}

// NewGetInteractions creates a new http.Handler for the get interactions operation
func NewGetInteractions(ctx *middleware.Context, handler GetInteractionsHandler) *GetInteractions {
	return &GetInteractions{Context: ctx, Handler: handler}
}

/*GetInteractions swagger:route GET /userdata/v1/interactions/{exhibitID}/{tokenID} getInteractions

Get the interaction data for the exhibitID and tokenID

*/
type GetInteractions struct {
	Context *middleware.Context
	Handler GetInteractionsHandler
}

func (o *GetInteractions) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetInteractionsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}