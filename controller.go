package gineng

import (
	"github.com/dinozor-io/interfaces"
)

type Controller struct {
	req interfaces.Request
	res interfaces.Response
}

func (c *Controller) Init(req interfaces.Request, res interfaces.Response) {
	c.req = req
	c.res = res
}

func (c *Controller) Req() interfaces.Request {
	return c.req
}

func (c *Controller) Res() interfaces.Response {
	return c.res
}
