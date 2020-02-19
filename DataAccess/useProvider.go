package DataAccess

import (
	"baidu/DBHelper"
	"baidu/DataEntity"
)

func GetUserInfo(userId string) *DataEntity.Tbl_User {
	db := DBHelper.NewInstance()
	defer db.Close()

	//插入练习
	//add := DataEntity.Tbl_User{UserId: "cc", UserName: "dd"}
	//db.Create(&add)

	var user DataEntity.Tbl_User
	rows := db.Where("userid=?", userId).Find(&user)
	if rows.Error != nil {
		panic(rows.Error.Error())
	}
	return &user
}
