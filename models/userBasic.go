package models

import (
	"ginchat/utils"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name          string `form:"name" json:"Name" gorm:"password,omitempty" binding:"required"`
	Password      string `form:"password" json:"Password" gorm:"password,omitempty" binding:"required"`
	Phone         string `form:"phone" json:"Phone" gorm:"phone,omitempty"`
	Email         string
	Identify      string
	ClientIp      string
	ClientPort    string
	Salt          string
	LoginTime     uint64
	HeartBeatTime uint64
	LoginOutTime  uint64
	IsLogout      bool
	DeviceInfo    string
	Avatar        string
}

func (table *UserBasic) TableName() string {
	return "userBasic"
}

type UserResult struct {
	ID    uint   `json:"userId"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func GetUserList(p *utils.Pagination) []*UserResult {

	data := make([]*UserResult, 10)
	utils.Db.Model(&UserBasic{}).Scopes(p.GormPaginate()).Find(&data)
	var total int64
	utils.Db.Model(&UserBasic{}).Count(&total)
	p.Total = cast.ToInt(total)
	return data
}
func CreateUser(user UserBasic) *gorm.DB {

	return utils.Db.Create(&user)
}
func UpdateUser(user UserBasic) *gorm.DB {

	return utils.Db.Model(&user).Updates(UserBasic{Name: user.Name, Avatar: user.Avatar})
}
func FindUserByNameAndPwd(name, password string) UserBasic {
	user := UserBasic{}
	utils.Db.Where("name = ? and password = ?", name, password).First(&user)
	return user
}
func FindUserByName(name string) UserBasic {
	user := UserBasic{}
	utils.Db.Where("name = ?", name).First(&user)
	return user
}

func FindUserByPhone(phone string) UserBasic {
	user := UserBasic{}
	utils.Db.Where("phone = ?", phone).First(&user)
	return user
}
func FindUserByEmail(email string) UserBasic {
	user := UserBasic{}
	utils.Db.Where("email = ?", email).First(&user)
	return user
}

func FindeUserById(userId uint) UserBasic {
	user := UserBasic{}
	utils.Db.Find(&user, userId)
	return user
}
