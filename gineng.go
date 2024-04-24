package gineng

import (
	"net/http"
	"os"

	"github.com/dinozor-io/consts/methods"
	"github.com/dinozor-io/consts/resptype"
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

				if res.Type() == resptype.JSON {
					c.JSON(res.StatusCode(), res.Data())
				} else if res.Type() == resptype.HTML {
					bytes, e := os.ReadFile(res.path)

					if e != nil {
						c.Data(http.StatusNotFound, "text/html", []byte("Not Found"))
						return
					}

					c.Data(res.StatusCode(), "text/html", bytes)
				}
			})
		}
	}

	e.gin.Run(e.srv.Port())
}
