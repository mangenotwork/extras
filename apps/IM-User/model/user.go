package model

import (
	"fmt"
	"github.com/mangenotwork/extras/apps/IM-User/global"
)

type UserBase struct {
	ID int `gorm:"primary_key;column:id;size:11" json:"id"` // 表自增id
	TableId int `gorm:"column:table_id" json:"table_id"`
	UId string `gorm:"column:uid" json:"uid"` // 用户id  拆表id(前4位)+表自增id
	UName string `gorm:"column:uname" json:"uname"` // 用户昵称
	Account string `gorm:"column:account" json:"account"` // 用户账号
	Password string `gorm:"column:password" json:"password"` // 用户密码
	Sex int `gorm:"column:sex" json:"sex"` // 性别
	Birthday int `gorm:"column:birthday" json:"birthday"` // 生日
	Site string `gorm:"column:site" json:"site"` // 所在地
	Mail string `gorm:"column:mail" json:"mail"` // 邮件
	Phone string `gorm:"column:phone" json:"phone"` // 电话
	Avatar string `gorm:"column:avatar" json:"avatar"`                         // 头像
	Updated int64  `gorm:"column:updated" json:"updated"` // 更新
	Created int64  `gorm:"column:created" json:"created"` // 入库时间
	IsDeleted int64  `gorm:"column:deleted" json:"is_deleted"` // 删除字段
}

func (u *UserBase) TableName() string {
	return fmt.Sprintf(global.UserBaseTableName, u.TableId)
}
