package main

import (
	"wolfgo/base"
	"wolfgo/business"

	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
)

func main() {
	//xorm database init
	base.InitDB()

	app := iris.New()

	app.Adapt(iris.DevLogger())
	app.Adapt(httprouter.New())

	app.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
		ctx.HTML(iris.StatusNotFound, "<h1>Custom not found handler </h1>")
	})

	app.Post("/wolfgo/activity/add", business.AddActivity)

	app.Listen(":8090")
}
