// Code generated by go-swagger; DO NOT EDIT.

package experiments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ListExperimentsHandlerFunc turns a function with the right signature into a list experiments handler
type ListExperimentsHandlerFunc func(ListExperimentsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ListExperimentsHandlerFunc) Handle(params ListExperimentsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ListExperimentsHandler interface for that can handle valid list experiments params
type ListExperimentsHandler interface {
	Handle(ListExperimentsParams, interface{}) middleware.Responder
}

// NewListExperiments creates a new http.Handler for the list experiments operation
func NewListExperiments(ctx *middleware.Context, handler ListExperimentsHandler) *ListExperiments {
	return &ListExperiments{Context: ctx, Handler: handler}
}

/*ListExperiments swagger:route GET /di/v1/experiments Experiments listExperiments

List Experiments by User.

List Experiments by User.

*/
type ListExperiments struct {
	Context *middleware.Context
	Handler ListExperimentsHandler
}

func (o *ListExperiments) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListExperimentsParams()

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
