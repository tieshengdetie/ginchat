package models

import (
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	FromId   uint64 `json:"userId"` //发送者
	TargetId uint64 //接受者
	Type     int    //发送类型  1 私聊 2 群聊 3 广播
	Media    int    //媒体类型 1 文字 2 表情包 3 图片 4 音频
	Content  string //消息内容
	Picture  string
	Url      string
	Desc     string
	Amount   int //其他的统计数字
}

func (table *Message) TableName() string {
	return "message"
}
