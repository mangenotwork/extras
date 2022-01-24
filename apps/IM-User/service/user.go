package service

import (
	"github.com/mangenotwork/extras/common/conf"
	"github.com/mangenotwork/extras/common/logger"
)

type UserAPI interface {
	Register()
}

type UserParam struct {
	Name string `json:"name"` // 昵称
	Account string `json:"account"` // 账号
	Password string `json:"password"`
}

func (param *UserParam) Register() {

	logger.Debug(conf.Arg.Mysql)

	for _, v := range conf.Arg.Mysql {
		logger.Debug(v)
	}

	// 昵称是否重复
	// 账号是否重复

	// 创建账号
}