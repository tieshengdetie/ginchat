package models

import (
	"ginchat/utils"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

type Relation struct {
	gorm.Model
	OwnerId  uint //谁的关系信息
	TargetId uint //对应的谁
	Type     uint // 对应的类型  1 好友 2 群  3
	Desc     string
}

func (table *Relation) TableName() string {
	return "relation"
}

func SearcFriends(userId uint, p *utils.Pagination) []UserBasic {

	relation := make([]Relation, 0)
	userIds := make([]uint, 0)
	utils.Db.Model(&Relation{}).Where("owner_id= ? and type =1", userId).Find(&relation)
	for _, user := range relation {
		userIds = append(userIds, user.TargetId)
	}
	userbasic := make([]UserBasic, 0)
	whereCondition := utils.Db.Model(&UserBasic{}).Where("id IN ?", userIds)

	whereCondition.Scopes(p.GormPaginate()).Find(&userbasic)

	var total int64
	whereCondition.Count(&total)
	p.Total = cast.ToInt(total)
	return userbasic
}

func AddFriends(userId uint, targetId uint) bool {

	relation := []Relation{
		{Type: 1, OwnerId: userId, TargetId: targetId},
		{Type: 1, OwnerId: targetId, TargetId: userId},
	}
	err := utils.Db.Create(&relation).Error
	if err != nil {
		return false
	}
	return true
}

func FindRelationByUserIdAndTargetId(userId uint, targetId uint) Relation {

	relation := Relation{}
	utils.Db.Where("owner_id = ? and target_id = ? and type=1", userId, targetId).First(&relation)
	return relation
}
