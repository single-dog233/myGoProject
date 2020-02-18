package controller

import (
	"baidu/handler"
	"baidu/model"
	context2 "github.com/kataras/iris/v12/context"
)

func GetUerInfo(ctx context2.Context) {
	var User model.User
	handler.GetUserInfo(&User)
	ctx.JSON(User)
}
