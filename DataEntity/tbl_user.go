package DataEntity

type Tbl_User struct {
	UserId   string `gorm:"column:userId"`
	UserName string `gorm:"column:userName"`
}
