package model

import (
	"fmt"
	"github.com/mangenotwork/extras/apps/IM-Msg/global"
	"github.com/mangenotwork/extras/common/utils"
)

type Message struct {
	ID int64 `gorm:"primary_key;column:id" json:"id"`
	SendType int `gorm:"column:send_type" json:"send_type"` // 发送类型 对应 SendType 枚举
	FromUid string `gorm:"column:from_uid" json:"from_uid"` // 发送人id
	FromName string `gorm:"column:from_name" json:"from_name"` // 发送人昵称
	FromAvatar string `gorm:"column:from_avatar" json:"from_avatar"` // 发送人头像
	ToUid string `gorm:"column:to_uid" json:"to_uid"` // 接收人id
	ToName string `gorm:"column:to_name" json:"to_name"` // 接收人昵称
	ToAvatar string `gorm:"column:to_avatar" json:"to_avatar"` // 接收人头像
	ToGroupId string `gorm:"column:to_group_id" json:"to_group_id"` // 群id
	ToGroupName string `gorm:"column:to_group_name" json:"to_group_name"` // 群昵称
	ToGroupAvatar string `gorm:"column:to_group_avatar" json:"to_group_avatar"` // 群头像
	ContentType int `gorm:"column:content_type" json:"content_type"` // 消息类型 对应 MsgType
	Content []byte `gorm:"column:content" json:"content"` // 消息内容
	SendTime int64 `gorm:"column:send_time" json:"send_time"` // 消息时间
	Device string `gorm:"column:device" json:"device"` // 发送端的 设备Code
	Source string `gorm:"column:source" json:"source"` // 发送端的 设备类型
}

// hashFunc
func hashFunc(fromUid, toUid, toGroupId string) int64 {
	fromUidInt := utils.Str2Int64(fromUid)
	toInt := utils.Str2Int64(toUid)
	if toInt == 0 {
		toInt = utils.Str2Int64(toGroupId)
	}
	if fromUidInt < 0 {
		fromUidInt = -fromUidInt
	}
	if toInt < 0 {
		toInt = -toInt
	}
	sid := fromUidInt % global.MaxMsgTable
	rid := toInt % global.MaxMsgTable
	id := sid - rid
	if id < 0 {
		id = -id
	}
	return id
}

// TableName
func (m *Message) TableName(fromUid, toUid, toGroupId string) string {
	return fmt.Sprintf(global.MsgTableName, hashFunc(fromUid, toUid, toGroupId))
}

// SendType
const (
	SendType_OneToOne = 1 	// 一对一聊天
	SendType_Group = 2		// 群聊
	SendType_SYS = 3		// 系统消息
	SendType_Inform = 4 	// 通知消息

)

// MsgType
const (
	MsgType_TEXT      = 0	// 文字消息
	MsgType_IMAGE     = 1	// 图片消息
	MsgType_AUDIO     = 2	// 音频消息
	MsgType_VIDEO     = 3	// 视频消息
	MsgType_FILE      = 4	// 文件消息
	MsgType_LOCATION  = 5	// 位置消息
	MsgType_COMMAND   = 6	// 指令消息
	MsgType_Custom    = 7	// 自定义消息
)

// 发送者结构体
type FromUser struct {
	Uid string
	Name string
	Avatar string
}

// 接收者结构体
type ToUser struct {
	Uid string
	Name string
	Avatar string
}

// 接收群信息
type ToGroup struct {
	Id int //群id
	Name string
	Avatar string
}

// 发送 one to one 消息
func (msg *Message) SendOneToOne(from *FromUser, to *ToUser){

}

