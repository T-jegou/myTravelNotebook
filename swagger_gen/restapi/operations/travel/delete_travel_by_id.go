// Code generated by go-swagger; DO NOT EDIT.

package travel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteTravelByIDHandlerFunc turns a function with the right signature into a delete travel by Id handler
type DeleteTravelByIDHandlerFunc func(DeleteTravelByIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteTravelByIDHandlerFunc) Handle(params DeleteTravelByIDParams) middleware.Responder {
	return fn(params)
}

// DeleteTravelByIDHandler interface for that can handle valid delete travel by Id params
type DeleteTravelByIDHandler interface {
	Handle(DeleteTravelByIDParams) middleware.Responder
}

// NewDeleteTravelByID creates a new http.Handler for the delete travel by Id operation
func NewDeleteTravelByID(ctx *middleware.Context, handler DeleteTravelByIDHandler) *DeleteTravelByID {
	return &DeleteTravelByID{Context: ctx, Handler: handler}
}

/* DeleteTravelByID swagger:route DELETE /travel/{id} travel deleteTravelById

Delete an existing travel using his Id

*/
type DeleteTravelByID struct {
	Context *middleware.Context
	Handler DeleteTravelByIDHandler
}

func (o *DeleteTravelByID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteTravelByIDParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
