package main

import "github.com/kataras/iris/v12"
import "baidu/controller"

func main() {
	app := iris.New()
	controller.Register(app)
	app.Run(iris.Addr(":5000"), iris.WithoutServerError(iris.ErrServerClosed))
}
