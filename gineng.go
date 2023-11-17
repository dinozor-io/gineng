package gineng

import (
	"net/http"

	"github.com/dinozor-io/consts/methods"
	"github.com/dinozor-io/interfaces"
	"github.com/gin-gonic/gin"
)

type Engine struct {
	gin *gin.Engine
	srv interfaces.Server
}

func (e *Engine) Init(srv interfaces.Server) {
	e.srv = srv
}

func (e *Engine) Run() {
	e.gin = gin.Default()

	routes := e.srv.Router().Routes()

	for i := 0; i < len(routes); i++ {

		callback := routes[i].Callback()
		group := routes[i].Group()

		if methods.GET == routes[i].Method() {
			e.gin.GET(routes[i].Path(), func(c *gin.Context) {

				contr := new(Controller)
				req := new(Request)
				res := new(Response)

				req.Init(contr)
				res.Init(contr)

				contr.Init(req, res)

				if !group.CheckCond(contr) {
					c.JSON(http.StatusForbidden, map[string]any{})
					return
				}

				callback(contr)

				c.JSON(res.StatusCode(), res.Data())
			})
		}
	}

	e.gin.Run(e.srv.Port())
}
