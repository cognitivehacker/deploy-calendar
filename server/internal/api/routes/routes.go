package routes

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

//GetRoutes return routes
func GetRoutes() *router.Router {

	r := router.New()
	r.GET("/", exampleRequestHandler)

	return r

}

func exampleRequestHandler(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("hello")
}
