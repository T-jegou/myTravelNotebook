// Code generated by go-swagger; DO NOT EDIT.

package login

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetLoginHandlerFunc turns a function with the right signature into a get login handler
type GetLoginHandlerFunc func(GetLoginParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetLoginHandlerFunc) Handle(params GetLoginParams) middleware.Responder {
	return fn(params)
}

// GetLoginHandler interface for that can handle valid get login params
type GetLoginHandler interface {
	Handle(GetLoginParams) middleware.Responder
}

// NewGetLogin creates a new http.Handler for the get login operation
func NewGetLogin(ctx *middleware.Context, handler GetLoginHandler) *GetLogin {
	return &GetLogin{Context: ctx, Handler: handler}
}

/* GetLogin swagger:route GET /login login getLogin

login through oauth2 server

*/
type GetLogin struct {
	Context *middleware.Context
	Handler GetLoginHandler
}

func (o *GetLogin) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetLoginParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetLoginOKBody get login o k body
//
// swagger:model GetLoginOKBody
type GetLoginOKBody struct {

	// access token
	AccessToken string `json:"access_token,omitempty"`
}

// Validate validates this get login o k body
func (o *GetLoginOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get login o k body based on context it is used
func (o *GetLoginOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetLoginOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetLoginOKBody) UnmarshalBinary(b []byte) error {
	var res GetLoginOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
