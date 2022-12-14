// Code generated by go-swagger; DO NOT EDIT.

package travel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/T-jegou/myTravelNotebook/swagger_gen/models"
)

// GetTravelsHandlerFunc turns a function with the right signature into a get travels handler
type GetTravelsHandlerFunc func(GetTravelsParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetTravelsHandlerFunc) Handle(params GetTravelsParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetTravelsHandler interface for that can handle valid get travels params
type GetTravelsHandler interface {
	Handle(GetTravelsParams, *models.Principal) middleware.Responder
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
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
