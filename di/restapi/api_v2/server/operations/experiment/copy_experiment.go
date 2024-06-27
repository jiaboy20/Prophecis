// Code generated by go-swagger; DO NOT EDIT.

package experiment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// CopyExperimentHandlerFunc turns a function with the right signature into a copy experiment handler
type CopyExperimentHandlerFunc func(CopyExperimentParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn CopyExperimentHandlerFunc) Handle(params CopyExperimentParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// CopyExperimentHandler interface for that can handle valid copy experiment params
type CopyExperimentHandler interface {
	Handle(CopyExperimentParams, interface{}) middleware.Responder
}

// NewCopyExperiment creates a new http.Handler for the copy experiment operation
func NewCopyExperiment(ctx *middleware.Context, handler CopyExperimentHandler) *CopyExperiment {
	return &CopyExperiment{Context: ctx, Handler: handler}
}

/*CopyExperiment swagger:route POST /di/v2/experiment/{exp_id}/copy Experiment copyExperiment

CopyExperiment

复制一个已有的实验来创建一个新的实验

*/
type CopyExperiment struct {
	Context *middleware.Context
	Handler CopyExperimentHandler
}

func (o *CopyExperiment) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCopyExperimentParams()

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
