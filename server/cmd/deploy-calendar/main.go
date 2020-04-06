package main

import (
	"fmt"

	"github.com/cognitivehacker/deploy-calendar/server/internal/api/routes"
	"github.com/valyala/fasthttp"
)

func main() {

	fmt.Println()

	r := routes.GetRoutes()

	if err := fasthttp.ListenAndServe(":8080", r.Handler); err != nil {
		fmt.Println("Error in ListenAndServe: ", err)
	}

}

func exampleRequestHandler(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("hello")
}
