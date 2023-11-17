package gineng

import "github.com/dinozor-io/interfaces"

type Response struct {
	contr      interfaces.Controller
	statusCode int
	data       map[string]any
}

func (r *Response) Init(contr interfaces.Controller) {
	r.contr = contr
}

func (r *Response) Contr() interfaces.Controller {
	return r.contr
}

func (r *Response) StatusCode() int {
	return r.statusCode
}

func (r *Response) Data() map[string]any {
	return r.data
}

func (r *Response) JSON(statusCode int, data map[string]any) {
	r.statusCode = statusCode
	r.data = data
}
