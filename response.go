package gineng

import (
	"github.com/dinozor-io/consts/resptype"
	"github.com/dinozor-io/interfaces"
)

type Response struct {
	respType   int
	contr      interfaces.Controller
	statusCode int
	data       map[string]any
	path       string
}

func (r *Response) Init(contr interfaces.Controller) {
	r.contr = contr
}

func (r *Response) Type() int {
	return r.respType
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
	r.respType = resptype.JSON
	r.statusCode = statusCode
	r.data = data
}

func (r *Response) HTML(statusCode int, path string) {
	r.respType = resptype.HTML
	r.statusCode = statusCode
	r.path = path
}
