// Code generated by go-swagger; DO NOT EDIT.

package travel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetTravelsHandlerFunc turns a function with the right signature into a get travels handler
type GetTravelsHandlerFunc func(GetTravelsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetTravelsHandlerFunc) Handle(params GetTravelsParams) middleware.Responder {
	return fn(params)
}

// GetTravelsHandler interface for that can handle valid get travels params
type GetTravelsHandler interface {
	Handle(GetTravelsParams) middleware.Responder
}

// NewGetTravels creates a new http.Handler for the get travels operation
func NewGetTravels(ctx *middleware.Context, handler GetTravelsHandler) *GetTravels {
	return &GetTravels{Context: ctx, Handler: handler}
}

/* GetTravels swagger:route GET /travels travel getTravels

Returns an array of travels object.

*/
type GetTravels struct {
	Context *middleware.Context
	Handler GetTravelsHandler
}

func (o *GetTravels) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetTravelsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
