package models

import (
	"ginchat/utils"
	"gorm.io/gorm"
)

const (
	SexWomen  = "W"
	SexMan    = "M"
	SexUnknow = "U"
)

type User struct {
	gorm.Model
	Nickname string `form:"nickname" json:"name nickname" gorm:"varchar(20)" binding:"required"`
	Password string `form:"password" json:"password omitempty" gorm:"password,omitempty" binding:"required"`
	Mobile   string `form:"phone" json:"phone omitempty" gorm:"phone,omitempty"`
	Avatar   string `gorm:"varchar(150)" form:"avatar" json:"avatar"`
	Sex      string `gorm:"varchar(2)" form:"sex" json:"sex"`
	Salt     string `gorm:"varchar(20)" form:"salt" json:"-"`
	Online   int    `gorm:"int(10)" form:"online" json:"online"`
	Token    string `gorm:"varchar(40)" form:"token" json:"token"`
	Memo     string `gorm:"varchar(140)" form:"memo" json:"memo"`
}

func (table *User) TableName() string {
	return "user"
}

func FindUserByMobile(mobile string) User {

	user := User{}
	utils.Db.Where("mobile = ?", mobile).First(&user)
	return user
}
