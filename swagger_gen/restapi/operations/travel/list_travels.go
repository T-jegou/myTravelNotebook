// Code generated by go-swagger; DO NOT EDIT.

package travel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListTravelsHandlerFunc turns a function with the right signature into a list travels handler
type ListTravelsHandlerFunc func(ListTravelsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListTravelsHandlerFunc) Handle(params ListTravelsParams) middleware.Responder {
	return fn(params)
}

// ListTravelsHandler interface for that can handle valid list travels params
type ListTravelsHandler interface {
	Handle(ListTravelsParams) middleware.Responder
}

// NewListTravels creates a new http.Handler for the list travels operation
func NewListTravels(ctx *middleware.Context, handler ListTravelsHandler) *ListTravels {
	return &ListTravels{Context: ctx, Handler: handler}
}

/* ListTravels swagger:route GET /travels travel listTravels

ListTravels list travels API

*/
type ListTravels struct {
	Context *middleware.Context
	Handler ListTravelsHandler
}

func (o *ListTravels) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewListTravelsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}