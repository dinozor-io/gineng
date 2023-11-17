package gineng

import "github.com/dinozor-io/interfaces"

type Route struct {
	method   int8
	path     string
	callback func(interfaces.Controller)
	group    interfaces.Group
}

func (r *Route) Init(method int8, path string, callback func(interfaces.Controller), group interfaces.Group) {
	r.method = method
	r.path = path
	r.callback = callback
	r.group = group
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

func (r *Route) Group() interfaces.Group {
	return r.group
}
