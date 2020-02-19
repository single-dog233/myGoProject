package handler

import (
	"baidu/DataAccess"
	"baidu/DataEntity"
)

func GetUserInfo(userId string) *DataEntity.Tbl_User {
	user := DataAccess.GetUserInfo(userId)
	return user
}
