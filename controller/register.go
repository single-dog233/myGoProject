package controller

import "github.com/kataras/iris/v12"

func Register(app *iris.Application) {
	userController := app.Party("/user")
	{
		userController.Get("/getUserInfo", GetUerInfo)
	}
}
