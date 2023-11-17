package gineng

import "github.com/dinozor-io/interfaces"

type Router struct {
	routes []interfaces.Route
}

func (r *Router) AddRoute(method int8, path string, callback func(interfaces.Controller)) {
	var route Route
	(&route).Init(method, path, callback)
	r.routes = append(r.routes, &route)
}

func (r *Router) Routes() []interfaces.Route {
	return r.routes
}
