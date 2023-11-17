package gineng

import "github.com/dinozor-io/interfaces"

type Route struct {
	method   int8
	path     string
	callback func(interfaces.Controller)
}

func (r *Route) Init(method int8, path string, callback func(interfaces.Controller)) {
	r.method = method
	r.path = path
	r.callback = callback
}

func (r *Route) Method() int8 {
	return r.method
}

func (r *Route) Path() string {
	return r.path
}

func (r *Route) Callback() func(interfaces.Controller) {
	return r.callback
}
