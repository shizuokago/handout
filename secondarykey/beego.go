package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func main() {
	beego.Get("/:user", func(ctx *context.Context) {
		ctx.Output.Body([]byte("Hello Beego..." + ctx.Input.Param(":user")))
	})
	beego.Run()
}
