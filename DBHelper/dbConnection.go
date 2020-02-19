package DBHelper

import (
	"baidu/AESHelper"
	"baidu/DataEntity"
	_ "baidu/DataEntity"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

type DBConnectParam struct {
	UserName string
	DBServer string
	PWD      string
	DBName   string
}

func NewInstance() *gorm.DB {
	config := viper.New()
	config.SetConfigName("config")
	config.SetConfigType("yaml")
	config.AddConfigPath("./")
	if err := config.ReadInConfig(); err != nil {
		fmt.Println(err.Error())
	}

	var configData DBConnectParam
	if err := config.Unmarshal(&configData); err != nil {
		fmt.Println(err.Error())
	}
	decryptedPwd := AESHelper.AESDecrypt(configData.PWD)
	decryptedPwd = "123456789123456789"
	var connectParam = fmt.Sprintf(
		//		"#{configData.UserName}:#{decryptedPwd}@#{configData.DBServer}/#{configData.DBName}",
		"%v:%v@tcp(%v)/%v?charset=utf8&parseTime=True&loc=Local", configData.UserName, decryptedPwd, configData.DBServer, configData.DBName)
	db, err := gorm.Open("mysql", connectParam)
	db.AutoMigrate(&DataEntity.Tbl_User{})
	db.SingularTable(true)
	if err != nil {
		fmt.Println(err.Error())
	}
	return db
}
