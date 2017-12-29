package main

import "github.com/kataras/iris"

func main() {
	app := iris.New()

	app.RegisterView(iris.HTML("./pages", ".html").Reload(true))
	app.StaticWeb("/statics", "./statics")

	app.Get("/", func(c iris.Context) {
		c.View("index.html")
	})
	app.Post("/eval_photo", func(c iris.Context) {

	})

	app.Run(iris.Addr(":3600"))
}
