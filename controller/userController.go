package controller

import (
	"baidu/handler"
	context2 "github.com/kataras/iris/v12/context"
)

func GetUerInfo(ctx context2.Context) {
	userId := ctx.URLParam("userId")
	user := handler.GetUserInfo(userId)
	ctx.JSON(user)
}
