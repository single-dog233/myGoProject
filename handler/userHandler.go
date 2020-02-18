package handler

import "baidu/model"

func GetUserInfo(user *model.User) {
	user.UserId = "123"
	user.UserName = "456"
}
