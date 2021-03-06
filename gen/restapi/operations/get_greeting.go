// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetGreetingHandlerFunc turns a function with the right signature into a get greeting handler
type GetGreetingHandlerFunc func(GetGreetingParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetGreetingHandlerFunc) Handle(params GetGreetingParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetGreetingHandler interface for that can handle valid get greeting params
type GetGreetingHandler interface {
	Handle(GetGreetingParams, interface{}) middleware.Responder
}

// NewGetGreeting creates a new http.Handler for the get greeting operation
func NewGetGreeting(ctx *middleware.Context, handler GetGreetingHandler) *GetGreeting {
	return &GetGreeting{Context: ctx, Handler: handler}
}

/*GetGreeting swagger:route GET /hello getGreeting

GetGreeting get greeting API

*/
type GetGreeting struct {
	Context *middleware.Context
	Handler GetGreetingHandler
}

func (o *GetGreeting) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetGreetingParams()

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
