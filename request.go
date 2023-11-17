package gineng

import "github.com/dinozor-io/interfaces"

type Request struct {
	contr interfaces.Controller
}

func (r *Request) Init(contr interfaces.Controller) {
	r.contr = contr
}

func (r *Request) Contr() interfaces.Controller {
	return r.contr
}
